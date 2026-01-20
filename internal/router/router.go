package router

import (
	"net/http"

	"github.com/so00yeon/go-health-k8s/internal/handler"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handler.Health)
	return mux
}
