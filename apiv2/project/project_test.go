// +build !integration

package project

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	projectapi "github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client/project"
	modelv2 "github.com/mittwald/goharbor-client/v3/apiv2/model"

	"github.com/go-openapi/runtime"
	v2client "github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/legacyapi/client"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/legacyapi/client/products"
	"github.com/mittwald/goharbor-client/v3/apiv2/mocks"
	model "github.com/mittwald/goharbor-client/v3/apiv2/model/legacy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	authInfo            = runtimeclient.BasicAuth("foo", "bar")
	exampleStorageLimit = int64(1)
	exampleProjectID    = int64(0)
	exampleUser         = "example-user"
	exampleUserRoleID   = int64(1)
	exampleProject      = &modelv2.Project{Name: "example-project", ProjectID: int32(exampleProjectID)}
	exampleProject2     = &modelv2.Project{Name: "example-project-2", ProjectID: int32(exampleProjectID + 1)}
	usr                 = &model.User{Username: exampleUser}
	sPtr                = exampleStorageLimit * 1024 * 1024
	pReq                = &modelv2.ProjectReq{
		ProjectName:  "example-project",
		StorageLimit: &sPtr,
	}
	pReq2 = &modelv2.ProjectReq{
		ProjectName: "example-project-2",
		Metadata:    &modelv2.ProjectMetadata{},
	}
	exampleMetadataKey   = ProjectMetadataKeyEnableContentTrust
	exampleMetadataValue = "true"
)

func BuildLegacyClientWithMock(service *mocks.MockProductsClientService) *client.Harbor {
	return &client.Harbor{
		Products: service,
	}
}

func BuildProjectClientWithMocks(project *mocks.MockProjectClientService) *v2client.Harbor {
	return &v2client.Harbor{
		Project: project,
	}
}

