package transport

import "net/http"

type transportWithAuthHeader struct {
	r     http.RoundTripper
	token string
}

var _ http.RoundTripper = transportWithAuthHeader{}

// RoundTrip attaches "Authorization: Token <token>" for every request.
func (t transportWithAuthHeader) RoundTrip(request *http.Request) (*http.Response, error) {
	request.Header.Set("Authorization", "Token "+t.token)
	return t.r.RoundTrip(request)
}

// NewTransportWithAuthHeader returns middleware that attaches "Authorization: Token <token>" header to every request.
func NewTransportWithAuthHeader(token string) http.RoundTripper {
	return transportWithAuthHeader{
		r:     http.DefaultTransport,
		token: token,
	}
}
