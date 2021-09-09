package testutils

import "os"

// MustExtractAPITokenFromEnv is a helper function that looks up PAPERSWITHCODE_API_TOKEN.
// It panics if not found.
func MustExtractAPITokenFromEnv() string {
	apiToken, ok := os.LookupEnv("PAPERSWITHCODE_API_TOKEN")

	if !ok {
		panic("expected PAPERSWITHCODE_API_TOKEN environment variable")
	}
	return apiToken
}

// ToStringPtr returns a pointer to the given string.
func ToStringPtr(s string) *string {
	return &s
}
