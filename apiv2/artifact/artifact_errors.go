package artifact

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/artifact"
)

const (
	ErrArtifactBadRequestMsg          = "bad request"
	ErrArtifactUnauthorizedMsg        = "unauthorized"
	ErrArtifactInternalServerErrorMsg = "internal server error"
	ErrArtifactTagNotFoundMsg         = "could not find project, repository or reference"
)

type ErrArtifactBadRequest struct{}

func (e *ErrArtifactBadRequest) Error() string {
	return ErrArtifactBadRequestMsg
}

type ErrArtifactUnauthorized struct{}

func (e *ErrArtifactUnauthorized) Error() string {
	return ErrArtifactUnauthorizedMsg
}

type ErrArtifactInternalServerError struct{}

func (e *ErrArtifactInternalServerError) Error() string {
	return ErrArtifactInternalServerErrorMsg
}

type ErrArtifactTagNotFound struct{}

func (e *ErrArtifactTagNotFound) Error() string {
	return ErrArtifactTagNotFoundMsg
}

func handleSwaggerArtifactErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case http.StatusBadRequest:
			return &ErrArtifactBadRequest{}
		case http.StatusUnauthorized:
			return &ErrArtifactUnauthorized{}
		case http.StatusInternalServerError:
			return &ErrArtifactInternalServerError{}
		}
	}

	switch in.(type) {
	case *artifact.ListArtifactsInternalServerError:
		return &ErrArtifactInternalServerError{}
	case *artifact.ListArtifactsBadRequest:
		return &ErrArtifactBadRequest{}
	case *artifact.ListArtifactsUnauthorized:
		return &ErrArtifactUnauthorized{}
	case *artifact.CreateTagNotFound:
		return &ErrArtifactTagNotFound{}
	default:
		return in
	}
}
