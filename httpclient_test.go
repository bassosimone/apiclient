package apiclient

import (
	"errors"
	"net/http"
)

var ErrMocked = errors.New("mocked error")

type MockableHTTPClient struct {
	Resp *http.Response
	Err  error
}

func (c *MockableHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return c.Resp, c.Err
}

type MockableBodyWithFailure struct{}

func (b *MockableBodyWithFailure) Read(d []byte) (int, error) {
	return 0, ErrMocked
}

func (b *MockableBodyWithFailure) Close() error {
	return nil
}
