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
	// ErrConflict describes an error when a resource already exists.
	ErrConflict struct{}
	// ErrMultipleResults describes an error when multiple
	// resources were found by an API call that should return only one.
	ErrMultipleResults struct{}
	// ErrUnsupportedMediaType describes an error when the request contains an unsupported media type.
	ErrUnsupportedMediaType struct{}
	// ErrInternalErrors describes a generic error that led to an internal server error.
	ErrInternalErrors struct{}
)

const (
	ErrUnauthorizedMsg         = "unauthorized"
	ErrBadRequestMsg           = "bad request"
	ErrForbiddenMsg            = "forbidden"
	ErrConflictMsg             = "conflict"
	ErrNotFoundMsg             = "not found"
	ErrMultipleResultsMsg      = "multiple results found"
	ErrUnsupportedMediaTypeMsg = "unsupported media type"
	ErrInternalErrorsMsg       = "internal server error"
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

func (e *ErrConflict) Error() string {
	return ErrConflictMsg
}

func (e *ErrUnsupportedMediaType) Error() string {
	return ErrUnsupportedMediaTypeMsg
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
