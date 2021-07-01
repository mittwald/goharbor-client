package quota

import (
	"context"
	"encoding/json"
	"reflect"
	"strconv"

	"github.com/go-openapi/runtime"
	v2client "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v4/apiv2/internal/legacyapi/client"
	"github.com/mittwald/goharbor-client/v4/apiv2/internal/legacyapi/client/products"
	legacymodel "github.com/mittwald/goharbor-client/v4/apiv2/model/legacy"
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

func (c *RESTClient) ListQuotas(ctx context.Context, referenceType, referenceID, sort *string) ([]*legacymodel.Quota, error) {
	resp, err := c.LegacyClient.Products.GetQuotas(&products.GetQuotasParams{
		Reference:   referenceType,
		ReferenceID: referenceID,
		Sort:        sort,
		Context:     ctx,
	}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerQuotaErrors(err)
	}

	return resp.Payload, nil
}

// GetQuotaByProjectID returns a quota object containing all configured quotas for a project.
func (c *RESTClient) GetQuotaByProjectID(ctx context.Context, projectID int64) (*legacymodel.Quota, error) {
	projectIDStr := strconv.Itoa(int(projectID))
	quotas, err := c.ListQuotas(ctx, nil, &projectIDStr, nil)
	if err != nil {
		return nil, handleSwaggerQuotaErrors(err)
	}

	// Assert that quota.Ref implements a map[string]interface{} type, as it holds json data.
	for _, quota := range quotas {
		if values, ok := quota.Ref.(map[string]interface{}); ok {
			if reflect.DeepEqual(values["id"], json.Number(projectIDStr)) {
				return quota, nil
			}
		}
	}

	return nil, &ErrQuotaRefNotFound{}
}

// UpdateStorageQuotaByProjectID updates the storageLimit quota of a project.
// A storageLimit value smaller than '0' will implicitly be set to '-1', equalling the 'unlimited' setting.
func (c *RESTClient) UpdateStorageQuotaByProjectID(ctx context.Context, projectID int64, storageLimit int64) error {
	if storageLimit <= 0 {
		storageLimit = -1
	}

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
