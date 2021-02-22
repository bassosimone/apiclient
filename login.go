package apiclient

import (
	"context"
	"crypto/rand"
	"encoding/base64"
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
	jsonCodec JSONCodec
	kvstore   KVStore
	state     loginState
}

// loginKey is the key with which loginState is saved
// into the key-value store used by Client.
const loginKey = "orchestra.state"

// newLoginManager always returns a valid loginManager structure
// that may contain content from the kvstore. If there's no content
// in the kvstore, or the content is corrupt, then we return an
// empty loginState data structure.
func (c *Client) newLoginManager() *loginManager {
	out := &loginManager{jsonCodec: c.jsonCodec, kvstore: c.kvStore}
	data, err := c.kvStore.Get(loginKey)
	if err != nil {
		return out
	}
	var ls loginState
	if err := c.jsonCodec.Decode(data, &ls); err != nil {
		return out
	}
	out.state = ls
	return out
}

// This list contains the errors returned by login code. Users should not
// see these errors until something's very wrong with the backend.
var (
	errWantLogin    = errors.New("apiclient: we need to login")
	errWantRegister = errors.New("apiclient: we need to register")
)

// TODO(bassosimone): it may be useful to hold a file-based mutex
// during the register and login process to protect the kvstore. This
// should probably be implemented into the kvstore itself.

func (lm *loginManager) writeback() error {
	data, err := lm.jsonCodec.Encode(lm.state)
	if err != nil {
		panic("json.Marshal should not fail here")
	}
	return lm.kvstore.Set(loginKey, data)
}

// newRandomPassword generates a new random password.
func (c *Client) newRandomPassword() string {
	const siz = 48
	b := make([]byte, siz)
	if _, err := rand.Read(b); err != nil {
		panic("rand.Read should not fail here")
	}
	return base64.StdEncoding.EncodeToString(b)
}

// newRegisterRequest creates a new RegisterRequest.
func (c *Client) newRegisterRequest() (*imodel.RegisterRequest, error) {
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
		Password: c.newRandomPassword(),
	}, nil
}

// loginAdapter adapts an API type to the login flow.
type loginAdapter interface {
	call(ctx context.Context, clnt *Client, token string) error
}

// doWithToken calls the specified API with the current token, if possible, and
// returns the error that occurred, if any.
func (c *Client) doWithToken(ctx context.Context, la loginAdapter) error {
	lm := c.newLoginManager()
	if lm.state.Token == "" {
		return errWantRegister // we never registered
	}
	if time.Now().Add(30 * time.Second).After(lm.state.Expire) {
		return errWantLogin // token has expired
	}
	switch err := la.call(ctx, c, lm.state.Token); err {
	case ErrUnauthorized:
		return errWantLogin // let us try with a relogin first
	case nil:
		return nil // api call successful
	default:
		return err // any other unrecoverable error
	}
}

// doLogin executes a login with the backend, if possible, and
// returns the result of doing this operation.
func (c *Client) doLogin(ctx context.Context) error {
	lm := c.newLoginManager()
	if lm.state.ClientID == "" || lm.state.Password == "" {
		return errWantRegister // we never registered
	}
	req := &imodel.LoginRequest{
		ClientID: lm.state.ClientID, Password: lm.state.Password}
	resp, err := newLoginAPI(c).call(ctx, req)
	switch err {
	case nil:
		lm.state.Token, lm.state.Expire = resp.Token, resp.Expire
		return lm.writeback() // this sounds like success
	case ErrUnauthorized:
		return errWantRegister // something changed in the server DB?
	default:
		return err // any other unrecoverable error
	}
}

// doRegister registers a new account with the backend and
// returns the result of attempting to do so.
func (c *Client) doRegister(ctx context.Context) error {
	req, err := c.newRegisterRequest()
	if err != nil {
		return err // unrecoverable error
	}
	resp, err := newRegisterAPI(c).call(ctx, req)
	if err != nil {
		return err // unrecoverable error
	}
	lm := c.newLoginManager()
	// start afresh with the saved state
	lm.state = loginState{
		ClientID: resp.ClientID,
		Password: req.Password,
	}
	return lm.writeback()
}

// doWithLoginAdapter attempts to call the logged-in API represented by the
// loginAdapter using the current login state. Depending on what happens this
// code may register a new account and call again the API.
func (c *Client) doWithLoginAdapter(ctx context.Context, la loginAdapter) error {
	switch err := c.doWithToken(ctx, la); err {
	case errWantRegister:
		if err := c.doRegister(ctx); err != nil {
			return err // unrecoverable error
		}
		if err := c.doLogin(ctx); err != nil {
			return err // unrecoverable error
		}
		return c.doWithToken(ctx, la) // token should be good
	case errWantLogin:
		switch err := c.doLogin(ctx); err {
		case errWantRegister:
			if err := c.doRegister(ctx); err != nil {
				return err // unrecoverable error
			}
			if err := c.doLogin(ctx); err != nil {
				return err // unrecoverable error
			}
			return c.doWithToken(ctx, la) // token should be good
		case nil:
			return c.doWithToken(ctx, la) // token should be good
		default:
			return err // unrecoverable error
		}
	case nil:
		return nil // we're all good
	default:
		return err // unrecoverable error
	}
}
