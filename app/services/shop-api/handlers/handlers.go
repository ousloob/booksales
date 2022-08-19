// Package handlers manage the different version of the API.
package handlers

import (
	"log"
	"net/http"
	"os"
)

// Options represent optional parameters.
type Options struct {
	corsOrigin string
}

// WithCORS provides configuration options for CORS.
func WithCORS(origin string) func(opts *Options) {
	return func(opts *Options) {
		opts.corsOrigin = origin
	}
}

// APIMuxConfig contains all the mandatory system required by handelers.
type APIMuxConfig struct {
	shutdown chan os.Signal
	Log      *log.Logger
}

// APIMux constructs a http.Handler with all application routes defined.
func APIMux(cfg APIMuxConfig, options ...func(opts *Options)) http.Handler {
	var opts Options
	for _, option := range options {
		option(&opts)
	}

	// Construct the web.App which holds all routes as well as common
	// Middleware.
	// var app *web.App

}
