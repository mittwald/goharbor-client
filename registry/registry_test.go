// +build !integration

package registry

import (
	"context"
	"testing"

	"github.com/go-openapi/runtime"

	"github.com/mittwald/goharbor-client/internal/api/v1_10_0/client"
	"github.com/mittwald/goharbor-client/internal/api/v1_10_0/client/products"
	"github.com/mittwald/goharbor-client/mocks"
	model "github.com/mittwald/goharbor-client/model/v1_10_0"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const (
	name string = "test-registry"
)

var authInfo = runtimeclient.BasicAuth("foo", "bar")

func TestNewClient(t *testing.T) {
	swaggerClient := client.New(runtimeclient.New("foobar:30002", "/api",
		[]string{"http"}), strfmt.Default)
	authInfo := runtimeclient.BasicAuth("foo", "bar")
	c := NewClient(swaggerClient, authInfo)

	require.NotNil(t, c)
	assert.NotNil(t, c.AuthInfo)
	assert.NotNil(t, c.Client)

	assert.Equal(t, swaggerClient, c.Client)
}

func TestRESTClient_NewRegistry(t *testing.T) {
	req := &model.Registry{
		Credential: &model.RegistryCredential{},
		Insecure:   true,
		Name:       name,
		Type:       "harbor",
		URL:        "http://test.reg",
	}
	ctx := context.Background()

	p := &mocks.MockClientService{}
	p.On("PostRegistries",
		&products.PostRegistriesParams{
			Registry: req,
			Context:  ctx,
		},
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc"),
	).Return(&products.PostRegistriesCreated{}, nil)
	p.On("GetRegistries",
		&products.GetRegistriesParams{
			Name:    &req.Name,
			Context: ctx,
		}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc"),
	).Return(
		&products.GetRegistriesOK{
			Payload: []*model.Registry{{
				CreationTime: "",
				Credential:   nil,
				Description:  "",
				ID:           10,
				Insecure:     req.Insecure,
				Name:         req.Name,
				Status:       "",
				Type:         req.Type,
				UpdateTime:   "",
				URL:          req.URL,
			}},
		}, nil)

	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)

	r, err := cl.NewRegistry(ctx, req.Name, req.Type,
		req.URL, &model.RegistryCredential{}, true)

	p.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, int64(10), r.ID)
	assert.Equal(t, req.Name, r.Name)
	assert.Equal(t, req.URL, r.URL)
	assert.Equal(t, req.Type, r.Type)
	assert.Equal(t, req.Insecure, r.Insecure)
}

func TestRESTClient_NewRegistry_ErrOnPOST(t *testing.T) {
	req := &model.Registry{
		Credential: &model.RegistryCredential{},
		Insecure:   true,
		Name:       name,
		Type:       "harbor",
		URL:        "http://test.reg",
	}
	ctx := context.Background()

	p := &mocks.MockClientService{}
	p.On("PostRegistries",
		&products.PostRegistriesParams{
			Registry: req,
			Context:  ctx,
		},
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc"),
	).Return(nil, &runtime.APIError{
		OperationName: "",
		Response:      nil,
		Code:          400,
	})

	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)

	r, err := cl.NewRegistry(ctx, req.Name, req.Type,
		req.URL, &model.RegistryCredential{}, true)

	p.AssertExpectations(t)
	assert.Nil(t, r)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrRegistryIllegalIDFormat{}, err)
	}
}

