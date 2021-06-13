package paperswithcode_go

import (
	"encoding/json"
	"fmt"
	"github.com/codingpot/paperswithcode-go/v2/models"
)

type MethodListParams struct {
	// Page (default: 1)
	Page int
	// ItemsPerPage (default: 50)
	ItemsPerPage int
}

func (m MethodListParams) String() string {
	return fmt.Sprintf("page=%d&items_per_page=%d", m.Page, m.ItemsPerPage)
}

func (c *Client) MethodList(params MethodListParams) (*models.MethodList, error) {
	url := c.baseURL + "/methods?" + params.String()

	response, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	var listResult models.MethodList

	err = json.NewDecoder(response.Body).Decode(&listResult)
	if err != nil {
		return nil, err
	}

	return &listResult, nil
}