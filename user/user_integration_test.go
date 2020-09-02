// +build integration

package user

import (
	"context"
	"flag"
	"github.com/go-openapi/strfmt"
	"github.com/mittwald/goharbor-client/internal/api/v1_10_4/client"
	"net/url"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	host     = "http://localhost:30002/api"
	user     = "admin"
	password = "Harbor12345"
)

var (
	u, _          = url.Parse(host)
	swaggerClient = client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	authInfo      = runtimeclient.BasicAuth(user, password)
	harborVersion = flag.String("version", "1.10.2",
		"Harbor version, used in conjunction with -integration, "+
			"defaults to 1.10.2")
	skipSpinUp = flag.Bool("skip-spinup", false,
		"Skip kind cluster creation")
)

func TestAPIUserNew(t *testing.T) {
	ctx := context.Background()
	username := "foobar"
	email := "foo@bar.com"
	realname := "Foo Bar"
	password := "1VerySeriousPassword"
	comments := "Some comments"

	c := NewClient(swaggerClient, authInfo)
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

	c := NewClient(swaggerClient, authInfo)
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

	c := NewClient(swaggerClient, authInfo)
	user, err := c.NewUser(ctx, username, email, realname, password, comments)

	require.NoError(t, err)
	require.NotNil(t, user)

	defer c.DeleteUser(ctx, user)

	user2, err := c.GetUser(ctx, username)

	require.NoError(t, err)
	require.NotNil(t, user2)

	assert.Equal(t, user, user2)
}

func TestAPIUserDelete(t *testing.T) {
	ctx := context.Background()
	username := "foobar"
	email := "foo@bar.com"
	realname := "Foo Bar"
	password := "1VerySeriousPassword"
	comments := "Some comments"

	c := NewClient(swaggerClient, authInfo)
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

	c := NewClient(swaggerClient, authInfo)
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
