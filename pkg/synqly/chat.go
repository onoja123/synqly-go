package synqly

// ChatService handles chat operations
type ChatService struct {
	client *Client
}

// Create creates a chat completion
func (s *ChatService) Create(params ChatCreateParams) (*ChatResponse, error) {
	if params.Model == "" {
		return nil, &APIError{StatusCode: 400, Message: "model is required"}
	}
	if len(params.Messages) == 0 {
		return nil, &APIError{StatusCode: 400, Message: "messages cannot be empty"}
	}

	var apiResp APIResponse
	err := s.client.request("POST", "/api/v1/chat/create-chat", params, &apiResp)
	if err != nil {
		return nil, err
	}

	if !apiResp.Success {
		return nil, &APIError{StatusCode: 500, Message: apiResp.Message}
	}

	return apiResp.Data, nil
}

// GetContent is a helper to get the response text
func (r *ChatResponse) GetContent() string {
	return r.Content
}
