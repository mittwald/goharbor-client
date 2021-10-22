package robotv1

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/robotv1"
	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/common"
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
	case *robotv1.CreateRobotV1BadRequest:
		return &common.ErrRobotAccountInvalid{}
	case *robotv1.UpdateRobotV1NotFound:
		return &common.ErrRobotAccountUnknownResource{}
	case *robotv1.UpdateRobotV1Conflict:
		return &common.ErrRobotAccountInvalid{}
	default:
		return in
	}

}
