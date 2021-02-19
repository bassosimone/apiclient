package apiclient

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/bassosimone/apiclient/internal/imodel"
	"github.com/bassosimone/apiclient/model"
	"github.com/google/go-cmp/cmp"
)

type handleRegisterAndLoginRecord struct {
	Method string
	Path   string
	Status int
}

type handleRegisterAndLogin struct {
	mu     sync.Mutex
	recs   []handleRegisterAndLoginRecord
	t      *testing.T
	tokens map[string]time.Time
	users  map[string]string
}

type handleRegisterAndLoginWriter struct {
	http.ResponseWriter
	parent *handleRegisterAndLogin
	method string
	path   string
}

func (w *handleRegisterAndLoginWriter) Write(b []byte) (int, error) {
	w.parent.t.Log("\t< 200 Ok")
	n, err := w.ResponseWriter.Write(b)
	defer w.parent.mu.Unlock()
	w.parent.mu.Lock()
	w.parent.recs = append(w.parent.recs, handleRegisterAndLoginRecord{
		Method: w.method,
		Path:   w.path,
		Status: 200,
	})
	return n, err
}

func (w *handleRegisterAndLoginWriter) WriteHeader(code int) {
	w.parent.t.Logf("\t< %d %s", code, http.StatusText(code))
	w.ResponseWriter.WriteHeader(code)
	defer w.parent.mu.Unlock()
	w.parent.mu.Lock()
	w.parent.recs = append(w.parent.recs, handleRegisterAndLoginRecord{
		Method: w.method,
		Path:   w.path,
		Status: code,
	})
}

func (h *handleRegisterAndLogin) dropAllAccounts() {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.users = make(map[string]string)
}

func (h *handleRegisterAndLogin) saveUser(name, pwd string) {
	defer h.mu.Unlock()
	h.mu.Lock()
	if h.users == nil {
		h.users = make(map[string]string)
	}
	h.users[name] = pwd
}

func (h *handleRegisterAndLogin) register(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	var rreq imodel.RegisterRequest
	if err := json.Unmarshal(data, &rreq); err != nil {
		w.WriteHeader(400)
		return
	}
	var username string
	ff := &fakeFill{}
	ff.fill(&username)
	rresp := imodel.RegisterResponse{ClientID: username}
	data, err = json.Marshal(rresp)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	h.saveUser(username, rreq.Password)
	w.Write(data)
}

func (h *handleRegisterAndLogin) checkPassword(name, pwd string) bool {
	defer h.mu.Unlock()
	h.mu.Lock()
	p := h.users[name]
	return p == pwd
}

func (h *handleRegisterAndLogin) login(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	var rreq imodel.LoginRequest
	if err := json.Unmarshal(data, &rreq); err != nil {
		w.WriteHeader(400)
		return
	}
	if !h.checkPassword(rreq.ClientID, rreq.Password) {
		w.WriteHeader(401)
		return
	}
	var token string
	ff := &fakeFill{}
	ff.fill(&token)
	t := time.Now().Add(3600 * time.Second)
	rresp := imodel.LoginResponse{Expire: t, Token: token}
	data, err = json.Marshal(rresp)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	h.saveToken(token, t)
	h.t.Log("\t< 200 Ok")
	w.Write(data)
}

func (h *handleRegisterAndLogin) saveToken(token string, expire time.Time) {
	defer h.mu.Unlock()
	h.mu.Lock()
	if h.tokens == nil {
		h.tokens = make(map[string]time.Time)
	}
	h.tokens[token] = expire
}

func (h *handleRegisterAndLogin) torTargets(w http.ResponseWriter, r *http.Request) {
	tok := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", 1)
	if !h.validToken(tok) {
		w.WriteHeader(401)
		return
	}
	var out model.TorTargetsResponse
	ff := &fakeFill{}
	ff.fill(&out)
	data, err := json.Marshal(out)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Write(data)
}

func (h *handleRegisterAndLogin) validToken(tok string) bool {
	defer h.mu.Unlock()
	h.mu.Lock()
	exp, found := h.tokens[tok]
	return found && exp.After(time.Now())
}

func (h *handleRegisterAndLogin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w = &handleRegisterAndLoginWriter{
		ResponseWriter: w,
		parent:         h,
		method:         r.Method,
		path:           r.URL.Path,
	}
	if r.Method == "POST" && r.URL.Path == "/api/v1/register" && r.Body != nil {
		h.t.Log("\t> POST /api/v1/register")
		h.register(w, r)
		return
	}
	if r.Method == "POST" && r.URL.Path == "/api/v1/login" && r.Body != nil {
		h.t.Log("\t> POST /api/v1/login")
		h.login(w, r)
		return
	}
	if r.Method == "GET" && r.URL.Path == "/api/v1/test-list/tor-targets" {
		h.t.Log("\t> GET /api/v1/test-list/tor-targets")
		h.torTargets(w, r)
		return
	}
	w.WriteHeader(404)
}

