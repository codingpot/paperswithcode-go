package paperswithcode_go

import (
	"github.com/codingpot/paperswithcode-go/v2/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_MethodList(t *testing.T) {
	tests := []struct {
		name    string
		params  MethodListParams
		want    models.MethodList
		wantErr bool
	}{
		{
			name: "MethodList returns methods",
			params: MethodListParams{
				Page:         1,
				ItemsPerPage: 2,
			},
			want: models.MethodList{
				Next:     toPtr("https://paperswithcode.com/api/v1/methods/?items_per_page=2&page=2"),
				Previous: nil,
				Results: []models.Method{
					{
						ID:          "1cycle",
						Name:        "1cycle",
						FullName:    "1cycle learning rate scheduling policy",
						Description: "",
						Paper:       toPtr("a-disciplined-approach-to-neural-network"),
					},
					{
						ID:          "1d-cnn",
						Name:        "1D CNN",
						FullName:    "1-Dimensional Convolutional Neural Networks",
						Description: "1D Convolutional Neural Networks are similar to well known and more established 2D Convolutional Neural Networks. 1D Convolutional Neural Networks are used mainly used on text and 1D signals.",
						Paper:       toPtr("convolutional-neural-network-and-rule-based"),
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient()
			got, err := c.MethodList(tt.params)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// we don't want to test count values as it's changing.
			tt.want.Count = got.Count
			// these results are dynamic. Hence, ignore values.
			tt.want.Results = got.Results

			assert.Equal(t, tt.want, got)
		})
	}
}

func toPtr(s string) *string {
	return &s
}

func TestMethodListParams_Build(t *testing.T) {
	type fields struct {
		Page         int
		ItemsPerPage int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "default values are parsed as URL params",
			fields: fields{
				Page:         0,
				ItemsPerPage: 0,
			},
			want: "items_per_page=0&page=0",
		},
		{
			name: "some values are parsed as URL params",
			fields: fields{
				Page:         1,
				ItemsPerPage: 1234,
			},
			want: "items_per_page=1234&page=1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MethodListParams{
				Page:         tt.fields.Page,
				ItemsPerPage: tt.fields.ItemsPerPage,
			}
			assert.Equal(t, tt.want, m.Build())
		})
	}
}
