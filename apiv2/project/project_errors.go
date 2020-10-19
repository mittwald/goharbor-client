package project

import (
	projectapi "github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client/project"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/legacyapi/client/products"
)

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

	// ErrProjectUserIsNoMemberMsg is the error message for ErrProjectUserIsNoMember.
	ErrProjectUserIsNoMemberMsg = "user is no member in project"

	// ErrProjectInvalidRequestMsg is the error message for ErrProjectInvalidRequest error.
	ErrProjectInvalidRequestMsg = "invalid request"

	// ErrProjectMetadataAlreadyExistsMsg is the error message for ErrProjectMetadataAlreadyExists error.
	ErrProjectMetadataAlreadyExistsMsg = "metadata key already exists"

	// ErrProjectUnknownResourceMsg is the error message for ErrProjectUnknownResource error.
	ErrProjectUnknownResourceMsg = "resource unknown"

	// ErrProjectNameNotProvidedMsg is the error message for ErrProjectNameNotProvided error.
	ErrProjectNameNotProvidedMsg = "project name not provided"

	// ErrProjectMetadataUndefinedMsg is the error message for ErrProjectMetadataUndefined error.
	ErrProjectMetadataUndefinedMsg = "project metadata undefined"
)

// ErrProjectNameNotProvided describes a missing project name.
type ErrProjectNameNotProvided struct{}

// Error returns the error message.
func (e *ErrProjectNameNotProvided) Error() string {
	return ErrProjectNameNotProvidedMsg
}

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

// ErrProjectMemberIllegalFormat describes an communication
// error when performing project member operations.
type ErrProjectInvalidRequest struct{}

// Error returns the error message.
func (e *ErrProjectInvalidRequest) Error() string {
	return ErrProjectInvalidRequestMsg
}

// ErrProjectMetadataUndefined describes an error accessing a project's metadata.
type ErrProjectMetadataUndefined struct{}

// Error returns the error message.
func (e *ErrProjectMetadataUndefined) Error() string {
	return ErrProjectMetadataUndefinedMsg
}

// ErrProjectMetadataAlreadyExists describes an error, which happens
// when a metadata key of a project is tried to be created a second time.
type ErrProjectMetadataAlreadyExists struct{}

// Error returns the error message.
func (e *ErrProjectMetadataAlreadyExists) Error() string {
	return ErrProjectMetadataAlreadyExistsMsg
}

// ErrProjectUnknownResource describes which happens,
// when requesting an unknown ressource.
type ErrProjectUnknownResource struct{}

// Error returns the error message.
func (e *ErrProjectUnknownResource) Error() string {
	return ErrProjectUnknownResourceMsg
}

// handleProjectErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with proper message.
func handleSwaggerProjectErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case http.StatusCreated:
			// Harbor sometimes return 201 instead of 200 despite the swagger spec
			// not declaring it.
			return nil
		case http.StatusBadRequest:
			return &ErrProjectIllegalIDFormat{}
		case http.StatusUnauthorized:
			return &ErrProjectUnauthorized{}
		case http.StatusForbidden:
			return &ErrProjectNoPermission{}
		case http.StatusNotFound:
			return &ErrProjectUnknownResource{}
		case http.StatusInternalServerError:
			return &ErrProjectInternalErrors{}
		}
	}

	switch in.(type) {
	case *projectapi.DeleteProjectNotFound:
		return &ErrProjectIDNotExists{}
	case *projectapi.UpdateProjectNotFound:
		return &ErrProjectIDNotExists{}
	case *projectapi.CreateProjectConflict:
		return &ErrProjectNameAlreadyExists{}
	case *products.PostProjectsProjectIDMembersBadRequest:
		return &ErrProjectInvalidRequest{}
	case *products.PostProjectsProjectIDMetadatasBadRequest:
		return &ErrProjectInvalidRequest{}
	default:
		return in
	}
}
