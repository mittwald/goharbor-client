package repository

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/testwill/goharbor-client/v5/apiv2/pkg/errors"
)

// handleSwaggerRepositoryErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with a proper message.
func handleSwaggerRepositoryErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case http.StatusOK:
			return nil
		case http.StatusBadRequest:
			return &errors.ErrBadRequest{}
		case http.StatusUnauthorized:
			return &errors.ErrUnauthorized{}
		case http.StatusForbidden:
			return &errors.ErrForbidden{}
		case http.StatusNotFound:
			return &errors.ErrNotFound{}
		case http.StatusInternalServerError:
			return &errors.ErrInternalErrors{}
		}
	}

	return in
}
