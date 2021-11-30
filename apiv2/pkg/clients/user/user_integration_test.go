//go:build integration

package user

import (
	"context"
	"strconv"
	"testing"

	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/user"
	modelv2 "github.com/mittwald/goharbor-client/v5/apiv2/model"
	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"

	"github.com/stretchr/testify/require"
)

var (
	username = "foobar"
	email    = "foo@bar.com"
	realname = "Foo Bar"
	password = "1VerySeriousPassword"
	comments = "Some comments"
)

func TestAPIUserNew(t *testing.T) {
	ctx := context.Background()

	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)
	err := c.NewUser(ctx, username, email, realname, password, comments)
	require.NoError(t, err)

	usr, err := c.GetUserByName(ctx, username)
	require.NoError(t, err)

	require.NotNil(t, usr)

	defer func() {
		_ = c.DeleteUser(ctx, usr.UserID)
	}()

	require.Equal(t, username, usr.Username)
	require.Equal(t, email, usr.Email)
	require.NotEqual(t, 0, usr.UserID)
}

func TestAPIUserAlreadyExists(t *testing.T) {
	ctx := context.Background()

	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)
	err := c.NewUser(ctx, username, email, realname, password, comments)
	require.NoError(t, err)

	usr, err := c.GetUserByName(ctx, username)
	require.NoError(t, err)

	defer func() {
		_ = c.DeleteUser(ctx, usr.UserID)
	}()

	err = c.NewUser(ctx, username, email, realname, password, comments)

	require.Error(t, err)
	require.IsType(t, &user.CreateUserConflict{}, err)
}

func TestAPIUserGet(t *testing.T) {
	ctx := context.Background()

	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)
	err := c.NewUser(ctx, username, email, realname, password, comments)
	require.NoError(t, err)

	usr, err := c.GetUserByName(ctx, username)
	require.NoError(t, err)

	require.NotNil(t, usr)

	defer func() {
		_ = c.DeleteUser(ctx, usr.UserID)
	}()

	usr, err = c.GetUserByName(ctx, username)

	require.NoError(t, err)
	require.NotNil(t, usr)

	require.Equal(t, usr, usr)
}

func TestAPIUserGet_2(t *testing.T) {
	ctx := context.Background()

	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)
	err := c.NewUser(ctx, username, email, realname, password, comments)
	require.NoError(t, err)

	usr, err := c.GetUserByName(ctx, username)
	require.NoError(t, err)

	require.NotNil(t, usr)

	defer func() {
		_ = c.DeleteUser(ctx, usr.UserID)
	}()

	user2, err := c.GetUserByName(ctx, username)

	require.NoError(t, err)
	require.NotNil(t, user2)

	user3, err := c.GetUserByID(ctx, usr.UserID)

	require.NoError(t, err)
	require.NotNil(t, user3)

	require.Equal(t, usr, user3)
}

func TestAPIUserList(t *testing.T) {
	ctx := context.Background()

	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)
	for i := 0; i < 50; i++ {
		err := c.NewUser(ctx, username+"-"+strconv.Itoa(i), strconv.Itoa(i)+"@bar.com", realname, password, comments)

		require.NoError(t, err)
	}

	users, err := c.ListUsers(ctx)
	require.NoError(t, err)

	require.Equal(t, 50, len(users))

	for _, u := range users {
		err := c.DeleteUser(ctx, u.UserID)
		require.NoError(t, err)
	}
}

func TestAPIUserDelete(t *testing.T) {
	ctx := context.Background()

	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)
	err := c.NewUser(ctx, username, email, realname, password, comments)
	require.NoError(t, err)

	usr, err := c.GetUserByName(ctx, username)
	require.NoError(t, err)

	require.NotNil(t, usr)

	err = c.DeleteUser(ctx, usr.UserID)
	require.NoError(t, err)

	usr, err = c.GetUserByName(ctx, username)
	require.Error(t, err)
	require.Contains(t, err.Error(), "user not found on server side")
	require.Nil(t, usr)
}

func TestAPIUserUpdate(t *testing.T) {
	ctx := context.Background()

	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)
	err := c.NewUser(ctx, username, email, realname, password, comments)
	require.NoError(t, err)

	usr, err := c.GetUserByName(ctx, username)
	require.NoError(t, err)

	require.NotNil(t, usr)

	defer func() {
		_ = c.DeleteUser(ctx, usr.UserID)
	}()

	err = c.UpdateUserProfile(ctx, usr.UserID, &modelv2.UserProfile{
		Email:    "foo@baz.com",
		Comment:  "Some other comment",
		Realname: "Foo Baz",
	})
	require.NoError(t, err)

	usr, err = c.GetUserByName(ctx, usr.Username)
	require.NoError(t, err)

	require.Equal(t, usr.Email, "foo@baz.com")
	require.Equal(t, usr.Comment, "Some other comment")
	require.Equal(t, usr.Realname, "Foo Baz")
}
