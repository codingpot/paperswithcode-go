package paperswithcode_go

import (
	"github.com/codingpot/paperswithcode-go/v2/models"
	"net/url"
	"strconv"
)

// PaperList returns multiple papers.
func (c *Client) PaperList(params PaperListParams) (models.PaperList, error) {
	papersListURL := c.baseURL + "/papers?" + params.Build()
	return getJson[models.PaperList](c, papersListURL)
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
	b := url.Values{}
	b.Set("page", strconv.Itoa(p.Page))
	b.Set("items_per_page", strconv.Itoa(p.ItemsPerPage))

	addParamsIfValid(b, "q", p.Q)
	addParamsIfValid(b, "arxiv_id", p.ArxivID)
	addParamsIfValid(b, "title", p.Title)
	addParamsIfValid(b, "abstract", p.Abstract)

	return b.Encode()
}

func addParamsIfValid(b url.Values, key string, value string) {
	if value != "" {
		b.Set(key, value)
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
