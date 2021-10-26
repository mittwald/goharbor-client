package errors

type (
	// ErrUnauthorized describes an unauthorized request.
	ErrUnauthorized struct{}
	// ErrBadRequest describes a malformed / invalid request.
	ErrBadRequest struct{}
	// ErrForbidden describes a forbidden request.
	ErrForbidden struct{}
	// ErrNotFound describes an error when no corresponding resources could be found.
	ErrNotFound struct{}
	// ErrMultipleResults describes an error when multiple
	// resources were found by an API call that should return only one.
	ErrMultipleResults struct{}
	// ErrInternalErrors describes a generic error that led to an internal server error.
	ErrInternalErrors struct{}
)

const (
	ErrUnauthorizedMsg = "unauthorized"
	ErrBadRequestMsg   = "bad request"
	ErrForbiddenMsg    = "forbidden"
	ErrNotFoundMsg     = "not found"

	ErrMultipleResultsMsg = "multiple results found"
	ErrInternalErrorsMsg  = "internal server error"
)

// Error returns the error message.
func (e *ErrUnauthorized) Error() string {
	return ErrUnauthorizedMsg
}

func (e *ErrBadRequest) Error() string {
	return ErrBadRequestMsg
}

func (e *ErrForbidden) Error() string {
	return ErrForbiddenMsg
}

func (e *ErrNotFound) Error() string {
	return ErrNotFoundMsg
}

func (e *ErrMultipleResults) Error() string {
	return ErrMultipleResultsMsg
}

func (e *ErrInternalErrors) Error() string {
	return ErrInternalErrorsMsg
}
