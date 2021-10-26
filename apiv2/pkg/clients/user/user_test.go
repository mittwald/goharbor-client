//go:build !integration

package user

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/user"
	"github.com/mittwald/goharbor-client/v5/apiv2/mocks"
	modelv2 "github.com/mittwald/goharbor-client/v5/apiv2/model"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/errors"
	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"
)

var (
	exampleUsername = "example-user"
	examplePassword = "password"
	exampleEmail    = "test@example.com"
	exampleUserID   = int64(1)
	ctx             = context.Background()
)

func APIandMockClientsForTests() (*RESTClient, *clienttesting.MockClients) {
	desiredMockClients := &clienttesting.MockClients{
		User: mocks.MockUserClientService{},
	}

	v2Client := clienttesting.BuildV2ClientWithMocks(desiredMockClients)

	cl := NewClient(v2Client, clienttesting.DefaultOpts, clienttesting.AuthInfo)

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

	createParams.WithTimeout(apiClient.Options.Timeout)

	listParams := &user.ListUsersParams{
		PageSize: &clienttesting.DefaultOpts.PageSize,
		Q:        &clienttesting.DefaultOpts.Query,
		Sort:     &clienttesting.DefaultOpts.Sort,
		Context:  ctx,
	}

	listParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.User.On("CreateUser", createParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&user.CreateUserCreated{}, nil)

	err := apiClient.NewUser(ctx, exampleUsername, exampleEmail, "", examplePassword, "")

	require.NoError(t, err)

	mockClient.User.AssertExpectations(t)
}

func TestRESTClient_GetUserByID(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := &user.GetUserParams{
		UserID:  exampleUserID,
		Context: ctx,
	}

	getParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.User.On("GetUser", getParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&user.GetUserOK{}, nil)

	_, err := apiClient.GetUserByID(ctx, exampleUserID)

	require.Error(t, err)
	require.ErrorIs(t, err, &errors.ErrUserNotFound{})

	mockClient.User.AssertExpectations(t)
}

func TestRESTClient_GetUserByID_2(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := &user.GetUserParams{
		UserID:  exampleUserID,
		Context: ctx,
	}

	getParams.WithTimeout(apiClient.Options.Timeout)

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

	getParams.WithTimeout(apiClient.Options.Timeout)

	for i := -1; i <= 0; i++ {
		t.Run(fmt.Sprintf("ErrUserInvalidID/%d", i), func(t *testing.T) {
			mockClient.User.On("GetUser", getParams,
				mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
				Return(&user.GetUserOK{Payload: &modelv2.UserResp{
					UserID: int64(i),
				}}, nil)

			_, err := apiClient.GetUserByID(ctx, exampleUserID)

			require.Error(t, err)
			require.ErrorIs(t, err, &errors.ErrUserInvalidID{})
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

	listParams.WithTimeout(apiClient.Options.Timeout)

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

	getParams.WithTimeout(apiClient.Options.Timeout)

	updateParams := &user.UpdateUserProfileParams{
		Profile: profile,
		UserID:  exampleUserID,
		Context: ctx,
	}

	updateParams.WithTimeout(apiClient.Options.Timeout)

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

	getParams.WithTimeout(apiClient.Options.Timeout)

	updatePWParams := &user.UpdateUserPasswordParams{
		Password: &modelv2.PasswordReq{
			NewPassword: "bar",
			OldPassword: "foo",
		},
		UserID:  exampleUserID,
		Context: ctx,
	}

	updatePWParams.WithTimeout(apiClient.Options.Timeout)

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

	getParams.WithTimeout(apiClient.Options.Timeout)

	deleteParams := &user.DeleteUserParams{
		UserID:  exampleUserID,
		Context: ctx,
	}

	deleteParams.WithTimeout(apiClient.Options.Timeout)

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

	getParams.WithTimeout(apiClient.Options.Timeout)

	listParams := &user.ListUsersParams{
		PageSize: &apiClient.Options.PageSize,
		Q:        &apiClient.Options.Query,
		Sort:     &apiClient.Options.Sort,
		Context:  ctx,
	}

	listParams.WithTimeout(apiClient.Options.Timeout)

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
