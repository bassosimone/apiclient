// Code generated by go generate; DO NOT EDIT.
// 2021-01-30 20:29:58.564702274 +0100 CET m=+0.000124347

package apiclient

import "context"

//go:generate go run ./internal/genapi/...

// GETCheckReportID implements the GET /api/_/check_report_id API
func (c Client) GETCheckReportID(ctx context.Context, in *CheckReportIDRequest) (*CheckReportIDResponse, error) {
	req, err := newCheckReportIDRequest(ctx, c.BaseURL, in)
	if err != nil {
		return nil, err
	}
	if c.Accept != "" {
		req.Header.Add("Accept", c.Accept)
	}
	if c.Authorization != "" {
		req.Header.Add("Authorization", c.Authorization)
	}
	req.Header.Add("User-Agent", c.UserAgent)
	return newCheckReportIDResponse(c.HTTPClient.Do(req))
}

// POSTCheckIn implements the POST /api/v1/check-in API
func (c Client) POSTCheckIn(ctx context.Context, in *CheckInRequest) (*CheckInResponse, error) {
	req, err := newCheckInRequest(ctx, c.BaseURL, in)
	if err != nil {
		return nil, err
	}
	if c.Accept != "" {
		req.Header.Add("Accept", c.Accept)
	}
	if c.Authorization != "" {
		req.Header.Add("Authorization", c.Authorization)
	}
	req.Header.Add("User-Agent", c.UserAgent)
	return newCheckInResponse(c.HTTPClient.Do(req))
}

// POSTLogin implements the POST /api/v1/login API
func (c Client) POSTLogin(ctx context.Context, in *LoginRequest) (*LoginResponse, error) {
	req, err := newLoginRequest(ctx, c.BaseURL, in)
	if err != nil {
		return nil, err
	}
	if c.Accept != "" {
		req.Header.Add("Accept", c.Accept)
	}
	if c.Authorization != "" {
		req.Header.Add("Authorization", c.Authorization)
	}
	req.Header.Add("User-Agent", c.UserAgent)
	return newLoginResponse(c.HTTPClient.Do(req))
}

// GETMeasurementMeta implements the GET /api/v1/measurement_meta API
func (c Client) GETMeasurementMeta(ctx context.Context, in *MeasurementMetaRequest) (*MeasurementMetaResponse, error) {
	req, err := newMeasurementMetaRequest(ctx, c.BaseURL, in)
	if err != nil {
		return nil, err
	}
	if c.Accept != "" {
		req.Header.Add("Accept", c.Accept)
	}
	if c.Authorization != "" {
		req.Header.Add("Authorization", c.Authorization)
	}
	req.Header.Add("User-Agent", c.UserAgent)
	return newMeasurementMetaResponse(c.HTTPClient.Do(req))
}

// GETPsiphonConfig implements the GET /api/v1/test-list/psiphon-config API
func (c Client) GETPsiphonConfig(ctx context.Context, in *PsiphonConfigRequest) (PsiphonConfigResponse, error) {
	req, err := newPsiphonConfigRequest(ctx, c.BaseURL, in)
	if err != nil {
		return nil, err
	}
	if c.Accept != "" {
		req.Header.Add("Accept", c.Accept)
	}
	if c.Authorization != "" {
		req.Header.Add("Authorization", c.Authorization)
	}
	req.Header.Add("User-Agent", c.UserAgent)
	return newPsiphonConfigResponse(c.HTTPClient.Do(req))
}

// GETTorTargets implements the GET /api/v1/test-list/tor-targets API
func (c Client) GETTorTargets(ctx context.Context, in *TorTargetsRequest) (TorTargetsResponse, error) {
	req, err := newTorTargetsRequest(ctx, c.BaseURL, in)
	if err != nil {
		return nil, err
	}
	if c.Accept != "" {
		req.Header.Add("Accept", c.Accept)
	}
	if c.Authorization != "" {
		req.Header.Add("Authorization", c.Authorization)
	}
	req.Header.Add("User-Agent", c.UserAgent)
	return newTorTargetsResponse(c.HTTPClient.Do(req))
}

// POSTRegister implements the POST /api/v1/register API
func (c Client) POSTRegister(ctx context.Context, in *RegisterRequest) (*RegisterResponse, error) {
	req, err := newRegisterRequest(ctx, c.BaseURL, in)
	if err != nil {
		return nil, err
	}
	if c.Accept != "" {
		req.Header.Add("Accept", c.Accept)
	}
	if c.Authorization != "" {
		req.Header.Add("Authorization", c.Authorization)
	}
	req.Header.Add("User-Agent", c.UserAgent)
	return newRegisterResponse(c.HTTPClient.Do(req))
}

// POSTOpenReport implements the POST /report API
func (c Client) POSTOpenReport(ctx context.Context, in *OpenReportRequest) (*OpenReportResponse, error) {
	req, err := newOpenReportRequest(ctx, c.BaseURL, in)
	if err != nil {
		return nil, err
	}
	if c.Accept != "" {
		req.Header.Add("Accept", c.Accept)
	}
	if c.Authorization != "" {
		req.Header.Add("Authorization", c.Authorization)
	}
	req.Header.Add("User-Agent", c.UserAgent)
	return newOpenReportResponse(c.HTTPClient.Do(req))
}

// POSTSubmitMeasurement implements the POST /report/{{ .ReportID }} API
func (c Client) POSTSubmitMeasurement(ctx context.Context, in *SubmitMeasurementRequest) (*SubmitMeasurementResponse, error) {
	req, err := newSubmitMeasurementRequest(ctx, c.BaseURL, in)
	if err != nil {
		return nil, err
	}
	if c.Accept != "" {
		req.Header.Add("Accept", c.Accept)
	}
	if c.Authorization != "" {
		req.Header.Add("Authorization", c.Authorization)
	}
	req.Header.Add("User-Agent", c.UserAgent)
	return newSubmitMeasurementResponse(c.HTTPClient.Do(req))
}

// GETTestHelpers implements the GET /api/v1/test-helpers API
func (c Client) GETTestHelpers(ctx context.Context, in *TestHelpersRequest) (TestHelpersResponse, error) {
	req, err := newTestHelpersRequest(ctx, c.BaseURL, in)
	if err != nil {
		return nil, err
	}
	if c.Accept != "" {
		req.Header.Add("Accept", c.Accept)
	}
	if c.Authorization != "" {
		req.Header.Add("Authorization", c.Authorization)
	}
	req.Header.Add("User-Agent", c.UserAgent)
	return newTestHelpersResponse(c.HTTPClient.Do(req))
}

// GETURLS implements the GET /api/v1/test-list/urls API
func (c Client) GETURLS(ctx context.Context, in *URLSRequest) (*URLSResponse, error) {
	req, err := newURLSRequest(ctx, c.BaseURL, in)
	if err != nil {
		return nil, err
	}
	if c.Accept != "" {
		req.Header.Add("Accept", c.Accept)
	}
	if c.Authorization != "" {
		req.Header.Add("Authorization", c.Authorization)
	}
	req.Header.Add("User-Agent", c.UserAgent)
	return newURLSResponse(c.HTTPClient.Do(req))
}

