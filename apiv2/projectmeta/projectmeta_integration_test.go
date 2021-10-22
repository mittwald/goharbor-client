//go:build integration

package projectmeta

import (
	"context"
	"net/url"
	"strconv"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/require"

	v2client "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/common"
	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/config"
	integrationtest "github.com/mittwald/goharbor-client/v4/apiv2/pkg/testing"
	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/util"
	"github.com/mittwald/goharbor-client/v4/apiv2/project"
)

var (
	u, _                       = url.Parse(integrationtest.Host)
	v2SwaggerClient            = v2client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	authInfo                   = runtimeclient.BasicAuth(integrationtest.User, integrationtest.Password)
	storageLimitPositive int64 = 1
	storageLimitNegative int64 = -1
	opts                       = config.Options{}
	defaultOpts                = opts.Defaults()
)

func TestAPIProjectMetadataAdd(t *testing.T) {
	projectName := "test-project"

	ctx := context.Background()
	metaClient := NewClient(v2SwaggerClient, defaultOpts, authInfo)

	projectClient := project.NewClient(v2SwaggerClient, defaultOpts, authInfo)

	p, err := projectClient.NewProject(ctx, projectName, &storageLimitPositive)
	defer projectClient.DeleteProject(ctx, p)
	require.NoError(t, err)

	err = metaClient.AddProjectMetadata(ctx, util.ProjectIDAsString(p.ProjectID), common.MetadataKeyAutoScan, "true")
	require.NoError(t, err)
	err = metaClient.AddProjectMetadata(ctx, util.ProjectIDAsString(p.ProjectID), common.MetadataKeyEnableContentTrust, "true")
	require.NoError(t, err)
	err = metaClient.AddProjectMetadata(ctx, util.ProjectIDAsString(p.ProjectID), common.MetadataKeyPreventVul, "true")
	require.NoError(t, err)
	// TODO: Re-introduce this, once https://github.com/goharbor/harbor/pull/15800/ has been adopted.
	// err = metaClient.AddProjectMetadata(ctx, util.ProjectIDAsString(p.ProjectID), common.MetadataKeyReuseSysCVEAllowlist, "true")
	// require.NoError(t, err)
	err = metaClient.AddProjectMetadata(ctx, util.ProjectIDAsString(p.ProjectID), common.MetadataKeySeverity, "medium")
	require.NoError(t, err)
}

func TestAPIProjectMetadataGet(t *testing.T) {
	projectName := "test-project"

	ctx := context.Background()
	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)

	projectClient := project.NewClient(v2SwaggerClient, defaultOpts, authInfo)

	p, err := projectClient.NewProject(ctx, projectName, &storageLimitPositive)
	defer projectClient.DeleteProject(ctx, p)
	require.NoError(t, err)

	m, err := c.GetProjectMetadataValue(ctx, strconv.Itoa(int(p.ProjectID)), common.MetadataKeyPublic)
	require.NoError(t, err)

	require.Equal(t, "false", m)
}

func TestAPIProjectMetadataGetInvalidKey(t *testing.T) {
	projectName := "test-project"

	ctx := context.Background()
	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)

	projectClient := project.NewClient(v2SwaggerClient, defaultOpts, authInfo)

	p, err := projectClient.NewProject(ctx, projectName, &storageLimitPositive)
	defer projectClient.DeleteProject(ctx, p)
	require.NoError(t, err)

	m, err := c.GetProjectMetadataValue(ctx, strconv.Itoa(int(p.ProjectID)), "foobar")

	require.Error(t, err)
	require.Equal(t, "invalid request", err.Error())
	require.IsType(t, &common.ErrProjectInvalidRequest{}, err)
	require.Equal(t, "", m)
}

func TestAPIProjectMetadataList(t *testing.T) {
	projectName := "test-project"

	ctx := context.Background()
	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)

	projectClient := project.NewClient(v2SwaggerClient, defaultOpts, authInfo)
	p, err := projectClient.NewProject(ctx, projectName, &storageLimitPositive)
	defer projectClient.DeleteProject(ctx, p)
	require.NoError(t, err)

	err = c.AddProjectMetadata(ctx, strconv.Itoa(int(p.ProjectID)), common.MetadataKeyAutoScan, "true")
	require.NoError(t, err)
	err = c.AddProjectMetadata(ctx, strconv.Itoa(int(p.ProjectID)), common.MetadataKeyEnableContentTrust, "true")
	require.NoError(t, err)
	err = c.AddProjectMetadata(ctx, strconv.Itoa(int(p.ProjectID)), common.MetadataKeyPreventVul, "true")
	require.NoError(t, err)
	// TODO: Re-introduce this, once https://github.com/goharbor/harbor/pull/15800/ has been adopted.
	// err = c.AddProjectMetadata(ctx, strconv.Itoa(int(p.ProjectID)), common.MetadataKeyReuseSysCVEAllowlist, "true")
	// require.NoError(t, err)
	err = c.AddProjectMetadata(ctx, strconv.Itoa(int(p.ProjectID)), common.MetadataKeySeverity, "medium")
	require.NoError(t, err)

	m, err := c.ListProjectMetadata(ctx, strconv.Itoa(int(p.ProjectID)))
	require.NoError(t, err)

	require.Equal(t, "true", m[common.MetadataKeyAutoScan.String()])
	require.Equal(t, "true", m[common.MetadataKeyEnableContentTrust.String()])
	require.Equal(t, "true", m[common.MetadataKeyPreventVul.String()])
	// require.Equal(t, "true", m[common.MetadataKeyReuseSysCVEAllowlist.String()])
	require.Equal(t, "medium", m[common.MetadataKeySeverity.String()])
}

func TestAPIProjectMetadataUpdate(t *testing.T) {
	projectName := "test-project"

	ctx := context.Background()
	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)

	projectClient := project.NewClient(v2SwaggerClient, defaultOpts, authInfo)

	p, err := projectClient.NewProject(ctx, projectName, &storageLimitPositive)
	defer projectClient.DeleteProject(ctx, p)
	require.NoError(t, err)

	err = c.AddProjectMetadata(ctx, strconv.Itoa(int(p.ProjectID)), common.MetadataKeyAutoScan, "true")
	require.NoError(t, err)

	err = c.UpdateProjectMetadata(ctx, strconv.Itoa(int(p.ProjectID)), common.MetadataKeyAutoScan, "false")

	k, err := c.GetProjectMetadataValue(ctx, strconv.Itoa(int(p.ProjectID)), common.MetadataKeyAutoScan)
	require.NoError(t, err)

	require.Equal(t, "false", k)
}

func TestAPIProjectMetadataDelete(t *testing.T) {
	projectName := "test-project"

	ctx := context.Background()
	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)

	projectClient := project.NewClient(v2SwaggerClient, defaultOpts, authInfo)

	p, err := projectClient.NewProject(ctx, projectName, &storageLimitPositive)
	defer projectClient.DeleteProject(ctx, p)
	require.NoError(t, err)

	err = c.AddProjectMetadata(ctx, strconv.Itoa(int(p.ProjectID)), common.MetadataKeyAutoScan, "true")
	require.NoError(t, err)

	err = c.DeleteProjectMetadataValue(ctx, strconv.Itoa(int(p.ProjectID)), common.MetadataKeyAutoScan)

	require.NoError(t, err)

	m, err := c.GetProjectMetadataValue(ctx, strconv.Itoa(int(p.ProjectID)), common.MetadataKeyAutoScan)

	require.Error(t, err)
	require.Equal(t, "project metadata value is empty: auto_scan", err.Error())
	require.Equal(t, "", m)
}
