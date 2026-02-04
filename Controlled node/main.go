package main

import (
	"controlled-node/config"
	"controlled-node/core"
	"controlled-node/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.Load()
	sys := core.NewSystem()

	log.Println("SYSTEM INIT")
	sys.Transition(core.Running)

	mux := http.NewServeMux()
	mux.Handle("/status", handlers.StatusHandler(sys))
	mux.Handle("/fault", handlers.FaultHandler(sys))

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: mux,
	}

	go func() {
		log.Println("SERVER STARTED on port", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("SYSTEM DRAINING")
	sys.Transition(core.Draining)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Shutdown(ctx)

	log.Println("SYSTEM STOPPED")
	sys.Transition(core.Stopped)
}
