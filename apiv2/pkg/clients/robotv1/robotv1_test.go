//go:build !integration

package robotv1

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/robotv1"
	"github.com/mittwald/goharbor-client/v5/apiv2/mocks"
	modelv2 "github.com/mittwald/goharbor-client/v5/apiv2/model"
	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"
)

var (
	ctx                    = context.Background()
	exampleProjectID       = 1
	exampleRobotID   int64 = 42
)

func APIandMockClientsForTests() (*RESTClient, *clienttesting.MockClients) {
	desiredMockClients := &clienttesting.MockClients{
		Robotv1: mocks.MockRobotv1ClientService{},
	}

	v2Client := clienttesting.BuildV2ClientWithMocks(desiredMockClients)

	cl := NewClient(v2Client, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	return cl, desiredMockClients
}

func TestRESTClient_ListProjectRobotsV1(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	expectedRobot := &modelv2.Robot{
		Description: "some robot account",
		Disable:     false,
		ID:          42,
		Name:        "robot$account",
	}

	params := &robotv1.ListRobotV1Params{
		Page:            &apiClient.Options.Page,
		PageSize:        &apiClient.Options.PageSize,
		ProjectNameOrID: strconv.Itoa(exampleProjectID),
		Q:               &apiClient.Options.Query,
		Sort:            &apiClient.Options.Sort,
		Context:         ctx,
	}

	params.WithTimeout(apiClient.Options.Timeout)

	mockClient.Robotv1.On("ListRobotV1", params, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&robotv1.ListRobotV1OK{Payload: []*modelv2.Robot{expectedRobot}}, nil)

	robots, err := apiClient.ListProjectRobotsV1(ctx, strconv.Itoa(exampleProjectID))

	assert.NoError(t, err)

	assert.Equal(t, expectedRobot, robots[0])

	mockClient.Robotv1.AssertExpectations(t)
}

func TestRESTClient_AddProjectRobotV1(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	newRobot := &modelv2.RobotCreateV1{
		Access: []*modelv2.Access{{
			Action:   "push",
			Effect:   "",
			Resource: fmt.Sprintf("/project/%d/repository", exampleProjectID),
		}},
		Name: "test-robot",
	}

	params := &robotv1.CreateRobotV1Params{
		ProjectNameOrID: strconv.Itoa(int(exampleProjectID)),
		Robot:           newRobot,
		Context:         ctx,
	}

	params.WithTimeout(apiClient.Options.Timeout)

	expectedPayload := &modelv2.RobotCreated{
		Name: "test-robot",
	}

	mockClient.Robotv1.On("CreateRobotV1", params, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&robotv1.CreateRobotV1Created{Payload: expectedPayload}, nil)

	err := apiClient.AddProjectRobotV1(ctx, strconv.Itoa(exampleProjectID), newRobot)

	assert.NoError(t, err)

	mockClient.Robotv1.AssertExpectations(t)
}

func TestRESTClient_UpdateProjectRobotV1(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	updateRobot := &modelv2.Robot{
		CreationTime: strfmt.DateTime{},
		Disable:      false,
		Editable:     true,
		ID:           exampleRobotID,
		Name:         "test-robot",
	}

	params := &robotv1.UpdateRobotV1Params{
		ProjectNameOrID: strconv.Itoa(int(exampleProjectID)),
		Robot:           updateRobot,
		RobotID:         exampleRobotID,
		Context:         ctx,
	}

	params.WithTimeout(apiClient.Options.Timeout)

	mockClient.Robotv1.On("UpdateRobotV1", params, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&robotv1.UpdateRobotV1OK{}, nil)

	err := apiClient.UpdateProjectRobotV1(ctx, strconv.Itoa(exampleProjectID), exampleRobotID, updateRobot)

	assert.NoError(t, err)

	mockClient.Robotv1.AssertExpectations(t)
}

func TestRESTClient_DeleteProjectRobotV1(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	params := &robotv1.DeleteRobotV1Params{
		ProjectNameOrID: strconv.Itoa(int(exampleProjectID)),
		RobotID:         exampleRobotID,
		Context:         ctx,
	}

	params.WithTimeout(apiClient.Options.Timeout)

	mockClient.Robotv1.On("DeleteRobotV1", params, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&robotv1.DeleteRobotV1OK{}, nil)

	err := apiClient.DeleteProjectRobotV1(ctx, strconv.Itoa(exampleProjectID), exampleRobotID)

	assert.NoError(t, err)

	mockClient.Robotv1.AssertExpectations(t)
}
