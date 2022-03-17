package paperswithcode_go

import (
	"github.com/codingpot/paperswithcode-go/v2/models"
	"net/url"
	"strconv"
)

// MethodListParams is an argument object to MethodList.
type MethodListParams struct {
	// Page (default: 1)
	Page int
	// ItemsPerPage (default: 50)
	ItemsPerPage int
}

// Build used to expand `MethodListParams` to URL path params.
func (m MethodListParams) Build() string {
	return url.Values{
		"page":           []string{strconv.Itoa(m.Page)},
		"items_per_page": []string{strconv.Itoa(m.ItemsPerPage)},
	}.Encode()
}

// MethodList fetches a list of "methods" that can be used in research papers.
func (c *Client) MethodList(params MethodListParams) (models.MethodList, error) {
	u := c.baseURL + "/methods?" + params.Build()
	var listResult models.MethodList
	err := c.sendGetRequest(u, &listResult)
	return listResult, err
}
