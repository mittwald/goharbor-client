package robotv1

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/robotv1"
	"github.com/testwill/goharbor-client/v5/apiv2/pkg/errors"
)

// handleSwaggerRobotV1Errors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with a proper message.
func handleSwaggerRobotV1Errors(in error) error {
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
	case *robotv1.CreateRobotV1BadRequest:
		return &errors.ErrRobotAccountInvalid{}
	case *robotv1.UpdateRobotV1NotFound:
		return &errors.ErrRobotAccountUnknownResource{}
	case *robotv1.UpdateRobotV1Conflict:
		return &errors.ErrRobotAccountInvalid{}
	default:
		return in
	}
}
