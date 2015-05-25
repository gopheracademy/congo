package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	stdlog "log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"time"

	"github.com/go-kit/kit/server"
	jsoncodec "github.com/go-kit/kit/transport/codec/json"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/streadway/handy/cors"
	"github.com/streadway/handy/encoding"
	"golang.org/x/net/context"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/expvar"
	"github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/kit/metrics/statsd"
	"github.com/go-kit/kit/tracing/zipkin"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	var (
		debugAddr = flag.String("debug.addr", ":8000", "Address for HTTP debug/instrumentation server")
		httpAddr  = flag.String("http.addr", ":8001", "Address for HTTP (JSON) server")
	)
	flag.Parse()

	// `package log` domain
	// This logger should get passed into other components
	// and wrapped with component specifics
	var logger kitlog.Logger
	logger = kitlog.NewPrefixLogger(os.Stderr)
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestampUTC)

	kitlog.DefaultLogger = logger                     // for other gokit components
	stdlog.SetOutput(kitlog.NewStdlibAdapter(logger)) // redirect stdlib logging to us
	stdlog.SetFlags(0)                                // flags are handled in our logger
	// `package metrics` domain
	requests := metrics.NewMultiCounter(
		expvar.NewCounter("requests"),
		statsd.NewCounter(ioutil.Discard, "requests_total", time.Second),
		prometheus.NewCounter(stdprometheus.CounterOpts{
			Namespace: "congo",
			Name:      "requests_total",
			Help:      "Total number of received requests.",
		}, []string{}),
	)
	duration := metrics.NewMultiHistogram(
		expvar.NewHistogram("duration_nanoseconds_total", 0, 100000000, 3),
		statsd.NewHistogram(ioutil.Discard, "duration_nanoseconds_total", time.Second),
		prometheus.NewSummary(stdprometheus.SummaryOpts{
			Namespace: "congo",
			Name:      "duration_nanoseconds_total",
			Help:      "Total nanoseconds spend serving requests.",
		}, []string{}),
	)

	var e server.Endpoint
	//	e = makeEndpoint(a)

	// Mechanical stuff
	root := context.Background()
	errc := make(chan error)

	go func() {
		errc <- interrupt()
	}()

	// Transport: HTTP (debug/instrumentation)
	go func() {
		logger.Log("congo", *debugAddr, "transport", "debug")
		errc <- http.ListenAndServe(*debugAddr, nil)
	}()

	// Transport: HTTP (JSON)
	go func() {
		ctx, cancel := context.WithCancel(root)
		defer cancel()

		field := metrics.Field{Key: "transport", Value: "http"}
		after := httptransport.After(httptransport.SetContentType("application/json"))

		var handler http.Handler
		handler = httptransport.NewBinding(ctx, reflect.TypeOf(request{}), jsoncodec.New(), e, nil, after)
		handler = encoding.Gzip(handler)
		handler = cors.Middleware(cors.Config{})(handler)
		handler = httpInstrument(requests.With(field), duration.With(field))(handler)
		// move this out so it is available for other components to add subroutes?
		mux := http.NewServeMux()
		// Add handlers here
		logger.Log("congo", *httpAddr, "transport", "HTTP")
		errc <- http.ListenAndServe(*httpAddr, mux)
	}()

	logger.Log("fatal", <-errc)
}
func httpInstrument(requests metrics.Counter, duration metrics.Histogram) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requests.Add(1)
			defer func(begin time.Time) { duration.Observe(time.Since(begin).Nanoseconds()) }(time.Now())
			next.ServeHTTP(w, r)
		})
	}
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

type request struct {
}

type response struct {
}
