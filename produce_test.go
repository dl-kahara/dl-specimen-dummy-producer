package main

import (
	"crypto/sha256"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_post(t *testing.T) {
	SetupMetrics()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Fatalf("wanted POST, got %s", r.Method)
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatal("Could not read request body")
		}
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(201)
		checksum := sha256.Sum256(body)
		w.Write(checksum[:])
	}))
	defer server.Close()

	config := Config{
		PrimeBits: 2048,
		Consumer:  server.URL,
	}

	for range 10 {
		post(config)
	}
}
