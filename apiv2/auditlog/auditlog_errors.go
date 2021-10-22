package auditlog

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/auditlog"
	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/common"
)

// handleProjectErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with a proper message.
func handleSwaggerAuditLogErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case http.StatusBadRequest:
			return &common.ErrAuditLogBadRequest{}
		case http.StatusUnauthorized:
			return &common.ErrAuditLogUnauthorized{}
		case http.StatusInternalServerError:
			return &common.ErrAuditLogInternalServerError{}
		}
	}

	switch in.(type) {
	case *auditlog.ListAuditLogsInternalServerError:
		return &common.ErrAuditLogInternalServerError{}
	case *auditlog.ListAuditLogsBadRequest:
		return &common.ErrAuditLogBadRequest{}
	case *auditlog.ListAuditLogsUnauthorized:
		return &common.ErrAuditLogUnauthorized{}
	default:
		return in
	}
}
