// +build !integration

package quota

import (
	"context"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	v2client "github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/legacyapi/client"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/legacyapi/client/products"
	"github.com/mittwald/goharbor-client/v3/apiv2/mocks"
	legacymodel "github.com/mittwald/goharbor-client/v3/apiv2/model/legacy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	authInfo                       = runtimeclient.BasicAuth("foo", "bar")
	testProjectID            int64 = 1
	testStorageLimitPositive int64 = 1
	testStorageLimitNegative int64 = -1
	testStorageLimitNull     int64 = 0
)

func BuildLegacyClientWithMock(service *mocks.MockProductsClientService) *client.Harbor {
	return &client.Harbor{
		Products: service,
	}
}

func BuildProjectClientWithMocks() *v2client.Harbor {
	return &v2client.Harbor{}
}

func TestRESTClient_GetQuotaByProjectID(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildProjectClientWithMocks()

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	getQuotasIDParams := &products.GetQuotasIDParams{
		ID:      testProjectID,
		Context: ctx,
	}

	p.On("GetQuotasID", getQuotasIDParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetQuotasIDOK{Payload: &legacymodel.Quota{}}, nil)

	_, err := cl.GetQuotaByProjectID(ctx, testProjectID)
	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_UpdateStorageQuotaByProjectID(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildProjectClientWithMocks()

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	t.Run("PositiveLimit", func(t *testing.T) {
		putQuotasIDParams := &products.PutQuotasIDParams{
			ID: testProjectID,
			Hard: &legacymodel.QuotaUpdateReq{
				Hard: map[string]int64{
					"storage": testStorageLimitPositive,
				},
			},
			Context: ctx,
		}

		p.On("PutQuotasID", putQuotasIDParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&products.PutQuotasIDOK{}, nil)

		err := cl.UpdateStorageQuotaByProjectID(ctx, testProjectID, testStorageLimitPositive)
		assert.NoError(t, err)

		p.AssertExpectations(t)
	})

	t.Run("NegativeLimit", func(t *testing.T) {
		putQuotasIDParams := &products.PutQuotasIDParams{
			ID: testProjectID,
			Hard: &legacymodel.QuotaUpdateReq{
				Hard: map[string]int64{
					"storage": testStorageLimitNegative,
				},
			},
			Context: ctx,
		}

		p.On("PutQuotasID", putQuotasIDParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&products.PutQuotasIDOK{}, nil)

		err := cl.UpdateStorageQuotaByProjectID(ctx, testProjectID, testStorageLimitNegative)
		assert.NoError(t, err)

		p.AssertExpectations(t)
	})

	t.Run("NullLimit", func(t *testing.T) {
		putQuotasIDParams := &products.PutQuotasIDParams{
			ID: testProjectID,
			Hard: &legacymodel.QuotaUpdateReq{
				Hard: map[string]int64{
					"storage": testStorageLimitNegative,
				},
			},
			Context: ctx,
		}

		p.On("PutQuotasID", putQuotasIDParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&products.PutQuotasIDOK{}, nil)

		err := cl.UpdateStorageQuotaByProjectID(ctx, testProjectID, testStorageLimitNull)
		assert.NoError(t, err)

		p.AssertExpectations(t)
	})
}
