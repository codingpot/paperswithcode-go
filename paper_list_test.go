package paperswithcode_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_PaperList(t *testing.T) {
	client := NewClient(WithAPIToken(apiToken))
	_, err := client.PaperList(PaperListParamsDefault())
	assert.NoError(t, err)
}
