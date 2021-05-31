package paperswithcode_go

import (
	"encoding/json"
	"fmt"
	"github.com/codingpot/paperswithcode-go/v2/models"
	"net/url"
)

// PaperResultList returns the evaluation results of the paper.
func (c *Client) PaperResultList(paperID string) (*models.ResultList, error) {
	pURL := fmt.Sprintf("%s/papers/%s/results", c.baseURL, url.QueryEscape(paperID))

	resp, err := c.httpClient.Get(pURL)
	if err != nil {
		return nil, err
	}

	var paperResultList models.ResultList
	err = json.NewDecoder(resp.Body).Decode(&paperResultList)
	if err != nil {
		return nil, err
	}

	return &paperResultList, nil
}
