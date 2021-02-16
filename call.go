// Code generated by go generate; DO NOT EDIT.
// 2021-02-16 17:30:52.26542541 +0100 CET m=+0.000271203

package apiclient

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

//go:generate go run ./internal/gencall/...

// CheckReportIDAPI is the CheckReportID API. The zero-value structure
// works as intended using suitable default values.
type CheckReportIDAPI struct {
	BaseURL    string
	HTTPClient HTTPClient
	NewRequest func(ctx context.Context, method, URL string, body io.Reader) (*http.Request, error)
	UserAgent  string
	marshal    func(v interface{}) ([]byte, error)
	unmarshal  func(b []byte, v interface{}) error
}

// Call calls GET /api/_/check_report_id. Arguments MUST NOT be nil. The return
// value is either a non-nil error or a non-nil result.
func (api CheckReportIDAPI) Call(ctx context.Context, in *CheckReportIDRequest) (*CheckReportIDResponse, error) {
	req, err := api.newRequest(ctx, api.BaseURL, in)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", api.UserAgent)
	var httpClient HTTPClient = http.DefaultClient
	if api.HTTPClient != nil {
		httpClient = api.HTTPClient
	}
	return api.newResponse(httpClient.Do(req))
}

// CheckInAPI is the CheckIn API. The zero-value structure
// works as intended using suitable default values.
type CheckInAPI struct {
	BaseURL    string
	HTTPClient HTTPClient
	NewRequest func(ctx context.Context, method, URL string, body io.Reader) (*http.Request, error)
	UserAgent  string
	marshal    func(v interface{}) ([]byte, error)
	unmarshal  func(b []byte, v interface{}) error
}

// Call calls POST /api/v1/check-in. Arguments MUST NOT be nil. The return
// value is either a non-nil error or a non-nil result.
func (api CheckInAPI) Call(ctx context.Context, in *CheckInRequest) (*CheckInResponse, error) {
	req, err := api.newRequest(ctx, api.BaseURL, in)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", api.UserAgent)
	var httpClient HTTPClient = http.DefaultClient
	if api.HTTPClient != nil {
		httpClient = api.HTTPClient
	}
	return api.newResponse(httpClient.Do(req))
}

// LoginAPI is the Login API. The zero-value structure
// works as intended using suitable default values.
type LoginAPI struct {
	BaseURL    string
	HTTPClient HTTPClient
	NewRequest func(ctx context.Context, method, URL string, body io.Reader) (*http.Request, error)
	UserAgent  string
	marshal    func(v interface{}) ([]byte, error)
	unmarshal  func(b []byte, v interface{}) error
}

// Call calls POST /api/v1/login. Arguments MUST NOT be nil. The return
// value is either a non-nil error or a non-nil result.
func (api LoginAPI) Call(ctx context.Context, in *LoginRequest) (*LoginResponse, error) {
	req, err := api.newRequest(ctx, api.BaseURL, in)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", api.UserAgent)
	var httpClient HTTPClient = http.DefaultClient
	if api.HTTPClient != nil {
		httpClient = api.HTTPClient
	}
	return api.newResponse(httpClient.Do(req))
}

// MeasurementMetaAPI is the MeasurementMeta API. The zero-value structure
// works as intended using suitable default values.
type MeasurementMetaAPI struct {
	BaseURL    string
	HTTPClient HTTPClient
	NewRequest func(ctx context.Context, method, URL string, body io.Reader) (*http.Request, error)
	UserAgent  string
	marshal    func(v interface{}) ([]byte, error)
	unmarshal  func(b []byte, v interface{}) error
}

// Call calls GET /api/v1/measurement_meta. Arguments MUST NOT be nil. The return
// value is either a non-nil error or a non-nil result.
func (api MeasurementMetaAPI) Call(ctx context.Context, in *MeasurementMetaRequest) (*MeasurementMetaResponse, error) {
	req, err := api.newRequest(ctx, api.BaseURL, in)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", api.UserAgent)
	var httpClient HTTPClient = http.DefaultClient
	if api.HTTPClient != nil {
		httpClient = api.HTTPClient
	}
	return api.newResponse(httpClient.Do(req))
}

