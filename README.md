# paperswithcode-go

[![Go](https://github.com/codingpot/paperswithcode-go/v2/actions/workflows/go.yaml/badge.svg)](https://github.com/codingpot/paperswithcode-go/v2/actions/workflows/go.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/codingpot/paperswithcode-go)](https://goreportcard.com/report/github.com/codingpot/paperswithcode-go)
[![Go Reference](https://pkg.go.dev/badge/github.com/codingpot/paperswithcode-go/v2.svg)](https://pkg.go.dev/github.com/codingpot/paperswithcode-go/v2)
[![codecov](https://codecov.io/gh/codingpot/paperswithcode-go/branch/main/graph/badge.svg?token=MhzDKZOtWK)](https://codecov.io/gh/codingpot/paperswithcode-go)

This is a client for PapersWithCode read/write API.

For Python version, see https://github.com/paperswithcode/paperswithcode-client

## Quick usage example

```go
import "github.com/codingpot/paperswithcode-go"
```

```go
c := paperswithcode_go.NewClient()
papers, _ := c.PaperList(paperswithcode_go.PaperListParamsDefault())
gan, _ := c.PaperGet(paperswithcode_go.GetPaperIDFromPaperTitle("Generative Adversarial Networks"))
```
