package common

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
