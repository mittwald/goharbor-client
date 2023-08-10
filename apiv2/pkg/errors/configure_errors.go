package errors

const (
	// ErrConfigureUnauthorizedMsg is the error message for ErrConfigureUnauthorized error.
	ErrConfigureUnauthorizedMsg = "unauthorized"

	// ErrConfigureNoPermissionMsg is the error message for ErrConfigureNoPermission error.
	ErrConfigureNoPermissionMsg = "user does not have permission of admin role"

	// ErrConfigureInternalServerErrorMsg is the error message for ErrConfigureInternalServerError error.
	ErrConfigureInternalServerErrorMsg = "unexpected internal errors"
)

type (
	// ErrConfigureInternalServerError describes server-side internal errors.
	ErrConfigureInternalServerError struct{}

	// ErrConfigureNoPermission describes a request error without permission.
	ErrConfigureNoPermission struct{}

	// ErrConfigureUnauthorized describes an unauthorized request.
	ErrConfigureUnauthorized struct{}
)

// Error returns the error message.
func (e *ErrConfigureUnauthorized) Error() string {
	return ErrConfigureUnauthorizedMsg
}

// Error returns the error message.
func (e *ErrConfigureNoPermission) Error() string {
	return ErrConfigureNoPermissionMsg
}

// Error returns the error message.
func (e *ErrConfigureInternalServerError) Error() string {
	return ErrConfigureInternalServerErrorMsg
}
