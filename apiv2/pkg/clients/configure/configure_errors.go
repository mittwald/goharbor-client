package configure

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/errors"
)

// handleSwaggerConfigurationsErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with a proper message.
func handleSwaggerConfigurationsErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case http.StatusUnauthorized:
			return &errors.ErrConfigureUnauthorized{}
		case http.StatusForbidden:
			return &errors.ErrConfigureNoPermission{}
		case http.StatusInternalServerError:
			return &errors.ErrConfigureInternalServerError{}
		}
	}
	return nil
}
