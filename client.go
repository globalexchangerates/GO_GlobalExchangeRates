package globalexchangerates

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	// BaseURL is the base URL for the Global Exchange Rates API.
	BaseURL = "https://api.globalexchangerates.org/v1"
)

// Client is the client for accessing the Global Exchange Rates API.
type Client struct {
	// httpClient is the HTTP client used to make requests.
	httpClient *http.Client
	
	// apiKey is the API key used to authenticate with the Global Exchange Rates API.
	apiKey string
}

// ClientOption represents a configuration option for the Client.
type ClientOption func(*Client)

// WithHTTPClient sets a custom HTTP client for the Client.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// NewClient creates a new Client instance with the provided API key.
func NewClient(apiKey string, options ...ClientOption) *Client {
	if apiKey == "" {
		panic("apiKey cannot be empty")
	}
	
	client := &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		apiKey: apiKey,
	}
	
	// Apply options
	for _, option := range options {
		option(client)
	}
	
	return client
}

// sendRequest sends a request to the Global Exchange Rates API and processes the response.
func (c *Client) sendRequest(ctx context.Context, endpoint string, params url.Values, v interface{}) error {
	// Construct the URL
	reqURL, err := url.Parse(BaseURL + endpoint)
	if err != nil {
		return err
	}
	
	// Add query parameters
	if params != nil {
		reqURL.RawQuery = params.Encode()
	}
	
	// Create the request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return err
	}
	
	// Add headers
	req.Header.Set("Subscription-Key", c.apiKey)
	req.Header.Set("X-Source", "GOLANG")
	req.Header.Set("Accept", "application/json")
	
	// Send the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	
	// Check for errors
	if resp.StatusCode >= 400 {
		apiErr := &APIError{
			StatusCode: resp.StatusCode,
		}
		
		var errResp ErrorResponse
		if err := json.Unmarshal(body, &errResp); err == nil {
			apiErr.ErrorCode = errResp.ErrorCode
			apiErr.Message = errResp.Message
		}
		
		return apiErr
	}
	
	// Unmarshal the response
	return json.Unmarshal(body, v)
}
