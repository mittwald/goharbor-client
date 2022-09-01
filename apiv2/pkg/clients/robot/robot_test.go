//go:build !integration

package robot

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/robot"
	"github.com/mittwald/goharbor-client/v5/apiv2/mocks"
	modelv2 "github.com/mittwald/goharbor-client/v5/apiv2/model"
	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"
)

var (
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
	exampleSec = "aVeryL0000ngSecret"
	ctx        = context.Background()
)

func APIandMockClientsForTests() (*RESTClient, *clienttesting.MockClients) {
	desiredMockClients := &clienttesting.MockClients{
		Robot: mocks.MockRobotClientService{},
	}

	v2Client := clienttesting.BuildV2ClientWithMocks(desiredMockClients)

	cl := NewClient(v2Client, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	return cl, desiredMockClients
}

func TestRESTClient_ListRobotAccounts(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	listRobotParams := &robot.ListRobotParams{
		Page:     &apiClient.Options.Page,
		PageSize: &apiClient.Options.PageSize,
		Q:        &apiClient.Options.Query,
		Sort:     &apiClient.Options.Sort,
		Context:  ctx,
	}
	listRobotParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Robot.On("ListRobot", listRobotParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&robot.ListRobotOK{Payload: []*modelv2.Robot{exampleRobotAccount}}, nil)

	robots, err := apiClient.ListRobotAccounts(ctx)
	require.NoError(t, err)
	require.NotNil(t, robots)
	require.Equal(t, 1, len(robots))

	mockClient.Robot.AssertExpectations(t)
}

func TestRESTClient_GetRobotAccountByName(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	listRobotParams := &robot.ListRobotParams{
		Page:     &apiClient.Options.Page,
		PageSize: &apiClient.Options.PageSize,
		Q:        &apiClient.Options.Query,
		Sort:     &apiClient.Options.Sort,
		Context:  ctx,
	}
	listRobotParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Robot.On("ListRobot", listRobotParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&robot.ListRobotOK{Payload: []*modelv2.Robot{exampleRobotAccount}}, nil)

	rAcc, err := apiClient.GetRobotAccountByName(ctx, "test-robot")
	require.NoError(t, err)
	require.NotNil(t, rAcc)

	mockClient.Robot.AssertExpectations(t)
}

func TestRESTClient_GetRobotAccountByID(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getRobotParams := &robot.GetRobotByIDParams{
		RobotID: 1,
		Context: ctx,
	}
	getRobotParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Robot.On("GetRobotByID", getRobotParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&robot.GetRobotByIDOK{Payload: exampleRobotAccount}, nil)

	rAcc, err := apiClient.GetRobotAccountByID(ctx, 1)
	require.NoError(t, err)
	require.NotNil(t, rAcc)

	mockClient.Robot.AssertExpectations(t)
}

func TestRESTClient_NewRobotAccount(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	createRobotParams := &robot.CreateRobotParams{Context: ctx, Robot: exampleRobotCreate}
	createRobotParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Robot.On("CreateRobot", createRobotParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&robot.CreateRobotCreated{Payload: &modelv2.RobotCreated{
			ExpiresAt: exampleRobotAccount.ExpiresAt,
			ID:        exampleRobotAccount.ID,
			Name:      exampleRobotAccount.Name,
			Secret:    exampleRobotCreate.Secret,
		}}, nil)

	_, err := apiClient.NewRobotAccount(ctx, exampleRobotCreate)
	require.NoError(t, err)

	mockClient.Robot.AssertExpectations(t)
}

func TestRESTClient_DeleteRobotAccountByName(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	listRobotParams := &robot.ListRobotParams{
		Page:     &apiClient.Options.Page,
		PageSize: &apiClient.Options.PageSize,
		Q:        &apiClient.Options.Query,
		Sort:     &apiClient.Options.Sort,
		Context:  ctx,
	}
	listRobotParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Robot.On("ListRobot", listRobotParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&robot.ListRobotOK{Payload: []*modelv2.Robot{exampleRobotAccount}}, nil)

	deleteRobotParams := &robot.DeleteRobotParams{
		Context: ctx,
		RobotID: exampleRobotAccount.ID,
	}
	deleteRobotParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Robot.On("DeleteRobot", deleteRobotParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&robot.DeleteRobotOK{}, nil)

	err := apiClient.DeleteRobotAccountByName(ctx, "test-robot")
	require.NoError(t, err)

	mockClient.Robot.AssertExpectations(t)
}

func TestRESTClient_DeleteRobotAccountByID(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	deleteRobotParams := &robot.DeleteRobotParams{Context: ctx, RobotID: exampleRobotAccount.ID}
	deleteRobotParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Robot.On("DeleteRobot", deleteRobotParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&robot.DeleteRobotOK{}, nil)

	err := apiClient.DeleteRobotAccountByID(ctx, 1)
	require.NoError(t, err)

	mockClient.Robot.AssertExpectations(t)
}

func TestRESTClient_UpdateRobotAccount(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	updateRobotParams := &robot.UpdateRobotParams{Context: ctx, RobotID: exampleRobotAccount.ID, Robot: exampleRobotUpdate}
	updateRobotParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Robot.On("UpdateRobot", updateRobotParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&robot.UpdateRobotOK{}, nil)

	err := apiClient.UpdateRobotAccount(ctx, exampleRobotUpdate)
	require.NoError(t, err)

	mockClient.Robot.AssertExpectations(t)
}

func TestRESTClient_RefreshRobotAccountSecretByID(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	refreshSecParams := &robot.RefreshSecParams{
		Context: ctx,
		RobotID: exampleRobotAccount.ID,
		RobotSec: &modelv2.RobotSec{
			Secret: exampleSec,
		},
	}
	refreshSecParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Robot.On("RefreshSec", refreshSecParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&robot.RefreshSecOK{Payload: &modelv2.RobotSec{Secret: exampleSec}}, nil)

	rSec, err := apiClient.RefreshRobotAccountSecretByID(ctx, 1, exampleSec)

	require.NoError(t, err)
	require.NotNil(t, rSec)

	mockClient.Robot.AssertExpectations(t)
}

func TestRESTClient_RefreshRobotAccountSecretByName(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	listRobotParams := &robot.ListRobotParams{
		Page:     &apiClient.Options.Page,
		PageSize: &apiClient.Options.PageSize,
		Q:        &apiClient.Options.Query,
		Sort:     &apiClient.Options.Sort,
		Context:  ctx,
	}
	listRobotParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Robot.On("ListRobot", listRobotParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&robot.ListRobotOK{Payload: []*modelv2.Robot{exampleRobotAccount}}, nil)

	refreshSecParams := &robot.RefreshSecParams{
		Context: ctx,
		RobotID: exampleRobotAccount.ID,
		RobotSec: &modelv2.RobotSec{
			Secret: exampleSec,
		},
	}
	refreshSecParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Robot.On("RefreshSec", refreshSecParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&robot.RefreshSecOK{Payload: &modelv2.RobotSec{Secret: exampleSec}}, nil)

	rSec, err := apiClient.RefreshRobotAccountSecretByName(ctx, "test-robot", exampleSec)

	require.NoError(t, err)
	require.NotNil(t, rSec)

	mockClient.Robot.AssertExpectations(t)
}
