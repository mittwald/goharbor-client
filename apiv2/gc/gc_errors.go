package gc

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/gc"
	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/common"
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
			return &common.ErrSystemInvalidSchedule{}
		case http.StatusUnauthorized:
			return &common.ErrSystemUnauthorized{}
		case http.StatusForbidden:
			return &common.ErrSystemNoPermission{}
		case http.StatusConflict:
			return &common.ErrSystemGcInProgress{}
		case http.StatusInternalServerError:
			return &common.ErrSystemInternalErrors{}
		}
	}

	switch in.(type) {
	case *gc.CreateGCScheduleConflict:
		return &common.ErrSystemGcInProgress{}
	case *gc.UpdateGCScheduleBadRequest:
		return &common.ErrSystemInvalidSchedule{}
	default:
		return in
	}
}
