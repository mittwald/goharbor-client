// +build !integration

package user

import (
	"context"
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

func TestRESTClient_NewUser_EmptyUserName(t *testing.T) {
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

	uReq.Username = ""

	p.On("PostUsers",
		postParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PostUsersCreated{},
			nil)

	_, err := cl.NewUser(ctx, "", exampleEmail, "", examplePassword, "")

	assert.Error(t, err)

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

	getParams := &products.GetUsersParams{
		Context:  ctx,
		Username: &exampleUser,
	}

	p.On("GetUsers", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetUsersOK{
			Payload: []*model.User{{Username: exampleUser}},
		}, nil)

	_, err := cl.GetUser(ctx, exampleUser)

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_GetUser_EmptyUserName(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	emptyUserName := ""

	_, err := cl.GetUser(ctx, emptyUserName)

	assert.Error(t, err)

	p.AssertExpectations(t)

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

func TestRESTClient_UpdateUser_EmptyUserName(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	u := &model.User{Username: ""}

	err := cl.UpdateUser(ctx, u)

	assert.Error(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_UpdateUser_UserNotProvided(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	u := &model.User{}
	u = nil

	err := cl.UpdateUser(ctx, u)

	assert.Error(t, err)

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
}

func TestRESTClient_DeleteUser_EmptyUserName(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	u := &model.User{Username: ""}

	err := cl.DeleteUser(ctx, u)

	assert.Error(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_DeleteUser_UserNotProvided(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	u := &model.User{}
	u = nil

	err := cl.DeleteUser(ctx, u)

	assert.Error(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_DeleteUser_UserIDMismatch(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	u := &model.User{
		Username: exampleUser,
		UserID:   1,
	}

	u2 := &model.User{
		Username: exampleUser,
		UserID:   2,
	}

	getParams := &products.GetUsersParams{
		Context:  ctx,
		Username: &u.Username,
	}

	p.On("GetUsers", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetUsersOK{
			Payload: []*model.User{u},
		}, &ErrUserMismatch{})

	err := cl.DeleteUser(ctx, u2)

	assert.Equal(t, err, &ErrUserMismatch{})

	p.AssertExpectations(t)
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

	p.AssertExpectations(t)
}

func TestErrUserAlreadyExists_Error(t *testing.T) {
	var e ErrUserAlreadyExists

	assert.Equal(t, ErrUserAlreadyExistsMsg, e.Error())
}

func TestErrUserBadRequest_Error(t *testing.T) {
	var e ErrUserBadRequest

	assert.Equal(t, ErrUserBadRequestMsg, e.Error())
}

func TestErrUserInvalidID_Error(t *testing.T) {
	var e ErrUserInvalidID

	assert.Equal(t, ErrUserInvalidIDMsg, e.Error())
}

func TestErrUserMismatch_Error(t *testing.T) {
	var e ErrUserMismatch

	assert.Equal(t, ErrUserMismatchMsg, e.Error())
}

func TestErrUserNotFound_Error(t *testing.T) {
	var e ErrUserNotFound

	assert.Equal(t, ErrUserNotFoundMsg, e.Error())
}