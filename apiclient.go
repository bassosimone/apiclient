// Package apiclient contains clients for the OONI API.
//
// Usage
//
// For each defined API Foobar, there is a structure called FoobarAPI. Instantiate
// a new FoobarAPI structure. In most cases, the zero structure is already valid. In
// some cases, you need to explicitly initialize the Auth token.
//
// You MAY reuse the same FoobarAPI structure to service multiple requests.
//
// Then, create and fill a FoobarRequest structure. Pass this structure along with
// a valid context to FoobarAPI's Call method.
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
// API documentation
//
// Please, refer to https://api.ooni.io/apidocs/ for more info.
package apiclient

import "errors"

// This package defines the following errors. In addition to these errors, this
// package may of course return any other stdlib specific error.
var (
	ErrHTTPFailure     = errors.New("apiclient: http request failed")
	ErrJSONLiteralNull = errors.New("apiclient: server returned us a literal null")
	ErrEmptyField      = errors.New("apiclient: empty field")
	ErrEmptyToken      = errors.New("apiclient: empty auth token")
)

// Swagger returns the API swagger v2.0 as a serialized JSON.
func Swagger() string {
	return swagger
}
