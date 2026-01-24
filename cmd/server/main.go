package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/so00yeon/go-health-k8s/internal/config"
	"github.com/so00yeon/go-health-k8s/internal/router"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg := config.Load()
	mux := router.New(cfg)

	addr := ":" + cfg.Port
	slog.Info("server starting", "addr", addr, "env", cfg.AppEnv)
	if err := http.ListenAndServe(addr, mux); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}
