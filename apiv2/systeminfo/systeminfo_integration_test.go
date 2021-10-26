////go:build integration

package systeminfo

import (
	"context"
	"net/url"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/require"

	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/config"
	integrationtest "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"

	v2client "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v5/apiv2/internal/legacyapi/client"
)

var (
	u, _                = url.Parse(integrationtest.Host)
	legacySwaggerClient = client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	v2SwaggerClient     = v2client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	authInfo            = runtimeclient.BasicAuth(integrationtest.User, integrationtest.Password)
	opts                = config.Options{}
	defaultOpts         = opts.Defaults()
)

func TestAPIGetSystemInfo(t *testing.T) {
	ctx := context.Background()
	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)

	resp, err := c.GetSystemInfo(ctx)
	require.NoError(t, err)

	require.Equal(t,false,  *resp.WithChartmuseum )
	require.Equal(t, false, *resp.WithNotary)
	require.Equal(t, *resp.RegistryURL, "localhost")
}
