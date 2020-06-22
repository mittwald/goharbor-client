package user

import (
	"context"
	"errors"
	"github.com/go-openapi/runtime"
	"github.com/mittwald/goharbor-client/internal/api/v1.10.0/client"

	"github.com/mittwald/goharbor-client/internal/api/v1.10.0/client/products"
	"github.com/mittwald/goharbor-client/internal/api/v1.10.0/model"
)

// UserRESTClient is a subclient for RESTClient handling user related actions.
type RESTClient struct {
	Client   *client.Harbor
	AuthInfo runtime.ClientAuthInfoWriter
}

func NewClient(cl *client.Harbor, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		Client:   cl,
		AuthInfo: authInfo,
	}
}

// New creates and returns a new user, or error in case of failure.
// Username is a unique username.
// email is the Email of the user.
// realname is the fullname of the user.
// password is the password for this user.
// comments as a comment attached to the user.
func (c *RESTClient) New(ctx context.Context, username, email, realname, password, comments string) (*model.User, error) {
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

	err = handleSwaggerUserErrors(err, username)
	if err != nil {
		return nil, err
	}

	user, err := c.Get(ctx, username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Get fetches a user from Harbor given by username,
// or an error in case of failure.
func (c *RESTClient) Get(ctx context.Context, username string) (*model.User, error) {
	if username == "" {
		return nil, errors.New("no username provided")
	}

	resp, err := c.Client.Products.GetUsers(&products.GetUsersParams{
		Context:  ctx,
		Username: &username,
	}, c.AuthInfo)

	err = handleSwaggerUserErrors(err, username)
	if err != nil {
		return nil, err
	}

	for _, v := range resp.Payload {
		if v.Username == username {
			return v, nil
		}
	}

	return nil, &ErrUserNotFound{}
}

// Delete deletes a user from Harbor given by a user model,
// or error in case of failure.
func (c *RESTClient) Delete(ctx context.Context, u *model.User) error {
	if u == nil {
		return errors.New("no user provided")
	}

	user, err := c.Get(ctx, u.Username)
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

	return handleSwaggerUserErrors(err, user.Username)
}

// Update updates a user given by a user model,
// or error in case of failure.
// Note that only realname, email and comment properties are updateable.
func (c *RESTClient) Update(ctx context.Context, u *model.User) error {
	if u == nil {
		return errors.New("no user provided")
	}

	user, err := c.Get(ctx, u.Username)
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

	return handleSwaggerUserErrors(err, user.Username)
}

// handleUserErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with proper message.
func handleSwaggerUserErrors(in error, username string) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case 409:
			return &ErrUserAlreadyExist{}
		}
	}

	switch in.(type) {
	case *products.PostUsersBadRequest:
		return &ErrUserBadRequest{}
	case *products.PutUsersUserIDBadRequest:
		return &ErrUserInvalidID{}
	default:
		return in
	}
}

// UserExists checks whether a user exists or not
func (c *RESTClient) UserExists(ctx context.Context, u *model.User) (bool, error) {
	_, err := c.Get(ctx, u.Username)

	if err != nil {
		if _, ok := err.(*ErrUserNotFound); ok {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}
