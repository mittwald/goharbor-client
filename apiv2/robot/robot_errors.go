package robot

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/robot"
)

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

// ErrRobotAccountInvalid describes an invalid robot account error.
type ErrRobotAccountInvalid struct{}

// Error returns the error message.
func (e *ErrRobotAccountInvalid) Error() string {
	return ErrRobotAccountInvalidMsg
}

// ErrRobotAccountUnauthorized describes an unauthorized request to the 'robots' API.
type ErrRobotAccountUnauthorized struct{}

// Error returns the error message.
func (e *ErrRobotAccountUnauthorized) Error() string {
	return ErrRobotAccountUnauthorizedMsg
}

// ErrRobotAccountNoPermission describes a request error without permission.
type ErrRobotAccountNoPermission struct{}

// Error returns the error message.
func (e *ErrRobotAccountNoPermission) Error() string {
	return ErrRobotAccountNoPermissionMsg
}

// ErrRobotAccountUnknownResource describes an error when
// the specified robot account could not be found.
type ErrRobotAccountUnknownResource struct{}

// Error returns the error message.
func (e *ErrRobotAccountUnknownResource) Error() string {
	return ErrRobotAccountUnknownResourceMsg
}

// ErrRobotAccountInternalErrors describes server-sided internal errors.
type ErrRobotAccountInternalErrors struct{}

// Error returns the error message.
func (e *ErrRobotAccountInternalErrors) Error() string {
	return ErrRobotAccountInternalErrorsMsg
}

// handleSwaggerRobotErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with proper message.
func handleSwaggerRobotErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case http.StatusOK:
			return nil
		case http.StatusCreated:
			return nil
		case http.StatusBadRequest:
			return &ErrRobotAccountInvalid{}
		case http.StatusUnauthorized:
			return &ErrRobotAccountUnauthorized{}
		case http.StatusForbidden:
			return &ErrRobotAccountNoPermission{}
		case http.StatusNotFound:
			return &ErrRobotAccountUnknownResource{}
		case http.StatusInternalServerError:
			return &ErrRobotAccountInternalErrors{}
		}
	}

	switch in.(type) {
	case *robot.CreateRobotBadRequest:
		return &ErrRobotAccountInvalid{}
	case *robot.UpdateRobotNotFound:
		return &ErrRobotAccountUnknownResource{}
	case *robot.UpdateRobotConflict:
		return &ErrRobotAccountInvalid{}
	default:
		return in
	}
}
