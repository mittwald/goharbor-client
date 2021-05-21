// +build !integration

package project

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/go-openapi/strfmt"
	projectapi "github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client/project"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client/robotv1"
	modelv2 "github.com/mittwald/goharbor-client/v3/apiv2/model"

	"github.com/go-openapi/runtime"
	v2client "github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/legacyapi/client"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/legacyapi/client/products"
	"github.com/mittwald/goharbor-client/v3/apiv2/mocks"
	legacymodel "github.com/mittwald/goharbor-client/v3/apiv2/model/legacy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
	usr                         = &legacymodel.User{Username: exampleUser}
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
	exampleMetadataKey   = ProjectMetadataKeyEnableContentTrust
	exampleMetadataValue = "true"
)

func BuildLegacyClientWithMock(service *mocks.MockProductsClientService) *client.Harbor {
	return &client.Harbor{
		Products: service,
	}
}

func BuildV2ClientWithMocks(project *mocks.MockProjectClientService,
	robot *mocks.MockRobotv1ClientService) *v2client.Harbor {
	return &v2client.Harbor{
		Project: project,
		Robotv1: robot,
	}
}

func TestRESTClient_NewProject(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	postProjectParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	getProjectParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	p.On("CreateProject", postProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, nil)

	p.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{Payload: exampleProject}, nil)

	_, err := cl.NewProject(ctx, exampleProject.Name, &exampleStorageLimitPositive)

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_NewProject_UnlimitedStorage(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	postProjectParams := &projectapi.CreateProjectParams{
		Project: pReq3,
		Context: ctx,
	}

	getProjectParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	p.On("CreateProject", postProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, nil)

	p.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{Payload: exampleProject3}, nil)

	_, err := cl.NewProject(ctx, exampleProject3.Name, &exampleStorageLimitNegative)

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectNotFound(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	postProjectParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	getProjectParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	p.On("CreateProject", postProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, nil)

	p.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &runtime.APIError{Code: http.StatusNotFound})

	_, err := cl.NewProject(ctx, exampleProject.Name, &exampleStorageLimitPositive)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrProjectNotFound{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectIllegalIDFormat(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	postProjectParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	getProjectParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	p.On("CreateProject", postProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, nil)

	p.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{Payload: nil}, &ErrProjectIllegalIDFormat{})

	_, err := cl.NewProject(ctx, exampleProject.Name, &exampleStorageLimitPositive)

	if assert.Error(t, err) {
		assert.Equal(t, &ErrProjectIllegalIDFormat{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectUnauthorized(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	postProjectParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	p.On("CreateProject", postProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, &runtime.APIError{Code: http.StatusUnauthorized})

	_, err := cl.NewProject(ctx, exampleProject.Name, &exampleStorageLimitPositive)

	if assert.Error(t, err) {
		assert.Equal(t, &ErrProjectUnauthorized{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectNoPermission(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	postProjectParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	getProjectParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	p.On("CreateProject", postProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, nil)

	p.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{Payload: nil}, &runtime.APIError{Code: http.StatusForbidden})

	_, err := cl.NewProject(ctx, exampleProject.Name, &exampleStorageLimitPositive)

	if assert.Error(t, err) {
		assert.Equal(t, &ErrProjectNoPermission{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectUnknownResource(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	postProjectParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	p.On("CreateProject", postProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &ErrProjectUnknownResource{})

	_, err := cl.NewProject(ctx, exampleProject.Name, &exampleStorageLimitPositive)

	if assert.Error(t, err) {
		assert.Equal(t, &ErrProjectUnknownResource{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectInternalErrors(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	postProjectParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	p.On("CreateProject", postProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, &runtime.APIError{Code: http.StatusInternalServerError})

	_, err := cl.NewProject(ctx, exampleProject.Name, &exampleStorageLimitPositive)

	if assert.Error(t, err) {
		assert.Equal(t, &ErrProjectInternalErrors{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectIDNotExists(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	postProjectParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	p.On("CreateProject", postProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, &runtime.APIError{Code: http.StatusNotFound})

	_, err := cl.NewProject(ctx, exampleProject.Name, &exampleStorageLimitPositive)

	if assert.Error(t, err) {
		assert.Equal(t, &ErrProjectUnknownResource{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectNameAlreadyExists(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	postProjectParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	p.On("CreateProject", postProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, &projectapi.CreateProjectConflict{})

	_, err := cl.NewProject(ctx, exampleProject.Name, &exampleStorageLimitPositive)

	if assert.Error(t, err) {
		assert.Equal(t, &ErrProjectNameAlreadyExists{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectInvalidRequest(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	postProjectParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	p.On("CreateProject", postProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, &products.PostProjectsProjectIDMembersBadRequest{})

	_, err := cl.NewProject(ctx, exampleProject.Name, &exampleStorageLimitPositive)

	if assert.Error(t, err) {
		assert.Equal(t, &ErrProjectInvalidRequest{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_NewProject_ErrProjectInvalidRequest_2(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	postProjectParams := &projectapi.CreateProjectParams{
		Project: pReq,
		Context: ctx,
	}

	p.On("CreateProject", postProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.CreateProjectCreated{}, &products.PostProjectsProjectIDMembersBadRequest{})

	_, err := cl.NewProject(ctx, exampleProject.Name, &exampleStorageLimitPositive)

	if assert.Error(t, err) {
		assert.Equal(t, &ErrProjectInvalidRequest{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_Project_ErrProjectNotProvided(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildV2ClientWithMocks(p, nil)

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
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	getProjectParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	deleteProjectParams := &projectapi.DeleteProjectParams{
		ProjectNameOrID: strconv.Itoa(int(exampleProjectID)),
		Context:         ctx,
	}

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
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	n := "example-nonexistent"
	nonExistentProject := &modelv2.Project{Name: n}

	getProjectParams := &projectapi.GetProjectParams{
		ProjectNameOrID: nonExistentProject.Name,
		Context:         ctx,
	}

	p.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &ErrProjectNotFound{})

	err := cl.DeleteProject(ctx, nonExistentProject)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrProjectMismatch{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_DeleteProject_ErrProjectUnknownResource(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	n := "example-nonexistent"
	nonExistentProject := &modelv2.Project{Name: n}

	getProjectParams := &projectapi.GetProjectParams{
		ProjectNameOrID: n,
		Context:         ctx,
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
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	getProjectParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	p.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{
			Payload: exampleProject,
		}, nil)

	_, err := cl.GetProject(ctx, exampleProject.Name)

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_GetProject_ErrProjectNameNotProvided(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	_, err := cl.GetProject(ctx, "")

	if assert.Error(t, err) {
		assert.IsType(t, &ErrProjectNameNotProvided{}, err)
	}
}

func TestRESTClient_ListProjects(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildV2ClientWithMocks(p, nil)

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
	v2Client := BuildV2ClientWithMocks(p, nil)

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
	v2Client := BuildV2ClientWithMocks(p, nil)

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
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	getProjectParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	updateProjectParams := &projectapi.UpdateProjectParams{
		Project:         pReq3,
		ProjectNameOrID: strconv.Itoa(int(exampleProjectID)),
		Context:         ctx,
	}

	p.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{
			Payload: exampleProject,
		}, nil)

	p.On("UpdateProject",
		updateProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.UpdateProjectOK{}, nil)

	project, err := cl.GetProject(ctx, exampleProject.Name)

	assert.NoError(t, err)

	err = cl.UpdateProject(ctx, project, &exampleStorageLimitNegative)

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_UpdateProject_ErrProjectInternalErrors(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	getProjectParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	p.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{}, &runtime.APIError{
			OperationName: "",
			Response:      nil,
			Code:          500,
		})

	err := cl.UpdateProject(ctx, exampleProject, &exampleStorageLimitPositive)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrProjectInternalErrors{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_UpdateProject_ErrProjectMismatch(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	project2 := *exampleProject

	getProjectParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	p.On("GetProject", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{Payload: exampleProject}, nil)

	project2.ProjectID = 100
	err := cl.UpdateProject(ctx, &project2, &exampleStorageLimitPositive)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrProjectMismatch{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_AddProjectMember(t *testing.T) {
	p := &mocks.MockProjectClientService{}
	l := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(l)
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	getProjectsParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	getUserParams := &products.GetUsersParams{
		Context:  ctx,
		Username: &exampleUser,
	}

	postProjectsProjectIDMembersParams := &products.PostProjectsProjectIDMembersParams{
		ProjectID: exampleProjectID,
		ProjectMember: &legacymodel.ProjectMember{
			MemberUser: &legacymodel.UserEntity{
				Username: usr.Username,
			},
			MemberGroup: &legacymodel.UserGroup{
				GroupName: "",
				GroupType: 0,
				ID:        0,
			},
			RoleID: exampleUserRoleID,
		},
		Context: ctx,
	}

	p.On("GetProject", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{
			Payload: exampleProject,
		}, nil)

	l.On("GetUsers", getUserParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetUsersOK{
			Payload: []*legacymodel.User{{Username: exampleUser}},
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
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	getProjectsParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

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
	v2Client := BuildV2ClientWithMocks(p, nil)

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
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	getProjectsParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	getUserParams := &products.GetUsersParams{
		Context:  ctx,
		Username: &exampleUser,
	}

	p.On("GetProject", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{
			Payload: exampleProject,
		}, nil)

	l.On("GetUsers", getUserParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetUsersOK{
			Payload: []*legacymodel.User{{Username: "example-nonexistent"}},
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
	v2Client := BuildV2ClientWithMocks(nil, nil)

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
			Payload: []*legacymodel.ProjectMemberEntity{{}},
		}, nil)

	_, err := cl.ListProjectMembers(ctx, exampleProject)

	assert.NoError(t, err)

	l.AssertExpectations(t)
}

func TestRESTClient_ListProjectMembers_ErrProjectUnknownResource(t *testing.T) {
	l := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(l)
	v2Client := BuildV2ClientWithMocks(nil, nil)

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
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	e := ""

	getProjectsParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	getUserParams := &products.GetUsersParams{
		Context:  ctx,
		Username: &exampleUser,
	}

	postProjectsProjectIDMembersParams := &products.PostProjectsProjectIDMembersParams{
		ProjectID: exampleProjectID,
		ProjectMember: &legacymodel.ProjectMember{
			MemberUser: &legacymodel.UserEntity{
				Username: usr.Username,
			},
			MemberGroup: &legacymodel.UserGroup{
				GroupName: "",
				GroupType: 0,
				ID:        0,
			},
			RoleID: exampleUserRoleID,
		},
		Context: ctx,
	}

	p.On("GetProject", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{
			Payload: exampleProject,
		}, nil)

	l.On("GetUsers", getUserParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetUsersOK{
			Payload: []*legacymodel.User{{Username: exampleUser}},
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
		Role:      &legacymodel.RoleRequest{RoleID: exampleUserRoleID},
		Context:   ctx,
	}

	p.On("GetProject", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{
			Payload: exampleProject,
		}, nil)

	l.On("GetProjectsProjectIDMembers",
		&getProjectsProjectIDMembersParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetProjectsProjectIDMembersOK{
			Payload: []*legacymodel.ProjectMemberEntity{{
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
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	e := ""

	getProjectsParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	getProjectsProjectIDMembersParams := products.GetProjectsProjectIDMembersParams{
		Entityname: &e,
		ProjectID:  exampleProjectID,
		Context:    ctx,
	}

	p.On("GetProject", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{
			Payload: exampleProject,
		}, nil)

	l.On("GetProjectsProjectIDMembers",
		&getProjectsProjectIDMembersParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetProjectsProjectIDMembersOK{
			Payload: []*legacymodel.ProjectMemberEntity{{}},
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
	v2Client := BuildV2ClientWithMocks(p, nil)

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
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	e := ""

	getProjectsParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	getUserParams := &products.GetUsersParams{
		Context:  ctx,
		Username: &exampleUser,
	}

	postProjectsProjectIDMembersParams := &products.PostProjectsProjectIDMembersParams{
		ProjectID: exampleProjectID,
		ProjectMember: &legacymodel.ProjectMember{
			MemberUser: &legacymodel.UserEntity{
				Username: usr.Username,
			},
			MemberGroup: &legacymodel.UserGroup{
				GroupName: "",
				GroupType: 0,
				ID:        0,
			},
			RoleID: exampleUserRoleID,
		},
		Context: ctx,
	}

	p.On("GetProject", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{
			Payload: exampleProject,
		}, nil)

	l.On("GetUsers", getUserParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetUsersOK{
			Payload: []*legacymodel.User{{Username: exampleUser}},
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
			Payload: []*legacymodel.ProjectMemberEntity{{
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
	v2Client := BuildV2ClientWithMocks(p, nil)

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
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	getProjectsParams := &projectapi.GetProjectParams{
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	p.On("GetProject", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{Payload: nil}, &ErrProjectNotFound{})

	err := cl.DeleteProjectMember(ctx, exampleProject, usr)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrProjectMismatch{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_AddProjectMetadata(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	mPtr := "true"

	pReq.Metadata = &modelv2.ProjectMetadata{}
	pReq.Metadata.EnableContentTrust = &mPtr
	pReq.StorageLimit = nil

	updateProjectParams := &projectapi.UpdateProjectParams{
		Project:         pReq,
		ProjectNameOrID: strconv.Itoa(int(exampleProjectID)),
		Context:         ctx,
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
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	mPtr := "true"
	pReq.Metadata.EnableContentTrust = &mPtr
	pReq.StorageLimit = nil

	updateProjectParams := &projectapi.UpdateProjectParams{
		Project:         pReq,
		ProjectNameOrID: strconv.Itoa(int(exampleProjectID)),
		Context:         ctx,
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
	v2Client := BuildV2ClientWithMocks(p, nil)

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

	sPtr := "test"

	for _, k := range keys {
		getProjectsParams := &projectapi.GetProjectParams{
			ProjectNameOrID: strconv.Itoa(int(exampleProjectID)),
			Context:         ctx,
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

		val, err := cl.GetProjectMetadataValue(ctx, strconv.Itoa(int(exampleProjectID)), k)

		assert.Equal(t, val, sPtr)

		assert.NoError(t, err)

		p.AssertExpectations(t)
	}
}

func TestRESTClient_GetProjectMetadataValue_ValuesUndefined(t *testing.T) {
	t.Parallel()

	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	getProjectsParams := &projectapi.GetProjectParams{
		ProjectNameOrID: strconv.Itoa(int(exampleProjectID)),
		Context:         ctx,
	}

	var k MetadataKey
	t.Run("ProjectMetadataValueEnableContentTrustUndefined", func(t *testing.T) {
		k = ProjectMetadataKeyEnableContentTrust
		p.On("GetProject", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&projectapi.GetProjectOK{Payload: &modelv2.Project{
				Metadata:  &modelv2.ProjectMetadata{},
				ProjectID: int32(exampleProjectID),
			}}, nil)
		val, err := cl.GetProjectMetadataValue(ctx, strconv.Itoa(int(exampleProjectID)), k)
		if assert.Error(t, err) {
			assert.Equal(t, val, "")
			assert.IsType(t, &ErrProjectMetadataValueEnableContentTrustUndefined{}, err)
		}
		p.AssertExpectations(t)
	})
	t.Run("ProjectMetadataValueAutoScanUndefined", func(t *testing.T) {
		k = ProjectMetadataKeyAutoScan
		p.On("GetProject", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&projectapi.GetProjectOK{Payload: &modelv2.Project{
				Metadata:  &modelv2.ProjectMetadata{},
				ProjectID: int32(exampleProjectID),
			}}, nil)
		val, err := cl.GetProjectMetadataValue(ctx, strconv.Itoa(int(exampleProjectID)), k)
		if assert.Error(t, err) {
			assert.Equal(t, val, "")
			assert.IsType(t, &ErrProjectMetadataValueAutoScanUndefined{}, err)
		}
		p.AssertExpectations(t)
	})
	t.Run("ProjectMetadataValueSeverityUndefined", func(t *testing.T) {
		k = ProjectMetadataKeySeverity
		p.On("GetProject", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&projectapi.GetProjectOK{Payload: &modelv2.Project{
				Metadata:  &modelv2.ProjectMetadata{},
				ProjectID: int32(exampleProjectID),
			}}, nil)
		val, err := cl.GetProjectMetadataValue(ctx, strconv.Itoa(int(exampleProjectID)), k)
		if assert.Error(t, err) {
			assert.Equal(t, val, "")
			assert.IsType(t, &ErrProjectMetadataValueSeverityUndefined{}, err)
		}
		p.AssertExpectations(t)
	})
	t.Run("ProjectMetadataValueReuseSysCveAllowlistUndefined", func(t *testing.T) {
		k = ProjectMetadataKeyReuseSysCveAllowlist
		p.On("GetProject", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&projectapi.GetProjectOK{Payload: &modelv2.Project{
				Metadata:  &modelv2.ProjectMetadata{},
				ProjectID: int32(exampleProjectID),
			}}, nil)
		val, err := cl.GetProjectMetadataValue(ctx, strconv.Itoa(int(exampleProjectID)), k)
		if assert.Error(t, err) {
			assert.Equal(t, val, "")
			assert.IsType(t, &ErrProjectMetadataValueReuseSysCveAllowlistUndefined{}, err)
		}
		p.AssertExpectations(t)
	})
	t.Run("ProjectMetadataValueRetentionIDUndefined", func(t *testing.T) {
		k = ProjectMetadataKeyRetentionID
		p.On("GetProject", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&projectapi.GetProjectOK{Payload: &modelv2.Project{
				Metadata:  &modelv2.ProjectMetadata{},
				ProjectID: int32(exampleProjectID),
			}}, nil)
		val, err := cl.GetProjectMetadataValue(ctx, strconv.Itoa(int(exampleProjectID)), k)
		if assert.Error(t, err) {
			assert.Equal(t, val, "")
			assert.IsType(t, &ErrProjectMetadataValueRetentionIDUndefined{}, err)
		}
		p.AssertExpectations(t)
	})
	t.Run("ProjectMetadataValuePublicUndefined", func(t *testing.T) {
		k = ProjectMetadataKeyPublic
		p.On("GetProject", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&projectapi.GetProjectOK{Payload: &modelv2.Project{
				Metadata:  &modelv2.ProjectMetadata{},
				ProjectID: int32(exampleProjectID),
			}}, nil)
		val, err := cl.GetProjectMetadataValue(ctx, strconv.Itoa(int(exampleProjectID)), k)
		if assert.Error(t, err) {
			assert.Equal(t, val, "")
			assert.IsType(t, &ErrProjectMetadataValuePublicUndefined{}, err)
		}
		p.AssertExpectations(t)
	})
	t.Run("ProjectMetadataValuePreventVulUndefined", func(t *testing.T) {
		k = ProjectMetadataKeyPreventVul
		p.On("GetProject", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&projectapi.GetProjectOK{Payload: &modelv2.Project{
				Metadata:  &modelv2.ProjectMetadata{},
				ProjectID: int32(exampleProjectID),
			}}, nil)
		val, err := cl.GetProjectMetadataValue(ctx, strconv.Itoa(int(exampleProjectID)), k)
		if assert.Error(t, err) {
			assert.Equal(t, val, "")
			assert.IsType(t, &ErrProjectMetadataValuePreventVulUndefined{}, err)
		}
		p.AssertExpectations(t)
	})
}

func TestRESTClient_GetProjectMetadataValue_MetadataNil(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	getProjectsParams := &projectapi.GetProjectParams{
		ProjectNameOrID: strconv.Itoa(int(exampleProjectID)),
		Context:         ctx,
	}

	p.On("GetProject", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectapi.GetProjectOK{Payload: &modelv2.Project{
			Metadata:  nil,
			ProjectID: int32(exampleProjectID),
		}}, nil)

	_, err := cl.GetProjectMetadataValue(ctx, strconv.Itoa(int(exampleProjectID)), ProjectMetadataKeyRetentionID)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrProjectMetadataUndefined{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_GetProjectMetadataValue_ErrProjectUnknownResource(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildV2ClientWithMocks(p, nil)

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

	sPtr := "test"

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
			ProjectNameOrID: strconv.Itoa(int(exampleProjectID)),
			Context:         ctx,
		}

		p.On("GetProject", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&projectapi.GetProjectOK{Payload: nil}, &runtime.APIError{Code: http.StatusNotFound})

		_, err := cl.GetProjectMetadataValue(ctx, strconv.Itoa(int(exampleProject.ProjectID)), keys[i])

		if assert.Error(t, err) {
			assert.IsType(t, &ErrProjectUnknownResource{}, err)
		}

		p.AssertExpectations(t)
	}
}

func TestRESTClient_ListProjectMetadata(t *testing.T) {
	p := &mocks.MockProjectClientService{}

	legacyClient := BuildLegacyClientWithMock(nil)
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	sPtr := "true"

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
		ProjectNameOrID: strconv.Itoa(int(exampleProjectID)),
		Context:         ctx,
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
	v2Client := BuildV2ClientWithMocks(p, nil)

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

	mPtr := "true"
	pReq2.Metadata.EnableContentTrust = &mPtr
	updateProjectParams := &projectapi.UpdateProjectParams{
		Project:         pReq2,
		ProjectNameOrID: strconv.Itoa(int(exampleProject2.ProjectID)),
		Context:         ctx,
	}

	l.On("GetProjectsProjectIDMetadatasMetaName",
		getProjectsProjectIDMetadatasMetaName, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetProjectsProjectIDMetadatasMetaNameOK{Payload: &legacymodel.ProjectMetadata{
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
	v2Client := BuildV2ClientWithMocks(p, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	metaPtr := "true"

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
		Return(&products.GetProjectsProjectIDMetadatasMetaNameOK{Payload: &legacymodel.ProjectMetadata{
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
	v2Client := BuildV2ClientWithMocks(nil, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	metaPtr := "true"

	project := &modelv2.Project{
		ProjectID: exampleProject.ProjectID,
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
		ProjectID: 1,
		Context:   ctx,
	}

	l.On("GetProjectsProjectIDMetadatasMetaName",
		getProjectsProjectIDMetadatasMetaName, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetProjectsProjectIDMetadatasMetaNameOK{Payload: &legacymodel.ProjectMetadata{
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
	v2Client := BuildV2ClientWithMocks(nil, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	metaPtr := "true"

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

func TestRESTClient_ListProjectWebhookPolicies(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(nil, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	expectedWebhookPolicies := []*legacymodel.WebhookPolicy{
		{
			ID:        42,
			Name:      "example-policy",
			ProjectID: exampleProjectID,
		},
	}

	params := &products.GetProjectsProjectIDWebhookPoliciesParams{
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	p.On("GetProjectsProjectIDWebhookPolicies", params, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetProjectsProjectIDWebhookPoliciesOK{Payload: expectedWebhookPolicies}, nil)

	webhookPolicies, err := cl.ListProjectWebhookPolicies(ctx, exampleProject)

	assert.NoError(t, err)

	assert.Equal(t, expectedWebhookPolicies, webhookPolicies)

	p.AssertExpectations(t)
}

func TestRESTClient_AddProjectWebhookPolicy(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(nil, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	newPolicy := &legacymodel.WebhookPolicy{
		Enabled: true,
		Name:    "my-policy",
		Targets: []*legacymodel.WebhookTargetObject{{
			Address: "http://example-webhook.com",
		}},
		EventTypes: []string{
			"SCANNING_FAILED",
			"SCANNING_COMPLETED",
		},
	}

	params := &products.PostProjectsProjectIDWebhookPoliciesParams{
		ProjectID: exampleProjectID,
		Policy:    newPolicy,
		Context:   ctx,
	}

	p.On("PostProjectsProjectIDWebhookPolicies", params, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PostProjectsProjectIDWebhookPoliciesCreated{}, nil)

	err := cl.AddProjectWebhookPolicy(ctx, exampleProject, newPolicy)

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_UpdateProjectWebhookPolicy(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(nil, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	const examplePolicyID = 42

	updatePolicy := &legacymodel.WebhookPolicy{
		Enabled: false,
		Name:    "my-policy",
		Targets: []*legacymodel.WebhookTargetObject{{
			Address: "http://example-webhook.com",
		}},
		EventTypes: []string{
			"SCANNING_FAILED",
			"SCANNING_COMPLETED",
		},
	}

	params := &products.PutProjectsProjectIDWebhookPoliciesPolicyIDParams{
		ProjectID: exampleProjectID,
		PolicyID:  examplePolicyID,
		Policy:    updatePolicy,
		Context:   ctx,
	}

	p.On("PutProjectsProjectIDWebhookPoliciesPolicyID", params, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PutProjectsProjectIDWebhookPoliciesPolicyIDOK{}, nil)

	err := cl.UpdateProjectWebhookPolicy(ctx, exampleProject, examplePolicyID, updatePolicy)

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_DeleteProjectWebhookPolicy(t *testing.T) {
	p := &mocks.MockProductsClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildV2ClientWithMocks(nil, nil)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	const examplePolicyID = 42

	params := &products.DeleteProjectsProjectIDWebhookPoliciesPolicyIDParams{
		ProjectID: exampleProjectID,
		PolicyID:  examplePolicyID,
		Context:   ctx,
	}

	p.On("DeleteProjectsProjectIDWebhookPoliciesPolicyID", params, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.DeleteProjectsProjectIDWebhookPoliciesPolicyIDOK{}, nil)

	err := cl.DeleteProjectWebhookPolicy(ctx, exampleProject, examplePolicyID)

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
