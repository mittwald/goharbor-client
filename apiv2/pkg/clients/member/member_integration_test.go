//go:build integration

package member

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/clients/project"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/clients/user"
	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"

	modelv2 "github.com/mittwald/goharbor-client/v5/apiv2/model"
)

var (
	storageLimitPositive int64 = 1
	projectName                = "test-project"
	memberUsername             = "foobar"
	memberEmail                = "foo@bar.com"
	memberRealname             = "Foo Bar"
	memberPassword             = "1VerySeriousPassword"
	memberComments             = "Some comments"
)

func TestAPIProjectUserMemberAdd(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	projectClient := project.NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := projectClient.NewProject(ctx, &modelv2.ProjectReq{
		ProjectName:  projectName,
		StorageLimit: &storageLimitPositive,
	})
	require.NoError(t, err)

	p, err := projectClient.GetProject(ctx, projectName)
	require.NoError(t, err)

	defer projectClient.DeleteProject(ctx, p.Name)

	userClient := user.NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err = userClient.NewUser(ctx, memberUsername, memberEmail, memberRealname, memberPassword, memberComments)
	require.NoError(t, err)

	u, err := userClient.GetUserByName(ctx, memberUsername)
	require.NoError(t, err)

	defer userClient.DeleteUser(ctx, u.UserID)

	err = c.AddProjectMember(ctx, p.Name, &modelv2.ProjectMember{
		RoleID: 1,
		MemberUser: &modelv2.UserEntity{
			UserID:   u.UserID,
			Username: u.Username,
		},
	})
	require.NoError(t, err)
}

func TestAPIProjectMemberList(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	projectClient := project.NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)
	err := projectClient.NewProject(ctx, &modelv2.ProjectReq{
		ProjectName:  projectName,
		StorageLimit: &storageLimitPositive,
	})
	require.NoError(t, err)

	p, err := projectClient.GetProject(ctx, projectName)
	require.NoError(t, err)
	defer projectClient.DeleteProject(ctx, projectName)

	userClient := user.NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err = userClient.NewUser(ctx, memberUsername, memberEmail, memberRealname, memberPassword, memberComments)
	require.NoError(t, err)

	u, err := userClient.GetUserByName(ctx, memberUsername)
	require.NoError(t, err)

	defer userClient.DeleteUser(ctx, u.UserID)

	members, err := c.ListProjectMembers(ctx, p.Name, u.Username)
	require.NoError(t, err)

	require.Len(t, members, 0)

	err = c.AddProjectMember(ctx, p.Name, &modelv2.ProjectMember{
		MemberUser: &modelv2.UserEntity{
			UserID:   u.UserID,
			Username: u.Username,
		},
		RoleID: 1,
	})

	require.NoError(t, err)

	members, err = c.ListProjectMembers(ctx, p.Name, u.Username)
	require.NoError(t, err)

	require.Len(t, members, 1)
}

func TestAPIProjectUserMemberUpdate(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	projectClient := project.NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := projectClient.NewProject(ctx, &modelv2.ProjectReq{
		ProjectName:  projectName,
		StorageLimit: &storageLimitPositive,
	})
	require.NoError(t, err)

	p, err := projectClient.GetProject(ctx, projectName)
	require.NoError(t, err)
	defer projectClient.DeleteProject(ctx, projectName)

	userClient := user.NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err = userClient.NewUser(ctx, memberUsername, memberEmail, memberRealname, memberPassword, memberComments)
	require.NoError(t, err)

	u, err := userClient.GetUserByName(ctx, memberUsername)
	require.NoError(t, err)

	defer userClient.DeleteUser(ctx, u.UserID)

	err = c.AddProjectMember(ctx, p.Name, &modelv2.ProjectMember{
		MemberUser: &modelv2.UserEntity{
			UserID:   u.UserID,
			Username: u.Username,
		},
		RoleID: 1,
	})
	require.NoError(t, err)

	err = c.UpdateProjectMember(ctx, p.Name, &modelv2.ProjectMember{
		MemberUser: &modelv2.UserEntity{
			UserID:   u.UserID,
			Username: u.Username,
		},
		RoleID: 2,
	})
	require.NoError(t, err)

	members, err := c.ListProjectMembers(ctx, p.Name, u.Username)
	require.NoError(t, err)

	for _, v := range members {
		if v.EntityType == "u" && v.ProjectID == int64(p.ProjectID) && v.EntityName == u.Username {
			require.Equal(t, int64(2), v.RoleID)
		}
	}
}

func TestAPIProjectUserMemberDelete(t *testing.T) {
	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	projectClient := project.NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err := projectClient.NewProject(ctx, &modelv2.ProjectReq{
		ProjectName:  projectName,
		StorageLimit: &storageLimitPositive,
	})
	require.NoError(t, err)

	p, err := projectClient.GetProject(ctx, projectName)
	require.NoError(t, err)
	defer projectClient.DeleteProject(ctx, projectName)

	userClient := user.NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	err = userClient.NewUser(ctx, memberUsername, memberEmail, memberRealname, memberPassword, memberComments)
	require.NoError(t, err)

	u, err := userClient.GetUserByName(ctx, memberUsername)
	require.NoError(t, err)

	defer userClient.DeleteUser(ctx, u.UserID)

	mem := &modelv2.ProjectMember{
		MemberUser: &modelv2.UserEntity{
			UserID:   u.UserID,
			Username: u.Username,
		},
		RoleID: 1,
	}

	err = c.AddProjectMember(ctx, p.Name, mem)
	require.NoError(t, err)

	err = c.DeleteProjectMember(ctx, p.Name, mem)
	require.NoError(t, err)

	members, err := c.ListProjectMembers(ctx, p.Name, u.Username)
	require.NoError(t, err)

	found := false
	for _, v := range members {
		if v.EntityType == "u" && v.ProjectID == int64(p.ProjectID) && v.EntityName == u.Username {
			require.Equal(t, int64(2), v.RoleID)
			found = true
		}
	}

	require.False(t, found)
}
