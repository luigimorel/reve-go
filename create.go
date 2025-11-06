package revego

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CreateImageRequest struct {
	Prompt      string `json:"prompt"`
	AspectRatio string `json:"aspect_ratio"`
	Version     string `json:"version"`
}

type ImageResponse struct {
	Image            string `json:"image"`
	ContentViolation bool   `json:"content_violation"`
	RequestID        string `json:"request_id"`
	Version          string `json:"version"`
	CreditsUsed      int    `json:"credits_used"`
	CreditsRemaining int    `json:"credits_remaining"`
}

// Create images from a text description. To learn more about the pricing for this endpoint, visit the pricing page.
func (c *Client) CreateImage(ctx context.Context, req CreateImageRequest) (*ImageResponse, error) {
	endpoint := c.APIBaseURL + "/image/create"
	b, err := json.Marshal(&req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("failed to send request with context: %w", err)
	}

	SetHeaders(httpReq, c.APIKey)

	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, &APIError{StatusCode: resp.StatusCode, Body: string(body)}
	}

	var out ImageResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return &out, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &out, nil
}
