// Package apiclient contains clients for the OONI API.
//
// Usage
//
// For each defined API Foobar, there is a structure called FoobarAPI. Instantiate
// a new FoobarAPI structure. In most cases, the zero structure is already valid. In
// some cases, you need to explicitly initialize the Authorizer.
//
// You MAY reuse the same FoobarAPI structure to service multiple requests.
//
// To call the API, create and fill a FoobarRequest structure. Pass this structure along
// with a valid context to FoobarAPI's Call method.
//
// You will get back either an error (and a nil FoobarResponse instance) or a
// valid FoobarResponse instance (and a nil error).
//
// Maintenance
//
// Edit internal/datamodel to change the request and response structures. Edit
// internal/apimodel to change the API specification. Run
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
	"context"
	"errors"
	"io"
	"net/http"
	"text/template"
)

// Errors defined by this package. In addition to these errors, this
// package may of course return any other stdlib specific error.
var (
	ErrHTTPFailure       = errors.New("apiclient: http request failed")
	ErrJSONLiteralNull   = errors.New("apiclient: server returned us a literal null")
	ErrEmptyField        = errors.New("apiclient: empty field")
	ErrMissingAuthorizer = errors.New("apiclient: missing Authorizer")
)

// Swagger returns the API swagger v2.0 as a serialized JSON.
func Swagger() string {
	return swagger
}

// HTTPClient is the interface of a generic HTTP client.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Authorizer authenticates specific client requests.
type Authorizer interface {
	// MaybeRefreshToken refreshes the token for Authorization and returns
	// either such a token, on success, or the error that occurred.
	MaybeRefreshToken(ctx context.Context) (string, error)
}

type staticAuthorizer struct {
	token string
}

func (sa *staticAuthorizer) MaybeRefreshToken(ctx context.Context) (string, error) {
	return sa.token, nil
}

// NewStaticAuthorizer creates a new Authorizer that always
// returns the specified token to the caller.
func NewStaticAuthorizer(token string) Authorizer {
	return &staticAuthorizer{token}
}

// Client is a client for the OONI API.
type Client struct {
	// BaseURL is the base URL for the OONI API.
	BaseURL string

	// HTTPClient is the HTTP client to use. If not set, we will
	// use the http.DefaultClient client.
	HTTPClient HTTPClient

	// UserAgent is the user agent for the OONI API.
	UserAgent string
}

// MaybeRefreshToken implements Authorizer.MaybeRefreshToken.
func (c *Client) MaybeRefreshToken(ctx context.Context) (string, error) {
	return c.maybeLogin(ctx)
}

type textTemplate interface {
	Parse(text string) (textTemplate, error)
	Execute(wr io.Writer, data interface{}) error
}

type stdlibTextTemplate struct {
	*template.Template
}

func (t *stdlibTextTemplate) Parse(text string) (textTemplate, error) {
	out, err := t.Template.Parse(text)
	if err != nil {
		return nil, err
	}
	return &stdlibTextTemplate{out}, nil
}

func newStdlibTextTemplate(name string) textTemplate {
	return &stdlibTextTemplate{template.New(name)}
}
