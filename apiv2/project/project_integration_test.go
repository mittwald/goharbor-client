// +build integration

package project

import (
	"context"
	"fmt"
	"net/url"
	"testing"

	"github.com/go-openapi/strfmt"
	v2client "github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/legacyapi/client"
	integrationtest "github.com/mittwald/goharbor-client/v3/apiv2/testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	uc "github.com/mittwald/goharbor-client/v3/apiv2/user"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	u, _                       = url.Parse(integrationtest.Host)
	legacySwaggerClient        = client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	v2SwaggerClient            = v2client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	authInfo                   = runtimeclient.BasicAuth(integrationtest.User, integrationtest.Password)
	storageLimitPositive int64 = 1
	storageLimitNegative int64 = -1
)

func TestAPIProjectNew(t *testing.T) {
	name := "test-project"

	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	p, err := c.NewProject(ctx, name, &storageLimitPositive)
	require.NoError(t, err)
	err = c.DeleteProject(ctx, p)

	require.NoError(t, err)
	assert.Equal(t, name, p.Name)
	assert.False(t, p.Deleted)
}

func TestAPIProjectNew_UnlimitedStorage(t *testing.T) {
	name := "test-project"

	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	p, err := c.NewProject(ctx, name, &storageLimitNegative)
	defer c.DeleteProject(ctx, p)

	require.NoError(t, err)
	assert.Equal(t, name, p.Name)
	assert.False(t, p.Deleted)
}

func TestAPIProjectGet(t *testing.T) {
	name := "test-project"

	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	p, err := c.NewProject(ctx, name, &storageLimitPositive)
	require.NoError(t, err)
	defer c.DeleteProject(ctx, p)

	p2, err := c.GetProject(ctx, name)
	require.NoError(t, err)
	assert.Equal(t, p, p2)
}

func TestAPIProjectDelete(t *testing.T) {
	name := "test-project"
	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	p, err := c.NewProject(ctx, name, &storageLimitPositive)
	require.NoError(t, err)

	err = c.DeleteProject(ctx, p)
	require.NoError(t, err)

	p, err = c.GetProject(ctx, name)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrProjectNotFound{}, err)
	}
}

func TestAPIProjectList(t *testing.T) {
	namePrefix := "test-project"
	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("%s-%d", namePrefix, i)
		p, err := c.NewProject(ctx, name, &storageLimitPositive)
		require.NoError(t, err)
		defer func() {
			err := c.DeleteProject(ctx, p)
			if err != nil {
				panic("error in cleanup routine: " + err.Error())
			}
		}()
	}

	projects, err := c.ListProjects(ctx, namePrefix)
	require.NoError(t, err)
	assert.Len(t, projects, 10)
	for _, v := range projects {
		assert.Contains(t, v.Name, namePrefix)
	}
}

func TestAPIProjectUpdate(t *testing.T) {
	name := "test-project"
	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	p, err := c.NewProject(ctx, name, &storageLimitPositive)
	require.NoError(t, err)
	defer c.DeleteProject(ctx, p)

	var sPtr *string

	require.Equal(t, sPtr, p.Metadata.AutoScan)

	mPtr := "true"
	p.Metadata.AutoScan = &mPtr
	err = c.UpdateProject(ctx, p, &storageLimitPositive)
	require.NoError(t, err)
	p2, err := c.GetProject(ctx, name)
	require.NoError(t, err)

	assert.NotEqual(t, p, p2)
}

func TestAPIProjectUserMemberAdd(t *testing.T) {
	projectName := "test-project"
	memberUsername := "foobar"
	memberEmail := "foo@bar.com"
	memberRealname := "Foo Bar"
	memberPassword := "1VerySeriousPassword"
	memberComments := "Some comments"

	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	p, err := c.NewProject(ctx, projectName, &storageLimitPositive)
	defer c.DeleteProject(ctx, p)
	require.NoError(t, err)

	uc := uc.NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	u, err := uc.NewUser(ctx, memberUsername, memberEmail, memberRealname, memberPassword, memberComments)
	require.NoError(t, err)

	defer uc.DeleteUser(ctx, u)

	err = c.AddProjectMember(ctx, p, u, 1)
	require.NoError(t, err)
}

