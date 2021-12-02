package user

import (
	"context"
	"errors"

	"k8s.io/apimachinery/pkg/util/intstr"

	v2client "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/user"
	modelv2 "github.com/mittwald/goharbor-client/v5/apiv2/model"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/config"
	clienterrors "github.com/mittwald/goharbor-client/v5/apiv2/pkg/errors"

	"github.com/go-openapi/runtime"
)

// RESTClient is a subclient for handling user related actions.
type RESTClient struct {
	// The new client of the harbor v2 API
	V2Client *v2client.Harbor

	// Options contains optional configuration when making API calls.
	Options *config.Options

	// AuthInfo contains the auth information that is provided on API calls.
	AuthInfo runtime.ClientAuthInfoWriter
}

func NewClient(v2Client *v2client.Harbor, opts *config.Options, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		Options:  opts,
		V2Client: v2Client,
		AuthInfo: authInfo,
	}
}

type Client interface {
	NewUser(ctx context.Context, username, email, realname, password, comments string) error
	GetUserByName(ctx context.Context, username string) (*modelv2.UserResp, error)
	GetUserByID(ctx context.Context, id int64) (*modelv2.UserResp, error)
	ListUsers(ctx context.Context) ([]*modelv2.UserResp, error)
	SearchUsers(ctx context.Context, name string) ([]*modelv2.UserSearchRespItem, error)
	GetCurrentUserInfo(ctx context.Context) (*modelv2.UserResp, error)
	GetCurrentUserPermisisons(ctx context.Context, relative bool, scope string) ([]*modelv2.Permission, error)
	SetUserSysAdmin(ctx context.Context, id int64, admin bool) error
	DeleteUser(ctx context.Context, id int64) error
	UpdateUserProfile(ctx context.Context, id int64, profile *modelv2.UserProfile) error
	UpdateUserPassword(ctx context.Context, userID int64, passwordRequest *modelv2.PasswordReq) error
	UserExists(ctx context.Context, idOrName intstr.IntOrString) (bool, error)
}

