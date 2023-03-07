package statistic

import (
	"context"

	"github.com/go-openapi/runtime"
	v2client "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/statistic"
	"github.com/mittwald/goharbor-client/v5/apiv2/model"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/config"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/errors"
)

// RESTClient is a subclient for handling statistic related actions.
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
	GetStatistic(ctx context.Context) (*model.Statistic, error)
}

func (c *RESTClient) GetStatistic(ctx context.Context) (*model.Statistic, error) {
	params := &statistic.GetStatisticParams{
		Context: ctx,
	}
	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.Statistic.GetStatistic(params, c.AuthInfo)
	if err != nil {
		return nil, err
	}

	if resp.Payload == nil {
		return nil, &errors.ErrInternalErrors{}
	}

	return resp.Payload, nil
}
