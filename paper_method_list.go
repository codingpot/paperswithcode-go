package paperswithcode_go

import (
	"encoding/json"
	"fmt"
	"github.com/codingpot/paperswithcode-go/models"
	"net/http"
	"net/url"
)

func (c *Client) PaperMethodList(paperID string) (*models.MethodList, error) {
	pURL := fmt.Sprintf("%s/papers/%s/methods", c.BaseURL, url.QueryEscape(paperID))
	response, err := http.Get(pURL)
	if err != nil {
		return nil, err
	}

	var methodList models.MethodList
	err = json.NewDecoder(response.Body).Decode(&methodList)
	if err != nil {
		return nil, err
	}

	return &methodList, err
}
