//go:build !integration

package user

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/util/intstr"

	runtimeclient "github.com/go-openapi/runtime/client"

	"github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/user"
	"github.com/mittwald/goharbor-client/v4/apiv2/mocks"
	modelv2 "github.com/mittwald/goharbor-client/v4/apiv2/model"
	unittesting "github.com/mittwald/goharbor-client/v4/apiv2/pkg/testing"
)

var (
	authInfo        = runtimeclient.BasicAuth("foo", "bar")
	exampleUsername = "example-user"
	examplePassword = "password"
	exampleEmail    = "test@example.com"
	exampleUserID   = int64(1)
	ctx             = context.Background()
)

func APIandMockClientsForTests() (*RESTClient, *unittesting.MockClients) {
	desiredMockClients := &unittesting.MockClients{
		User: mocks.MockUserClientService{},
	}

	v2Client := unittesting.BuildV2ClientWithMocks(desiredMockClients)

	cl := NewClient(v2Client, &unittesting.DefaultOpts, authInfo)

	return cl, desiredMockClients
}

func TestTRESTClient_NewUser(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	createParams := &user.CreateUserParams{
		UserReq: &modelv2.UserCreationReq{
			Comment:  "",
			Email:    exampleEmail,
			Password: examplePassword,
			Realname: "",
			Username: exampleUsername,
		},
		Context: ctx,
	}

	listParams := &user.ListUsersParams{
		PageSize: &unittesting.DefaultOpts.PageSize,
		Q:        &unittesting.DefaultOpts.Query,
		Sort:     &unittesting.DefaultOpts.Sort,
		Context:  ctx,
	}

	mockClient.User.On("CreateUser", createParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&user.CreateUserCreated{}, nil)

	mockClient.User.On("ListUsers", listParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&user.ListUsersOK{Payload: []*modelv2.UserResp{{
			Username: exampleUsername,
		}}}, nil)

	_, err := apiClient.NewUser(ctx, exampleUsername, exampleEmail, "", examplePassword, "")

	require.NoError(t, err)

	mockClient.User.AssertExpectations(t)
}

func TestRESTClient_GetUserByID(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := &user.GetUserParams{
		UserID:  exampleUserID,
		Context: ctx,
	}

	mockClient.User.On("GetUser", getParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&user.GetUserOK{}, nil)

	_, err := apiClient.GetUserByID(ctx, exampleUserID)

	require.Error(t, err)
	require.ErrorIs(t, err, &ErrUserNotFound{})

	mockClient.User.AssertExpectations(t)
}

func TestRESTClient_GetUserByID_2(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := &user.GetUserParams{
		UserID:  exampleUserID,
		Context: ctx,
	}

	mockClient.User.On("GetUser", getParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&user.GetUserOK{Payload: &modelv2.UserResp{
			UserID: exampleUserID,
		}}, nil)

	_, err := apiClient.GetUserByID(ctx, exampleUserID)

	require.NoError(t, err)

	mockClient.User.AssertExpectations(t)
}

