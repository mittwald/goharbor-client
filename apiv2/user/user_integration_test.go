// +build integration

package user

import (
	"context"
	"net/url"
	"testing"

	"github.com/go-openapi/strfmt"
	v2client "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v4/apiv2/internal/legacyapi/client"
	integrationtest "github.com/mittwald/goharbor-client/v4/apiv2/testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	u, _                = url.Parse(integrationtest.Host)
	legacySwaggerClient = client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	v2SwaggerClient     = v2client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	authInfo            = runtimeclient.BasicAuth(integrationtest.User, integrationtest.Password)
)

func TestAPIUserNew(t *testing.T) {
	ctx := context.Background()
	username := "foobar"
	email := "foo@bar.com"
	realname := "Foo Bar"
	password := "1VerySeriousPassword"
	comments := "Some comments"

	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)
	user, err := c.NewUser(ctx, username, email, realname, password, comments)

	require.NoError(t, err)
	require.NotNil(t, user)

	defer c.DeleteUser(ctx, user)

	assert.Equal(t, username, user.Username)
	assert.Equal(t, email, user.Email)
	assert.NotEqual(t, 0, user.UserID)
}

func TestAPIUserAlreadyExists(t *testing.T) {
	ctx := context.Background()
	username := "foobar"
	email := "foo@bar.com"
	realname := "Foo Bar"
	password := "1VerySeriousPassword"
	comments := "Some comments"

	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)
	user, err := c.NewUser(ctx, username, email, realname, password, comments)

	require.NoError(t, err)
	require.NotNil(t, user)

	defer c.DeleteUser(ctx, user)

	_, err = c.NewUser(ctx, username, email, realname, password, comments)

	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "user with this username already exists")
	}
}

func TestAPIUserGet(t *testing.T) {
	ctx := context.Background()
	username := "foobar"
	email := "foo@bar.com"
	realname := "Foo Bar"
	password := "1VerySeriousPassword"
	comments := "Some comments"

	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)
	user, err := c.NewUser(ctx, username, email, realname, password, comments)

	require.NoError(t, err)
	require.NotNil(t, user)

	defer c.DeleteUser(ctx, user)

	user2, err := c.GetUser(ctx, username)

	require.NoError(t, err)
	require.NotNil(t, user2)

	assert.Equal(t, user, user2)
}

func TestAPIUserGet_2(t *testing.T) {
	ctx := context.Background()
	username := "foobar"
	email := "foo@bar.com"
	realname := "Foo Bar"
	password := "1VerySeriousPassword"
	comments := "Some comments"

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

	assert.Equal(t, user, user3)
}

func TestAPIUserDelete(t *testing.T) {
	ctx := context.Background()
	username := "foobar"
	email := "foo@bar.com"
	realname := "Foo Bar"
	password := "1VerySeriousPassword"
	comments := "Some comments"

	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)
	user, err := c.NewUser(ctx, username, email, realname, password, comments)

	require.NoError(t, err)
	require.NotNil(t, user)

	err = c.DeleteUser(ctx, user)
	require.NoError(t, err)

	user, err = c.GetUser(ctx, username)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "user not found on server side")
	require.Nil(t, user)
}

func TestAPIUserUpdate(t *testing.T) {
	ctx := context.Background()
	username := "foobar"
	email := "foo@bar.com"
	realname := "Foo Bar"
	password := "1VerySeriousPassword"
	comments := "Some comments"

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

	assert.Equal(t, user.Email, user2.Email)
	assert.Equal(t, user.Comment, user2.Comment)
	assert.Equal(t, user.Realname, user2.Realname)
}
