// +build integration

package robot

import (
	"context"
	"net/url"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/require"

	v2client "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client"
	modelv2 "github.com/mittwald/goharbor-client/v4/apiv2/model"
	integrationtest "github.com/mittwald/goharbor-client/v4/apiv2/testing"
)

var (
	u, _                   = url.Parse(integrationtest.Host)
	v2SwaggerClient        = v2client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	authInfo               = runtimeclient.BasicAuth(integrationtest.User, integrationtest.Password)
	testRobotAccountCreate = &modelv2.RobotCreate{
		Description: "test",
		Disable:     false,
		Duration:    30,
		Level:       LevelSystem.String(),
		Name:        "test-robot",
		Permissions: []*modelv2.RobotPermission{{
			Access: []*modelv2.Access{{
				Action:   ActionPull.String(),
				Resource: ResourceRepository.String(),
			}},
			Kind:      LevelProject.String(),
			Namespace: "library",
		}},
		Secret: "",
	}
)

func TestAPINewRobotAccount(t *testing.T) {
	ctx := context.Background()
	c := NewClient(v2SwaggerClient, authInfo)

	defer c.DeleteRobotAccountByName(ctx, "test-robot")

	r, err := c.NewRobotAccount(ctx, testRobotAccountCreate)

	require.NoError(t, err)
	require.NotNil(t, r)
}

func TestAPIListRobots(t *testing.T) {
	ctx := context.Background()
	c := NewClient(v2SwaggerClient, authInfo)

	defer c.DeleteRobotAccountByName(ctx, "test-robot")

	_, err := c.NewRobotAccount(ctx, testRobotAccountCreate)
	require.NoError(t, err)

	robots, err := c.ListRobotAccounts(ctx)
	require.NoError(t, err)
	require.NotNil(t, robots)

	require.Equal(t, 1, len(robots))
}

func TestAPIGetRobotByName(t *testing.T) {
	ctx := context.Background()
	c := NewClient(v2SwaggerClient, authInfo)

	defer c.DeleteRobotAccountByName(ctx, "test-robot")

	_, err := c.NewRobotAccount(ctx, testRobotAccountCreate)
	require.NoError(t, err)

	robot, err := c.GetRobotAccountByName(ctx, "test-robot")
	require.NoError(t, err)
	require.NotNil(t, robot)
}

func TestAPIUpdateRobotAccount(t *testing.T) {
	ctx := context.Background()
	c := NewClient(v2SwaggerClient, authInfo)

	defer c.DeleteRobotAccountByName(ctx, "test-robot")

	_, err := c.NewRobotAccount(ctx, testRobotAccountCreate)
	require.NoError(t, err)

	r, err := c.GetRobotAccountByName(ctx, "test-robot")
	require.NoError(t, err)
	require.NotNil(t, r)

	err = c.UpdateRobotAccount(ctx, &modelv2.Robot{
		Name:        r.Name,
		Level:       r.Level,
		ID:          r.ID,
		Description: "test-updated",
		Disable:     false,
		Duration:    30,
		Permissions: []*modelv2.RobotPermission{{
			Access: []*modelv2.Access{{
				Action:   ActionPush.String(),
				Resource: ResourceRepository.String(),
			}},
			Namespace: "*",
		}},
	})

	require.NoError(t, err)

	r, err = c.GetRobotAccountByName(ctx, "test-robot")
	require.NoError(t, err)
	require.NotNil(t, r)

	require.Equal(t, "test-updated", r.Description)
}

func TestAPIRefreshRobotAccountSecret(t *testing.T) {
	ctx := context.Background()
	c := NewClient(v2SwaggerClient, authInfo)

	defer c.DeleteRobotAccountByName(ctx, "test-robot")

	_, err := c.NewRobotAccount(ctx, testRobotAccountCreate)
	require.NoError(t, err)

	r, err := c.GetRobotAccountByName(ctx, "test-robot")
	require.NoError(t, err)
	require.NotNil(t, r)

    rSec, err := c.RefreshRobotAccountSecretByName(ctx, "test-robot", "aVeryL0000ngSecret")
	require.NoError(t, err)
	require.NotNil(t, rSec)

	r, err = c.GetRobotAccountByName(ctx, "test-robot")
	require.NoError(t, err)
	require.NotNil(t, r)
}
