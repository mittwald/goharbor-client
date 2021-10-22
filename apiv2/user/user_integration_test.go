//go:build integration

package user

import (
	"context"
	"net/url"
	"strconv"
	"testing"

	"github.com/go-openapi/strfmt"

	v2client "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/user"
	modelv2 "github.com/mittwald/goharbor-client/v4/apiv2/model"
	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/config"
	integrationtest "github.com/mittwald/goharbor-client/v4/apiv2/pkg/testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/require"
)

var (
	u, _            = url.Parse(integrationtest.Host)
	v2SwaggerClient = v2client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	authInfo        = runtimeclient.BasicAuth(integrationtest.User, integrationtest.Password)

	username    = "foobar"
	email       = "foo@bar.com"
	realname    = "Foo Bar"
	password    = "1VerySeriousPassword"
	comments    = "Some comments"
	opts        = config.Options{}
	defaultOpts = opts.Defaults()
)

func TestAPIUserNew(t *testing.T) {
	ctx := context.Background()

	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)
	usr, err := c.NewUser(ctx, username, email, realname, password, comments)

	require.NoError(t, err)
	require.NotNil(t, usr)

	defer c.DeleteUser(ctx, usr.UserID)

	require.Equal(t, username, usr.Username)
	require.Equal(t, email, usr.Email)
	require.NotEqual(t, 0, usr.UserID)
}

func TestAPIUserAlreadyExists(t *testing.T) {
	ctx := context.Background()

	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)
	usr, err := c.NewUser(ctx, username, email, realname, password, comments)

	require.NoError(t, err)
	require.NotNil(t, usr)

	defer c.DeleteUser(ctx, usr.UserID)

	_, err = c.NewUser(ctx, username, email, realname, password, comments)

	require.Error(t, err)
	require.IsType(t, &user.CreateUserConflict{}, err)
}

func TestAPIUserGet(t *testing.T) {
	ctx := context.Background()

	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)
	usr, err := c.NewUser(ctx, username, email, realname, password, comments)

	require.NoError(t, err)
	require.NotNil(t, usr)

	defer c.DeleteUser(ctx, usr.UserID)

	usr, err = c.GetUserByName(ctx, username)

	require.NoError(t, err)
	require.NotNil(t, usr)

	require.Equal(t, usr, usr)
}

func TestAPIUserGet_2(t *testing.T) {
	ctx := context.Background()

	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)
	usr, err := c.NewUser(ctx, username, email, realname, password, comments)

	require.NoError(t, err)
	require.NotNil(t, usr)

	defer c.DeleteUser(ctx, usr.UserID)

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

	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)
	for i := 0; i < 50; i++ {
		_, err := c.NewUser(ctx, username+"-"+strconv.Itoa(i), strconv.Itoa(i)+"@bar.com", realname, password, comments)
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

	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)
	usr, err := c.NewUser(ctx, username, email, realname, password, comments)

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

	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)
	usr, err := c.NewUser(ctx, username, email, realname, password, comments)

	require.NoError(t, err)
	require.NotNil(t, usr)

	defer c.DeleteUser(ctx, usr.UserID)

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
