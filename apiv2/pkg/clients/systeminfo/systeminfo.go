package systeminfo

import (
	"context"

	"github.com/go-openapi/runtime"

	v2client "github.com/testwill/goharbor-client/v5/apiv2/internal/api/client"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/systeminfo"
	modelv2 "github.com/testwill/goharbor-client/v5/apiv2/model"
	"github.com/testwill/goharbor-client/v5/apiv2/pkg/config"
)

// RESTClient is a subclient for handling system related actions.
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
	GetSystemInfo(ctx context.Context) (*modelv2.GeneralInfo, error)
}

// GetSystemInfo returns information.
func (c *RESTClient) GetSystemInfo(ctx context.Context) (*modelv2.GeneralInfo, error) {
	params := &systeminfo.GetSystemInfoParams{
		Context: ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.Systeminfo.GetSystemInfo(params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerSystemInfoErrors(err)
	}

	return resp.Payload, nil
}
