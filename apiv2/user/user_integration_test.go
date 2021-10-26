// +build integration

package user

import (
	"context"
	"net/url"
	"strconv"
	"testing"

	"github.com/go-openapi/strfmt"

	v2client "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v5/apiv2/internal/legacyapi/client"
	integrationtest "github.com/mittwald/goharbor-client/v5/apiv2/testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/require"
)

var (
	u, _                = url.Parse(integrationtest.Host)
	legacySwaggerClient = client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	v2SwaggerClient     = v2client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	authInfo            = runtimeclient.BasicAuth(integrationtest.User, integrationtest.Password)

	username = "foobar"
	email    = "foo@bar.com"
	realname = "Foo Bar"
	password = "1VerySeriousPassword"
	comments = "Some comments"
)

func TestAPIUserNew(t *testing.T) {
	ctx := context.Background()

	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)
	user, err := c.NewUser(ctx, username, email, realname, password, comments)

	require.NoError(t, err)
	require.NotNil(t, user)

	defer c.DeleteUser(ctx, user)

	require.Equal(t, username, user.Username)
	require.Equal(t, email, user.Email)
	require.NotEqual(t, 0, user.UserID)
}

func TestAPIUserAlreadyExists(t *testing.T) {
	ctx := context.Background()

	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)
	user, err := c.NewUser(ctx, username, email, realname, password, comments)

	require.NoError(t, err)
	require.NotNil(t, user)

	defer c.DeleteUser(ctx, user)

	_, err = c.NewUser(ctx, username, email, realname, password, comments)

	require.Error(t, err)
	require.Contains(t, err.Error(), "user with this username already exists")
}

func TestAPIUserGet(t *testing.T) {
	ctx := context.Background()

	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)
	user, err := c.NewUser(ctx, username, email, realname, password, comments)

	require.NoError(t, err)
	require.NotNil(t, user)

	defer c.DeleteUser(ctx, user)

	user2, err := c.GetUser(ctx, username)

	require.NoError(t, err)
	require.NotNil(t, user2)

	require.Equal(t, user, user2)
}

func TestAPIUserGet_2(t *testing.T) {
	ctx := context.Background()

	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)
	user, err := c.NewUser(ctx, username, email, realname, password, comments)

	require.NoError(t, err)
	require.NotNil(t, user)

	defer c.DeleteUser(ctx, user)

	user2, err := c.GetUser(ctx, username)

	require.NoError(t, err)
	require.NotNil(t, user2)

	user3, err := c.GetUserByID(ctx, user2.UserID)

	require.NoError(t, err)
	require.NotNil(t, user3)

	require.Equal(t, user, user3)
}

func TestAPIUserList(t *testing.T) {
	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)
	for i := 0; i < 50; i++ {
		_, err := c.NewUser(ctx, username+"-"+strconv.Itoa(i), strconv.Itoa(i)+"@bar.com", realname, password, comments)
		require.NoError(t, err)
	}

	users, err := c.ListUsers(ctx)
	require.NoError(t, err)

	require.Equal(t, 50, len(users))

	for _, u := range users {
		err := c.DeleteUser(ctx, u)
		require.NoError(t, err)
	}
}

func TestAPIUserDelete(t *testing.T) {
	ctx := context.Background()

	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)
	user, err := c.NewUser(ctx, username, email, realname, password, comments)

	require.NoError(t, err)
	require.NotNil(t, user)

	err = c.DeleteUser(ctx, user)
	require.NoError(t, err)

	user, err = c.GetUser(ctx, username)
	require.Error(t, err)
	require.Contains(t, err.Error(), "user not found on server side")
	require.Nil(t, user)
}

func TestAPIUserUpdate(t *testing.T) {
	ctx := context.Background()

	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)
	user, err := c.NewUser(ctx, username, email, realname, password, comments)

	require.NoError(t, err)
	require.NotNil(t, user)

	defer c.DeleteUser(ctx, user)

	user.Email = "foo@baz.com"
	user.Comment = "Some other comment"
	user.Realname = "Foo Baz"
	err = c.UpdateUser(ctx, user)
	require.NoError(t, err)

	user2, err := c.GetUser(ctx, user.Username)
	require.NoError(t, err)

	require.Equal(t, user.Email, user2.Email)
	require.Equal(t, user.Comment, user2.Comment)
	require.Equal(t, user.Realname, user2.Realname)
}
