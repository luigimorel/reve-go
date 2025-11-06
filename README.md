# reve-go

Go client library for the [Reve.com](https://reve.com) image generation API.

## Installation

```bash
go get github.com/luigimorel/reve-go
```

## Quick Start

```go
package main

import (
    "context"
    "log"

    revego "github.com/luigimorel/reve-go"
)

func main() {
    client, err := revego.NewClient("your-api-key")
    if err != nil {
        log.Fatal(err)
    }

    resp, err := client.CreateImage(context.Background(), revego.CreateImageRequest{
        Prompt:      "A beautiful sunset over mountains",
        AspectRatio: "16:9",
    })
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Generated image: %s", resp.Image)
    log.Printf("Credits remaining: %d", resp.CreditsRemaining)
}
```

## API Methods

### Create Image

Generate images from text descriptions.

```go
resp, err := client.CreateImage(ctx, revego.CreateImageRequest{
    Prompt:      "A cat sitting on a windowsill",
    AspectRatio: "1:1",        // Optional: "1:1", "16:9", etc.
    Version:     "latest",     // Optional: API version
})
```

### Edit Image

Modify existing images using text instructions.

```go
resp, err := client.EditImage(ctx, revego.EditImageRequest{
    EditInstruction: "Make the sky more blue",
    ReferenceImage:  base64EncodedImage, // Base64 encoded image data
    Version:         "latest",           // Optional: API version
})
```

### Remix Image

Create new images using text prompts and reference images.

```go
resp, err := client.RemixImage(ctx, revego.RemixImageRequest{
    Prompt:          "A futuristic version of this scene",
    ReferenceImages: []string{base64Image1, base64Image2}, // Base64 encoded images
    AspectRatio:     "16:9",                               // Optional
    Version:         "latest",                             // Optional
})
```

## Response Structure

All image generation methods return an `ImageResponse`:

```go
type ImageResponse struct {
    Image            string `json:"image"`             // Generated image (base64 or URL)
    ContentViolation bool   `json:"content_violation"` // Content policy violation flag
    RequestID        string `json:"request_id"`        // Unique request identifier
    Version          string `json:"version"`           // API version used
    CreditsUsed      int    `json:"credits_used"`      // Credits consumed
    CreditsRemaining int    `json:"credits_remaining"` // Credits left in account
}
```

## Error Handling

The client returns structured API errors:

```go
resp, err := client.CreateImage(ctx, req)
if err != nil {
    if apiErr, ok := err.(*revego.APIError); ok {
        log.Printf("API error: status=%d, body=%s", apiErr.StatusCode, apiErr.Body)
    } else {
        log.Printf("Client error: %v", err)
    }
    return
}
```

## Configuration

The client uses sensible defaults but can be customized:

```go
client := &revego.Client{
    APIKey:     "your-api-key",
    APIBaseURL: "https://api.reve.com/v1/", // Default
    HTTPClient: &http.Client{Timeout: 60 * time.Second}, // Default timeout
}
```

## License

This project is licensed under the MIT [License](./LICENSE).
