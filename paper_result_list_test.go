package paperswithcode_go

import (
	"github.com/codingpot/paperswithcode-go/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

var one int = 1

func TestClient_PaperResultList(t *testing.T) {
	paperID := "hierarchical-multi-scale-attention-for"
	bestMetric := "Mean IoU (class)"

	c := NewClient()
	got, err := c.PaperResultList(paperID)
	assert.NoError(t, err)

	expected := &models.ResultList{
		Count:    2,
		Next:     nil,
		Previous: nil,
		Results: []*models.Result{
			{
				ID:       "657bb1c5-de55-4afc-ba54-0af74eaab5e2",
				BestRank: &one,
				Metrics: map[string]string{
					bestMetric: "85.1%",
				},
				Methodology:        "HRNet-OCR (Hierarchical Multi-Scale Attention)",
				UsesAdditionalData: false,
				Paper:              &paperID,
				BestMetric:         &bestMetric,
				EvaluatedOn:        models.NewYyyyMmDdDashed(2020, 5, 21),
			},
			{
				ID:       "a0db982b-68a8-4685-921d-a8a557304c85",
				BestRank: nil,
				Metrics: map[string]string{
					"PQ": "17.6",
				},
				Methodology:        "HRNet-OCR (Hierarchical Multi-Scale Attention)",
				UsesAdditionalData: false,
				Paper:              &paperID,
				BestMetric:         nil,
				EvaluatedOn:        models.NewYyyyMmDdDashed(2020, 5, 21),
			},
		},
	}

	assert.Equal(t, expected, got)
}
