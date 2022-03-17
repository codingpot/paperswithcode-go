package paperswithcode_go

import (
	"github.com/codingpot/paperswithcode-go/v2/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_PaperTaskList(t *testing.T) {
	c := NewClient()
	got, err := c.PaperTaskList("generative-adversarial-networks")
	assert.NoError(t, err)

	expected := models.TaskList{
		Count:    0,
		Next:     nil,
		Previous: nil,
		Results:  []models.Task{},
	}
	assert.Equal(t, expected, got)
}
