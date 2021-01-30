package apimodel

import "github.com/bassosimone/apiclient/internal/datamodel"

type descriptor struct {
	Method   string
	URLPath  string
	Request  interface{}
	Response interface{}
}

var descriptors = []descriptor{{
	Method:   "GET",
	URLPath:  "/api/_/check_report_id",
	Request:  datamodel.CheckReportIDRequest{},
	Response: datamodel.CheckReportIDResponse{},
}, {
	Method:   "GET",
	URLPath:  "/api/v1/measurement_meta",
	Request:  datamodel.MeasurementMetaRequest{},
	Response: datamodel.MeasurementMetaResponse{},
}, {
	Method:   "POST",
	URLPath:  "/report",
	Request:  datamodel.OpenReportRequest{},
	Response: datamodel.OpenReportResponse{},
}, {
	Method:   "POST",
	URLPath:  "/report/{{ .ReportID }}",
	Request:  datamodel.SubmitMeasurementRequest{},
	Response: datamodel.SubmitMeasurementResponse{},
}, {
	Method:   "GET",
	URLPath:  "/api/v1/test-helpers",
	Request:  datamodel.TestHelpersRequest{},
	Response: datamodel.TestHelpersResponse{},
}, {
	Method:   "GET",
	URLPath:  "/api/v1/test-list/urls",
	Request:  datamodel.URLSRequest{},
	Response: datamodel.URLSResponse{},
}}
