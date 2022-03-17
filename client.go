package paperswithcode_go

import (
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/codingpot/paperswithcode-go/v2/internal/transport"
)

const (
	BaseURL = "https://paperswithcode.com/api/v1"
)

var whiteSpaceRegexp = regexp.MustCompile(`\s+`)

// ClientOption can be used to swap the default http client or swap the API key
type ClientOption func(*Client)

// WithAPIToken sets the client API token.
func WithAPIToken(apiToken string) ClientOption {
	return func(client *Client) {
		client.apiToken = apiToken
		client.httpClient.Transport = transport.NewTransportWithAuthHeader(apiToken)
	}
}

// NewClient creates a Client object.
func NewClient(opts ...ClientOption) Client {
	defaultClient := Client{
		baseURL: BaseURL,
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}

	for _, opt := range opts {
		opt(&defaultClient)
	}

	return defaultClient
}

type Client struct {
	baseURL    string
	httpClient http.Client
	apiToken   string
}

// GetPaperIDFromPaperTitle generates a paper ID from paper title.
// WARNING: This function does not cover all cases.
func GetPaperIDFromPaperTitle(paperTitle string) string {
	return strings.ToLower(whiteSpaceRegexp.ReplaceAllString(paperTitle, "-"))
}
