package status

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/svdx9/conjugate-cc/backend/internal/api/v1"
)

func TestHandler_GetStatus(t *testing.T) {
	h := NewHandler("sha", "time")
	req := httptest.NewRequest(http.MethodGet, "/v1/status", nil)
	w := httptest.NewRecorder()

	h.GetStatus(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var resp api.Status
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Status != "ok" {
		t.Errorf("expected status ok, got %s", resp.Status)
	}
}

func TestHandler_GetMetadata(t *testing.T) {
	gitSHA := "d32e693"
	buildTime := "2026-03-14T17:33:30Z"
	h := NewHandler(gitSHA, buildTime)
	req := httptest.NewRequest(http.MethodGet, "/v1/metadata", nil)
	w := httptest.NewRecorder()

	h.GetMetadata(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var resp api.Metadata
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.GitSha != gitSHA {
		t.Errorf("expected git_sha %s, got %s", gitSHA, resp.GitSha)
	}

	if resp.BuildTime != buildTime {
		t.Errorf("expected build_time %s, got %s", buildTime, resp.BuildTime)
	}
}
