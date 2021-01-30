// Code generated by go generate; DO NOT EDIT.
// 2021-01-30 17:35:30.699471593 +0100 CET m=+0.000237412

package apiclient

import "context"

//go:generate go run ./internal/genapi/...

// GETCheckReportID implements the GET /api/_/check_report_id API
func (c Client) GETCheckReportID(ctx context.Context, in *CheckReportIDRequest) (*CheckReportIDResponse, error) {
	req, err := NewCheckReportIDRequest(ctx, c.BaseURL, in)
	if err != nil {
		return nil, err
	}
	return NewCheckReportIDResponse(c.HTTPClient.Do(req))
}

// GETMeasurementMeta implements the GET /api/v1/measurement_meta API
func (c Client) GETMeasurementMeta(ctx context.Context, in *MeasurementMetaRequest) (*MeasurementMetaResponse, error) {
	req, err := NewMeasurementMetaRequest(ctx, c.BaseURL, in)
	if err != nil {
		return nil, err
	}
	return NewMeasurementMetaResponse(c.HTTPClient.Do(req))
}

// POSTOpenReport implements the POST /report API
func (c Client) POSTOpenReport(ctx context.Context, in *OpenReportRequest) (*OpenReportResponse, error) {
	req, err := NewOpenReportRequest(ctx, c.BaseURL, in)
	if err != nil {
		return nil, err
	}
	return NewOpenReportResponse(c.HTTPClient.Do(req))
}

// POSTSubmitMeasurement implements the POST /report/{{ .ReportID }} API
func (c Client) POSTSubmitMeasurement(ctx context.Context, in *SubmitMeasurementRequest) (*SubmitMeasurementResponse, error) {
	req, err := NewSubmitMeasurementRequest(ctx, c.BaseURL, in)
	if err != nil {
		return nil, err
	}
	return NewSubmitMeasurementResponse(c.HTTPClient.Do(req))
}

// GETTestHelpers implements the GET /api/v1/test-helpers API
func (c Client) GETTestHelpers(ctx context.Context, in *TestHelpersRequest) (*TestHelpersResponse, error) {
	req, err := NewTestHelpersRequest(ctx, c.BaseURL, in)
	if err != nil {
		return nil, err
	}
	return NewTestHelpersResponse(c.HTTPClient.Do(req))
}

// GETURLS implements the GET /api/v1/test-list/urls API
func (c Client) GETURLS(ctx context.Context, in *URLSRequest) (*URLSResponse, error) {
	req, err := NewURLSRequest(ctx, c.BaseURL, in)
	if err != nil {
		return nil, err
	}
	return NewURLSResponse(c.HTTPClient.Do(req))
}

