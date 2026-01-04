package synqly

// Message represents a chat message
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatCreateParams holds parameters for creating a chat completion
type ChatCreateParams struct {
	Provider    string    `json:"provider,omitempty"`
	Model       string    `json:"modelType"`
	Messages    []Message `json:"messages"`
	Temperature *float64  `json:"temperature,omitempty"`
	MaxTokens   *int      `json:"max_tokens,omitempty"`
	TopP        *float64  `json:"top_p,omitempty"`
}

// ChatResponse is the response from a chat completion
type ChatResponse struct {
	ID        string `json:"id"`
	Provider  string `json:"provider"`
	ModelType string `json:"modelType"`
	Content   string `json:"content"`
	Usage     struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	FinishReason string `json:"finish_reason"`
	CreatedAt    string `json:"created_at"`
}

// APIResponse wraps the API response
type APIResponse struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    *ChatResponse `json:"data,omitempty"`
}

// APIError represents an API error
type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return e.Message
}
