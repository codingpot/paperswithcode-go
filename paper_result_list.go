package paperswithcode_go

import (
	"encoding/json"
	"fmt"
	"github.com/codingpot/paperswithcode-go/models"
	"net/url"
)

func (c *Client) PaperResultList(paperID string) (*models.ResultList, error) {
	pURL := fmt.Sprintf("%s/papers/%s/results", c.BaseURL, url.QueryEscape(paperID))

	resp, err := c.HTTPClient.Get(pURL)
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
