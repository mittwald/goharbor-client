//go:build integration

package health

import (
	"context"
	"net/url"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/require"

	v2client "github.com/testwill/goharbor-client/v5/apiv2/internal/api/client"
	"github.com/testwill/goharbor-client/v5/apiv2/pkg/config"
	integrationtest "github.com/testwill/goharbor-client/v5/apiv2/pkg/testing"
)

var (
	u, _            = url.Parse(integrationtest.Host)
	v2SwaggerClient = v2client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	authInfo        = runtimeclient.BasicAuth(integrationtest.User, integrationtest.Password)
	defaultOpts     = config.Defaults()
)

func TestAPIGetHealth(t *testing.T) {
	ctx := context.Background()
	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)

	health, err := c.GetHealth(ctx)
	require.NoError(t, err)

	for _, c := range health.Components {
		require.NotEmpty(t, c.Status)
	}
}
