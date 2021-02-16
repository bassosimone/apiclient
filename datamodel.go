// Code generated by go generate; DO NOT EDIT.
// 2021-02-16 09:52:11.739636946 +0100 CET m=+0.000163765

package apiclient

//go:generate go run ./internal/gendatamodel/...

// CheckInRequest is the check-in API request
type CheckInRequest struct {
	Charging        bool                           `json:"charging"`
	OnWiFi          bool                           `json:"on_wifi"`
	Platform        string                         `json:"platform"`
	ProbeASN        string                         `json:"probe_asn"`
	ProbeCC         string                         `json:"probe_cc"`
	RunType         string                         `json:"run_type"`
	SoftwareName    string                         `json:"software_name"`
	SoftwareVersion string                         `json:"software_version"`
	WebConnectivity *CheckInRequestWebConnectivity `json:"web_connectivity"`
}

// CheckInRequestWebConnectivity contains WebConnectivity
// specific parameters to include into CheckInRequest
type CheckInRequestWebConnectivity struct {
	CategoryCodes []string `json:"category_codes"`
}

// CheckInResponse is the check-in API response
type CheckInResponse struct {
	WebConnectivity *CheckInResponseWebConnectivity `json:"web_connectivity"`
}

// CheckInResponseURLInfo contains information about an URL.
type CheckInResponseURLInfo struct {
	CategoryCode string `json:"category_code"`
	CountryCode  string `json:"country_code"`
	URL          string `json:"url"`
}

// CheckInResponseWebConnectivity contains WebConnectivity
// specific information of a CheckInResponse
type CheckInResponseWebConnectivity struct {
	ReportID string                   `json:"report_id"`
	URLs     []CheckInResponseURLInfo `json:"urls"`
}

// CheckReportIDRequest is the CheckReportID request.
type CheckReportIDRequest struct {
	ReportID string `query:"report_id" required:"true"`
}

// CheckReportIDResponse is the CheckReportID response.
type CheckReportIDResponse struct {
	Error string `json:"error"`
	Found bool   `json:"found"`
	V     int64  `json:"v"`
}

// LoginRequest is the login API request
type LoginRequest struct {
	ClientID string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse is the login API response
type LoginResponse struct {
	Expire string `json:"expire"`
	Token  string `json:"token"`
}

// MeasurementMetaRequest is the MeasurementMeta Request.
type MeasurementMetaRequest struct {
	ReportID string `query:"report_id" required:"true"`
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

// PsiphonConfigRequest is the request for the PsiphonConfig API
type PsiphonConfigRequest struct{}

// PsiphonConfigResponse is the response from the PsiphonConfig API
type PsiphonConfigResponse map[string]interface{}

// RegisterRequest is the request for the Register API.
type RegisterRequest struct {
	AvailableBandwidth string   `json:"available_bandwidth,omitempty"`
	DeviceToken        string   `json:"device_token,omitempty"`
	Language           string   `json:"language,omitempty"`
	NetworkType        string   `json:"network_type,omitempty"`
	Platform           string   `json:"platform"`
	ProbeASN           string   `json:"probe_asn"`
	ProbeCC            string   `json:"probe_cc"`
	ProbeFamily        string   `json:"probe_family,omitempty"`
	ProbeTimezone      string   `json:"probe_timezone,omitempty"`
	SoftwareName       string   `json:"software_name"`
	SoftwareVersion    string   `json:"software_version"`
	SupportedTests     []string `json:"supported_tests"`
}

// RegisterResponse is the response from the Register API.
type RegisterResponse struct {
	ClientID string `json:"client_id"`
}

// SubmitMeasurementRequest is the SubmitMeasurement request.
type SubmitMeasurementRequest struct {
	ReportID string      `path:"report_id"`
	Format   string      `json:"format"`
	Content  interface{} `json:"content"`
}

// SubmitMeasurementResponse is the SubmitMeasurement response.
type SubmitMeasurementResponse struct {
	MeasurementUID string `json:"measurement_uid"`
}

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
type TestHelpersResponse map[string][]TestHelpersHelperInfo

// TorTargetsRequest is a request for the TorTargets API.
type TorTargetsRequest struct{}

// TorTargetsResponse is the response from the TorTargets API.
type TorTargetsResponse map[string]TorTargetsTarget

// TorTargetsTarget is a target for the tor experiment.
type TorTargetsTarget struct {
	Address  string              `json:"address"`
	Name     string              `json:"name"`
	Params   map[string][]string `json:"params"`
	Protocol string              `json:"protocol"`
	Source   string              `json:"source"`
}

// URLsMetadata contains metadata in the URLs response.
type URLsMetadata struct {
	Count int64 `json:"count"`
}

// URLsRequest is the URLs request.
type URLsRequest struct {
	CategoryCodes string `query:"category_codes"`
	CountryCode   string `query:"country_code"`
	Limit         int64  `query:"limit"`
}

// URLsResponse is the URLs response.
type URLsResponse struct {
	Metadata URLsMetadata      `json:"metadata"`
	Results  []URLsResponseURL `json:"results"`
}

// URLsResponseURL is a single URL in the URLs response.
type URLsResponseURL struct {
	CategoryCode string `json:"category_code"`
	CountryCode  string `json:"country_code"`
	URL          string `json:"url"`
}

