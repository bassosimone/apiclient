package apiclient

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"time"

	"github.com/bassosimone/apiclient/internal/imodel"
)

// loginState contains the login state. This structure is saved
// into the KVStore and tracks whether we need to create a new
// account and/or whether we need to refresh the token.
type loginState struct {
	ClientID string
	Expire   time.Time
	Password string
	Token    string
}

type loginManager struct {
	kvstore KVStore
	state   loginState
}

// loginKey is the key with which loginState is saved
// into the key-value store used by Client.
const loginKey = "orchestra.state"

// newLoginManager always returns a valid loginManager structure
// that may contain content from the kvstore. If there's no content
// in the kvstore, or the content is corrupt, then we return an
// empty loginState data structure.
func newLoginManager(kvstore KVStore) *loginManager {
	data, err := kvstore.Get(loginKey)
	if err != nil {
		return &loginManager{kvstore: kvstore}
	}
	var ls loginState
	if err := json.Unmarshal(data, &ls); err != nil {
		return &loginManager{kvstore: kvstore}
	}
	return &loginManager{kvstore: kvstore, state: ls}
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
func (lm *loginManager) token() (string, error) {
	if lm.state.Token == "" {
		return "", errLoginTokenEmpty
	}
	if time.Now().Add(30 * time.Second).After(lm.state.Expire) {
		return "", errLoginTokenExpired
	}
	return lm.state.Token, nil
}

// loginRequest returns a LoginRequest for the current loginState
// or an error if we don't have enough information.
func (lm *loginManager) loginRequest() (*imodel.LoginRequest, error) {
	if lm.state.ClientID == "" || lm.state.Password == "" {
		return nil, errLoginNotRegistered
	}
	return &imodel.LoginRequest{
		ClientID: lm.state.ClientID,
		Password: lm.state.Password,
	}, nil
}

func (lm *loginManager) writeback() error {
	data, err := json.Marshal(lm.state)
	if err != nil {
		return err
	}
	return lm.kvstore.Set(loginKey, data)
}

// doLogin executes the login flow and returns the token or an error.
func (c *Client) doLogin(ctx context.Context, lm *loginManager) (string, error) {
	req, err := lm.loginRequest()
	if err != nil {
		return "", err
	}
	resp, err := newLoginAPI(c).call(ctx, req)
	if err != nil {
		if errors.Is(err, ErrHTTPFailure) {
			// This happens if we get a 401 Unauthorized because for
			// some reason the backend database has changed.
			// TODO(bassosimone): need to check for 401 explicitly?
			err = errLoginBackendChanged
		}
		return "", err
	}
	lm.state.Token, lm.state.Expire = resp.Token, resp.Expire
	if err := lm.writeback(); err != nil {
		return "", err
	}
	return lm.state.Token, nil
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
func (c *Client) newRegisterRequest() (*imodel.RegisterRequest, error) {
	password, err := c.newRandomPassword()
	if err != nil {
		return nil, err
	}
	return &imodel.RegisterRequest{
		Metadata: imodel.RegisterRequestMetadata{
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
func (c *Client) doRegisterAndLogin(ctx context.Context, lm *loginManager) (string, error) {
	req, err := c.newRegisterRequest()
	if err != nil {
		return "", err
	}
	resp, err := newRegisterAPI(c).call(ctx, req)
	if err != nil {
		return "", err
	}
	lm.state.ClientID, lm.state.Password = resp.ClientID, req.Password
	if err := lm.writeback(); err != nil {
		return "", err
	}
	return c.doLogin(ctx, lm)
}

// maybeRefreshToken implements authorizer.maybeRefreshToken.
//
// When invoked, this method will roughly do the following:
//
// 1. if we already have a valid token, just return it;
//
// 2. if we already have valid orchestra credentials, then
// login again so to refresh the token, then return the token;
//
// 3. otherwise, create a new account, and then login with
// such an account, so we have a token to return.
//
// This implementation should be robust to a change in
// the backend database where all logins are lost.
func (c *Client) maybeRefreshToken(ctx context.Context) (string, error) {
	lm := newLoginManager(c.kvstore())
	if token, err := lm.token(); err == nil {
		return token, nil // we already have a good token to use
	}
	if token, err := c.doLogin(ctx, lm); err == nil {
		return token, nil // we have relogged in
	}
	return c.doRegisterAndLogin(ctx, lm)
}
