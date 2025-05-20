package globalexchangerates

import (
	"fmt"
)

// APIError represents an error returned by the Global Exchange Rates API.
type APIError struct {
	// StatusCode is the HTTP status code of the failed API request.
	StatusCode int
	
	// ErrorCode is the error code returned by the API, if available.
	ErrorCode int
	
	// Message is the error message returned by the API, if available.
	Message string
}

// Error implements the error interface for APIError.
func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("API request failed with status code %d: %s (error code: %d)", 
			e.StatusCode, e.Message, e.ErrorCode)
	}
	return fmt.Sprintf("API request failed with status code %d", e.StatusCode)
}

// IsAPIError checks if an error is an APIError.
func IsAPIError(err error) (*APIError, bool) {
	apiErr, ok := err.(*APIError)
	return apiErr, ok
}
