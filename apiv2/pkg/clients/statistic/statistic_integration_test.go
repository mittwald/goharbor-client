//go:build integration

package statistic

import (
	"context"
	"testing"

	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"
	"github.com/stretchr/testify/require"
)

func TestAPINewRobotAccount(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	statistic, err := c.GetStatistic(ctx)
	require.NoError(t, err)
	require.NotNil(t, statistic)
}
