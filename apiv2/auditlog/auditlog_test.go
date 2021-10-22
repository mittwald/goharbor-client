//go:build !integration

package auditlog

import (
	"context"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/auditlog"
	"github.com/mittwald/goharbor-client/v4/apiv2/mocks"
	unittesting "github.com/mittwald/goharbor-client/v4/apiv2/pkg/testing"
)

var (
	authInfo = runtimeclient.BasicAuth("foo", "bar")
	ctx      = context.Background()
)

func APIandMockClientsForTests() (*RESTClient, *unittesting.MockClients) {
	desiredMockClients := &unittesting.MockClients{
		Project: mocks.MockProjectClientService{},
	}

	v2Client := unittesting.BuildV2ClientWithMocks(desiredMockClients)

	cl := NewClient(v2Client, &unittesting.DefaultOpts, authInfo)

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

	assert.NoError(t, err)

	mockClient.Auditlog.AssertExpectations(t)
}
