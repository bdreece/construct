package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/bdreece/construct"
)

func main() {
	mux := construct.Apply(
		http.NewServeMux(),
		WithRoute("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
			// ...
		}),
		WithRoute("GET /foo", func(w http.ResponseWriter, r *http.Request) {
			// ...
		}),
		WithRoute("GET /bar", func(w http.ResponseWriter, r *http.Request) {
			// ...
		}),
		WithRoute("GET /baz", func(w http.ResponseWriter, r *http.Request) {
			// ...
		}),
	)

	srv := construct.New(
		WithPort(4433),
		WithHandler(mux),
		WithReadTimeout(5*time.Second),
		WithWriteTimeout(5*time.Second),
	)

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}

func WithPort(port int) construct.Option[http.Server] {
	return construct.With(func(srv *http.Server) {
		srv.Addr = net.JoinHostPort("", fmt.Sprint(port))
	})
}

func WithHandler(handler http.Handler) construct.Option[http.Server] {
	return construct.With(func(srv *http.Server) {
		srv.Handler = handler
	})
}

func WithReadTimeout(timeout time.Duration) construct.Option[http.Server] {
	return construct.With(func(srv *http.Server) {
		srv.ReadTimeout = timeout
	})
}

func WithWriteTimeout(timeout time.Duration) construct.Option[http.Server] {
	return construct.With(func(srv *http.Server) {
		srv.WriteTimeout = timeout
	})
}

func WithRoute(pattern string, handler http.HandlerFunc) construct.Option[http.ServeMux] {
	return construct.With(func(mux *http.ServeMux) {
		mux.Handle(pattern, handler)
	})
}
