package goharborclient

const (
	// ErrUserNotFoundMsg is the error message for ErrUserNotFound error.
	ErrUserNotFoundMsg = "user not found on server side"

	// ErrUserBadRequestMsg is the error message for ErrUserBadRequest error.
	ErrUserBadRequestMsg = "Unsatisfied with constraints of the user creation/modification."

	// ErrUserMismatchMsg is the error message for ErrUserMismatch error.
	ErrUserMismatchMsg = "id/name pair not found on server side"

	// ErrUserAlreadyExistMsg is the error message for ErrUserAlreadyExist error.
	ErrUserAlreadyExistMsg = "user with this username already exists"

	// ErrUserInvalidIDMsg is the error message for ErrUserInvalidID error.
	ErrUserInvalidIDMsg = "invalid user ID"
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

// ErrUserAlreadyExist describes an error indicating that this user already exists.
type ErrUserAlreadyExist struct{}

// Error returns the error message.
func (e *ErrUserAlreadyExist) Error() string {
	return ErrUserAlreadyExistMsg
}

// ErrUserInvalidID describes an error indicating an invalid user id.
type ErrUserInvalidID struct{}

// Error returns the error message.
func (e *ErrUserInvalidID) Error() string {
	return ErrUserInvalidIDMsg
}
