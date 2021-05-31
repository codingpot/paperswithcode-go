package main

import (
	"fmt"
	"github.com/codingpot/paperswithcode-go/v2"
	"os"
	"strings"
)

func main() {
	token, ok := os.LookupEnv("PAPERSWITHCODE_API_TOKEN")
	if !ok {
		panic("PAPERSWITHCODE_API_TOKEN environment variable is not found")
	}

	c := paperswithcode_go.NewClient(paperswithcode_go.WithAPIToken(token))
	paper := "generative adversarial networks"
	paper = strings.ReplaceAll(paper, " ", "-")

	paperList, _ := c.PaperGet(paper)
	fmt.Println(paperList)
	fmt.Println()
	fmt.Println()
	fmt.Println()

	methodList, _ := c.PaperMethodList(paper)
	fmt.Println(methodList)
	fmt.Println()
	fmt.Println()
	fmt.Println()

	repositoryList, _ := c.PaperRepositoryList(paper)
	fmt.Println(repositoryList)
	fmt.Println()
	fmt.Println()
	fmt.Println()
}
