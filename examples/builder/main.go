package main

import (
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"github.com/bdreece/construct"
)

type config struct {
	src    string
	secret string
}

type handler struct {
	tmpl   *template.Template
	client *http.Client
}

var env string

func init() {
	flag.StringVar(&env, "e", "development", "environment")
}

func main() {
	flag.Parse()
	builder := construct.NewBuilder(newLogger)
	if env == "development" {
		builder.Apply(withLevel(slog.LevelDebug))
	} else {
		builder.Apply(withLevel(slog.LevelInfo))
	}

	logger := builder.Apply(withSource(true)).Build()
	logger.Info("hello, world!")
}

func newLogger(opts slog.HandlerOptions) *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &opts))
}

func withLevel(level slog.Level) construct.Option[slog.HandlerOptions] {
	return construct.With(func(opts *slog.HandlerOptions) {
		opts.Level = level
	})
}

func withSource(addSource bool) construct.Option[slog.HandlerOptions] {
	return construct.With(func(opts *slog.HandlerOptions) {
		opts.AddSource = addSource
	})
}
