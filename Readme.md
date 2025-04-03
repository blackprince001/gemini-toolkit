# Gemini Toolkit Documentation

## Overview

The Gemini Toolkit provides a simple interface to interact with Google's Gemini AI models for chat, document processing, and search functionality.

## Installation

```bash
go get github.com/blackprince001/gemini-toolkit
```

## Quick Start

### 1. Import the package

```go
import (
    "github.com/blackprince001/gemini-toolkit/tools"
    "github.com/blackprince001/gemini-toolkit/types"
)
```

### 2. Initialize the service

Create a configuration with your API key and preferred model:

```go
config := types.ToolKitConfig{
    ApiKey:    "YOUR_API_KEY",    // Replace with your actual API key
    BaseModel: "gemini-pro",      // Or your preferred model
}
```

### 3. Use the services

#### Chat Service

```go
chat := tools.NewToolKitChatService(config)

// Simple chat
response, err := chat.GetChatResponse(context.Background(), "Hello Gemini!")
if err != nil {
    // handle error
}
fmt.Println(response)

// Streaming chat
stream, err := chat.GetChatResponseStream(context.Background(), "Hello Gemini!")
if err != nil {
    // handle error
}
// Process the stream...
```

#### Document Service

```go
doc := tools.NewToolKitDocumentService(config)

// Get document summary
summary, err := doc.GetDocumentSummary(context.Background(), "https://example.com/document.pdf")
if err != nil {
    // handle error
}
fmt.Println(summary)
```

## Available Services

### Chat Service

- `GetChatResponse(ctx, prompt)` - Get a single response
- `GetChatResponseStream(ctx, prompt)` - Get a streaming response
- `GetSearchResponse(ctx, prompt)` - Get web-search augmented response

### Document Service

- `GetDocumentSummary(ctx, fileUrl)` - Get summary of a document
- `GetDocumentType(ctx, fileUrl)` - Get category/type of a document

## Configuration Options

The `ToolKitConfig` struct accepts:

- `ApiKey`: Your Gemini API key (required)
- `BaseModel`: The model to use (e.g., "gemini-pro")

## Error Handling

All methods return errors that should be checked. Common errors include:

- Invalid API key
- Network issues
- Invalid document URLs
- Model errors

## Examples

See the `examples/` directory in the package for complete usage examples.

## Dependencies

- Go 1.23+
- google.golang.org/genai v0.7.0

---

This documentation covers just what you need to get started. For advanced usage, refer to the source code and Google's Gemini API documentation.
