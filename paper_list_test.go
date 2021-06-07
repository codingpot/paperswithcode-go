package paperswithcode_go

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_PaperList(t *testing.T) {
	client := NewClient(WithAPIToken(apiToken))
	_, err := client.PaperList(PaperListParamsDefault())
	assert.NoError(t, err)
}

func TestPaperListParams_Build(t *testing.T) {
	type fields struct {
		Q            string
		ArxivID      string
		Title        string
		Abstract     string
		Page         int
		ItemsPerPage int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Q is given, it passes Q",
			fields: fields{
				Q:            "wow",
				Page:         1,
				ItemsPerPage: 50,
			},
			want: "page=1&items_per_page=50&q=wow",
		},
		{
			name: "Q is not given, it shouldn't add Q param",
			fields: fields{
				Page:         1,
				ItemsPerPage: 50,
			},
			want: "page=1&items_per_page=50",
		},
		{
			name:   "Default Param is valid",
			fields: fields(PaperListParamsDefault()),
			want:   "page=1&items_per_page=50",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PaperListParams{
				Q:            tt.fields.Q,
				ArxivID:      tt.fields.ArxivID,
				Title:        tt.fields.Title,
				Abstract:     tt.fields.Abstract,
				Page:         tt.fields.Page,
				ItemsPerPage: tt.fields.ItemsPerPage,
			}
			assert.Equal(t, tt.want, p.Build())
		})
	}
}
