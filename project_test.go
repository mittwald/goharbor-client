package goharborclient

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAPIProjectNew(t *testing.T) {
	if !*integrationTest {
		t.Skip()
	}

	name := "test-project"

	ctx := context.Background()
	c := NewClient(host, defaultUser, defaultPassword)

	p, err := c.Projects().NewProject(ctx, name, 3, 3)
	defer c.Projects().Delete(ctx, p)

	require.NoError(t, err)
	assert.Equal(t, name, p.Name)
	assert.False(t, p.Deleted)
}

func TestAPIProjectGet(t *testing.T) {
	if !*integrationTest {
		t.Skip()
	}

	name := "test-project"

	ctx := context.Background()
	c := NewClient(host, defaultUser, defaultPassword)

	p, err := c.Projects().NewProject(ctx, name, 3, 3)
	require.NoError(t, err)
	defer c.Projects().Delete(ctx, p)

	p2, err := c.Projects().Get(ctx, name)
	require.NoError(t, err)
	assert.Equal(t, p, p2)
}

func TestAPIProjectDelete(t *testing.T) {
	if !*integrationTest {
		t.Skip()
	}

	name := "test-project"
	ctx := context.Background()
	c := NewClient(host, defaultUser, defaultPassword)

	p, err := c.Projects().NewProject(ctx, name, 3, 3)
	require.NoError(t, err)

	err = c.Projects().Delete(ctx, p)
	require.NoError(t, err)

	p, err = c.Projects().Get(ctx, name)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrProjectNotFound{}, err)
	}
}

func TestAPIProjectList(t *testing.T) {
	if !*integrationTest {
		t.Skip()
	}

	namePrefix := "test-project"
	ctx := context.Background()
	c := NewClient(host, defaultUser, defaultPassword)

	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("%s-%d", namePrefix, i)
		p, err := c.Projects().NewProject(ctx, name, 3, 3)
		require.NoError(t, err)
		defer func() {
			err := c.Projects().Delete(ctx, p)
			if err != nil {
				panic("error in cleanup routine: " + err.Error())
			}
		}()
	}

	projects, err := c.Projects().List(ctx, namePrefix)
	require.NoError(t, err)
	assert.Len(t, projects, 10)
	for _, v := range projects {
		assert.Contains(t, v.Name, namePrefix)
	}
}

func TestAPIProjectUpdate(t *testing.T) {
	if !*integrationTest {
		t.Skip()
	}

	name := "test-project"
	ctx := context.Background()
	c := NewClient(host, defaultUser, defaultPassword)

	p, err := c.Projects().NewProject(ctx, name, 3, 3)
	require.NoError(t, err)
	defer c.Projects().Delete(ctx, p)
	require.Equal(t, "", p.Metadata.AutoScan)

	p.Metadata.AutoScan = "true"
	err = c.Projects().Update(ctx, p, 2, 2)
	require.NoError(t, err)
	p2, err := c.Projects().Get(ctx, name)
	require.NoError(t, err)
	assert.Equal(t, p, p2)
}

func TestAPIProjectUserMemberAdd(t *testing.T) {
	if !*integrationTest {
		t.Skip()
	}

	projectName := "test-project"
	memberUsername := "foobar"
	memberEmail := "foo@bar.com"
	memberRealname := "Foo Bar"
	memberPassword := "1VerySeriousPassword"
	memberComments := "Some comments"

	ctx := context.Background()
	c := NewClient(host, defaultUser, defaultPassword)

	p, err := c.Projects().NewProject(ctx, projectName, 3, 3)
	defer c.Projects().Delete(ctx, p)
	require.NoError(t, err)

	u, err := c.Users().NewUser(ctx, memberUsername, memberEmail, memberRealname, memberPassword, memberComments)
	require.NoError(t, err)
	defer c.Users().Delete(ctx, u)

	err = c.Projects().AddUserMember(ctx, p, u, 1)
	require.NoError(t, err)

}

func TestAPIProjectMemberList(t *testing.T) {
	if !*integrationTest {
		t.Skip()
	}

	projectName := "test-project"
	memberUsername := "foobar"
	memberEmail := "foo@bar.com"
	memberRealname := "Foo Bar"
	memberPassword := "1VerySeriousPassword"
	memberComments := "Some comments"

	ctx := context.Background()
	c := NewClient(host, defaultUser, defaultPassword)

	p, err := c.Projects().NewProject(ctx, projectName, 3, 3)
	defer c.Projects().Delete(ctx, p)
	require.NoError(t, err)

	u, err := c.Users().NewUser(ctx, memberUsername, memberEmail, memberRealname, memberPassword, memberComments)
	require.NoError(t, err)
	defer c.Users().Delete(ctx, u)

	members, err := c.Projects().ListMembers(ctx, p)
	require.NoError(t, err)

	assert.Len(t, members, 1)

	err = c.Projects().AddUserMember(ctx, p, u, 1)
	require.NoError(t, err)

	members, err = c.Projects().ListMembers(ctx, p)
	require.NoError(t, err)

	assert.Len(t, members, 2)

}

func TestAPIProjectUserMemberUpdate(t *testing.T) {
	if !*integrationTest {
		t.Skip()
	}

	projectName := "test-project"
	memberUsername := "foobar"
	memberEmail := "foo@bar.com"
	memberRealname := "Foo Bar"
	memberPassword := "1VerySeriousPassword"
	memberComments := "Some comments"

	ctx := context.Background()
	c := NewClient(host, defaultUser, defaultPassword)

	p, err := c.Projects().NewProject(ctx, projectName, 3, 3)
	defer c.Projects().Delete(ctx, p)
	require.NoError(t, err)

	u, err := c.Users().NewUser(ctx, memberUsername, memberEmail, memberRealname, memberPassword, memberComments)
	require.NoError(t, err)
	defer c.Users().Delete(ctx, u)

	err = c.Projects().AddUserMember(ctx, p, u, 1)
	require.NoError(t, err)

	err = c.Projects().UpdateUserMemberRole(ctx, p, u, 2)
	require.NoError(t, err)

	members, err := c.Projects().ListMembers(ctx, p)
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
	if !*integrationTest {
		t.Skip()
	}

	projectName := "test-project"
	memberUsername := "foobar"
	memberEmail := "foo@bar.com"
	memberRealname := "Foo Bar"
	memberPassword := "1VerySeriousPassword"
	memberComments := "Some comments"

	ctx := context.Background()
	c := NewClient(host, defaultUser, defaultPassword)

	p, err := c.Projects().NewProject(ctx, projectName, 3, 3)
	defer c.Projects().Delete(ctx, p)
	require.NoError(t, err)

	u, err := c.Users().NewUser(ctx, memberUsername, memberEmail, memberRealname, memberPassword, memberComments)
	require.NoError(t, err)
	defer c.Users().Delete(ctx, u)

	err = c.Projects().AddUserMember(ctx, p, u, 1)
	require.NoError(t, err)

	err = c.Projects().DeleteUserMember(ctx, p, u)
	require.NoError(t, err)

	members, err := c.Projects().ListMembers(ctx, p)
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