func TestRESTClient_NewProject(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	postProjectParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	listProjectParams := &projectapi.ListProjectsParams{
		Name:    &pReq.ProjectName,
		Context: ctx,
	}

	getProjectParams := &projectapi.GetProjectParams{
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	p.On("CreateProject", postProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, nil)

	p.On("ListProjects", listProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.ListProjectsOK{Payload: []*modelv2.Project{exampleProject}}, nil)

	p.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{Payload: exampleProject}, nil)

	_, err := cl.NewProject(ctx, exampleProject.Name, int(exampleStorageLimit))

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

// A workaround to test the successful return of the "201" status on a NewProject() call
func TestRESTClient_NewProject_201(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	postProjectParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	listProjectParams := &projectapi.ListProjectsParams{
		Name:    &pReq.ProjectName,
		Context: ctx,
	}

	getProjectParams := &projectapi.GetProjectParams{
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	p.On("CreateProject", postProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, nil)

	p.On("ListProjects", listProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.ListProjectsOK{Payload: []*modelv2.Project{exampleProject}}, nil)

	p.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{Payload: exampleProject}, nil)

	_, err := cl.NewProject(ctx, exampleProject.Name, int(exampleStorageLimit))

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectNotFound(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	postProjectParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	listProjectParams := &projectapi.ListProjectsParams{
		Name:    &pReq.ProjectName,
		Context: ctx,
	}

	getProjectParams := &projectapi.GetProjectParams{
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	p.On("CreateProject", postProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, nil)

	p.On("ListProjects", listProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.ListProjectsOK{Payload: []*modelv2.Project{exampleProject}}, nil)

	p.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{Payload: nil}, nil)

	_, err := cl.NewProject(ctx, exampleProject.Name, int(exampleStorageLimit))

	if assert.Error(t, err) {
		assert.IsType(t, &ErrProjectNotFound{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectIllegalIDFormat(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	postProjectParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	listProjectParams := &projectapi.ListProjectsParams{
		Name:    &pReq.ProjectName,
		Context: ctx,
	}

	getProjectParams := &projectapi.GetProjectParams{
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	p.On("CreateProject", postProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, nil)

	p.On("ListProjects", listProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.ListProjectsOK{Payload: []*modelv2.Project{exampleProject}}, nil)

	p.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{Payload: nil}, &ErrProjectIllegalIDFormat{})

	_, err := cl.NewProject(ctx, exampleProject.Name, int(exampleStorageLimit))

	if assert.Error(t, err) {
		assert.Equal(t, &ErrProjectIllegalIDFormat{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectUnauthorized(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	postProjectParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	p.On("CreateProject", postProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, &runtime.APIError{Code: http.StatusUnauthorized})

	_, err := cl.NewProject(ctx, exampleProject.Name, int(exampleStorageLimit))

	if assert.Error(t, err) {
		assert.Equal(t, &ErrProjectUnauthorized{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectNoPermission(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	postProjectParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	listProjectParams := &projectapi.ListProjectsParams{
		Name:    &pReq.ProjectName,
		Context: ctx,
	}

	getProjectParams := &projectapi.GetProjectParams{
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	p.On("CreateProject", postProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, nil)

	p.On("ListProjects", listProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.ListProjectsOK{Payload: []*modelv2.Project{exampleProject}}, nil)

	p.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{Payload: nil}, &runtime.APIError{Code: http.StatusForbidden})

	_, err := cl.NewProject(ctx, exampleProject.Name, int(exampleStorageLimit))

	if assert.Error(t, err) {
		assert.Equal(t, &ErrProjectNoPermission{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectUnknownResource(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	postProjectParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	p.On("CreateProject", postProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &ErrProjectUnknownResource{})

	_, err := cl.NewProject(ctx, exampleProject.Name, int(exampleStorageLimit))

	if assert.Error(t, err) {
		assert.Equal(t, &ErrProjectUnknownResource{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectInternalErrors(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	postProjectParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	p.On("CreateProject", postProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, &runtime.APIError{Code: http.StatusInternalServerError})

	_, err := cl.NewProject(ctx, exampleProject.Name, int(exampleStorageLimit))

	if assert.Error(t, err) {
		assert.Equal(t, &ErrProjectInternalErrors{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectIDNotExists(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	postProjectParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	p.On("CreateProject", postProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, &runtime.APIError{Code: http.StatusNotFound})

	_, err := cl.NewProject(ctx, exampleProject.Name, int(exampleStorageLimit))

	if assert.Error(t, err) {
		assert.Equal(t, &ErrProjectUnknownResource{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectNameAlreadyExists(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	postProjectParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	p.On("CreateProject", postProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, &projectapi.CreateProjectConflict{})

	_, err := cl.NewProject(ctx, exampleProject.Name, int(exampleStorageLimit))

	if assert.Error(t, err) {
		assert.Equal(t, &ErrProjectNameAlreadyExists{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectInvalidRequest(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	postProjectParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	p.On("CreateProject", postProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, &products.PostProjectsProjectIDMembersBadRequest{})

	_, err := cl.NewProject(ctx, exampleProject.Name, int(exampleStorageLimit))

	if assert.Error(t, err) {
		assert.Equal(t, &ErrProjectInvalidRequest{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectInvalidRequest_2(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	postProjectParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	p.On("CreateProject", postProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, &products.PostProjectsProjectIDMembersBadRequest{})

	_, err := cl.NewProject(ctx, exampleProject.Name, int(exampleStorageLimit))

	if assert.Error(t, err) {
		assert.Equal(t, &ErrProjectInvalidRequest{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_Project_ErrProjectNotProvided(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	t.Run("DeleteProject_ErrProjectNotProvided", func(t *testing.T) {
		err := cl.DeleteProject(ctx, nil)

		if assert.Error(t, err) {
			assert.IsType(t, &ErrProjectNotProvided{}, err)
		}
	})

	t.Run("AddProjectMember_ErrProjectNotProvided", func(t *testing.T) {
		err := cl.AddProjectMember(ctx, nil, usr, 1)

		if assert.Error(t, err) {
			assert.Equal(t, &ErrProjectNotProvided{}, err)
		}
	})

	t.Run("AddProjectMember_ErrProjectNotProvided", func(t *testing.T) {
		err := cl.AddProjectMember(ctx, nil, usr, 1)

		if assert.Error(t, err) {
			assert.Equal(t, &ErrProjectNotProvided{}, err)
		}
	})

	t.Run("ListProjectMembers_ErrProjectNotProvided", func(t *testing.T) {
		_, err := cl.ListProjectMembers(ctx, nil)

		if assert.Error(t, err) {
			assert.IsType(t, &ErrProjectNotProvided{}, err)
		}
	})

	t.Run("UpdateProjectMemberRole_ErrProjectNotProvided", func(t *testing.T) {
		err := cl.UpdateProjectMemberRole(ctx, nil, usr, int(exampleUserRoleID))

		if assert.Error(t, err) {
			assert.Equal(t, &ErrProjectNotProvided{}, err)
		}
	})

	t.Run("DeleteProjectMember_ErrProjectNotProvided", func(t *testing.T) {
		err := cl.DeleteProjectMember(ctx, nil, usr)

		if assert.Error(t, err) {
			assert.IsType(t, &ErrProjectNotProvided{}, err)
		}
	})

	t.Run("UpdateProjectMetadata_ErrProjectNotProvided", func(t *testing.T) {
		err := cl.UpdateProjectMetadata(ctx, nil, exampleMetadataKey, exampleMetadataValue)

		if assert.Error(t, err) {
			assert.IsType(t, &ErrProjectNotProvided{}, err)
		}
	})

	t.Run("ListProjectMetadata_ErrProjectNotProvided", func(t *testing.T) {
		_, err := cl.ListProjectMetadata(ctx, nil)

		if assert.Error(t, err) {
			assert.Equal(t, &ErrProjectNotProvided{}, err)
		}
	})

	t.Run("DeleteProjectMetadataValue_ErrProjectNotProvided", func(t *testing.T) {
		err := cl.DeleteProjectMetadataValue(ctx, nil, exampleMetadataKey)

		if assert.Error(t, err) {
			assert.IsType(t, &ErrProjectNotProvided{}, err)
		}
	})

	t.Run("AddProjectMetadata_ErrProjectNotProvided", func(t *testing.T) {
		err := cl.AddProjectMetadata(ctx, nil, exampleMetadataKey, exampleMetadataValue)

		if assert.Error(t, err) {
			assert.IsType(t, &ErrProjectNotProvided{}, err)
		}
	})
}

func TestRESTClient_DeleteProject(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	listProjectParams := &projectapi.ListProjectsParams{
		Name:    &pReq.ProjectName,
		Context: ctx,
	}

	getProjectParams := &projectapi.GetProjectParams{
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	deleteProjectParams := &projectapi.DeleteProjectParams{
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	p.On("ListProjects", listProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.ListProjectsOK{
			Payload: []*modelv2.Project{exampleProject},
		}, nil)

	p.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{
			Payload: exampleProject,
		}, nil)

	p.On("DeleteProject",
		deleteProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.DeleteProjectOK{}, nil)

	err := cl.DeleteProject(ctx, exampleProject)

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_DeleteProject_ErrProjectMismatch(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	n := "example-nonexistent"
	nonExistentProject := &modelv2.Project{Name: n}

	listProjectParams := &projectapi.ListProjectsParams{
		Name:    &n,
		Context: ctx,
	}

	getProjectParams := &projectapi.GetProjectParams{
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	p.On("ListProjects", listProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.ListProjectsOK{Payload: []*modelv2.Project{exampleProject}}, nil)

	p.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{Payload: nil}, nil)

	err := cl.DeleteProject(ctx, nonExistentProject)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrProjectMismatch{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_DeleteProject_ErrProjectUnknownResource(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	n := "example-nonexistent"
	nonExistentProject := &modelv2.Project{Name: n}

	listProjectParams := &projectapi.ListProjectsParams{
		Name:    &n,
		Context: ctx,
	}

	p.On("ListProjects", listProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.ListProjectsOK{Payload: []*modelv2.Project{exampleProject}}, nil)

	getProjectParams := &projectapi.GetProjectParams{
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	p.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{
			Payload: nil,
		}, &runtime.APIError{Code: http.StatusNotFound})

	err := cl.DeleteProject(ctx, nonExistentProject)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrProjectUnknownResource{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_GetProject(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	getProjectParams := &projectapi.GetProjectParams{
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	listProjectParams := &projectapi.ListProjectsParams{
		Name:    &exampleProject.Name,
		Context: ctx,
	}

	p.On("ListProjects", listProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.ListProjectsOK{Payload: []*modelv2.Project{exampleProject}}, nil)

	p.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{
			Payload: exampleProject,
		}, nil)

	_, err := cl.GetProjectByName(ctx, exampleProject.Name)

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_GetProject_ErrProjectNameNotProvided(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	_, err := cl.GetProjectByName(ctx, "")

	if assert.Error(t, err) {
		assert.IsType(t, &ErrProjectNameNotProvided{}, err)
	}
}

func TestRESTClient_ListProjects(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	listProjectParams := &projectapi.ListProjectsParams{
		Name:    &exampleProject.Name,
		Context: ctx,
	}

	p.On("ListProjects", listProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.ListProjectsOK{Payload: []*modelv2.Project{exampleProject}}, nil)

	_, err := cl.ListProjects(ctx, exampleProject.Name)

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_ListProjectsErrProjectNotFound(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	listProjectsParams := &projectapi.ListProjectsParams{
		Name:    &exampleProject.Name,
		Context: ctx,
	}

	p.On("ListProjects", listProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.ListProjectsOK{}, nil)

	resp, err := cl.ListProjects(ctx, exampleProject.Name)

	if assert.Error(t, err) {
		assert.Equal(t, &ErrProjectNotFound{}, err)
		assert.Nil(t, resp)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_ListProjects_ErrProjectUnknownResource(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	listProjectsParams := &projectapi.ListProjectsParams{
		Name:    &exampleProject.Name,
		Context: ctx,
	}

	p.On("ListProjects", listProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.ListProjectsOK{Payload: []*modelv2.Project{exampleProject}}, &runtime.APIError{Code: http.StatusNotFound})

	_, err := cl.ListProjects(ctx, exampleProject.Name)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrProjectUnknownResource{}, err)
	}
}

func TestRESTClient_UpdateProject(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	listProjectsParams := &projectapi.ListProjectsParams{
		Name:    &exampleProject.Name,
		Context: ctx,
	}

	getProjectParams := &projectapi.GetProjectParams{
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	updateProjectParams := &projectapi.UpdateProjectParams{
		Project:   pReq,
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	p.On("ListProjects", listProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.ListProjectsOK{Payload: []*modelv2.Project{exampleProject}}, nil)

	p.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{
			Payload: exampleProject,
		}, nil)

	p.On("UpdateProject",
		updateProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.UpdateProjectOK{}, nil)

	project, err := cl.GetProjectByName(ctx, exampleProject.Name)

	assert.NoError(t, err)

	err = cl.UpdateProject(ctx, project, int(exampleStorageLimit))

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_UpdateProject_ErrProjectInternalErrors(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	getProjectParams := &projectapi.GetProjectParams{
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	listProjectsParams := &projectapi.ListProjectsParams{
		Name:    &exampleProject.Name,
		Context: ctx,
	}

	p.On("ListProjects", listProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.ListProjectsOK{Payload: []*modelv2.Project{exampleProject}}, nil)

	p.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{}, &runtime.APIError{
			OperationName: "",
			Response:      nil,
			Code:          500,
		})

	err := cl.UpdateProject(ctx, exampleProject, int(exampleStorageLimit))

	if assert.Error(t, err) {
		assert.IsType(t, &ErrProjectInternalErrors{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_UpdateProject_ErrProjectMismatch(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	project2 := *exampleProject

	getProjectParams := &projectapi.GetProjectParams{
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	listProjectsParams := &projectapi.ListProjectsParams{
		Name:    &exampleProject.Name,
		Context: ctx,
	}

	p.On("ListProjects", listProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.ListProjectsOK{Payload: []*modelv2.Project{exampleProject}}, nil)

	p.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{Payload: exampleProject}, nil)

	project2.ProjectID = 100
	err := cl.UpdateProject(ctx, &project2, int(exampleStorageLimit))

	if assert.Error(t, err) {
		assert.IsType(t, &ErrProjectMismatch{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_AddProjectMember(t *testing.T) {
	p := &mocks.MockProjectClientService{}
	l := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(l)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	listProjectsParams := &projectapi.ListProjectsParams{
		Name:    &exampleProject.Name,
		Context: ctx,
	}

	getProjectsParams := &projectapi.GetProjectParams{
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	getUserParams := &products.GetUsersParams{
		Context:  ctx,
		Username: &exampleUser,
	}

	postProjectsProjectIDMembersParams := &products.PostProjectsProjectIDMembersParams{
		ProjectID: exampleProjectID,
		ProjectMember: &model.ProjectMember{
			MemberUser: &model.UserEntity{
				Username: usr.Username,
			},
			MemberGroup: &model.UserGroup{
				GroupName: "",
				GroupType: 0,
				ID:        0,
			},
			RoleID: exampleUserRoleID,
		},
		Context: ctx,
	}

	p.On("ListProjects", listProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.ListProjectsOK{Payload: []*modelv2.Project{exampleProject}}, nil)

	p.On("GetProject", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{
			Payload: exampleProject,
		}, nil)

	l.On("GetUsers", getUserParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetUsersOK{
			Payload: []*model.User{{Username: exampleUser}},
		}, nil)

	l.On("PostProjectsProjectIDMembers", postProjectsProjectIDMembersParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PostProjectsProjectIDMembersCreated{}, nil)

	err := cl.AddProjectMember(ctx, exampleProject, usr, 1)

	assert.NoError(t, err)

	p.AssertExpectations(t)
	l.AssertExpectations(t)
}

func TestRESTClient_AddProjectMember_ErrProjectUnknownResource(t *testing.T) {
	p := &mocks.MockProjectClientService{}
	l := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(l)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	listProjectsParams := &projectapi.ListProjectsParams{
		Name:    &exampleProject.Name,
		Context: ctx,
	}

	getProjectsParams := &projectapi.GetProjectParams{
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	p.On("ListProjects", listProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.ListProjectsOK{Payload: []*modelv2.Project{exampleProject}}, nil)

	p.On("GetProject", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{
			Payload: nil,
		}, &runtime.APIError{Code: http.StatusNotFound})

	err := cl.AddProjectMember(ctx, exampleProject, usr, 1)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrProjectUnknownResource{}, err)
	}

	p.AssertExpectations(t)
	l.AssertExpectations(t)
}

func TestRESTClient_AddProjectMember_ErrProjectNoMemberProvided(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	err := cl.AddProjectMember(ctx, exampleProject, nil, 1)

	if assert.Error(t, err) {
		assert.Equal(t, &ErrProjectNoMemberProvided{}, err)
	}
}

func TestRESTClient_AddProjectMember_ErrProjectMemberMismatch(t *testing.T) {
	p := &mocks.MockProjectClientService{}
	l := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(l)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	getProjectsParams := &projectapi.GetProjectParams{
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	getUserParams := &products.GetUsersParams{
		Context:  ctx,
		Username: &exampleUser,
	}

	listProjectsParams := &projectapi.ListProjectsParams{
		Name:    &exampleProject.Name,
		Context: ctx,
	}

	p.On("ListProjects", listProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.ListProjectsOK{Payload: []*modelv2.Project{exampleProject}}, nil)

	p.On("GetProject", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{
			Payload: exampleProject,
		}, nil)

	l.On("GetUsers", getUserParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetUsersOK{
			Payload: []*model.User{{Username: "example-nonexistent"}},
		}, nil)

	err := cl.AddProjectMember(ctx, exampleProject, usr, 1)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrProjectMemberMismatch{}, err)
	}

	p.AssertExpectations(t)
	l.AssertExpectations(t)
}

func TestRESTClient_ListProjectMembers(t *testing.T) {
	l := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(l)
	v2Client := BuildProjectClientWithMocks(nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	e := ""

	getProjectsProjectIDMembersParams := products.GetProjectsProjectIDMembersParams{
		Entityname: &e,
		ProjectID:  exampleProjectID,
		Context:    ctx,
	}

	l.On("GetProjectsProjectIDMembers",
		&getProjectsProjectIDMembersParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetProjectsProjectIDMembersOK{
			Payload: []*model.ProjectMemberEntity{{}},
		}, nil)

	_, err := cl.ListProjectMembers(ctx, exampleProject)

	assert.NoError(t, err)

	l.AssertExpectations(t)
}

func TestRESTClient_ListProjectMembers_ErrProjectUnknownResource(t *testing.T) {
	l := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(l)
	v2Client := BuildProjectClientWithMocks(nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	e := ""

	getProjectsProjectIDMembersParams := products.GetProjectsProjectIDMembersParams{
		Entityname: &e,
		ProjectID:  exampleProjectID,
		Context:    ctx,
	}

	l.On("GetProjectsProjectIDMembers",
		&getProjectsProjectIDMembersParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetProjectsProjectIDMembersOK{
			Payload: nil,
		}, &runtime.APIError{Code: http.StatusNotFound})

	_, err := cl.ListProjectMembers(ctx, exampleProject)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrProjectUnknownResource{}, err)
	}

	l.AssertExpectations(t)
}

func TestRESTClient_UpdateProjectMemberRole(t *testing.T) {
	p := &mocks.MockProjectClientService{}
	l := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(l)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	e := ""

	getProjectsParams := &projectapi.GetProjectParams{
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	getUserParams := &products.GetUsersParams{
		Context:  ctx,
		Username: &exampleUser,
	}

	postProjectsProjectIDMembersParams := &products.PostProjectsProjectIDMembersParams{
		ProjectID: exampleProjectID,
		ProjectMember: &model.ProjectMember{
			MemberUser: &model.UserEntity{
				Username: usr.Username,
			},
			MemberGroup: &model.UserGroup{
				GroupName: "",
				GroupType: 0,
				ID:        0,
			},
			RoleID: exampleUserRoleID,
		},
		Context: ctx,
	}

	listProjectsParams := &projectapi.ListProjectsParams{
		Name:    &exampleProject.Name,
		Context: ctx,
	}

	p.On("ListProjects", listProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.ListProjectsOK{Payload: []*modelv2.Project{exampleProject}}, nil)

	p.On("GetProject", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{
			Payload: exampleProject,
		}, nil)

	l.On("GetUsers", getUserParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetUsersOK{
			Payload: []*model.User{{Username: exampleUser}},
		}, nil)

	l.On("PostProjectsProjectIDMembers", postProjectsProjectIDMembersParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PostProjectsProjectIDMembersCreated{}, nil)

	err := cl.AddProjectMember(ctx, exampleProject, usr, int(exampleUserRoleID))

	assert.NoError(t, err)

	getProjectsProjectIDMembersParams := products.GetProjectsProjectIDMembersParams{
		Entityname: &e,
		ProjectID:  exampleProjectID,
		Context:    ctx,
	}

	putProjectsProjectIDMembersMidParams := products.PutProjectsProjectIDMembersMidParams{
		Mid:       1,
		ProjectID: exampleProjectID,
		Role:      &model.RoleRequest{RoleID: exampleUserRoleID},
		Context:   ctx,
	}

	p.On("ListProjects", listProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.ListProjectsOK{Payload: []*modelv2.Project{exampleProject}}, nil)

	p.On("GetProject", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{
			Payload: exampleProject,
		}, nil)

	l.On("GetProjectsProjectIDMembers",
		&getProjectsProjectIDMembersParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetProjectsProjectIDMembersOK{
			Payload: []*model.ProjectMemberEntity{{
				EntityType: "u",
				EntityName: exampleUser,
				ID:         exampleUserRoleID,
			}},
		}, nil)

	l.On("PutProjectsProjectIDMembersMid",
		&putProjectsProjectIDMembersMidParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PutProjectsProjectIDMembersMidOK{}, nil)

	err = cl.UpdateProjectMemberRole(ctx, exampleProject, usr, int(exampleUserRoleID))

	assert.NoError(t, err)

	p.AssertExpectations(t)
	l.AssertExpectations(t)
}

func TestRESTClient_UpdateProjectMemberRole_UserIsNoMember(t *testing.T) {
	p := &mocks.MockProjectClientService{}
	l := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(l)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	e := ""

	getProjectsParams := &projectapi.GetProjectParams{
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	getProjectsProjectIDMembersParams := products.GetProjectsProjectIDMembersParams{
		Entityname: &e,
		ProjectID:  exampleProjectID,
		Context:    ctx,
	}

	listProjectsParams := &projectapi.ListProjectsParams{
		Name:    &exampleProject.Name,
		Context: ctx,
	}

	p.On("ListProjects", listProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.ListProjectsOK{Payload: []*modelv2.Project{exampleProject}}, nil)

	p.On("GetProject", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{
			Payload: exampleProject,
		}, nil)

	l.On("GetProjectsProjectIDMembers",
		&getProjectsProjectIDMembersParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetProjectsProjectIDMembersOK{
			Payload: []*model.ProjectMemberEntity{{}},
		}, nil)

	err := cl.UpdateProjectMemberRole(ctx, exampleProject, usr, int(exampleUserRoleID))

	if assert.Error(t, err) {
		assert.Equal(t, &ErrProjectUserIsNoMember{}, err)
	}

	p.AssertExpectations(t)
	l.AssertExpectations(t)
}

func TestRESTClient_UpdateProjectMemberRole_ErrProjectNoMemberProvided(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	err := cl.UpdateProjectMemberRole(ctx, exampleProject, nil, int(exampleUserRoleID))

	if assert.Error(t, err) {
		assert.Equal(t, &ErrProjectNoMemberProvided{}, err)
	}
}

func TestRESTClient_DeleteProjectMember(t *testing.T) {
	p := &mocks.MockProjectClientService{}
	l := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(l)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	e := ""

	getProjectsParams := &projectapi.GetProjectParams{
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	getUserParams := &products.GetUsersParams{
		Context:  ctx,
		Username: &exampleUser,
	}

	postProjectsProjectIDMembersParams := &products.PostProjectsProjectIDMembersParams{
		ProjectID: exampleProjectID,
		ProjectMember: &model.ProjectMember{
			MemberUser: &model.UserEntity{
				Username: usr.Username,
			},
			MemberGroup: &model.UserGroup{
				GroupName: "",
				GroupType: 0,
				ID:        0,
			},
			RoleID: exampleUserRoleID,
		},
		Context: ctx,
	}

	listProjectsParams := &projectapi.ListProjectsParams{
		Name:    &exampleProject.Name,
		Context: ctx,
	}

	p.On("ListProjects", listProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.ListProjectsOK{Payload: []*modelv2.Project{exampleProject}}, nil)

	p.On("GetProject", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{
			Payload: exampleProject,
		}, nil)

	l.On("GetUsers", getUserParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetUsersOK{
			Payload: []*model.User{{Username: exampleUser}},
		}, nil)

	l.On("PostProjectsProjectIDMembers", postProjectsProjectIDMembersParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PostProjectsProjectIDMembersCreated{}, nil)

	err := cl.AddProjectMember(ctx, exampleProject, usr, int(exampleUserRoleID))

	assert.NoError(t, err)

	getProjectsProjectIDMembersParams := products.GetProjectsProjectIDMembersParams{
		Entityname: &e,
		ProjectID:  exampleProjectID,
		Context:    ctx,
	}

	deleteProjectsProjectIDMembersMidParams := &products.DeleteProjectsProjectIDMembersMidParams{
		Mid:       1,
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	p.On("GetProject", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{
			Payload: exampleProject,
		}, nil)

	l.On("GetProjectsProjectIDMembers",
		&getProjectsProjectIDMembersParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetProjectsProjectIDMembersOK{
			Payload: []*model.ProjectMemberEntity{{
				EntityType: "u",
				EntityName: exampleUser,
				ID:         exampleUserRoleID,
			}},
		}, nil)

	l.On("DeleteProjectsProjectIDMembersMid",
		deleteProjectsProjectIDMembersMidParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.DeleteProjectsProjectIDMembersMidOK{}, nil)

	err = cl.DeleteProjectMember(ctx, exampleProject, usr)

	assert.NoError(t, err)

	p.AssertExpectations(t)
	l.AssertExpectations(t)
}

func TestRESTClient_DeleteProjectMember_ErrProjectNoMemberProvided(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	err := cl.DeleteProjectMember(ctx, exampleProject, nil)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrProjectNoMemberProvided{}, err)
	}
}

func TestRESTClient_DeleteProjectMember_ErrProjectMismatch(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	getProjectsParams := &projectapi.GetProjectParams{
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	listProjectsParams := &projectapi.ListProjectsParams{
		Name:    &exampleProject.Name,
		Context: ctx,
	}

	p.On("ListProjects", listProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.ListProjectsOK{Payload: []*modelv2.Project{exampleProject}}, nil)

	p.On("GetProject", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{Payload: nil}, nil)

	err := cl.DeleteProjectMember(ctx, exampleProject, usr)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrProjectMismatch{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_AddProjectMetadata(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	var mPtr = "true"

	pReq.Metadata = &modelv2.ProjectMetadata{}
	pReq.Metadata.EnableContentTrust = &mPtr
	pReq.StorageLimit = nil

	updateProjectParams := &projectapi.UpdateProjectParams{
		Project:   pReq,
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	p.On("UpdateProject", updateProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.UpdateProjectOK{}, nil)

	err := cl.AddProjectMetadata(ctx, exampleProject, exampleMetadataKey, exampleMetadataValue)

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_AddProjectMetadata_ErrProjectMetadataAlreadyExists(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	var mPtr = "true"
	pReq.Metadata.EnableContentTrust = &mPtr
	pReq.StorageLimit = nil

	updateProjectParams := &projectapi.UpdateProjectParams{
		Project:   pReq,
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	p.On("UpdateProject", updateProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.UpdateProjectOK{}, &runtime.APIError{Code: http.StatusConflict})

	err := cl.AddProjectMetadata(ctx, exampleProject, exampleMetadataKey, exampleMetadataValue)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrProjectMetadataAlreadyExists{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_GetProjectMetadataValue(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	keys := []MetadataKey{
		ProjectMetadataKeyEnableContentTrust,
		ProjectMetadataKeyAutoScan,
		ProjectMetadataKeySeverity,
		ProjectMetadataKeyReuseSysCveAllowlist,
		ProjectMetadataKeyRetentionID,
		ProjectMetadataKeyPublic,
		ProjectMetadataKeyPreventVul,
	}

	var sPtr = "test"
	for i := range keys {
		getProjectsParams := &projectapi.GetProjectParams{
			ProjectID: exampleProjectID,
			Context:   ctx,
		}

		p.On("GetProject", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&projectapi.GetProjectOK{Payload: &modelv2.Project{
				Metadata: &modelv2.ProjectMetadata{
					AutoScan:             &sPtr,
					EnableContentTrust:   &sPtr,
					PreventVul:           &sPtr,
					Public:               sPtr,
					RetentionID:          &sPtr,
					ReuseSysCveAllowlist: &sPtr,
					Severity:             &sPtr,
				},
				ProjectID: int32(exampleProjectID),
			}}, nil)

		val, err := cl.GetProjectMetadataValue(ctx, exampleProjectID, keys[i])

		assert.Equal(t, val, sPtr)

		assert.NoError(t, err)

		p.AssertExpectations(t)
	}
}

func TestRESTClient_GetProjectMetadataValue_ErrProjectUnknownResource(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	keys := []MetadataKey{
		ProjectMetadataKeyEnableContentTrust,
		ProjectMetadataKeyAutoScan,
		ProjectMetadataKeySeverity,
		ProjectMetadataKeyReuseSysCveAllowlist,
		ProjectMetadataKeyRetentionID,
		ProjectMetadataKeyPublic,
		ProjectMetadataKeyPreventVul,
	}

	var sPtr = "test"

	exampleProject.Metadata = &modelv2.ProjectMetadata{
		AutoScan:             &sPtr,
		EnableContentTrust:   &sPtr,
		PreventVul:           &sPtr,
		Public:               sPtr,
		RetentionID:          &sPtr,
		ReuseSysCveAllowlist: &sPtr,
		Severity:             &sPtr,
	}

	for i := range keys {
		getProjectsParams := &projectapi.GetProjectParams{
			ProjectID: exampleProjectID,
			Context:   ctx,
		}

		p.On("GetProject", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&projectapi.GetProjectOK{Payload: nil}, &runtime.APIError{Code: http.StatusNotFound})

		_, err := cl.GetProjectMetadataValue(ctx, int64(exampleProject.ProjectID), keys[i])

		if assert.Error(t, err) {
			assert.IsType(t, &ErrProjectUnknownResource{}, err)
		}

		p.AssertExpectations(t)
	}
}

func TestRESTClient_ListProjectMetadata(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	var sPtr = "true"

	exampleProject2 := exampleProject

	exampleProject2.Metadata = &modelv2.ProjectMetadata{
		AutoScan:             &sPtr,
		EnableContentTrust:   &sPtr,
		PreventVul:           &sPtr,
		Public:               sPtr,
		RetentionID:          &sPtr,
		ReuseSysCveAllowlist: &sPtr,
		Severity:             &sPtr,
	}

	getProjectParams := &projectapi.GetProjectParams{
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	p.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{Payload: exampleProject}, nil)

	meta, err := cl.ListProjectMetadata(ctx, exampleProject)

	assert.NoError(t, err)

	assert.Equal(t, meta.EnableContentTrust, exampleProject.Metadata.EnableContentTrust)

	p.AssertExpectations(t)
}

func TestRESTClient_UpdateProjectMetadata(t *testing.T) {
	p := &mocks.MockProjectClientService{}
	l := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(l)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	getProjectsProjectIDMetadatasMetaName := &products.GetProjectsProjectIDMetadatasMetaNameParams{
		MetaName:  string(ProjectMetadataKeyEnableContentTrust),
		ProjectID: int64(exampleProject2.ProjectID),
		Context:   ctx,
	}

	deleteProjectsProjectIDMetadatasMetaName := &products.DeleteProjectsProjectIDMetadatasMetaNameParams{
		MetaName:  string(ProjectMetadataKeyEnableContentTrust),
		ProjectID: int64(exampleProject2.ProjectID),
		Context:   ctx,
	}

	var mPtr = "true"
	pReq2.Metadata.EnableContentTrust = &mPtr
	updateProjectParams := &projectapi.UpdateProjectParams{
		Project:   pReq2,
		ProjectID: int64(exampleProject2.ProjectID),
		Context:   ctx,
	}

	l.On("GetProjectsProjectIDMetadatasMetaName",
		getProjectsProjectIDMetadatasMetaName, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetProjectsProjectIDMetadatasMetaNameOK{Payload: &model.ProjectMetadata{
			EnableContentTrust: exampleMetadataValue,
		}}, nil)

	l.On("DeleteProjectsProjectIDMetadatasMetaName",
		deleteProjectsProjectIDMetadatasMetaName, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.DeleteProjectsProjectIDMetadatasMetaNameOK{}, nil)

	p.On("UpdateProject",
		updateProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.UpdateProjectOK{}, nil)

	err := cl.UpdateProjectMetadata(ctx, exampleProject2, exampleMetadataKey, exampleMetadataValue)

	assert.NoError(t, err)

	p.AssertExpectations(t)
	l.AssertExpectations(t)
}

func TestRESTClient_UpdateProjectMetadata_GetProjectMeta_ErrProjectUnknownResource(t *testing.T) {
	p := &mocks.MockProjectClientService{}
	l := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(l)
	v2Client := BuildProjectClientWithMocks(p)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	var metaPtr = "true"

	project := &modelv2.Project{
		ProjectID: int32(exampleProjectID),
		Metadata: &modelv2.ProjectMetadata{
			EnableContentTrust: &metaPtr,
		},
	}

	getProjectsProjectIDMetadatasMetaName := &products.GetProjectsProjectIDMetadatasMetaNameParams{
		MetaName:  string(ProjectMetadataKeyEnableContentTrust),
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	l.On("GetProjectsProjectIDMetadatasMetaName",
		getProjectsProjectIDMetadatasMetaName, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetProjectsProjectIDMetadatasMetaNameOK{Payload: &model.ProjectMetadata{
			EnableContentTrust: exampleMetadataValue,
		}}, &runtime.APIError{Code: http.StatusNotFound})

	err := cl.UpdateProjectMetadata(ctx, project, exampleMetadataKey, exampleMetadataValue)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrProjectUnknownResource{}, err)
	}

	p.AssertExpectations(t)
	l.AssertExpectations(t)
}

func TestRESTClient_UpdateProjectMetadata_DeleteProjectMeta_ErrProjectUnknownResource(t *testing.T) {
	l := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(l)
	v2Client := BuildProjectClientWithMocks(nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	var metaPtr = "true"

	project := &modelv2.Project{
		ProjectID: int32(exampleProjectID),
		Metadata: &modelv2.ProjectMetadata{
			EnableContentTrust: &metaPtr,
		},
	}

	getProjectsProjectIDMetadatasMetaName := &products.GetProjectsProjectIDMetadatasMetaNameParams{
		MetaName:  string(ProjectMetadataKeyEnableContentTrust),
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	deleteProjectsProjectIDMetadatasMetaName := &products.DeleteProjectsProjectIDMetadatasMetaNameParams{
		MetaName:  string(ProjectMetadataKeyEnableContentTrust),
		ProjectID: 0,
		Context:   ctx,
	}

	l.On("GetProjectsProjectIDMetadatasMetaName",
		getProjectsProjectIDMetadatasMetaName, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetProjectsProjectIDMetadatasMetaNameOK{Payload: &model.ProjectMetadata{
			EnableContentTrust: exampleMetadataValue,
		}}, nil)

	l.On("DeleteProjectsProjectIDMetadatasMetaName",
		deleteProjectsProjectIDMetadatasMetaName, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.DeleteProjectsProjectIDMetadatasMetaNameOK{}, &runtime.APIError{Code: http.StatusNotFound})

	err := cl.UpdateProjectMetadata(ctx, project, exampleMetadataKey, exampleMetadataValue)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrProjectUnknownResource{}, err)
	}

	l.AssertExpectations(t)
}

func TestRESTClient_DeleteProjectMetadataValue(t *testing.T) {
	l := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(l)
	v2Client := BuildProjectClientWithMocks(nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	var metaPtr = "true"

	project := &modelv2.Project{
		ProjectID: int32(exampleProjectID),
		Metadata: &modelv2.ProjectMetadata{
			EnableContentTrust: &metaPtr,
		},
	}

	deleteProjectsProjectIDMetadatasMetaName := &products.DeleteProjectsProjectIDMetadatasMetaNameParams{
		MetaName:  string(ProjectMetadataKeyEnableContentTrust),
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	l.On("DeleteProjectsProjectIDMetadatasMetaName",
		deleteProjectsProjectIDMetadatasMetaName, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.DeleteProjectsProjectIDMetadatasMetaNameOK{}, nil)

	err := cl.DeleteProjectMetadataValue(ctx, project, exampleMetadataKey)

	assert.NoError(t, err)

	l.AssertExpectations(t)
}

func TestRESTClient_ListProjectRobots(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildProjectClientWithMocks(nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	expectedRobots := []*model.RobotAccount{
		&model.RobotAccount{
			Description: "some robot account",
			Disabled:    false,
			ID:          42,
			Name:        "robot$account",
			ProjectID:   exampleProjectID,
		},
	}

	params := &products.GetProjectsProjectIDRobotsParams{
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	p.On("GetProjectsProjectIDRobots", params, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetProjectsProjectIDRobotsOK{Payload: expectedRobots}, nil)

	robots, err := cl.ListProjectRobots(ctx, exampleProject)

	assert.NoError(t, err)

	assert.Equal(t, expectedRobots, robots)

	p.AssertExpectations(t)
}

func TestRESTClient_AddProjectRobot(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildProjectClientWithMocks(nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	newRobot := &model.RobotAccountCreate{
		Access: []*model.RobotAccountAccess{
			{
				Action:   "push",
				Resource: fmt.Sprintf("/project/%d/repository", exampleProjectID),
			},
		},
		Description: "some robot account",
		ExpiresAt:   0,
		Name:        "my-account",
	}

	params := &products.PostProjectsProjectIDRobotsParams{
		ProjectID: exampleProjectID,
		Robot:     newRobot,
		Context:   ctx,
	}

	expectedPayload := &model.RobotAccountPostRep{
		Name:  "robot$my-account",
		Token: "very-secret-token-here",
	}

	p.On("PostProjectsProjectIDRobots", params, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PostProjectsProjectIDRobotsCreated{Payload: expectedPayload}, nil)

	token, err := cl.AddProjectRobot(ctx, exampleProject, newRobot)

	assert.NoError(t, err)

	assert.Equal(t, expectedPayload.Token, token)

	p.AssertExpectations(t)
}

func TestRESTClient_UpdateProjectRobot(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildProjectClientWithMocks(nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	const exampleRobotID = 42

	updateRobot := &model.RobotAccountUpdate{
		Disabled: true,
	}

	params := &products.PutProjectsProjectIDRobotsRobotIDParams{
		ProjectID: exampleProjectID,
		RobotID:   exampleRobotID,
		Robot:     updateRobot,
		Context:   ctx,
	}

	p.On("PutProjectsProjectIDRobotsRobotID", params, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PutProjectsProjectIDRobotsRobotIDOK{}, nil)

	err := cl.UpdateProjectRobot(ctx, exampleProject, exampleRobotID, updateRobot)

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_DeleteProjectRobot(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildProjectClientWithMocks(nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	const exampleRobotID = 42

	params := &products.DeleteProjectsProjectIDRobotsRobotIDParams{
		ProjectID: exampleProjectID,
		RobotID:   exampleRobotID,
		Context:   ctx,
	}

	p.On("DeleteProjectsProjectIDRobotsRobotID", params, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.DeleteProjectsProjectIDRobotsRobotIDOK{}, nil)

	err := cl.DeleteProjectRobot(ctx, exampleProject, exampleRobotID)

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestErrProjectNameNotProvided_Error(t *testing.T) {
	var e ErrProjectNameNotProvided

	assert.Equal(t, ErrProjectNameNotProvidedMsg, e.Error())
}

func TestErrProjectIDNotExists_Error(t *testing.T) {
	var e ErrProjectIDNotExists

	assert.Equal(t, ErrProjectIDNotExistsMsg, e.Error())
}

func TestErrProjectIllegalIDFormat_Error(t *testing.T) {
	var e ErrProjectIllegalIDFormat

	assert.Equal(t, ErrProjectIllegalIDFormatMsg, e.Error())
}

func TestErrProjectMetadataUndefined_Error(t *testing.T) {
	var e ErrProjectMetadataUndefined

	assert.Equal(t, ErrProjectMetadataUndefinedMsg, e.Error())
}

func TestErrProjectInternalErrors_Error(t *testing.T) {
	var e ErrProjectInternalErrors

	assert.Equal(t, ErrProjectInternalErrorsMsg, e.Error())
}

func TestErrProjectInvalidRequest_Error(t *testing.T) {
	var e ErrProjectInvalidRequest

	assert.Equal(t, ErrProjectInvalidRequestMsg, e.Error())
}

func TestErrProjectMemberIllegalFormat_Error(t *testing.T) {
	var e ErrProjectMemberIllegalFormat

	assert.Equal(t, ErrProjectMemberIllegalFormatMsg, e.Error())
}

func TestErrProjectMemberMismatch_Error(t *testing.T) {
	var e ErrProjectMemberMismatch

	assert.Equal(t, ErrProjectMemberMismatchMsg, e.Error())
}

func TestErrProjectMetadataAlreadyExists_Error(t *testing.T) {
	var e ErrProjectMetadataAlreadyExists

	assert.Equal(t, ErrProjectMetadataAlreadyExistsMsg, e.Error())
}

func TestErrProjectMismatch_Error(t *testing.T) {
	var e ErrProjectMismatch

	assert.Equal(t, ErrProjectMismatchMsg, e.Error())
}

func TestErrProjectNameAlreadyExists_Error(t *testing.T) {
	var e ErrProjectNameAlreadyExists

	assert.Equal(t, ErrProjectNameAlreadyExistsMsg, e.Error())
}

func TestErrProjectNoMemberProvided_Error(t *testing.T) {
	var e ErrProjectNoMemberProvided

	assert.Equal(t, ErrProjectNoMemberProvidedMsg, e.Error())
}

func TestErrProjectNoPermission_Error(t *testing.T) {
	var e ErrProjectNoPermission

	assert.Equal(t, ErrProjectNoPermissionMsg, e.Error())
}

func TestErrProjectNotFound_Error(t *testing.T) {
	var e ErrProjectNotFound

	assert.Equal(t, ErrProjectNotFoundMsg, e.Error())
}

func TestErrProjectNotProvided_Error(t *testing.T) {
	var e ErrProjectNotProvided

	assert.Equal(t, ErrProjectNotProvidedMsg, e.Error())
}

func TestErrProjectUnauthorized_Error(t *testing.T) {
	var e ErrProjectUnauthorized

	assert.Equal(t, ErrProjectUnauthorizedMsg, e.Error())
}

func TestErrProjectUnknownResource_Error(t *testing.T) {
	var e ErrProjectUnknownResource

	assert.Equal(t, ErrProjectUnknownResourceMsg, e.Error())
}

func TestErrProjectUserIsNoMember_Error(t *testing.T) {
	var e ErrProjectUserIsNoMember

	assert.Equal(t, ErrProjectUserIsNoMemberMsg, e.Error())
}
