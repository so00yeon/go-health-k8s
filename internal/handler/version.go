package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type versionResponse struct {
	Version string `json:"version"`
}

func Version(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(versionResponse{Version: "0.1.0"}); err != nil {
		slog.Error("failed to encode response", "error", err)
	}
}
