package statistic

import (
	"context"
	"testing"

	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/statistic"
	"github.com/mittwald/goharbor-client/v5/apiv2/mocks"
	"github.com/mittwald/goharbor-client/v5/apiv2/model"
	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	ctx context.Context
)

func APIandMockClientsForTests() (*RESTClient, *clienttesting.MockClients) {
	desiredMockClients := &clienttesting.MockClients{
		Statistic: mocks.MockStatisticClientService{},
	}

	v2Client := clienttesting.BuildV2ClientWithMocks(desiredMockClients)

	cl := NewClient(v2Client, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	return cl, desiredMockClients
}

func TestRESTClient_GetStatistic(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := &statistic.GetStatisticParams{
		Context: ctx,
	}
	getParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Statistic.On("GetStatistic", getParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&statistic.GetStatisticOK{Payload: &model.Statistic{}}, nil)

	_, err := apiClient.GetStatistic(ctx)
	require.NoError(t, err)

	mockClient.Statistic.AssertExpectations(t)
}
