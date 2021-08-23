// +build !integration

package auditlog

import (
	"context"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	v2client "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/auditlog"
	"github.com/mittwald/goharbor-client/v4/apiv2/mocks"
)

var authInfo = runtimeclient.BasicAuth("foo", "bar")

func BuildV2ClientWithMocks(audit *mocks.MockAuditlogClientService) *v2client.Harbor {
	return &v2client.Harbor{
		Auditlog: audit,
	}
}

func TestRESTClient_ListAuditLogs(t *testing.T) {
	a := &mocks.MockAuditlogClientService{}

	v2Client := BuildV2ClientWithMocks(a)

	cl := NewClient(v2Client, authInfo)

	ctx := context.Background()

	listAuditLogsParams := &auditlog.ListAuditLogsParams{
		PageSize: nil,
		Q:        nil,
		Context:  ctx,
	}

	a.On("ListAuditLogs", listAuditLogsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&auditlog.ListAuditLogsOK{}, nil)

	_, err := cl.ListAuditLogs(ctx, nil, nil)

	assert.NoError(t, err)

	a.AssertExpectations(t)
}