// RegisterAPI is the Register API. The zero-value structure
// works as intended using suitable default values.
type RegisterAPI struct {
	BaseURL    string
	HTTPClient HTTPClient
	NewRequest func(ctx context.Context, method, URL string, body io.Reader) (*http.Request, error)
	UserAgent  string
	marshal    func(v interface{}) ([]byte, error)
	unmarshal  func(b []byte, v interface{}) error
}

// Call calls POST /api/v1/register. Arguments MUST NOT be nil. The return
// value is either a non-nil error or a non-nil result.
func (api RegisterAPI) Call(ctx context.Context, in *RegisterRequest) (*RegisterResponse, error) {
	req, err := api.newRequest(ctx, api.BaseURL, in)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", api.UserAgent)
	var httpClient HTTPClient = http.DefaultClient
	if api.HTTPClient != nil {
		httpClient = api.HTTPClient
	}
	return api.newResponse(httpClient.Do(req))
}

// TestHelpersAPI is the TestHelpers API. The zero-value structure
// works as intended using suitable default values.
type TestHelpersAPI struct {
	BaseURL    string
	HTTPClient HTTPClient
	NewRequest func(ctx context.Context, method, URL string, body io.Reader) (*http.Request, error)
	UserAgent  string
	marshal    func(v interface{}) ([]byte, error)
	unmarshal  func(b []byte, v interface{}) error
}

// Call calls GET /api/v1/test-helpers. Arguments MUST NOT be nil. The return
// value is either a non-nil error or a non-nil result.
func (api TestHelpersAPI) Call(ctx context.Context, in *TestHelpersRequest) (TestHelpersResponse, error) {
	req, err := api.newRequest(ctx, api.BaseURL, in)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", api.UserAgent)
	var httpClient HTTPClient = http.DefaultClient
	if api.HTTPClient != nil {
		httpClient = api.HTTPClient
	}
	return api.newResponse(httpClient.Do(req))
}

// PsiphonConfigAPI is the PsiphonConfig API. The zero-value structure
// is not valid because Authorizer is always required. We use
// suitable defaults for any other zero-initialized field.
type PsiphonConfigAPI struct {
	Authorizer Authorizer
	BaseURL    string
	HTTPClient HTTPClient
	NewRequest func(ctx context.Context, method, URL string, body io.Reader) (*http.Request, error)
	UserAgent  string
	marshal    func(v interface{}) ([]byte, error)
	unmarshal  func(b []byte, v interface{}) error
}

// Call calls GET /api/v1/test-list/psiphon-config. Arguments MUST NOT be nil. The return
// value is either a non-nil error or a non-nil result.
func (api PsiphonConfigAPI) Call(ctx context.Context, in *PsiphonConfigRequest) (PsiphonConfigResponse, error) {
	req, err := api.newRequest(ctx, api.BaseURL, in)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	if api.Authorizer == nil {
		return nil, ErrMissingAuthorizer
	}
	token, err := api.Authorizer.MaybeRefreshToken(ctx)
	if err != nil {
		return nil, err
	}
	authorization := fmt.Sprintf("Bearer %s", token)
	req.Header.Add("Authorization", authorization)
	req.Header.Add("User-Agent", api.UserAgent)
	var httpClient HTTPClient = http.DefaultClient
	if api.HTTPClient != nil {
		httpClient = api.HTTPClient
	}
	return api.newResponse(httpClient.Do(req))
}

// TorTargetsAPI is the TorTargets API. The zero-value structure
// is not valid because Authorizer is always required. We use
// suitable defaults for any other zero-initialized field.
type TorTargetsAPI struct {
	Authorizer Authorizer
	BaseURL    string
	HTTPClient HTTPClient
	NewRequest func(ctx context.Context, method, URL string, body io.Reader) (*http.Request, error)
	UserAgent  string
	marshal    func(v interface{}) ([]byte, error)
	unmarshal  func(b []byte, v interface{}) error
}