func TestAPIProjectMemberList(t *testing.T) {
	projectName := "test-project"
	memberUsername := "foobar"
	memberEmail := "foo@bar.com"
	memberRealname := "Foo Bar"
	memberPassword := "1VerySeriousPassword"
	memberComments := "Some comments"

	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	p, err := c.NewProject(ctx, projectName, &storageLimitPositive)
	defer c.DeleteProject(ctx, p)
	require.NoError(t, err)

	uc := uc.NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	u, err := uc.NewUser(ctx, memberUsername, memberEmail, memberRealname, memberPassword, memberComments)
	require.NoError(t, err)
	defer uc.DeleteUser(ctx, u)

	members, err := c.ListProjectMembers(ctx, p)
	require.NoError(t, err)

	assert.Len(t, members, 1)

	err = c.AddProjectMember(ctx, p, u, 1)
	require.NoError(t, err)

	members, err = c.ListProjectMembers(ctx, p)
	require.NoError(t, err)

	assert.Len(t, members, 2)
}

func TestAPIProjectUserMemberUpdate(t *testing.T) {
	projectName := "test-project"
	memberUsername := "foobar"
	memberEmail := "foo@bar.com"
	memberRealname := "Foo Bar"
	memberPassword := "1VerySeriousPassword"
	memberComments := "Some comments"

	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	p, err := c.NewProject(ctx, projectName, &storageLimitPositive)
	defer c.DeleteProject(ctx, p)
	require.NoError(t, err)

	uc := uc.NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	u, err := uc.NewUser(ctx, memberUsername, memberEmail, memberRealname, memberPassword, memberComments)
	require.NoError(t, err)
	defer uc.DeleteUser(ctx, u)

	err = c.AddProjectMember(ctx, p, u, 1)
	require.NoError(t, err)

	err = c.UpdateProjectMemberRole(ctx, p, u, 2)
	require.NoError(t, err)

	members, err := c.ListProjectMembers(ctx, p)
	require.NoError(t, err)

	found := false
	for _, v := range members {
		if v.EntityType == "u" && v.ProjectID == int64(p.ProjectID) && v.EntityName == u.Username {
			assert.Equal(t, int64(2), v.RoleID)
			found = true
		}
	}

	if !found {
		t.Error("did not find member in project")
	}
}

func TestAPIProjectUserMemberDelete(t *testing.T) {
	projectName := "test-project"
	memberUsername := "foobar"
	memberEmail := "foo@bar.com"
	memberRealname := "Foo Bar"
	memberPassword := "1VerySeriousPassword"
	memberComments := "Some comments"

	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	p, err := c.NewProject(ctx, projectName, &storageLimitPositive)
	defer c.DeleteProject(ctx, p)
	require.NoError(t, err)

	uc := uc.NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	u, err := uc.NewUser(ctx, memberUsername, memberEmail, memberRealname, memberPassword, memberComments)
	require.NoError(t, err)
	defer uc.DeleteUser(ctx, u)

	err = c.AddProjectMember(ctx, p, u, 1)
	require.NoError(t, err)

	err = c.DeleteProjectMember(ctx, p, u)
	require.NoError(t, err)

	members, err := c.ListProjectMembers(ctx, p)
	require.NoError(t, err)

	found := false
	for _, v := range members {
		if v.EntityType == "u" && v.ProjectID == int64(p.ProjectID) && v.EntityName == u.Username {
			assert.Equal(t, int64(2), v.RoleID)
			found = true
		}
	}

	assert.False(t, found)
}

func TestAPIProjectMetadataAdd(t *testing.T) {
	projectName := "test-project"

	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	p, err := c.NewProject(ctx, projectName, &storageLimitPositive)
	defer c.DeleteProject(ctx, p)
	require.NoError(t, err)

	err = c.AddProjectMetadata(ctx, p, ProjectMetadataKeyAutoScan, "true")
	require.NoError(t, err)
	err = c.AddProjectMetadata(ctx, p, ProjectMetadataKeyEnableContentTrust, "true")
	require.NoError(t, err)
	err = c.AddProjectMetadata(ctx, p, ProjectMetadataKeyPreventVul, "true")
	require.NoError(t, err)
	err = c.AddProjectMetadata(ctx, p, ProjectMetadataKeyReuseSysCveAllowlist, "true")
	require.NoError(t, err)
	err = c.AddProjectMetadata(ctx, p, ProjectMetadataKeySeverity, "medium")
	require.NoError(t, err)
	err = c.AddProjectMetadata(ctx, p, ProjectMetadataKeyRetentionID, "1")
}

