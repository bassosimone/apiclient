// Package apiclient contains a client for the OONI API.
package apiclient

import (
	"errors"
	"net/http"
)

// The following errors may be returned by this implementation in
// addition to the errors returned by APIs we call.
var (
	ErrHTTPFailure     = errors.New("apiclient: http request failed")
	ErrJSONLiteralNull = errors.New("apiclient: server returned us a literal null")
	ErrEmptyField      = errors.New("apiclient: empty field")
)

// Client is a client for the OONI API.
type Client struct {
	Accept        string
	Authorization string
	BaseURL       string
	HTTPClient    *http.Client
	UserAgent     string
}
