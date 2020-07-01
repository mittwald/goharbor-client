// +build !integration

package project

import (
	"context"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/mittwald/goharbor-client/internal/api/v1.10.0/client"
	"github.com/mittwald/goharbor-client/internal/api/v1.10.0/client/products"
	"github.com/mittwald/goharbor-client/internal/api/v1.10.0/model"
	"github.com/mittwald/goharbor-client/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	authInfo            = runtimeclient.BasicAuth("foo", "bar")
	exampleProject      = "example-project"
	exampleCountLimit   = int64(1)
	exampleStorageLimit = int64(1)
	exampleProjectID    = int64(0)
	exampleUser         = "example-user"
	exampleUserRoleID   = int64(1)
)

func TestRESTClient_NewProject(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	pReq := &model.ProjectReq{
		CveWhitelist: nil,
		Metadata:     nil,
		ProjectName:  exampleProject,
		CountLimit:   exampleCountLimit,
		StorageLimit: exampleStorageLimit * 1024 * 1024,
	}

	postProjectParams := &products.PostProjectsParams{
		Project: pReq,
		Context: ctx,
	}

	getProjectParams := &products.GetProjectsParams{
		Name:    &pReq.ProjectName,
		Context: ctx,
	}

	p.On("PostProjects", postProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PostProjectsCreated{}, nil)

	p.On("GetProjects", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetProjectsOK{
			Payload: []*model.Project{{Name: exampleProject}}}, nil)

	_, err := cl.NewProject(ctx, exampleProject, int(exampleCountLimit), int(exampleStorageLimit))

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_DeleteProject(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	pReq := &model.ProjectReq{
		CveWhitelist: nil,
		Metadata:     nil,
		ProjectName:  exampleProject,
		CountLimit:   exampleCountLimit,
		StorageLimit: exampleStorageLimit * 1024 * 1024,
	}

	project := &model.Project{
		Name:      exampleProject,
		ProjectID: int32(exampleProjectID),
	}

	getProjectParams := &products.GetProjectsParams{
		Name:    &pReq.ProjectName,
		Context: ctx,
	}

	deleteProjectParams := &products.DeleteProjectsProjectIDParams{
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	p.On("GetProjects", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetProjectsOK{
			Payload: []*model.Project{{Name: exampleProject}}}, nil)

	p.On("DeleteProjectsProjectID", deleteProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.DeleteProjectsProjectIDOK{}, nil)

	err := cl.DeleteProject(ctx, project)

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_GetProject(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	pReq := &model.ProjectReq{
		CveWhitelist: nil,
		Metadata:     nil,
		ProjectName:  exampleProject,
		CountLimit:   exampleCountLimit,
		StorageLimit: exampleStorageLimit * 1024 * 1024,
	}

	getProjectParams := &products.GetProjectsParams{
		Name:    &pReq.ProjectName,
		Context: ctx,
	}

	p.On("GetProjects", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetProjectsOK{
			Payload: []*model.Project{{Name: exampleProject}}}, nil)

	_, err := cl.GetProject(ctx, exampleProject)

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_ListProjects(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	pReq := &model.ProjectReq{
		CveWhitelist: nil,
		Metadata:     nil,
		ProjectName:  exampleProject,
		CountLimit:   exampleCountLimit,
		StorageLimit: exampleStorageLimit * 1024 * 1024,
	}

	getProjectParams := &products.GetProjectsParams{
		Name:    &pReq.ProjectName,
		Context: ctx,
	}

	p.On("GetProjects", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetProjectsOK{
			Payload: []*model.Project{{Name: exampleProject}}}, nil)

	_, err := cl.ListProjects(ctx, exampleProject)

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_UpdateProject(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	pReq := &model.ProjectReq{
		CveWhitelist: nil,
		Metadata:     nil,
		ProjectName:  exampleProject,
		CountLimit:   exampleCountLimit,
		StorageLimit: exampleStorageLimit * 1024 * 1024,
	}

	getProjectParams := &products.GetProjectsParams{
		Name:    &pReq.ProjectName,
		Context: ctx,
	}

	putProjectParams := &products.PutProjectsProjectIDParams{
		Project:   pReq,
		ProjectID: exampleProjectID,
		Context:   ctx,
	}

	p.On("GetProjects", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetProjectsOK{
			Payload: []*model.Project{{Name: exampleProject}}}, nil)

	p.On("GetProjects", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetProjectsOK{
			Payload: []*model.Project{{Name: exampleProject}}}, nil)

	p.On("PutProjectsProjectID", putProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PutProjectsProjectIDOK{}, nil)

	project, err := cl.GetProject(ctx, exampleProject)

	assert.NoError(t, err)

	err = cl.UpdateProject(ctx, project, int(exampleCountLimit), int(exampleStorageLimit))

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_UpdateProject_IDMismatch(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	pReq := &model.ProjectReq{
		CveWhitelist: nil,
		Metadata:     nil,
		ProjectName:  exampleProject,
		CountLimit:   exampleCountLimit,
		StorageLimit: exampleStorageLimit * 1024 * 1024,
	}

	getProjectParams := &products.GetProjectsParams{
		Name:    &pReq.ProjectName,
		Context: ctx,
	}

	project := &model.Project{
		Name:      exampleProject,
		ProjectID: int32(exampleProjectID),
	}

	p.On("GetProjects", getProjectParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetProjectsOK{
			Payload: []*model.Project{{Name: exampleProject}}}, &ErrProjectMismatch{})

	err := cl.UpdateProject(ctx, project, int(exampleCountLimit), int(exampleStorageLimit))

	if assert.Error(t, err) {
		assert.Equal(t, &ErrProjectMismatch{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_AddProjectMember(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	project := &model.Project{
		Name:      exampleProject,
		ProjectID: int32(exampleProjectID),
	}

	usr := &model.User{Username: exampleUser}

	pReq := &model.ProjectReq{
		CveWhitelist: nil,
		Metadata:     nil,
		ProjectName:  exampleProject,
		CountLimit:   exampleCountLimit,
		StorageLimit: exampleStorageLimit * 1024 * 1024,
	}

	getProjectsParams := &products.GetProjectsParams{
		Name:    &pReq.ProjectName,
		Context: ctx,
	}

	getUserParams := &products.GetUsersParams{
		Context:  ctx,
		Username: &exampleUser}

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

	p.On("GetProjects", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetProjectsOK{
			Payload: []*model.Project{{Name: exampleProject}}}, nil)

	p.On("GetUsers", getUserParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetUsersOK{
			Payload: []*model.User{{Username: exampleUser}},
		}, nil)

	p.On("PostProjectsProjectIDMembers", postProjectsProjectIDMembersParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PostProjectsProjectIDMembersCreated{}, nil)

	err := cl.AddProjectMember(ctx, project, usr, 1)

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_ListProjectMembers(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	project := &model.Project{
		Name:      exampleProject,
		ProjectID: int32(exampleProjectID),
	}

	e := ""

	getProjectsProjectIDMembersParams := products.GetProjectsProjectIDMembersParams{
		Entityname: &e,
		ProjectID:  exampleProjectID,
		Context:    ctx,
	}

	p.On("GetProjectsProjectIDMembers", &getProjectsProjectIDMembersParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetProjectsProjectIDMembersOK{
			Payload: []*model.ProjectMemberEntity{{}},
		}, nil)

	_, err := cl.ListProjectMembers(ctx, project)

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_UpdateProjectMemberRole(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	e := ""

	pReq := &model.ProjectReq{
		CveWhitelist: nil,
		Metadata:     nil,
		ProjectName:  exampleProject,
		CountLimit:   exampleCountLimit,
		StorageLimit: exampleStorageLimit * 1024 * 1024,
	}

	getProjectsParams := &products.GetProjectsParams{
		Name:    &pReq.ProjectName,
		Context: ctx,
	}

	getUserParams := &products.GetUsersParams{
		Context:  ctx,
		Username: &exampleUser}

	usr := &model.User{Username: exampleUser}

	project := &model.Project{
		Name:      exampleProject,
		ProjectID: int32(exampleProjectID),
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

	p.On("GetProjects", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetProjectsOK{
			Payload: []*model.Project{{Name: exampleProject}}}, nil)

	p.On("GetUsers", getUserParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetUsersOK{
			Payload: []*model.User{{Username: exampleUser}},
		}, nil)

	p.On("PostProjectsProjectIDMembers", postProjectsProjectIDMembersParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PostProjectsProjectIDMembersCreated{}, nil)

	err := cl.AddProjectMember(ctx, project, usr, int(exampleUserRoleID))

	assert.NoError(t, err)

	getProjectsProjectIDMembersParams := products.GetProjectsProjectIDMembersParams{
		Entityname: &e,
		ProjectID:  exampleProjectID,
		Context:    ctx,
	}

	putProjectsProjectIDMembersMidParams := products.PutProjectsProjectIDMembersMidParams{
		Mid: 1,
		ProjectID: exampleProjectID,
		Role:      &model.RoleRequest{RoleID: exampleUserRoleID},
		Context:   ctx,
	}

	p.On("GetProjects", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetProjectsOK{
			Payload: []*model.Project{{Name: exampleProject}}}, nil)

	p.On("GetProjectsProjectIDMembers", &getProjectsProjectIDMembersParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetProjectsProjectIDMembersOK{
			Payload: []*model.ProjectMemberEntity{{
				EntityType: "u",
				EntityName: exampleUser,
				ID: exampleUserRoleID,
			}},
		}, nil)

	p.On("PutProjectsProjectIDMembersMid", &putProjectsProjectIDMembersMidParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PutProjectsProjectIDMembersMidOK{}, nil)

	err = cl.UpdateProjectMemberRole(ctx, project, usr, int(exampleUserRoleID))

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_UpdateProjectMemberRole_UserIsNoMember(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	e := ""

	pReq := &model.ProjectReq{
		CveWhitelist: nil,
		Metadata:     nil,
		ProjectName:  exampleProject,
		CountLimit:   exampleCountLimit,
		StorageLimit: exampleStorageLimit * 1024 * 1024,
	}

	getProjectsParams := &products.GetProjectsParams{
		Name:    &pReq.ProjectName,
		Context: ctx,
	}

	getProjectsProjectIDMembersParams := products.GetProjectsProjectIDMembersParams{
		Entityname: &e,
		ProjectID:  exampleProjectID,
		Context:    ctx,
	}

	project := &model.Project{
		Name:      exampleProject,
		ProjectID: int32(exampleProjectID),
	}

	usr := &model.User{Username: exampleUser}

	p.On("GetProjects", getProjectsParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetProjectsOK{
			Payload: []*model.Project{{Name: exampleProject}}}, nil)

	p.On("GetProjectsProjectIDMembers", &getProjectsProjectIDMembersParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetProjectsProjectIDMembersOK{
			Payload: []*model.ProjectMemberEntity{{}},
		}, nil)

	err := cl.UpdateProjectMemberRole(ctx, project, usr, int(exampleUserRoleID))

	if assert.Error(t, err) {
		assert.Equal(t, &ErrProjectUserIsNoMember{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_DeleteProjectMember(t *testing.T) {

}

func TestRESTClient_AddProjectMetadata(t *testing.T) {

}

func TestRESTClient_ListProjectMetadata(t *testing.T) {

}

func TestRESTClient_UpdateProjectMetadata(t *testing.T) {

}

func TestRESTClient_DeleteProjectMetadataValue(t *testing.T) {

}

func TestErrProjectIDNotExists_Error(t *testing.T) {
	var e ErrProjectIDNotExists

	assert.Equal(t, ErrProjectIDNotExistsMsg, e.Error())
}

func TestErrProjectIllegalIDFormat_Error(t *testing.T) {
	var e ErrProjectIllegalIDFormat

	assert.Equal(t, ErrProjectIllegalIDFormatMsg, e.Error())
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
