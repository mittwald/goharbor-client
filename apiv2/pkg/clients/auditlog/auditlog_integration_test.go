//go:build integration

package auditlog

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	modelv2 "github.com/mittwald/goharbor-client/v5/apiv2/model"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/clients/project"

	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"
)

// TestAPIListAuditLogs tests listing the latest auditlog entry by creating
// a project and expecting the audit log entry to contain the proper metadata.
func TestAPIListAuditLogs(t *testing.T) {
	ctx := context.Background()
	storageLimit := int64(0)

	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	pc := project.NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := pc.NewProject(ctx, &modelv2.ProjectReq{
		ProjectName:  "test-auditlog",
		StorageLimit: &storageLimit,
	})
	require.NoError(t, err)

	p, err := pc.GetProject(ctx, "test-auditlog")
	require.NoError(t, err)

	defer pc.DeleteProject(ctx, p.Name)

	c.Options.PageSize = 100

	a, err := c.ListAuditLogs(ctx)
	require.NoError(t, err)

	require.Greater(t, len(a), 0)

	require.Equal(t, "create", a[0].Operation)
	require.Equal(t, "test-auditlog", a[0].Resource)
	require.Equal(t, "project", a[0].ResourceType)
	require.Equal(t, "admin", a[0].Username)
}
