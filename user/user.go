package user

import (
	"context"
	"errors"

	"github.com/mittwald/goharbor-client/internal/api/v1_10_4/client"

	"github.com/go-openapi/runtime"

	"github.com/mittwald/goharbor-client/internal/api/v1_10_4/client/products"
	model "github.com/mittwald/goharbor-client/model/v1_10_4"
)

// RESTClient is a subclient for handling user related actions.
type RESTClient struct {
	// The swagger client
	Client *client.Harbor

	// AuthInfo contain auth information, which are provided on API calls.
	AuthInfo runtime.ClientAuthInfoWriter
}

func NewClient(cl *client.Harbor, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		Client:   cl,
		AuthInfo: authInfo,
	}
}

type Client interface {
	NewUser(ctx context.Context, username, email, realname, password, comments string)
	GetUser(ctx context.Context, username string) (*model.User, error)
	DeleteUser(ctx context.Context, u *model.User) error
	UpdateUser(ctx context.Context, u *model.User) error
	UpdateUserPassword(ctx context.Context, id int64, password *model.Password) error
}

// NewUser creates and returns a new user, or error in case of failure.
// Username is a unique username.
// email is the Email of the user.
// realname is the fullname of the user.
// password is the password for this user.
// comments as a comment attached to the user.
func (c *RESTClient) NewUser(ctx context.Context, username, email, realname, password,
	comments string) (*model.User, error) {
	uReq := &model.User{
		Username: username,
		Password: password,
		Email:    email,
		Realname: realname,
		Comment:  comments,
	}

	_, err := c.Client.Products.PostUsers(&products.PostUsersParams{
		User:    uReq,
		Context: ctx,
	}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerUserErrors(err)
	}

	user, err := c.GetUser(ctx, username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUser returns an existing user or an error in case of failure.
func (c *RESTClient) GetUser(ctx context.Context, username string) (*model.User, error) {
	if username == "" {
		return nil, errors.New("no username provided")
	}

	resp, err := c.Client.Products.GetUsers(&products.GetUsersParams{
		Context:  ctx,
		Username: &username,
	}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerUserErrors(err)
	}

	for _, v := range resp.Payload {
		if v.Username == username {
			return v, nil
		}
	}

	return nil, &ErrUserNotFound{}
}

// DeleteUser deletes the specified user.
func (c *RESTClient) DeleteUser(ctx context.Context, u *model.User) error {
	if u == nil {
		return errors.New("no user provided")
	}

	user, err := c.GetUser(ctx, u.Username)
	if err != nil {
		return err
	}

	if u.UserID != user.UserID {
		return &ErrUserMismatch{}
	}

	_, err = c.Client.Products.DeleteUsersUserID(&products.DeleteUsersUserIDParams{
		UserID:  user.UserID,
		Context: ctx,
	}, c.AuthInfo)

	return handleSwaggerUserErrors(err)
}

// UpdateUser updates a user with the specified data.
// Note that only realname, email and comment properties are updateable.
func (c *RESTClient) UpdateUser(ctx context.Context, u *model.User) error {
	if u == nil {
		return errors.New("no user provided")
	}

	user, err := c.GetUser(ctx, u.Username)
	if err != nil {
		return err
	}

	profile := &model.UserProfile{
		Comment:  u.Comment,
		Email:    u.Email,
		Realname: u.Realname,
	}

	if u.UserID != user.UserID {
		return &ErrUserMismatch{}
	}

	_, err = c.Client.Products.PutUsersUserID(&products.PutUsersUserIDParams{
		UserID:  user.UserID,
		Profile: profile,
		Context: ctx,
	}, c.AuthInfo)

	return handleSwaggerUserErrors(err)
}

// UpdateUserPassword updates a users password
func (c *RESTClient) UpdateUserPassword(ctx context.Context, id int64, password *model.Password) error {
	if password == nil {
		return errors.New("no password provided")
	}

	_, err := c.Client.Products.PutUsersUserIDPassword(&products.PutUsersUserIDPasswordParams{
		Password: password,
		UserID:   id,
		Context:  ctx,
	}, c.AuthInfo)

	return handleSwaggerUserErrors(err)
}

func (c *RESTClient) UserExists(ctx context.Context, u *model.User) (bool, error) {
	_, err := c.GetUser(ctx, u.Username)
	if err != nil {
		if _, ok := err.(*ErrUserNotFound); ok {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}
