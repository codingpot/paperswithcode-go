# paperswithcode-go
client code repository for paperswithcode's official APIs

```go
import "github.com/codingpot/paperswithcode-go"
```

```go
APIToken := "<your-api-token>" // from https://paperswithcode.com/accounts/generate_api_token
c := paperswithcode_go.NewClient(paperswithcode_go.WithAPIToken(token))
```
