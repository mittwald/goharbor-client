package health

import (
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/health"
	"github.com/testwill/goharbor-client/v5/apiv2/pkg/errors"
)

// handleSwaggerHealthErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with a proper message.
func handleSwaggerHealthErrors(in error) error {
	switch in.(type) {
	case *health.GetHealthInternalServerError:
		return &errors.ErrInternalErrors{}
	default:
		return in
	}
}
