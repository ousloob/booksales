package main

import (
	"expvar"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ardanlabs/conf/v3"
)

var build = "develop"

func main() {
	log := log.New(os.Stdout, "SHOP-API: ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	// Perform the startup and shutdown sequence.
	if err := run(log); err != nil {
		os.Exit(1)
	}
}

func run(log *log.Logger) error {

	// =========================================================================
	// Configuration

	cfg := struct {
		conf.Version
		Web struct {
			ReadTimeout     time.Duration `conf:"default:5s"`
			WriteTimeOut    time.Duration `conf:"default:10s"`
			IdleTimeout     time.Duration `conf:"default:12s"`
			ShutdownTimeout time.Duration `conf:"default:20s"`
			APIHost         string        `conf:"default:0.0.0.0:3000"`
		}
	}{
		Version: conf.Version{
			Build: build,
			Desc:  "© 2021 Oussama Moulana",
		},
	}

	const prefix = "SHOP"
	_, err := conf.Parse(prefix, &cfg)
	if err != nil {
		return fmt.Errorf("parsing config: %w", err)
	}

	// =========================================================================
	// App Starting

	log.Printf("starting service version %s", build)
	defer log.Println("shutdown complete")

	out, err := conf.String(&cfg)
	if err != nil {
		return fmt.Errorf("generating config for output: %w", err)
	}
	log.Printf("startup config %s", out)

	expvar.NewString("build").Set(build)

	// =========================================================================
	// Start API Service

	log.Print("startup status initalizing V1 API support")

	// Make a channel to listen for an interrupt or terminal signal from the OS.
	// Use a buffered channel because the signal package require it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	// Construct the mux for the API calls.
	// apiMux :=

	// Construct a server to service the request against the mux.
	api := http.Server{
		Addr:         cfg.Web.APIHost,
		ReadTimeout:  cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeOut,
		IdleTimeout:  cfg.Web.IdleTimeout,
		ErrorLog:     log,
	}

	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this
	// error.

	serveErrors := make(chan error, 1)

	// Start the service listening for api requests.
	go func() {
		log.Printf("startup status api router started host %s", api.Addr)
		serveErrors <- api.ListenAndServe()
	}()

	return nil
}
