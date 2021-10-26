// +build !integration

package system

import (
	"context"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	legacyclient "github.com/mittwald/goharbor-client/v5/apiv2/internal/legacyapi/client"
	"github.com/mittwald/goharbor-client/v5/apiv2/internal/legacyapi/client/products"
	"github.com/mittwald/goharbor-client/v5/apiv2/mocks"
	legacymodel "github.com/mittwald/goharbor-client/v5/apiv2/model/legacy"
)

var authInfo = runtimeclient.BasicAuth("foo", "bar")

func BuildLegacyClientWithMock(service *mocks.MockProductsClientService) *legacyclient.Harbor {
	return &legacyclient.Harbor{
		Products: service,
	}
}

func TestRESTClient_Health(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)

	cl := NewClient(legacyClient, nil, authInfo)

	ctx := context.Background()

	healthParams := &products.GetHealthParams{
		Context: ctx,
	}

	p.On("GetHealth", healthParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetHealthOK{Payload: &legacymodel.OverallHealthStatus{}}, nil)

	_, err := cl.Health(ctx)

	require.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_GetSystemCVEAllowList(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)

	cl := NewClient(legacyClient, nil, authInfo)

	ctx := context.Background()

	getParams := &products.GetSystemCVEAllowlistParams{
		Context: ctx,
	}

	p.On("GetSystemCVEAllowlist", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetSystemCVEAllowlistOK{}, nil)

	_, err := cl.GetSystemCVEAllowList(ctx)

	require.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_UpdateSystemCVEAllowList(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)

	cl := NewClient(legacyClient, nil, authInfo)

	ctx := context.Background()

	putParams := &products.PutSystemCVEAllowlistParams{
		Allowlist: &legacymodel.CVEAllowlist{
			ExpiresAt: 1640995200,
			ID:        0,
			Items: []*legacymodel.CVEAllowlistItem{{
				CveID: "CVE-1999-0095",
			}},
			ProjectID: 0,
		},
		Context: ctx,
	}

	p.On("PutSystemCVEAllowlist", putParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PutSystemCVEAllowlistOK{}, nil)

	err := cl.UpdateSystemCVEAllowList(ctx, []string{"CVE-1999-0095"}, 1640995200)

	require.NoError(t, err)

	p.AssertExpectations(t)
}
