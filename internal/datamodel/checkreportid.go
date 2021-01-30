package datamodel

// CheckReportIDRequest is the CheckReportID request.
type CheckReportIDRequest struct {
	ReportID string `query:"report_id"`
}

// CheckReportIDResponse is the CheckReportID response.
type CheckReportIDResponse struct {
	Found bool `json:"found"`
}
