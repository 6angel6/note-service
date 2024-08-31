package main

import (
	"Zametki-go/internal/config"
	"Zametki-go/internal/handler"
	"Zametki-go/internal/repository"
	"Zametki-go/internal/service"
	"context"
	"errors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.Read()

	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	//DI
	repo := repository.NewRepository(db)
	s := service.NewService(repo)
	h := handler.NewHandler(s)

	srv := http.Server{
		Addr:           cfg.HTTPAddr,
		Handler:        h.InitRoutes(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// listen to OS signals and gracefully shutdown HTTP server
	stopped := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("HTTP Server Shutdown Error: %v", err)
		}
		close(stopped)
	}()

	log.Printf("Starting HTTP server on %s", cfg.HTTPAddr)

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("HTTP server ListenAndServe Error: %v", err)
	}

	return nil
}
