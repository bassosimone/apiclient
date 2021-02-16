// Code generated by go generate; DO NOT EDIT.
// 2021-02-16 09:43:24.625213665 +0100 CET m=+0.000329023

package apiclient

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

//go:generate go run ./internal/gencalltest/...

func TestCheckReportIDInvalidURL(t *testing.T) {
	api := &CheckReportIDAPI{
		BaseURL: "\t", // invalid
	}
	ctx := context.Background()
	req := &CheckReportIDRequest{}
	resp, err := api.Call(ctx, req)
	if err == nil || !strings.HasSuffix(err.Error(), "invalid control character in URL") {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestCheckReportIDWithHTTPErr(t *testing.T) {
	clnt := &MockableHTTPClient{Err: ErrMocked}
	api := &CheckReportIDAPI{
		BaseURL:    "https://ps1.ooni.io",
		HTTPClient: clnt,
	}
	ctx := context.Background()
	req := &CheckReportIDRequest{
		ReportID: "antani",
	}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrMocked) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestCheckReportIDWithNewRequestErr(t *testing.T) {
	api := &CheckReportIDAPI{
		BaseURL:    "https://ps1.ooni.io",
		NewRequest: func(ctx context.Context, method, URL string, body io.Reader) (*http.Request, error) {
			return nil, ErrMocked
		},
	}
	ctx := context.Background()
	req := &CheckReportIDRequest{
		ReportID: "antani",
	}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrMocked) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestCheckReportIDWith400(t *testing.T) {
	clnt := &MockableHTTPClient{Resp: &http.Response{StatusCode: 400}}
	api := &CheckReportIDAPI{
		BaseURL:    "https://ps1.ooni.io",
		HTTPClient: clnt,
	}
	ctx := context.Background()
	req := &CheckReportIDRequest{
		ReportID: "antani",
	}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrHTTPFailure) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestCheckInInvalidURL(t *testing.T) {
	api := &CheckInAPI{
		BaseURL: "\t", // invalid
	}
	ctx := context.Background()
	req := &CheckInRequest{}
	resp, err := api.Call(ctx, req)
	if err == nil || !strings.HasSuffix(err.Error(), "invalid control character in URL") {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestCheckInWithHTTPErr(t *testing.T) {
	clnt := &MockableHTTPClient{Err: ErrMocked}
	api := &CheckInAPI{
		BaseURL:    "https://ps1.ooni.io",
		HTTPClient: clnt,
	}
	ctx := context.Background()
	req := &CheckInRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrMocked) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestCheckInMarshalErr(t *testing.T) {
	api := &CheckInAPI{
		BaseURL: "https://ps1.ooni.io",
		marshal: func(v interface{}) ([]byte, error) {
			return nil, ErrMocked
		},
	}
	ctx := context.Background()
	req := &CheckInRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrMocked) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestCheckInWithNewRequestErr(t *testing.T) {
	api := &CheckInAPI{
		BaseURL:    "https://ps1.ooni.io",
		NewRequest: func(ctx context.Context, method, URL string, body io.Reader) (*http.Request, error) {
			return nil, ErrMocked
		},
	}
	ctx := context.Background()
	req := &CheckInRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrMocked) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestCheckInWith400(t *testing.T) {
	clnt := &MockableHTTPClient{Resp: &http.Response{StatusCode: 400}}
	api := &CheckInAPI{
		BaseURL:    "https://ps1.ooni.io",
		HTTPClient: clnt,
	}
	ctx := context.Background()
	req := &CheckInRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrHTTPFailure) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestLoginInvalidURL(t *testing.T) {
	api := &LoginAPI{
		BaseURL: "\t", // invalid
	}
	ctx := context.Background()
	req := &LoginRequest{}
	resp, err := api.Call(ctx, req)
	if err == nil || !strings.HasSuffix(err.Error(), "invalid control character in URL") {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestLoginWithHTTPErr(t *testing.T) {
	clnt := &MockableHTTPClient{Err: ErrMocked}
	api := &LoginAPI{
		BaseURL:    "https://ps1.ooni.io",
		HTTPClient: clnt,
	}
	ctx := context.Background()
	req := &LoginRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrMocked) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestLoginMarshalErr(t *testing.T) {
	api := &LoginAPI{
		BaseURL: "https://ps1.ooni.io",
		marshal: func(v interface{}) ([]byte, error) {
			return nil, ErrMocked
		},
	}
	ctx := context.Background()
	req := &LoginRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrMocked) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestLoginWithNewRequestErr(t *testing.T) {
	api := &LoginAPI{
		BaseURL:    "https://ps1.ooni.io",
		NewRequest: func(ctx context.Context, method, URL string, body io.Reader) (*http.Request, error) {
			return nil, ErrMocked
		},
	}
	ctx := context.Background()
	req := &LoginRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrMocked) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestLoginWith400(t *testing.T) {
	clnt := &MockableHTTPClient{Resp: &http.Response{StatusCode: 400}}
	api := &LoginAPI{
		BaseURL:    "https://ps1.ooni.io",
		HTTPClient: clnt,
	}
	ctx := context.Background()
	req := &LoginRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrHTTPFailure) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestMeasurementMetaInvalidURL(t *testing.T) {
	api := &MeasurementMetaAPI{
		BaseURL: "\t", // invalid
	}
	ctx := context.Background()
	req := &MeasurementMetaRequest{}
	resp, err := api.Call(ctx, req)
	if err == nil || !strings.HasSuffix(err.Error(), "invalid control character in URL") {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestMeasurementMetaWithHTTPErr(t *testing.T) {
	clnt := &MockableHTTPClient{Err: ErrMocked}
	api := &MeasurementMetaAPI{
		BaseURL:    "https://ps1.ooni.io",
		HTTPClient: clnt,
	}
	ctx := context.Background()
	req := &MeasurementMetaRequest{
		ReportID: "antani",
	}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrMocked) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestMeasurementMetaWithNewRequestErr(t *testing.T) {
	api := &MeasurementMetaAPI{
		BaseURL:    "https://ps1.ooni.io",
		NewRequest: func(ctx context.Context, method, URL string, body io.Reader) (*http.Request, error) {
			return nil, ErrMocked
		},
	}
	ctx := context.Background()
	req := &MeasurementMetaRequest{
		ReportID: "antani",
	}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrMocked) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestMeasurementMetaWith400(t *testing.T) {
	clnt := &MockableHTTPClient{Resp: &http.Response{StatusCode: 400}}
	api := &MeasurementMetaAPI{
		BaseURL:    "https://ps1.ooni.io",
		HTTPClient: clnt,
	}
	ctx := context.Background()
	req := &MeasurementMetaRequest{
		ReportID: "antani",
	}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrHTTPFailure) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestRegisterInvalidURL(t *testing.T) {
	api := &RegisterAPI{
		BaseURL: "\t", // invalid
	}
	ctx := context.Background()
	req := &RegisterRequest{}
	resp, err := api.Call(ctx, req)
	if err == nil || !strings.HasSuffix(err.Error(), "invalid control character in URL") {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestRegisterWithHTTPErr(t *testing.T) {
	clnt := &MockableHTTPClient{Err: ErrMocked}
	api := &RegisterAPI{
		BaseURL:    "https://ps1.ooni.io",
		HTTPClient: clnt,
	}
	ctx := context.Background()
	req := &RegisterRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrMocked) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestRegisterMarshalErr(t *testing.T) {
	api := &RegisterAPI{
		BaseURL: "https://ps1.ooni.io",
		marshal: func(v interface{}) ([]byte, error) {
			return nil, ErrMocked
		},
	}
	ctx := context.Background()
	req := &RegisterRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrMocked) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestRegisterWithNewRequestErr(t *testing.T) {
	api := &RegisterAPI{
		BaseURL:    "https://ps1.ooni.io",
		NewRequest: func(ctx context.Context, method, URL string, body io.Reader) (*http.Request, error) {
			return nil, ErrMocked
		},
	}
	ctx := context.Background()
	req := &RegisterRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrMocked) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestRegisterWith400(t *testing.T) {
	clnt := &MockableHTTPClient{Resp: &http.Response{StatusCode: 400}}
	api := &RegisterAPI{
		BaseURL:    "https://ps1.ooni.io",
		HTTPClient: clnt,
	}
	ctx := context.Background()
	req := &RegisterRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrHTTPFailure) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestTestHelpersInvalidURL(t *testing.T) {
	api := &TestHelpersAPI{
		BaseURL: "\t", // invalid
	}
	ctx := context.Background()
	req := &TestHelpersRequest{}
	resp, err := api.Call(ctx, req)
	if err == nil || !strings.HasSuffix(err.Error(), "invalid control character in URL") {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestTestHelpersWithHTTPErr(t *testing.T) {
	clnt := &MockableHTTPClient{Err: ErrMocked}
	api := &TestHelpersAPI{
		BaseURL:    "https://ps1.ooni.io",
		HTTPClient: clnt,
	}
	ctx := context.Background()
	req := &TestHelpersRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrMocked) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestTestHelpersWithNewRequestErr(t *testing.T) {
	api := &TestHelpersAPI{
		BaseURL:    "https://ps1.ooni.io",
		NewRequest: func(ctx context.Context, method, URL string, body io.Reader) (*http.Request, error) {
			return nil, ErrMocked
		},
	}
	ctx := context.Background()
	req := &TestHelpersRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrMocked) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestTestHelpersWith400(t *testing.T) {
	clnt := &MockableHTTPClient{Resp: &http.Response{StatusCode: 400}}
	api := &TestHelpersAPI{
		BaseURL:    "https://ps1.ooni.io",
		HTTPClient: clnt,
	}
	ctx := context.Background()
	req := &TestHelpersRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrHTTPFailure) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestPsiphonConfigInvalidURL(t *testing.T) {
	api := &PsiphonConfigAPI{
		BaseURL: "\t", // invalid
	}
	ctx := context.Background()
	req := &PsiphonConfigRequest{}
	resp, err := api.Call(ctx, req)
	if err == nil || !strings.HasSuffix(err.Error(), "invalid control character in URL") {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestPsiphonConfigWithEmptyToken(t *testing.T) {
	api := &PsiphonConfigAPI{
		BaseURL: "https://ps1.ooni.io",
	}
	ctx := context.Background()
	req := &PsiphonConfigRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrEmptyToken) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestPsiphonConfigWithHTTPErr(t *testing.T) {
	clnt := &MockableHTTPClient{Err: ErrMocked}
	api := &PsiphonConfigAPI{
		BaseURL:    "https://ps1.ooni.io",
		HTTPClient: clnt,
		Token:      "fakeToken",
	}
	ctx := context.Background()
	req := &PsiphonConfigRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrMocked) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestPsiphonConfigWithNewRequestErr(t *testing.T) {
	api := &PsiphonConfigAPI{
		BaseURL:    "https://ps1.ooni.io",
		NewRequest: func(ctx context.Context, method, URL string, body io.Reader) (*http.Request, error) {
			return nil, ErrMocked
		},
		Token:      "fakeToken",
	}
	ctx := context.Background()
	req := &PsiphonConfigRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrMocked) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestPsiphonConfigWith400(t *testing.T) {
	clnt := &MockableHTTPClient{Resp: &http.Response{StatusCode: 400}}
	api := &PsiphonConfigAPI{
		BaseURL:    "https://ps1.ooni.io",
		HTTPClient: clnt,
		Token:      "fakeToken",
	}
	ctx := context.Background()
	req := &PsiphonConfigRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrHTTPFailure) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestTorTargetsInvalidURL(t *testing.T) {
	api := &TorTargetsAPI{
		BaseURL: "\t", // invalid
	}
	ctx := context.Background()
	req := &TorTargetsRequest{}
	resp, err := api.Call(ctx, req)
	if err == nil || !strings.HasSuffix(err.Error(), "invalid control character in URL") {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestTorTargetsWithEmptyToken(t *testing.T) {
	api := &TorTargetsAPI{
		BaseURL: "https://ps1.ooni.io",
	}
	ctx := context.Background()
	req := &TorTargetsRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrEmptyToken) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestTorTargetsWithHTTPErr(t *testing.T) {
	clnt := &MockableHTTPClient{Err: ErrMocked}
	api := &TorTargetsAPI{
		BaseURL:    "https://ps1.ooni.io",
		HTTPClient: clnt,
		Token:      "fakeToken",
	}
	ctx := context.Background()
	req := &TorTargetsRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrMocked) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestTorTargetsWithNewRequestErr(t *testing.T) {
	api := &TorTargetsAPI{
		BaseURL:    "https://ps1.ooni.io",
		NewRequest: func(ctx context.Context, method, URL string, body io.Reader) (*http.Request, error) {
			return nil, ErrMocked
		},
		Token:      "fakeToken",
	}
	ctx := context.Background()
	req := &TorTargetsRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrMocked) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestTorTargetsWith400(t *testing.T) {
	clnt := &MockableHTTPClient{Resp: &http.Response{StatusCode: 400}}
	api := &TorTargetsAPI{
		BaseURL:    "https://ps1.ooni.io",
		HTTPClient: clnt,
		Token:      "fakeToken",
	}
	ctx := context.Background()
	req := &TorTargetsRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrHTTPFailure) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestURLsInvalidURL(t *testing.T) {
	api := &URLsAPI{
		BaseURL: "\t", // invalid
	}
	ctx := context.Background()
	req := &URLsRequest{}
	resp, err := api.Call(ctx, req)
	if err == nil || !strings.HasSuffix(err.Error(), "invalid control character in URL") {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestURLsWithHTTPErr(t *testing.T) {
	clnt := &MockableHTTPClient{Err: ErrMocked}
	api := &URLsAPI{
		BaseURL:    "https://ps1.ooni.io",
		HTTPClient: clnt,
	}
	ctx := context.Background()
	req := &URLsRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrMocked) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestURLsWithNewRequestErr(t *testing.T) {
	api := &URLsAPI{
		BaseURL:    "https://ps1.ooni.io",
		NewRequest: func(ctx context.Context, method, URL string, body io.Reader) (*http.Request, error) {
			return nil, ErrMocked
		},
	}
	ctx := context.Background()
	req := &URLsRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrMocked) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestURLsWith400(t *testing.T) {
	clnt := &MockableHTTPClient{Resp: &http.Response{StatusCode: 400}}
	api := &URLsAPI{
		BaseURL:    "https://ps1.ooni.io",
		HTTPClient: clnt,
	}
	ctx := context.Background()
	req := &URLsRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrHTTPFailure) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestOpenReportInvalidURL(t *testing.T) {
	api := &OpenReportAPI{
		BaseURL: "\t", // invalid
	}
	ctx := context.Background()
	req := &OpenReportRequest{}
	resp, err := api.Call(ctx, req)
	if err == nil || !strings.HasSuffix(err.Error(), "invalid control character in URL") {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestOpenReportWithHTTPErr(t *testing.T) {
	clnt := &MockableHTTPClient{Err: ErrMocked}
	api := &OpenReportAPI{
		BaseURL:    "https://ps1.ooni.io",
		HTTPClient: clnt,
	}
	ctx := context.Background()
	req := &OpenReportRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrMocked) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestOpenReportMarshalErr(t *testing.T) {
	api := &OpenReportAPI{
		BaseURL: "https://ps1.ooni.io",
		marshal: func(v interface{}) ([]byte, error) {
			return nil, ErrMocked
		},
	}
	ctx := context.Background()
	req := &OpenReportRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrMocked) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestOpenReportWithNewRequestErr(t *testing.T) {
	api := &OpenReportAPI{
		BaseURL:    "https://ps1.ooni.io",
		NewRequest: func(ctx context.Context, method, URL string, body io.Reader) (*http.Request, error) {
			return nil, ErrMocked
		},
	}
	ctx := context.Background()
	req := &OpenReportRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrMocked) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestOpenReportWith400(t *testing.T) {
	clnt := &MockableHTTPClient{Resp: &http.Response{StatusCode: 400}}
	api := &OpenReportAPI{
		BaseURL:    "https://ps1.ooni.io",
		HTTPClient: clnt,
	}
	ctx := context.Background()
	req := &OpenReportRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrHTTPFailure) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestSubmitMeasurementInvalidURL(t *testing.T) {
	api := &SubmitMeasurementAPI{
		BaseURL: "\t", // invalid
	}
	ctx := context.Background()
	req := &SubmitMeasurementRequest{}
	resp, err := api.Call(ctx, req)
	if err == nil || !strings.HasSuffix(err.Error(), "invalid control character in URL") {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestSubmitMeasurementWithHTTPErr(t *testing.T) {
	clnt := &MockableHTTPClient{Err: ErrMocked}
	api := &SubmitMeasurementAPI{
		BaseURL:    "https://ps1.ooni.io",
		HTTPClient: clnt,
	}
	ctx := context.Background()
	req := &SubmitMeasurementRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrMocked) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestSubmitMeasurementMarshalErr(t *testing.T) {
	api := &SubmitMeasurementAPI{
		BaseURL: "https://ps1.ooni.io",
		marshal: func(v interface{}) ([]byte, error) {
			return nil, ErrMocked
		},
	}
	ctx := context.Background()
	req := &SubmitMeasurementRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrMocked) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestSubmitMeasurementWithNewRequestErr(t *testing.T) {
	api := &SubmitMeasurementAPI{
		BaseURL:    "https://ps1.ooni.io",
		NewRequest: func(ctx context.Context, method, URL string, body io.Reader) (*http.Request, error) {
			return nil, ErrMocked
		},
	}
	ctx := context.Background()
	req := &SubmitMeasurementRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrMocked) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

func TestSubmitMeasurementWith400(t *testing.T) {
	clnt := &MockableHTTPClient{Resp: &http.Response{StatusCode: 400}}
	api := &SubmitMeasurementAPI{
		BaseURL:    "https://ps1.ooni.io",
		HTTPClient: clnt,
	}
	ctx := context.Background()
	req := &SubmitMeasurementRequest{}
	resp, err := api.Call(ctx, req)
	if !errors.Is(err, ErrHTTPFailure) {
		t.Fatalf("not the error we expected: %+v", err)
	}
	if resp != nil {
		t.Fatal("expected nil resp")
	}
}

