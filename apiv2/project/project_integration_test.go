//go:build integration

package project

import (
	"context"
	"fmt"
	"net/url"
	"testing"

	"github.com/go-openapi/strfmt"

	clienterrors "github.com/mittwald/goharbor-client/v4/apiv2/pkg/errors"

	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/config"
	integrationtest "github.com/mittwald/goharbor-client/v4/apiv2/pkg/testing"

	runtimeclient "github.com/go-openapi/runtime/client"

	v2client "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client"

	"github.com/stretchr/testify/require"
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

func TestAPIProjectNew(t *testing.T) {
	name := "test-project"

	ctx := context.Background()
	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)

	p, err := c.NewProject(ctx, name, &storageLimitPositive)
	require.NoError(t, err)
	err = c.DeleteProject(ctx, p)

	require.NoError(t, err)
	require.Equal(t, name, p.Name)
	require.False(t, p.Deleted)
}

func TestAPIProjectNew_UnlimitedStorage(t *testing.T) {
	name := "test-project"

	ctx := context.Background()
	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)

	p, err := c.NewProject(ctx, name, &storageLimitNegative)
	defer c.DeleteProject(ctx, p)

	require.NoError(t, err)
	require.Equal(t, name, p.Name)
	require.False(t, p.Deleted)
}

func TestAPIProjectGet(t *testing.T) {
	name := "test-project"

	ctx := context.Background()
	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)

	p, err := c.NewProject(ctx, name, &storageLimitPositive)
	require.NoError(t, err)
	defer c.DeleteProject(ctx, p)

	p2, err := c.GetProject(ctx, name)
	require.NoError(t, err)
	require.Equal(t, p, p2)
}

func TestAPIProjectDelete(t *testing.T) {
	name := "test-project"
	ctx := context.Background()
	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)

	p, err := c.NewProject(ctx, name, &storageLimitPositive)
	require.NoError(t, err)

	err = c.DeleteProject(ctx, p)
	require.NoError(t, err)

	p, err = c.GetProject(ctx, name)
	require.Error(t, err)
	require.ErrorIs(t, err, &clienterrors.ErrProjectNotFound{})
}

func TestAPIProjectList(t *testing.T) {
	namePrefix := "test-project"
	ctx := context.Background()
	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)

	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("%s-%d", namePrefix, i)
		p, err := c.NewProject(ctx, name, &storageLimitPositive)
		require.NoError(t, err)
		defer func() {
			err := c.DeleteProject(ctx, p)
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
	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)

	p, err := c.NewProject(ctx, name, &storageLimitPositive)
	require.NoError(t, err)

	defer c.DeleteProject(ctx, p)

	require.Equal(t, false, p.Togglable)

	p.Togglable = true

	err = c.UpdateProject(ctx, p, &storageLimitPositive)
	require.NoError(t, err)
	p2, err := c.GetProject(ctx, name)
	require.NoError(t, err)

	require.NotEqual(t, p, p2)
}
