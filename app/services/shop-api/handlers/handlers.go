// Package handlers manges the different versions of the API.
package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *log.Logger
}

func APIMux(cfg APIMuxConfig) http.Handler {
	app := chi.NewRouter()
	app.Use(middleware.Logger)

	return app
}
