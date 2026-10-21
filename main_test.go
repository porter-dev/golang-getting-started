package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRoutes(t *testing.T) {
	tests := []struct {
		method     string
		path       string
		wantStatus int
		wantBody   string
	}{
		{http.MethodGet, "/", http.StatusOK, `{"message":"Hello from Go on Porter!"}`},
		{http.MethodGet, "/healthz", http.StatusOK, `{"status":"ok"}`},
		{http.MethodGet, "/missing", http.StatusNotFound, "404 page not found"},
		{http.MethodPost, "/healthz", http.StatusMethodNotAllowed, "Method Not Allowed"},
	}

	router := newRouter()
	for _, tt := range tests {
		t.Run(tt.method+" "+tt.path, func(t *testing.T) {
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, httptest.NewRequest(tt.method, tt.path, nil))

			if rec.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
			if got := strings.TrimSpace(rec.Body.String()); got != tt.wantBody {
				t.Fatalf("body = %q, want %q", got, tt.wantBody)
			}
		})
	}
}
