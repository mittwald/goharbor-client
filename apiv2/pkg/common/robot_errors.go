package common

const (
	// ErrRobotAccountInvalidMsg is the error message for ErrRobotAccountInvalid error.
	ErrRobotAccountInvalidMsg = "the robot account is invalid"
	// ErrRobotAccountUnauthorizedMsg is the error message for ErrRobotAccountUnauthorized error.
	ErrRobotAccountUnauthorizedMsg = "unauthorized"
	// ErrRobotAccountNoPermissionMsg is the error message for ErrRobotAccountNoPermission error.
	ErrRobotAccountNoPermissionMsg = "user does not have permission to the robot account"
	// ErrRobotAccountUnknownResourceMsg is the error message for ErrRobotAccountUnknownResource error.
	ErrRobotAccountUnknownResourceMsg = "resource unknown"
	// ErrRobotAccountInternalErrorsMsg is the error message for ErrRobotAccountInternalErrors error.
	ErrRobotAccountInternalErrorsMsg = "internal server error"
)

type (
	// ErrRobotAccountInvalid describes an invalid robot account error.
	ErrRobotAccountInvalid struct{}
	// ErrRobotAccountUnauthorized describes an unauthorized request to the 'robots' API.
	ErrRobotAccountUnauthorized struct{}
	// ErrRobotAccountNoPermission describes a request error without permission.
	ErrRobotAccountNoPermission struct{}
	// ErrRobotAccountUnknownResource describes an error when
	// the specified robot account could not be found.
	ErrRobotAccountUnknownResource struct{}
	// ErrRobotAccountInternalErrors describes server-sided internal errors.
	ErrRobotAccountInternalErrors struct{}
)

// Error returns the error message.
func (e *ErrRobotAccountInvalid) Error() string {
	return ErrRobotAccountInvalidMsg
}

// Error returns the error message.
func (e *ErrRobotAccountUnauthorized) Error() string {
	return ErrRobotAccountUnauthorizedMsg
}

// Error returns the error message.
func (e *ErrRobotAccountNoPermission) Error() string {
	return ErrRobotAccountNoPermissionMsg
}

// Error returns the error message.
func (e *ErrRobotAccountUnknownResource) Error() string {
	return ErrRobotAccountUnknownResourceMsg
}

// Error returns the error message.
func (e *ErrRobotAccountInternalErrors) Error() string {
	return ErrRobotAccountInternalErrorsMsg
}
