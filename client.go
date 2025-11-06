package revego

import (
	"errors"
	"net/http"
	"time"
)

var APIBaseURL = "https://api.reve.com/v1/"

type Client struct {
	ClientID   string
	HTTPClient *http.Client
	APIKey     string
	APIBaseURL string
}

func NewClient(apiKey string) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("API Key is missing")
	}

	client := &Client{
		APIKey:     apiKey,
		APIBaseURL: APIBaseURL,
		HTTPClient: &http.Client{Timeout: 60 * time.Second},
	}

	return client, nil
}
