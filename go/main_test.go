package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func captureStatusCheckLog(url string) string {
	var output bytes.Buffer
	originalOutput := log.Writer()
	log.SetOutput(&output)
	defer log.SetOutput(originalOutput)

	logStatusCheck(url)
	return output.String()
}

func TestLogStatusCheckDoesNotLogRequestURL(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	statusURL, err := url.Parse(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	statusURL.User = url.UserPassword("username", "password")
	statusURL.RawQuery = "token=sensitive"
	logged := captureStatusCheckLog(statusURL.String())

	if !strings.Contains(logged, "status check -> 204") {
		t.Fatalf("logStatusCheck() output = %q; want status code", logged)
	}
	for _, sensitive := range []string{"username", "password", "token=sensitive"} {
		if strings.Contains(logged, sensitive) {
			t.Errorf("logStatusCheck() output contains %q: %q", sensitive, logged)
		}
	}
}

func TestLogStatusCheckDoesNotLogRequestError(t *testing.T) {
	statusURL := url.URL{
		Scheme:   "http",
		User:     url.UserPassword("username", "password"),
		Host:     "127.0.0.1:0",
		Path:     "/status",
		RawQuery: "token=sensitive",
	}
	logged := captureStatusCheckLog(statusURL.String())
	if !strings.Contains(logged, "status check failed") {
		t.Fatalf("logStatusCheck() output = %q; want a generic failure message", logged)
	}
	for _, sensitive := range []string{"username", "password", "token=sensitive"} {
		if strings.Contains(logged, sensitive) {
			t.Errorf("logStatusCheck() output contains %q: %q", sensitive, logged)
		}
	}
}

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
