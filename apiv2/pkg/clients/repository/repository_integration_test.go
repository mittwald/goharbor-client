//go:build integration

package repository

import (
	"context"
	"testing"

	"github.com/testwill/goharbor-client/v5/apiv2/model"
	"github.com/testwill/goharbor-client/v5/apiv2/pkg/clients/project"
	clienttesting "github.com/testwill/goharbor-client/v5/apiv2/pkg/testing"
	"github.com/stretchr/testify/require"
)

var projectName = "test-project"

func TestAPIRepositoryListAllRepositories(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	repositories, err := c.ListAllRepositories(ctx)
	require.NoError(t, err)
	// A repository should exist, as the Makefile target `upload-test-image` pushes a test image.
	require.NotEmpty(t, repositories)
}

func TestAPIRepositoryListRepositories(t *testing.T) {
	ctx := context.Background()

	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)
	pc := project.NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := pc.NewProject(ctx, &model.ProjectReq{
		ProjectName: projectName,
	})
	require.NoError(t, err)

	defer pc.DeleteProject(ctx, projectName)

	projectRepositories, err := c.ListRepositories(ctx, projectName)
	require.NoError(t, err)
	require.Empty(t, projectRepositories)
}
