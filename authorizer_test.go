package apiclient

import "context"

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

type failingAuthorizer struct{}

func (fa *failingAuthorizer) maybeRefreshToken(ctx context.Context) (string, error) {
	return "", errMocked
}
