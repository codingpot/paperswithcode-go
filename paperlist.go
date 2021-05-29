package paperswithcode_go

import (
	"encoding/json"
	"fmt"
	"github.com/codingpot/paperswithcode-go/models"
	"net/url"
)

func (c *Client) PaperList(params PaperListParams) (*models.PaperListResult, error) {
	papersListURL := c.BaseURL + "/papers?" + params.Build()

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

func PaperListParamsDefault() PaperListParams {
	return PaperListParams{
		Query: "",
		Page:  1,
		Limit: 50,
	}
}
