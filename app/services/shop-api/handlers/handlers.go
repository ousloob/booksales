// Package handlers manges the different versions of the API.
package handlers

import (
	"log"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	v1 "github.com/oussamm/bookstore/app/services/shop-api/handlers/v1"
)

type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *log.Logger
}

func APIMux(cfg APIMuxConfig) *chi.Mux {
	app := chi.NewRouter()
	app.Use(middleware.Logger)

	v1.Routes(app, v1.Config{
		Log: cfg.Log,
	})

	return app
}
