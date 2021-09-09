package paperswithcode_go

import (
	"fmt"
	"github.com/codingpot/paperswithcode-go/v2/models"
	"net/url"
)

// PaperTaskList returns tasks (an area of research) for the given paper.
func (c *Client) PaperTaskList(paperID string) (*models.TaskList, error) {
	pURL := fmt.Sprintf("%s/papers/%s/tasks/", c.baseURL, url.QueryEscape(paperID))
	var taskList models.TaskList
	err := c.sendGetRequest(pURL, &taskList)
	return &taskList, err
}
