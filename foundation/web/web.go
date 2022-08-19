// Package web contains a small web framwork extension.
package web

import (
	"context"
	"net/http"
	"os"

	"github.com/dimfeld/httptreemux/v5"
)

// A Handler is a type that handles a http request within our own little mini
// framework.
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers.
type App struct {
	mux      *httptreemux.ContextMux
	shutdown chan os.Signal
	mw       []Middleware
}

// NewApp creates an App value that handle a set of routes for the application.
func NewApp(shutdown chan os.Signal, mw ...Middleware) *App {

	// Create an OpenTelemetry HTTP Handler which wraps our router. This will
	// start the initial span and annotate it with information about the
	// request/response.
	//
	// This is configured to use the W3C TraceContext standard to set the remote
	// parent if a client request inclides the appropriate headers.
	// https://w3c.github.io/trace-context/

	mux := httptreemux.NewContextMux()

	return &App{
		mux: mux,
	}
}
