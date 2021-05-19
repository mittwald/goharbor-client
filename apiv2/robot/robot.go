package robot

import (
	"context"

	"github.com/go-openapi/runtime"
	v2client "github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client/robot"
	modelv2 "github.com/mittwald/goharbor-client/v3/apiv2/model"
)

const (
	// Definitions in this block may be used to interact with the package methods.

	// LevelProject defines a project-wide access level for a robot account.
	LevelProject Level = "project"
	// LevelSystem defines a system-wide access level for a robot account.
	LevelSystem Level = "system"

	ResourceRepository       AccessResource = "repository"
	ResourceArtifact         AccessResource = "artifact"
	ResourceHelmChart        AccessResource = "helm-chart"
	ResourceHelmChartVersion AccessResource = "helm-chart-version"
	ResourceTag              AccessResource = "tag"
	ResourceArtifactLabel    AccessResource = "artifact-label"
	ResourceScan             AccessResource = "scan"

	ActionPush   AccessAction = "push"
	ActionPull   AccessAction = "pull"
	ActionCreate AccessAction = "create"
	ActionDelete AccessAction = "delete"
	ActionRead   AccessAction = "read"
)

// RESTClient is a subclient for handling project related actions.
type RESTClient struct {
	// The new client of the harbor v2 API
	V2Client *v2client.Harbor

	// AuthInfo contains the auth information that is provided on API calls.
	AuthInfo runtime.ClientAuthInfoWriter
}

func NewClient(v2Client *v2client.Harbor, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		V2Client: v2Client,
		AuthInfo: authInfo,
	}
}

type Client interface {
	ListRobotAccounts(ctx context.Context) ([]*modelv2.Robot, error)
	GetRobotAccountByName(ctx context.Context, name string) (*modelv2.Robot, error)
	GetRobotAccountByID(ctx context.Context, id int64) (*modelv2.Robot, error)
	NewRobotAccount(ctx context.Context, r *modelv2.RobotCreate) (*modelv2.RobotCreated, error)
	DeleteRobotAccountByName(ctx context.Context, name string) error
	DeleteRobotAccountByID(ctx context.Context, id int64) error
	UpdateRobotAccount(ctx context.Context, r *modelv2.Robot) error
}

type Level string

func (in Level) String() string {
	return string(in)
}

type AccessResource string

func (in AccessResource) String() string {
	return string(in)
}

type AccessAction string

func (in AccessAction) String() string {
	return string(in)
}

// ListProjectRobots returns a list of all robot accounts.
func (c *RESTClient) ListRobotAccounts(ctx context.Context) ([]*modelv2.Robot, error) {
	resp, err := c.V2Client.Robot.ListRobot(&robot.ListRobotParams{
		Context: ctx,
	}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerRobotErrors(err)
	}

	return resp.Payload, nil
}

// GetRobotByName lists all existing robot accounts and returns the one matching the provided name.
// Note that the generic 'robot$'-prefix of the robot name is implicitly used for getting the resource.
func (c *RESTClient) GetRobotAccountByName(ctx context.Context, name string) (*modelv2.Robot, error) {
	robots, err := c.ListRobotAccounts(ctx)
	if err != nil {
		return nil, err
	}

	for _, r := range robots {
		if r.Name == "robot$"+name {
			return r, nil
		}
	}

	return nil, &ErrRobotAccountUnknownResource{}
}

// GetRobotByID returns a robot account identified by its 'id'.
func (c *RESTClient) GetRobotAccountByID(ctx context.Context, id int64) (*modelv2.Robot, error) {
	resp, err := c.V2Client.Robot.GetRobotByID(&robot.GetRobotByIDParams{
		RobotID: id,
		Context: ctx,
	}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerRobotErrors(err)
	}

	return resp.Payload, nil
}

// NewRobotAccount creates a new robot account from the specification of 'r' and returns a 'RobotCreated' response.
func (c *RESTClient) NewRobotAccount(ctx context.Context, r *modelv2.RobotCreate) (*modelv2.RobotCreated, error) {
	resp, err := c.V2Client.Robot.CreateRobot(&robot.CreateRobotParams{
		Robot:   r,
		Context: ctx,
	}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerRobotErrors(err)
	}

	return resp.Payload, nil
}

// DeleteRobotAccountByName deletes a robot account identified by its 'name'.
// Note that the generic 'robot$'-prefix of the robot name is implicitly used for deletion.
func (c *RESTClient) DeleteRobotAccountByName(ctx context.Context, name string) error {
	robots, err := c.ListRobotAccounts(ctx)
	if err != nil {
		return err
	}

	for _, r := range robots {
		if r.Name == "robot$"+name {
			return c.DeleteRobotAccountByID(ctx, r.ID)
		}
	}

	return &ErrRobotAccountUnknownResource{}
}

// DeleteProjectRobotByID deletes a robot account identified by its id.
func (c *RESTClient) DeleteRobotAccountByID(ctx context.Context, id int64) error {
	_, err := c.V2Client.Robot.DeleteRobot(&robot.DeleteRobotParams{
		RobotID: id,
		Context: ctx,
	}, c.AuthInfo)
	if err != nil {
		return handleSwaggerRobotErrors(err)
	}

	return nil
}

// UpdateRobotAccount updates the robot account 'r' with the provided specification.
// Note that modelv2.Robot.Name & modelv2.Robot.Level are immutable by API definitions.
func (c *RESTClient) UpdateRobotAccount(ctx context.Context, r *modelv2.Robot) error {
	_, err := c.V2Client.Robot.UpdateRobot(&robot.UpdateRobotParams{
		Robot:   r,
		RobotID: r.ID,
		Context: ctx,
	}, c.AuthInfo)
	if err != nil {
		return handleSwaggerRobotErrors(err)
	}

	return nil
}
