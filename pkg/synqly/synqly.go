package synqly

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client is the main Synqly client
type Client struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
	Chat       *ChatService
	Analytics  *AnalyticsService
	Keys       *KeysService
}

// Config holds the client configuration
type Config struct {
	APIKey  string
	BaseURL string
}

// UsageStats represents usage statistics
type UsageStats struct {
	Success bool `json:"success"`
	Data    []struct {
		Provider      string  `json:"provider"`
		TotalRequests int     `json:"totalRequests"`
		TotalTokens   int     `json:"totalTokens"`
		TotalCost     float64 `json:"totalCost"`
		AvgLatency    float64 `json:"avgLatency"`
	} `json:"data"`
}

// NewClient creates a new Synqly client
func NewClient(config Config) *Client {
	if config.BaseURL == "" {
		config.BaseURL = "https://api.synqly.com/api/v1"
	}

	client := &Client{
		APIKey:  config.APIKey,
		BaseURL: config.BaseURL,
		HTTPClient: &http.Client{
			Timeout: 60 * time.Second,
		},
	}

	// Initialize services
	client.Chat = &ChatService{client: client}
	client.Analytics = &AnalyticsService{client: client}
	client.Keys = &KeysService{client: client}

	return client
}

// request makes an HTTP request
func (c *Client) request(method, endpoint string, body interface{}, result interface{}) error {
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
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

// AnalyticsService handles analytics operations
type AnalyticsService struct {
	client *Client
}

// UsageParams holds parameters for usage queries
type UsageParams struct {
	StartDate string
	EndDate   string
	Provider  string
}

// Usage retrieves usage statistics
func (s *AnalyticsService) Usage(params UsageParams) (*UsageStats, error) {
	endpoint := "/analytics/usage"
	if s.client.BaseURL == "https://api.synqly.com/api/v1" && endpoint[:4] != "/api" {
		endpoint = "/api/v1" + endpoint
	}
	if params.StartDate != "" || params.EndDate != "" || params.Provider != "" {
		endpoint += "?"
		if params.StartDate != "" {
			endpoint += "startDate=" + params.StartDate + "&"
		}
		if params.EndDate != "" {
			endpoint += "endDate=" + params.EndDate + "&"
		}
		if params.Provider != "" {
			endpoint += "provider=" + params.Provider
		}
	}

	var result UsageStats
	err := s.client.request("GET", endpoint, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Costs retrieves cost breakdown
func (s *AnalyticsService) Costs(startDate, endDate string) (map[string]interface{}, error) {
	endpoint := "/analytics/costs"
	if s.client.BaseURL == "https://api.synqly.com/api/v1" && endpoint[:4] != "/api" {
		endpoint = "/api/v1" + endpoint
	}
	if startDate != "" || endDate != "" {
		endpoint += "?"
		if startDate != "" {
			endpoint += "startDate=" + startDate + "&"
		}
		if endDate != "" {
			endpoint += "endDate=" + endDate
		}
	}

	var result map[string]interface{}
	err := s.client.request("GET", endpoint, nil, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Latency retrieves latency statistics
func (s *AnalyticsService) Latency(provider string) (map[string]interface{}, error) {
	endpoint := "/analytics/latency"
	if s.client.BaseURL == "https://api.synqly.com/api/v1" && endpoint[:4] != "/api" {
		endpoint = "/api/v1" + endpoint
	}
	if provider != "" {
		endpoint += "?provider=" + provider
	}

	var result map[string]interface{}
	err := s.client.request("GET", endpoint, nil, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// KeysService handles API key operations
type KeysService struct {
	client *Client
}

// CreateKeyParams holds parameters for creating an API key
type CreateKeyParams struct {
	Email         string                   `json:"email"`
	Name          string                   `json:"name"`
	KeyName       string                   `json:"keyName"`
	Providers     []map[string]interface{} `json:"providers"`
	FallbackOrder []string                 `json:"fallbackOrder,omitempty"`
}

// Create creates a new API key
func (s *KeysService) Create(params CreateKeyParams) (map[string]interface{}, error) {
	var result map[string]interface{}
	endpoint := "/keys/create"
	if s.client.BaseURL == "https://api.synqly.com/api/v1" && endpoint[:4] != "/api" {
		endpoint = "/api/v1" + endpoint
	}
	err := s.client.request("POST", endpoint, params, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// List retrieves all API keys for a user
func (s *KeysService) List(email string) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("/keys/list?email=%s", email)
	if s.client.BaseURL == "https://api.synqly.com/api/v1" && endpoint[:4] != "/api" {
		endpoint = "/api/v1" + endpoint[:]
	}

	var result map[string]interface{}
	err := s.client.request("GET", endpoint, nil, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Delete deletes an API key
func (s *KeysService) Delete(keyID, email string) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("/keys/%s", keyID)
	if s.client.BaseURL == "https://api.synqly.com/api/v1" && endpoint[:4] != "/api" {
		endpoint = "/api/v1" + endpoint[:]
	}
	body := map[string]string{"email": email}

	var result map[string]interface{}
	err := s.client.request("DELETE", endpoint, body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
