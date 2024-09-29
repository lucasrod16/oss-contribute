package main

import (
	"context"
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lucasrod16/oss-contribute/internal/cache"
	ihttp "github.com/lucasrod16/oss-contribute/internal/http"
)

//go:embed ui/build/*
var ui embed.FS

func main() {
	fs, err := fs.Sub(ui, "ui/build")
	if err != nil {
		log.Fatalf("failed to load UI assets: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := cache.New()

	ticker := time.NewTicker(12 * time.Hour)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Fatal(ctx.Err())
			case <-ticker.C:
				if err := c.RepoData(ctx); err != nil {
					log.Printf("Error fetching GitHub repo data: %v", err)
				}
			}
		}
	}()

	// initial fetch on startup
	if err := c.RepoData(ctx); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(fs)))
	mux.Handle("/repos", ihttp.GetRepos(c))

	rl := ihttp.NewRateLimiter()
	limitedMux := rl.Limit(mux)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: limitedMux,
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("API server listening on port 8080")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-shutdown
	log.Println("Shutting down server...")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("Server gracefully shutdown")
}