func TestAPIProjectMetadataGet(t *testing.T) {
	projectName := "test-project"

	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	p, err := c.NewProject(ctx, projectName, &storageLimitPositive)
	defer c.DeleteProject(ctx, p)
	require.NoError(t, err)

	m, err := c.GetProjectMetadataValue(ctx, fmt.Sprint(p.ProjectID), ProjectMetadataKeyPublic)
	require.NoError(t, err)

	assert.Equal(t, "false", m)
}

func TestAPIProjectMetadataGetInvalidKey(t *testing.T) {
	projectName := "test-project"

	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	p, err := c.NewProject(ctx, projectName, &storageLimitPositive)
	defer c.DeleteProject(ctx, p)
	require.NoError(t, err)

	m, err := c.GetProjectMetadataValue(ctx, fmt.Sprint(p.ProjectID), "foobar")

	if assert.Error(t, err) {
		assert.Equal(t, "invalid request", err.Error())
		assert.IsType(t, &ErrProjectInvalidRequest{}, err)
	}

	assert.Equal(t, "", m)
}

func TestAPIProjectMetadataList(t *testing.T) {
	projectName := "test-project"

	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	p, err := c.NewProject(ctx, projectName, &storageLimitPositive)
	defer c.DeleteProject(ctx, p)
	require.NoError(t, err)

	err = c.AddProjectMetadata(ctx, p, ProjectMetadataKeyAutoScan, "true")
	require.NoError(t, err)
	err = c.AddProjectMetadata(ctx, p, ProjectMetadataKeyEnableContentTrust, "true")
	require.NoError(t, err)
	err = c.AddProjectMetadata(ctx, p, ProjectMetadataKeyPreventVul, "true")
	require.NoError(t, err)
	err = c.AddProjectMetadata(ctx, p, ProjectMetadataKeyReuseSysCveAllowlist, "true")
	require.NoError(t, err)
	err = c.AddProjectMetadata(ctx, p, ProjectMetadataKeySeverity, "medium")
	require.NoError(t, err)
	err = c.AddProjectMetadata(ctx, p, ProjectMetadataKeyRetentionID, "1")

	m, err := c.ListProjectMetadata(ctx, p)
	require.NoError(t, err)

	sPtrTrue := "true"
	sPtrMedium := "medium"
	sPtrRID := "1"

	assert.Equal(t, &sPtrTrue, m.AutoScan)
	assert.Equal(t, &sPtrTrue, m.EnableContentTrust)
	assert.Equal(t, &sPtrTrue, m.PreventVul)
	assert.Equal(t, &sPtrTrue, m.ReuseSysCveAllowlist)
	assert.Equal(t, &sPtrMedium, m.Severity)
	assert.Equal(t, &sPtrRID, m.RetentionID)
}

/* TODO: re-introduce this, once https://github.com/goharbor/harbor/issues/12570 is resolved.
func TestAPIProjectMetadataUpdate(t *testing.T) {
	projectName := "test-project"

	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	p, err := c.NewProject(ctx, projectName, 3, 3)
	defer c.DeleteProject(ctx, p)
	require.NoError(t, err)

	err = c.AddProjectMetadata(ctx, p, ProjectMetadataKeyAutoScan, "true")
	require.NoError(t, err)

	err = c.UpdateProjectMetadata(ctx, p, ProjectMetadataKeyAutoScan, "false")

	k, err := c.GetProjectMetadataValue(ctx, p, ProjectMetadataKeyAutoScan)
	require.NoError(t, err)

	assert.Equal(t, "false", k)
}*/

/* TODO: re-introduce this, once https://github.com/goharbor/harbor/issues/12570 is resolved.
func TestAPIProjectMetadataDelete(t *testing.T) {
	projectName := "test-project"

	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	p, err := c.NewProject(ctx, projectName, 3, 3)
	defer c.DeleteProject(ctx, p)
	require.NoError(t, err)

	err = c.AddProjectMetadata(ctx, p, ProjectMetadataKeyAutoScan, "true")
	require.NoError(t, err)

	err = c.DeleteProjectMetadataValue(ctx, p, ProjectMetadataKeyAutoScan)

	require.NoError(t, err)

	m, err := c.GetProjectMetadataValue(ctx, p, ProjectMetadataKeyAutoScan)

	if assert.Error(t, err) {
		assert.Equal(t, "resource unknown", err.Error())
	}
	assert.Equal(t, "", m)
}*/
