package apiclient

import (
	"errors"
	"io"
	"net/http"
	"sync"
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

type MockableEmptyBody struct {
	done bool
	mu   sync.Mutex
}

func (b *MockableEmptyBody) Read(d []byte) (int, error) {
	defer b.mu.Unlock()
	b.mu.Lock()
	if b.done == false {
		b.done = true
		var out = []byte("{}")
		if len(d) < len(out) {
			panic("unexpected very small slice")
		}
		copy(d, out)
		return len(out), nil
	}
	return 0, io.EOF
}

func (b *MockableEmptyBody) Close() error {
	return nil
}

type MockableLiteralNull struct {
	done bool
	mu   sync.Mutex
}

func (b *MockableLiteralNull) Read(d []byte) (int, error) {
	defer b.mu.Unlock()
	b.mu.Lock()
	if b.done == false {
		b.done = true
		var out = []byte("null")
		if len(d) < len(out) {
			panic("unexpected very small slice")
		}
		copy(d, out)
		return len(out), nil
	}
	return 0, io.EOF
}

func (b *MockableLiteralNull) Close() error {
	return nil
}
