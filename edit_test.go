package revego_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	revego "github.com/luigimorel/reve-go"
)

func TestEditImage(t *testing.T) {
	expectedResponse := revego.ImageResponse{
		Image:            "an edited cute cat",
		ContentViolation: false,
		RequestID:        "req_456",
		Version:          "v1",
		CreditsUsed:      300,
		CreditsRemaining: 5,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/image/edit" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer "+TEST_API_KEY {
			t.Errorf("missing or incorrect Authorization header")
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(expectedResponse)
	}))

	defer server.Close()

	client := &revego.Client{
		APIKey:     TEST_API_KEY,
		APIBaseURL: server.URL,
		HTTPClient: server.Client(),
	}

	resp, err := client.EditImage(context.Background(), revego.EditImageRequest{
		EditInstruction: "Make the cat wear a hat",
		ReferenceImage:  "base64encodedstring",
		Version:         "v1",
	})
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
	if resp.ContentViolation != expectedResponse.ContentViolation {
		t.Errorf(
			"expected content violation %v, got %v",
			expectedResponse.ContentViolation,
			resp.ContentViolation,
		)
	}
}
