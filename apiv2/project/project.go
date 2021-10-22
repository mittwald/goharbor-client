package project

import (
	"context"
	"errors"

	projectapi "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/project"
	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/common"
	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/config"
	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/util"

	v2client "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client"
	modelv2 "github.com/mittwald/goharbor-client/v4/apiv2/model"

	"github.com/go-openapi/runtime"
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
	NewProject(ctx context.Context, name string, storageLimit *int64) (*modelv2.Project, error)
	DeleteProject(ctx context.Context, p *modelv2.Project) error
	GetProject(ctx context.Context, nameOrID string) (*modelv2.Project, error)
	ListProjects(ctx context.Context, nameFilter string) ([]*modelv2.Project, error)
	UpdateProject(ctx context.Context, p *modelv2.Project, storageLimit *int64) error
	ProjectExists(ctx context.Context, projectNameOrID string) (bool, error)
}

// NewProject creates a new project with name as the project's name.
// Returns the project as it is stored inside Harbor or an error,
// if the project could not be created.
// CountLimit limits the number of repositories for this project.
// StorageLimit limits the allocatable space for this project.
func (c *RESTClient) NewProject(ctx context.Context, name string, storageLimit *int64) (*modelv2.Project, error) {
	pReq := &modelv2.ProjectReq{
		ProjectName:  name,
		StorageLimit: storageLimit,
	}

	params := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Project.CreateProject(params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerProjectErrors(err)
	}

	project, err := c.GetProject(ctx, name)
	if err != nil {
		return nil, err
	}

	return project, nil
}

// DeleteProject deletes the specified project.
// Returns an error when no matching project is found or when
// having difficulties talking to the API.
func (c *RESTClient) DeleteProject(ctx context.Context, p *modelv2.Project) error {
	if p == nil {
		return &common.ErrProjectNotProvided{}
	}

	projectExists, err := c.ProjectExists(ctx, p.Name)
	if err != nil {
		return err
	}

	if !projectExists {
		return &common.ErrProjectMismatch{}
	}

	params := &projectapi.DeleteProjectParams{
		ProjectNameOrID: util.ProjectIDAsString(p.ProjectID),
		Context:         ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err = c.V2Client.Project.DeleteProject(params, c.AuthInfo)

	return handleSwaggerProjectErrors(err)
}

// GetProject returns an existing project identified by nameOrID.
// nameOrID may contain a unique project name or its unique ID.
// Returns an error if it cannot find a matching project or when
// having difficulties talking to the API.
func (c *RESTClient) GetProject(ctx context.Context, nameOrID string) (*modelv2.Project, error) {
	if nameOrID == "" {
		return nil, &common.ErrProjectNameNotProvided{}
	}

	params := &projectapi.GetProjectParams{
		ProjectNameOrID: nameOrID,
		Context:         ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.Project.GetProject(params, c.AuthInfo)
	if err != nil {
		if resp == nil {
			return nil, &common.ErrProjectNotFound{}
		}
		return nil, handleSwaggerProjectErrors(err)
	}

	return resp.Payload, nil
}

// ListProjects returns a list of projects based on a name filter.
// Returns all projects if name is an empty string.
// Returns an error if no projects were found.
func (c *RESTClient) ListProjects(ctx context.Context, nameFilter string) ([]*modelv2.Project, error) {
	params := &projectapi.ListProjectsParams{
		Name:     &nameFilter,
		PageSize: &c.Options.PageSize,
		Q:        &c.Options.Query,
		Sort:     &c.Options.Sort,
		Context:  ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.Project.ListProjects(params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerProjectErrors(err)
	}

	if len(resp.Payload) == 0 {
		return nil, &common.ErrProjectNotFound{}
	}

	return resp.Payload, nil
}

// UpdateProject updates a project with the specified data.
// Returns an error if name/ID pair of p does not match a stored project.
// Note: Only positive values of storageLimit are supported through this method.
// If you want to set an infinite storageLimit (-1),
// please refer to the quota client's 'UpdateStorageQuotaByProjectID' method.
func (c *RESTClient) UpdateProject(ctx context.Context, p *modelv2.Project, storageLimit *int64) error {
	project, err := c.GetProject(ctx, p.Name)
	if err != nil {
		return err
	}

	if p.ProjectID != project.ProjectID {
		return &common.ErrProjectMismatch{}
	}

	pReq := &modelv2.ProjectReq{
		CVEAllowlist: p.CVEAllowlist,
		Metadata:     p.Metadata,
		ProjectName:  p.Name,
		StorageLimit: storageLimit,
	}

	params := &projectapi.UpdateProjectParams{
		Project:         pReq,
		ProjectNameOrID: util.ProjectIDAsString(p.ProjectID),
		Context:         ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err = c.V2Client.Project.UpdateProject(params, c.AuthInfo)

	return handleSwaggerProjectErrors(err)
}

// ProjectExists returns true, if p matches a project on server side.
// Returns false, if not found.
// Returns an error in case of communication problems.
func (c *RESTClient) ProjectExists(ctx context.Context, projectNameOrID string) (bool, error) {
	_, err := c.GetProject(ctx, projectNameOrID)
	if err != nil {
		if errors.Is(err, &common.ErrProjectNotFound{}) {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}
