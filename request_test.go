package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMakeRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello from test server"))
	}))
	defer server.Close()

	body, resp, err := makeRequest(server.URL)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	expectedBody := "Hello from test server"

	if body != expectedBody {
		t.Errorf("expected body %q, got %q", expectedBody, body)
	}
}

func TestMakeRequestInvalidURL(t *testing.T) {
	_, _, err := makeRequest("://invalid-url")

	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}

func TestMakeRequestNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	}))
	defer server.Close()

	_, resp, err := makeRequest(server.URL)
	if err != nil {
		t.Fatalf("expected no request error, got %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, resp.StatusCode)
	}
}
