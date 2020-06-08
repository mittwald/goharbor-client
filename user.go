package goharborclient

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-openapi/runtime"

	"github.com/mittwald/goharbor-client/api/v1.10.0/client/products"
	"github.com/mittwald/goharbor-client/api/v1.10.0/model"
)

const (
	ErrUserNotFound     = "user not found on server side"
	ErrUserBadRequest   = "Unsatisfied with constraints of the user creation/modification."
	ErrUserMismatch     = "id/name pair not found on server side"
	ErrUserAlreadyExist = "User with this username already exists."
	ErrUserInvalidID    = "Invalid user ID."
)

// UserError is an error describing a errors related to project operations
// and implements the error interface.
type UserError struct {
	// Name of the related project. Empty string means undefined.
	Username string

	// Error message of the related project.
	errorMessage string
}

// Error implements the Error interface.
func (p *UserError) Error() string {
	return fmt.Sprintf("%s (username: %s)",
		p.errorMessage, p.Username)
}

// NewUserError creates a new ProjectError.
func NewUserError(msg string, name string) error {
	return &UserError{
		Username:     name,
		errorMessage: msg,
	}
}

// UserRESTClient is a subclient for RESTClient handling user related actions.
type UserRESTClient struct {
	parent *RESTClient
}

// NewUser creates and returns a new user, or error in case of failure.
// Username is a unique username.
// email is the Email of the user.
// realname is the fullname of the user.
// password is the password for this user.
// comments as a comment attached to the user.
func (c *UserRESTClient) NewUser(ctx context.Context, username, email, realname, password, comments string) (*model.User, error) {
	uReq := &model.User{
		Username: username,
		Password: password,
		Email:    email,
		Realname: realname,
		Comment:  comments,
	}

	_, err := c.parent.Client.Products.PostUsers(&products.PostUsersParams{
		User:    uReq,
		Context: ctx,
	}, c.parent.AuthInfo)

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
func (c *UserRESTClient) Get(ctx context.Context, username string) (*model.User, error) {
	if username == "" {
		return nil, errors.New("no username provided")
	}

	resp, err := c.parent.Client.Products.GetUsers(&products.GetUsersParams{
		Context:  ctx,
		Username: &username,
	}, c.parent.AuthInfo)

	err = handleSwaggerUserErrors(err, username)
	if err != nil {
		return nil, err
	}

	for _, v := range resp.Payload {
		if v.Username == username {
			return v, nil
		}
	}

	return nil, NewUserError(ErrUserNotFound, username)
}

// Delete deletes a user from Harbor given by a user model,
// or error in case of failure.
func (c *UserRESTClient) Delete(ctx context.Context, u *model.User) error {
	if u == nil {
		return errors.New("no user provided")
	}

	user, err := c.Get(ctx, u.Username)
	if err != nil {
		return err
	}

	if u.UserID != user.UserID {
		return NewUserError(ErrUserMismatch, u.Username)
	}

	_, err = c.parent.Client.Products.DeleteUsersUserID(&products.DeleteUsersUserIDParams{
		UserID:  user.UserID,
		Context: ctx,
	}, c.parent.AuthInfo)

	return handleSwaggerUserErrors(err, user.Username)
}

// Update updates a user given by a user model,
// or error in case of failure.
// Note that only realname, email and comment properties are updateable.
func (c *UserRESTClient) Update(ctx context.Context, u *model.User) error {
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
		return NewUserError(ErrUserMismatch, u.Username)
	}

	_, err = c.parent.Client.Products.PutUsersUserID(&products.PutUsersUserIDParams{
		UserID:  user.UserID,
		Profile: profile,
		Context: ctx,
	}, c.parent.AuthInfo)

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
			return NewUserError(ErrUserAlreadyExist, username)
		}
	}

	switch in.(type) {
	case *products.PostUsersBadRequest:
		return NewUserError(ErrUserBadRequest, username)
	case *products.PutUsersUserIDBadRequest:
		return NewUserError(ErrUserInvalidID, username)
	default:
		return in
	}
}
