package apiclient

import (
	"bytes"
	"encoding/gob"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestRequestMethodDefault(t *testing.T) {
	if requestMethod("") != "GET" {
		t.Fatal("requestMethod is not WAI")
	}
}

func TestNewCacheRecordKeyCannotReadBody(t *testing.T) {
	crc, err := newCacheRecordKey(&http.Request{
		Body: &mockableBodyWithFailure{},
		URL:  &url.URL{},
	})
	if !errors.Is(err, errMocked) {
		t.Fatal("not the error we expected", err)
	}
	if crc != nil {
		t.Fatal("expected nil here")
	}
}

func TestCacheClientDoWithNewCacheRecordFailure(t *testing.T) {
	req := &http.Request{
		Body: &mockableBodyWithFailure{},
		URL:  &url.URL{},
	}
	cc := &cacheClient{}
	resp, err := cc.Do(req)
	if !errors.Is(err, errMocked) {
		t.Fatal("not the error we expected", err)
	}
	if resp != nil {
		t.Fatal("expected nil here")
	}
}

func TestCacheClientDoWithClientDoFailure(t *testing.T) {
	req := &http.Request{URL: &url.URL{}}
	cc := &cacheClient{
		HTTPClient: &mockableHTTPClient{
			Err: errMocked,
		},
		KVStore: &memkvstore{},
	}
	resp, err := cc.Do(req)
	if !errors.Is(err, errMocked) {
		t.Fatal("not the error we expected", err)
	}
	if resp != nil {
		t.Fatal("expected nil here")
	}
}

func TestCacheClientDoWithReadBodyErr(t *testing.T) {
	req := &http.Request{URL: &url.URL{}}
	cc := &cacheClient{
		HTTPClient: &mockableHTTPClient{
			Resp: &http.Response{
				StatusCode: 200,
				Body:       &mockableBodyWithFailure{},
			},
		},
		KVStore: &memkvstore{},
	}
	resp, err := cc.Do(req)
	if !errors.Is(err, errMocked) {
		t.Fatal("not the error we expected", err)
	}
	if resp != nil {
		t.Fatal("expected nil here")
	}
}

func TestCacheClientDoWithCachedResponse(t *testing.T) {
	req := &http.Request{
		URL: &url.URL{
			Scheme:   "http",
			Host:     "www.example.com",
			Path:     "/antani",
			RawQuery: "foo=bar&baz=1",
		},
		Method: "GET",
		Body:   ioutil.NopCloser(strings.NewReader("{\"antani\":true}")),
	}
	recs := []cacheRecord{{
		Key: cacheRecordKey{
			Method:  "GET",
			URLPath: "/antani",
			Query:   "foo=bar&baz=1",
			Request: "{\"antani\":true}",
		},
		Response: "42",
	}}
	var bw bytes.Buffer
	if err := gob.NewEncoder(&bw).Encode(recs); err != nil {
		t.Fatal(err)
	}
	kvstore := &memkvstore{}
	if err := kvstore.Set(cacheKey, bw.Bytes()); err != nil {
		t.Fatal(err)
	}
	cc := &cacheClient{
		HTTPClient: &mockableHTTPClient{
			Err: errMocked,
		},
		KVStore: kvstore,
	}
	resp, err := cc.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatal("unexpected statusCode")
	}
	if resp.Header.Get("X-OONI-APIClient-Cache") != "FALLBACK" {
		t.Fatal("missing cache header")
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "42" {
		t.Fatal("reading unexpected data")
	}
}

func TestAddToCacheBoundedSize(t *testing.T) {
	kvstore := &memkvstore{}
	cc := &cacheClient{
		HTTPClient: &mockableHTTPClient{
			Err: errMocked,
		},
		KVStore: kvstore,
	}
	ff := &fakeFill{}
	for idx := 0; idx < maxCacheIndex*2; idx++ {
		var (
			crk  cacheRecordKey
			body string
		)
		ff.fill(&crk)
		ff.fill(&body)
		if err := cc.addToCache(&crk, body); err != nil {
			t.Fatal(err)
		}
	}
	out := cc.readCache()
	if len(out) != maxCacheIndex {
		t.Fatal("unexpected number of cached entries", len(out))
	}
}

type alwaysFailingDecoder struct{}

func (afd *alwaysFailingDecoder) Decode(e interface{}) error {
	return errMocked
}

func TestCacheDecodeFailure(t *testing.T) {
	kvstore := &memkvstore{}
	cc := &cacheClient{
		HTTPClient: &mockableHTTPClient{
			Err: errMocked,
		},
		KVStore: kvstore,
		newDecoder: func(r io.Reader) decoder {
			return &alwaysFailingDecoder{}
		},
	}
	ff := &fakeFill{}
	var (
		crk  cacheRecordKey
		body string
	)
	ff.fill(&crk)
	ff.fill(&body)
	if err := cc.addToCache(&crk, body); err != nil {
		t.Fatal(err)
	}
	out := cc.readCache()
	if len(out) != 0 {
		t.Fatal("unexpected result", out)
	}
}

type alwaysFailingEncoder struct{}

func (afe *alwaysFailingEncoder) Encode(e interface{}) error {
	return errMocked
}

func TestCacheEncodeFailure(t *testing.T) {
	kvstore := &memkvstore{}
	cc := &cacheClient{
		HTTPClient: &mockableHTTPClient{
			Err: errMocked,
		},
		KVStore: kvstore,
		newEncoder: func(w io.Writer) encoder {
			return &alwaysFailingEncoder{}
		},
	}
	ff := &fakeFill{}
	var (
		crk  cacheRecordKey
		body string
	)
	ff.fill(&crk)
	ff.fill(&body)
	if err := cc.addToCache(&crk, body); !errors.Is(err, errMocked) {
		t.Fatal("not the error we expected", err)
	}
}
