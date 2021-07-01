// +build !integration

package robot

import (
	"context"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/robot"
	"github.com/mittwald/goharbor-client/v4/apiv2/mocks"
	modelv2 "github.com/mittwald/goharbor-client/v4/apiv2/model"
)

var (
	authInfo            = runtimeclient.BasicAuth("foo", "bar")
	exampleRobotAccount = &modelv2.Robot{
		Description: "test-robot",
		Disable:     false,
		Duration:    30,
		Editable:    true,
		ID:          1,
		Level:       LevelProject.String(),
		Name:        "robot$test-robot",
		Permissions: []*modelv2.RobotPermission{{
			Access: []*modelv2.Access{
				{
					Action:   ActionPush.String(),
					Resource: ResourceRepository.String(),
				},
				{
					Action:   ActionPull.String(),
					Resource: ResourceRepository.String(),
				},
			},
			Kind:      LevelProject.String(),
			Namespace: "library",
		}},
	}
	exampleRobotCreate = &modelv2.RobotCreate{
		Description: exampleRobotAccount.Description,
		Disable:     exampleRobotAccount.Disable,
		Duration:    exampleRobotAccount.Duration,
		Level:       exampleRobotAccount.Level,
		Name:        exampleRobotAccount.Name,
		Permissions: exampleRobotAccount.Permissions,
		Secret:      exampleRobotAccount.Secret,
	}
	exampleRobotUpdate = &modelv2.Robot{
		Description: "test-updated",
		Disable:     true,
		ID:          exampleRobotAccount.ID,
		Duration:    10,
		Editable:    false,
		Level:       exampleRobotAccount.Level,
		Name:        "robot$test-robot",
		Permissions: []*modelv2.RobotPermission{{
			Access:    []*modelv2.Access{},
			Kind:      exampleRobotAccount.Level,
			Namespace: "library",
		}},
	}
)

func BuildV2ClientWithMock(r *mocks.MockRobotClientService) *client.Harbor {
	return &client.Harbor{
		Robot: r,
	}
}

func TestRESTClient_ListRobotAccounts(t *testing.T) {
	r := &mocks.MockRobotClientService{}

	v2Client := BuildV2ClientWithMock(r)

	cl := NewClient(v2Client, authInfo)

	ctx := context.Background()

	r.On("ListRobot", &robot.ListRobotParams{Context: ctx}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&robot.ListRobotOK{Payload: []*modelv2.Robot{exampleRobotAccount}}, nil)

	robots, err := cl.ListRobotAccounts(ctx)
	require.NoError(t, err)
	require.NotNil(t, robots)
	require.Equal(t, 1, len(robots))

	r.AssertExpectations(t)
}

func TestRESTClient_GetRobotAccountByName(t *testing.T) {
	r := &mocks.MockRobotClientService{}

	v2Client := BuildV2ClientWithMock(r)

	cl := NewClient(v2Client, authInfo)

	ctx := context.Background()

	r.On("ListRobot", &robot.ListRobotParams{Context: ctx}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&robot.ListRobotOK{Payload: []*modelv2.Robot{exampleRobotAccount}}, nil)

	rAcc, err := cl.GetRobotAccountByName(ctx, "test-robot")
	require.NoError(t, err)
	require.NotNil(t, rAcc)

	r.AssertExpectations(t)
}

func TestRESTClient_GetRobotAccountByID(t *testing.T) {
	r := &mocks.MockRobotClientService{}

	v2Client := BuildV2ClientWithMock(r)

	cl := NewClient(v2Client, authInfo)

	ctx := context.Background()

	r.On("GetRobotByID", &robot.GetRobotByIDParams{Context: ctx, RobotID: 1}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&robot.GetRobotByIDOK{Payload: exampleRobotAccount}, nil)

	rAcc, err := cl.GetRobotAccountByID(ctx, 1)
	require.NoError(t, err)
	require.NotNil(t, rAcc)

	r.AssertExpectations(t)
}

func TestRESTClient_NewRobotAccount(t *testing.T) {
	r := &mocks.MockRobotClientService{}

	v2Client := BuildV2ClientWithMock(r)

	cl := NewClient(v2Client, authInfo)

	ctx := context.Background()

	r.On("CreateRobot", &robot.CreateRobotParams{Context: ctx, Robot: exampleRobotCreate},
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&robot.CreateRobotCreated{Payload: &modelv2.RobotCreated{
			ExpiresAt: exampleRobotAccount.ExpiresAt,
			ID:        exampleRobotAccount.ID,
			Name:      exampleRobotAccount.Name,
			Secret:    exampleRobotCreate.Secret,
		}}, nil)

	rCreated, err := cl.NewRobotAccount(ctx, exampleRobotCreate)
	require.NoError(t, err)
	require.NotNil(t, rCreated)

	r.AssertExpectations(t)
}

func TestRESTClient_DeleteRobotAccountByName(t *testing.T) {
	r := &mocks.MockRobotClientService{}

	v2Client := BuildV2ClientWithMock(r)

	cl := NewClient(v2Client, authInfo)

	ctx := context.Background()

	r.On("ListRobot", &robot.ListRobotParams{Context: ctx}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&robot.ListRobotOK{Payload: []*modelv2.Robot{exampleRobotAccount}}, nil)

	r.On("DeleteRobot", &robot.DeleteRobotParams{Context: ctx, RobotID: exampleRobotAccount.ID}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&robot.DeleteRobotOK{}, nil)

	err := cl.DeleteRobotAccountByName(ctx, "test-robot")
	require.NoError(t, err)

	r.AssertExpectations(t)
}

func TestRESTClient_DeleteRobotAccountByID(t *testing.T) {
	r := &mocks.MockRobotClientService{}

	v2Client := BuildV2ClientWithMock(r)

	cl := NewClient(v2Client, authInfo)

	ctx := context.Background()

	r.On("DeleteRobot", &robot.DeleteRobotParams{Context: ctx, RobotID: exampleRobotAccount.ID}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&robot.DeleteRobotOK{}, nil)

	err := cl.DeleteRobotAccountByID(ctx, 1)
	require.NoError(t, err)

	r.AssertExpectations(t)
}

func TestRESTClient_UpdateRobotAccount(t *testing.T) {
	r := &mocks.MockRobotClientService{}

	v2Client := BuildV2ClientWithMock(r)

	cl := NewClient(v2Client, authInfo)

	ctx := context.Background()

	r.On("UpdateRobot", &robot.UpdateRobotParams{Context: ctx, RobotID: exampleRobotAccount.ID, Robot: exampleRobotUpdate}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&robot.UpdateRobotOK{}, nil)

	err := cl.UpdateRobotAccount(ctx, exampleRobotUpdate)
	require.NoError(t, err)

	r.AssertExpectations(t)
}
