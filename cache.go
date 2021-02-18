package apiclient

import (
	"bytes"
	"encoding/gob"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

// cacheClient is an HTTPClient that caches responses
// on disk (or memory) using a KVStore. Note that this
// client will read the bodies of http.Requests and
// read the bodies of http.Responses. These bodies will
// be replaced by suitable bytes.Readers.
type cacheClient struct {
	HTTPClient
	KVStore
}

// cacheRecordKey is the key of a cacheRecord.
type cacheRecordKey struct {
	// Method is the request method.
	Method string

	// URLPath is the request URL path.
	URLPath string

	// Query is the URL query.
	Query string

	// Request is the request body.
	Request string
}

// requestMethod ensures that the request method is not empty.
func requestMethod(meth string) string {
	if meth == "" {
		meth = "GET"
	}
	return meth
}

// newCacheRecordKey creates a new cacheRecordKey from the
// provided http.Request. This function will read the request
// body and replace it with a string reader body.
func newCacheRecordKey(req *http.Request) (*cacheRecordKey, error) {
	out := &cacheRecordKey{
		Method:  requestMethod(req.Method),
		URLPath: req.URL.Path,
		Query:   req.URL.RawQuery,
	}
	if req.Body != nil {
		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		out.Request = string(data)
		req.Body = ioutil.NopCloser(bytes.NewReader(data))
	}
	return out, nil
}

// cacheRecord is the record stored in the cache.
type cacheRecord struct {
	// Key is the key in the cache.
	Key cacheRecordKey

	// Response is the response body.
	Response string
}

// cacheList is the list of cached records
type cacheList []cacheRecord

// Do implements HTTPClient.Do.
func (c *cacheClient) Do(req *http.Request) (*http.Response, error) {
	ckey, err := newCacheRecordKey(req)
	if err != nil {
		return nil, err
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return c.cachedResponse(ckey, err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close() // tell the real transport we're done here
	if errors.Is(err, io.EOF) && resp.Close {
		err = nil // treat EOF here like success
	}
	if err != nil {
		return c.cachedResponse(ckey, err)
	}
	defer c.addToCache(ckey, string(data))
	resp.Body = ioutil.NopCloser(bytes.NewReader(data))
	return resp, nil
}

const cacheKey = "apicache.state"

func (c *cacheClient) cachedResponse(ckey *cacheRecordKey, httpErr error) (*http.Response, error) {
	orig := c.readCache()
	for _, entry := range orig {
		if reflect.DeepEqual(*ckey, entry.Key) {
			resp := &http.Response{
				StatusCode: 200,
				Header:     http.Header{},
			}
			resp.Header.Set("X-APIClient-Cache", "HIT")
			resp.Body = ioutil.NopCloser(strings.NewReader(entry.Response))
			return resp, nil
		}
	}
	return nil, httpErr // return the original error
}

const maxCacheIndex = 128

func (c *cacheClient) addToCache(ckey *cacheRecordKey, body string) error {
	orig := c.readCache()
	var out cacheList
	out = append(out, cacheRecord{Key: *ckey, Response: body})
	for idx, entry := range orig {
		if !reflect.DeepEqual(*ckey, entry.Key) {
			out = append(out, entry)
		}
		if idx >= maxCacheIndex {
			// we wanna keep the cache size bounded
			break
		}
	}
	return c.writeCache(out)
}

// readCache always returns a valid cacheList. If we have already a cache, of
// course we return it. If we don't have a cache yet or there is any other error,
// we return an empty cache. We will discover about errors on write.
func (c *cacheClient) readCache() cacheList {
	data, err := c.KVStore.Get(cacheKey)
	if err != nil {
		return cacheList{}
	}
	var cl cacheList
	if err := gob.NewDecoder(bytes.NewReader(data)).Decode(&cl); err != nil {
		return cacheList{}
	}
	return cl
}

func (c *cacheClient) writeCache(cl cacheList) error {
	var bw bytes.Buffer
	if err := gob.NewEncoder(&bw).Encode(cl); err != nil {
		return err
	}
	return c.KVStore.Set(cacheKey, bw.Bytes())
}
