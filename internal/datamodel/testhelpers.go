package datamodel

// TestHelpersRequest is the TestHelpers request.
type TestHelpersRequest struct{}

// TestHelpersResponse is the TestHelpers response.
type TestHelpersResponse struct {
	HTTPReturnJSONHeaders []TestHelpersHelperInfo `json:"http-return-json-headers"`
	TCPEcho               []TestHelpersHelperInfo `json:"tcp-echo"`
	WebConnectivity       []TestHelpersHelperInfo `json:"web-connectivity"`
}

// TestHelpersHelperInfo is a single helper within the
// response returned by the TestHelpers API.
type TestHelpersHelperInfo struct {
	Address string `json:"address"`
	Type    string `json:"type"`
	Front   string `json:"front,omitempty"`
}
