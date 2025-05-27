package repository

import (
	"io"
	"net/http"

	"github.com/wiremock/go-wiremock"
)

type WiremockClient interface {
	StubFor(stub *wiremock.StubRule) error
	Reset() error
}

type HTTPClient interface {
	Get(url string) (*http.Response, error)
	Post(url, contentType string, body io.Reader) (*http.Response, error)
}
