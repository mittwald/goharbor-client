//go:build !integration

package project

import (
	"context"
	"net/http"
	"strconv"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/stretchr/testify/require"

	projectapi "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/project"
	modelv2 "github.com/mittwald/goharbor-client/v4/apiv2/model"
	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/common"
	unittesting "github.com/mittwald/goharbor-client/v4/apiv2/pkg/testing"

	runtimeclient "github.com/go-openapi/runtime/client"

	"github.com/stretchr/testify/mock"

	"github.com/mittwald/goharbor-client/v4/apiv2/mocks"
)

var (
	authInfo                    = runtimeclient.BasicAuth("foo", "bar")
	exampleStorageLimitPositive = int64(1)
	exampleStorageLimitNegative = int64(-1)
	exampleProjectID            = int64(1)
	exampleUser                 = "example-user"
	exampleUserRoleID           = int64(1)
	exampleProject              = &modelv2.Project{Name: "example-project", ProjectID: int32(exampleProjectID)}
	exampleProject2             = &modelv2.Project{Name: "example-project", ProjectID: int32(exampleProjectID + 1)}
	exampleProject3             = &modelv2.Project{Name: "example-project", ProjectID: int32(exampleProjectID)}
	pReq                        = &modelv2.ProjectReq{
		ProjectName:  "example-project",
		StorageLimit: &exampleStorageLimitPositive,
	}
	pReq2 = &modelv2.ProjectReq{
		ProjectName: "example-project",
		Metadata:    &modelv2.ProjectMetadata{},
	}
	pReq3 = &modelv2.ProjectReq{
		ProjectName:  "example-project",
		StorageLimit: &exampleStorageLimitNegative,
	}
	exampleMetadataValue = "true"
	ctx                  = context.Background()
)

func APIandMockClientsForTests() (*RESTClient, *unittesting.MockClients) {
	desiredMockClients := &unittesting.MockClients{
		Project: mocks.MockProjectClientService{},
	}

	v2Client := unittesting.BuildV2ClientWithMocks(desiredMockClients)

	cl := NewClient(v2Client, &unittesting.DefaultOpts, authInfo)

	return cl, desiredMockClients
}

