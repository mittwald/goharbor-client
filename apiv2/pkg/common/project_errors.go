package common

const (
	// ErrProjectIllegalIDFormatMsg is the error message for ErrProjectIllegalIDFormat error.
	ErrProjectIllegalIDFormatMsg = "illegal format of provided ID value"

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

	// ErrProjectInvalidRequestMsg is the error message for ErrProjectInvalidRequest error.
	ErrProjectInvalidRequestMsg = "invalid request"

	// ErrProjectUnknownResourceMsg is the error message for ErrProjectUnknownResource error.
	ErrProjectUnknownResourceMsg = "resource unknown"

	// ErrProjectNameNotProvidedMsg is the error message for ErrProjectNameNotProvided error.
	ErrProjectNameNotProvidedMsg = "project name not provided"

	// ErrProjectNoWebhookPolicyProvidedMsg is the error message for ErrProjectNoWebhookPolicyProvided error.
	ErrProjectNoWebhookPolicyProvidedMsg = "no webhook policy provided"
)

type (
	// ErrProjectNameNotProvided describes a missing project name.
	ErrProjectNameNotProvided struct{}
	// ErrProjectIllegalIDFormat describes an illegal request format.
	ErrProjectIllegalIDFormat struct{}
	// ErrProjectInternalErrors describes server-side internal errors.
	ErrProjectInternalErrors struct{}
	// ErrProjectNoPermission describes a request error without permission.
	ErrProjectNoPermission struct{}
	// ErrProjectIDNotExists describes an error
	// when no proper project ID is found.
	ErrProjectIDNotExists struct{}
	// ErrProjectNameAlreadyExists describes a duplicate project name error.
	ErrProjectNameAlreadyExists struct{}

	// ErrProjectMismatch describes a failed lookup
	// of a project with name/id pair.
	ErrProjectMismatch struct{}
	// ErrProjectNotFound describes an error
	// when a specific project is not found.
	ErrProjectNotFound struct{}

	ErrProjectNotProvided struct{}
	// ErrProjectNoMemberProvided describes an error when a project's member is not provided.
	ErrProjectNoMemberProvided struct{}
	// ErrProjectMemberMismatch describes an error
	// when user does not exist in context of
	// project member operations.
	ErrProjectMemberMismatch struct{}

	// ErrProjectMemberIllegalFormat describes a communication error when performing project member operations.
	ErrProjectMemberIllegalFormat struct{}
	// ErrProjectInvalidRequest describes a communication
	// error when performing project member operations.
	ErrProjectInvalidRequest struct{}
	// ErrProjectUnknownResource describes an error after requesting an unknown resource.
	ErrProjectUnknownResource struct{}
	// ErrProjectNoWebhookPolicyProvided describes an error when no webhook policy is provided.
	ErrProjectNoWebhookPolicyProvided struct{}
)

// Error returns the error message.
func (e *ErrProjectNameNotProvided) Error() string {
	return ErrProjectNameNotProvidedMsg
}

// Error returns the error message.
func (e *ErrProjectIllegalIDFormat) Error() string {
	return ErrProjectIllegalIDFormatMsg
}

// Error returns the error message.
func (e *ErrProjectInternalErrors) Error() string {
	return ErrProjectInternalErrorsMsg
}

// Error returns the error message.
func (e *ErrProjectNoPermission) Error() string {
	return ErrProjectNoPermissionMsg
}

// Error returns the error message.
func (e *ErrProjectIDNotExists) Error() string {
	return ErrProjectIDNotExistsMsg
}

// Error returns the error message.
func (e *ErrProjectNameAlreadyExists) Error() string {
	return ErrProjectNameAlreadyExistsMsg
}

// Error returns the error message.
func (e *ErrProjectMismatch) Error() string {
	return ErrProjectMismatchMsg
}

// Error returns the error message.
func (e *ErrProjectNotFound) Error() string {
	return ErrProjectNotFoundMsg
}

// Error returns the error message.
func (e *ErrProjectNotProvided) Error() string {
	return ErrProjectNotProvidedMsg
}

// Error returns the error message.
func (e *ErrProjectInvalidRequest) Error() string {
	return ErrProjectInvalidRequestMsg
}

// Error returns the error message.
func (e *ErrProjectUnknownResource) Error() string {
	return ErrProjectUnknownResourceMsg
}

// Error returns the error message.
func (e *ErrProjectNoWebhookPolicyProvided) Error() string {
	return ErrProjectNoWebhookPolicyProvidedMsg
}
