package paperswithcode_go

import (
	"github.com/codingpot/paperswithcode-go/v2/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_PaperMethodList(t *testing.T) {
	c := NewClient()
	paperID := "generative-adversarial-networks"
	got, err := c.PaperMethodList(paperID)
	assert.NoError(t, err)

	expected := &models.MethodList{
		Count:    2,
		Next:     nil,
		Previous: nil,
		Results: []*models.Method{
			{
				ID:          "gan",
				Name:        "GAN",
				FullName:    "Generative Adversarial Network",
				Description: "A **GAN**, or **Generative Adversarial Network**, is a generative model that simultaneously trains\r\ntwo models: a generative model $G$ that captures the data distribution, and a discriminative model $D$ that estimates the\r\nprobability that a sample came from the training data rather than $G$.\r\n\r\nThe training procedure for $G$ is to maximize the probability of $D$ making\r\na mistake. This framework corresponds to a minimax two-player game. In the\r\nspace of arbitrary functions $G$ and $D$, a unique solution exists, with $G$\r\nrecovering the training data distribution and $D$ equal to $\\frac{1}{2}$\r\neverywhere. In the case where $G$ and $D$ are defined by multilayer perceptrons,\r\nthe entire system can be trained with backpropagation. \r\n\r\n(Image Source: [here](http://www.kdnuggets.com/2017/01/generative-adversarial-networks-hot-topic-machine-learning.html))",
				Paper:       &paperID,
			},
			{
				ID:          "convolution",
				Name:        "Convolution",
				FullName:    "Convolution",
				Description: "A **convolution** is a type of matrix operation, consisting of a kernel, a small matrix of weights, that slides over input data performing element-wise multiplication with the part of the input it is on, then summing the results into an output.\r\n\r\nIntuitively, a convolution allows for weight sharing - reducing the number of effective parameters - and image translation (allowing for the same feature to be detected in different parts of the input space).\r\n\r\nImage Source: [https://arxiv.org/pdf/1603.07285.pdf](https://arxiv.org/pdf/1603.07285.pdf)",
				Paper:       nil,
			},
		},
	}
	assert.Equal(t, expected, got)
}
