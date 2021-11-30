//go:build integration

package systeminfo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"
)

func TestAPIGetSystemInfo(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	resp, err := c.GetSystemInfo(ctx)
	require.NoError(t, err)

	require.Equal(t, false, *resp.WithChartmuseum)
	require.Equal(t, false, *resp.WithNotary)
	require.Equal(t, *resp.RegistryURL, "localhost")
}
