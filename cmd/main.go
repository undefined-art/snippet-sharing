package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"snippet-sharing/config"
	"snippet-sharing/internal/routes"
	"syscall"
	"time"
)

func main() {
	environment := flag.String("e", "development", "Environment to run in (development, production)")

	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		fmt.Println("  -e string")
		fmt.Println("        Environment mode: development, production (default: development)")
		os.Exit(1)
	}

	flag.Parse()
	config.Init(*environment)

	Init()
}

func Init() {
	cfg := config.GetConfig()
	r := routes.NewRouter()

	addr := cfg.GetString("server.address")
	if addr == "" {
		addr = "localhost:8080"
	}

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	go func() {
		log.Printf("Server starting on %s", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
