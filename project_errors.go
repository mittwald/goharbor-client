package goharborclient

const (
	// ErrProjectIllegalIDFormatMsg is the error message for ErrProjectIllegalIDFormat error.
	ErrProjectIllegalIDFormatMsg = "illegal format of provided ID value"

	// ErrProjectUnauthorizedMsg is the error message for ErrProjectUnauthorized error.
	ErrProjectUnauthorizedMsg = "unauthorized"

	// ErrProjectInternalErrorsMsg is the error message for ErrProjectInternalErrors error.
	ErrProjectInternalErrorsMsg = "unexpected internal errors"

	// ErrProjectNoPermissionMsg is the error message for ErrProjectNoPermission error.
	ErrProjectNoPermissionMsg = "user does not have permission to the project"

	// ErrProjectIDNotExistsMsg is the error message for ErrProjectIDNotExists error.
	ErrProjectIDNotExistsMsg = "project ID does not exist"

	// ErrProjectNameAlreadyExistsMsg is the error message for ErrProjectNameAlreadyExists error.
	ErrProjectNameAlreadyExistsMsg = "project name already exists"

	// ErrProjectMismatchMsg is the error message for ErrProjectMismatch error.
	ErrProjectMismatchMsg = "id/name pair not found on server side"

	// ErrProjectNotFoundMsg is the error message for ErrProjectNotFound error.
	ErrProjectNotFoundMsg = "project not found on server side"

	// ErrProjectNotProvidedMsg is the error message for ErrProjectNotProvided error.
	ErrProjectNotProvidedMsg = "no project provided"

	// ErrProjectNoMemberProvidedMsg is the error message for ErrProjectNoMemberProvided error.
	ErrProjectNoMemberProvidedMsg = "no project member provided"

	// ErrProjectMemberMismatchMsg is the error message for ErrProjectMemberMismatch error.
	ErrProjectMemberMismatchMsg = "no user with id/name pair found on server side"

	// ErrProjectMemberIllegalFormatMsg is the error message for ErrProjectMemberIllegalFormat error.
	ErrProjectMemberIllegalFormatMsg = "illegal format of project member or project id is invalid, or LDAP DN is invalid"

	ErrProjectUserIsNoMemberMsg = "user is no member in project"
)

// ErrProjectIllegalIDFormat describes an illegal request format.
type ErrProjectIllegalIDFormat struct{}

// Error returns the error message.
func (e *ErrProjectIllegalIDFormat) Error() string {
	return ErrProjectIllegalIDFormatMsg
}

// ErrProjectUnauthorized describes an unauthorized request.
type ErrProjectUnauthorized struct{}

// Error returns the error message.
func (e *ErrProjectUnauthorized) Error() string {
	return ErrProjectUnauthorizedMsg
}

// ErrProjectInternalErrors describes server-side internal errors.
type ErrProjectInternalErrors struct{}

// Error returns the error message.
func (e *ErrProjectInternalErrors) Error() string {
	return ErrProjectInternalErrorsMsg
}

// ErrProjectNoPermission describes a request error without permission.
type ErrProjectNoPermission struct{}

// Error returns the error message.
func (e *ErrProjectNoPermission) Error() string {
	return ErrProjectNoPermissionMsg
}

// ErrProjectIDNotExists describes an error
// when no proper project ID is found.
type ErrProjectIDNotExists struct{}

// Error returns the error message.
func (e *ErrProjectIDNotExists) Error() string {
	return ErrProjectIDNotExistsMsg
}

// ErrProjectNameAlreadyExists describes a duplicate project name error.
type ErrProjectNameAlreadyExists struct{}

// Error returns the error message.
func (e *ErrProjectNameAlreadyExists) Error() string {
	return ErrProjectNameAlreadyExistsMsg
}

// ErrProjectMismatch describes a failed lookup
// of a project with name/id pair.
type ErrProjectMismatch struct{}

// Error returns the error message.
func (e *ErrProjectMismatch) Error() string {
	return ErrProjectMismatchMsg
}

// ErrProjectNotFound describes an error
// when a specific project is not found.
type ErrProjectNotFound struct{}

// Error returns the error message.
func (e *ErrProjectNotFound) Error() string {
	return ErrProjectNotFoundMsg
}

type ErrProjectNotProvided struct{}

// Error returns the error message.
func (e *ErrProjectNotProvided) Error() string {
	return ErrProjectNotProvidedMsg
}

// ErrProjectNoMemberProvided
type ErrProjectNoMemberProvided struct{}

// Error returns the error message.
func (e *ErrProjectNoMemberProvided) Error() string {
	return ErrProjectNoMemberProvidedMsg
}

// ErrProjectMemberMismatch describes an error
// when user does not exist in context of
// project member operations.
type ErrProjectMemberMismatch struct{}

// Error returns the error message.
func (e *ErrProjectMemberMismatch) Error() string {
	return ErrProjectMemberMismatchMsg
}

// ErrProjectMemberIllegalFormat describes an communication
// error when performing project member operations.
type ErrProjectMemberIllegalFormat struct{}

// Error returns the error message.
func (e *ErrProjectMemberIllegalFormat) Error() string {
	return ErrProjectMemberIllegalFormatMsg
}

// ErrProjectUserIsNoMember describes an error case,
// where a given user is no member of a given project.
type ErrProjectUserIsNoMember struct{}

// Error returns the error message.
func (e *ErrProjectUserIsNoMember) Error() string {
	return ErrProjectUserIsNoMemberMsg
}
