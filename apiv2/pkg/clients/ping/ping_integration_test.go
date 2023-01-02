//go:build integration

package ping

import (
	"context"
	"testing"

	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"
	"github.com/stretchr/testify/require"
)

func TestAPIGetPing(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	res, err := c.GetPing(ctx)
	require.NoError(t, err)
	require.Equal(t, "Pong", res)
}
