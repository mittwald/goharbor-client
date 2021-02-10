package quota

import (
	"context"

	"github.com/go-openapi/runtime"
	v2client "github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/legacyapi/client"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/legacyapi/client/products"
	legacymodel "github.com/mittwald/goharbor-client/v3/apiv2/model/legacy"
)

// RESTClient is a subclient for handling project related actions.
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
	GetQuotaByProjectID(ctx context.Context, projectID int64) (*legacymodel.Quota, error)
	UpdateStorageQuotaByProjectID(ctx context.Context, projectID int64, storageLimit int64) error
}

// GetQuotaByProjectID returns a quota object containing all configured quotas for a project.
func (c *RESTClient) GetQuotaByProjectID(ctx context.Context, projectID int64) (*legacymodel.Quota, error) {
	quota, err := c.LegacyClient.Products.GetQuotasID(&products.GetQuotasIDParams{
		ID:      projectID,
		Context: ctx,
	}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerQuotaErrors(err)
	}

	return quota.Payload, nil
}

// UpdateStorageQuotaByProjectID updates the storageLimit quota of a project.
func (c *RESTClient) UpdateStorageQuotaByProjectID(ctx context.Context, projectID int64, storageLimit int64) error {
	params := &products.PutQuotasIDParams{
		Hard: &legacymodel.QuotaUpdateReq{
			Hard: map[string]int64{
				"storage": storageLimit,
			},
		},
		ID:      projectID,
		Context: ctx,
	}

	_, err := c.LegacyClient.Products.PutQuotasID(params, c.AuthInfo)
	if err != nil {
		return handleSwaggerQuotaErrors(err)
	}

	return nil
}
