package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type versionResponse struct {
	Version string `json:"version"`
	Env     string `json:"env"`
}

func (h *Handler) Version(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(versionResponse{
		Version: h.cfg.AppVersion,
		Env:     h.cfg.AppEnv,
	}); err != nil {
		slog.Error("failed to encode response", "error", err)
	}
}
