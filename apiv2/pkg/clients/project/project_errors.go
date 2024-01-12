package project

import (
	"net/http"

	"github.com/go-openapi/runtime"

	projectapi "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/project"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/errors"
)

// handleSwaggerProjectErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with a proper message.
func handleSwaggerProjectErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case http.StatusCreated:
			// Harbor sometimes returns 201 instead of 200 despite the swagger spec
			// not declaring it.
			return nil
		case http.StatusBadRequest:
			return &errors.ErrProjectIllegalIDFormat{}
		case http.StatusUnauthorized:
			return &errors.ErrUnauthorized{}
		case http.StatusForbidden:
			return &errors.ErrProjectNoPermission{}
		case http.StatusNotFound:
			return &errors.ErrProjectNotFound{}
		case http.StatusInternalServerError:
			return &errors.ErrProjectInternalErrors{}
		}
	}

	switch in.(type) {
	case *projectapi.DeleteProjectNotFound:
		return &errors.ErrProjectIDNotExists{}
	case *projectapi.UpdateProjectNotFound:
		return &errors.ErrProjectIDNotExists{}
	case *projectapi.CreateProjectConflict:
		return &errors.ErrProjectNameAlreadyExists{}
	default:
		return in
	}
}
