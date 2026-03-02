package api

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	apphttp "github.com/svdx9/conjugate-cc/backend/internal/http"
	statusservice "github.com/svdx9/conjugate-cc/backend/internal/status/service"
)

func TestHandlerServesStatusEndpoints(t *testing.T) {
	t.Parallel()

	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	service := statusservice.New("git-sha-123", "2026-03-02T18:00:00Z")
	handler := NewHandler(service)
	router := apphttp.NewRouter(logger, handler)

	testCases := []struct {
		name           string
		path           string
		expectedStatus int
		expectedBody   map[string]string
	}{
		{
			name:           "health",
			path:           "/v1/health",
			expectedStatus: http.StatusOK,
			expectedBody: map[string]string{
				"status": "ok",
			},
		},
		{
			name:           "build info",
			path:           "/v1/build-info",
			expectedStatus: http.StatusOK,
			expectedBody: map[string]string{
				"gitSha":    "git-sha-123",
				"buildTime": "2026-03-02T18:00:00Z",
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			request := httptest.NewRequest(http.MethodGet, testCase.path, nil)
			responseRecorder := httptest.NewRecorder()

			router.ServeHTTP(responseRecorder, request)

			response := responseRecorder.Result()
			defer response.Body.Close()

			if response.StatusCode != testCase.expectedStatus {
				t.Fatalf("status code = %d, want %d", response.StatusCode, testCase.expectedStatus)
			}

			body := map[string]string{}
			decoder := json.NewDecoder(response.Body)
			err := decoder.Decode(&body)
			if err != nil {
				t.Fatalf("decode response: %v", err)
			}

			for key, expectedValue := range testCase.expectedBody {
				if body[key] != expectedValue {
					t.Fatalf("%s = %q, want %q", key, body[key], expectedValue)
				}
			}
		})
	}
}
