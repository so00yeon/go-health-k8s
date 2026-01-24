package handler

import "github.com/so00yeon/go-health-k8s/internal/config"

type Handler struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Handler {
	return &Handler{cfg: cfg}
}
