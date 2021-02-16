// Code generated by go generate; DO NOT EDIT.
// 2021-02-16 09:09:21.553316863 +0100 CET m=+0.000284599

package apiclient

import (
	"context"
	"errors"
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

