package router

import (
	"net/http"

	"github.com/so00yeon/go-health-k8s/internal/config"
	"github.com/so00yeon/go-health-k8s/internal/handler"
)

func New(cfg *config.Config) *http.ServeMux {
	mux := http.NewServeMux()

	h := handler.New(cfg)
	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("GET /live", h.Live)
	mux.HandleFunc("GET /ready", h.Ready)
	mux.HandleFunc("GET /version", h.Version)

	return mux
}
