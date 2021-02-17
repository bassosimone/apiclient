package apiclient

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"time"
)

// loginState contains the login state. This structure is saved
// into the KVStore and tracks whether we need to create a new
// account and/or whether we need to refresh the token.
type loginState struct {
	ClientID string
	Expire   time.Time
	Password string
	Token    string
	kvstore  KVStore
}

// loginKey is the key with which loginState is saved
// into the key-value store used by Client.
const loginKey = "orchestra.state"

// newLoginState always returns a valid loginState data structure
// that may contain content from the kvstore. If there's no content
// in the kvstore, or the content is corrupt, then we return an
// empty loginState data structure.
func newLoginState(kvstore KVStore) *loginState {
	data, err := kvstore.Get(loginKey)
	if err != nil {
		return &loginState{kvstore: kvstore}
	}
	var ls loginState
	if err := json.Unmarshal(data, &ls); err != nil {
		return &loginState{kvstore: kvstore}
	}
	ls.kvstore = kvstore
	return &ls
}

// This list contains the errors returned by login code.
var (
	errLoginBackendChanged = errors.New("apiclient: login: backend changed")
	errLoginNotRegistered  = errors.New("apiclient: login: not registered")
	errLoginTokenEmpty     = errors.New("apiclient: login: token empty")
	errLoginTokenExpired   = errors.New("apiclient: login: token expired")
)

// token returns the loginState token, if valid, or an
// error if the token has expired or is not valid.
func (ls *loginState) token() (string, error) {
	if ls.Token == "" {
		return "", errLoginTokenEmpty
	}
	if time.Now().Add(30 * time.Second).After(ls.Expire) {
		return "", errLoginTokenExpired
	}
	return ls.Token, nil
}

// loginRequest returns a LoginRequest for the current loginState
// or an error pointer if we don't have enough information.
func (ls *loginState) loginRequest() (*LoginRequest, error) {
	if ls.ClientID == "" || ls.Password == "" {
		return nil, errLoginNotRegistered
	}
	return &LoginRequest{ClientID: ls.ClientID, Password: ls.Password}, nil
}

func (ls *loginState) writeback() error {
	data, err := json.Marshal(ls)
	if err != nil {
		return err
	}
	return ls.kvstore.Set(loginKey, data)
}

// doLogin executes the login flow.
func (c *Client) doLogin(ctx context.Context, state *loginState) (string, error) {
	req, err := state.loginRequest()
	if err != nil {
		return "", err
	}
	resp, err := newLoginAPI(c).Call(ctx, req)
	if err != nil {
		if errors.Is(err, ErrHTTPFailure) {
			// This happens if we get a 401 Unauthorized because for
			// some reason the backend database has changed.
			return "", errLoginBackendChanged
		}
		return "", err
	}
	state.Token, state.Expire = resp.Token, resp.Expire
	if err := state.writeback(); err != nil {
		return "", err
	}
	return state.Token, nil
}

// TODO(bassosimone): it may be useful to hold a file-based mutex
// during the register and login process to protect the kvstore. This
// should probably be implemented into the kvstore itself.

// newRandomPassword generates a new random password.
func (c *Client) newRandomPassword() (string, error) {
	const siz = 48
	b := make([]byte, siz)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

// newRegisterRequest creates a new RegisterRequest.
func (c *Client) newRegisterRequest() (*RegisterRequest, error) {
	password, err := c.newRandomPassword()
	if err != nil {
		return nil, err
	}
	return &RegisterRequest{
		Metadata: RegisterRequestMetadata{
			// The original implementation has as its only use case that we
			// were registering and logging in for sending an update regarding
			// the probe whereabouts. Yet here in probe-engine, the orchestra
			// is currently only used to fetch inputs. For this purpose, we don't
			// need to communicate any specific information. The code that will
			// perform an update used to be responsible of doing that. Now, we
			// are not using orchestra for this purpose anymore.
			Platform:        "miniooni",
			ProbeASN:        "AS0",
			ProbeCC:         "ZZ",
			SoftwareName:    "miniooni",
			SoftwareVersion: "0.1.0-dev",
			SupportedTests:  []string{"web_connectivity"},
		},
		Password: password,
	}, nil
}

// doRegisterAndLogin executes the register and login flows.
func (c *Client) doRegisterAndLogin(ctx context.Context, state *loginState) (string, error) {
	req, err := c.newRegisterRequest()
	if err != nil {
		return "", err
	}
	resp, err := newRegisterAPI(c).Call(ctx, req)
	if err != nil {
		return "", err
	}
	state.ClientID, state.Password = resp.ClientID, req.Password
	if err := state.writeback(); err != nil {
		return "", err
	}
	return c.doLogin(ctx, state)
}

// maybeLogin returns the authorization token on success and
// the error that occurred in case of failure.
func (c *Client) maybeLogin(ctx context.Context) (string, error) {
	state := newLoginState(c.kvstore())
	if token, err := state.token(); err == nil {
		return token, nil // we already have a good token to use
	}
	if token, err := c.doLogin(ctx, state); err == nil {
		return token, nil
	}
	return c.doRegisterAndLogin(ctx, state)
}
