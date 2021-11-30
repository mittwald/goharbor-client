package registry

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/registry"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/errors"
)

// handleRegistryErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with a proper message.
func handleSwaggerRegistryErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case http.StatusBadRequest:
			return &errors.ErrRegistryIllegalIDFormat{}
		case http.StatusUnauthorized:
			return &errors.ErrRegistryUnauthorized{}
		case http.StatusForbidden:
			return &errors.ErrRegistryNoPermission{}
		case http.StatusInternalServerError:
			return &errors.ErrRegistryInternalErrors{}
		}
	}

	switch in.(type) {
	case *registry.DeleteRegistryNotFound:
		return &errors.ErrRegistryIDNotExists{}
	case *registry.UpdateRegistryNotFound:
		return &errors.ErrRegistryIDNotExists{}
	case *registry.CreateRegistryConflict:
		return &errors.ErrRegistryNameAlreadyExists{}
	default:
		return in
	}
}
