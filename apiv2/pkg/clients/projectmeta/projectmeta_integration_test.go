//go:build integration

package projectmeta

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/clients/project"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/common"

	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"

	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/errors"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/util"
)

var (
	storageLimitPositive int64 = 1
	storageLimitNegative int64 = -1
)

func TestAPIProjectMetadataAdd(t *testing.T) {
	projectName := "test-project"

	ctx := context.Background()
	metaClient := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	projectClient := project.NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := projectClient.NewProject(ctx, projectName, &storageLimitPositive)
	require.NoError(t, err)

	p, err := projectClient.GetProject(ctx, projectName)
	require.NoError(t, err)
	defer projectClient.DeleteProject(ctx, projectName)

	err = metaClient.AddProjectMetadata(ctx, util.ProjectIDAsString(p.ProjectID), common.ProjectMetadataKeyAutoScan, "true")
	require.NoError(t, err)
	err = metaClient.AddProjectMetadata(ctx, util.ProjectIDAsString(p.ProjectID), common.ProjectMetadataKeyEnableContentTrust, "true")
	require.NoError(t, err)
	err = metaClient.AddProjectMetadata(ctx, util.ProjectIDAsString(p.ProjectID), common.ProjectMetadataKeyPreventVul, "true")
	require.NoError(t, err)
	// TODO: Re-introduce this, once https://github.com/goharbor/harbor/pull/15800/ has been added.
	// err = metaClient.AddProjectMetadata(ctx, util.ProjectIDAsString(p.ProjectID), common.MetadataKeyReuseSysCVEAllowlist, "true")
	// require.NoError(t, err)
	err = metaClient.AddProjectMetadata(ctx, util.ProjectIDAsString(p.ProjectID), common.ProjectMetadataKeySeverity, "medium")
	require.NoError(t, err)
}

func TestAPIProjectMetadataGet(t *testing.T) {
	projectName := "test-project"

	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	projectClient := project.NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := projectClient.NewProject(ctx, projectName, &storageLimitPositive)
	require.NoError(t, err)

	p, err := projectClient.GetProject(ctx, projectName)
	require.NoError(t, err)
	defer projectClient.DeleteProject(ctx, projectName)

	m, err := c.GetProjectMetadataValue(ctx, strconv.Itoa(int(p.ProjectID)), common.ProjectMetadataKeyPublic)
	require.NoError(t, err)

	require.Equal(t, "false", m)
}

func TestAPIProjectMetadataGetInvalidKey(t *testing.T) {
	projectName := "test-project"

	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	projectClient := project.NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := projectClient.NewProject(ctx, projectName, &storageLimitPositive)
	require.NoError(t, err)

	p, err := projectClient.GetProject(ctx, projectName)
	require.NoError(t, err)
	defer projectClient.DeleteProject(ctx, projectName)

	m, err := c.GetProjectMetadataValue(ctx, strconv.Itoa(int(p.ProjectID)), "foobar")

	require.Error(t, err)
	require.Equal(t, "invalid request", err.Error())
	require.IsType(t, &errors.ErrProjectInvalidRequest{}, err)
	require.Equal(t, "", m)
}

func TestAPIProjectMetadataList(t *testing.T) {
	projectName := "test-project"

	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	projectClient := project.NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)
	err := projectClient.NewProject(ctx, projectName, &storageLimitPositive)
	require.NoError(t, err)

	p, err := projectClient.GetProject(ctx, projectName)
	require.NoError(t, err)
	defer projectClient.DeleteProject(ctx, projectName)

	err = c.AddProjectMetadata(ctx, strconv.Itoa(int(p.ProjectID)), common.ProjectMetadataKeyAutoScan, "true")
	require.NoError(t, err)
	err = c.AddProjectMetadata(ctx, strconv.Itoa(int(p.ProjectID)), common.ProjectMetadataKeyEnableContentTrust, "true")
	require.NoError(t, err)
	err = c.AddProjectMetadata(ctx, strconv.Itoa(int(p.ProjectID)), common.ProjectMetadataKeyPreventVul, "true")
	require.NoError(t, err)
	// TODO: Re-introduce this, once https://github.com/goharbor/harbor/pull/15800/ has been adopted.
	// err = c.AddProjectMetadata(ctx, strconv.Itoa(int(p.ProjectID)), common.MetadataKeyReuseSysCVEAllowlist, "true")
	// require.NoError(t, err)
	err = c.AddProjectMetadata(ctx, strconv.Itoa(int(p.ProjectID)), common.ProjectMetadataKeySeverity, "medium")
	require.NoError(t, err)

	m, err := c.ListProjectMetadata(ctx, strconv.Itoa(int(p.ProjectID)))
	require.NoError(t, err)

	require.Equal(t, "true", m[common.ProjectMetadataKeyAutoScan.String()])
	require.Equal(t, "true", m[common.ProjectMetadataKeyEnableContentTrust.String()])
	require.Equal(t, "true", m[common.ProjectMetadataKeyPreventVul.String()])
	// require.Equal(t, "true", m[common.MetadataKeyReuseSysCVEAllowlist.String()])
	require.Equal(t, "medium", m[common.ProjectMetadataKeySeverity.String()])
}

func TestAPIProjectMetadataUpdate(t *testing.T) {
	projectName := "test-project"

	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	projectClient := project.NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := projectClient.NewProject(ctx, projectName, &storageLimitPositive)
	require.NoError(t, err)

	p, err := projectClient.GetProject(ctx, projectName)
	require.NoError(t, err)
	defer projectClient.DeleteProject(ctx, projectName)

	err = c.AddProjectMetadata(ctx, strconv.Itoa(int(p.ProjectID)), common.ProjectMetadataKeyAutoScan, "true")
	require.NoError(t, err)

	err = c.UpdateProjectMetadata(ctx, strconv.Itoa(int(p.ProjectID)), common.ProjectMetadataKeyAutoScan, "false")

	k, err := c.GetProjectMetadataValue(ctx, strconv.Itoa(int(p.ProjectID)), common.ProjectMetadataKeyAutoScan)
	require.NoError(t, err)

	require.Equal(t, "false", k)
}

func TestAPIProjectMetadataDelete(t *testing.T) {
	projectName := "test-project"

	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	projectClient := project.NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := projectClient.NewProject(ctx, projectName, &storageLimitPositive)
	require.NoError(t, err)

	p, err := projectClient.GetProject(ctx, projectName)
	require.NoError(t, err)
	defer projectClient.DeleteProject(ctx, projectName)

	err = c.AddProjectMetadata(ctx, strconv.Itoa(int(p.ProjectID)), common.ProjectMetadataKeyAutoScan, "true")
	require.NoError(t, err)

	err = c.DeleteProjectMetadataValue(ctx, strconv.Itoa(int(p.ProjectID)), common.ProjectMetadataKeyAutoScan)

	require.NoError(t, err)

	m, err := c.GetProjectMetadataValue(ctx, strconv.Itoa(int(p.ProjectID)), common.ProjectMetadataKeyAutoScan)

	require.Error(t, err)
	require.Equal(t, "project metadata value is empty: auto_scan", err.Error())
	require.Equal(t, "", m)
}
