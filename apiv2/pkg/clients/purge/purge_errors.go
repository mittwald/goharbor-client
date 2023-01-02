package purge

import (
	"github.com/go-openapi/runtime"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/errors"
	"net/http"
)

// handleSwaggerPurgeErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with a proper message.
func handleSwaggerPurgeErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case http.StatusCreated:
			// Harbor sometimes returns 201 instead of 200 despite the swagger spec
			// not declaring it.
			return nil
		case http.StatusBadRequest:
			return &errors.ErrBadRequest{}
		case http.StatusUnauthorized:
			return &errors.ErrUnauthorized{}
		case http.StatusForbidden:
			return &errors.ErrQuotaNoPermission{}
		case http.StatusNotFound:
			return &errors.ErrQuotaUnknownResource{}
		case http.StatusInternalServerError:
			return &errors.ErrProjectInternalErrors{}
		}
	}

	return in
}
