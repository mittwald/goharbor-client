//go:build integration

package project

import (
	"context"
	"fmt"
	"testing"

	modelv2 "github.com/testwill/goharbor-client/v5/apiv2/model"
	"github.com/testwill/goharbor-client/v5/apiv2/pkg/clients/registry"
	clienterrors "github.com/testwill/goharbor-client/v5/apiv2/pkg/errors"
	clienttesting "github.com/testwill/goharbor-client/v5/apiv2/pkg/testing"

	"github.com/stretchr/testify/require"
)

var (
	storageLimitPositive int64 = 1
	storageLimitNegative int64 = -1
)

func TestAPIProjectNew(t *testing.T) {
	name := "test-project"

	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := c.NewProject(ctx, &modelv2.ProjectReq{ProjectName: name})
	require.NoError(t, err)

	p, err := c.GetProject(ctx, name)
	require.NoError(t, err)

	defer c.DeleteProject(ctx, name)

	require.Equal(t, name, p.Name)
	require.False(t, p.Deleted)
}

func TestAPINewProxyCacheProject(t *testing.T) {
	ctx := context.Background()

	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)
	rc := registry.NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := rc.NewRegistry(ctx, &modelv2.Registry{
		Insecure: true,
		Name:     "proxy-cache",
		Type:     "harbor",
		URL:      "harbor-registry:5000",
	})

	r, err := rc.GetRegistryByName(ctx, "proxy-cache")
	require.NoError(t, err)

	defer rc.DeleteRegistryByID(ctx, r.ID)

	err = c.NewProject(ctx, &modelv2.ProjectReq{
		ProjectName:  "proxy-cache",
		RegistryID:   &r.ID,
		StorageLimit: &storageLimitPositive,
	})
	require.NoError(t, err)

	p, err := c.GetProject(ctx, "proxy-cache")
	require.NoError(t, err)

	defer c.DeleteProject(ctx, "proxy-cache")

	err = c.UpdateProject(ctx, p, &storageLimitPositive)
	require.NoError(t, err)

	projectAfterUpdate, err := c.GetProject(ctx, "proxy-cache")
	require.NoError(t, err)

	require.Equal(t, r.ID, projectAfterUpdate.RegistryID)
}

func TestAPIProjectNew_UnlimitedStorage(t *testing.T) {
	name := "test-project"

	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := c.NewProject(ctx, &modelv2.ProjectReq{ProjectName: name})
	require.NoError(t, err)

	p, err := c.GetProject(ctx, name)
	require.NoError(t, err)
	defer c.DeleteProject(ctx, name)

	require.Equal(t, name, p.Name)
	require.False(t, p.Deleted)
}

func TestAPIProjectGet(t *testing.T) {
	name := "test-project"

	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := c.NewProject(ctx, &modelv2.ProjectReq{ProjectName: name})
	require.NoError(t, err)

	var pID int32
	var byName, byID *modelv2.Project

	t.Run("ByName", func(t *testing.T) {
		byName, err := c.GetProject(ctx, name)
		require.NoError(t, err)
		require.NoError(t, err)
		pID = byName.ProjectID
	})

	t.Run("ByID", func(t *testing.T) {
		byID, err := c.GetProject(ctx, fmt.Sprint(pID))
		require.NoError(t, err)
		require.NotNil(t, byID)
	})

	require.Equal(t, byName, byID)

	defer c.DeleteProject(ctx, name)
}

func TestAPIProjectDelete(t *testing.T) {
	name := "test-project"
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := c.NewProject(ctx, &modelv2.ProjectReq{ProjectName: name})
	require.NoError(t, err)

	_, err = c.GetProject(ctx, name)
	require.NoError(t, err)

	err = c.DeleteProject(ctx, name)
	require.NoError(t, err)

	_, err = c.GetProject(ctx, name)
	require.Error(t, err)
	require.ErrorIs(t, err, &clienterrors.ErrProjectNotFound{})
}

func TestAPIProjectList(t *testing.T) {
	namePrefix := "test-project"
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("%s-%d", namePrefix, i)
		err := c.NewProject(ctx, &modelv2.ProjectReq{ProjectName: name})
		require.NoError(t, err)

		_, err = c.GetProject(ctx, name)
		require.NoError(t, err)
		require.NoError(t, err)
		defer func() {
			defer c.DeleteProject(ctx, name)
			if err != nil {
				panic("error in cleanup routine: " + err.Error())
			}
		}()
	}

	projects, err := c.ListProjects(ctx, namePrefix)
	require.NoError(t, err)
	require.Len(t, projects, 10)
	for _, v := range projects {
		require.Contains(t, v.Name, namePrefix)
	}
}

func TestAPIProjectUpdate(t *testing.T) {
	name := "test-project"
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := c.NewProject(ctx, &modelv2.ProjectReq{ProjectName: name})
	require.NoError(t, err)

	p, err := c.GetProject(ctx, name)
	require.NoError(t, err)
	require.NoError(t, err)

	defer c.DeleteProject(ctx, name)

	require.Equal(t, false, p.Togglable)

	p.Togglable = true

	err = c.UpdateProject(ctx, p, &storageLimitPositive)
	require.NoError(t, err)
	p2, err := c.GetProject(ctx, name)
	require.NoError(t, err)

	require.NotEqual(t, p, p2)
}
