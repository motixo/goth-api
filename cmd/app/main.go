package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mot0x0/goth-api/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	server, err := InitializeApp()
	if err != nil {
		panic("failed to initialize app: " + err.Error())
	}

	go func() {
		log.Printf("Server starting on port %s", cfg.ServerPort)
		if err := server.Run(cfg.ServerPort); err != nil {
			panic("server failed" + err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down...")
}
