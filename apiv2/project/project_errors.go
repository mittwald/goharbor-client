package project

import (
	"net/http"

	"github.com/go-openapi/runtime"

	projectapi "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/project"
	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/common"
)

// handleSwaggerProjectErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with a proper message.
func handleSwaggerProjectErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case http.StatusCreated:
			// Harbor sometimes return 201 instead of 200 despite the swagger spec
			// not declaring it.
			return nil
		case http.StatusBadRequest:
			return &common.ErrProjectIllegalIDFormat{}
		case http.StatusUnauthorized:
			return &common.ErrUnauthorized{}
		case http.StatusForbidden:
			return &common.ErrProjectNoPermission{}
		case http.StatusNotFound:
			return &common.ErrProjectUnknownResource{}
		case http.StatusInternalServerError:
			return &common.ErrProjectInternalErrors{}
		}
	}

	switch in.(type) {
	case *projectapi.DeleteProjectNotFound:
		return &common.ErrProjectIDNotExists{}
	case *projectapi.UpdateProjectNotFound:
		return &common.ErrProjectIDNotExists{}
	case *projectapi.CreateProjectConflict:
		return &common.ErrProjectNameAlreadyExists{}
	default:
		return in
	}
}
