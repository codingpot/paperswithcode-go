package paperswithcode_go

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_PaperResultList(t *testing.T) {
	paperID := "hierarchical-multi-scale-attention-for"

	c := NewClient()
	got, err := c.PaperResultList(paperID)
	assert.NoError(t, err)
	assert.Greater(t, got.Count, 0)
}
