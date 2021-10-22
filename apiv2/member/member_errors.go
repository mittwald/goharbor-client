package member

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/member"
	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/common"
)

// handleSwaggerProjectErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with a proper message.
func handleSwaggerMemberErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case http.StatusCreated:
			// Harbor sometimes return 201 instead of 200 despite the swagger spec
			// not declaring it.
			return nil
		case http.StatusBadRequest:
			return &common.ErrBadRequest{}
		case http.StatusUnauthorized:
			return &common.ErrUnauthorized{}
		case http.StatusForbidden:
			return &common.ErrForbidden{}
		case http.StatusNotFound:
			return &common.ErrNotFound{}
		}
	}

	switch in.(type) {
	default:
		return in
	case *member.CreateProjectMemberConflict:
		return &common.ErrMemberAlreadyExists{}
	}
}
