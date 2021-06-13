package paperswithcode_go

import (
	"encoding/json"
	"github.com/codingpot/paperswithcode-go/v2/models"
)

// MethodGet returns a method in a paper.
// See https://paperswithcode-client.readthedocs.io/en/latest/api/client.html#paperswithcode.client.PapersWithCodeClient.method_list
func (c *Client) MethodGet(methodID string) (*models.Method, error) {
	url := c.baseURL + "/methods/" + methodID

	response, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	var result models.Method

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}