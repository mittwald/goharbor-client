package gc

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/gc"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/errors"
)

// handleSystemErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with a proper message.
func handleSwaggerSystemErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		// As per documentation '200' should be the status code for success,
		// yet the API returns status code '201' when creating a GC schedule succeeds.
		case http.StatusCreated:
			return nil
		case http.StatusBadRequest:
			return &errors.ErrSystemInvalidSchedule{}
		case http.StatusUnauthorized:
			return &errors.ErrSystemUnauthorized{}
		case http.StatusForbidden:
			return &errors.ErrSystemNoPermission{}
		case http.StatusConflict:
			return &errors.ErrSystemGcInProgress{}
		case http.StatusInternalServerError:
			return &errors.ErrSystemInternalErrors{}
		}
	}

	switch in.(type) {
	case *gc.CreateGCScheduleConflict:
		return &errors.ErrSystemGcInProgress{}
	case *gc.UpdateGCScheduleBadRequest:
		return &errors.ErrSystemInvalidSchedule{}
	default:
		return in
	}
}
