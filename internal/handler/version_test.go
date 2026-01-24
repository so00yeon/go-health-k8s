package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/so00yeon/go-health-k8s/internal/config"
)

func TestVersion(t *testing.T) {
	cfg := &config.Config{
		AppVersion: "1.2.3",
		AppEnv:     "test",
	}
	h := New(cfg)

	req := httptest.NewRequest(http.MethodGet, "/version", nil)
	rec := httptest.NewRecorder()

	h.Version(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}

	var body versionResponse
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if body.Version != "1.2.3" {
		t.Errorf("expected version '1.2.3', got '%s'", body.Version)
	}

	if body.Env != "test" {
		t.Errorf("expected env 'test', got '%s'", body.Env)
	}
}
