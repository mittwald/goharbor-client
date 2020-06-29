// +build !integration

package user

import (
	"context"
	"fmt"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/mittwald/goharbor-client/internal/api/v1.10.0/client"
	"github.com/mittwald/goharbor-client/internal/api/v1.10.0/client/products"
	"github.com/mittwald/goharbor-client/internal/api/v1.10.0/model"
	"github.com/mittwald/goharbor-client/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	authInfo        = runtimeclient.BasicAuth("foo", "bar")
	exampleUser     = "example-user"
	examplePassword = "password"
	exampleEmail    = "test@example.com"
	exampleUserID   = int64(0)
)

func TestRESTClient_NewUser(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	uReq := &model.User{
		Username: exampleUser,
		Password: examplePassword,
		Email:    exampleEmail,
		Realname: "",
		Comment:  "",
	}

	postParams := &products.PostUsersParams{
		User:    uReq,
		Context: ctx,
	}

	getParams := &products.GetUsersParams{
		Context:  ctx,
		Username: &exampleUser}

	p.On("PostUsers",
		postParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PostUsersCreated{},
			nil)

	p.On("GetUsers", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetUsersOK{
			Payload: []*model.User{{Username: exampleUser}},
		}, nil)

	_, err := cl.NewUser(ctx, exampleUser, exampleEmail, "", examplePassword, "")

	assert.NoError(t, err)

	p.AssertExpectations(t)

}

func TestRESTClient_GetUser(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	t.Run("NoUsernameProvided", func(t *testing.T) {
		emptyUserName := ""

		_, err := cl.GetUser(ctx, emptyUserName)

		assert.Error(t, err)

		p.AssertExpectations(t)
	})

	t.Run("GetUser", func(t *testing.T) {
		getParams := &products.GetUsersParams{
			Context:  ctx,
			Username: &exampleUser}

		p.On("GetUsers", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&products.GetUsersOK{
				Payload: []*model.User{{Username: ""}},
			}, nil)

		_, err := cl.GetUser(ctx, exampleUser)

		assert.Error(t, err)

		p.AssertExpectations(t)
	})

	t.Run("UserNotFound", func(t *testing.T) {
		nonexistentUser := "nonexistent-user"
		getParams := &products.GetUsersParams{
			Context:  ctx,
			Username: &exampleUser}

		p.On("GetUsers", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&products.GetUsersOK{
				Payload: []*model.User{{Username: nonexistentUser}},
			}, nil)

		_, err := cl.GetUser(ctx, exampleUser)

		fmt.Println(err)
		assert.Error(t, err)

		p.AssertExpectations(t)

	})
}

func TestRESTClient_UpdateUser(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	getParams := &products.GetUsersParams{
		Context:  ctx,
		Username: &exampleUser}

	putParams := &products.PutUsersUserIDParams{
		UserID: exampleUserID,
		Profile: &model.UserProfile{
			Comment:  "",
			Email:    exampleEmail,
			Realname: "",
		},
		Context: ctx,
	}

	p.On("GetUsers", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetUsersOK{
			Payload: []*model.User{{Username: exampleUser, UserID: exampleUserID, Email: exampleEmail, Password: examplePassword}},
		}, nil)

	u, err := cl.GetUser(ctx, exampleUser)

	p.On("PutUsersUserID", putParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PutUsersUserIDOK{}, nil)

	err = cl.UpdateUser(ctx, u)

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_DeleteUser(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	t.Run("NoUsernameProvided", func(t *testing.T) {
		emptyUserName := ""

		_, err := cl.GetUser(ctx, emptyUserName)

		assert.Error(t, err)

		p.AssertExpectations(t)
	})

	t.Run("DeleteUser", func(t *testing.T) {
		u := &model.User{
			Username: exampleUser,
		}

		deleteParams := &products.DeleteUsersUserIDParams{
			UserID:  u.UserID,
			Context: ctx,
		}

		getParams := &products.GetUsersParams{
			Context:  ctx,
			Username: &u.Username}

		p.On("GetUsers", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&products.GetUsersOK{
				Payload: []*model.User{{Username: exampleUser}},
			}, nil)

		p.On("DeleteUsersUserID", deleteParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&products.DeleteUsersUserIDOK{}, nil)

		err := cl.DeleteUser(ctx, u)

		assert.NoError(t, err)

		p.AssertExpectations(t)
	})
}

func TestRESTClient_UserExists(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	u := &model.User{
		Username: exampleUser,
	}

	getParams := &products.GetUsersParams{
		Context:  ctx,
		Username: &u.Username}

	p.On("GetUsers", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetUsersOK{
			Payload: []*model.User{{Username: exampleUser}},
		}, nil)

	exists, err := cl.UserExists(ctx, u)

	assert.NoError(t, err)

	assert.Equal(t, exists, true)
}
