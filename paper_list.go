package paperswithcode_go

import (
	"encoding/json"
	"fmt"
	"github.com/codingpot/paperswithcode-go/v2/models"
	"net/url"
)

// PaperList returns multiple papers.
func (c *Client) PaperList(params PaperListParams) (*models.PaperList, error) {
	papersListURL := c.baseURL + "/papers?" + params.build()

	response, err := c.httpClient.Get(papersListURL)
	if err != nil {
		return nil, err
	}

	var paperListResult models.PaperList

	err = json.NewDecoder(response.Body).Decode(&paperListResult)
	if err != nil {
		return nil, err
	}

	return &paperListResult, nil
}

// PaperListParams is the parameter for PaperList method.
type PaperListParams struct {
	// Query to search papers (default: "")
	// If empty, it returns all papers.
	Query string
	// Page is the number of page to search (default: 1)
	Page int
	// Limit returns how many papers are returned in a single response.
	Limit int
}

func (p PaperListParams) build() string {
	if p.Query == "" {
		return fmt.Sprintf("items_per_page=%d&page=%d", p.Limit, p.Page)
	}
	return fmt.Sprintf("q=%s&items_per_page=%d&page=%d", url.QueryEscape(p.Query), p.Limit, p.Page)
}

// PaperListParamsDefault returns the default PaperListParams.
func PaperListParamsDefault() PaperListParams {
	return PaperListParams{
		Query: "",
		Page:  1,
		Limit: 50,
	}
}
