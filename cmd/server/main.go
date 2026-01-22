package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/so00yeon/go-health-k8s/internal/config"
	"github.com/so00yeon/go-health-k8s/internal/router"
)

func main() {
	cfg := config.Load()
	mux := router.New()

	addr := ":" + cfg.Port
	fmt.Printf("Server listening on %s (env=%s)\n", addr, cfg.AppEnv)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