func TestRESTClient_GetUserByID_3(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := &user.GetUserParams{
		UserID:  exampleUserID,
		Context: ctx,
	}

	for i := -1; i <= 0; i++ {
		t.Run(fmt.Sprintf("ErrUserInvalidID/%d", i), func(t *testing.T) {
			mockClient.User.On("GetUser", getParams,
				mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
				Return(&user.GetUserOK{Payload: &modelv2.UserResp{
					UserID: int64(i),
				}}, nil)

			_, err := apiClient.GetUserByID(ctx, exampleUserID)

			require.Error(t, err)
			require.ErrorIs(t, err, &ErrUserInvalidID{})
		})
	}

	mockClient.User.AssertExpectations(t)
}

func TestRESTClient_GetUserByName(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	listParams := &user.ListUsersParams{
		PageSize: &apiClient.Options.PageSize,
		Q:        &apiClient.Options.Query,
		Sort:     &apiClient.Options.Sort,
		Context:  ctx,
	}

	mockClient.User.On("ListUsers", listParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&user.ListUsersOK{
			Payload: []*modelv2.UserResp{{
				Username: exampleUsername,
			}},
		}, nil)

	_, err := apiClient.GetUserByName(ctx, exampleUsername)
	require.NoError(t, err)

	mockClient.User.AssertExpectations(t)
}

func TestRESTClient_GetUserByName_2(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	_, err := apiClient.GetUserByName(ctx, "")
	require.Error(t, err)

	mockClient.User.AssertExpectations(t)
}

func TestRESTClient_UpdateUserProfile(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	profile := &modelv2.UserProfile{
		Comment:  "",
		Email:    exampleEmail,
		Realname: "",
	}

	getParams := &user.GetUserParams{
		UserID:  exampleUserID,
		Context: ctx,
	}

	updateParams := &user.UpdateUserProfileParams{
		Profile: profile,
		UserID:  exampleUserID,
		Context: ctx,
	}

	t.Run("", func(t *testing.T) {
		mockClient.User.On("GetUser", getParams,
			mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&user.GetUserOK{Payload: &modelv2.UserResp{UserID: exampleUserID}}, nil)

		mockClient.User.On("UpdateUserProfile", updateParams,
			mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&user.UpdateUserProfileOK{}, nil)

		err := apiClient.UpdateUserProfile(ctx, exampleUserID, profile)

		require.NoError(t, err)
	})

	mockClient.User.AssertExpectations(t)
}

func TestRESTClient_UpdateUserPassword(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := &user.GetUserParams{
		UserID:  exampleUserID,
		Context: ctx,
	}

	updatePWParams := &user.UpdateUserPasswordParams{
		Password: &modelv2.PasswordReq{
			NewPassword: "bar",
			OldPassword: "foo",
		},
		UserID:  exampleUserID,
		Context: ctx,
	}

	t.Run("WithPassword", func(t *testing.T) {
		mockClient.User.On("GetUser", getParams,
			mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&user.GetUserOK{Payload: &modelv2.UserResp{UserID: exampleUserID}}, nil)

		mockClient.User.On("UpdateUserPassword", updatePWParams,
			mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&user.UpdateUserPasswordOK{}, nil)

		err := apiClient.UpdateUserPassword(ctx, exampleUserID, "foo", "bar")
		require.NoError(t, err)
	})

	t.Run("NoOldPassword", func(t *testing.T) {
		mockClient.User.On("GetUser", getParams,
			mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&user.GetUserOK{Payload: &modelv2.UserResp{UserID: exampleUserID}}, nil)

		mockClient.User.On("UpdateUserPassword", updatePWParams,
			mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&user.UpdateUserPasswordOK{}, nil)

		err := apiClient.UpdateUserPassword(ctx, exampleUserID, "", "bar")
		require.Error(t, err)
		require.Contains(t, err.Error(), "no old password provided")
	})

	t.Run("NoNewPassword", func(t *testing.T) {
		mockClient.User.On("GetUser", getParams,
			mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&user.GetUserOK{Payload: &modelv2.UserResp{UserID: exampleUserID}}, nil)

		mockClient.User.On("UpdateUserPassword", updatePWParams,
			mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&user.UpdateUserPasswordOK{}, nil)

		err := apiClient.UpdateUserPassword(ctx, exampleUserID, "foo", "")
		require.Error(t, err)
		require.Contains(t, err.Error(), "no new password provided")
	})

	mockClient.User.AssertExpectations(t)
}

func TestRESTClient_DeleteUser(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := &user.GetUserParams{
		UserID:  exampleUserID,
		Context: ctx,
	}

	deleteParams := &user.DeleteUserParams{
		UserID:  exampleUserID,
		Context: ctx,
	}

	mockClient.User.On("GetUser", getParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&user.GetUserOK{Payload: &modelv2.UserResp{UserID: exampleUserID}}, nil)

	mockClient.User.On("DeleteUser", deleteParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&user.DeleteUserOK{}, nil)

	err := apiClient.DeleteUser(ctx, exampleUserID)
	require.NoError(t, err)

	mockClient.User.AssertExpectations(t)
}

func TestRESTClient_UserExists(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := &user.GetUserParams{
		UserID:  exampleUserID,
		Context: ctx,
	}

	listParams := &user.ListUsersParams{
		PageSize: &apiClient.Options.PageSize,
		Q:        &apiClient.Options.Query,
		Sort:     &apiClient.Options.Sort,
		Context:  ctx,
	}

	t.Run("ByID", func(t *testing.T) {
		mockClient.User.On("GetUser", getParams,
			mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&user.GetUserOK{Payload: &modelv2.UserResp{UserID: exampleUserID}}, nil)

		exists, err := apiClient.UserExists(ctx, intstr.FromInt(int(exampleUserID)))
		require.NoError(t, err)
		require.True(t, exists)
	})

	t.Run("ByName", func(t *testing.T) {
		mockClient.User.On("ListUsers", listParams,
			mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&user.ListUsersOK{Payload: []*modelv2.UserResp{{Username: exampleUsername}}}, nil)

		exists, err := apiClient.UserExists(ctx, intstr.FromString(exampleUsername))
		require.NoError(t, err)
		require.True(t, exists)
	})

	mockClient.User.AssertExpectations(t)
}

//
// func TestRESTClient_UserExists_ErrUserNotFound(t *testing.T) {
// 	p := &mocks.MockProductsClientService{}
//
// 	legacyClient := BuildLegacyClientWithMock(p)
// 	v2Client := BuildV2ClientWithMocks()
//
// 	cl := NewClient(legacyClient, v2Client, authInfo)
//
// 	ctx := context.Background()
//
// 	u := &legacymodel.User{
// 		Username: exampleUser,
// 	}
//
// 	getUserParams := &products.GetUsersParams{
// 		Context:  ctx,
// 		Username: &u.Username,
// 	}
//
// 	p.On("GetUsers", getUserParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
// 		Return(&products.GetUsersOK{
// 			Payload: nil,
// 		}, &ErrUserNotFound{})
//
// 	exists, err := cl.UserExists(ctx, u)
//
// 	assert.Equal(t, false, exists)
// 	assert.NoError(t, err)
//
// 	u2 := &legacymodel.User{
// 		Username: "",
// 	}
//
// 	getUserParams2 := &products.GetUsersParams{
// 		Context:  ctx,
// 		Username: &u.Username,
// 	}
//
// 	p.On("GetUsers", getUserParams2, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
// 		Return(&products.GetUsersOK{
// 			Payload: nil,
// 		}, &ErrUserNotFound{})
//
// 	exists, err = cl.UserExists(ctx, u2)
//
// 	assert.Equal(t, false, exists)
// 	if assert.Error(t, err) {
// 		assert.Equal(t, errors.New("no username provided"), err)
// 	}
//
// 	p.AssertExpectations(t)
// }
//
// func TestErrUserAlreadyExists_Error(t *testing.T) {
// 	var e ErrUserAlreadyExists
//
// 	assert.Equal(t, ErrUserAlreadyExistsMsg, e.Error())
// }
//
// func TestErrUserBadRequest_Error(t *testing.T) {
// 	var e ErrUserBadRequest
//
// 	assert.Equal(t, ErrUserBadRequestMsg, e.Error())
// }
//
// func TestErrUserInvalidID_Error(t *testing.T) {
// 	var e ErrUserInvalidID
//
// 	assert.Equal(t, ErrUserInvalidIDMsg, e.Error())
// }
//
// func TestErrUserMismatch_Error(t *testing.T) {
// 	var e ErrUserMismatch
//
// 	assert.Equal(t, ErrUserMismatchMsg, e.Error())
// }
//
// func TestErrUserNotFound_Error(t *testing.T) {
// 	var e ErrUserNotFound
//
// 	assert.Equal(t, ErrUserNotFoundMsg, e.Error())
// }
//
// func TestErrUserPasswordInvalid_Error(t *testing.T) {
// 	var e ErrUserPasswordInvalid
//
// 	assert.Equal(t, ErrUserPasswordInvalidMsg, e.Error())
// }
//
// func TestErrUserIDNotExists_Error(t *testing.T) {
// 	var e ErrUserIDNotExists
//
// 	assert.Equal(t, ErrUserIDNotExistsMsg, e.Error())
// }
