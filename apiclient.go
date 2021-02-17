// Package apiclient contains clients for the OONI API.
//
// Usage
//
// Create a Client instance. Even though the zero Client works, you most
// likely want to fill in all its fields to customize its behavior.
//
// For each defined API Foobar, there is a structure called FoobarAPI. Instantiate
// a new FoobarAPI structure using NewFoobarAPI factor, which takes in input
// a Client instance. This will register the Client as the Authorizer for the
// given API. (As low-level alternative, instantiate the API directly.)
//
// You MAY reuse the same FoobarAPI structure to service multiple requests. But
// in general we expect you to create a new structure whenever you need one.
//
// To call the API, create and fill a FoobarRequest structure. Pass this structure
// along with a valid context to FoobarAPI's Call method.
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
	"fmt"
	"net/http"
)

// Errors defined by this package. In addition to these errors, this
// package may of course return any other stdlib specific error.
var (
	ErrHTTPFailure       = errors.New("apiclient: http request failed")
	ErrJSONLiteralNull   = errors.New("apiclient: server returned us a literal null")
	ErrEmptyField        = errors.New("apiclient: empty field")
	errMissingAuthorizer = errors.New("apiclient: missing Authorizer")
)

func newErrEmptyField(field string) error {
	return fmt.Errorf("%w: %s", ErrEmptyField, field)
}

func newHTTPFailure(status int) error {
	return fmt.Errorf("%w: %d", ErrHTTPFailure, status)
}

func newQueryFieldInt64(v int64) string {
	return fmt.Sprintf("%d", v)
}

func newAuthorizationHeader(token string) string {
	return fmt.Sprintf("Bearer %s", token)
}

// Swagger returns the API swagger v2.0 as a serialized JSON.
func Swagger() string {
	return swagger
}

// HTTPClient is the interface of a generic HTTP client.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client is a client for the OONI API.
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

// maybeRefreshToken implements authorizer.maybeRefreshToken.
//
// You typically do not call this method directly. Rather, you create
// an API using NewFoobarAPI(c). This will register the client as
// the Authorizer for the specified API.
//
// When invoked, this method will roughly do the following:
//
// 1. if we already have a valid token, just return it;
//
// 2. if we already have valid orchestra credentials, then
// login in again so to refresh the token, then return the token;
//
// 3. otherwise, create a new account, and then login with
// such an account, so we have a token to return.
//
// This implementation should be robust to a change in
// the backend database where all logins are lost.
func (c *Client) maybeRefreshToken(ctx context.Context) (string, error) {
	return c.maybeLogin(ctx)
}
