// +build integration

package retention

import (
	"context"
	"flag"
	"net/url"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	v2client "github.com/mittwald/goharbor-client/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/apiv2/internal/legacyapi/client"
	pc "github.com/mittwald/goharbor-client/apiv2/project"
	integrationtest "github.com/mittwald/goharbor-client/apiv2/testing"
	"github.com/stretchr/testify/require"
)

var (
	u, _                = url.Parse(integrationtest.Host)
	legacySwaggerClient = client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	v2SwaggerClient     = v2client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	authInfo            = runtimeclient.BasicAuth(integrationtest.User, integrationtest.Password)
	harborVersion       = flag.String("version", "2.0.2",
		"Harbor version, used in conjunction with -integration, "+
			"defaults to 2.0.2")
	skipSpinUp = flag.Bool("skip-spinup", false,
		"Skip kind cluster creation")
	projectName = "test-project"
)

func TestAPIRetentionNew(t *testing.T) {
	ctx := context.Background()

	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	pc := pc.NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	p, err := pc.NewProject(ctx, projectName, 0, 0)
	require.NoError(t, err)

	defer pc.DeleteProject(ctx, p)

	p, err = pc.GetProject(ctx, projectName)

	err = c.NewRetentionPolicy(ctx, ScopeSelectorRepoMatches, int64(p.ProjectID), PolicyTemplateDaysSinceLastPush, TagSelectorMatches,
		map[PolicyTemplate]interface{}{PolicyTemplateDaysSinceLastPush: 1}, "**", "**", "0 * * * *", true)

	require.NoError(t, err)

	require.Nil(t, err)
}

func TestAPIRetentionGet(t *testing.T) {
	ctx := context.Background()

	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	pc := pc.NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	p, err := pc.NewProject(ctx, projectName, 0, 0)
	require.NoError(t, err)

	defer pc.DeleteProject(ctx, p)

	p, err = pc.GetProject(ctx, projectName)

	err = c.NewRetentionPolicy(ctx, ScopeSelectorRepoMatches, int64(p.ProjectID), PolicyTemplateDaysSinceLastPush, TagSelectorMatches,
		map[PolicyTemplate]interface{}{PolicyTemplateDaysSinceLastPush: 1}, "**", "**", "0 * * * *", true)

	require.NoError(t, err)
	require.Nil(t, err)

	rp, err := c.GetRetentionPolicyByProjectID(ctx, int64(p.ProjectID))
	require.NoError(t, err)
	require.NotNil(t, rp)
}
