package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerRoot(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api", nil)
	res := httptest.NewRecorder()

	Handler(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, res.Code)
	}
	if got := res.Header().Get("Content-Type"); got == "" {
		t.Fatal("expected a JSON content type")
	}
}

func TestHandlerHealth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	res := httptest.NewRecorder()

	Handler(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, res.Code)
	}
	if body := res.Body.String(); body != `{"status":"ok"}` {
		t.Fatalf("unexpected response body: %q", body)
	}
}
