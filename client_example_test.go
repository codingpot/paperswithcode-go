package paperswithcode_go_test

import (
	"fmt"

	"github.com/codingpot/paperswithcode-go/v2"
)

func Example() {
	c := paperswithcode_go.NewClient()
	papers, _ := c.PaperList(paperswithcode_go.PaperListParamsDefault())
	fmt.Println(len(papers.Results))

	gan, _ := c.PaperGet(paperswithcode_go.GetPaperIDFromPaperTitle("Generative Adversarial Networks"))
	fmt.Println(gan.Title)
	fmt.Println(gan.Authors)

	// Output:
	// 50
	// Generative Adversarial Networks
	// [Ian J. Goodfellow Jean Pouget-Abadie Mehdi Mirza Bing Xu David Warde-Farley Sherjil Ozair Aaron Courville Yoshua Bengio]
}
