//go:build !integration

package repository

import (
	"context"
	"testing"

	"github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/repository"
	"github.com/testwill/goharbor-client/v5/apiv2/mocks"
	"github.com/testwill/goharbor-client/v5/apiv2/model"
	clienttesting "github.com/testwill/goharbor-client/v5/apiv2/pkg/testing"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var ctx = context.Background()

func APIandMockClientsForTests() (*RESTClient, *clienttesting.MockClients) {
	desiredMockClients := &clienttesting.MockClients{
		Repository:      mocks.MockRepositoryClientService{},
		ProjectMetadata: mocks.MockProject_metadataClientService{},
	}

	v2Client := clienttesting.BuildV2ClientWithMocks(desiredMockClients)

	cl := NewClient(v2Client, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	return cl, desiredMockClients
}

var (
	projectName    = "test-project"
	repositoryName = "test-repository"
)

func TestRESTClient_GetRepository(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := &repository.GetRepositoryParams{
		ProjectName:    projectName,
		RepositoryName: repositoryName,
		Context:        ctx,
	}

	getParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Repository.On("GetRepository", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&repository.GetRepositoryOK{Payload: &model.Repository{}}, nil)

	_, err := apiClient.GetRepository(ctx, projectName, repositoryName)
	require.NoError(t, err)

	mockClient.Retention.AssertExpectations(t)
}

func TestRESTClient_UpdateRepository(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	updateParams := &repository.UpdateRepositoryParams{
		ProjectName:    projectName,
		RepositoryName: repositoryName,
		Repository:     &model.Repository{},
		Context:        ctx,
	}

	updateParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Repository.On("UpdateRepository", updateParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&repository.UpdateRepositoryOK{}, nil)

	err := apiClient.UpdateRepository(ctx, projectName, repositoryName, &model.Repository{})
	require.NoError(t, err)

	mockClient.Retention.AssertExpectations(t)
}

func TestRESTClient_DeleteRepository(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	deleteParams := &repository.DeleteRepositoryParams{
		ProjectName:    projectName,
		RepositoryName: repositoryName,
		Context:        ctx,
	}

	deleteParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Repository.On("DeleteRepository", deleteParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&repository.DeleteRepositoryOK{}, nil)

	err := apiClient.DeleteRepository(ctx, projectName, repositoryName)
	require.NoError(t, err)

	mockClient.Retention.AssertExpectations(t)
}

func TestRESTClient_ListAllRepositories(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	listParams := &repository.ListAllRepositoriesParams{
		Page:     &apiClient.Options.Page,
		PageSize: &apiClient.Options.PageSize,
		Q:        &apiClient.Options.Query,
		Sort:     &apiClient.Options.Sort,
		Context:  ctx,
	}

	listParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Repository.On("ListAllRepositories", listParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&repository.ListAllRepositoriesOK{Payload: []*model.Repository{}}, nil)

	_, err := apiClient.ListAllRepositories(ctx)
	require.NoError(t, err)

	mockClient.Retention.AssertExpectations(t)
}

func TestRESTClient_ListRepositories(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	listParams := &repository.ListRepositoriesParams{
		Page:        &apiClient.Options.Page,
		PageSize:    &apiClient.Options.PageSize,
		ProjectName: projectName,
		Q:           &apiClient.Options.Query,
		Sort:        &apiClient.Options.Sort,
		Context:     ctx,
	}
	listParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Repository.On("ListRepositories", listParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&repository.ListRepositoriesOK{Payload: []*model.Repository{}}, nil)

	_, err := apiClient.ListRepositories(ctx, projectName)
	require.NoError(t, err)

	mockClient.Retention.AssertExpectations(t)
}
