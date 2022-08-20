// Package handlers manges the different versions of the API.
package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func App() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

}
