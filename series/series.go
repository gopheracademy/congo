package main

import (
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/server"
)

type SeriesStorage interface {
	List() ([]Series, error)
	Get() (Series, error)
	Update(series) error
	Create(series) error
}

//Series represents a repeating event.  If GopherCon 2015 is an event
//then GopherCon is a series.  Series will
type Series struct {
	ID   string // UUID or int, tbd
	Name string
}

type SeriesService struct {
	log     kitlog.Log
	mux     gorilla.Mux
	path    string // "/series"
	storage SeriesStorage
}
type SeriesPgStorage struct {
	DSN string
}

func (ss *SeriesService) List() ([]Series, error) {
	return ss.storage.List()
}
func (ss *SeriesService) Get(r getrequest) (Series, error) {
	return ss.storage.Get(getrequest.id)
}

func (pgs *SeriesPgStorage) List() ([]Series, error) {
}
func (pgs *SeriesPgStorage) Get(id string) (Series, error) {
}
func (pgs *SeriesPgStorage) Update(s Series) error {
}
func (pgs *SeriesPgStorage) Create(s Series) error {
}

func (ss *SeriesService) ListEndpoint() server.Endpoint {
	return func(ctx context.Context, req server.Request) (server.Response, error) {
		select {
		case <-ctx.Done():
			return nil, server.ErrContextCanceled
		default:
		}

		listReq, ok := req.(*listrequest)
		if !ok {
			return nil, server.ErrBadCast
		}

		s, err := ss.List()

		return response{
			series: s,
		}, err
	}
}

func (ss *SeriesService) GetEndpoint() server.Endpoint {
	return func(ctx context.Context, req server.Request) (server.Response, error) {
		select {
		case <-ctx.Done():
			return nil, server.ErrContextCanceled
		default:
		}

		getReq, ok := req.(*getrequest)
		if !ok {
			return nil, server.ErrBadCast
		}

		s, err := ss.Get(getReq)

		return response{
			series: s,
		}, err
	}
}
