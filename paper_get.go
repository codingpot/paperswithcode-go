package paperswithcode_go

import (
	"encoding/json"
	"fmt"
	"github.com/codingpot/paperswithcode-go/models"
	"net/url"
)

// PaperGet returns a single paper. Note that paperID is hyphen cased (e.g., generative-adversarial-networks).
func (c *Client) PaperGet(paperID string) (*models.Paper, error) {
	paperGetURL := fmt.Sprintf("%s/papers/%s/", c.BaseURL, url.QueryEscape(paperID))
	response, err := c.HTTPClient.Get(paperGetURL)
	if err != nil {
		return nil, err
	}

	var paperGetResult models.Paper
	err = json.NewDecoder(response.Body).Decode(&paperGetResult)
	if err != nil {
		return nil, err
	}

	return &paperGetResult, nil
}
