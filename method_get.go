package paperswithcode_go

import (
	"github.com/codingpot/paperswithcode-go/v2/models"
)

// MethodGet returns a method in a paper.
// See https://paperswithcode-client.readthedocs.io/en/latest/api/client.html#paperswithcode.client.PapersWithCodeClient.method_list
func (c *Client) MethodGet(methodID string) (models.Method, error) {
	url := c.baseURL + "/methods/" + methodID
	return getJson[models.Method](c, url)
}
