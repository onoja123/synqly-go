package main

import (
	"os"

	"github.com/onoja123/synqly-go/pkg/synqly"
)

func main() {
	apiKey := os.Getenv("SYNQLY_API_KEY")
	logger := synqly.Logger{}

	if apiKey == "" {
		logger.Fatal("missing SYNQLY_API_KEY environment variable")
	}

	client := synqly.NewClient(synqly.Config{
		APIKey: apiKey,
	})

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

	logger.Printf("%s", response.GetContent())

	temp := 0.7
	maxTokens := 100

	response, err = client.Chat.Create(synqly.ChatCreateParams{
		Provider: "openai",
		Model:    "gpt-4",
		Messages: []synqly.Message{
			{Role: "system", Content: "You are a helpful assistant."},
			{Role: "user", Content: "What is 2+2?"},
		},
		Temperature: &temp,
		MaxTokens:   &maxTokens,
	})

	if err != nil {
		logger.Fatal(err)
	}

	logger.Printf("%s", response.GetContent())
	logger.Printf("Tokens used: %d", response.Usage.TotalTokens)
}
