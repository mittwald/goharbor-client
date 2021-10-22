package robotv1

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/robotv1"
	"github.com/mittwald/goharbor-client/v4/apiv2/mocks"
	modelv2 "github.com/mittwald/goharbor-client/v4/apiv2/model"
)

func TestRESTClient_ListProjectRobots(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	r := &mocks.MockRobotv1ClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(nil, r)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	expectedRobot := &modelv2.Robot{
		Description: "some robot account",
		Disable:     false,
		ID:          42,
		Name:        "robot$account",
	}

	params := &robotv1.ListRobotV1Params{
		ProjectNameOrID: strconv.Itoa(int(exampleProjectID)),
		Context:         ctx,
	}

	r.On("ListRobotV1", params, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&robotv1.ListRobotV1OK{Payload: []*modelv2.Robot{expectedRobot}}, nil)

	robots, err := cl.ListProjectRobots(ctx, exampleProject)

	assert.NoError(t, err)

	assert.Equal(t, expectedRobot, robots[0])

	r.AssertExpectations(t)
}

func TestRESTClient_AddProjectRobot(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	r := &mocks.MockRobotv1ClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(nil, r)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

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

	expectedPayload := &modelv2.RobotCreated{
		Name: "test-robot",
	}

	r.On("CreateRobotV1", params, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&robotv1.CreateRobotV1Created{Payload: expectedPayload}, nil)

	createdRobot, err := cl.AddProjectRobot(ctx, exampleProject, newRobot)

	assert.NoError(t, err)

	assert.NotNil(t, createdRobot)

	r.AssertExpectations(t)
}

func TestRESTClient_UpdateProjectRobot(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	r := &mocks.MockRobotv1ClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(nil, r)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	const exampleRobotID = 42

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

	r.On("UpdateRobotV1", params, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&robotv1.UpdateRobotV1OK{}, nil)

	err := cl.UpdateProjectRobot(ctx, exampleProject, exampleRobotID, updateRobot)

	assert.NoError(t, err)

	r.AssertExpectations(t)
}

func TestRESTClient_DeleteProjectRobot(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	r := &mocks.MockRobotv1ClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(nil, r)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	const exampleRobotID = 42

	params := &robotv1.DeleteRobotV1Params{
		ProjectNameOrID: strconv.Itoa(int(exampleProjectID)),
		RobotID:         exampleRobotID,
		Context:         ctx,
	}

	r.On("DeleteRobotV1", params, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&robotv1.DeleteRobotV1OK{}, nil)

	err := cl.DeleteProjectRobot(ctx, exampleProject, exampleRobotID)

	assert.NoError(t, err)

	r.AssertExpectations(t)
}
