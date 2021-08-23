package user

import (
	"context"
	"errors"

	v2client "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v4/apiv2/internal/legacyapi/client"

	"github.com/go-openapi/runtime"

	"github.com/mittwald/goharbor-client/v4/apiv2/internal/legacyapi/client/products"
	legacymodel "github.com/mittwald/goharbor-client/v4/apiv2/model/legacy"
)

// RESTClient is a subclient for handling user related actions.
type RESTClient struct {
	// The legacy swagger client
	LegacyClient *client.Harbor

	// The new client of the harbor v2 API
	V2Client *v2client.Harbor

	// AuthInfo contains the auth information that is provided on API calls.
	AuthInfo runtime.ClientAuthInfoWriter
}

func NewClient(legacyClient *client.Harbor, v2Client *v2client.Harbor, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		LegacyClient: legacyClient,
		V2Client:     v2Client,
		AuthInfo:     authInfo,
	}
}

type Client interface {
	NewUser(ctx context.Context, username, email, realname, password,
		comments string) (*legacymodel.User, error)
	GetUser(ctx context.Context, username string) (*legacymodel.User, error)
	GetUserByID(ctx context.Context, id int64) (*legacymodel.User, error)
	ListUsers(ctx context.Context) ([]*legacymodel.User, error)
	DeleteUser(ctx context.Context, u *legacymodel.User) error
	UpdateUser(ctx context.Context, u *legacymodel.User) error
	UpdateUserPassword(ctx context.Context, id int64, password *legacymodel.Password) error
	UserExists(ctx context.Context, u *legacymodel.User) (bool, error)
}

// NewUser creates and returns a new user, or error in case of failure.
// Username is a unique username.
// email is the Email of the user.
// realname is the fullname of the user.
// password is the password for this user.
// comments as a comment attached to the user.
func (c *RESTClient) NewUser(ctx context.Context, username, email, realname, password,
	comments string) (*legacymodel.User, error) {
	uReq := &legacymodel.User{
		Username: username,
		Password: password,
		Email:    email,
		Realname: realname,
		Comment:  comments,
	}

	_, err := c.LegacyClient.Products.PostUsers(&products.PostUsersParams{
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
func (c *RESTClient) GetUser(ctx context.Context, username string) (*legacymodel.User, error) {
	if username == "" {
		return nil, errors.New("no username provided")
	}

	resp, err := c.LegacyClient.Products.GetUsers(&products.GetUsersParams{
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

// ListUsers lists and returns all registered Harbor users.
func (c *RESTClient) ListUsers(ctx context.Context) ([]*legacymodel.User, error) {
	resp, err := c.LegacyClient.Products.GetUsers(&products.GetUsersParams{
		Context: ctx,
	}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerUserErrors(err)
	}

	return resp.Payload, nil
}

// GetUserByID fetches a registered user by the provided user id.
// Returns an error if no user could be found, or if the id is '0'.
func (c *RESTClient) GetUserByID(ctx context.Context, id int64) (*legacymodel.User, error) {
	if id <= 0 {
		return nil, &ErrUserInvalidID{}
	}

	resp, err := c.LegacyClient.Products.GetUsersUserID(&products.GetUsersUserIDParams{
		UserID:  id,
		Context: ctx,
	}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerUserErrors(err)
	}

	if resp.Payload.UserID != id {
		return nil, &ErrUserMismatch{}
	}

	return resp.Payload, nil
}

// DeleteUser deletes the specified user.
func (c *RESTClient) DeleteUser(ctx context.Context, u *legacymodel.User) error {
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

	_, err = c.LegacyClient.Products.DeleteUsersUserID(&products.DeleteUsersUserIDParams{
		UserID:  user.UserID,
		Context: ctx,
	}, c.AuthInfo)

	return handleSwaggerUserErrors(err)
}

// UpdateUser updates a user with the specified data.
// Note that only realname, email and comment properties are updateable.
func (c *RESTClient) UpdateUser(ctx context.Context, u *legacymodel.User) error {
	if u == nil {
		return errors.New("no user provided")
	}

	user, err := c.GetUser(ctx, u.Username)
	if err != nil {
		return err
	}

	profile := &legacymodel.UserProfile{
		Comment:  u.Comment,
		Email:    u.Email,
		Realname: u.Realname,
	}

	if u.UserID != user.UserID {
		return &ErrUserMismatch{}
	}

	_, err = c.LegacyClient.Products.PutUsersUserID(&products.PutUsersUserIDParams{
		UserID:  user.UserID,
		Profile: profile,
		Context: ctx,
	}, c.AuthInfo)

	return handleSwaggerUserErrors(err)
}

// UpdateUserPassword updates a users password
func (c *RESTClient) UpdateUserPassword(ctx context.Context, id int64, password *legacymodel.Password) error {
	if password == nil {
		return errors.New("no password provided")
	}

	_, err := c.LegacyClient.Products.PutUsersUserIDPassword(&products.PutUsersUserIDPasswordParams{
		Password: password,
		UserID:   id,
		Context:  ctx,
	}, c.AuthInfo)

	return handleSwaggerUserErrors(err)
}

func (c *RESTClient) UserExists(ctx context.Context, u *legacymodel.User) (bool, error) {
	_, err := c.GetUser(ctx, u.Username)
	if err != nil {
		if _, ok := err.(*ErrUserNotFound); ok {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
