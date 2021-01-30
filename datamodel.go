// Code generated by go generate; DO NOT EDIT.
// 2021-01-30 17:47:46.347301185 +0100 CET m=+0.000276319

package apiclient

//go:generate go run ./internal/gendatamodel/...

// CheckReportIDRequest is the CheckReportID request.
type CheckReportIDRequest struct {
	ReportID string `query:"report_id" mandatory:"true"`
}

// CheckReportIDResponse is the CheckReportID response.
type CheckReportIDResponse struct {
	Found bool `json:"found"`
}

// MeasurementMetaRequest is the MeasurementMeta Request.
type MeasurementMetaRequest struct {
	ReportID string `query:"report_id" mandatory:"true"`
	Full     bool   `query:"full"`
	Input    string `query:"input"`
}

// MeasurementMetaResponse is the MeasurementMeta Response.
type MeasurementMetaResponse struct {
	Anomaly              bool   `json:"anomaly"`
	CategoryCode         string `json:"category_code"`
	Confirmed            bool   `json:"confirmed"`
	Failure              bool   `json:"failure"`
	Input                string `json:"input"`
	MeasurementStartTime string `json:"measurement_start_time"`
	ProbeASN             int64  `json:"probe_asn"`
	ProbeCC              string `json:"probe_cc"`
	RawMeasurement       string `json:"raw_measurement"`
	ReportID             string `json:"report_id"`
	Scores               string `json:"scores"`
	TestName             string `json:"test_name"`
	TestStartTime        string `json:"test_start_time"`
}

// OpenReportRequest is the OpenReport request.
type OpenReportRequest struct {
	DataFormatVersion string `json:"data_format_version"`
	Format            string `json:"format"`
	ProbeASN          string `json:"probe_asn"`
	ProbeCC           string `json:"probe_cc"`
	SoftwareName      string `json:"software_name"`
	SoftwareVersion   string `json:"software_version"`
	TestName          string `json:"test_name"`
	TestStartTime     string `json:"test_start_time"`
	TestVersion       string `json:"test_version"`
}

// OpenReportResponse is the OpenReport response.
type OpenReportResponse struct {
	ReportID         string   `json:"report_id"`
	SupportedFormats []string `json:"supported_formats"`
}

// SubmitMeasurementRequest is the SubmitMeasurement request.
type SubmitMeasurementRequest struct {
	ReportID string
	Format   string      `json:"format"`
	Content  interface{} `json:"content"`
}

// SubmitMeasurementResponse is the SubmitMeasurement response.
type SubmitMeasurementResponse struct{}

// TestHelpersHelperInfo is a single helper within the
// response returned by the TestHelpers API.
type TestHelpersHelperInfo struct {
	Address string `json:"address"`
	Type    string `json:"type"`
	Front   string `json:"front,omitempty"`
}

// TestHelpersRequest is the TestHelpers request.
type TestHelpersRequest struct{}

// TestHelpersResponse is the TestHelpers response.
type TestHelpersResponse struct {
	HTTPReturnJSONHeaders []TestHelpersHelperInfo `json:"http-return-json-headers"`
	TCPEcho               []TestHelpersHelperInfo `json:"tcp-echo"`
	WebConnectivity       []TestHelpersHelperInfo `json:"web-connectivity"`
}

// URLSRequest is the URLS request.
type URLSRequest struct {
	Categories  string `query:"categories"`
	CountryCode string `query:"country_code"`
	Limit       int64  `query:"limit"`
}

// URLSResponse is the URLS response.
type URLSResponse struct {
	Results []URLSResponseURL `json:"results"`
}

// URLSResponseURL is a single URL in the URLS response.
type URLSResponseURL struct {
	CategoryCode string `json:"category_code"`
	CountryCode  string `json:"country_code"`
	URL          string `json:"url"`
}

