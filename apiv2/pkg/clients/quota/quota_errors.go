package quota

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/testwill/goharbor-client/v5/apiv2/pkg/errors"
)

// handleSwaggerQuotaErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with a proper message.
func handleSwaggerQuotaErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case http.StatusBadRequest:
			return &errors.ErrQuotaIllegalIDFormat{}
		case http.StatusUnauthorized:
			return &errors.ErrQuotaUnauthorized{}
		case http.StatusForbidden:
			return &errors.ErrQuotaNoPermission{}
		case http.StatusNotFound:
			return &errors.ErrQuotaUnknownResource{}
		case http.StatusInternalServerError:
			return &errors.ErrQuotaInternalServerErrors{}
		}
	}
	return nil
}
