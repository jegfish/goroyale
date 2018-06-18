package goroyale

import (
	"fmt"
	"time"
)

// APIError represents an error returned from the API.
// https://docs.royaleapi.com/#/errors
type APIError struct {
	StatusCode int    `json:"status"` // http reponse code
	Message    string // human readable message explaining the error
}

func (err APIError) Error() string {
	return err.Message
}

// RatelimitError is returned when it is detected that you will hit the ratelimit.
type RatelimitError struct {
	RetryAfter time.Duration
}

func (err RatelimitError) Error() string {
	return fmt.Sprintf("ratelimit, retry in: %v", err.RetryAfter)
}

func newRatelimitError(retryAfter time.Duration) *RatelimitError {
	err := &RatelimitError{
		RetryAfter: retryAfter,
	}
	return err
}
