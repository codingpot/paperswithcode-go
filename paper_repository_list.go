package paperswithcode_go

import (
	"encoding/json"
	"fmt"
	"github.com/codingpot/paperswithcode-go/models"
	"net/url"
)

func (c *Client) PaperRepositoryList(paperID string) (*models.RepositoryList, error) {
	paperURL := fmt.Sprintf("%s/papers/%s/repositories", c.BaseURL, url.QueryEscape(paperID))
	response, err := c.HTTPClient.Get(paperURL)
	if err != nil {
		return nil, err
	}

	var repoList models.RepositoryList
	err = json.NewDecoder(response.Body).Decode(&repoList)
	if err != nil {
		return nil, err
	}

	return &repoList, nil
}
