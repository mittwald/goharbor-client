package configure

import (
	"context"

	"github.com/go-openapi/runtime"

	v2client "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/configure"
	"github.com/mittwald/goharbor-client/v5/apiv2/model"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/config"
)

// RESTClient is a subclient for handling project related actions.
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
	GetConfigs(ctx context.Context) (*model.ConfigurationsResponse, error)
	GetInternalConfigs(ctx context.Context) (*model.InternalConfigurationsResponse, error)
	UpdateConfigs(ctx context.Context, newConfiguration *model.Configurations) error
}

// GetConfigs returns a system configurations object.
func (c *RESTClient) GetConfigs(ctx context.Context) (*model.ConfigurationsResponse, error) {
	params := &configure.GetConfigurationsParams{
		Context: ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.Configure.GetConfigurations(params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerConfigurationsErrors(err)
	}

	return resp.Payload, nil
}

// UpdateConfigs modifying system configurations that only provides for admin users.
func (c *RESTClient) UpdateConfigs(ctx context.Context, cfg *model.Configurations) error {
	params := &configure.UpdateConfigurationsParams{
		Configurations: cfg,
		Context:        ctx,
	}
	params.WithTimeout(c.Options.Timeout)
	_, err := c.V2Client.Configure.UpdateConfigurations(params, c.AuthInfo)
	return handleSwaggerConfigurationsErrors(err)
}
