// Code generated by go generate; DO NOT EDIT.
// 2021-01-31 01:54:50.014865434 +0100 CET m=+0.000135612

package apiclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"text/template"
	"net/http"
	"net/url"
	"strings"
)

//go:generate go run ./internal/gennewrequest/...

func newCheckReportIDRequest(ctx context.Context, baseURL string, req *CheckReportIDRequest) (*http.Request, error) {
	URL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	URL.Path = "/api/_/check_report_id"
	query := url.Values{}
	if req.ReportID == "" {
		return nil, fmt.Errorf("%w: ReportID", ErrEmptyField)
	}
	query.Add("report_id", req.ReportID)
	URL.RawQuery = query.Encode()
	return http.NewRequestWithContext(ctx, "GET", URL.String(), nil)
}

func newCheckInRequest(ctx context.Context, baseURL string, req *CheckInRequest) (*http.Request, error) {
	URL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	URL.Path = "/api/v1/check-in"
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	out, err := http.NewRequestWithContext(ctx, "POST", URL.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	out.Header.Set("Content-Type", "application/json")
	return out, nil
}

func newLoginRequest(ctx context.Context, baseURL string, req *LoginRequest) (*http.Request, error) {
	URL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	URL.Path = "/api/v1/login"
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	out, err := http.NewRequestWithContext(ctx, "POST", URL.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	out.Header.Set("Content-Type", "application/json")
	return out, nil
}

func newMeasurementMetaRequest(ctx context.Context, baseURL string, req *MeasurementMetaRequest) (*http.Request, error) {
	URL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	URL.Path = "/api/v1/measurement_meta"
	query := url.Values{}
	if req.ReportID == "" {
		return nil, fmt.Errorf("%w: ReportID", ErrEmptyField)
	}
	query.Add("report_id", req.ReportID)
	if req.Full {
		query.Add("full", "true")
	}
	if req.Input != "" {
		query.Add("input", req.Input)
	}
	URL.RawQuery = query.Encode()
	return http.NewRequestWithContext(ctx, "GET", URL.String(), nil)
}

func newRegisterRequest(ctx context.Context, baseURL string, req *RegisterRequest) (*http.Request, error) {
	URL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	URL.Path = "/api/v1/register"
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	out, err := http.NewRequestWithContext(ctx, "POST", URL.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	out.Header.Set("Content-Type", "application/json")
	return out, nil
}

func newTestHelpersRequest(ctx context.Context, baseURL string, req *TestHelpersRequest) (*http.Request, error) {
	URL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	URL.Path = "/api/v1/test-helpers"
	return http.NewRequestWithContext(ctx, "GET", URL.String(), nil)
}

func newPsiphonConfigRequest(ctx context.Context, baseURL string, req *PsiphonConfigRequest) (*http.Request, error) {
	URL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	URL.Path = "/api/v1/test-list/psiphon-config"
	return http.NewRequestWithContext(ctx, "GET", URL.String(), nil)
}

func newTorTargetsRequest(ctx context.Context, baseURL string, req *TorTargetsRequest) (*http.Request, error) {
	URL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	URL.Path = "/api/v1/test-list/tor-targets"
	return http.NewRequestWithContext(ctx, "GET", URL.String(), nil)
}

func newURLSRequest(ctx context.Context, baseURL string, req *URLSRequest) (*http.Request, error) {
	URL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	URL.Path = "/api/v1/test-list/urls"
	query := url.Values{}
	if req.Categories != "" {
		query.Add("categories", req.Categories)
	}
	if req.CountryCode != "" {
		query.Add("country_code", req.CountryCode)
	}
	if req.Limit != 0 {
		query.Add("limit", fmt.Sprintf("%d", req.Limit))
	}
	URL.RawQuery = query.Encode()
	return http.NewRequestWithContext(ctx, "GET", URL.String(), nil)
}

func newOpenReportRequest(ctx context.Context, baseURL string, req *OpenReportRequest) (*http.Request, error) {
	URL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	URL.Path = "/report"
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	out, err := http.NewRequestWithContext(ctx, "POST", URL.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	out.Header.Set("Content-Type", "application/json")
	return out, nil
}

func newSubmitMeasurementRequest(ctx context.Context, baseURL string, req *SubmitMeasurementRequest) (*http.Request, error) {
	URL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	tmpl, err := template.New("urlpath").Parse("/report/{{ .ReportID }}")
	if err != nil {
		return nil, err
	}
	var urlpath strings.Builder
	err = tmpl.Execute(&urlpath, req)
	if err != nil {
		return nil, err
	}
	URL.Path = urlpath.String()
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	out, err := http.NewRequestWithContext(ctx, "POST", URL.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	out.Header.Set("Content-Type", "application/json")
	return out, nil
}

