package paperswithcode_go

import (
	"fmt"
	"github.com/codingpot/paperswithcode-go/v2/models"
	"net/url"
)

// PaperRepositoryList returns repositories related to the given paper.
func (c *Client) PaperRepositoryList(paperID string) (models.RepositoryList, error) {
	paperURL := fmt.Sprintf("%s/papers/%s/repositories", c.baseURL, url.QueryEscape(paperID))
	var repoList models.RepositoryList
	err := c.sendGetRequest(paperURL, &repoList)
	return repoList, err
}
