// Code generated by go generate; DO NOT EDIT.
// 2021-02-16 17:26:54.827558415 +0100 CET m=+0.000143712

package apiclient

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

//go:generate go run ./internal/gennewresponse/...

func (api *CheckReportIDAPI) newResponse(resp *http.Response, err error) (*CheckReportIDResponse, error) {
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
	unmarshal := json.Unmarshal
	if api.unmarshal != nil {
		unmarshal = api.unmarshal
	}
	if err := unmarshal(data, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (api *CheckInAPI) newResponse(resp *http.Response, err error) (*CheckInResponse, error) {
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
	unmarshal := json.Unmarshal
	if api.unmarshal != nil {
		unmarshal = api.unmarshal
	}
	if err := unmarshal(data, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (api *LoginAPI) newResponse(resp *http.Response, err error) (*LoginResponse, error) {
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
	unmarshal := json.Unmarshal
	if api.unmarshal != nil {
		unmarshal = api.unmarshal
	}
	if err := unmarshal(data, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (api *MeasurementMetaAPI) newResponse(resp *http.Response, err error) (*MeasurementMetaResponse, error) {
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
	unmarshal := json.Unmarshal
	if api.unmarshal != nil {
		unmarshal = api.unmarshal
	}
	if err := unmarshal(data, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (api *RegisterAPI) newResponse(resp *http.Response, err error) (*RegisterResponse, error) {
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
	unmarshal := json.Unmarshal
	if api.unmarshal != nil {
		unmarshal = api.unmarshal
	}
	if err := unmarshal(data, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (api *TestHelpersAPI) newResponse(resp *http.Response, err error) (TestHelpersResponse, error) {
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
	unmarshal := json.Unmarshal
	if api.unmarshal != nil {
		unmarshal = api.unmarshal
	}
	if err := unmarshal(data, &out); err != nil {
		return nil, err
	}
	if out == nil {
		return nil, ErrJSONLiteralNull
	}
	return out, nil
}

func (api *PsiphonConfigAPI) newResponse(resp *http.Response, err error) (PsiphonConfigResponse, error) {
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
	unmarshal := json.Unmarshal
	if api.unmarshal != nil {
		unmarshal = api.unmarshal
	}
	if err := unmarshal(data, &out); err != nil {
		return nil, err
	}
	if out == nil {
		return nil, ErrJSONLiteralNull
	}
	return out, nil
}

func (api *TorTargetsAPI) newResponse(resp *http.Response, err error) (TorTargetsResponse, error) {
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
	unmarshal := json.Unmarshal
	if api.unmarshal != nil {
		unmarshal = api.unmarshal
	}
	if err := unmarshal(data, &out); err != nil {
		return nil, err
	}
	if out == nil {
		return nil, ErrJSONLiteralNull
	}
	return out, nil
}

func (api *URLsAPI) newResponse(resp *http.Response, err error) (*URLsResponse, error) {
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
	unmarshal := json.Unmarshal
	if api.unmarshal != nil {
		unmarshal = api.unmarshal
	}
	if err := unmarshal(data, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (api *OpenReportAPI) newResponse(resp *http.Response, err error) (*OpenReportResponse, error) {
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
	unmarshal := json.Unmarshal
	if api.unmarshal != nil {
		unmarshal = api.unmarshal
	}
	if err := unmarshal(data, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (api *SubmitMeasurementAPI) newResponse(resp *http.Response, err error) (*SubmitMeasurementResponse, error) {
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
	unmarshal := json.Unmarshal
	if api.unmarshal != nil {
		unmarshal = api.unmarshal
	}
	if err := unmarshal(data, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
