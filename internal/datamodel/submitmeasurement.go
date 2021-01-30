package datamodel

// SubmitMeasurementRequest is the SubmitMeasurement request.
type SubmitMeasurementRequest struct {
	ReportID string
	Format   string      `json:"format"`
	Content  interface{} `json:"content"`
}

// SubmitMeasurementResponse is the SubmitMeasurement response.
type SubmitMeasurementResponse struct{}
