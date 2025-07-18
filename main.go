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

	"github.com/lucasrod16/oss-projects/internal/cache"
	ihttp "github.com/lucasrod16/oss-projects/internal/http"
)

//go:embed ui/build/*
var ui embed.FS

func main() {
	fs, err := fs.Sub(ui, "ui/build")
	if err != nil {
		log.Fatalf("failed to load UI assets: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	c := cache.New()

	ticker := time.NewTicker(12 * time.Hour)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println(ctx.Err())
				return
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
	mux.Handle("GET /", http.FileServer(http.FS(fs)))
	mux.Handle("GET /repos", ihttp.GetRepos(c))

	rl := ihttp.NewRateLimiter()
	limitedMux := rl.Limit(mux)

	server := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      limitedMux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	go func() {
		log.Println("API server listening on port 8080")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	stop()
	log.Println("Shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatal(err)
	}
	log.Println("Server gracefully shutdown")
}
