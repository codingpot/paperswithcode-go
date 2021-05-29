package paperswithcode_go

import (
	"github.com/codingpot/paperswithcode-go/internal/testutils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var apiToken string

func init() {
	apiToken = testutils.MustExtractAPITokenFromEnv()
}

func TestWithAPIToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("ok"))
		assert.NoError(t, err)
	}))
	defer server.Close()

	c := NewClient(WithAPIToken("MY_TOKEN"))
	emptyRequest, err := http.NewRequest(http.MethodGet, server.URL, nil)

	assert.NoError(t, err)
	_, err = c.HTTPClient.Transport.RoundTrip(emptyRequest)
	assert.NoError(t, err)
	assert.Equal(t, "Token MY_TOKEN", emptyRequest.Header.Get("Authorization"))
}

func TestTransportIsNotProvidedWhenNoAPIIsProvided(t *testing.T) {
	c := NewClient()
	assert.Nil(t, c.HTTPClient.Transport)
}
