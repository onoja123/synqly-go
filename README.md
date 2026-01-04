# Synqly Go SDK

**One API for Every AI Model**

Official Go client for the **Synqly API** — a unified LLM gateway that lets you interact with multiple AI providers (OpenAI, Anthropic, Google, and more) using a single, consistent interface.

---

## Features

* Unified access to OpenAI, Anthropic, Google, and more
* Simple, idiomatic Go API
* Multi-turn conversations
* Configurable parameters (temperature, max tokens, etc.)
* Built-in error handling
* Uses Synqly production endpoints by default
* No vendor lock-in

---

## Installation

```bash
go get github.com/onoja123/synqly-go/pkg/synqly
```

---

## Quick Start

```go
package main

import (
	"github.com/onoja123/synqly-go/pkg/synqly"
)

func main() {
	logger := synqly.Logger{}

	client := synqly.NewClient(synqly.Config{
		APIKey: "YOUR_API_KEY",
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
}
```

---

## Usage

### Basic Chat Completion

```go
response, err := client.Chat.Create(synqly.ChatCreateParams{
	Provider: "openai",
	Model:    "gpt-4",
	Messages: []synqly.Message{
		{Role: "user", Content: "What is the capital of France?"},
	},
})

if err != nil {
	logger.Fatal(err)
}
```

---

### Chat with Parameters

```go
temp := 0.7
maxTokens := 500

response, err := client.Chat.Create(synqly.ChatCreateParams{
	Provider:    "anthropic",
	Model:       "claude-sonnet-4",
	Messages: []synqly.Message{
		{Role: "system", Content: "You are a helpful assistant."},
		{Role: "user", Content: "Explain quantum computing in simple terms."},
	},
	Temperature: &temp,
	MaxTokens:   &maxTokens,
})

if err != nil {
	logger.Fatal(err)
}

logger.Printf("%s", response.GetContent())
```

---

### Switching Providers

```go
// OpenAI
response, _ := client.Chat.Create(synqly.ChatCreateParams{
	Provider: "openai",
	Model:    "gpt-4",
	Messages: messages,
})

// Anthropic
response, _ = client.Chat.Create(synqly.ChatCreateParams{
	Provider: "anthropic",
	Model:    "claude-sonnet-4",
	Messages: messages,
})
```

---

## API Reference

### `synqly.NewClient(config)`

Creates a new Synqly client.

| Field     | Type   | Required | Description                                  |
| --------- | ------ | -------- | -------------------------------------------- |
| `APIKey`  | string | Yes      | Your Synqly API key                          |

---

### `client.Chat.Create(params)`

Creates a chat completion.

| Field         | Type        | Required | Description                                   |
| ------------- | ----------- | -------- | --------------------------------------------- |
| `Provider`    | string      | No       | AI provider (`openai`, `anthropic`, `google`) |
| `Model`       | string      | Yes      | Model name                                    |
| `Messages`    | `[]Message` | Yes      | Conversation messages                         |
| `Temperature` | `*float64`  | No       | Sampling temperature                          |
| `MaxTokens`   | `*int`      | No       | Max tokens in response                        |
| `TopP`        | `*float64`  | No       | Nucleus sampling                              |

---

## Response

```go
type ChatResponse struct {
	ID           string
	Provider     string
	ModelType    string
	Content      string
	Usage        Usage
	FinishReason string
	CreatedAt    string
}

response.GetContent()
```

---

## Supported Providers

| Provider    | Models                                         |
| ----------- | ---------------------------------------------- |
| `openai`    | gpt-4, gpt-4-turbo, gpt-3.5-turbo              |
| `anthropic` | claude-sonnet-4, claude-3-opus, claude-3-haiku |
| `google`    | gemini-pro, gemini-ultra                       |

---

## Error Handling

```go
response, err := client.Chat.Create(params)

if err != nil {

	if apiErr, ok := err.(*synqly.APIError); 

    ok {
		fmt.Printf("API Error: %s (status: %d)\n", apiErr.Message, apiErr.StatusCode)
	} else {
		fmt.Printf("Error: %v\n", err)
	}
}
```

---

## Get an API Key

1. Visit **synqly.xyz**
2. Sign up and create an API key
3. Use the key in your application

---

## Contributing

Contributions are welcome. Feel free to open an issue or submit a pull request.

---

## License

MIT

---
