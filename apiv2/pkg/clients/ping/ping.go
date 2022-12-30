package ping

import (
	"context"

	"github.com/go-openapi/runtime"
	v2client "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/ping"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/config"
)

// RESTClient is a subclient for handling ping related actions.
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
	GetPing(ctx context.Context) (string, error)
}

func (c *RESTClient) GetPing(ctx context.Context) (string, error) {
	params := &ping.GetPingParams{
		Context: ctx,
	}
	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.Ping.GetPing(params, c.AuthInfo)
	if err != nil {
		return "", err
	}

	return resp.Payload, nil
}
