package registry

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/registry"
	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/common"
)

// handleRegistryErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with a proper message.
func handleSwaggerRegistryErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case http.StatusBadRequest:
			return &common.ErrRegistryIllegalIDFormat{}
		case http.StatusUnauthorized:
			return &common.ErrRegistryUnauthorized{}
		case http.StatusForbidden:
			return &common.ErrRegistryNoPermission{}
		case http.StatusInternalServerError:
			return &common.ErrRegistryInternalErrors{}
		}
	}

	switch in.(type) {
	case *registry.DeleteRegistryNotFound:
		return &common.ErrRegistryIDNotExists{}
	case *registry.UpdateRegistryNotFound:
		return &common.ErrRegistryIDNotExists{}
	case *registry.CreateRegistryConflict:
		return &common.ErrRegistryNameAlreadyExists{}
	default:
		return in
	}
}
