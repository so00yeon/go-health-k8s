package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type probeResponse struct {
	Status string `json:"status"`
}

func (h *Handler) Live(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(probeResponse{Status: "live"}); err != nil {
		slog.Error("failed to encode response", "error", err)
	}
}

func (h *Handler) Ready(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(probeResponse{Status: "ready"}); err != nil {
		slog.Error("failed to encode response", "error", err)
	}
}

