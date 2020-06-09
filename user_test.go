package goharborclient

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAPIUserNew(t *testing.T) {
	if !*integrationTest {
		t.Skip()
	}

	ctx := context.Background()
	username := "foobar"
	email := "foo@bar.com"
	realname := "Foo Bar"
	password := "1VerySeriousPassword"
	comments := "Some comments"

	c := NewClient(host, defaultUser, defaultPassword)
	user, err := c.Users().NewUser(ctx, username, email, realname, password, comments)

	require.NoError(t, err)
	require.NotNil(t, user)

	defer c.Users().Delete(ctx, user)

	assert.Equal(t, username, user.Username)
	assert.Equal(t, email, user.Email)
	assert.NotEqual(t, 0, user.UserID)
}

func TestAPIUserAlreadyExists(t *testing.T) {
	if !*integrationTest {
		t.Skip()
	}

	ctx := context.Background()
	username := "foobar"
	email := "foo@bar.com"
	realname := "Foo Bar"
	password := "1VerySeriousPassword"
	comments := "Some comments"

	c := NewClient(host, defaultUser, defaultPassword)
	user, err := c.Users().NewUser(ctx, username, email, realname, password, comments)

	require.NoError(t, err)
	require.NotNil(t, user)

	defer c.Users().Delete(ctx, user)

	_, err = c.Users().NewUser(ctx, username, email, realname, password, comments)

	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "User with this username already exists")
	}
}

func TestAPIUserGet(t *testing.T) {
	if !*integrationTest {
		t.Skip()
	}

	ctx := context.Background()
	username := "foobar"
	email := "foo@bar.com"
	realname := "Foo Bar"
	password := "1VerySeriousPassword"
	comments := "Some comments"

	c := NewClient(host, defaultUser, defaultPassword)
	user, err := c.Users().NewUser(ctx, username, email, realname, password, comments)

	require.NoError(t, err)
	require.NotNil(t, user)

	defer c.Users().Delete(ctx, user)

	user2, err := c.Users().Get(ctx, username)

	require.NoError(t, err)
	require.NotNil(t, user2)

	assert.Equal(t, user, user2)
}

func TestAPIUserDelete(t *testing.T) {
	if !*integrationTest {
		t.Skip()
	}

	ctx := context.Background()
	username := "foobar"
	email := "foo@bar.com"
	realname := "Foo Bar"
	password := "1VerySeriousPassword"
	comments := "Some comments"

	c := NewClient(host, defaultUser, defaultPassword)
	user, err := c.Users().NewUser(ctx, username, email, realname, password, comments)

	require.NoError(t, err)
	require.NotNil(t, user)

	err = c.Users().Delete(ctx, user)
	require.NoError(t, err)

	user, err = c.Users().Get(ctx, username)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "user not found on server side")
	require.Nil(t, user)
}

func TestAPIUserUpdate(t *testing.T) {
	if !*integrationTest {
		t.Skip()
	}

	ctx := context.Background()
	username := "foobar"
	email := "foo@bar.com"
	realname := "Foo Bar"
	password := "1VerySeriousPassword"
	comments := "Some comments"

	c := NewClient(host, defaultUser, defaultPassword)
	user, err := c.Users().NewUser(ctx, username, email, realname, password, comments)

	require.NoError(t, err)
	require.NotNil(t, user)

	defer c.Users().Delete(ctx, user)

	user.Email = "foo@baz.com"
	user.Comment = "Some other comment"
	user.Realname = "Foo Baz"
	err = c.Users().Update(ctx, user)
	require.NoError(t, err)

	user2, err := c.Users().Get(ctx, user.Username)
	require.NoError(t, err)

	assert.Equal(t, user.Email, user2.Email)
	assert.Equal(t, user.Comment, user2.Comment)
	assert.Equal(t, user.Realname, user2.Realname)
}
