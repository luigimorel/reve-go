package revego_test

import (
	"testing"
	"time"

	revego "github.com/luigimorel/reve-go"
)

var TEST_API_KEY = "papi.6ba7d95b-641c-499f-8cd0-cdb324b42ca5"

func TestNewClientSuccess(t *testing.T) {
	origBase := revego.APIBaseURL
	revego.APIBaseURL = "https://example.test/v2/"
	t.Cleanup(func() { revego.APIBaseURL = origBase })

	apiKey := TEST_API_KEY
	c, err := revego.NewClient(apiKey)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if c == nil {
		t.Fatal("expected client, got nil")
	}

	if c.APIKey != apiKey {
		t.Errorf("expected APIKey %q, got %q", apiKey, c.APIKey)
	}
	if c.APIBaseURL != "https://example.test/v2/" {
		t.Errorf("expected APIBaseURL to follow global var, got %q", c.APIBaseURL)
	}
	if c.HTTPClient == nil {
		t.Fatal("expected HTTPClient to be initialized")
	}
	if c.HTTPClient.Timeout != 60*time.Second {
		t.Errorf("expected HTTP client timeout 60s, got %v", c.HTTPClient.Timeout)
	}
}

func TestNewClientMissingAPIKey(t *testing.T) {
	c, err := revego.NewClient("")
	if err == nil {
		t.Fatal("expected error for missing API key, got nil")
	}
	if c != nil {
		t.Fatalf("expected nil client on error, got %#v", c)
	}
}
