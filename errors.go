package synqly

import "fmt"

// APIError represents an API error
type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("synqly API error (status %d): %s", e.StatusCode, e.Message)
}
