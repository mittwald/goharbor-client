// +build integration

package quota

import (
	"context"
	"net/url"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	v2client "github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/legacyapi/client"
	"github.com/mittwald/goharbor-client/v3/apiv2/project"
	integrationtest "github.com/mittwald/goharbor-client/v3/apiv2/testing"
	"github.com/stretchr/testify/require"
)

var (
	u, _                       = url.Parse(integrationtest.Host)
	legacySwaggerClient        = client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	v2SwaggerClient            = v2client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	authInfo                   = runtimeclient.BasicAuth(integrationtest.User, integrationtest.Password)
	storageLimitPositive int64 = 1
	storageLimitNegative int64 = -1
	testProjectName            = "test-project"
)

func TestAPIGetQuotaByProjectID_PositiveQuota(t *testing.T) {
	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	pc := project.NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)
	p, err := pc.NewProject(ctx, testProjectName, &storageLimitPositive)
	defer pc.DeleteProject(ctx, p)

	project, err := pc.GetProjectByName(ctx, testProjectName)
	require.NoError(t, err)

	q, err := c.GetQuotaByProjectID(ctx, int64(project.ProjectID))
	require.NoError(t, err)
	require.NotNil(t, q)

	require.Equal(t, q.Hard["storage"], storageLimitPositive)
}

func TestAPIGetQuotaByProjectID_NegativeQuota(t *testing.T) {
	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	pc := project.NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)
	p, err := pc.NewProject(ctx, testProjectName, &storageLimitNegative)
	defer pc.DeleteProject(ctx, p)

	project, err := pc.GetProjectByName(ctx, testProjectName)
	require.NoError(t, err)

	q, err := c.GetQuotaByProjectID(ctx, int64(project.ProjectID))
	require.NoError(t, err)
	require.NotNil(t, q)
	require.Equal(t, q.Hard["storage"], storageLimitNegative)
}
