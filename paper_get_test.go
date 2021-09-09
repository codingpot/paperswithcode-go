package paperswithcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_PaperGet(t *testing.T) {
	c := NewClient()
	got, err := c.PaperGet("generative-adversarial-networks")
	assert.NoError(t, err)
	assert.Equal(t, "generative-adversarial-networks", got.ID)
}
