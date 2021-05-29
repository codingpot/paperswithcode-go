package testutils

import "os"

func MustExtractAPITokenFromEnv() string {
	apiToken, ok := os.LookupEnv("PAPERSWITHCODE_API_TOKEN")

	if !ok {
		panic("expected PAPERSWITHCODE_API_TOKEN environment variable")
	}
	return apiToken
}
