//go:build !integration

package project

import (
	"context"
	"net/http"
	"strconv"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/stretchr/testify/require"

	projectapi "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/project"
	modelv2 "github.com/mittwald/goharbor-client/v5/apiv2/model"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/errors"
	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"

	"github.com/stretchr/testify/mock"

	"github.com/mittwald/goharbor-client/v5/apiv2/mocks"
)

var (
	exampleStorageLimitPositive = int64(1)
	exampleStorageLimitNegative = int64(-1)
	exampleProjectID            = int64(1)
	exampleProject              = &modelv2.Project{Name: "example-project", ProjectID: int32(exampleProjectID)}
	pReq                        = &modelv2.ProjectReq{
		ProjectName:  "example-project",
		StorageLimit: &exampleStorageLimitPositive,
	}
	pReq3 = &modelv2.ProjectReq{
		ProjectName:  "example-project",
		StorageLimit: &exampleStorageLimitNegative,
		RegistryID:   int64Ptr(0),
	}
	ctx = context.Background()
)

// int64Ptr returns a pointer to the given int64 value.
func int64Ptr(i int64) *int64 {
	return &i
}

func APIandMockClientsForTests() (*RESTClient, *clienttesting.MockClients) {
	desiredMockClients := &clienttesting.MockClients{
		Project: mocks.MockProjectClientService{},
	}

	v2Client := clienttesting.BuildV2ClientWithMocks(desiredMockClients)

	cl := NewClient(v2Client, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	return cl, desiredMockClients
}

func TestRESTClient_NewProject(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	postParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	postParams.WithTimeout(apiClient.Options.Timeout)

	getParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	getParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Project.On("CreateProject", postParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, nil)

	err := apiClient.NewProject(ctx, &modelv2.ProjectReq{
		ProjectName:  exampleProject.Name,
		StorageLimit: &exampleStorageLimitPositive,
	})

	require.NoError(t, err)

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_NewProject_UnlimitedStorage(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	createParams := &projectapi.CreateProjectParams{
		Project: pReq3,
		Context: ctx,
	}

	createParams.WithTimeout(apiClient.Options.Timeout)

	getParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	getParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Project.On("CreateProject", createParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, nil)

	err := apiClient.NewProject(ctx, &modelv2.ProjectReq{
		ProjectName:  exampleProject.Name,
		StorageLimit: &exampleStorageLimitNegative,
		RegistryID:   int64Ptr(0),
	})

	require.NoError(t, err)

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectUnauthorized(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	createParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	createParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Project.On("CreateProject", createParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, &runtime.APIError{Code: http.StatusUnauthorized})

	err := apiClient.NewProject(ctx, &modelv2.ProjectReq{
		ProjectName:  exampleProject.Name,
		StorageLimit: &exampleStorageLimitPositive,
	})

	require.Error(t, err)
	require.ErrorIs(t, err, &errors.ErrUnauthorized{})

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectNoPermission(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	createParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	createParams.WithTimeout(apiClient.Options.Timeout)

	getParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	getParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Project.On("GetProject", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{Payload: nil}, &runtime.APIError{Code: http.StatusForbidden})

	_, err := apiClient.GetProject(ctx, exampleProject.Name)

	require.Error(t, err)
	require.ErrorIs(t, err, &errors.ErrProjectNoPermission{})

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectUnknownResource(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	createParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	createParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Project.On("CreateProject", createParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &errors.ErrProjectUnknownResource{})

	err := apiClient.NewProject(ctx, &modelv2.ProjectReq{
		ProjectName:  exampleProject.Name,
		StorageLimit: &exampleStorageLimitPositive,
	})

	require.Error(t, err)
	require.ErrorIs(t, err, &errors.ErrProjectUnknownResource{})

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectInternalErrors(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	createParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	createParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Project.On("CreateProject", createParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, &runtime.APIError{Code: http.StatusInternalServerError})

	err := apiClient.NewProject(ctx, &modelv2.ProjectReq{
		ProjectName:  exampleProject.Name,
		StorageLimit: &exampleStorageLimitPositive,
	})
	require.Error(t, err)
	require.ErrorIs(t, err, &errors.ErrProjectInternalErrors{})

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectIDNotExists(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	createParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	createParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Project.On("CreateProject", createParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, &runtime.APIError{Code: http.StatusNotFound})

	err := apiClient.NewProject(ctx, &modelv2.ProjectReq{
		ProjectName:  exampleProject.Name,
		StorageLimit: &exampleStorageLimitPositive,
	})
	require.Error(t, err)
	require.ErrorIs(t, err, &errors.ErrProjectUnknownResource{})

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectNameAlreadyExists(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	createParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	createParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Project.On("CreateProject", createParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, &projectapi.CreateProjectConflict{})

	err := apiClient.NewProject(ctx, &modelv2.ProjectReq{
		ProjectName:  exampleProject.Name,
		StorageLimit: &exampleStorageLimitPositive,
	})

	require.Error(t, err)
	require.ErrorIs(t, err, &errors.ErrProjectNameAlreadyExists{})

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_Project_ErrProjectNotProvided(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	err := apiClient.DeleteProject(ctx, "")

	require.Error(t, err)
	require.ErrorIs(t, err, &errors.ErrProjectNameNotProvided{})

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_DeleteProject(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	getParams.WithTimeout(apiClient.Options.Timeout)

	deleteParams := &projectapi.DeleteProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	deleteParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Project.On("GetProject", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{
			Payload: exampleProject,
		}, nil)

	mockClient.Project.On("DeleteProject",
		deleteParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.DeleteProjectOK{}, nil)

	err := apiClient.DeleteProject(ctx, exampleProject.Name)

	require.NoError(t, err)

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_DeleteProject_ErrProjectMismatch(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := &projectapi.GetProjectParams{
		ProjectNameOrID: "bar",
		Context:         ctx,
	}

	getParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Project.On("GetProject", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &errors.ErrProjectNotFound{})

	err := apiClient.DeleteProject(ctx, "bar")

	require.Error(t, err)
	require.ErrorIs(t, err, &errors.ErrProjectMismatch{})

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_DeleteProject_ErrProjectUnknownResource(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := &projectapi.GetProjectParams{
		ProjectNameOrID: "foo",
		Context:         ctx,
	}

	getParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Project.On("GetProject", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{
			Payload: nil,
		}, &runtime.APIError{Code: http.StatusNotFound})

	err := apiClient.DeleteProject(ctx, "foo")
	require.Error(t, err)
	require.ErrorIs(t, err, &errors.ErrProjectUnknownResource{})

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_GetProject(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	getParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Project.On("GetProject", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{
			Payload: exampleProject,
		}, nil)

	_, err := apiClient.GetProject(ctx, exampleProject.Name)

	require.NoError(t, err)

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_GetProject_ErrProjectNameNotProvided(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	_, err := apiClient.GetProject(ctx, "")

	require.Error(t, err)
	require.ErrorIs(t, err, &errors.ErrProjectNameNotProvided{})

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_ListProjects(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	listParams := &projectapi.ListProjectsParams{
		Name:     &exampleProject.Name,
		PageSize: &apiClient.Options.PageSize,
		Q:        &apiClient.Options.Query,
		Sort:     &apiClient.Options.Sort,
		Context:  ctx,
	}

	listParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Project.On("ListProjects", listParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.ListProjectsOK{Payload: []*modelv2.Project{exampleProject}}, nil)

	_, err := apiClient.ListProjects(ctx, exampleProject.Name)

	require.NoError(t, err)

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_ListProjectsErrProjectNotFound(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	listParams := &projectapi.ListProjectsParams{
		Name:     &exampleProject.Name,
		PageSize: &apiClient.Options.PageSize,
		Q:        &apiClient.Options.Query,
		Sort:     &apiClient.Options.Sort,
		Context:  ctx,
	}

	listParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Project.On("ListProjects", listParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.ListProjectsOK{}, nil)

	resp, err := apiClient.ListProjects(ctx, exampleProject.Name)

	require.Error(t, err)
	require.ErrorIs(t, err, &errors.ErrProjectNotFound{})
	require.Nil(t, resp)

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_ListProjects_ErrProjectUnknownResource(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	listParams := &projectapi.ListProjectsParams{
		Name:     &exampleProject.Name,
		PageSize: &apiClient.Options.PageSize,
		Q:        &apiClient.Options.Query,
		Sort:     &apiClient.Options.Sort,
		Context:  ctx,
	}

	listParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Project.On("ListProjects", listParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.ListProjectsOK{Payload: []*modelv2.Project{exampleProject}}, &runtime.APIError{Code: http.StatusNotFound})

	_, err := apiClient.ListProjects(ctx, exampleProject.Name)

	require.Error(t, err)
	require.ErrorIs(t, err, &errors.ErrProjectUnknownResource{})
}

func TestRESTClient_UpdateProject(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	getParams.WithTimeout(apiClient.Options.Timeout)

	updateProjectParams := &projectapi.UpdateProjectParams{
		Project:         pReq3,
		ProjectNameOrID: strconv.Itoa(int(exampleProjectID)),
		Context:         ctx,
	}

	updateProjectParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Project.On("GetProject", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{
			Payload: exampleProject,
		}, nil)

	mockClient.Project.On("UpdateProject",
		updateProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.UpdateProjectOK{}, nil)

	project, err := apiClient.GetProject(ctx, exampleProject.Name)

	require.NoError(t, err)

	err = apiClient.UpdateProject(ctx, project, &exampleStorageLimitNegative)

	require.NoError(t, err)

	mockClient.Project.AssertExpectations(t)
}
