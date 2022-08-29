//go:build !integration

package quota

import (
	"context"
	"encoding/json"
	"strconv"
	"testing"

	"github.com/goharbor/harbor/src/pkg/quota/types"
	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/quota"
	"github.com/mittwald/goharbor-client/v5/apiv2/mocks"
	modelv2 "github.com/mittwald/goharbor-client/v5/apiv2/model"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/errors"
	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	testStorageLimitPositive int64 = 1
	testStorageLimitNegative int64 = -1
	testStorageLimitNull     int64 = 0
	exampleQuotaID           int64 = 1
	exampleProjectID         int64 = 1
	ctx                            = context.Background()
)

func APIandMockClientsForTests() (*RESTClient, *clienttesting.MockClients) {
	desiredMockClients := &clienttesting.MockClients{
		Quota: mocks.MockQuotaClientService{},
	}

	v2Client := clienttesting.BuildV2ClientWithMocks(desiredMockClients)

	cl := NewClient(v2Client, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	return cl, desiredMockClients
}

func TestRESTClient_GetQuotaByProjectID_Unexpected(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	refID := strconv.Itoa(int(exampleProjectID))
	listParams := &quota.ListQuotasParams{
		Page:        &apiClient.Options.Page,
		PageSize:    &apiClient.Options.PageSize,
		ReferenceID: &refID,
		Sort:        &apiClient.Options.Sort,
		Context:     ctx,
	}

	listParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Quota.On("ListQuotas", listParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&quota.ListQuotasOK{}, nil)

	_, err := apiClient.GetQuotaByProjectID(ctx, exampleProjectID)
	require.Error(t, err)
	require.ErrorIs(t, err, &errors.ErrQuotaRefNotFound{})

	mockClient.Quota.AssertExpectations(t)
}

func TestRESTClient_GetQuotaByProjectID(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	refID := strconv.Itoa(int(exampleProjectID))

	listParams := &quota.ListQuotasParams{
		Page:        &apiClient.Options.Page,
		PageSize:    &apiClient.Options.PageSize,
		ReferenceID: &refID,
		Sort:        &apiClient.Options.Sort,
		Context:     ctx,
	}

	listParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Quota.On("ListQuotas", listParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&quota.ListQuotasOK{
			Payload: []*modelv2.Quota{{
				Hard: modelv2.ResourceList{
					string(types.ResourceStorage): 10,
				},
				ID: exampleProjectID,
				Ref: modelv2.QuotaRefObject(map[string]interface{}{
					"id": json.Number(strconv.Itoa(1)),
				}),
			}},
		}, nil)

	q, err := apiClient.GetQuotaByProjectID(ctx, exampleProjectID)

	require.NoError(t, err)

	require.NotNil(t, q)
	require.Equal(t, int64(10), q.Hard["storage"])

	mockClient.Quota.AssertExpectations(t)
}

func TestRESTClient_UpdateStorageQuotaByProjectID(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	t.Run("PositiveLimit", func(t *testing.T) {
		updateParams := &quota.UpdateQuotaParams{
			Hard: &modelv2.QuotaUpdateReq{
				Hard: modelv2.ResourceList{
					string(types.ResourceStorage): testStorageLimitPositive,
				},
			},
			ID:      exampleQuotaID,
			Context: ctx,
		}

		updateParams.WithTimeout(apiClient.Options.Timeout)

		mockClient.Quota.On("UpdateQuota", updateParams,
			mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&quota.UpdateQuotaOK{}, nil)

		err := apiClient.UpdateStorageQuotaByProjectID(ctx, exampleProjectID, testStorageLimitPositive)

		require.NoError(t, err)

		mockClient.Quota.AssertExpectations(t)
	})

	t.Run("NegativeLimit", func(t *testing.T) {
		updateParams := &quota.UpdateQuotaParams{
			Hard: &modelv2.QuotaUpdateReq{
				Hard: modelv2.ResourceList{
					string(types.ResourceStorage): testStorageLimitNegative,
				},
			},
			ID:      exampleQuotaID,
			Context: ctx,
		}

		updateParams.WithTimeout(apiClient.Options.Timeout)

		mockClient.Quota.On("UpdateQuota", updateParams,
			mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&quota.UpdateQuotaOK{}, nil)

		err := apiClient.UpdateStorageQuotaByProjectID(ctx, exampleProjectID, testStorageLimitNegative)

		require.NoError(t, err)

		mockClient.Quota.AssertExpectations(t)
	})

	t.Run("NullLimit", func(t *testing.T) {
		updateParams := &quota.UpdateQuotaParams{
			Hard: &modelv2.QuotaUpdateReq{
				Hard: modelv2.ResourceList{
					string(types.ResourceStorage): testStorageLimitNegative,
				},
			},
			ID:      exampleQuotaID,
			Context: ctx,
		}

		updateParams.WithTimeout(apiClient.Options.Timeout)

		mockClient.Quota.On("UpdateQuota", updateParams,
			mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&quota.UpdateQuotaOK{}, nil)

		err := apiClient.UpdateStorageQuotaByProjectID(ctx, exampleProjectID, testStorageLimitNull)

		require.NoError(t, err)

		mockClient.Quota.AssertExpectations(t)
	})
}
