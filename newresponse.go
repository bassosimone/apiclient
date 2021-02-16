// Code generated by go generate; DO NOT EDIT.
// 2021-02-16 07:23:27.150087263 +0100 CET m=+0.000268151

package apiclient

import (
	"encoding/json"
	"fmt"
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
		return nil, fmt.Errorf("%w: %d", ErrHTTPFailure, resp.StatusCode)
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

func newCheckInResponse(resp *http.Response, err error) (*CheckInResponse, error) {
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%w: %d", ErrHTTPFailure, resp.StatusCode)
	}
	defer resp.Body.Close()
	reader := io.LimitReader(resp.Body, 4<<20)
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var out CheckInResponse
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func newLoginResponse(resp *http.Response, err error) (*LoginResponse, error) {
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%w: %d", ErrHTTPFailure, resp.StatusCode)
	}
	defer resp.Body.Close()
	reader := io.LimitReader(resp.Body, 4<<20)
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var out LoginResponse
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
		return nil, fmt.Errorf("%w: %d", ErrHTTPFailure, resp.StatusCode)
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

func newRegisterResponse(resp *http.Response, err error) (*RegisterResponse, error) {
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%w: %d", ErrHTTPFailure, resp.StatusCode)
	}
	defer resp.Body.Close()
	reader := io.LimitReader(resp.Body, 4<<20)
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var out RegisterResponse
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func newTestHelpersResponse(resp *http.Response, err error) (TestHelpersResponse, error) {
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%w: %d", ErrHTTPFailure, resp.StatusCode)
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
	if out == nil {
		return nil, ErrJSONLiteralNull
	}
	return out, nil
}

func newPsiphonConfigResponse(resp *http.Response, err error) (PsiphonConfigResponse, error) {
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%w: %d", ErrHTTPFailure, resp.StatusCode)
	}
	defer resp.Body.Close()
	reader := io.LimitReader(resp.Body, 4<<20)
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var out PsiphonConfigResponse
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, err
	}
	if out == nil {
		return nil, ErrJSONLiteralNull
	}
	return out, nil
}

func newTorTargetsResponse(resp *http.Response, err error) (TorTargetsResponse, error) {
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%w: %d", ErrHTTPFailure, resp.StatusCode)
	}
	defer resp.Body.Close()
	reader := io.LimitReader(resp.Body, 4<<20)
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var out TorTargetsResponse
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, err
	}
	if out == nil {
		return nil, ErrJSONLiteralNull
	}
	return out, nil
}

func newURLsResponse(resp *http.Response, err error) (*URLsResponse, error) {
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%w: %d", ErrHTTPFailure, resp.StatusCode)
	}
	defer resp.Body.Close()
	reader := io.LimitReader(resp.Body, 4<<20)
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var out URLsResponse
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
		return nil, fmt.Errorf("%w: %d", ErrHTTPFailure, resp.StatusCode)
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
		return nil, fmt.Errorf("%w: %d", ErrHTTPFailure, resp.StatusCode)
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

