package paperswithcode_go

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var apiToken string

func init() {
	var ok bool
	apiToken, ok = os.LookupEnv("PAPERSWITHCODE_API_TOKEN")

	if !ok {
		panic("expected PAPERSWITHCODE_API_TOKEN environment variable")
	}
}

func TestClient_PaperList(t *testing.T) {
	client := NewClient(WithAPIToken(apiToken))
	_, err := client.PaperList(PaperListParamsDefault())
	assert.NoError(t, err)
}
