package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/so00yeon/go-health-k8s/internal/config"
)

type versionResponse struct {
	Version string `json:"version"`
	Env     string `json:"env"`
}

func Version(w http.ResponseWriter, r *http.Request) {
	cfg := config.Load()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(versionResponse{
		Version: cfg.AppVersion,
		Env:     cfg.AppEnv,
	}); err != nil {
		slog.Error("failed to encode response", "error", err)
	}
}
