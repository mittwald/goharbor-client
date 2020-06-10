package goharborclient

import (
	"github.com/go-openapi/runtime"
	"github.com/mittwald/goharbor-client/api/v1.10.0/client/products"
)

const (
	// ErrReplicationIllegalIDFormat describes an illegal request format
	ErrReplicationIllegalIDFormatMsg = "illegal format of provided ID value"

	// ErrReplicationUnauthorized describes an unauthorized request
	ErrReplicationUnauthorizedMsg = "unauthorized"

	// ErrReplicationInternalErrors describes server-side internal errors
	ErrReplicationInternalErrorsMsg = "unexpected internal errors"

	// ErrReplicationNoPermission describes a request error without permission
	ErrReplicationNoPermissionMsg = "user does not have permission to the replication"

	// ErrReplicationIDNotExists describes an error
	// when no proper replication ID is found
	ErrReplicationIDNotExistsMsg = "replication ID does not exist"

	// ErrReplicationNameAlreadyExists describes a duplicate replication name error
	ErrReplicationNameAlreadyExistsMsg = "replication name already exists"

	// ErrReplicationMismatch describes a failed lookup
	// of a replication with name/id pair
	ErrReplicationMismatchMsg = "id/name pair not found on server side"

	// ErrReplicationNotFound describes an error
	// when a specific replication is not found
	ErrReplicationNotFoundMsg = "replication not found on server side"

	// ErrReplicationNotProvidedMsg describes an error
	// caused by a missing replication object
	ErrReplicationNotProvidedMsg = "no replication provided"
)

// ErrReplicationIllegalIDFormat describes an illegal request format.
type ErrReplicationIllegalIDFormat struct{}

// Error returns the error message.
func (e *ErrReplicationIllegalIDFormat) Error() string {
	return ErrReplicationIllegalIDFormatMsg
}

// ErrReplicationUnauthorized describes an unauthorized request.
type ErrReplicationUnauthorized struct{}

// Error returns the error message.
func (e *ErrReplicationUnauthorized) Error() string {
	return ErrReplicationUnauthorizedMsg
}

// ErrReplicationInternalErrors describes server-side internal errors.
type ErrReplicationInternalErrors struct{}

// Error returns the error message.
func (e *ErrReplicationInternalErrors) Error() string {
	return ErrReplicationInternalErrorsMsg
}

// ErrReplicationNoPermission describes a request error without permission.
type ErrReplicationNoPermission struct{}

// Error returns the error message.
func (e *ErrReplicationNoPermission) Error() string {
	return ErrReplicationNoPermissionMsg
}

// ErrReplicationIDNotExists describes an error
// when no proper replication ID is found.
type ErrReplicationIDNotExists struct{}

// Error returns the error message.
func (e *ErrReplicationIDNotExists) Error() string {
	return ErrReplicationIDNotExistsMsg
}

// ErrReplicationNameAlreadyExists describes a duplicate replication name error.
type ErrReplicationNameAlreadyExists struct{}

// Error returns the error message.
func (e *ErrReplicationNameAlreadyExists) Error() string {
	return ErrReplicationNameAlreadyExistsMsg
}

// ErrReplicationMismatch describes a failed lookup
// of a replication with name/id pair.
type ErrReplicationMismatch struct{}

// Error returns the error message.
func (e *ErrReplicationMismatch) Error() string {
	return ErrReplicationMismatchMsg
}

// ErrReplicationNotFound describes an error
// when a specific replication is not found.
type ErrReplicationNotFound struct{}

// Error returns the error message.
func (e *ErrReplicationNotFound) Error() string {
	return ErrReplicationNotFoundMsg
}

type ErrReplicationNotProvided struct{}

// Error returns the error message.
func (e *ErrReplicationNotProvided) Error() string {
	return ErrReplicationNotProvidedMsg
}

// handleReplicationErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with proper message.
func handleSwaggerReplicationErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case 400:
			return &ErrReplicationIllegalIDFormat{}
		case 401:
			return &ErrReplicationUnauthorized{}
		case 403:
			return &ErrReplicationNoPermission{}
		case 500:
			return &ErrReplicationInternalErrors{}
		}
	}

	switch in.(type) {
	case *products.DeleteRegistriesIDNotFound:
		return &ErrReplicationIDNotExists{}
	case *products.PutRegistriesIDNotFound:
		return &ErrReplicationIDNotExists{}
	case *products.PostRegistriesConflict:
		return &ErrReplicationNameAlreadyExists{}
	default:
		return in
	}
}
