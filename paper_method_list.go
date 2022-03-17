package paperswithcode_go

import (
	"fmt"
	"github.com/codingpot/paperswithcode-go/v2/models"
	"net/url"
)

// PaperMethodList returns the methods used in the given paper.
func (c *Client) PaperMethodList(paperID string) (models.MethodList, error) {
	pURL := fmt.Sprintf("%s/papers/%s/methods", c.baseURL, url.QueryEscape(paperID))
	return getJson[models.MethodList](c, pURL)
}
