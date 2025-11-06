package revego

import (
	"fmt"
	"net/http"
)

type APIError struct {
	StatusCode int
	Body       any
}

func (e *APIError) Error() string {
	return fmt.Sprintf("reve api error: status=%d body=%s", e.StatusCode, e.Body)
}

func SetHeaders(req *http.Request, apiKey string) error {
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	return nil
}
