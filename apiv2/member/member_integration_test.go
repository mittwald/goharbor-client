//go:build integration

package member

import (
	"context"
	"net/url"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/require"

	v2client "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client"
	modelv2 "github.com/mittwald/goharbor-client/v4/apiv2/model"
	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/config"
	integrationtest "github.com/mittwald/goharbor-client/v4/apiv2/pkg/testing"
	"github.com/mittwald/goharbor-client/v4/apiv2/project"
	"github.com/mittwald/goharbor-client/v4/apiv2/user"
)

var (
	u, _                       = url.Parse(integrationtest.Host)
	authInfo                   = runtimeclient.BasicAuth(integrationtest.User, integrationtest.Password)
	v2SwaggerClient            = v2client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	opts                       = config.Options{}
	defaultOpts                = opts.Defaults()
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
	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)

	projectClient := project.NewClient(v2SwaggerClient, defaultOpts, authInfo)

	p, err := projectClient.NewProject(ctx, projectName, &storageLimitPositive)
	defer projectClient.DeleteProject(ctx, p)
	require.NoError(t, err)

	userClient := user.NewClient(v2SwaggerClient, defaultOpts, authInfo)

	u, err := userClient.NewUser(ctx, memberUsername, memberEmail, memberRealname, memberPassword, memberComments)
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
	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)

	projectClient := project.NewClient(v2SwaggerClient, defaultOpts, authInfo)
	p, err := projectClient.NewProject(ctx, projectName, &storageLimitPositive)
	defer projectClient.DeleteProject(ctx, p)

	userClient := user.NewClient(v2SwaggerClient, defaultOpts, authInfo)

	u, err := userClient.NewUser(ctx, memberUsername, memberEmail, memberRealname, memberPassword, memberComments)
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
	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)

	projectClient := project.NewClient(v2SwaggerClient, defaultOpts, authInfo)

	p, err := projectClient.NewProject(ctx, projectName, &storageLimitPositive)
	defer projectClient.DeleteProject(ctx, p)
	require.NoError(t, err)

	userClient := user.NewClient(v2SwaggerClient, defaultOpts, authInfo)

	u, err := userClient.NewUser(ctx, memberUsername, memberEmail, memberRealname, memberPassword, memberComments)
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
	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)

	projectClient := project.NewClient(v2SwaggerClient, defaultOpts, authInfo)

	p, err := projectClient.NewProject(ctx, projectName, &storageLimitPositive)
	defer projectClient.DeleteProject(ctx, p)
	require.NoError(t, err)

	userClient := user.NewClient(v2SwaggerClient, defaultOpts, authInfo)

	u, err := userClient.NewUser(ctx, memberUsername, memberEmail, memberRealname, memberPassword, memberComments)
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
