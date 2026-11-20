package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusCode(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	code, err := statusCode(server.URL)
	if err != nil || code != http.StatusNoContent {
		t.Fatalf("statusCode() = %d, %v; want %d, nil", code, err, http.StatusNoContent)
	}
}

func TestStatusCodeRejectsUntrustedTLS(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	defer server.Close()

	if _, err := statusCode(server.URL); err == nil {
		t.Fatal("statusCode accepted an untrusted TLS certificate")
	}
}
