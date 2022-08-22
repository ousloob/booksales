// Package v1 contains the full set of handler functions and routes supported
// by the v1 web api.
package v1

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Config struct {
	Log *log.Logger
}

// Routes binds all the version 1 routes.
func Routes(router *chi.Mux, cfg Config) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Salam Alaikoum"))
	})
}