func TestRESTClient_GetRegistry(t *testing.T) {
	name := name
	insecure := true
	regType := "harbor"
	url := "http://foo.bar"
	id := int64(11)
	ctx := context.Background()

	p := &mocks.MockClientService{}
	p.On("GetRegistries",
		&products.GetRegistriesParams{
			Name:    &name,
			Context: ctx,
		}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc"),
	).Return(&products.GetRegistriesOK{
		Payload: []*model.Registry{
			{
				CreationTime: "",
				Credential:   nil,
				Description:  "",
				ID:           id,
				Insecure:     insecure,
				Name:         name,
				Status:       "",
				Type:         regType,
				UpdateTime:   "",
				URL:          url,
			},
		},
	}, nil)

	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)
	r, err := cl.GetRegistry(ctx, name)

	p.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, id, r.ID)
	assert.Equal(t, name, r.Name)
	assert.Equal(t, url, r.URL)
	assert.Equal(t, regType, r.Type)
	assert.Equal(t, insecure, r.Insecure)
}

func TestRESTClient_GetRegistry_NotFound(t *testing.T) {
	name := name
	ctx := context.Background()

	p := &mocks.MockClientService{}
	p.On("GetRegistries",
		&products.GetRegistriesParams{
			Name:    &name,
			Context: ctx,
		}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc"),
	).Return(&products.GetRegistriesOK{
		Payload: []*model.Registry{},
	}, nil)

	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)
	r, err := cl.GetRegistry(ctx, name)

	p.AssertExpectations(t)
	assert.Nil(t, r)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrRegistryNotFound{}, err)
	}
}

func TestRESTClient_GetRegistry_ErrOnGet(t *testing.T) {
	name := name
	ctx := context.Background()
	errorMsg := "error on server side"
	errorCode := 500

	p := &mocks.MockClientService{}
	p.On("GetRegistries",
		&products.GetRegistriesParams{
			Name:    &name,
			Context: ctx,
		}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc"),
	).Return(nil, &runtime.APIError{
		OperationName: "",
		Response:      errorMsg,
		Code:          errorCode,
	})

	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)
	r, err := cl.GetRegistry(ctx, name)

	p.AssertExpectations(t)
	assert.Nil(t, r)
	if assert.Error(t, err) {
		require.IsType(t, &ErrRegistryInternalErrors{}, err)
	}
}

func TestRESTClient_GetRegistry_ErrRegistryNotProvided(t *testing.T) {
	p := &mocks.MockClientService{}
	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)
	r, err := cl.GetRegistry(context.Background(), "")

	assert.Nil(t, r)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrRegistryNotProvided{}, err)
	}
}

func TestRESTClient_DeleteRegistry(t *testing.T) {
	ctx := context.Background()
	registry := &model.Registry{
		CreationTime: "",
		Credential:   nil,
		Description:  "",
		ID:           10,
		Insecure:     false,
		Name:         "restregistry",
		Status:       "",
		Type:         "harbor",
		UpdateTime:   "",
		URL:          "http://foo.bar",
	}

	p := &mocks.MockClientService{}
	p.On("GetRegistries",
		&products.GetRegistriesParams{
			Name:    &registry.Name,
			Context: ctx,
		}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc"),
	).Return(&products.GetRegistriesOK{
		Payload: []*model.Registry{registry},
	}, nil)
	p.On("DeleteRegistriesID",
		&products.DeleteRegistriesIDParams{
			ID:      registry.ID,
			Context: ctx,
		}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc"),
	).Return(&products.DeleteRegistriesIDOK{}, nil)

	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)
	err := cl.DeleteRegistry(ctx, registry)

	p.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestRESTClient_DeleteRegistry_NotFound(t *testing.T) {
	ctx := context.Background()
	registry := &model.Registry{
		CreationTime: "",
		Credential:   nil,
		Description:  "",
		ID:           10,
		Insecure:     false,
		Name:         "restregistry",
		Status:       "",
		Type:         "harbor",
		UpdateTime:   "",
		URL:          "http://foo.bar",
	}

	p := &mocks.MockClientService{}
	p.On("GetRegistries",
		&products.GetRegistriesParams{
			Name:    &registry.Name,
			Context: ctx,
		}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc"),
	).Return(nil, &runtime.APIError{
		OperationName: "",
		Response:      "not found",
		Code:          404,
	})

	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)
	err := cl.DeleteRegistry(ctx, registry)

	p.AssertExpectations(t)
	if assert.Error(t, err) {
		require.IsType(t, &runtime.APIError{}, err)
		ty, _ := err.(*runtime.APIError)
		assert.Equal(t, 404, ty.Code)
		assert.Equal(t, "not found", ty.Response)
	}
}

