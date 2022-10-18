package configurations

import (
	"context"
	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/configure"
	"github.com/mittwald/goharbor-client/v5/apiv2/mocks"
	"github.com/mittwald/goharbor-client/v5/apiv2/model"
	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
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

func TestRESTClient_GetConfigurationsInfo(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := configure.NewGetConfigurationsParams()

	getParams.WithTimeout(apiClient.Options.Timeout)
	getParams.WithContext(ctx)

	mockClient.Configure.On("GetConfigurations", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&configure.GetConfigurationsOK{Payload: &model.ConfigurationsResponse{}}, nil)

	resp, err := apiClient.GetConfigurationsInfo(ctx)
	require.NoError(t, err)
	require.NotNil(t, resp)

	mockClient.Configure.AssertExpectations(t)
}

func TestRESTClient_UpdateConfigurationsInfo(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := &configure.GetConfigurationsParams{
		Context: ctx,
	}
	getParams.WithTimeout(apiClient.Options.Timeout)

	authMode := model.StringConfigItem{
		Editable: false,
		Value:    "oidc_auth",
	}
	data, _ := authMode.MarshalBinary()
	auth := string(data)
	updateParams := &configure.UpdateConfigurationsParams{
		Configurations: &model.Configurations{
			AuthMode: &auth,
		},
		Context: ctx,
	}
	updateParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Configure.On("GetConfigurations", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&configure.GetConfigurationsOK{Payload: &model.ConfigurationsResponse{}}, nil)

	mockClient.Configure.On("UpdateConfigurations",
		updateParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&configure.UpdateConfigurationsOK{}, nil)

	resp, err := apiClient.GetConfigurationsInfo(ctx)
	require.NoError(t, err)
	require.NotNil(t, resp)

	err = apiClient.UpdateConfigurationsInfo(ctx, &model.Configurations{AuthMode: &auth})
	require.NoError(t, err)

	mockClient.Configure.AssertExpectations(t)
}
