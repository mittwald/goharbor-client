package common

const (
	ErrRegistryIllegalIDFormatMsg   = "illegal format of provided ID value"
	ErrRegistryUnauthorizedMsg      = "unauthorized"
	ErrRegistryInternalErrorsMsg    = "unexpected internal errors"
	ErrRegistryNoPermissionMsg      = "user does not have permission to the registry"
	ErrRegistryIDNotExistsMsg       = "registry ID does not exist"
	ErrRegistryNameAlreadyExistsMsg = "registry name already exists"
	ErrRegistryMismatchMsg          = "id/name pair not found on server side"
	ErrRegistryNotFoundMsg          = "registry not found on server side"
	ErrRegistryNotProvidedMsg       = "no registry provided"
)

type (
	// ErrRegistryIllegalIDFormat describes an illegal request format.
	ErrRegistryIllegalIDFormat struct{}
	// ErrRegistryUnauthorized describes an unauthorized request.
	ErrRegistryUnauthorized struct{}
	// ErrRegistryInternalErrors describes server-side internal errors.
	ErrRegistryInternalErrors struct{}
	// ErrRegistryNoPermission describes a request error without permission.
	ErrRegistryNoPermission struct{}
	// ErrRegistryIDNotExists describes an error
	// when no proper registry ID is found.
	ErrRegistryIDNotExists struct{}
	// ErrRegistryNameAlreadyExists describes a duplicate registry name error.
	ErrRegistryNameAlreadyExists struct{}
	// ErrRegistryMismatch describes a failed lookup
	// of a registry with name/id pair.
	ErrRegistryMismatch struct{}
	// ErrRegistryNotFound describes an error
	// when a specific registry is not found.
	ErrRegistryNotFound struct{}
	// ErrRegistryNotProvided describes an error
	// when no registry was provided.
	ErrRegistryNotProvided struct{}
)

// Error returns the error message.
func (e *ErrRegistryIllegalIDFormat) Error() string {
	return ErrRegistryIllegalIDFormatMsg
}

// Error returns the error message.
func (e *ErrRegistryUnauthorized) Error() string {
	return ErrRegistryUnauthorizedMsg
}

// Error returns the error message.
func (e *ErrRegistryInternalErrors) Error() string {
	return ErrRegistryInternalErrorsMsg
}

// Error returns the error message.
func (e *ErrRegistryNoPermission) Error() string {
	return ErrRegistryNoPermissionMsg
}

// Error returns the error message.
func (e *ErrRegistryIDNotExists) Error() string {
	return ErrRegistryIDNotExistsMsg
}

// Error returns the error message.
func (e *ErrRegistryNameAlreadyExists) Error() string {
	return ErrRegistryNameAlreadyExistsMsg
}

// Error returns the error message.
func (e *ErrRegistryMismatch) Error() string {
	return ErrRegistryMismatchMsg
}

// Error returns the error message.
func (e *ErrRegistryNotFound) Error() string {
	return ErrRegistryNotFoundMsg
}

// Error returns the error message.
func (e *ErrRegistryNotProvided) Error() string {
	return ErrRegistryNotProvidedMsg
}
