package paperswithcode_go

import (
	"encoding/json"
	"fmt"
	"github.com/codingpot/paperswithcode-go/models"
	"net/url"
)

// PaperRepositoryList returns repositories related to the given paper.
func (c *Client) PaperRepositoryList(paperID string) (*models.RepositoryList, error) {
	paperURL := fmt.Sprintf("%s/papers/%s/repositories", c.baseURL, url.QueryEscape(paperID))
	response, err := c.httpClient.Get(paperURL)
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
