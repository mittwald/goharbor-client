//go:build !integration

package configure

import (
	"context"
	"testing"

	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/configure"
	"github.com/mittwald/goharbor-client/v5/apiv2/mocks"
	"github.com/mittwald/goharbor-client/v5/apiv2/model"
	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	authMode      = "oidc"
	OIDCName      = "example"
	exampleConfig = &model.Configurations{
		AuthMode: &authMode,
		OIDCName: &OIDCName,
	}
	ctx = context.Background()
)

func APIandMockClientsForTests() (*RESTClient, *clienttesting.MockClients) {
	desiredMockClients := &clienttesting.MockClients{
		Configure: mocks.MockConfigureClientService{},
	}

	v2Client := clienttesting.BuildV2ClientWithMocks(desiredMockClients)

	cl := NewClient(v2Client, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	return cl, desiredMockClients
}

func TestRESTClient_GetConfigurations(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	params := &configure.GetConfigurationsParams{
		Context: ctx,
	}
	params.WithTimeout(apiClient.Options.Timeout)

	mockClient.Configure.On("GetConfigurations", params, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(&configure.GetConfigurationsOK{}, nil)

	_, err := apiClient.GetConfigs(ctx)

	require.NoError(t, err)
	mockClient.Configure.AssertExpectations(t)
}

func TestRESTClient_UpdateConfigs(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	params := &configure.UpdateConfigurationsParams{
		Configurations: exampleConfig,
		Context:        ctx,
	}
	params.WithTimeout(apiClient.Options.Timeout)

	mockClient.Configure.On("UpdateConfigurations", params, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(&configure.UpdateConfigurationsOK{}, nil)

	err := apiClient.UpdateConfigs(ctx, exampleConfig)
	require.NoError(t, err)
	mockClient.Configure.AssertExpectations(t)
}
