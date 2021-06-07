package paperswithcode_go

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/codingpot/paperswithcode-go/v2/models"
)

// PaperList returns multiple papers.
func (c *Client) PaperList(params PaperListParams) (*models.PaperList, error) {
	papersListURL := c.baseURL + "/papers?" + params.Build()

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
	// Q to search papers (default: "")
	// If empty, it returns all papers.
	Q        string
	ArxivID  string
	Title    string
	Abstract string
	// Page is the number of page to search (default: 1)
	Page int
	// ItemsPerPage returns how many papers are returned in a single response.
	ItemsPerPage int
}

func (p PaperListParams) Build() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("page=%d&items_per_page=%d", p.Page, p.ItemsPerPage))

	addParamsIfValid(&b, "q", p.Q)
	addParamsIfValid(&b, "arxiv_id", p.ArxivID)
	addParamsIfValid(&b, "title", p.Title)
	addParamsIfValid(&b, "abstract", p.Abstract)

	return b.String()
}

func addParamsIfValid(b *strings.Builder, key string, value string) {
	if value != "" {
		b.WriteString(fmt.Sprintf("&%s=%s", key, value))
	}
}

// PaperListParamsDefault returns the default PaperListParams.
func PaperListParamsDefault() PaperListParams {
	return PaperListParams{
		Q:            "",
		Page:         1,
		ItemsPerPage: 50,
	}
}
