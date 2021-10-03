package paperswithcode_go

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	_, err = c.httpClient.Transport.RoundTrip(emptyRequest)
	assert.NoError(t, err)
	assert.Equal(t, "Token MY_TOKEN", emptyRequest.Header.Get("Authorization"))
}

func TestTransportIsNotProvidedWhenNoAPIIsProvided(t *testing.T) {
	c := NewClient()
	assert.Nil(t, c.httpClient.Transport)
}

func ExampleGetPaperIDFromPaperTitle() {
	paperTitle := "Generative Adversarial Networks"
	fmt.Println(GetPaperIDFromPaperTitle(paperTitle))
	// Output: generative-adversarial-networks
}

func TestGetPaperIDFromPaperTitle(t *testing.T) {
	tests := []struct {
		paperTitle  string
		wantPaperID string
	}{
		{
			paperTitle:  "This Is Paper",
			wantPaperID: "this-is-paper",
		},
		{
			paperTitle:  "This        Is Paper",
			wantPaperID: "this-is-paper",
		},
		{
			paperTitle:  "This-Is-Paper",
			wantPaperID: "this-is-paper",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("The ID of %s should be %s", tt.paperTitle, tt.wantPaperID),
			func(t *testing.T) {
				assert.Equal(t, tt.wantPaperID, GetPaperIDFromPaperTitle(tt.paperTitle))
			})
	}
}
