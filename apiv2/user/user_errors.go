package user

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/legacyapi/client/products"
)

const (
	// ErrUserNotFoundMsg is the error message for ErrUserNotFound error.
	ErrUserNotFoundMsg = "user not found on server side"

	// ErrUserBadRequestMsg is the error message for ErrUserBadRequest error.
	ErrUserBadRequestMsg = "Unsatisfied with constraints of the user creation/modification."

	// ErrUserMismatchMsg is the error message for ErrUserMismatch error.
	ErrUserMismatchMsg = "id/name pair not found on server side"

	// ErrUserAlreadyExistMsg is the error message for ErrUserAlreadyExists error.
	ErrUserAlreadyExistsMsg = "user with this username already exists"

	// ErrUserInvalidIDMsg is the error message for ErrUserInvalidID error.
	ErrUserInvalidIDMsg = "invalid user ID"

	// ErrUserPasswordInvalid  is the error message for ErrUserPasswordInvalid error.
	ErrUserPasswordInvalidMsg = "invalid user password"
)

// ErrUserNotFound describes an error when a specific user was not found on server side.
type ErrUserNotFound struct{}

// Error returns the error message.
func (e *ErrUserNotFound) Error() string {
	return ErrUserNotFoundMsg
}

// ErrUserBadRequest describes a formal error when creating or updating a user (such as bad password).
type ErrUserBadRequest struct{}

// Error returns the error message.
func (e *ErrUserBadRequest) Error() string {
	return ErrUserBadRequestMsg
}

// ErrUserMismatch describes an error when the id and name of a user do not match on server side.
type ErrUserMismatch struct{}

// Error returns the error message.
func (e *ErrUserMismatch) Error() string {
	return ErrUserMismatchMsg
}

// ErrUserAlreadyExists describes an error indicating that this user already exists.
type ErrUserAlreadyExists struct{}

// Error returns the error message.
func (e *ErrUserAlreadyExists) Error() string {
	return ErrUserAlreadyExistsMsg
}

// ErrUserInvalidID describes an error indicating an invalid user id.
type ErrUserInvalidID struct{}

// Error returns the error message.
func (e *ErrUserInvalidID) Error() string {
	return ErrUserInvalidIDMsg
}

// ErrUserPasswordInvalid describes an error indicating an invalid password
type ErrUserPasswordInvalid struct{}

func (e *ErrUserPasswordInvalid) Error() string {
	return ErrUserPasswordInvalidMsg
}

// handleUserErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with proper message.
func handleSwaggerUserErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case http.StatusConflict:
			return &ErrUserAlreadyExists{}
		}
	}

	switch in.(type) {
	case *products.PostUsersBadRequest:
		return &ErrUserBadRequest{}
	case *products.PutUsersUserIDBadRequest:
		return &ErrUserInvalidID{}
	case *products.PutUsersUserIDPasswordBadRequest:
		return &ErrUserPasswordInvalid{}
	default:
		return in
	}
}
