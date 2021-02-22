// Package apiclient contains the OONI-API Client.
//
// Usage
//
// Create a new Config instance. The zero Config instance represents a
// valid configuration. You may still want to change the defaults.
//
// Call New to obtain a Client instance.
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
	"context"
	"errors"
	"io"
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
	// Do should work like http.Client.Do.
	Do(req *http.Request) (*http.Response, error)
}

// KVStore is a key-value store. The Client uses a KVStore to
// persist on disk authentication information.
type KVStore interface {
	// Get returns the value for the specified key.
	Get(key string) ([]byte, error)

	// Set sets the value for the specified key.
	Set(key string, value []byte) error
}

// JSONCodec is a JSON encoder and decoder.
type JSONCodec interface {
	// Encode encodes v as a serialized JSON byte slice.
	Encode(v interface{}) ([]byte, error)

	// Decode decodes the serialized JSON byte slice into v.
	Decode(b []byte, v interface{}) error
}

// RequestMaker makes an HTTP request.
type RequestMaker interface {
	// NewRequest creates a new HTTP request.
	NewRequest(ctx context.Context, method, URL string, body io.Reader) (*http.Request, error)
}

// GobCodec is a codec for Go's gob format.
type GobCodec interface {
	// Encode encodes v as a serialized gob byte slice.
	Encode(v interface{}) ([]byte, error)

	// Decode decodes the serialized gob byte slice into v.
	Decode(b []byte, v interface{}) error
}

// TemplateExecutor executes a text template.
type TemplateExecutor interface {
	// Execute takes in input a template string and some piece of data. It
	// returns either a string where template parameters have been replaced,
	// on success, or an error, on failure.
	Execute(tmpl string, v interface{}) (string, error)
}

// Config contains configuration for creating a client.
type Config struct {
	// BaseURL is the base URL to use. If not set we will
	// configure the Client to use a default URL.
	BaseURL string

	// GobCodec is the gob codec to use. If not set we will
	// configure Client to use the stdlib's codec.
	GobCodec GobCodec

	// HTTPClient is the HTTP client to use. If not set we will
	// configure the Client to use a default HTTP client.
	HTTPClient HTTPClient

	// JSONCodec is the JSON codec to use. If not set we will
	// use the Go standard library's JSON codec.
	JSONCodec JSONCodec

	// KVStore is the key-value store to use. If not set, we will
	// configure the Client to use a memory based KVStore.
	KVStore KVStore

	// RequestMaker is the HTTP request maker to use. If not set, we will
	// configure the Client to use the standard library's maker.
	RequestMaker RequestMaker

	// TemplateExecutor is the text template executor to use. If not set, we
	// will use the standard library to perform this task.
	TemplateExecutor TemplateExecutor

	// UserAgent is the user agent to use. If not set, we will
	// configure the Client to use a default user agent.
	UserAgent string
}

// Client is a client for the OONI API. The client does not keep
// any on memory state, so it's cheap to create and destroy.
// You must create a new client using the New factory function.
type Client struct {
	baseURL          string
	gobCodec         GobCodec
	httpClient       HTTPClient
	jsonCodec        JSONCodec
	kvStore          KVStore
	requestMaker     RequestMaker
	templateExecutor TemplateExecutor
	userAgent        string
}

const defaultBaseURL = "https://ps1.ooni.io"

// New creates a new instance of Client. You must use this
// function to create a new valid instance.
func New(config *Config) *Client {
	baseURL := config.BaseURL
	if baseURL == "" {
		baseURL = defaultBaseURL
	}
	gobCodec := config.GobCodec
	if gobCodec == nil {
		gobCodec = &stdlibGobCodec{}
	}
	httpClient := config.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	jsonCodec := config.JSONCodec
	if jsonCodec == nil {
		jsonCodec = &stdlibJSONCodec{}
	}
	kvStore := config.KVStore
	if kvStore == nil {
		kvStore = &memkvstore{}
	}
	requestMaker := config.RequestMaker
	if requestMaker == nil {
		requestMaker = &stdlibRequestMaker{}
	}
	templateExecutor := config.TemplateExecutor
	if templateExecutor == nil {
		templateExecutor = &stdlibTemplateExecutor{}
	}
	// TODO(bassosimone): make sure this leads to the empty
	// user agent and otherwise change the docs.
	userAgent := config.UserAgent
	return &Client{
		baseURL:      baseURL,
		gobCodec:     gobCodec,
		httpClient:   httpClient,
		jsonCodec:    jsonCodec,
		kvStore:      kvStore,
		requestMaker: requestMaker,
		userAgent:    userAgent,
	}
}
