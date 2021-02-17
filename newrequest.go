// Code generated by go generate; DO NOT EDIT.
// 2021-02-17 09:49:09.912551912 +0100 CET m=+0.000147540

package apiclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

//go:generate go run ./internal/gennewrequest/...

func (api *checkReportIDAPI) newRequest(ctx context.Context, baseURL string, req *CheckReportIDRequest) (*http.Request, error) {
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
	newRequest := http.NewRequestWithContext
	if api.NewRequest != nil {
		newRequest = api.NewRequest
	}
	return newRequest(ctx, "GET", URL.String(), nil)
}

func (api *checkInAPI) newRequest(ctx context.Context, baseURL string, req *CheckInRequest) (*http.Request, error) {
	URL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	URL.Path = "/api/v1/check-in"
	marshal := json.Marshal
	if api.marshal != nil {
		marshal = api.marshal
	}
	body, err := marshal(req)
	if err != nil {
		return nil, err
	}
	newRequest := http.NewRequestWithContext
	if api.NewRequest != nil {
		newRequest = api.NewRequest
	}
	out, err := newRequest(ctx, "POST", URL.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	out.Header.Set("Content-Type", "application/json")
	return out, nil
}

func (api *loginAPI) newRequest(ctx context.Context, baseURL string, req *LoginRequest) (*http.Request, error) {
	URL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	URL.Path = "/api/v1/login"
	marshal := json.Marshal
	if api.marshal != nil {
		marshal = api.marshal
	}
	body, err := marshal(req)
	if err != nil {
		return nil, err
	}
	newRequest := http.NewRequestWithContext
	if api.NewRequest != nil {
		newRequest = api.NewRequest
	}
	out, err := newRequest(ctx, "POST", URL.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	out.Header.Set("Content-Type", "application/json")
	return out, nil
}

func (api *measurementMetaAPI) newRequest(ctx context.Context, baseURL string, req *MeasurementMetaRequest) (*http.Request, error) {
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
	newRequest := http.NewRequestWithContext
	if api.NewRequest != nil {
		newRequest = api.NewRequest
	}
	return newRequest(ctx, "GET", URL.String(), nil)
}

func (api *registerAPI) newRequest(ctx context.Context, baseURL string, req *RegisterRequest) (*http.Request, error) {
	URL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	URL.Path = "/api/v1/register"
	marshal := json.Marshal
	if api.marshal != nil {
		marshal = api.marshal
	}
	body, err := marshal(req)
	if err != nil {
		return nil, err
	}
	newRequest := http.NewRequestWithContext
	if api.NewRequest != nil {
		newRequest = api.NewRequest
	}
	out, err := newRequest(ctx, "POST", URL.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	out.Header.Set("Content-Type", "application/json")
	return out, nil
}

func (api *testHelpersAPI) newRequest(ctx context.Context, baseURL string, req *TestHelpersRequest) (*http.Request, error) {
	URL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	URL.Path = "/api/v1/test-helpers"
	newRequest := http.NewRequestWithContext
	if api.NewRequest != nil {
		newRequest = api.NewRequest
	}
	return newRequest(ctx, "GET", URL.String(), nil)
}

func (api *psiphonConfigAPI) newRequest(ctx context.Context, baseURL string, req *PsiphonConfigRequest) (*http.Request, error) {
	URL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	URL.Path = "/api/v1/test-list/psiphon-config"
	newRequest := http.NewRequestWithContext
	if api.NewRequest != nil {
		newRequest = api.NewRequest
	}
	return newRequest(ctx, "GET", URL.String(), nil)
}

func (api *torTargetsAPI) newRequest(ctx context.Context, baseURL string, req *TorTargetsRequest) (*http.Request, error) {
	URL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	URL.Path = "/api/v1/test-list/tor-targets"
	newRequest := http.NewRequestWithContext
	if api.NewRequest != nil {
		newRequest = api.NewRequest
	}
	return newRequest(ctx, "GET", URL.String(), nil)
}

func (api *urlsAPI) newRequest(ctx context.Context, baseURL string, req *URLsRequest) (*http.Request, error) {
	URL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	URL.Path = "/api/v1/test-list/urls"
	query := url.Values{}
	if req.CategoryCodes != "" {
		query.Add("category_codes", req.CategoryCodes)
	}
	if req.CountryCode != "" {
		query.Add("country_code", req.CountryCode)
	}
	if req.Limit != 0 {
		query.Add("limit", fmt.Sprintf("%d", req.Limit))
	}
	URL.RawQuery = query.Encode()
	newRequest := http.NewRequestWithContext
	if api.NewRequest != nil {
		newRequest = api.NewRequest
	}
	return newRequest(ctx, "GET", URL.String(), nil)
}

func (api *openReportAPI) newRequest(ctx context.Context, baseURL string, req *OpenReportRequest) (*http.Request, error) {
	URL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	URL.Path = "/report"
	marshal := json.Marshal
	if api.marshal != nil {
		marshal = api.marshal
	}
	body, err := marshal(req)
	if err != nil {
		return nil, err
	}
	newRequest := http.NewRequestWithContext
	if api.NewRequest != nil {
		newRequest = api.NewRequest
	}
	out, err := newRequest(ctx, "POST", URL.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	out.Header.Set("Content-Type", "application/json")
	return out, nil
}

func (api *submitMeasurementAPI) newRequest(ctx context.Context, baseURL string, req *SubmitMeasurementRequest) (*http.Request, error) {
	URL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	newTemplate := newStdlibTextTemplate
	if api.newTemplate != nil {
		newTemplate = api.newTemplate
	}
	tmpl, err := newTemplate("urlpath").Parse("/report/{{ .ReportID }}")
	if err != nil {
		return nil, err
	}
	var urlpath strings.Builder
	err = tmpl.Execute(&urlpath, req)
	if err != nil {
		return nil, err
	}
	URL.Path = urlpath.String()
	marshal := json.Marshal
	if api.marshal != nil {
		marshal = api.marshal
	}
	body, err := marshal(req)
	if err != nil {
		return nil, err
	}
	newRequest := http.NewRequestWithContext
	if api.NewRequest != nil {
		newRequest = api.NewRequest
	}
	out, err := newRequest(ctx, "POST", URL.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	out.Header.Set("Content-Type", "application/json")
	return out, nil
}

