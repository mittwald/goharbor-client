package oidc

import (
	"context"
	"github.com/go-openapi/runtime"
	v2client "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/oidc"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/config"
)

// RESTClient is a subclient for handling oidc related actions.
type RESTClient struct {
	// Options contains optional configuration when making API calls.
	Options *config.Options

	// The new client of the harbor v2 API
	V2Client *v2client.Harbor

	// AuthInfo contains the auth information that is provided on API calls.
	AuthInfo runtime.ClientAuthInfoWriter
}

func NewClient(v2Client *v2client.Harbor, opts *config.Options, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		Options:  opts,
		V2Client: v2Client,
		AuthInfo: authInfo,
	}
}

type Client interface {
	PingOIDC(ctx context.Context) error
}

func (c *RESTClient) PingOIDC(ctx context.Context, body oidc.PingOIDCBody) error {
	params := &oidc.PingOIDCParams{
		Endpoint: body,
		Context:  ctx,
	}

	_, err := c.V2Client.OIDC.PingOIDC(params, c.AuthInfo)
	if err != nil {
		return err
	}

	return nil
}
