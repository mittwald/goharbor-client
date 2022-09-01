//go:build !integration

package auditlog

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/mock"

	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/auditlog"
	"github.com/testwill/goharbor-client/v5/apiv2/mocks"
	clienttesting "github.com/testwill/goharbor-client/v5/apiv2/pkg/testing"
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

func TestRESTClient_ListAuditLogs(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	listAuditLogsParams := &auditlog.ListAuditLogsParams{
		PageSize: &apiClient.Options.PageSize,
		Q:        &apiClient.Options.Query,
		Sort:     &apiClient.Options.Sort,
		Context:  ctx,
	}

	mockClient.Auditlog.On("ListAuditLogs", listAuditLogsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&auditlog.ListAuditLogsOK{}, nil)

	_, err := apiClient.ListAuditLogs(ctx)

	require.NoError(t, err)

	mockClient.Auditlog.AssertExpectations(t)
}
