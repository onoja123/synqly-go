package synqly

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Client is the Synqly API client
type Client struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
	Chat       *ChatService
}

// Config holds client configuration
type Config struct {
	APIKey  string
	BaseURL string
}

// NewClient creates a new Synqly client
func NewClient(config Config) *Client {
	if config.BaseURL == "" {
		config.BaseURL = "https://api.synqly.xyz"
	}
	config.BaseURL = strings.TrimRight(config.BaseURL, "/")

	client := &Client{
		APIKey:  config.APIKey,
		BaseURL: config.BaseURL,
		HTTPClient: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
	client.Chat = &ChatService{client: client}

	return client
}

// request makes an HTTP request to the Synqly API
func (c *Client) request(method, endpoint string, body interface{}, result interface{}) error {
	if !strings.HasPrefix(endpoint, "/") {
		endpoint = "/" + endpoint
	}
	if !strings.HasPrefix(endpoint, "/api/v1") {
		endpoint = "/api/v1" + endpoint
	}

	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, c.BaseURL+endpoint, reqBody)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("x-synqly-key", c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error (%d): %s", resp.StatusCode, string(bodyBytes))
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}
