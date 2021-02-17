package apiclient

import (
	"context"
	"testing"

	"github.com/bassosimone/apiclient/model"
)

// TODO(bassosimone): write test for URL path

func TestMeasurementMetaNewRequestURLOkay(t *testing.T) {
	api := &measurementMetaAPI{}
	apireq := &model.MeasurementMetaRequest{
		ReportID: "abc",
		Full:     true,
		Input:    "xyz",
	}
	ctx := context.Background()
	req, err := api.newRequest(ctx, apireq)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	if q.Get("report_id") != "abc" {
		t.Fatal("invalid report_id")
	}
	if q.Get("full") != "true" {
		t.Fatal("invalid full")
	}
	if q.Get("input") != "xyz" {
		t.Fatal("invalid xyz")
	}
}

func TestURLsNewRequestURLOkay(t *testing.T) {
	api := &urlsAPI{}
	apireq := &model.URLsRequest{
		CategoryCodes: "HUMR,HACK",
		CountryCode:   "IT",
		Limit:         128,
	}
	ctx := context.Background()
	req, err := api.newRequest(ctx, apireq)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	if q.Get("category_codes") != "HUMR,HACK" {
		t.Fatal("invalid category_codes")
	}
	if q.Get("country_code") != "IT" {
		t.Fatal("invalid country_code")
	}
	if q.Get("limit") != "128" {
		t.Fatal("invalid limit")
	}
}
