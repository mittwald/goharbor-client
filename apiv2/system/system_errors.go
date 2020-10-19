package system

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/legacyapi/client/products"
)

const (
	// ErrSystemInvalidScheduleMsg describes an invalid schedule type request
	ErrSystemInvalidScheduleMsg = "invalid schedule type"

	// ErrSystemGcInProgressMsg describes that a gc progress is already running
	ErrSystemGcInProgressMsg = "the system is already running a gc job"

	// ErrSystemUnauthorizedMsg describes an unauthorized request
	ErrSystemUnauthorizedMsg = "unauthorized"

	// ErrSystemInternalErrorsMsg describes server-side internal errors
	ErrSystemInternalErrorsMsg = "unexpected internal errors"

	// ErrSystemNoPermissionMsg describes a request error without permission
	ErrSystemNoPermissionMsg = "user does not have permission to the System"

	// ErrSystemGcUndefinedMsg describes a server-side response returning an empty GC schedule
	ErrSystemGcUndefinedMsg = "no schedule defined"

	// ErrSystemGcScheduleIdenticalMsg describes equality between two GC schedules
	ErrSystemGcScheduleIdenticalMsg = "the provided schedule is identical to the existing schedule"

	// ErrSystemGcScheduleNotProvidedMsg describes the absence of a required schedule
	ErrSystemGcScheduleNotProvidedMsg = "no schedule provided"
)

// ErrSystemInvalidSchedule describes an invalid schedule type request.
type ErrSystemInvalidSchedule struct{}

// Error returns the error message.
func (e *ErrSystemInvalidSchedule) Error() string {
	return ErrSystemInvalidScheduleMsg
}

// ErrSystemGcInProgress describes that a gc progress is already running.
type ErrSystemGcInProgress struct{}

// Error returns the error message.
func (e *ErrSystemGcInProgress) Error() string {
	return ErrSystemGcInProgressMsg
}

// ErrSystemUnauthorized describes an unauthorized request.
type ErrSystemUnauthorized struct{}

// Error returns the error message.
func (e *ErrSystemUnauthorized) Error() string {
	return ErrSystemUnauthorizedMsg
}

// ErrSystemInternalErrors describes server-side internal errors.
type ErrSystemInternalErrors struct{}

// Error returns the error message.
func (e *ErrSystemInternalErrors) Error() string {
	return ErrSystemInternalErrorsMsg
}

// ErrSystemNoPermission describes a request error without permission.
type ErrSystemNoPermission struct{}

// Error returns the error message.
func (e *ErrSystemNoPermission) Error() string {
	return ErrSystemNoPermissionMsg
}

// ErrSystemGcUndefined describes a server-side response returning an empty GC schedule.
type ErrSystemGcUndefined struct{}

// Error returns the error message.
func (e *ErrSystemGcUndefined) Error() string {
	return ErrSystemGcUndefinedMsg
}

// ErrSystemGcScheduleIdentical describes equality between two GC schedules.
type ErrSystemGcScheduleIdentical struct{}

// Error returns the error message.
func (e *ErrSystemGcScheduleIdentical) Error() string {
	return ErrSystemGcScheduleIdenticalMsg
}

// ErrSystemGcScheduleNotProvided describes the absence of a required schedule.
type ErrSystemGcScheduleNotProvided struct{}

// Error returns the error message.
func (e *ErrSystemGcScheduleNotProvided) Error() string {
	return ErrSystemGcScheduleNotProvidedMsg
}

// handleSystemErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with proper message.
func handleSwaggerSystemErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		// As per documentation '200' should be the status code for success,
		// yet the API returns status code '201' when creating a GC schedule succeeds.
		case http.StatusCreated:
			return nil
		case http.StatusBadRequest:
			return &ErrSystemInvalidSchedule{}
		case http.StatusUnauthorized:
			return &ErrSystemUnauthorized{}
		case http.StatusForbidden:
			return &ErrSystemNoPermission{}
		case http.StatusConflict:
			return &ErrSystemGcInProgress{}
		case http.StatusInternalServerError:
			return &ErrSystemInternalErrors{}
		}
	}

	switch in.(type) {
	case *products.PostSystemGcScheduleConflict:
		return &ErrSystemGcInProgress{}
	case *products.PutSystemGcScheduleBadRequest:
		return &ErrSystemInvalidSchedule{}
	default:
		return in
	}
}
