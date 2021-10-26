package auditlog

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/auditlog"
)

const (
	// ErrAuditLogBadRequestMsg is the error message for ErrAuditLogBadRequest error.
	ErrAuditLogBadRequestMsg = "unsatisfied with constraints of the auditlog request"

	// ErrAuditLogUnauthorizedMsg is the error message for ErrAuditLogUnauthorized error.
	ErrAuditLogUnauthorizedMsg = "unauthorized"

	// ErrAuditLogInternalServerErrorMsg is the error message for ErrAuditLogInternalServerError error.
	ErrAuditLogInternalServerErrorMsg = "unexpected internal errors"
)

// ErrAuditLogBadRequest describes an error when a request to the auditlog API is malformed.
type ErrAuditLogBadRequest struct{}

// Error returns the error message.
func (e *ErrAuditLogBadRequest) Error() string {
	return ErrAuditLogBadRequestMsg
}

// ErrAuditLogUnauthorized describes an unauthorized request.
type ErrAuditLogUnauthorized struct{}

// Error returns the error message.
func (e *ErrAuditLogUnauthorized) Error() string {
	return ErrAuditLogUnauthorizedMsg
}

// ErrAuditLogInternalServerError describes server-side internal errors.
type ErrAuditLogInternalServerError struct{}

// Error returns the error message.
func (e *ErrAuditLogInternalServerError) Error() string {
	return ErrAuditLogInternalServerErrorMsg
}

// handleProjectErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with a proper message.
func handleSwaggerAuditLogErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case http.StatusBadRequest:
			return &ErrAuditLogBadRequest{}
		case http.StatusUnauthorized:
			return &ErrAuditLogUnauthorized{}
		case http.StatusInternalServerError:
			return &ErrAuditLogInternalServerError{}
		}
	}

	switch in.(type) {
	case *auditlog.ListAuditLogsInternalServerError:
		return &ErrAuditLogInternalServerError{}
	case *auditlog.ListAuditLogsBadRequest:
		return &ErrAuditLogBadRequest{}
	case *auditlog.ListAuditLogsUnauthorized:
		return &ErrAuditLogUnauthorized{}
	default:
		return in
	}
}
