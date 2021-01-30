// Package apiclient contains a client for the OONI API.
package apiclient

import "net/http"

// Client is a client for the OONI API.
type Client struct {
	Authorization string
	BaseURL       string
	HTTPClient    *http.Client
	UserAgent     string
}
