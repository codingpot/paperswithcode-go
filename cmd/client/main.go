package main

import (
	"context"
	"fmt"
	"github.com/codingpot/paperswithcode-go"
	"os"
	"strings"
)

func main() {
	token, ok := os.LookupEnv("PAPERSWITHCODE_API_TOKEN")
	if !ok {
		panic("PAPERSWITHCODE_API_TOKEN environment variable is not found")
	}

	c := paperswithcode_go.NewClient(paperswithcode_go.WithAPIToken(token))
	ctx := context.Background()
	paper := "generative adversarial networks"
	paper = strings.ReplaceAll(paper, " ", "-")

	paperList, _ := c.PaperGet(paper)
	fmt.Println(paperList)
	fmt.Println()
	fmt.Println()
	fmt.Println()

	methodList, _ := c.GetMethodList(ctx, paper)
	fmt.Println(methodList)
	fmt.Println()
	fmt.Println()
	fmt.Println()

	repositoryList, _ := c.GetRepositoryList(ctx, paper)
	fmt.Println(repositoryList)
	fmt.Println()
	fmt.Println()
	fmt.Println()
}
