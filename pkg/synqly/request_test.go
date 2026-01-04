package synqly

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestChatCreate_RequestsExpectedURLAndHeaders(t *testing.T) {
	// Arrange: mock Synqly API
	var gotMethod, gotPath, gotKeyHeader, gotContentType string
	var gotBody map[string]interface{}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		gotKeyHeader = r.Header.Get("x-synqly-key")
		gotContentType = r.Header.Get("Content-Type")

		b, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(b, &gotBody)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(APIResponse{
			Success: true,
			Message: "ok",
			Data: &ChatResponse{
				ID:        "test_1",
				Provider:  "openai",
				ModelType: "gpt-4",
				Content:   "hello",
			},
		})
	}))
	defer srv.Close()

	client := NewClient(Config{APIKey: "test_key", BaseURL: srv.URL})

	// Act
	resp, err := client.Chat.Create(ChatCreateParams{
		Provider: "openai",
		Model:    "gpt-4",
		Messages: []Message{{Role: "user", Content: "Hello"}},
	})

	// Assert
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp == nil || resp.Content != "hello" {
		t.Fatalf("unexpected response: %#v", resp)
	}
	if gotMethod != http.MethodPost {
		t.Fatalf("expected POST, got %q", gotMethod)
	}
	if gotPath != "/api/v1/chat/create-chat" {
		t.Fatalf("expected path /api/v1/chat/create-chat, got %q", gotPath)
	}
	if gotKeyHeader != "test_key" {
		t.Fatalf("expected x-synqly-key header to be set")
	}
	if gotContentType != "application/json" {
		t.Fatalf("expected Content-Type application/json, got %q", gotContentType)
	}
	if gotBody == nil {
		t.Fatalf("expected JSON body")
	}
}

func TestRequest_AuthErrorReturnsErr(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"success":false,"error":"Invalid or inactive API key."}`))
	}))
	defer srv.Close()

	client := NewClient(Config{APIKey: "bad_key", BaseURL: srv.URL})

	_, err := client.Chat.Create(ChatCreateParams{
		Provider: "openai",
		Model:    "gpt-4",
		Messages: []Message{{Role: "user", Content: "Hello"}},
	})

	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
