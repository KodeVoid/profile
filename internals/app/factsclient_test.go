package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// mock response struct
type mockResponse struct {
	Fact string `json:"fact"`
}

func TestGetFact_Success(t *testing.T) {
	expectedFact := "Cats sleep for 70% of their lives."
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := mockResponse{Fact: expectedFact}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			t.Fatalf("failed to encode JSON response: %v", err)
		}

	}))
	defer server.Close()

	got, err := GetFact(server.URL)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if got != expectedFact {
		t.Errorf("expected %q, got %q", expectedFact, got)
	}
}

func TestGetFact_NonOKStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}))
	defer server.Close()

	_, err := GetFact(server.URL)
	if err == nil {
		t.Fatal("expected an error for non-OK status, got nil")
	}
}

func TestGetFact_InvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("not a valid json")); err != nil {
			t.Fatalf("failed to write invalid JSON response: %v", err)
		}

	}))
	defer server.Close()

	_, err := GetFact(server.URL)
	if err == nil {
		t.Fatal("expected JSON unmarshal error, got nil")
	}
}
