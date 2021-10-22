//go:build !integration

package member

import (
	"context"
	"net/http"
	"testing"

	"github.com/go-openapi/runtime"
	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/member"
	"github.com/mittwald/goharbor-client/v4/apiv2/mocks"
	modelv2 "github.com/mittwald/goharbor-client/v4/apiv2/model"
	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/common"
	unittesting "github.com/mittwald/goharbor-client/v4/apiv2/pkg/testing"
)

var (
	authInfo         = runtimeclient.BasicAuth("foo", "bar")
	exampleProjectID = int64(1)
	exampleProject   = &modelv2.Project{Name: "example-project", ProjectID: int32(exampleProjectID)}
	ctx              = context.Background()
	exampleMember    = modelv2.ProjectMember{
		MemberUser: &modelv2.UserEntity{},
		RoleID:     1,
	}
)

func APIandMockClientsForTests() (*RESTClient, *unittesting.MockClients) {
	desiredMockClients := &unittesting.MockClients{
		Project: mocks.MockProjectClientService{},
	}

	v2Client := unittesting.BuildV2ClientWithMocks(desiredMockClients)

	cl := NewClient(v2Client, &unittesting.DefaultOpts, authInfo)

	return cl, desiredMockClients
}

func TestRESTClient_AddProjectMember(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	addParams := &member.CreateProjectMemberParams{
		ProjectMember:   &exampleMember,
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	mockClient.Member.On("CreateProjectMember", addParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&member.CreateProjectMemberCreated{}, nil)

	err := apiClient.AddProjectMember(ctx, exampleProject.Name, &exampleMember)

	require.NoError(t, err)
	mockClient.Member.AssertExpectations(t)
}

func TestRESTClient_AddProjectMember_ErrProjectNoMemberProvided(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()
	err := apiClient.AddProjectMember(ctx, exampleProject.Name, nil)
	require.Error(t, err)
	require.ErrorIs(t, &common.ErrProjectNoMemberProvided{}, err)
	mockClient.Member.AssertExpectations(t)
}

func TestRESTClient_ListProjectMembers(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	listParams := &member.ListProjectMembersParams{
		Entityname:      &exampleMember.MemberUser.Username,
		PageSize:        &apiClient.Options.PageSize,
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	mockClient.Member.On("ListProjectMembers", listParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&member.ListProjectMembersOK{}, nil)

	_, err := apiClient.ListProjectMembers(ctx, exampleProject.Name, exampleMember.MemberUser.Username)

	require.NoError(t, err)

	mockClient.Member.AssertExpectations(t)
}

func TestRESTClient_ListProjectMembers_ErrProjectUnknownResource(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	listParams := &member.ListProjectMembersParams{
		Entityname:      &exampleMember.MemberUser.Username,
		PageSize:        &apiClient.Options.PageSize,
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	mockClient.Member.On("ListProjectMembers", listParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &runtime.APIError{Code: http.StatusNotFound})

	_, err := apiClient.ListProjectMembers(ctx, exampleProject.Name, exampleMember.MemberUser.Username)

	require.Error(t, err)
	require.ErrorIs(t, &common.ErrNotFound{}, err)

	mockClient.Member.AssertExpectations(t)
}

func TestRESTClient_UpdateProjectMember(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	q := ""
	listParams := &member.ListProjectMembersParams{
		Entityname:      &q,
		PageSize:        &apiClient.Options.PageSize,
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	updateParams := &member.UpdateProjectMemberParams{
		Mid:             exampleMember.MemberUser.UserID,
		ProjectNameOrID: exampleProject.Name,
		Role:            &modelv2.RoleRequest{RoleID: 1},
		Context:         ctx,
	}

	mockClient.Member.On("ListProjectMembers", listParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&member.ListProjectMembersOK{Payload: []*modelv2.ProjectMemberEntity{{
			EntityID:   exampleMember.MemberUser.UserID,
			EntityName: exampleMember.MemberUser.Username,
			EntityType: EntityTypeUser.String(),
			ID:         exampleMember.MemberUser.UserID,
			ProjectID:  exampleProjectID,
			RoleID:     exampleMember.RoleID,
			RoleName:   "projectAdmin",
		}}}, nil)

	mockClient.Member.On("UpdateProjectMember", updateParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&member.UpdateProjectMemberOK{}, nil)

	err := apiClient.UpdateProjectMember(ctx, exampleProject.Name, &exampleMember)
	require.NoError(t, err)

	mockClient.Member.AssertExpectations(t)
}

func TestRESTClient_DeleteProjectMember(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	q := ""
	listParams := &member.ListProjectMembersParams{
		Entityname:      &q,
		PageSize:        &apiClient.Options.PageSize,
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	mockClient.Member.On("ListProjectMembers", listParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&member.ListProjectMembersOK{Payload: []*modelv2.ProjectMemberEntity{{
			EntityID:   exampleMember.MemberUser.UserID,
			EntityName: exampleMember.MemberUser.Username,
			EntityType: EntityTypeUser.String(),
			ID:         exampleMember.MemberUser.UserID,
			ProjectID:  exampleProjectID,
			RoleID:     exampleMember.RoleID,
			RoleName:   "projectAdmin",
		}}}, nil)

	deleteParams := &member.DeleteProjectMemberParams{
		Mid:             exampleMember.MemberUser.UserID,
		ProjectNameOrID: exampleProject.Name,
		Context:         ctx,
	}

	mockClient.Member.On("DeleteProjectMember", deleteParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&member.DeleteProjectMemberOK{}, nil)

	err := apiClient.DeleteProjectMember(ctx, exampleProject.Name, &exampleMember)
	require.NoError(t, err)
	mockClient.Member.AssertExpectations(t)
}
