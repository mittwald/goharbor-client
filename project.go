package goharborclient

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/mittwald/goharbor-client/api/v1.10.0/client/products"
	"github.com/mittwald/goharbor-client/api/v1.10.0/model"
)

const (
	// ErrProjectIlligalIDFormat describes an illegal request format
	ErrProjectIlligalIDFormat = "illegal format of provided ID value"

	// ErrProjectUnauthorized describes an unauthorized request
	ErrProjectUnauthorized = "unauthorized"

	// ErrProjectInternalErrors describes server-side internal errors
	ErrProjectInternalErrors = "unexpected internal errors"

	// ErrProjectNoPermission describes a request error without permission
	ErrProjectNoPermission = "user does not have permission to the project"

	// ErrProjectIDNotExists describes an error
	// when no proper project ID is found
	ErrProjectIDNotExists = "project ID does not exist"

	// ErrProjectNameAlreadyExists describes a duplicate project name error
	ErrProjectNameAlreadyExists = "project name already exists"

	// ErrProjectMismatch describes a failed lookup
	// of a project with name/id pair
	ErrProjectMismatch = "id/name pair not found on server side"

	// ErrProjectNotFound describes an error
	// when a specific project is not found
	ErrProjectNotFound = "project not found on server side"
)

// ProjectError is an error describing a errors related to project operations
// and implements the error interface.
type ProjectError struct {
	// ID of the related project. -1 means undefined.
	ProjectID int32

	// Name of the related project. Empty string means undefined.
	ProjectName string

	// Error message of the related project.
	errorMessage string
}

// Error implements the Error interface.
func (p *ProjectError) Error() string {
	return fmt.Sprintf("%s (project: %s, id: %d)",
		p.errorMessage, p.ProjectName, p.ProjectID)
}

// NewProjectError creates a new ProjectError.
func NewProjectError(msg string, id int32, name string) error {
	return &ProjectError{
		ProjectID:    id,
		ProjectName:  name,
		errorMessage: msg,
	}
}

// ProjectRESTClient is a subclient for RESTClient handling project related
// actions.
type ProjectRESTClient struct {
	parent *RESTClient
}

// NewProject creates a new project with name as project name.
// CountLimit and StorageLimit limits space and access for this project.
// Returns the project as it is stored inside Harbor or an error,
// if the project could not be created.
func (c *ProjectRESTClient) NewProject(ctx context.Context, name string,
	countLimit int, storageLimit int) (*model.Project, error) {
	pReq := &model.ProjectReq{
		CveWhitelist: nil,
		Metadata:     nil,
		ProjectName:  name,
		CountLimit:   int64(countLimit),
		StorageLimit: int64(storageLimit) * 1024 * 1024,
	}

	_, err := c.parent.Client.Products.PostProjects(
		&products.PostProjectsParams{
			Project: pReq,
			Context: ctx,
		}, c.parent.AuthInfo)

	err = handleSwaggerProjectErrors(err, -1, name)
	if err != nil {
		return nil, err
	}

	project, err := c.Get(ctx, name)
	if err != nil {
		return nil, err
	}

	return project, nil
}

// Delete deletes a project.
// Returns an error when no matching project is found or when
// having difficulties talking to the API.
func (c *ProjectRESTClient) Delete(ctx context.Context,
	p *model.Project) error {
	if p == nil {
		return errors.New("no project provided")
	}

	project, err := c.Get(ctx, p.Name)
	if err != nil {
		return err
	}

	if p.ProjectID != project.ProjectID {
		return NewProjectError(ErrProjectMismatch, p.ProjectID, p.Name)
	}

	_, err = c.parent.Client.Products.DeleteProjectsProjectID(
		&products.DeleteProjectsProjectIDParams{
			ProjectID: int64(project.ProjectID),
			Context:   ctx,
		}, c.parent.AuthInfo)

	return handleSwaggerProjectErrors(err, p.ProjectID, p.Name)
}

// Get returns a project identified by name.
// Returns an error if it cannot find a matching project or when
// having difficulties talking to the API.
func (c *ProjectRESTClient) Get(ctx context.Context,
	name string) (*model.Project, error) {
	if name == "" {
		return nil, errors.New("no name provided")
	}
	resp, err := c.parent.Client.Products.GetProjects(
		&products.GetProjectsParams{
			Name:    &name,
			Context: ctx,
		}, c.parent.AuthInfo)

	err = handleSwaggerProjectErrors(err, -1, name)
	if err != nil {
		return nil, err
	}

	for _, p := range resp.Payload {
		if p.Name == name {
			return p, nil
		}
	}

	return nil, NewProjectError(ErrProjectNotFound, -1, name)
}

// List retrieves projects filtered by name (all if name is empty string).
// Returns an error if no projects were found.
func (c *ProjectRESTClient) List(ctx context.Context,
	nameFilter string) ([]*model.Project, error) {
	resp, err := c.parent.Client.Products.GetProjects(
		&products.GetProjectsParams{
			Name:    &nameFilter,
			Context: ctx,
		}, c.parent.AuthInfo)

	err = handleSwaggerProjectErrors(err, -1, "")
	if err != nil {
		return nil, err
	}

	if len(resp.Payload) == 0 {
		return nil, NewProjectError(ErrProjectNotFound, -1, nameFilter)
	}

	return resp.Payload, nil
}

// Update overwrites properties of a stored project with properties of p.Update.
// Return an error if Name/ID pair of p does not match a stored project.
func (c *ProjectRESTClient) Update(ctx context.Context, p *model.Project,
	countLimit int, storageLimit int) error {
	project, err := c.Get(ctx, p.Name)
	if err != nil {
		return err
	}

	if p.ProjectID != project.ProjectID {
		return NewProjectError(ErrProjectMismatch, p.ProjectID, p.Name)
	}

	pReq := &model.ProjectReq{
		CveWhitelist: p.CveWhitelist,
		Metadata:     p.Metadata,
		ProjectName:  p.Name,
		CountLimit:   int64(countLimit),
		StorageLimit: int64(storageLimit) * 1024 * 1024,
	}

	_, err = c.parent.Client.Products.PutProjectsProjectID(
		&products.PutProjectsProjectIDParams{
			Project:   pReq,
			ProjectID: int64(p.ProjectID),
			Context:   ctx,
		}, c.parent.AuthInfo)

	return handleSwaggerProjectErrors(err, p.ProjectID, p.Name)
}

// handleProjectErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with proper message.
func handleSwaggerProjectErrors(in error, id int32, name string) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case 400:
			return NewProjectError(ErrProjectIlligalIDFormat, id, name)
		case 401:
			return NewProjectError(ErrProjectUnauthorized, id, name)
		case 403:
			return NewProjectError(ErrProjectNoPermission, id, name)
		case 500:
			return NewProjectError(ErrProjectInternalErrors, id, name)
		}
	}

	switch in.(type) {
	case *products.DeleteProjectsProjectIDNotFound:
		return NewProjectError(ErrProjectIDNotExists, id, name)
	case *products.PutProjectsProjectIDNotFound:
		return NewProjectError(ErrProjectIDNotExists, id, name)
	case *products.PostProjectsConflict:
		return NewProjectError(ErrProjectNameAlreadyExists, id, name)
	default:
		return in
	}
}
