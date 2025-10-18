package revego

import "fmt"

type APIError struct {
	StatusCode int
	Body       any
}

func (e *APIError) Error() string {
	return fmt.Sprintf("reve api error: status=%d body=%s", e.StatusCode, e.Body)
}
