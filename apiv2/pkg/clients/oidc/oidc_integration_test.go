//go:build integration

package oidc

import (
	"context"
	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/oidc"
	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAPIPingOIDC(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	body := oidc.PingOIDCBody{
		URL:        "",
		VerifyCert: false,
	}

	err := c.PingOIDC(ctx, body)
	require.NoError(t, err)
}
