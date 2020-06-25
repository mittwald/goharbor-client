package goharborclient

import (
	"github.com/go-openapi/runtime"
	"github.com/mittwald/goharbor-client/internal/api/v1.10.0/client"
)

func ExampleNewRESTClient() {
	var h *client.Harbor
	var authInfo runtime.ClientAuthInfoWriter

	cl := NewRESTClient(h, authInfo)

	_ = cl
}
