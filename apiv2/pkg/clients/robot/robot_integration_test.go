//go:build integration

package robot

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	modelv2 "github.com/mittwald/goharbor-client/v5/apiv2/model"
	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"
)

var testRobotAccountCreate = &modelv2.RobotCreate{
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

func TestAPINewRobotAccount(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	defer c.DeleteRobotAccountByName(ctx, "test-robot")

	robotCreated, err := c.NewRobotAccount(ctx, testRobotAccountCreate)
	require.NoError(t, err)
	require.NotNil(t, robotCreated)

	r, err := c.GetRobotAccountByName(ctx, testRobotAccountCreate.Name)
	require.NoError(t, err)

	require.NotNil(t, r)
}

func TestAPIListRobots(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

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
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	defer c.DeleteRobotAccountByName(ctx, "test-robot")

	_, err := c.NewRobotAccount(ctx, testRobotAccountCreate)
	require.NoError(t, err)

	robot, err := c.GetRobotAccountByName(ctx, "test-robot")
	require.NoError(t, err)
	require.NotNil(t, robot)
}

func TestAPIUpdateRobotAccount(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

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
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

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
