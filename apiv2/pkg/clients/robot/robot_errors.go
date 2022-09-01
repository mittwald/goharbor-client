package robot

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/robot"
	"github.com/testwill/goharbor-client/v5/apiv2/pkg/errors"
)

// handleSwaggerRobotErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with a proper message.
func handleSwaggerRobotErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case http.StatusOK:
			return nil
		case http.StatusCreated:
			return nil
		case http.StatusBadRequest:
			return &errors.ErrRobotAccountInvalid{}
		case http.StatusUnauthorized:
			return &errors.ErrRobotAccountUnauthorized{}
		case http.StatusForbidden:
			return &errors.ErrRobotAccountNoPermission{}
		case http.StatusNotFound:
			return &errors.ErrRobotAccountUnknownResource{}
		case http.StatusInternalServerError:
			return &errors.ErrRobotAccountInternalErrors{}
		}
	}

	switch in.(type) {
	case *robot.CreateRobotBadRequest:
		return &errors.ErrRobotAccountInvalid{}
	case *robot.UpdateRobotNotFound:
		return &errors.ErrRobotAccountUnknownResource{}
	case *robot.UpdateRobotConflict:
		return &errors.ErrRobotAccountInvalid{}
	default:
		return in
	}
}