// Call calls GET /api/v1/test-list/tor-targets. Arguments MUST NOT be nil. The return
// value is either a non-nil error or a non-nil result.
func (api TorTargetsAPI) Call(ctx context.Context, in *TorTargetsRequest) (TorTargetsResponse, error) {
	req, err := api.newRequest(ctx, api.BaseURL, in)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	if api.Authorizer == nil {
		return nil, ErrMissingAuthorizer
	}
	token, err := api.Authorizer.MaybeRefreshToken(ctx)
	if err != nil {
		return nil, err
	}
	authorization := fmt.Sprintf("Bearer %s", token)
	req.Header.Add("Authorization", authorization)
	req.Header.Add("User-Agent", api.UserAgent)
	var httpClient HTTPClient = http.DefaultClient
	if api.HTTPClient != nil {
		httpClient = api.HTTPClient
	}
	return api.newResponse(httpClient.Do(req))
}

// URLsAPI is the URLs API. The zero-value structure
// works as intended using suitable default values.
type URLsAPI struct {
	BaseURL    string
	HTTPClient HTTPClient
	NewRequest func(ctx context.Context, method, URL string, body io.Reader) (*http.Request, error)
	UserAgent  string
	marshal    func(v interface{}) ([]byte, error)
	unmarshal  func(b []byte, v interface{}) error
}

// Call calls GET /api/v1/test-list/urls. Arguments MUST NOT be nil. The return
// value is either a non-nil error or a non-nil result.
func (api URLsAPI) Call(ctx context.Context, in *URLsRequest) (*URLsResponse, error) {
	req, err := api.newRequest(ctx, api.BaseURL, in)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", api.UserAgent)
	var httpClient HTTPClient = http.DefaultClient
	if api.HTTPClient != nil {
		httpClient = api.HTTPClient
	}
	return api.newResponse(httpClient.Do(req))
}

// OpenReportAPI is the OpenReport API. The zero-value structure
// works as intended using suitable default values.
type OpenReportAPI struct {
	BaseURL    string
	HTTPClient HTTPClient
	NewRequest func(ctx context.Context, method, URL string, body io.Reader) (*http.Request, error)
	UserAgent  string
	marshal    func(v interface{}) ([]byte, error)
	unmarshal  func(b []byte, v interface{}) error
}

// Call calls POST /report. Arguments MUST NOT be nil. The return
// value is either a non-nil error or a non-nil result.
func (api OpenReportAPI) Call(ctx context.Context, in *OpenReportRequest) (*OpenReportResponse, error) {
	req, err := api.newRequest(ctx, api.BaseURL, in)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", api.UserAgent)
	var httpClient HTTPClient = http.DefaultClient
	if api.HTTPClient != nil {
		httpClient = api.HTTPClient
	}
	return api.newResponse(httpClient.Do(req))
}

// SubmitMeasurementAPI is the SubmitMeasurement API. The zero-value structure
// works as intended using suitable default values.
type SubmitMeasurementAPI struct {
	BaseURL     string
	HTTPClient  HTTPClient
	NewRequest  func(ctx context.Context, method, URL string, body io.Reader) (*http.Request, error)
	UserAgent   string
	marshal     func(v interface{}) ([]byte, error)
	newTemplate func(s string) textTemplate
	unmarshal   func(b []byte, v interface{}) error
}

// Call calls POST /report/{{ .ReportID }}. Arguments MUST NOT be nil. The return
// value is either a non-nil error or a non-nil result.
func (api SubmitMeasurementAPI) Call(ctx context.Context, in *SubmitMeasurementRequest) (*SubmitMeasurementResponse, error) {
	req, err := api.newRequest(ctx, api.BaseURL, in)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", api.UserAgent)
	var httpClient HTTPClient = http.DefaultClient
	if api.HTTPClient != nil {
		httpClient = api.HTTPClient
	}
	return api.newResponse(httpClient.Do(req))
}
