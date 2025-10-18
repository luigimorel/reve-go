package revego_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	revego "github.com/luigimorel/reve-go"
)

func TestCreateImage(t *testing.T) {
	expectedResponse := revego.CreateImageResponse{
		Image:            "a cute cat",
		ContentViolation: false,
		RequestID:        "req_123",
		Version:          "v1",
		CreditsUsed:      655,
		CreditsRemaining: 1,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/image/create" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer test_api_key" {
			t.Errorf("missing or incorrect Authorization header")
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(expectedResponse)
	}))

	defer server.Close()

	client := &revego.Client{
		APIKey:     "api_key",
		APIBaseURL: server.URL,
		HTTPClient: server.Client(),
	}

	req := revego.CreateImageRequest{
		Prompt:      "a cute cat",
		AspectRatio: "1:1",
		Version:     "v1",
	}

	resp, err := client.CreateImage(context.Background(), req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if resp.Image != expectedResponse.Image {
		t.Errorf("expected image %s, got %s", expectedResponse.Image, resp.Image)
	}
	if resp.CreditsRemaining != expectedResponse.CreditsRemaining {
		t.Errorf(
			"expected credits remaining %d, got %d",
			expectedResponse.CreditsRemaining,
			resp.CreditsRemaining,
		)
	}
	if resp.RequestID != expectedResponse.RequestID {
		t.Errorf("expected request ID %s, got %s", expectedResponse.RequestID, resp.RequestID)
	}
}
