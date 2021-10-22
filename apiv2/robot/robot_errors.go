package robot

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/robot"
	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/common"
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
			return &common.ErrRobotAccountInvalid{}
		case http.StatusUnauthorized:
			return &common.ErrRobotAccountUnauthorized{}
		case http.StatusForbidden:
			return &common.ErrRobotAccountNoPermission{}
		case http.StatusNotFound:
			return &common.ErrRobotAccountUnknownResource{}
		case http.StatusInternalServerError:
			return &common.ErrRobotAccountInternalErrors{}
		}
	}

	switch in.(type) {
	case *robot.CreateRobotBadRequest:
		return &common.ErrRobotAccountInvalid{}
	case *robot.UpdateRobotNotFound:
		return &common.ErrRobotAccountUnknownResource{}
	case *robot.UpdateRobotConflict:
		return &common.ErrRobotAccountInvalid{}
	default:
		return in
	}
}
