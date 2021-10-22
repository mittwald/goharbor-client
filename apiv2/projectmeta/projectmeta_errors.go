package projectmeta

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/common"
)

// handleSwaggerProjectMetaErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with a proper message.
func handleSwaggerProjectMetaErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case http.StatusCreated:
			// Harbor sometimes return 201 instead of 200 despite the swagger spec
			// not declaring it.
			return nil

		case http.StatusConflict:
			return &common.ErrProjectMetadataAlreadyExists{}
		}
	}

	switch in.(type) {
	default:
		return in
	}
}