func TestRESTClient_DeleteRegistry_ErrRegistryNotProvided(t *testing.T) {
	p := &mocks.MockClientService{}
	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)
	err := cl.DeleteRegistry(context.Background(), nil)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrRegistryNotProvided{}, err)
	}
}

func TestRESTClient_UpdateRegistry(t *testing.T) {
	ctx := context.Background()
	registry := &model.Registry{
		CreationTime: "",
		Credential: &model.RegistryCredential{
			AccessKey:    "",
			AccessSecret: "",
			Type:         "",
		},
		Description: "",
		ID:          10,
		Insecure:    false,
		Name:        "restregistry",
		Status:      "",
		Type:        "harbor",
		UpdateTime:  "",
		URL:         "http://foo.bar",
	}

	rReq := &model.PutRegistry{
		AccessKey:      registry.Credential.AccessKey,
		AccessSecret:   registry.Credential.AccessSecret,
		CredentialType: registry.Credential.Type,
		Description:    registry.Description,
		Insecure:       registry.Insecure,
		Name:           registry.Name,
		URL:            registry.URL,
	}

	p := &mocks.MockClientService{}

	cl := NewClient(&client.Harbor{Products: p, Transport: nil}, authInfo)

	p.On("GetRegistries",
		&products.GetRegistriesParams{
			Name:    &registry.Name,
			Context: ctx,
		}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.GetRegistriesOK{
			Payload: []*model.Registry{registry},
		}, nil)

	p.On("PutRegistriesID",
		&products.PutRegistriesIDParams{
			ID:         registry.ID,
			RepoTarget: rReq,
			Context:    ctx,
		}, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).Return(
		&products.PutRegistriesIDOK{}, nil)

	err := cl.UpdateRegistry(ctx, registry)

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestErrRegistryIDNotExists_Error(t *testing.T) {
	var e ErrRegistryIDNotExists

	assert.Equal(t, ErrRegistryIDNotExistsMsg, e.Error())
}

func TestErrRegistryIllegalIDFormat_Error(t *testing.T) {
	var e ErrRegistryIllegalIDFormat

	assert.Equal(t, ErrRegistryIllegalIDFormatMsg, e.Error())
}

func TestErrRegistryInternalErrors_Error(t *testing.T) {
	var e ErrRegistryInternalErrors

	assert.Equal(t, ErrRegistryInternalErrorsMsg, e.Error())
}

func TestErrRegistryMismatch_Error(t *testing.T) {
	var e ErrRegistryMismatch

	assert.Equal(t, ErrRegistryMismatchMsg, e.Error())
}

func TestErrRegistryNameAlreadyExists_Error(t *testing.T) {
	var e ErrRegistryNameAlreadyExists

	assert.Equal(t, ErrRegistryNameAlreadyExistsMsg, e.Error())
}

func TestErrRegistryNoPermission_Error(t *testing.T) {
	var e ErrRegistryNoPermission

	assert.Equal(t, ErrRegistryNoPermissionMsg, e.Error())
}

func TestErrRegistryNotFound_Error(t *testing.T) {
	var e ErrRegistryNotFound

	assert.Equal(t, ErrRegistryNotFoundMsg, e.Error())
}

func TestErrRegistryNotProvided_Error(t *testing.T) {
	var e ErrRegistryNotProvided

	assert.Equal(t, ErrRegistryNotProvidedMsg, e.Error())
}

func TestErrRegistryUnauthorized_Error(t *testing.T) {
	var e ErrRegistryUnauthorized

	assert.Equal(t, ErrRegistryUnauthorizedMsg, e.Error())
}
