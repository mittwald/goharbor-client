package auditlog

import (
	"context"

	"github.com/go-openapi/runtime"

	v2client "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/auditlog"
	"github.com/mittwald/goharbor-client/v4/apiv2/internal/legacyapi/client"
	"github.com/mittwald/goharbor-client/v4/apiv2/model"
)

// RESTClient is a subclient for handling user related actions.
type RESTClient struct {
	// The legacy swagger client
	LegacyClient *client.Harbor

	// The new client of the harbor v2 API
	V2Client *v2client.Harbor

	// AuthInfo contains the auth information that is provided on API calls.
	AuthInfo runtime.ClientAuthInfoWriter
}

func NewClient(v2Client *v2client.Harbor, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		V2Client: v2Client,
		AuthInfo: authInfo,
	}
}

type Client interface {
	ListAuditLogs(ctx context.Context, pageSize *int64, query *string) ([]*model.AuditLog, error)
}

func (c *RESTClient) ListAuditLogs(ctx context.Context, pageSize *int64, query *string) ([]*model.AuditLog, error) {
	params := auditlog.ListAuditLogsParams{
		Context:  ctx,
		PageSize: pageSize,
		Q:        query,
	}

	resp, err := c.V2Client.Auditlog.ListAuditLogs(&params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerAuditLogErrors(err)
	}

	return resp.Payload, nil
}
