package synqly

// Message represents a chat message
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatCreateParams is the request for creating a chat completion
type ChatCreateParams struct {
	Provider    string    `json:"provider,omitempty"`
	Model       string    `json:"modelType"`
	Messages    []Message `json:"messages"`
	Temperature *float64  `json:"temperature,omitempty"`
	MaxTokens   *int      `json:"max_tokens,omitempty"`
	TopP        *float64  `json:"top_p,omitempty"`
	Fallback    *bool     `json:"fallback,omitempty"`
	Stream      *bool     `json:"stream,omitempty"`
}

// ChatResponse is the API response
type ChatResponse struct {
	ID        string `json:"id"`
	Provider  string `json:"provider"`
	ModelType string `json:"modelType"`
	Content   string `json:"content"`
	Usage     struct {
		PromptTokens        int `json:"prompt_tokens"`
		CompletionTokens    int `json:"completion_tokens"`
		TotalTokens         int `json:"total_tokens"`
		PromptTokensDetails struct {
			CachedTokens int `json:"cached_tokens"`
			AudioTokens  int `json:"audio_tokens"`
		} `json:"prompt_tokens_details"`
		CompletionTokensDetails struct {
			ReasoningTokens          int `json:"reasoning_tokens"`
			AudioTokens              int `json:"audio_tokens"`
			AcceptedPredictionTokens int `json:"accepted_prediction_tokens"`
			RejectedPredictionTokens int `json:"rejected_prediction_tokens"`
		} `json:"completion_tokens_details"`
	} `json:"usage"`
	FinishReason string `json:"finish_reason"`
	CreatedAt    string `json:"created_at"`
	Cached       bool   `json:"cached"`
	CacheHit     bool   `json:"cache_hit"`
}

// APIResponse wraps the response
type APIResponse struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Data    *ChatResponse `json:"data,omitempty"`
}

// ErrorResponse for error cases
type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// APIError for error handling in chat.go
type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return e.Message
}
