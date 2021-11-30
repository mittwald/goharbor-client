package errors

const (
	// ErrUserNotFoundMsg is the error message for ErrUserNotFound error.
	ErrUserNotFoundMsg = "user not found on server side"

	// ErrUserBadRequestMsg is the error message for ErrUserBadRequest error.
	ErrUserBadRequestMsg = "unsatisfied with constraints of the user creation/modification."

	// ErrUserMismatchMsg is the error message for ErrUserMismatch error.
	ErrUserMismatchMsg = "id/name pair not found on server side"

	// ErrUserAlreadyExistsMsg is the error message for ErrUserAlreadyExists error.
	ErrUserAlreadyExistsMsg = "user with this username already exists"

	// ErrUserInvalidIDMsg is the error message for ErrUserInvalidID error.
	ErrUserInvalidIDMsg = "invalid user ID"

	// ErrUserIDNotExistsMsg is the error message for ErrUserIDNotExists error.
	ErrUserIDNotExistsMsg = "user id does not exist"

	// ErrUserPasswordInvalidMsg  is the error message for ErrUserPasswordInvalid error.
	ErrUserPasswordInvalidMsg = "invalid user password"
)

type (
	// ErrUserNotFound describes an error when a specific user was not found on server side.
	ErrUserNotFound struct{}
	// ErrUserBadRequest describes a formal error when creating or updating a user (such as bad password).
	ErrUserBadRequest struct{}
	// ErrUserMismatch describes an error when the id and name of a user do not match on server side.
	ErrUserMismatch struct{}
	// ErrUserAlreadyExists describes an error indicating that this user already exists.
	ErrUserAlreadyExists struct{}
	// ErrUserInvalidID describes an error indicating an invalid user id.
	ErrUserInvalidID struct{}
	// ErrUserIDNotExists describes an error indicating a nonexistent user id.
	ErrUserIDNotExists struct{}
	// ErrUserPasswordInvalid describes an error indicating an invalid password.
	ErrUserPasswordInvalid struct{}
)

func (e *ErrUserNotFound) Error() string {
	return ErrUserNotFoundMsg
}

func (e *ErrUserBadRequest) Error() string {
	return ErrUserBadRequestMsg
}

func (e *ErrUserMismatch) Error() string {
	return ErrUserMismatchMsg
}

func (e *ErrUserAlreadyExists) Error() string {
	return ErrUserAlreadyExistsMsg
}

func (e *ErrUserInvalidID) Error() string {
	return ErrUserInvalidIDMsg
}

func (e *ErrUserIDNotExists) Error() string {
	return ErrUserIDNotExistsMsg
}

func (e *ErrUserPasswordInvalid) Error() string {
	return ErrUserPasswordInvalidMsg
}
