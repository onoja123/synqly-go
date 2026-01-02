# Synqly Go SDK

Official Go client for Synqly API - Unified LLM Gateway

## Installation
```bash
go get github.com/onoja123/synqly-go
```

## Quick Start
```go
package main

import (
    "fmt"
    "github.com/onoja123/synqly-go/pkg/synqly"
)

func main() {
    logger := synqly.Logger{}
    client := synqly.NewClient(synqly.Config{
        APIKey: "sk_synqly_your_key_here",
    })

    response, err := client.Chat.Create(synqly.ChatCreateParams{
        Provider: "anthropic",
        Model:    "claude-sonnet-4",
        Messages: []synqly.Message{
            {Role: "user", Content: "Hello!"},
        },
    })

    if err != nil {
        logger.Fatal(err)
    }

    logger.Printf("Response: %s", response.GetContent())
}
```

## Usage

### Basic Chat
```go
response, err := client.Chat.Create(synqly.ChatCreateParams{
    Model: "gpt-4",
    Messages: []synqly.Message{
        {Role: "user", Content: "What is AI?"},
    },
})
```

### With Options
```go
temp := 0.7
maxTokens := 500

response, err := client.Chat.Create(synqly.ChatCreateParams{
    Provider:    "openai",
    Model:       "gpt-4",
    Messages:    []synqly.Message{
        {Role: "system", Content: "You are helpful."},
        {Role: "user", Content: "Explain quantum computing."},
    },
    Temperature: &temp,
    MaxTokens:   &maxTokens,
})
```

### Multi-turn Conversation
```go
response, err := client.Chat.Create(synqly.ChatCreateParams{
    Model: "gpt-4",
    Messages: []synqly.Message{
        {Role: "user", Content: "Hi!"},
        {Role: "assistant", Content: "Hello! How can I help?"},
        {Role: "user", Content: "Tell me a joke."},
    },
})
```

## Error Handling
```go
response, err := client.Chat.Create(params)
if err != nil {
    if apiErr, ok := err.(*synqly.APIError); ok {
        fmt.Printf("API Error: %s (status: %d)\n", apiErr.Message, apiErr.StatusCode)
    } else {
        fmt.Printf("Error: %v\n", err)
    }
    return
}
```

# synqly-go

A simple Go SDK for Synqly's chat API.

## Installation

```
go get github.com/onoja123/synqly-go/pkg/synqly
```

## Usage

```go
package main

import (
    "github.com/onoja123/synqly-go/pkg/synqly"
)

func main() {
    logger := synqly.Logger{}
    client := synqly.NewClient(synqly.Config{
        APIKey: "sk_synqly_your_key_here",
    })

    response, err := client.Chat.Create(synqly.ChatCreateParams{
        Provider: "anthropic",
        Model:    "claude-sonnet-4",
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
```

- No need to set the BaseURL for normal usage.
- The SDK will use the correct production endpoint by default.
- Use the built-in logger for consistent output.

## License

MIT