//go:build integration

package configure

import (
	"context"
	"testing"

	modelv2 "github.com/mittwald/goharbor-client/v5/apiv2/model"
	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"

	"github.com/stretchr/testify/require"
)

func TestAPIGetConfig(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	resp, err := c.GetConfigs(ctx)
	require.NoError(t, err)

	require.Equal(t, "db_auth", *resp.AuthMode)
}

func TestAPIUpdateConfigs(t *testing.T) {
	authMode := "oidc"
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := c.UpdateConfigs(ctx, &modelv2.Configurations{AuthMode: authMode})
	require.NoError(t, err)
}
