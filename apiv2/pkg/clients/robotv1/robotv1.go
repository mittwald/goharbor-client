package robotv1

import (
	"context"

	"github.com/go-openapi/runtime"

	v2client "github.com/testwill/goharbor-client/v5/apiv2/internal/api/client"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/robotv1"
	modelv2 "github.com/testwill/goharbor-client/v5/apiv2/model"
	"github.com/testwill/goharbor-client/v5/apiv2/pkg/config"
)

// RESTClient is a subclient for handling project related actions.
type RESTClient struct {
	// Options contains optional configuration when making API calls.
	Options *config.Options

	// The new client of the harbor v2 API
	V2Client *v2client.Harbor

	// AuthInfo contains the auth information that is provided on API calls.
	AuthInfo runtime.ClientAuthInfoWriter
}

func NewClient(v2Client *v2client.Harbor, opts *config.Options, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		Options:  opts,
		V2Client: v2Client,
		AuthInfo: authInfo,
	}
}

type Client interface {
	ListProjectRobotsV1(ctx context.Context, projectNameOrID string) ([]*modelv2.Robot, error)
	AddProjectRobotV1(ctx context.Context, projectNameOrID string, r *modelv2.RobotCreateV1) error
	UpdateProjectRobotV1(ctx context.Context, projectNameOrID string, robotID int64, r *modelv2.Robot) error
	DeleteProjectRobotV1(ctx context.Context, projectNameOrID string, robotID int64) error
}

// ListProjectRobotsV1 returns a list of all robot accounts in project p.
func (c *RESTClient) ListProjectRobotsV1(ctx context.Context, projectNameOrID string) ([]*modelv2.Robot, error) {
	params := &robotv1.ListRobotV1Params{
		Page:            &c.Options.Page,
		PageSize:        &c.Options.PageSize,
		ProjectNameOrID: projectNameOrID,
		Q:               &c.Options.Query,
		Sort:            &c.Options.Sort,
		Context:         ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.Robotv1.ListRobotV1(params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerRobotV1Errors(err)
	}

	return resp.Payload, nil
}

// AddProjectRobotV1 creates the robot account 'r' and adds it to the project 'p'.
// and returns a 'RobotCreated' response.
func (c *RESTClient) AddProjectRobotV1(ctx context.Context, projectNameOrID string, r *modelv2.RobotCreateV1) error {
	params := &robotv1.CreateRobotV1Params{
		ProjectNameOrID: projectNameOrID,
		Robot:           r,
		Context:         ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Robotv1.CreateRobotV1(params, c.AuthInfo)
	if err != nil {
		return handleSwaggerRobotV1Errors(err)
	}

	return nil
}

// UpdateProjectRobotV1 updates a robot account 'r' in project 'p' using the 'robotID'.
func (c *RESTClient) UpdateProjectRobotV1(ctx context.Context, projectNameOrID string, robotID int64, r *modelv2.Robot) error {
	params := &robotv1.UpdateRobotV1Params{
		ProjectNameOrID: projectNameOrID,
		Robot:           r,
		RobotID:         robotID,
		Context:         ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Robotv1.UpdateRobotV1(params, c.AuthInfo)
	if err != nil {
		return handleSwaggerRobotV1Errors(err)
	}

	return nil
}

// DeleteProjectRobotV1 deletes a robot account from project p.
func (c *RESTClient) DeleteProjectRobotV1(ctx context.Context, projectNameOrID string, robotID int64) error {
	params := &robotv1.DeleteRobotV1Params{
		ProjectNameOrID: projectNameOrID,
		RobotID:         robotID,
		Context:         ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Robotv1.DeleteRobotV1(params, c.AuthInfo)
	if err != nil {
		return handleSwaggerRobotV1Errors(err)
	}

	return nil
}
