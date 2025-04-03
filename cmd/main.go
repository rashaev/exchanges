package main

import (
	"context"
	"exchanges/internal/config"
	"exchanges/internal/logger"
	"exchanges/internal/server"
	"log"
	"os"
	"os/signal"
)

func main() {
	// Load Config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize logger
	logger := logger.New(cfg)

	// Создаем базовый контекст и добавляем в него логгер
	baseCtx := context.WithValue(context.Background(), "logger", logger)

	ctx, cancel := signal.NotifyContext(baseCtx, os.Interrupt)
	defer cancel()

	// Initialize server
	srv := server.New(ctx, cfg)

	go func() {
		// Start server
		log.Printf("Starting server on %s", cfg.Server.Address)
		log.Fatal(srv.Start())
	}()

	// Wait for shutdown signal and initiate graceful shutdown once received.
	<-ctx.Done()
	log.Print("shutdown signal received, initiate shutdown")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("failed shutting down gracefully: %v", err)
	}
	log.Println("server stopped")
}
