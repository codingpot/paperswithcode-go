package paperswithcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_PaperRepositoryList(t *testing.T) {
	c := NewClient()
	list, err := c.PaperRepositoryList("generative-adversarial-networks")
	assert.NoError(t, err)
	assert.NotEmpty(t, list.Results[0].URL)
}
