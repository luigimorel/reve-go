package revego

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RemixImageRequest struct {
	Prompt          string   `json:"prompt"`
	ReferenceImages []string `json:"reference_images"` // Base64 encoded image data
	AspectRatio     string   `json:"aspect_ratio"`
	Version         string   `json:"version"`
}

func (c *Client) RemixImage(ctx context.Context, req RemixImageRequest) (*ImageResponse, error) {
	endpoint := c.APIBaseURL + "/image/remix"
	b, err := json.Marshal(&req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal the request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("failed to send request with context: %w", err)
	}

	SetHeaders(httpReq, c.APIKey)

	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
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
		return &out, fmt.Errorf("failed to unmarshal json response: %w", err)
	}

	return &out, nil
}
