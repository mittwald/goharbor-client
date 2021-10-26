package user

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/user"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/errors"
)

// handleUserErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with a proper message.
func handleSwaggerUserErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case http.StatusConflict:
			return &errors.ErrUserAlreadyExists{}
		case http.StatusCreated:
			return nil
		}
	}

	switch in.(type) {
	default:
		return in
	case *user.CreateUserBadRequest:
		return &errors.ErrUserBadRequest{}
	case *user.UpdateUserPasswordBadRequest:
		return &errors.ErrUserPasswordInvalid{}
	}
}
