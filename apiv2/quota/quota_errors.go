package quota

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/common"
)

// handleSwaggerQuotaErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with a proper message.
func handleSwaggerQuotaErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case http.StatusBadRequest:
			return &common.ErrQuotaIllegalIDFormat{}
		case http.StatusUnauthorized:
			return &common.ErrQuotaUnauthorized{}
		case http.StatusForbidden:
			return &common.ErrQuotaNoPermission{}
		case http.StatusNotFound:
			return &common.ErrQuotaUnknownResource{}
		case http.StatusInternalServerError:
			return &common.ErrQuotaInternalServerErrors{}
		}
	}
	return nil
}
