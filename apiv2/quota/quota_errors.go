package quota

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

const (
	// ErrQuotaIllegalIDFormatMsg is the error for message for ErrQuotaIllegalIDFormat errors.
	ErrQuotaIllegalIDFormatMsg = "illegal format in quota update request"

	// ErrQuotaUnauthorizedMsg is the error for message for ErrQuotaUnauthorized errors.
	ErrQuotaUnauthorizedMsg = "unauthorized"

	// ErrQuotaInternalServerErrorsMsg is the error message for ErrQuotaInternalServerErrors errors.
	ErrQuotaInternalServerErrorsMsg = "unexpected internal errors"

	// ErrQuotaNoPermissionMsg is the error message for ErrQuotaNoPermission errors.
	ErrQuotaNoPermissionMsg = "user does not have permission to the quota"

	// ErrQuotaUnknownResourceMsg is the errors message for ErrQuotaUnknownResource errors.
	ErrQuotaUnknownResourceMsg = "quota does not exist"

	ErrQuotaRefNotFoundMsg = "quota could not be found or contains unexpected reference object"
)

// ErrQuotaIllegalIDFormat describes an error due to an illegal request format.
type ErrQuotaIllegalIDFormat struct{}

// Error returns the error message.
func (e *ErrQuotaIllegalIDFormat) Error() string {
	return ErrQuotaIllegalIDFormatMsg
}

// ErrQuotaUnauthorized describes an unauthorized request.
type ErrQuotaUnauthorized struct{}

// Error returns the error message.
func (e *ErrQuotaUnauthorized) Error() string {
	return ErrQuotaUnauthorizedMsg
}

// ErrQuotaNoPermission describes an error in the request due to the lack of permissions.
type ErrQuotaNoPermission struct{}

// Error returns the error message.
func (e *ErrQuotaNoPermission) Error() string {
	return ErrQuotaNoPermissionMsg
}

// ErrQuotaUnknownResource describes an error when the specified quota could not be found.
type ErrQuotaUnknownResource struct{}

// Error returns the error message.
func (e *ErrQuotaUnknownResource) Error() string {
	return ErrQuotaUnknownResourceMsg
}

// ErrQuotaInternalServerErrors describes miscellaneous internal server errors.
type ErrQuotaInternalServerErrors struct{}

// Error returns the error message.
func (e *ErrQuotaInternalServerErrors) Error() string {
	return ErrQuotaInternalServerErrorsMsg
}

// ErrQuotaRefNotFound describes an error when the quota reference could not be found.
type ErrQuotaRefNotFound struct{}

func (e *ErrQuotaRefNotFound) Error() string {
	return ErrQuotaRefNotFoundMsg
}

// handleSwaggerQuotaErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with proper message.
func handleSwaggerQuotaErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case http.StatusBadRequest:
			return &ErrQuotaIllegalIDFormat{}
		case http.StatusUnauthorized:
			return &ErrQuotaUnauthorized{}
		case http.StatusForbidden:
			return &ErrQuotaNoPermission{}
		case http.StatusNotFound:
			return &ErrQuotaUnknownResource{}
		case http.StatusInternalServerError:
			return &ErrQuotaInternalServerErrors{}
		}
	}
	return nil
}
