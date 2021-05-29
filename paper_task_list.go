package paperswithcode_go

import (
	"encoding/json"
	"fmt"
	"github.com/codingpot/paperswithcode-go/models"
	"net/url"
)

func (c *Client) PaperTaskList(paperID string) (*models.TaskList, error) {
	pURL := fmt.Sprintf("%s/papers/%s/tasks/", c.BaseURL, url.QueryEscape(paperID))
	resp, err := c.HTTPClient.Get(pURL)
	if err != nil {
		return nil, err
	}

	var taskList models.TaskList
	err = json.NewDecoder(resp.Body).Decode(&taskList)
	if err != nil {
		return nil, err
	}

	return &taskList, nil
}
