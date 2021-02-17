package apiclient

import (
	"errors"
	"io"
	"net/http"
	"sync"
)

var errMocked = errors.New("mocked error")

type mockableHTTPClient struct {
	Resp *http.Response
	Err  error
}

func (c *mockableHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return c.Resp, c.Err
}

type mockableBodyWithFailure struct{}

func (b *mockableBodyWithFailure) Read(d []byte) (int, error) {
	return 0, errMocked
}

func (b *mockableBodyWithFailure) Close() error {
	return nil
}

type mockableEmptyBody struct {
	done bool
	mu   sync.Mutex
}

func (b *mockableEmptyBody) Read(d []byte) (int, error) {
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

func (b *mockableEmptyBody) Close() error {
	return nil
}

type mockableLiteralNull struct {
	done bool
	mu   sync.Mutex
}

func (b *mockableLiteralNull) Read(d []byte) (int, error) {
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

func (b *mockableLiteralNull) Close() error {
	return nil
}
