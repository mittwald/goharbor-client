package auditlog

import (
	"context"

	"github.com/go-openapi/runtime"

	v2client "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/auditlog"
	"github.com/mittwald/goharbor-client/v5/apiv2/model"
)

// RESTClient is a subclient for handling user related actions.
type RESTClient struct {
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

// ListAuditLogs lists the audit logs of all projects the current user is a member of.
// The 'pageSize' specifies how many audit log entries will be listed, where ...
// a value > '0' will list the specified number of log entries,
// a value of <= '0' will list all audit log entries,
// a 'nil' value will list the ten most recent log entries.
// Specifying 'query' will return the audit logs matching the specified query, e.g. 'operation=create'.
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
