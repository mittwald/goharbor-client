package auditlog

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/auditlog"
	"github.com/testwill/goharbor-client/v5/apiv2/pkg/errors"
)

// handleSwaggerAuditLogErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with a proper message.
func handleSwaggerAuditLogErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case http.StatusBadRequest:
			return &errors.ErrAuditLogBadRequest{}
		case http.StatusUnauthorized:
			return &errors.ErrAuditLogUnauthorized{}
		case http.StatusInternalServerError:
			return &errors.ErrAuditLogInternalServerError{}
		}
	}

	switch in.(type) {
	case *auditlog.ListAuditLogsInternalServerError:
		return &errors.ErrAuditLogInternalServerError{}
	case *auditlog.ListAuditLogsBadRequest:
		return &errors.ErrAuditLogBadRequest{}
	case *auditlog.ListAuditLogsUnauthorized:
		return &errors.ErrAuditLogUnauthorized{}
	default:
		return in
	}
}
