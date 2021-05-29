# paperswithcode-go
[![Go](https://github.com/codingpot/paperswithcode-go/actions/workflows/go.yaml/badge.svg)](https://github.com/codingpot/paperswithcode-go/actions/workflows/go.yaml)
[![Go Reference](https://pkg.go.dev/badge/github.com/codingpot/paperswithcode-go.svg)](https://pkg.go.dev/github.com/codingpot/paperswithcode-go)
[![codecov](https://codecov.io/gh/codingpot/paperswithcode-go/branch/main/graph/badge.svg?token=MhzDKZOtWK)](https://codecov.io/gh/codingpot/paperswithcode-go)

client code repository for paperswithcode's official APIs

```go
import "github.com/codingpot/paperswithcode-go"
```

```go
APIToken := "<your-api-token>" // from https://paperswithcode.com/accounts/generate_api_token
c := paperswithcode_go.NewClient(paperswithcode_go.WithAPIToken(token))
```