func (h *handleRegisterAndLogin) expireTokens() {
	defer h.mu.Unlock()
	h.mu.Lock()
	h.tokens = make(map[string]time.Time) // start afresh
}

func TestRegisterAndLoginWorkflow(t *testing.T) {
	handler := &handleRegisterAndLogin{t: t}
	srvr := httptest.NewServer(handler)
	defer srvr.Close()
	clnt := &Client{BaseURL: srvr.URL}
	ctx := context.Background()
	resp, err := clnt.TorTargets(ctx, &model.TorTargetsRequest{})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("expected non nil response here")
	}
	expect := []handleRegisterAndLoginRecord{{
		Method: "POST",
		Path:   "/api/v1/register",
		Status: 200,
	}, {
		Method: "POST",
		Path:   "/api/v1/login",
		Status: 200,
	}, {
		Method: "GET",
		Path:   "/api/v1/test-list/tor-targets",
		Status: 200,
	}}
	if diff := cmp.Diff(expect, handler.recs); diff != "" {
		t.Fatal(diff)
	}
}

func TestLoginWithExpiredToken(t *testing.T) {
	handler := &handleRegisterAndLogin{t: t}
	srvr := httptest.NewServer(handler)
	defer srvr.Close()
	kvstore := &memkvstore{}
	clnt := &Client{BaseURL: srvr.URL, KVStore: kvstore}
	ctx := context.Background()

	// step 1: create an account and login as a side effect
	resp, err := clnt.TorTargets(ctx, &model.TorTargetsRequest{})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("expected non nil response here")
	}

	// step 2: expire all tokens
	handler.expireTokens()

	// step 3: call API requiring token again, so we re-login
	resp, err = clnt.TorTargets(ctx, &model.TorTargetsRequest{})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("expected non nil response here")
	}

	// step 4: final checks
	expect := []handleRegisterAndLoginRecord{{
		Method: "POST",
		Path:   "/api/v1/register",
		Status: 200,
	}, {
		Method: "POST",
		Path:   "/api/v1/login",
		Status: 200,
	}, {
		Method: "GET",
		Path:   "/api/v1/test-list/tor-targets",
		Status: 200,
	}, {
		Method: "GET",
		Path:   "/api/v1/test-list/tor-targets",
		Status: 401,
	}, {
		Method: "POST",
		Path:   "/api/v1/login",
		Status: 200,
	}, {
		Method: "GET",
		Path:   "/api/v1/test-list/tor-targets",
		Status: 200,
	}}
	if diff := cmp.Diff(expect, handler.recs); diff != "" {
		t.Fatal(diff)
	}
}

func TestLoginWithDroppedDB(t *testing.T) {
	handler := &handleRegisterAndLogin{t: t}
	srvr := httptest.NewServer(handler)
	defer srvr.Close()
	kvstore := &memkvstore{}
	clnt := &Client{BaseURL: srvr.URL, KVStore: kvstore}
	ctx := context.Background()

	// step 1: create an account and login as a side effect
	resp, err := clnt.TorTargets(ctx, &model.TorTargetsRequest{})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("expected non nil response here")
	}

	// step 2: drop all accounts and tokens
	handler.expireTokens()
	handler.dropAllAccounts()

	// step 3: call API requiring token again, so we need
	// to create a new account and then to relogin.
	resp, err = clnt.TorTargets(ctx, &model.TorTargetsRequest{})
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("expected non nil response here")
	}

	// step 4: final checks
	expect := []handleRegisterAndLoginRecord{{
		Method: "POST",
		Path:   "/api/v1/register",
		Status: 200,
	}, {
		Method: "POST",
		Path:   "/api/v1/login",
		Status: 200,
	}, {
		Method: "GET",
		Path:   "/api/v1/test-list/tor-targets",
		Status: 200,
	}, {
		Method: "GET",
		Path:   "/api/v1/test-list/tor-targets",
		Status: 401,
	}, {
		Method: "POST",
		Path:   "/api/v1/login",
		Status: 401,
	}, {
		Method: "POST",
		Path:   "/api/v1/register",
		Status: 200,
	}, {
		Method: "POST",
		Path:   "/api/v1/login",
		Status: 200,
	}, {
		Method: "GET",
		Path:   "/api/v1/test-list/tor-targets",
		Status: 200,
	}}
	if diff := cmp.Diff(expect, handler.recs); diff != "" {
		t.Fatal(diff)
	}
}
