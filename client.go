package paperswithcode_go

import (
	"github.com/codingpot/paperswithcode-go/internal/transport"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const (
	BaseURL = "https://paperswithcode.com/api/v1"
)

var whiteSpaceRegexp = regexp.MustCompile("\\s+")

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

func GetPaperIDFromPaperTitle(paperTitle string) string {
	return strings.ToLower(whiteSpaceRegexp.ReplaceAllString(paperTitle, "-"))
}
