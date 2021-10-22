package apiv2

import (
	"context"

	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/config"
)

func ExampleNewRESTClientForHost() {
	ctx := context.Background()

	// Create a RESTClient for a Harbor instance
	harborClient, err := NewRESTClientForHost("harbor.domain.com/api", "admin", "password", &config.Options{})
	if err != nil {
		panic(err)
	}

	harborClient.system.Health(ctx)
}
