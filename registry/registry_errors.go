package registry

import (
	"github.com/go-openapi/runtime"
	"github.com/mittwald/goharbor-client/internal/api/v1_10_0/client/products"
)

const (
	// ErrRegistryIllegalIDFormat describes an illegal request format
	ErrRegistryIllegalIDFormatMsg = "illegal format of provided ID value"
	// ErrRegistryUnauthorized describes an unauthorized request
	ErrRegistryUnauthorizedMsg = "unauthorized"
	// ErrRegistryInternalErrors describes server-side internal errors
	ErrRegistryInternalErrorsMsg = "unexpected internal errors"
	// ErrRegistryNoPermission describes a request error without permission
	ErrRegistryNoPermissionMsg = "user does not have permission to the registry"
	// ErrRegistryIDNotExists describes an error
	// when no proper registry ID is found
	ErrRegistryIDNotExistsMsg = "registry ID does not exist"
	// ErrRegistryNameAlreadyExists describes a duplicate registry name error
	ErrRegistryNameAlreadyExistsMsg = "registry name already exists"
	// ErrRegistryMismatch describes a failed lookup
	// of a registry with name/id pair
	ErrRegistryMismatchMsg = "id/name pair not found on server side"
	// ErrRegistryNotFound describes an error
	// when a specific registry is not found
	ErrRegistryNotFoundMsg    = "registry not found on server side"
	ErrRegistryNotProvidedMsg = "no registry provided"

	Status400 int = 400
	Status401 int = 401
	Status403 int = 403
	Status404 int = 404
	Status500 int = 500
)

// ErrRegistryIllegalIDFormat describes an illegal request format.
type ErrRegistryIllegalIDFormat struct{}

// Error returns the error message.
func (e *ErrRegistryIllegalIDFormat) Error() string {
	return ErrRegistryIllegalIDFormatMsg
}

// ErrRegistryUnauthorized describes an unauthorized request.
type ErrRegistryUnauthorized struct{}

// Error returns the error message.
func (e *ErrRegistryUnauthorized) Error() string {
	return ErrRegistryUnauthorizedMsg
}

// ErrRegistryInternalErrors describes server-side internal errors.
type ErrRegistryInternalErrors struct{}

// Error returns the error message.
func (e *ErrRegistryInternalErrors) Error() string {
	return ErrRegistryInternalErrorsMsg
}

// ErrRegistryNoPermission describes a request error without permission.
type ErrRegistryNoPermission struct{}

// Error returns the error message.
func (e *ErrRegistryNoPermission) Error() string {
	return ErrRegistryNoPermissionMsg
}

// ErrRegistryIDNotExists describes an error
// when no proper registry ID is found.
type ErrRegistryIDNotExists struct{}

// Error returns the error message.
func (e *ErrRegistryIDNotExists) Error() string {
	return ErrRegistryIDNotExistsMsg
}

// ErrRegistryNameAlreadyExists describes a duplicate registry name error.
type ErrRegistryNameAlreadyExists struct{}

// Error returns the error message.
func (e *ErrRegistryNameAlreadyExists) Error() string {
	return ErrRegistryNameAlreadyExistsMsg
}

// ErrRegistryMismatch describes a failed lookup
// of a registry with name/id pair.
type ErrRegistryMismatch struct{}

// Error returns the error message.
func (e *ErrRegistryMismatch) Error() string {
	return ErrRegistryMismatchMsg
}

// ErrRegistryNotFound describes an error
// when a specific registry is not found.
type ErrRegistryNotFound struct{}

// Error returns the error message.
func (e *ErrRegistryNotFound) Error() string {
	return ErrRegistryNotFoundMsg
}

type ErrRegistryNotProvided struct{}

// Error returns the error message.
func (e *ErrRegistryNotProvided) Error() string {
	return ErrRegistryNotProvidedMsg
}

// handleRegistryErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with proper message.
func handleSwaggerRegistryErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case Status400:
			return &ErrRegistryIllegalIDFormat{}
		case Status401:
			return &ErrRegistryUnauthorized{}
		case Status403:
			return &ErrRegistryNoPermission{}
		case Status500:
			return &ErrRegistryInternalErrors{}
		}
	}

	switch in.(type) {
	case *products.DeleteRegistriesIDNotFound:
		return &ErrRegistryIDNotExists{}
	case *products.PutRegistriesIDNotFound:
		return &ErrRegistryIDNotExists{}
	case *products.PostRegistriesConflict:
		return &ErrRegistryNameAlreadyExists{}
	default:
		return in
	}
}
