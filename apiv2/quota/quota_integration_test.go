// +build integration

package quota

import (
	"context"
	"net/url"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	v2client "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v4/apiv2/internal/legacyapi/client"
	"github.com/mittwald/goharbor-client/v4/apiv2/project"
	integrationtest "github.com/mittwald/goharbor-client/v4/apiv2/testing"
	"github.com/stretchr/testify/require"
)

var (
	u, _                        = url.Parse(integrationtest.Host)
	legacySwaggerClient         = client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	v2SwaggerClient             = v2client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	authInfo                    = runtimeclient.BasicAuth(integrationtest.User, integrationtest.Password)
	storageLimitPositive  int64 = 1
	storageLimitNegative  int64 = -1
	storageLimitNegative2 int64 = -1000
	storageLimitNull      int64 = 0
	testProjectName             = "test-project"
)

// TestAPIGetQuotaByProjectID_PositiveQuota creates a project with a positive storage limit set,
// gets the storage quota and compares the fetched value.
func TestAPIGetQuotaByProjectID_PositiveQuota(t *testing.T) {
	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	pc := project.NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)
	p, err := pc.NewProject(ctx, testProjectName, &storageLimitPositive)
	defer pc.DeleteProject(ctx, p)

	project, err := pc.GetProject(ctx, testProjectName)
	require.NoError(t, err)
	require.NotNil(t, project)

	q, err := c.GetQuotaByProjectID(ctx, int64(project.ProjectID))
	require.NoError(t, err)
	require.NotNil(t, q)

	require.Equal(t, storageLimitPositive, q.Hard["storage"])
}

// TestAPIGetQuotaByProjectID_NegativeQuota creates a project with a negative storage limit set,
// gets the storage quota and compares the fetched value.
func TestAPIGetQuotaByProjectID_NegativeQuota(t *testing.T) {
	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	pc := project.NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)
	p, err := pc.NewProject(ctx, testProjectName, &storageLimitNegative)
	defer pc.DeleteProject(ctx, p)

	project, err := pc.GetProject(ctx, testProjectName)
	require.NoError(t, err)

	q, err := c.GetQuotaByProjectID(ctx, int64(project.ProjectID))
	require.NoError(t, err)
	require.NotNil(t, q)
	require.Equal(t, storageLimitNegative, q.Hard["storage"])
}

func TestAPIUpdateQuotaByProjectID(t *testing.T) {
	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	// TestAPIUpdateQuotaByProjectID_PositiveQuota creates a project with a negative storage limit.
	// Updates the projects storage quota to a positive value and compares the observed values.
	t.Run("PositiveQuota", func(t *testing.T) {
		pc := project.NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)
		p, err := pc.NewProject(ctx, testProjectName, &storageLimitPositive)
		defer pc.DeleteProject(ctx, p)

		project, err := pc.GetProject(ctx, testProjectName)
		require.NoError(t, err)

		err = c.UpdateStorageQuotaByProjectID(ctx, int64(project.ProjectID), storageLimitPositive)
		require.NoError(t, err)

		q, err := c.GetQuotaByProjectID(ctx, int64(project.ProjectID))
		require.NoError(t, err)
		require.NotNil(t, q)
		require.Equal(t, storageLimitPositive, q.Hard["storage"])
	})

	// TestAPIUpdateQuotaByProjectID_NegativeQuota creates a project with a positive storage limit.
	// Updates the projects storage quota to a negative value and compares the observed values.
	t.Run("NegativeQuota", func(t *testing.T) {
		pc := project.NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)
		p, err := pc.NewProject(ctx, testProjectName, &storageLimitNegative)
		defer pc.DeleteProject(ctx, p)

		project, err := pc.GetProject(ctx, testProjectName)
		require.NoError(t, err)

		err = c.UpdateStorageQuotaByProjectID(ctx, int64(project.ProjectID), storageLimitNegative)
		require.NoError(t, err)

		q, err := c.GetQuotaByProjectID(ctx, int64(project.ProjectID))
		require.NoError(t, err)
		require.NotNil(t, q)
		require.Equal(t, storageLimitNegative, q.Hard["storage"])
	})

	// TestAPIUpdateQuotaByProjectID_NegativeQuota_2 creates a project with a storage limit set.
	// Tries updating the storage quota to a value of "-1000",
	// which is expected to result in the quota being implicitly set to '-1'.
	t.Run("NegativeQuota_2", func(t *testing.T) {
		pc := project.NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)
		p, err := pc.NewProject(ctx, testProjectName, &storageLimitNegative)
		defer pc.DeleteProject(ctx, p)

		project, err := pc.GetProject(ctx, testProjectName)
		require.NoError(t, err)

		err = c.UpdateStorageQuotaByProjectID(ctx, int64(project.ProjectID), storageLimitNegative2)
		require.NoError(t, err)

		q, err := c.GetQuotaByProjectID(ctx, int64(project.ProjectID))
		require.NoError(t, err)
		require.NotNil(t, q)
		require.Equal(t, storageLimitNegative, q.Hard["storage"])
	})

	// TestAPIUpdateQuotaByProjectID_NullQuota creates a project with a storage limit set.
	// Tries updating the storage quota to a value of "0",
	// which is expected to result in the quota being implicitly set to '-1'.
	t.Run("NullQuota", func(t *testing.T) {
		pc := project.NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)
		p, err := pc.NewProject(ctx, testProjectName, &storageLimitNegative)
		defer pc.DeleteProject(ctx, p)

		project, err := pc.GetProject(ctx, testProjectName)
		require.NoError(t, err)

		err = c.UpdateStorageQuotaByProjectID(ctx, int64(project.ProjectID), storageLimitNull)
		require.NoError(t, err)

		q, err := c.GetQuotaByProjectID(ctx, int64(project.ProjectID))
		require.NoError(t, err)
		require.NotNil(t, q)
		require.Equal(t, storageLimitNegative, q.Hard["storage"])
	})
}
