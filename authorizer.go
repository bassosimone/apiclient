package apiclient

import "context"

// authorizer authenticates specific client requests.
type authorizer interface {
	// maybeRefreshToken refreshes the token for Authorization and returns
	// either such a token, on success, or the error that occurred.
	maybeRefreshToken(ctx context.Context) (string, error)
}
