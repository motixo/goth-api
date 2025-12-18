package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/motixo/goat-api/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	app, err := InitializeApp()
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}

	go func() {
		if err := app.Server.Run(cfg.ServerPort); err != nil {
			if err.Error() != "http: Server closed" {
				log.Fatalf("Server failed to run: %v", err)
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// a) Shut down HTTP server (stop accepting new requests)
	log.Println("Shutting down HTTP server...")
	if err := app.Server.Shutdown(shutdownCtx); err != nil { // NOTE: Your http.Server needs a Shutdown method
		log.Printf("HTTP Server forced to shutdown: %v", err)
	}

	// b) Wait for background Event Handlers to complete
	log.Println("Waiting for background event handlers to finish...")
	app.EventBus.Wait()
}
