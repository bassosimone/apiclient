package datamodel

// CheckReportIDRequest is the CheckReportID request.
type CheckReportIDRequest struct {
	ReportID string `query:"report_id" mandatory:"true"`
}

// CheckReportIDResponse is the CheckReportID response.
type CheckReportIDResponse struct {
	Found bool `json:"found"`
}
