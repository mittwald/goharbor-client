package systeminfo

import (
	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/systeminfo"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/errors"
)

// handleSwaggerSystemInfoErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with a proper message.
func handleSwaggerSystemInfoErrors(in error) error {
	switch in.(type) {
	case *systeminfo.GetSystemInfoInternalServerError:
		return &errors.ErrInternalErrors{}
	default:
		return in
	}
}
