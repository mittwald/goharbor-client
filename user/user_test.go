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
)

func TestRESTClient_NewUser(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	uReq := &model.User{
		Username: exampleUser,
		Password: examplePassword,
		Email:    exampleEmail,
		Realname: "",
		Comment:  "",
	}

	ctx := context.Background()

	postParams := &products.PostUsersParams{
		User:    uReq,
		Context: ctx,
	}

	p.On("PostUsers",
		postParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PostUsersCreated{},
			nil)

	getParams := &products.GetUsersParams{
		Context:  ctx,
		Username: &exampleUser}

	p.On("GetUsers", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetUsersOK{
			Payload: []*model.User{{Username: "example-user"}},
		}, nil)

	_, err := cl.NewUser(ctx, exampleUser, exampleEmail, "", examplePassword, "")

	assert.NoError(t, err)

	p.AssertExpectations(t)
}
