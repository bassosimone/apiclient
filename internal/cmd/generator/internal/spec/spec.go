// Package spec describes all the implemented APIs. You need to
// edit this package to add/remove/modify APIs. Once you are done editing,
// remember to run `go generate ./...` to regenerate apiclient files.
package spec

import (
	"github.com/bassosimone/apiclient/internal/imodel"
	"github.com/bassosimone/apiclient/model"
)

// URLPath describes the URLPath.
type URLPath struct {
	IsTemplate bool
	Value      string
	InSwagger  string
}

// Descriptor is an API descriptor.
type Descriptor struct {
	Name          string
	Method        string
	URLPath       URLPath
	Private       bool
	Request       interface{}
	Response      interface{}
	RequiresLogin bool
}

// Descriptors contains all descriptors.
var Descriptors = []Descriptor{{
	Name:     "CheckReportID",
	Method:   "GET",
	URLPath:  URLPath{Value: "/api/_/check_report_id"},
	Request:  &model.CheckReportIDRequest{},
	Response: &model.CheckReportIDResponse{},
}, {
	Name:     "CheckIn",
	Method:   "POST",
	URLPath:  URLPath{Value: "/api/v1/check-in"},
	Request:  &model.CheckInRequest{},
	Response: &model.CheckInResponse{},
}, {
	Name:     "Login",
	Method:   "POST",
	URLPath:  URLPath{Value: "/api/v1/login"},
	Private:  true,
	Request:  &imodel.LoginRequest{},
	Response: &imodel.LoginResponse{},
}, {
	Name:     "MeasurementMeta",
	Method:   "GET",
	URLPath:  URLPath{Value: "/api/v1/measurement_meta"},
	Request:  &model.MeasurementMetaRequest{},
	Response: &model.MeasurementMetaResponse{},
}, {
	Name:     "Register",
	Method:   "POST",
	URLPath:  URLPath{Value: "/api/v1/register"},
	Private:  true,
	Request:  &imodel.RegisterRequest{},
	Response: &imodel.RegisterResponse{},
}, {
	Name:     "TestHelpers",
	Method:   "GET",
	URLPath:  URLPath{Value: "/api/v1/test-helpers"},
	Request:  &model.TestHelpersRequest{},
	Response: model.TestHelpersResponse{},
}, {
	Name:          "PsiphonConfig",
	Method:        "GET",
	URLPath:       URLPath{Value: "/api/v1/test-list/psiphon-config"},
	Request:       &model.PsiphonConfigRequest{},
	Response:      model.PsiphonConfigResponse{},
	RequiresLogin: true,
}, {
	Name:          "TorTargets",
	Method:        "GET",
	URLPath:       URLPath{Value: "/api/v1/test-list/tor-targets"},
	Request:       &model.TorTargetsRequest{},
	Response:      model.TorTargetsResponse{},
	RequiresLogin: true,
}, {
	Name:     "URLs",
	Method:   "GET",
	URLPath:  URLPath{Value: "/api/v1/test-list/urls"},
	Request:  &model.URLsRequest{},
	Response: &model.URLsResponse{},
}, {
	Name:     "OpenReport",
	Method:   "POST",
	URLPath:  URLPath{Value: "/report"},
	Request:  &model.OpenReportRequest{},
	Response: &model.OpenReportResponse{},
}, {
	Name:   "SubmitMeasurement",
	Method: "POST",
	URLPath: URLPath{
		InSwagger:  "/report/{report_id}",
		IsTemplate: true,
		Value:      "/report/{{ .ReportID }}",
	},
	Request:  &model.SubmitMeasurementRequest{},
	Response: &model.SubmitMeasurementResponse{},
}}
