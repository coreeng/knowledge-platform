package main

import (
	"errors"
	"github.com/coreeng/minimal-reference-application-go/pkg/router"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	"github.com/slok/go-http-metrics/middleware/std"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("minimal-reference-app is starting")
	mw := middleware.New(middleware.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})

	// External server
	httpServer := &http.Server{
		Addr:              ":8080",
		Handler:           std.Handler("", mw, router.NewServer()),
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       5 * time.Second,
	}

	// Internal server that exposes promQL metrics
	internalServer := &http.Server{
		Addr:              ":8081",
		Handler:           std.Handler("", mw, router.NewInternalServer()),
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       5 * time.Second,
	}

	go start(internalServer, "metric server")
	start(httpServer, "canary test")
}

func start(server *http.Server, name string) {
	log.Printf("%s server has started listening on %s", name, server.Addr)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("start server: %w", err)
		return
	}
	log.Println("Application shut down gracefully")
}