// NewUser creates and returns a new user, or error in case of failure.
// Username is a unique username.
// email is the Email of the user.
// realname is the fullname of the user.
// password is the password for this user.
// comments as a comment attached to the user.
func (c *RESTClient) NewUser(ctx context.Context, username, email, realname, password, comments string) error {
	params := &user.CreateUserParams{
		UserReq: &modelv2.UserCreationReq{
			Username: username,
			Password: password,
			Email:    email,
			Realname: realname,
			Comment:  comments,
		},
		Context: ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.User.CreateUser(params, c.AuthInfo)
	if err != nil {
		return handleSwaggerUserErrors(err)
	}

	return nil
}

// GetUserByName returns an existing user identified by name.
func (c *RESTClient) GetUserByName(ctx context.Context, username string) (*modelv2.UserResp, error) {
	if username == "" {
		return nil, errors.New("no username provided")
	}

	c.Options.PageSize = 100

	resp, err := c.ListUsers(ctx)
	if err != nil {
		return nil, handleSwaggerUserErrors(err)
	}

	for _, u := range resp {
		if u.Username == username {
			return u, nil
		}
	}

	return nil, &clienterrors.ErrUserNotFound{}
}

// GetUserByID returns an existing user identified by ID.
func (c *RESTClient) GetUserByID(ctx context.Context, id int64) (*modelv2.UserResp, error) {
	if id <= 0 {
		return nil, &clienterrors.ErrUserInvalidID{}
	}

	c.Options.PageSize = 100

	params := &user.GetUserParams{
		UserID:  id,
		Context: ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.User.GetUser(params, c.AuthInfo)

	if resp.Payload == nil {
		return nil, &clienterrors.ErrUserNotFound{}
	}

	if resp.Payload.UserID != id {
		return nil, &clienterrors.ErrUserInvalidID{}
	}

	if err != nil {
		return nil, handleSwaggerUserErrors(err)
	}

	return resp.Payload, nil
}

// ListUsers lists and returns all registered Harbor users.
// The maximum number of users listed is bound to the RESTClient's configured PageSize.
func (c *RESTClient) ListUsers(ctx context.Context) ([]*modelv2.UserResp, error) {
	params := user.ListUsersParams{
		PageSize: &c.Options.PageSize,
		Q:        &c.Options.Query,
		Sort:     &c.Options.Sort,
		Context:  ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.User.ListUsers(&params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerUserErrors(err)
	}

	return resp.Payload, nil
}

// SearchUsers searches all existing users by the provided username 'name' and returns matching users.
func (c *RESTClient) SearchUsers(ctx context.Context, name string) ([]*modelv2.UserSearchRespItem, error) {
	params := &user.SearchUsersParams{
		PageSize: &c.Options.PageSize,
		Username: name,
		Context:  ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.User.SearchUsers(params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerUserErrors(err)
	}

	if len(resp.Payload) == 0 {
		return nil, &clienterrors.ErrUserNotFound{}
	}

	return resp.Payload, nil
}

// GetCurrentUserInfo returns information of currently active user.
func (c *RESTClient) GetCurrentUserInfo(ctx context.Context) (*modelv2.UserResp, error) {
	params := &user.GetCurrentUserInfoParams{
		Context: ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.User.GetCurrentUserInfo(params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerUserErrors(err)
	}

	return resp.Payload, nil
}

// GetCurrentUserPermisisons returns the permissions of the currently active user.
func (c *RESTClient) GetCurrentUserPermisisons(ctx context.Context, relative bool, scope string) ([]*modelv2.Permission, error) {
	params := &user.GetCurrentUserPermissionsParams{
		Relative: &relative,
		Scope:    &scope,
		Context:  ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.User.GetCurrentUserPermissions(params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerUserErrors(err)
	}

	return resp.Payload, nil
}

// SetUserSysAdmin updates a user's administrator privileges.
func (c *RESTClient) SetUserSysAdmin(ctx context.Context, id int64, admin bool) error {
	params := &user.SetUserSysAdminParams{
		SysadminFlag: &modelv2.UserSysAdminFlag{
			SysadminFlag: admin,
		},
		UserID:  id,
		Context: ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.User.SetUserSysAdmin(params, c.AuthInfo)

	return handleSwaggerUserErrors(err)
}

// DeleteUser deletes the specified user, first ensuring its existence.
func (c *RESTClient) DeleteUser(ctx context.Context, id int64) error {
	params := &user.DeleteUserParams{
		UserID:  id,
		Context: ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.GetUserByID(ctx, id)
	if err != nil {
		return err
	}

	_, err = c.V2Client.User.DeleteUser(params, c.AuthInfo)

	return handleSwaggerUserErrors(err)
}

// UpdateUserProfile updates a user identified by id with the specified profile data.
func (c *RESTClient) UpdateUserProfile(ctx context.Context, id int64, profile *modelv2.UserProfile) error {
	_, err := c.GetUserByID(ctx, id)
	if err != nil {
		return err
	}

	params := &user.UpdateUserProfileParams{
		Profile: profile,
		UserID:  id,
		Context: ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err = c.V2Client.User.UpdateUserProfile(params, c.AuthInfo)

	return handleSwaggerUserErrors(err)
}

// UpdateUserPassword updates a user's password from 'old' to 'new'.
// 'old' is an optional parameter when called by an administrator.
func (c *RESTClient) UpdateUserPassword(ctx context.Context, userID int64, passwordRequest *modelv2.PasswordReq) error {
	if passwordRequest.NewPassword == "" {
		return errors.New("no new password provided")
	}

	_, err := c.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	params := &user.UpdateUserPasswordParams{
		Password: passwordRequest,
		UserID:   userID,
		Context:  ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err = c.V2Client.User.UpdateUserPassword(params, c.AuthInfo)

	return handleSwaggerUserErrors(err)
}

// UserExists checks the user with the provided 'idOrName' for existence.
func (c *RESTClient) UserExists(ctx context.Context, idOrName intstr.IntOrString) (bool, error) {
	switch idOrName.Type {
	default:
		return false, nil
	case intstr.Int:
		if idOrName.Type == intstr.Int {
			_, err := c.GetUserByID(ctx, int64(idOrName.IntVal))
			if err != nil {
				if _, ok := err.(*clienterrors.ErrUserNotFound); ok {
					return false, nil
				}
				return false, err
			}
		}
	case intstr.String:
		_, err := c.GetUserByName(ctx, idOrName.StrVal)
		if err != nil {
			if _, ok := err.(*clienterrors.ErrUserNotFound); ok {
				return false, nil
			}
			return false, err
		}
	}

	return true, nil
}
