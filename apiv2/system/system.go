package system

import (
	"context"

	"github.com/go-openapi/runtime"
	v2client "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v4/apiv2/internal/legacyapi/client"
	"github.com/mittwald/goharbor-client/v4/apiv2/internal/legacyapi/client/products"
	legacymodel "github.com/mittwald/goharbor-client/v4/apiv2/model/legacy"
)

// RESTClient is a subclient for handling system related actions.
type RESTClient struct {
	// The legacy swagger client
	LegacyClient *client.Harbor

	// The new client of the harbor v2 API
	V2Client *v2client.Harbor

	// AuthInfo contains the auth information that is provided on API calls.
	AuthInfo runtime.ClientAuthInfoWriter
}

func NewClient(legacyClient *client.Harbor, v2Client *v2client.Harbor, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		LegacyClient: legacyClient,
		V2Client:     v2Client,
		AuthInfo:     authInfo,
	}
}

type Client interface {
	Health(ctx context.Context) (*legacymodel.OverallHealthStatus, error)
}

// Health reports Harbor system health information.
func (c *RESTClient) Health(ctx context.Context) (*legacymodel.OverallHealthStatus, error) {
	resp, err := c.LegacyClient.Products.GetHealth(&products.GetHealthParams{
		Context: ctx,
	}, c.AuthInfo)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
