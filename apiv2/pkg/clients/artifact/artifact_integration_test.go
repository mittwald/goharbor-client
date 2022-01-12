//go:build integration

package artifact

import (
	"context"
	"testing"

	"github.com/mittwald/goharbor-client/v5/apiv2/model"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/clients/label"
	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/util"
	"github.com/stretchr/testify/require"
)

var (
	projectName    = "library"
	repositoryName = "image"
	reference      = "test"
)

func TestAPICreateTag(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	tag := &model.Tag{
		Name: "v1.0.0",
	}

	err := c.CreateTag(ctx, projectName, repositoryName, reference, tag)

	defer c.DeleteTag(ctx, projectName, repositoryName, reference, tag.Name)

	require.NoError(t, err)
}

func TestAPIDeleteTag(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	tag := &model.Tag{
		Name: "v1.0.0",
	}

	err := c.CreateTag(ctx, projectName, repositoryName, reference, tag)
	require.NoError(t, err)

	err = c.DeleteTag(ctx, projectName, repositoryName, reference, tag.Name)
	require.NoError(t, err)
}

func TestAPIGetArtifact(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	artifact, err := c.GetArtifact(ctx, projectName, repositoryName, reference)
	require.NoError(t, err)
	require.NotNil(t, artifact)
}

func TestAPIListArtifacts(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	artifacts, err := c.ListArtifacts(ctx, projectName, repositoryName)
	require.NoError(t, err)
	require.NotNil(t, artifacts)
	require.Equal(t, 1, len(artifacts))
	require.Equal(t, "IMAGE", artifacts[0].Type)
	require.Equal(t, "sha256:d04a9a8116092a18007983459ba36800d083a079a5e11087f6b85885ec95c7f7", artifacts[0].Digest)
}

func TestAPIListTags(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	resp, err := c.ListTags(ctx, projectName, repositoryName, reference)
	require.NoError(t, err)
	require.Equal(t, 1, len(resp))
	require.Equal(t, "test", resp[0].Name)
}

func TestAPIAddLabel(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	lc := label.NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := lc.CreateLabel(ctx, &model.Label{
		// Hexadecimal color code
		Color:     "#0072bc",
		Name:      "test",
		Scope:     label.ScopeProject.String(),
		ProjectID: 1,
	})
	require.NoError(t, err)

	labels, err := lc.ListLabels(ctx, "test", util.Int64Ptr(1), label.ScopeProject)
	require.NoError(t, err)
	require.Equal(t, 1, len(labels))

	err = c.AddArtifactLabel(ctx, projectName, repositoryName, reference, labels[0])
	require.NoError(t, err)

	defer lc.DeleteLabel(ctx, labels[0].ID)
}

// TODO: Introduce this, once https://github.com/goharbor/harbor/issues/13468 is resolved.
//func TestAPIGetAddition(t *testing.T) {
//	ctx := context.Background()
//	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)
//
//	artifact, err := c.GetAddition(ctx, projectName, repositoryName, reference, AdditionBuildHistory)
//	fmt.Println(err.Error())
//	require.NoError(t, err)
//
//	fmt.Println(artifact)
//}

//func TestAPIGetVulnerabilitiesAddition(t *testing.T) {
//	ctx := context.Background()
//	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)
//
//	_, err := c.GetVulnerabilitiesAddition(ctx, projectName, repositoryName, reference)
//	fmt.Println(err)
//	require.Error(t, err)
//}
