package repository

import (
	"context"

	"github.com/go-openapi/runtime"
	v2client "github.com/testwill/goharbor-client/v5/apiv2/internal/api/client"
	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/repository"
	"github.com/testwill/goharbor-client/v5/apiv2/model"
	"github.com/testwill/goharbor-client/v5/apiv2/pkg/config"
	"github.com/testwill/goharbor-client/v5/apiv2/pkg/errors"
)

// RESTClient is a subclient for handling repository related actions.
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
	GetRepository(ctx context.Context, projectName, repositoryName string) (*model.Repository, error)
	UpdateRepository(ctx context.Context, projectName, repositoryName string, update *model.Repository) error
	ListAllRepositories(ctx context.Context) ([]*model.Repository, error)
	ListRepositories(ctx context.Context, projectName string) ([]*model.Repository, error)
	DeleteRepository(ctx context.Context, projectName, repositoryName string) error
}

func (c *RESTClient) GetRepository(ctx context.Context, projectName, repositoryName string) (*model.Repository, error) {
	params := &repository.GetRepositoryParams{
		ProjectName:    projectName,
		RepositoryName: repositoryName,
		Context:        ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.Repository.GetRepository(params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerRepositoryErrors(err)
	}

	if resp.Payload != nil {
		return resp.Payload, nil
	}

	return nil, &errors.ErrNotFound{}
}

func (c *RESTClient) UpdateRepository(ctx context.Context, projectName, repositoryName string, update *model.Repository) error {
	params := &repository.UpdateRepositoryParams{
		ProjectName:    projectName,
		Repository:     update,
		RepositoryName: repositoryName,
		Context:        ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Repository.UpdateRepository(params, c.AuthInfo)
	if err != nil {
		return handleSwaggerRepositoryErrors(err)
	}

	return nil
}

func (c *RESTClient) ListAllRepositories(ctx context.Context) ([]*model.Repository, error) {
	params := &repository.ListAllRepositoriesParams{
		Page:     &c.Options.Page,
		PageSize: &c.Options.PageSize,
		Q:        &c.Options.Query,
		Sort:     &c.Options.Sort,
		Context:  ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.Repository.ListAllRepositories(params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerRepositoryErrors(err)
	}

	if resp.Payload != nil {
		return resp.Payload, nil
	}

	return nil, &errors.ErrNotFound{}
}

func (c *RESTClient) ListRepositories(ctx context.Context, projectName string) ([]*model.Repository, error) {
	params := &repository.ListRepositoriesParams{
		Page:        &c.Options.Page,
		PageSize:    &c.Options.PageSize,
		ProjectName: projectName,
		Q:           &c.Options.Query,
		Sort:        &c.Options.Sort,
		Context:     ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.Repository.ListRepositories(params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerRepositoryErrors(err)
	}

	if resp.Payload != nil {
		return resp.Payload, nil
	}

	return nil, &errors.ErrNotFound{}
}

func (c *RESTClient) DeleteRepository(ctx context.Context, projectName, repositoryName string) error {
	params := &repository.DeleteRepositoryParams{
		ProjectName:    projectName,
		RepositoryName: repositoryName,
		Context:        ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Repository.DeleteRepository(params, c.AuthInfo)
	if err != nil {
		return handleSwaggerRepositoryErrors(err)
	}

	return nil
}
