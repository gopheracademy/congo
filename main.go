package main

import (
	"flag"
	"fmt"
	stdlog "log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"reflect"
	"syscall"

	"github.com/streadway/handy/cors"
	"github.com/streadway/handy/encoding"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/go-kit/kit/addsvc/pb"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/tracing/zipkin"
	jsoncodec "github.com/go-kit/kit/transport/codec/json"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	var (
		debugAddr = flag.String("debug.addr", ":8000", "Address for HTTP debug/instrumentation server")
		httpAddr  = flag.String("http.addr", ":8001", "Address for HTTP (JSON) server")
		grpcAddr  = flag.String("grpc.addr", ":8002", "Address for gRPC server")
	)
	flag.Parse()

	// `package log` domain
	var logger kitlog.Logger
	logger = kitlog.NewPrefixLogger(os.Stderr)
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestampUTC)
	kitlog.DefaultLogger = logger                     // for other gokit components
	stdlog.SetOutput(kitlog.NewStdlibAdapter(logger)) // redirect stdlib logging to us
	stdlog.SetFlags(0)                                // flags are handled in our logger

	// Mechanical stuff
	root := context.Background()
	errc := make(chan error)

	go func() {
		errc <- interrupt()
	}()

	// Transport: HTTP (debug/instrumentation)
	go func() {
		logger.Log("addr", *debugAddr, "transport", "debug")
		errc <- http.ListenAndServe(*debugAddr, nil)
	}()

	// Transport: HTTP (JSON)
	go func() {
		ctx, cancel := context.WithCancel(root)
		defer cancel()

		field := metrics.Field{Key: "transport", Value: "http"}
		before := httptransport.Before(zipkin.ToContext(zipkin.FromHTTP(zipkinAddSpanFunc)))
		after := httptransport.After(httptransport.SetContentType("application/json"))

		var handler http.Handler
		handler = httptransport.NewBinding(ctx, reflect.TypeOf(request{}), jsoncodec.New(), e, before, after)
		handler = encoding.Gzip(handler)
		handler = cors.Middleware(cors.Config{})(handler)
		handler = httpInstrument(requests.With(field), duration.With(field))(handler)

		mux := http.NewServeMux()
		mux.Handle("/add", handler)
		logger.Log("addr", *httpAddr, "transport", "HTTP")
		errc <- http.ListenAndServe(*httpAddr, mux)
	}()

	// Transport: gRPC
	go func() {
		ln, err := net.Listen("tcp", *grpcAddr)
		if err != nil {
			errc <- err
			return
		}
		s := grpc.NewServer() // uses its own context?
		field := metrics.Field{Key: "transport", Value: "grpc"}

		var addServer pb.AddServer
		addServer = grpcBinding{e}
		addServer = grpcInstrument(requests.With(field), duration.With(field))(addServer)

		pb.RegisterAddServer(s, addServer)
		logger.Log("addr", *grpcAddr, "transport", "gRPC")
		errc <- s.Serve(ln)
	}()

	logger.Log("fatal", <-errc)
}

func interrupt() error {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	return fmt.Errorf("%s", <-c)
}

type loggingCollector struct{}

func (loggingCollector) Collect(s *zipkin.Span) error {
	kitlog.With(kitlog.DefaultLogger, "caller", kitlog.DefaultCaller).Log(
		"trace_id", s.TraceID(),
		"span_id", s.SpanID(),
		"parent_span_id", s.ParentSpanID(),
	)
	return nil
}
