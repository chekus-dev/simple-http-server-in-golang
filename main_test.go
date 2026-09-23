package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootRouteReturnsJSONWelcome(t *testing.T) {
	mux := newServerMux()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	mux.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, res.Code)
	}

	var payload map[string]string
	if err := json.Unmarshal(res.Body.Bytes(), &payload); err != nil {
		t.Fatalf("expected valid JSON response: %v", err)
	}

	if payload["message"] != "Welcome" {
		t.Fatalf("expected welcome message, got %q", payload["message"])
	}
}

func TestCustomerEndpointRejectsNonGetMethods(t *testing.T) {
	mux := newServerMux()
	req := httptest.NewRequest(http.MethodPost, "/customer", nil)
	res := httptest.NewRecorder()

	mux.ServeHTTP(res, req)

	if res.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status %d, got %d", http.StatusMethodNotAllowed, res.Code)
	}
}
