package apiclient

import (
	"context"
	"errors"
)

func (c *Client) maybeLogin(ctx context.Context) (string, error) {
	return "", errors.New("apiclient: unauthorized")
}
