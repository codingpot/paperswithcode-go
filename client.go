package paperswithcode_go

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/codingpot/paperswithcode-go/models"
	"net/http"
	"net/url"
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

// Paper represents a specific Paper's Information by the title
type Paper struct {
	ID        string   `json:"id"`
	ArxivID   string   `json:"arxiv_id,omitempty"`
	URLAbs    string   `json:"url_abs"`
	URLPDF    string   `json:"url_pdf"`
	Title     string   `json:"title"`
	Abstract  string   `json:"abstract"`
	Authors   []string `json:"authors"`
	Published string   `json:"published"`
}

func (c *Client) GetPaper(ctx context.Context, id string) (*Paper, error) {
	fmt.Println(id)
	url := fmt.Sprintf("%s/papers/%s/", c.BaseURL, url.QueryEscape(id))
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := Paper{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Method list used by Paper's ID
type MethodList struct {
	Count   int      `json:"count"`
	Methods []Method `json:"results"`
}

type Method struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	Paper       string `json:"paper"`
}

func (c *Client) GetMethodList(ctx context.Context, id string) (*MethodList, error) {
	fmt.Println(id)
	url := fmt.Sprintf("%s/papers/%s/methods", c.BaseURL, url.QueryEscape(id))
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := MethodList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Repository List implementing by Paper's ID
type RepositoryList struct {
	Count   int          `json:"count"`
	Results []Repository `json:"results"`
}

type Repository struct {
	URL         string `json:"url"`
	IsOfficial  bool   `json:"is_official"`
	Description string `json:"description"`
	Stars       int    `json:"stars"`
	Framework   string `json:"framework"`
}

func (c *Client) GetRepositoryList(ctx context.Context, id string) (*RepositoryList, error) {
	fmt.Println(id)
	url := fmt.Sprintf("%s/papers/%s/repositories", c.BaseURL, url.QueryEscape(id))
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := RepositoryList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
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

type PaperListParams struct {
	Query string
	Page  int
	Limit int
}

func (p PaperListParams) Build() string {
	if p.Query == "" {
		return fmt.Sprintf("items_per_page=%d&page=%d", p.Limit, p.Page)
	}
	return fmt.Sprintf("q=%s&items_per_page=%d&page=%d", url.QueryEscape(p.Query), p.Limit, p.Page)
}

func (c *Client) PaperList(params PaperListParams) (*models.PaperListResult, error) {
	papersListURL := c.BaseURL + "/papers?" + params.Build()

	request, err := http.NewRequest(http.MethodGet, papersListURL, nil /*body*/)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", "Token "+c.apiToken)

	response, err := c.HTTPClient.Get(papersListURL)
	if err != nil {
		return nil, err
	}

	var paperListResult models.PaperListResult

	err = json.NewDecoder(response.Body).Decode(&paperListResult)
	if err != nil {
		return nil, err
	}

	return &paperListResult, nil
}

func PaperListParamsDefault() PaperListParams {
	return PaperListParams{
		Query: "",
		Page:  1,
		Limit: 50,
	}
}
