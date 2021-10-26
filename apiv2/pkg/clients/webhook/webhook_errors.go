package webhook

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// handleSwaggerWebhookErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with a proper message.
func handleSwaggerWebhookErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case http.StatusCreated:
			// Harbor sometimes return 201 instead of 200 despite the swagger spec
			// not declaring it.
			return nil
		}
	}

	switch in.(type) {
	default:
		return in
	}
}
