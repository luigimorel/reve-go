package revego

import (
	"errors"
	"net/http"
	"sync"
	"time"
)

var (
	APIBaseURL     = "https://api.reve.com/v1/"
	DefaultTimeout = 30 * time.Second
)

type Client struct {
	clientID    string
	HTTPClient  *http.Client
	APIKey      string
	mu          sync.Mutex
	apiBaseURL  string
	tokenExpiry time.Time
	Token       string
}

func NewClient(apiKey string) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("API Key is missing")
	}

	return &Client{
		APIKey: apiKey,
		HTTPClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}, nil
}

func (c *Client) SetBaseURL(url string) {
	c.apiBaseURL = url
}
