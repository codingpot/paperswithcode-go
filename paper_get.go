package paperswithcode_go

import (
	"fmt"
	"github.com/codingpot/paperswithcode-go/v2/models"
)

// PaperGet returns a single paper. Note that paperID is hyphen cased (e.g., generative-adversarial-networks).
func (c *Client) PaperGet(paperID string) (models.Paper, error) {
	paperGetURL := fmt.Sprintf("%s/papers/%s/", c.baseURL, paperID)
	return getJson[models.Paper](c, paperGetURL)
}
