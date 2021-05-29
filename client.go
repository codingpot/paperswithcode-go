package paperswithcode_go

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/codingpot/paperswithcode-go/internal/transport"
	"net/http"
	"time"
)

const (
	BaseURL = "https://paperswithcode.com/api/v1"
)

// ClientOption can be used to swap the default http client or swap the API key
type ClientOption func(*Client)

// WithAPIToken sets the client API token.
func WithAPIToken(apiToken string) ClientOption {
	return func(client *Client) {
		client.apiToken = apiToken
		client.HTTPClient.Transport = transport.NewTransportWithAuthHeader(apiToken)
	}
}

// NewClient creates a Client object.
func NewClient(opts ...ClientOption) *Client {
	defaultClient := &Client{
		BaseURL: BaseURL,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}

	for _, opt := range opts {
		opt(defaultClient)
	}

	return defaultClient
}

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	apiToken   string
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	// fmt.Println(fullResponse)

	return nil
}
