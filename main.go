package openhab

import (
	"context"
	"net/http"
)

//go:generate /bin/sh generate.sh

// WithToken adds an openHAB API token to each request.
func WithToken(token string) ClientOption {
	return WithRequestEditorFn(func(_ context.Context, req *http.Request) error {
		req.Header.Add("Authorization", "Bearer "+token)
		return nil
	})
}
