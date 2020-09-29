package retention

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/mittwald/goharbor-client/v2/apiv2/internal/legacyapi/client/products"
)

const (
	// ErrRetentionUnauthorizedMsg is the error message for ErrRetentionUnauthorized error.
	ErrRetentionUnauthorizedMsg = "unauthorized"

	// ErrRetentionNoPermissionMsg is the error message for ErrRetentionNoPermission error.
	ErrRetentionNoPermissionMsg = "user does not have permission to the retention"

	// ErrRetentionInternalErrorsMsg is the error message for ErrRetentionInternalErrors error.
	ErrRetentionInternalErrorsMsg = "unexpected internal errors"

	// ErrRetentionNotProvidedMsg is the error message for ErrRetentionNotProvided error.
	ErrRetentionNotProvidedMsg = "no retention policy provided"
)

// ErrRetentionUnauthorized describes an unauthorized request.
type ErrRetentionUnauthorized struct{}

// Error returns the error message.
func (e *ErrRetentionUnauthorized) Error() string {
	return ErrRetentionUnauthorizedMsg
}

// ErrRetentionNotProvided describes a missing retention instance
type ErrRetentionNotProvided struct{}

// Error returns the error message.
func (e *ErrRetentionNotProvided) Error() string {
	return ErrRetentionNotProvidedMsg
}

// ErrRetentionNoPermission describes a request error without permission.
type ErrRetentionNoPermission struct{}

// Error returns the error message.
func (e *ErrRetentionNoPermission) Error() string {
	return ErrRetentionNoPermissionMsg
}

// ErrRetentionInternalErrors describes server-side internal errors.
type ErrRetentionInternalErrors struct{}

// Error returns the error message.
func (e *ErrRetentionInternalErrors) Error() string {
	return ErrRetentionInternalErrorsMsg
}

// handleProjectErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with proper message.
func handleSwaggerRetentionErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case http.StatusCreated:
			// Harbor sometimes return 201 instead of 200 despite the swagger spec
			// not declaring it.
			return nil
		case http.StatusUnauthorized:
			return &ErrRetentionUnauthorized{}
		case http.StatusForbidden:
			return &ErrRetentionNoPermission{}
		case http.StatusInternalServerError:
			return &ErrRetentionInternalErrors{}
		}
	}

	switch in.(type) {
	case *products.PostRetentionsIDExecutionsInternalServerError:
		return &ErrRetentionInternalErrors{}
	case *products.GetRetentionsIDExecutionsUnauthorized:
		return &ErrRetentionUnauthorized{}
	case *products.PostRetentionsForbidden:
		return &ErrRetentionUnauthorized{}
	default:
		return in
	}
}
