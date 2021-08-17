// +build integration

package auditlog

import (
	"context"
	"net/url"
	"strconv"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/require"

	v2client "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v4/apiv2/internal/legacyapi/client"
	"github.com/mittwald/goharbor-client/v4/apiv2/project"
	integrationtest "github.com/mittwald/goharbor-client/v4/apiv2/testing"
)

var (
	u, _                = url.Parse(integrationtest.Host)
	legacySwaggerClient = client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	v2SwaggerClient     = v2client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	authInfo            = runtimeclient.BasicAuth(integrationtest.User, integrationtest.Password)
)

// TestAPIListAuditLogs tests listing the latest auditlog entry by creating
// a project and expecting the audit log entry to contain the proper metadata.
func TestAPIListAuditLogs(t *testing.T) {
	ctx := context.Background()
	pageSize := int64(1)
	storageLimit := int64(0)

	c := NewClient(v2SwaggerClient, authInfo)

	pc := project.NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	p, err := pc.NewProject(ctx, "test-auditlog", &storageLimit)

	defer pc.DeleteProject(ctx, p)

	a, err := c.ListAuditLogs(ctx, &pageSize, nil)
	require.NoError(t, err)

	require.Equal(t, 1, len(a))

	require.Equal(t, "create", a[0].Operation)
	require.Equal(t, "test-auditlog", a[0].Resource)
	require.Equal(t, "project", a[0].ResourceType)
	require.Equal(t, "admin", a[0].Username)
}

func TestAPIListAuditLogs_BigPageSize(t *testing.T) {
	ctx := context.Background()
	pageSize := int64(42)
	storageLimit := int64(0)

	c := NewClient(v2SwaggerClient, authInfo)

	pc := project.NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	for i := 0; i < 42; i++ {
		p, err := pc.NewProject(ctx, "test-auditlog-"+strconv.Itoa(i), &storageLimit)
		require.NoError(t, err)

		defer pc.DeleteProject(ctx, p)
	}

	a, err := c.ListAuditLogs(ctx, &pageSize, nil)
	require.NoError(t, err)

	require.Equal(t, 42, len(a))
}
