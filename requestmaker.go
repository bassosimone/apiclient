package apiclient

import (
	"context"
	"io"
	"net/http"
)

type stdlibRequestMaker struct{}

func (*stdlibRequestMaker) NewRequest(ctx context.Context, method, URL string, body io.Reader) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, method, URL, body)
}