func TestRESTClient_NewProject(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	postParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	getParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	mockClient.Project.On("CreateProject", postParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, nil)

	mockClient.Project.On("GetProject", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{Payload: exampleProject}, nil)

	_, err := apiClient.NewProject(ctx, exampleProject.Name, &exampleStorageLimitPositive)

	require.NoError(t, err)

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_NewProject_UnlimitedStorage(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	createParams := &projectapi.CreateProjectParams{
		Project: pReq3,
		Context: ctx,
	}

	getParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	mockClient.Project.On("CreateProject", createParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, nil)

	mockClient.Project.On("GetProject", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{Payload: exampleProject3}, nil)

	_, err := apiClient.NewProject(ctx, exampleProject3.Name, &exampleStorageLimitNegative)

	require.NoError(t, err)

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectNotFound(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	createParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	getParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	mockClient.Project.On("CreateProject", createParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, nil)

	mockClient.Project.On("GetProject", getParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &runtime.APIError{Code: http.StatusNotFound})

	_, err := apiClient.NewProject(ctx, exampleProject.Name, &exampleStorageLimitPositive)

	require.Error(t, err)
	require.ErrorIs(t, err, &common.ErrProjectNotFound{})

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectIllegalIDFormat(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	createParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	getParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	mockClient.Project.On("CreateProject", createParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, nil)

	mockClient.Project.On("GetProject", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{Payload: nil}, &common.ErrProjectIllegalIDFormat{})

	_, err := apiClient.NewProject(ctx, exampleProject.Name, &exampleStorageLimitPositive)

	require.Error(t, err)
	require.ErrorIs(t, err, &common.ErrProjectIllegalIDFormat{})

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectUnauthorized(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	createParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	mockClient.Project.On("CreateProject", createParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, &runtime.APIError{Code: http.StatusUnauthorized})

	_, err := apiClient.NewProject(ctx, exampleProject.Name, &exampleStorageLimitPositive)

	require.Error(t, err)
	require.ErrorIs(t, err, &common.ErrUnauthorized{})

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectNoPermission(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	createParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	getParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	mockClient.Project.On("CreateProject", createParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, nil)

	mockClient.Project.On("GetProject", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{Payload: nil}, &runtime.APIError{Code: http.StatusForbidden})

	_, err := apiClient.NewProject(ctx, exampleProject.Name, &exampleStorageLimitPositive)

	require.Error(t, err)
	require.ErrorIs(t, err, &common.ErrProjectNoPermission{})

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectUnknownResource(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	createParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	mockClient.Project.On("CreateProject", createParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &common.ErrProjectUnknownResource{})

	_, err := apiClient.NewProject(ctx, exampleProject.Name, &exampleStorageLimitPositive)

	require.Error(t, err)
	require.ErrorIs(t, err, &common.ErrProjectUnknownResource{})

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectInternalErrors(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	createParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	mockClient.Project.On("CreateProject", createParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, &runtime.APIError{Code: http.StatusInternalServerError})

	_, err := apiClient.NewProject(ctx, exampleProject.Name, &exampleStorageLimitPositive)

	require.Error(t, err)
	require.ErrorIs(t, err, &common.ErrProjectInternalErrors{})

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectIDNotExists(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	createParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	mockClient.Project.On("CreateProject", createParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, &runtime.APIError{Code: http.StatusNotFound})

	_, err := apiClient.NewProject(ctx, exampleProject.Name, &exampleStorageLimitPositive)

	require.Error(t, err)
	require.ErrorIs(t, err, &common.ErrProjectUnknownResource{})

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectNameAlreadyExists(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	createParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	mockClient.Project.On("CreateProject", createParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, &projectapi.CreateProjectConflict{})

	_, err := apiClient.NewProject(ctx, exampleProject.Name, &exampleStorageLimitPositive)

	require.Error(t, err)
	require.ErrorIs(t, err, &common.ErrProjectNameAlreadyExists{})

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_Project_ErrProjectNotProvided(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	err := apiClient.DeleteProject(ctx, nil)

	require.Error(t, err)
	require.ErrorIs(t, err, &common.ErrProjectNotProvided{})

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_DeleteProject(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	deleteParams := &projectapi.DeleteProjectParams{
		ProjectNameOrID: strconv.Itoa(int(exampleProjectID)),
		Context:         ctx,
	}

	mockClient.Project.On("GetProject", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{
			Payload: exampleProject,
		}, nil)

	mockClient.Project.On("DeleteProject",
		deleteParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.DeleteProjectOK{}, nil)

	err := apiClient.DeleteProject(ctx, exampleProject)

	require.NoError(t, err)

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_DeleteProject_ErrProjectMismatch(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	n := "example-nonexistent"
	nonExistentProject := &modelv2.Project{Name: n}

	getParams := &projectapi.GetProjectParams{
		ProjectNameOrID: nonExistentProject.Name,
		Context:         ctx,
	}

	mockClient.Project.On("GetProject", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &common.ErrProjectNotFound{})

	err := apiClient.DeleteProject(ctx, nonExistentProject)

	require.Error(t, err)
	require.ErrorIs(t, err, &common.ErrProjectMismatch{})

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_DeleteProject_ErrProjectUnknownResource(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()
	n := "example-nonexistent"
	nonExistentProject := &modelv2.Project{Name: n}

	getParams := &projectapi.GetProjectParams{
		ProjectNameOrID: n,
		Context:         ctx,
	}

	mockClient.Project.On("GetProject", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{
			Payload: nil,
		}, &runtime.APIError{Code: http.StatusNotFound})

	err := apiClient.DeleteProject(ctx, nonExistentProject)
	require.Error(t, err)
	require.ErrorIs(t, err, &common.ErrProjectUnknownResource{})

	mockClient.Project.AssertExpectations(t)
}

func TestRESTClient_GetProject(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

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
	require.ErrorIs(t, err, &common.ErrProjectNameNotProvided{})

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

	mockClient.Project.On("ListProjects", listParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.ListProjectsOK{}, nil)

	resp, err := apiClient.ListProjects(ctx, exampleProject.Name)

	require.Error(t, err)
	require.ErrorIs(t, err, &common.ErrProjectNotFound{})
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

	mockClient.Project.On("ListProjects", listParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.ListProjectsOK{Payload: []*modelv2.Project{exampleProject}}, &runtime.APIError{Code: http.StatusNotFound})

	_, err := apiClient.ListProjects(ctx, exampleProject.Name)

	require.Error(t, err)
	require.ErrorIs(t, err, &common.ErrProjectUnknownResource{})
}

func TestRESTClient_UpdateProject(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	updateProjectParams := &projectapi.UpdateProjectParams{
		Project:         pReq3,
		ProjectNameOrID: strconv.Itoa(int(exampleProjectID)),
		Context:         ctx,
	}

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

//
// func TestRESTClient_UpdateProject_ErrProjectInternalErrors(t *testing.T) {
// 	p := &mocks.MockProjectClientService{}
//
// 	legacyClient := BuildLegacyClientWithMock(nil)
// 	v2Client := BuildV2ClientWithMocks(p, nil)
//
// 	cl := NewClient(legacyClient, v2Client, authInfo)
//
// 	ctx := context.Background()
//
// 	getParams := &projectapi.GetProjectParams{
// 		ProjectNameOrID: exampleProject.Name,
// 		Context:         ctx,
// 	}
//
// 	mockClient.Project.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
// 		Return(&projectapi.GetProjectOK{}, &runtime.APIError{
// 			OperationName: "",
// 			Response:      nil,
// 			Code:          500,
// 		})
//
// 	err := apiClient.UpdateProject(ctx, exampleProject, &exampleStorageLimitPositive)
//
// 	require.Error(t, err)
// 		require.IsType(t, &common.ErrProjectInternalErrors{}, err)
// 	}
//
// 	mockClient.Project.AssertExpectations(t)
// }
//
// func TestRESTClient_UpdateProject_ErrProjectMismatch(t *testing.T) {
// 	p := &mocks.MockProjectClientService{}
//
// 	legacyClient := BuildLegacyClientWithMock(nil)
// 	v2Client := BuildV2ClientWithMocks(p, nil)
//
// 	cl := NewClient(legacyClient, v2Client, authInfo)
//
// 	ctx := context.Background()
//
// 	project2 := *exampleProject
//
// 	getParams := &projectapi.GetProjectParams{
// 		ProjectNameOrID: exampleProject.Name,
// 		Context:         ctx,
// 	}
//
// 	mockClient.Project.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
// 		Return(&projectapi.GetProjectOK{Payload: exampleProject}, nil)
//
// 	project2.ProjectID = 100
// 	err := apiClient.UpdateProject(ctx, &project2, &exampleStorageLimitPositive)
//
// 	require.Error(t, err)
// 		require.IsType(t, &common.ErrProjectMismatch{}, err)
// 	}
//
// 	mockClient.Project.AssertExpectations(t)
// }
