package immutable

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/immutable"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/errors"
)

func handleSwaggerImmutableRuleErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case http.StatusOK:
			return nil
		case http.StatusCreated:
			return nil
		case http.StatusBadRequest:
			return &errors.ErrBadRequest{}
		case http.StatusUnauthorized:
			return &errors.ErrUnauthorized{}
		case http.StatusForbidden:
			return &errors.ErrForbidden{}
		case http.StatusNotFound:
			return &errors.ErrNotFound{}
		case http.StatusInternalServerError:
			return &errors.ErrInternalErrors{}
		}
	}

	switch in.(type) {
	case *immutable.UpdateImmuRuleBadRequest:
		return &errors.ErrBadRequest{}
	case *immutable.CreateImmuRuleBadRequest:
		return &errors.ErrBadRequest{}
	case *immutable.DeleteImmuRuleBadRequest:
		return &errors.ErrBadRequest{}
	default:
		return in
	}
}