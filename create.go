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

type CreateImageResponse struct {
	Image            string `json:"image"`
	ContentViolation bool   `json:"content_violation"`
	RequestID        string `json:"request_id"`
	Version          string `json:"version"`
	CreditsUsed      int    `json:"credits_used"`
	CreditsRemaining int    `json:"credits_remaining"`
}

type APIError struct {
	StatusCode int
	Body       any
}

func (e *APIError) Error() string {
	return fmt.Sprintf("reve api error: status=%d body=%s", e.StatusCode, e.Body)
}

func (c *Client) CreateImage(
	ctx context.Context,
	req CreateImageRequest,
) (*CreateImageResponse, error) {
	endpoint := c.APIBaseURL + "/image/create"
	b, err := json.Marshal(&req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		endpoint,
		bytes.NewReader(b),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to send request with context: %w", err)
	}

	httpReq.Header.Set("Authorization", "Bearer "+c.APIKey)
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Content-Type", "application/json")

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

	var out CreateImageResponse
	if err := json.Unmarshal(body, &out); err != nil {
		return &out, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &out, nil
}
