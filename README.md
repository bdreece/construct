# construct

A generic library for implementing the functional options pattern in Go.

## Installation

```sh
go get -u github.com/bdreece/construct
```

## Usage

```go
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
		withRoute("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
			// ...
		}),
		withRoute("GET /foo", func(w http.ResponseWriter, r *http.Request) {
			// ...
		}),
		withRoute("GET /bar", func(w http.ResponseWriter, r *http.Request) {
			// ...
		}),
		withRoute("GET /baz", func(w http.ResponseWriter, r *http.Request) {
			// ...
		}),
	)

	srv := construct.New(
		withPort(4433),
		withHandler(mux),
		withReadTimeout(5*time.Second),
		withWriteTimeout(5*time.Second),
	)

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}

func withPort(port int) construct.Option[http.Server] {
	return construct.With(func(srv *http.Server) {
		srv.Addr = net.JoinHostPort("", fmt.Sprint(port))
	})
}

func withHandler(handler http.Handler) construct.Option[http.Server] {
	return construct.With(func(srv *http.Server) {
		srv.Handler = handler
	})
}

func withReadTimeout(timeout time.Duration) construct.Option[http.Server] {
	return construct.With(func(srv *http.Server) {
		srv.ReadTimeout = timeout
	})
}

func withWriteTimeout(timeout time.Duration) construct.Option[http.Server] {
	return construct.With(func(srv *http.Server) {
		srv.WriteTimeout = timeout
	})
}

func withRoute(pattern string, handler http.HandlerFunc) construct.Option[http.ServeMux] {
	return construct.With(func(mux *http.ServeMux) {
		mux.Handle(pattern, handler)
	})
}
```

## Examples

See more examples under the [examples/](./examples) directory.

---

MIT License

Copyright (c) 2025 Brian Reece
