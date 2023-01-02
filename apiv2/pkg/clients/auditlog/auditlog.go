package auditlog

import (
	"context"

	"github.com/go-openapi/runtime"

	v2client "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/auditlog"
	"github.com/mittwald/goharbor-client/v5/apiv2/model"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/config"
)

// RESTClient is a subclient for handling user related actions.
type RESTClient struct {
	// Options contains optional configuration when making API calls.
	Options *config.Options

	// The new client of the harbor v2 API
	V2Client *v2client.Harbor

	// AuthInfo contains the auth information that is provided on API calls.
	AuthInfo runtime.ClientAuthInfoWriter
}

func NewClient(v2Client *v2client.Harbor, opts *config.Options, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		Options:  opts,
		V2Client: v2Client,
		AuthInfo: authInfo,
	}
}

type Client interface {
	ListAuditLogs(ctx context.Context) ([]*model.AuditLog, error)
}

// ListAuditLogs lists the audit logs of all projects the current user is a member of.
func (c *RESTClient) ListAuditLogs(ctx context.Context) ([]*model.AuditLog, error) {
	var auditLogs []*model.AuditLog
	page := c.Options.Page

	params := auditlog.ListAuditLogsParams{
		Page:     &page,
		PageSize: &c.Options.PageSize,
		Q:        &c.Options.Query,
		Sort:     &c.Options.Sort,
		Context:  ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	for {
		resp, err := c.V2Client.Auditlog.ListAuditLogs(&params, c.AuthInfo)
		if err != nil {
			return nil, handleSwaggerAuditLogErrors(err)
		}

		if len(resp.Payload) == 0 {
			break
		}

		totalCount := resp.XTotalCount

		auditLogs = append(auditLogs, resp.Payload...)

		if int64(len(auditLogs)) >= totalCount {
			break
		}

		page++
	}

	return auditLogs, nil
}
