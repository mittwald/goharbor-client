package scanall

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/errors"
)

// handleRegistryErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with a proper message.
func handleSwaggerScanallErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case http.StatusBadRequest:
			return &errors.ErrBadRequest{}
		case http.StatusUnauthorized:
			return &errors.ErrUnauthorized{}
		case http.StatusForbidden:
			return &errors.ErrForbidden{}
		case http.StatusInternalServerError:
			return &errors.ErrInternalErrors{}
		}
	}
	return in
}
