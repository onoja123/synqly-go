package main

import (
	"github.com/onoja123/synqly-go/pkg/synqly"
)

func main() {
	logger := synqly.Logger{}
	// Create client
	client := synqly.NewClient(synqly.Config{
		APIKey: "sk_synqly_your_key_here",
	})

	// Simple chat
	response, err := client.Chat.Create(synqly.ChatCreateParams{
		Provider: "openai",
		Model:    "gpt-4",
		Messages: []synqly.Message{
			{Role: "user", Content: "Hello!"},
		},
	})

	if err != nil {
		logger.Fatal(err)
	}

	logger.Printf("Response: %s", response.GetContent())
	logger.Printf("Cached: %v", response.Cached)
	logger.Printf("Tokens: %d", response.Usage.TotalTokens)
}
