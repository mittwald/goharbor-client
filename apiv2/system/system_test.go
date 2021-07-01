// +build !integration

package system

import (
	"context"
	"errors"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	legacyclient "github.com/mittwald/goharbor-client/v4/apiv2/internal/legacyapi/client"
	"github.com/mittwald/goharbor-client/v4/apiv2/internal/legacyapi/client/products"
	"github.com/mittwald/goharbor-client/v4/apiv2/mocks"
	legacymodel "github.com/mittwald/goharbor-client/v4/apiv2/model/legacy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_Health_Error(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)

	cl := NewClient(legacyClient, nil, authInfo)

	ctx := context.Background()

	healthParams := &products.GetHealthParams{
		Context: ctx,
	}

	p.On("GetHealth", healthParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetHealthOK{Payload: &legacymodel.OverallHealthStatus{}},
			errors.New("err"))

	_, err := cl.Health(ctx)

	if assert.Error(t, err) {
		assert.Equal(t, errors.New("err"), err)
	}

	p.AssertExpectations(t)
}
