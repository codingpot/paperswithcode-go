package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	client "paperswithcode/internal"
)

const (
	BASE_URL = "https://paperswithcode.com/api/v1"
)

func NewClient() *client.Client {
	return &client.Client{
		BaseURL: BASE_URL,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func main() {
	c := NewClient()
	ctx := context.Background()
	paper := "generative adversarial networks"
	paper = strings.ReplaceAll(paper, " ", "-")

	paperList, _ := c.GetPaper(ctx, paper)
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
