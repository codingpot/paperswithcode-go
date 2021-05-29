package transport

import "net/http"

type transportWithAuthHeader struct {
	r     http.RoundTripper
	token string
}

var _ http.RoundTripper = transportWithAuthHeader{}

func (t transportWithAuthHeader) RoundTrip(request *http.Request) (*http.Response, error) {
	request.Header.Set("Authorization", "Token "+t.token)
	return t.r.RoundTrip(request)
}

func NewTransportWithAuthHeader(token string) http.RoundTripper {
	return transportWithAuthHeader{
		r:     http.DefaultTransport,
		token: token,
	}
}
