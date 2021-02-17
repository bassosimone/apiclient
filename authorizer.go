package apiclient

import "context"

// authorizer authenticates specific client requests.
type authorizer interface {
	// MaybeRefreshToken refreshes the token for Authorization and returns
	// either such a token, on success, or the error that occurred.
	maybeRefreshToken(ctx context.Context) (string, error)
}

type staticAuthorizer struct {
	token string
}

func (sa *staticAuthorizer) maybeRefreshToken(ctx context.Context) (string, error) {
	return sa.token, nil
}

// newStaticAuthorizer creates a new Authorizer that always
// returns the specified token to the caller.
func newStaticAuthorizer(token string) authorizer {
	return &staticAuthorizer{token}
}
