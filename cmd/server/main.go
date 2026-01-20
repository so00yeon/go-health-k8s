package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/so00yeon/go-health-k8s/internal/router"
)

func main() {
	mux := router.New()

	addr := ":8080"
	fmt.Printf("Server listening on %s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
