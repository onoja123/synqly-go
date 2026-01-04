package synqly

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient(Config{
		APIKey: "test_key",
	})

	if client.APIKey != "test_key" {
		t.Errorf("Expected API key 'test_key', got '%s'", client.APIKey)
	}

	if client.BaseURL != "https://synqly.onrender.com" {
		t.Errorf("Expected default base URL, got '%s'", client.BaseURL)
	}

	if client.Chat == nil {
		t.Error("Chat service not initialized")
	}
}

func TestNewClientCustomBaseURL(t *testing.T) {
	client := NewClient(Config{
		APIKey:  "test_key",
		BaseURL: "https://custom.api.com",
	})

	if client.BaseURL != "https://custom.api.com" {
		t.Errorf("Expected custom base URL, got '%s'", client.BaseURL)
	}
}
