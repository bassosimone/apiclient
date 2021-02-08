// Package apimodel describes all the implemented APIs. You need to
// edit this package to add/remove/modify APIs. Once you are done editing,
// remember to run `go generate ./...` to regenerate apiclient files.
package apimodel

import "github.com/bassosimone/apiclient/internal/datamodel"

// URLPath describes the URLPath.
type URLPath struct {
	IsTemplate bool
	Value      string
	InSwagger  string
}

// Descriptor is an API descriptor.
type Descriptor struct {
	Method        string
	URLPath       URLPath
	Request       interface{}
	Response      interface{}
	RequiresLogin bool
}

// Descriptors contains all descriptors.
var Descriptors = []Descriptor{{
	Method:   "GET",
	URLPath:  URLPath{Value: "/api/_/check_report_id"},
	Request:  datamodel.CheckReportIDRequest{},
	Response: datamodel.CheckReportIDResponse{},
}, {
	Method:   "POST",
	URLPath:  URLPath{Value: "/api/v1/check-in"},
	Request:  datamodel.CheckInRequest{},
	Response: datamodel.CheckInResponse{},
}, {
	Method:   "POST",
	URLPath:  URLPath{Value: "/api/v1/login"},
	Request:  datamodel.LoginRequest{},
	Response: datamodel.LoginResponse{},
}, {
	Method:   "GET",
	URLPath:  URLPath{Value: "/api/v1/measurement_meta"},
	Request:  datamodel.MeasurementMetaRequest{},
	Response: datamodel.MeasurementMetaResponse{},
}, {
	Method:   "POST",
	URLPath:  URLPath{Value: "/api/v1/register"},
	Request:  datamodel.RegisterRequest{},
	Response: datamodel.RegisterResponse{},
}, {
	Method:   "GET",
	URLPath:  URLPath{Value: "/api/v1/test-helpers"},
	Request:  datamodel.TestHelpersRequest{},
	Response: datamodel.TestHelpersResponse{},
}, {
	Method:        "GET",
	URLPath:       URLPath{Value: "/api/v1/test-list/psiphon-config"},
	Request:       datamodel.PsiphonConfigRequest{},
	Response:      datamodel.PsiphonConfigResponse{},
	RequiresLogin: true,
}, {
	Method:        "GET",
	URLPath:       URLPath{Value: "/api/v1/test-list/tor-targets"},
	Request:       datamodel.TorTargetsRequest{},
	Response:      datamodel.TorTargetsResponse{},
	RequiresLogin: true,
}, {
	Method:   "GET",
	URLPath:  URLPath{Value: "/api/v1/test-list/urls"},
	Request:  datamodel.URLsRequest{},
	Response: datamodel.URLsResponse{},
}, {
	Method:   "POST",
	URLPath:  URLPath{Value: "/report"},
	Request:  datamodel.OpenReportRequest{},
	Response: datamodel.OpenReportResponse{},
}, {
	Method: "POST",
	URLPath: URLPath{
		InSwagger:  "/report/{report_id}",
		IsTemplate: true,
		Value:      "/report/{{ .ReportID }}",
	},
	Request:  datamodel.SubmitMeasurementRequest{},
	Response: datamodel.SubmitMeasurementResponse{},
}}
