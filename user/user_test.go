// +build !integration

package user

import (
	"context"
	"errors"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/mittwald/goharbor-client/internal/api/v1_10_0/client"
	"github.com/mittwald/goharbor-client/internal/api/v1_10_0/client/products"
	"github.com/mittwald/goharbor-client/mocks"
	model "github.com/mittwald/goharbor-client/model/v1_10_0"
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

	postUserParams := &products.PostUsersParams{
		User:    uReq,
		Context: ctx,
	}

	getUserParams := &products.GetUsersParams{
		Context:  ctx,
		Username: &exampleUser,
	}

	p.On("PostUsers",
		postUserParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PostUsersCreated{},
			nil)

	p.On("GetUsers", getUserParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
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

	postUserParams := &products.PostUsersParams{
		User:    uReq,
		Context: ctx,
	}

	uReq.Username = ""

	p.On("PostUsers",
		postUserParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PostUsersCreated{},
			nil)

	_, err := cl.NewUser(ctx, "", exampleEmail, "", examplePassword, "")

	if assert.Error(t, err) {
		assert.Equal(t, errors.New("no username provided"), err)
	}

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

	getUserParams := &products.GetUsersParams{
		Context:  ctx,
		Username: &exampleUser,
	}

	p.On("GetUsers", getUserParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
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

	if assert.Error(t, err) {
		assert.Equal(t, errors.New("no username provided"), err)
	}

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

	getUserParams := &products.GetUsersParams{
		Context:  ctx,
		Username: &exampleUser,
	}

	putUserParams := &products.PutUsersUserIDParams{
		UserID: exampleUserID,
		Profile: &model.UserProfile{
			Comment:  "",
			Email:    exampleEmail,
			Realname: "",
		},
		Context: ctx,
	}

	p.On("GetUsers", getUserParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetUsersOK{
			Payload: []*model.User{{
				Username: exampleUser, UserID: exampleUserID, Email: exampleEmail,
				Password: examplePassword,
			}},
		}, nil)

	u, err := cl.GetUser(ctx, exampleUser)
	assert.NoError(t, err)

	p.On("PutUsersUserID", putUserParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PutUsersUserIDOK{}, nil)

	err = cl.UpdateUser(ctx, u)

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_UpdateUserPassword(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	password := &model.Password{
		NewPassword: "foo",
		OldPassword: "bar",
	}
	putUserPasswordParams := &products.PutUsersUserIDPasswordParams{
		Password: password,
		UserID:   0,
		Context:  ctx,
	}

	p.On("PutUsersUserIDPassword", putUserPasswordParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PutUsersUserIDPasswordOK{}, nil)

	err := cl.UpdateUserPassword(ctx, 0, password)

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_UpdateUserPassword_NoPasswordProvided(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	password := &model.Password{}

	password = nil

	err := cl.UpdateUserPassword(ctx, 0, password)

	assert.Errorf(t, err, "no password provided")

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

	if assert.Error(t, err) {
		assert.Equal(t, errors.New("no username provided"), err)
	}

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

	deleteUserParams := &products.DeleteUsersUserIDParams{
		UserID:  u.UserID,
		Context: ctx,
	}

	getUserParams := &products.GetUsersParams{
		Context:  ctx,
		Username: &u.Username,
	}

	p.On("GetUsers", getUserParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetUsersOK{
			Payload: []*model.User{{Username: exampleUser}},
		}, nil)

	p.On("DeleteUsersUserID", deleteUserParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
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

	if assert.Error(t, err) {
		assert.Equal(t, errors.New("no username provided"), err)
	}

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

	err := cl.DeleteUser(ctx, nil)

	if assert.Error(t, err) {
		assert.Equal(t, errors.New("no user provided"), err)
	}

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

	getUserParams := &products.GetUsersParams{
		Context:  ctx,
		Username: &u.Username,
	}

	p.On("GetUsers", getUserParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetUsersOK{
			Payload: []*model.User{u},
		}, &ErrUserMismatch{})

	err := cl.DeleteUser(ctx, u2)

	assert.IsType(t, err, &ErrUserMismatch{})

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

	getUserParams := &products.GetUsersParams{
		Context:  ctx,
		Username: &u.Username,
	}

	p.On("GetUsers", getUserParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
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

func TestErrUserPasswordInvalid_Error(t *testing.T) {
	var e ErrUserPasswordInvalid

	assert.Equal(t, ErrUserPasswordInvalidMsg, e.Error())
}
