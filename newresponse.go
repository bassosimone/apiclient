// Code generated by go generate; DO NOT EDIT.
// 2021-01-30 17:47:46.726621073 +0100 CET m=+0.000215574

package apiclient

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"io"
	"net/http"
)

//go:generate go run ./internal/gennewresponse/...

func newCheckReportIDResponse(resp *http.Response, err error) (*CheckReportIDResponse, error) {
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("apiclient: http request failed")
	}
	defer resp.Body.Close()
	reader := io.LimitReader(resp.Body, 4<<20)
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var out CheckReportIDResponse
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func newMeasurementMetaResponse(resp *http.Response, err error) (*MeasurementMetaResponse, error) {
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("apiclient: http request failed")
	}
	defer resp.Body.Close()
	reader := io.LimitReader(resp.Body, 4<<20)
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var out MeasurementMetaResponse
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func newOpenReportResponse(resp *http.Response, err error) (*OpenReportResponse, error) {
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("apiclient: http request failed")
	}
	defer resp.Body.Close()
	reader := io.LimitReader(resp.Body, 4<<20)
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var out OpenReportResponse
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func newSubmitMeasurementResponse(resp *http.Response, err error) (*SubmitMeasurementResponse, error) {
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("apiclient: http request failed")
	}
	defer resp.Body.Close()
	reader := io.LimitReader(resp.Body, 4<<20)
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var out SubmitMeasurementResponse
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func newTestHelpersResponse(resp *http.Response, err error) (*TestHelpersResponse, error) {
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("apiclient: http request failed")
	}
	defer resp.Body.Close()
	reader := io.LimitReader(resp.Body, 4<<20)
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var out TestHelpersResponse
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func newURLSResponse(resp *http.Response, err error) (*URLSResponse, error) {
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("apiclient: http request failed")
	}
	defer resp.Body.Close()
	reader := io.LimitReader(resp.Body, 4<<20)
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var out URLSResponse
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

