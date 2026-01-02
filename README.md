# Synqly Go SDK

Official Go client for the **Synqly API** — a unified LLM gateway that lets you interact with multiple AI providers using a single interface.

---

## Features

* Unified access to OpenAI, Anthropic, and more
* Simple, idiomatic Go API
* Multi-turn conversations
* Configurable parameters (temperature, max tokens, etc.)
* Built-in error handling
* Automatic use of Synqly production endpoints

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

---

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

---

### With Options

```go
temp := 0.7
maxTokens := 500

response, err := client.Chat.Create(synqly.ChatCreateParams{
    Provider: "openai",
    Model:    "gpt-4",
    Messages: []synqly.Message{
        {Role: "system", Content: "You are helpful."},
        {Role: "user", Content: "Explain quantum computing."},
    },
    Temperature: &temp,
    MaxTokens:   &maxTokens,
})
```

---

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

---

## Response Metadata

```go
logger.Printf("Response: %s", response.GetContent())
logger.Printf("Cached: %v", response.Cached)
logger.Printf("Tokens Used: %d", response.Usage.TotalTokens)
```

---

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

---

## Notes

* No need to set a `BaseURL` — the SDK uses Synqly’s production endpoint by default
* Supports multiple providers through a single API
* Designed to be lightweight and developer-friendly

---

## License

MIT
