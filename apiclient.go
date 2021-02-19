// Package apiclient contains the OONI-API Client.
//
// Usage
//
// Create a Client instance. Even though the zero Client works, you most
// likely want to fill in all its fields to customize its behavior.
//
// To call the API, create and fill a request structure. Pass this structure
// along with a valid context to the proper Client method.
//
// You will get back either an error (and a nil response instance) or a
// valid response instance (and a nil error).
//
// Maintenance
//
// Edit ./model to change the request and response structures. Edit
// internal/cmd/generator/internal/spec.go to change the API specification. Run
//
//     go generate ./...
//
// To regenerate all the definitions exported by this package.
//
// Continuous integration
//
// Running tests
//
//     go test ./...
//
// includes a check that verifies that we and the server are using the
// same definition for the exchanged data structures.
//
// API documentation
//
// Please, refer to https://api.ooni.io/apidocs/ for more info.
package apiclient

import (
	"errors"
	"net/http"
)

// Errors defined by this package. In addition to these errors, this
// package may of course return any other stdlib specific error.
var (
	ErrHTTPFailure     = errors.New("apiclient: http request failed")
	ErrJSONLiteralNull = errors.New("apiclient: server returned us a literal null")
	ErrEmptyField      = errors.New("apiclient: empty field")
	ErrUnauthorized    = errors.New("apiclient: not authorized")
	errMissingToken    = errors.New("apiclient: missing authorization token")
)

// Swagger returns the API swagger v2.0 as a serialized JSON.
func Swagger() string {
	return swagger
}

// HTTPClient is the interface of a generic HTTP client. We use this
// interface to abstract the HTTP client on which Client depends.
type HTTPClient interface {
	// Do should work exactly like http.DefaultClient.Do.
	Do(req *http.Request) (*http.Response, error)
}

// Client is a client for the OONI API. The client does not keep
// any on memory state, so it's cheap to create and destroy.
type Client struct {
	// BaseURL is the base URL for the OONI API. If not set, we will
	// use the default API-base-URL.
	BaseURL string

	// HTTPClient is the HTTP client to use. If not set, we will
	// use the http.DefaultClient client.
	HTTPClient HTTPClient

	// KVStore is the key-value store to use. If not set, we will
	// configure a memory-based, ephemeral key-value store.
	KVStore KVStore

	// UserAgent is the user agent for the OONI API. If not set, we
	// will send no User Agent to the server.
	UserAgent string
}

// httpClient returns the configured client or the default.
func (c *Client) httpClient() HTTPClient {
	if c.HTTPClient != nil {
		return c.HTTPClient
	}
	return http.DefaultClient
}
