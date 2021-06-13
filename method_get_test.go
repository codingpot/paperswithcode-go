package paperswithcode_go

import (
	"github.com/codingpot/paperswithcode-go/v2/internal/testutils"
	"github.com/codingpot/paperswithcode-go/v2/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_MethodGet(t *testing.T) {
	tests := []struct {
		name     string
		methodID string
		want     *models.Method
		wantErr  bool
	}{
		{
			name:     "With a correct methodID, it returns a method",
			methodID: "multi-head-attention",
			want: &models.Method{
				ID:          "multi-head-attention",
				Name:        "Multi-Head Attention",
				FullName:    "Multi-Head Attention",
				Description: "**Multi-head Attention** is a module for attention mechanisms which runs through an attention mechanism several times in parallel. The independent attention outputs are then concatenated and linearly transformed into the expected dimension. Intuitively, multiple attention heads allows for attending to parts of the sequence differently (e.g. longer-term dependencies versus shorter-term dependencies). \r\n\r\n$$ \\text{MultiHead}\\left(\\textbf{Q}, \\textbf{K}, \\textbf{V}\\right) = \\left[\\text{head}\\_{1},\\dots,\\text{head}\\_{h}\\right]\\textbf{W}_{0}$$\r\n\r\n$$\\text{where} \\text{ head}\\_{i} = \\text{Attention} \\left(\\textbf{Q}\\textbf{W}\\_{i}^{Q}, \\textbf{K}\\textbf{W}\\_{i}^{K}, \\textbf{V}\\textbf{W}\\_{i}^{V} \\right) $$\r\n\r\nAbove $\\textbf{W}$ are all learnable parameter matrices.\r\n\r\nNote that [scaled dot-product attention](https://paperswithcode.com/method/scaled) is most commonly used in this module, although in principle it can be swapped out for other types of attention mechanism.\r\n\r\nSource: [Lilian Weng](https://lilianweng.github.io/lil-log/2018/06/24/attention-attention.html#a-family-of-attention-mechanisms)",
				Paper:       testutils.ToStringPtr("attention-is-all-you-need"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient()
			got, err := c.MethodGet(tt.methodID)
			if tt.wantErr {
				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}