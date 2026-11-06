package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestRoutes exercises handlers in-process with httptest: no ports, no network.
func TestRoutes(t *testing.T) {
	h := routes(newStore())

	tests := []struct {
		name, method, path, body string
		want                     int
	}{
		{"health", "GET", "/healthz", "", http.StatusOK},
		{"create", "POST", "/items", `{"id":"1","name":"mug"}`, http.StatusCreated},
		{"fetch", "GET", "/items/1", "", http.StatusOK},
		{"missing", "GET", "/items/999", "", http.StatusNotFound},
		{"bad json", "POST", "/items", `{"nope":true}`, http.StatusBadRequest},
		{"wrong method", "PUT", "/items/1", "", http.StatusMethodNotAllowed},
	}
	for _, tt := range tests { // sequential: later cases depend on "create"
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, strings.NewReader(tt.body))
			rec := httptest.NewRecorder()
			h.ServeHTTP(rec, req)
			if rec.Code != tt.want {
				t.Fatalf("%s %s = %d, want %d (body %q)", tt.method, tt.path, rec.Code, tt.want, rec.Body.String())
			}
		})
	}
}
