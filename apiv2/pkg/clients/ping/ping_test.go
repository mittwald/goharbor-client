//go:build !integration

package ping

import (
	"context"
	"testing"

	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/ping"
	"github.com/mittwald/goharbor-client/v5/apiv2/mocks"
	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var ctx = context.Background()

func APIandMockClientsForTests() (*RESTClient, *clienttesting.MockClients) {
	desiredMockClients := &clienttesting.MockClients{
		Project: mocks.MockProjectClientService{},
	}

	v2Client := clienttesting.BuildV2ClientWithMocks(desiredMockClients)

	cl := NewClient(v2Client, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	return cl, desiredMockClients
}
func TestRESTClient_GetPing(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := &ping.GetPingParams{
		Context: ctx,
	}

	getParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Ping.On("GetPing", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(&ping.GetPingOK{}, nil)

	_, err := apiClient.GetPing(ctx)
	require.NoError(t, err)

	mockClient.Ping.AssertExpectations(t)
}
