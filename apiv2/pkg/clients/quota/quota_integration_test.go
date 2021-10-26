//go:build integration

package quota

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/clients/project"

	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"
)

var (
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
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	pc := project.NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)
	err := pc.NewProject(ctx, testProjectName, &storageLimitPositive)
	require.NoError(t, err)

	defer pc.DeleteProject(ctx, testProjectName)

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
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	pc := project.NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)
	err := pc.NewProject(ctx, testProjectName, &storageLimitNegative)
	require.NoError(t, err)
	defer pc.DeleteProject(ctx, testProjectName)

	project, err := pc.GetProject(ctx, testProjectName)
	require.NoError(t, err)

	q, err := c.GetQuotaByProjectID(ctx, int64(project.ProjectID))
	require.NoError(t, err)
	require.NotNil(t, q)
	require.Equal(t, storageLimitNegative, q.Hard["storage"])
}

func TestAPIUpdateQuotaByProjectID(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	// TestAPIUpdateQuotaByProjectID_PositiveQuota creates a project with a negative storage limit.
	// Updates the projects storage quota to a positive value and compares the observed values.
	t.Run("PositiveQuota", func(t *testing.T) {
		pc := project.NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)
		err := pc.NewProject(ctx, testProjectName, &storageLimitPositive)
		require.NoError(t, err)
		defer pc.DeleteProject(ctx, testProjectName)

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
		pc := project.NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)
		err := pc.NewProject(ctx, testProjectName, &storageLimitNegative)
		defer pc.DeleteProject(ctx, testProjectName)

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
		pc := project.NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)
		err := pc.NewProject(ctx, testProjectName, &storageLimitNegative)
		require.NoError(t, err)

		defer pc.DeleteProject(ctx, testProjectName)

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
		pc := project.NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)
		err := pc.NewProject(ctx, testProjectName, &storageLimitNegative)
		require.NoError(t, err)
		defer pc.DeleteProject(ctx, testProjectName)

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
