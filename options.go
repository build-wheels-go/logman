package logman

import (
	"io"
	"os"
)

type options struct {
	output        io.Writer
	level         Level
	formatter     Formatter
	disableCaller bool
}

type Option func(*options)

func initOptions(opts ...Option) (o *options) {
	o = &options{}
	for _, opt := range opts {
		opt(o)
	}
	if o.output == nil {
		o.output = os.Stderr
	}
	if o.formatter == nil {
		o.formatter = &TextFormatter{}
	}
	return
}

func WithOutput(output io.Writer) Option {
	return func(o *options) {
		o.output = output
	}
}

func WithLevel(level Level) Option  {
	return func(o *options) {
		o.level = level
	}
}

func WithFormatter(formatter Formatter) Option {
	return func(o *options) {
		o.formatter = formatter
	}
}

func WithDisableCaller(disableCaller bool) Option  {
	return func(o *options) {
		o.disableCaller = disableCaller
	}
}
