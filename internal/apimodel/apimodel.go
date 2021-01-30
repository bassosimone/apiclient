// Package apimodel describes all the implemented APIs.
package apimodel

import "github.com/bassosimone/apiclient/internal/datamodel"

// Descriptor is an API descriptor.
type Descriptor struct {
	Method   string
	URLPath  string
	Request  interface{}
	Response interface{}
}

// Descriptors contains all descriptors.
var Descriptors = []Descriptor{{
	Method:   "GET",
	URLPath:  "/api/_/check_report_id",
	Request:  datamodel.CheckReportIDRequest{},
	Response: datamodel.CheckReportIDResponse{},
}, {
	Method:   "POST",
	URLPath:  "/api/v1/check-in",
	Request:  datamodel.CheckInRequest{},
	Response: datamodel.CheckInResponse{},
}, {
	Method:   "POST",
	URLPath:  "/api/v1/login",
	Request:  datamodel.LoginRequest{},
	Response: datamodel.LoginResponse{},
}, {
	Method:   "GET",
	URLPath:  "/api/v1/measurement_meta",
	Request:  datamodel.MeasurementMetaRequest{},
	Response: datamodel.MeasurementMetaResponse{},
}, {
	Method:   "GET",
	URLPath:  "/api/v1/test-list/psiphon-config",
	Request:  datamodel.PsiphonConfigRequest{},
	Response: datamodel.PsiphonConfigResponse{},
}, {
	Method:   "POST",
	URLPath:  "/api/v1/register",
	Request:  datamodel.RegisterRequest{},
	Response: datamodel.RegisterResponse{},
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
