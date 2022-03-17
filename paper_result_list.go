package paperswithcode_go

import (
	"fmt"
	"github.com/codingpot/paperswithcode-go/v2/models"
	"net/url"
)

// PaperResultList returns the evaluation results of the paper.
func (c *Client) PaperResultList(paperID string) (models.ResultList, error) {
	pURL := fmt.Sprintf("%s/papers/%s/results", c.baseURL, url.QueryEscape(paperID))
	var paperResultList models.ResultList
	err := c.sendGetRequest(pURL, &paperResultList)
	return paperResultList, err
}
