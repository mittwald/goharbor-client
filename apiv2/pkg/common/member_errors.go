package common

type (
	ErrNoMemberFound struct{}

	ErrMemberAlreadyExists struct{}
)

const (
	// ErrProjectNoMemberProvidedMsg is the error message for ErrProjectNoMemberProvided error.
	ErrProjectNoMemberProvidedMsg = "no project member provided"

	// ErrProjectMemberMismatchMsg is the error message for ErrProjectMemberMismatch error.
	ErrProjectMemberMismatchMsg = "no user or group with id/name pair found on server side"
	// ErrProjectMemberIllegalFormatMsg is the error message for ErrProjectMemberIllegalFormat error.
	ErrProjectMemberIllegalFormatMsg = "illegal format of project member or project id is invalid, or LDAP DN is invalid"

	ErrNoMemberFoundMsg = "no project members found"

	ErrMemberAlreadyExistsMsg = "member already exists"
)

// Error returns the error message.
func (e *ErrProjectNoMemberProvided) Error() string {
	return ErrProjectNoMemberProvidedMsg
}

// Error returns the error message.
func (e *ErrProjectMemberMismatch) Error() string {
	return ErrProjectMemberMismatchMsg
}

// Error returns the error message.
func (e *ErrProjectMemberIllegalFormat) Error() string {
	return ErrProjectMemberIllegalFormatMsg
}

func (e *ErrNoMemberFound) Error() string {
	return ErrNoMemberFoundMsg
}

func (e *ErrMemberAlreadyExists) Error() string {
	return ErrMemberAlreadyExistsMsg
}
