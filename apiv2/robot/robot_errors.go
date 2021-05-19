package robot

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client/robot"
)

// ErrRobotAccountInvalid
type ErrRobotAccountInvalid struct{}

func (e *ErrRobotAccountInvalid) Error() string {
	return ErrRobotAccountInvalidMsg
}

type ErrRobotAccountUnauthorized struct{}

func (e *ErrRobotAccountUnauthorized) Error() string {
	return ErrRobotAccountUnauthorizedMsg
}

type ErrRobotAccountNoPermission struct{}

func (e *ErrRobotAccountNoPermission) Error() string {
	return ErrRobotAccountNoPermissionMsg
}

type ErrRobotAccountUnknownResource struct{}

func (e *ErrRobotAccountUnknownResource) Error() string {
	return ErrRobotAccountUnknownResourceMsg
}

type ErrRobotAccountInternalErrors struct{}

func (e *ErrRobotAccountInternalErrors) Error() string {
	return ErrRobotAccountInternalErrorsMsg
}

const (
	ErrRobotAccountInvalidMsg         = "the robot account is invalid"
	ErrRobotAccountUnauthorizedMsg    = "unauthorized"
	ErrRobotAccountNoPermissionMsg    = "user does not have permission to the robot account"
	ErrRobotAccountUnknownResourceMsg = "resource unknown"
	ErrRobotAccountInternalErrorsMsg  = "internal server error"
)

// handleSwagerRobotErrors takes a swagger generated error as input,
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
