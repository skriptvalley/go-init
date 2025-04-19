package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	
	log "github.com/skriptvalley/go-init/pkg/logger"
	"github.com/skriptvalley/go-init/internal/middleware"
	"github.com/skriptvalley/go-init/internal/config"
)

func Start() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	if err := cfg.Validate(); err != nil {
		panic(err)
	}

	logger, err := log.NewLogger(cfg.LogLevel)
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Infow("Starting HTTP server", "port", cfg.Port)
	
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		sugar.Infow("Health check endpoint hit")
		fmt.Fprintln(w, "ok")
	})
	
	handler := middleware.RecoveryMiddleware(sugar)(
		middleware.LoggingMiddleware(sugar)(mux),
	)
	
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: handler,
	}
	
	// Run server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			sugar.Fatalw("Server failed unexpectedly", "error", err)
		}
	}()
	
	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	sugar.Infow("Shutting down server...")
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	if err := srv.Shutdown(ctx); err != nil {
		sugar.Errorw("Server forced to shutdown", "error", err)
	} else {
		sugar.Infow("Server exited properly")
	}
}
