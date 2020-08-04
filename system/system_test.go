// +build !integration

package system

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/go-openapi/runtime"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/mittwald/goharbor-client/internal/api/v1_10_0/client"
	"github.com/mittwald/goharbor-client/internal/api/v1_10_0/client/products"
	"github.com/mittwald/goharbor-client/mocks"
	model "github.com/mittwald/goharbor-client/model/v1_10_0"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	authInfo            = runtimeclient.BasicAuth("foo", "bar")
	exampleCron         = "0 * * * *"
	exampleScheduleType = "Hourly"
	gcReq               = &model.AdminJobSchedule{
		Schedule: &model.AdminJobScheduleObj{
			Cron: exampleCron,
			Type: exampleScheduleType,
		},
	}
	getGcParams = &products.GetSystemGcScheduleParams{
		Context: context.Background(),
	}
	postGcParams = &products.PostSystemGcScheduleParams{
		Schedule: gcReq,
		Context:  context.Background(),
	}
)

func TestRESTClient_NewSystemGarbageCollection(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	p.On("PostSystemGcSchedule", postGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PostSystemGcScheduleOK{}, nil)

	p.On("GetSystemGcSchedule", getGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetSystemGcScheduleOK{Payload: gcReq}, nil)

	_, err := cl.NewSystemGarbageCollection(ctx, exampleCron, exampleScheduleType)

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_NewSystemGarbageCollection_ErrSystemInvalidSchedule(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	p.On("PostSystemGcSchedule", postGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PostSystemGcScheduleOK{}, &runtime.APIError{Code: http.StatusBadRequest})

	_, err := cl.NewSystemGarbageCollection(ctx, exampleCron, exampleScheduleType)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrSystemInvalidSchedule{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_NewSystemGarbageCollection_StatusUnauthorized(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	p.On("PostSystemGcSchedule", postGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PostSystemGcScheduleOK{}, nil)

	p.On("GetSystemGcSchedule", getGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetSystemGcScheduleOK{Payload: gcReq}, &runtime.APIError{Code: http.StatusUnauthorized})

	_, err := cl.NewSystemGarbageCollection(ctx, exampleCron, exampleScheduleType)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrSystemUnauthorized{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_NewSystemGarbageCollection_StatusCreated(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	p.On("PostSystemGcSchedule", postGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PostSystemGcScheduleOK{}, nil)

	p.On("GetSystemGcSchedule", getGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetSystemGcScheduleOK{Payload: gcReq}, &runtime.APIError{Code: http.StatusCreated})

	_, err := cl.NewSystemGarbageCollection(ctx, exampleCron, exampleScheduleType)

	assert.Nil(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_NewSystemGarbageCollection_ErrSystemGcInProgress(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	p.On("PostSystemGcSchedule", postGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PostSystemGcScheduleOK{}, nil)

	p.On("GetSystemGcSchedule", getGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetSystemGcScheduleOK{Payload: gcReq}, &runtime.APIError{Code: http.StatusConflict})

	_, err := cl.NewSystemGarbageCollection(ctx, exampleCron, exampleScheduleType)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrSystemGcInProgress{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_NewSystemGarbageCollection_ErrSystemInternalErrors(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	p.On("PostSystemGcSchedule", postGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PostSystemGcScheduleOK{}, nil)

	p.On("GetSystemGcSchedule", getGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetSystemGcScheduleOK{Payload: gcReq}, &runtime.APIError{Code: http.StatusInternalServerError})

	_, err := cl.NewSystemGarbageCollection(ctx, exampleCron, exampleScheduleType)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrSystemInternalErrors{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_NewSystemGarbageCollection_ErrSystemGcInProgress_2(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	p.On("PostSystemGcSchedule", postGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PostSystemGcScheduleOK{}, nil)

	p.On("GetSystemGcSchedule", getGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetSystemGcScheduleOK{Payload: gcReq}, &products.PostSystemGcScheduleConflict{})

	_, err := cl.NewSystemGarbageCollection(ctx, exampleCron, exampleScheduleType)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrSystemGcInProgress{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_NewSystemGarbageCollection_ErrSystemInvalidSchedule_2(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	p.On("PostSystemGcSchedule", postGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PostSystemGcScheduleOK{}, nil)

	p.On("GetSystemGcSchedule", getGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetSystemGcScheduleOK{Payload: gcReq}, &products.PutSystemGcScheduleBadRequest{})

	_, err := cl.NewSystemGarbageCollection(ctx, exampleCron, exampleScheduleType)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrSystemInvalidSchedule{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_GetSystemGarbageCollection(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	p.On("GetSystemGcSchedule", getGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetSystemGcScheduleOK{Payload: gcReq}, nil)

	gc, err := cl.GetSystemGarbageCollection(ctx)

	assert.NoError(t, err)

	assert.IsType(t, &model.AdminJobSchedule{}, gc)
	assert.Equal(t, gc.Schedule, gcReq.Schedule)

	p.AssertExpectations(t)
}

func TestRESTClient_GetSystemGarbageCollection_ScheduleNil(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	nilSchedulegcReq := &model.AdminJobSchedule{
		Schedule: nil,
	}

	p.On("GetSystemGcSchedule", getGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetSystemGcScheduleOK{Payload: nilSchedulegcReq}, nil)

	_, err := cl.GetSystemGarbageCollection(ctx)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrSystemGcUndefined{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_UpdateSystemGarbageCollection(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	newGcReq := model.AdminJobSchedule{
		Schedule: &model.AdminJobScheduleObj{
			Cron: "1 * * * *",
			Type: "Hourly",
		},
	}

	putGcParams := &products.PutSystemGcScheduleParams{
		Schedule: &newGcReq,
		Context:  ctx,
	}

	p.On("GetSystemGcSchedule", getGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetSystemGcScheduleOK{Payload: gcReq}, nil)

	p.On("PutSystemGcSchedule", putGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PutSystemGcScheduleOK{}, nil)

	err := cl.UpdateSystemGarbageCollection(ctx, newGcReq.Schedule)

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_UpdateSystemGarbageCollection_ScheduleNil(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	nilGcReq := model.AdminJobSchedule{
		Schedule: nil,
	}

	err := cl.UpdateSystemGarbageCollection(ctx, nilGcReq.Schedule)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrSystemGcScheduleNotProvided{}, err)
	}
}

func TestRESTClient_UpdateSystemGarbageCollection_ScheduleIdentical(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	p.On("GetSystemGcSchedule", getGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetSystemGcScheduleOK{Payload: gcReq}, nil)

	err := cl.UpdateSystemGarbageCollection(ctx, gcReq.Schedule)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrSystemGcScheduleIdentical{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_UpdateSystemGarbageCollection_ErrSystemNoPermission(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	newGcReq := model.AdminJobSchedule{
		Schedule: &model.AdminJobScheduleObj{
			Cron: "1 * * * *",
			Type: "Hourly",
		},
	}

	p.On("GetSystemGcSchedule", getGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetSystemGcScheduleOK{Payload: gcReq}, &runtime.APIError{Code: http.StatusForbidden})

	err := cl.UpdateSystemGarbageCollection(ctx, newGcReq.Schedule)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrSystemNoPermission{}, err)
	}

	p.AssertExpectations(t)
}

func TestRESTClient_ResetSystemGarbageCollection(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	putGcParams := &products.PutSystemGcScheduleParams{
		Context: ctx,
		Schedule: &model.AdminJobSchedule{
			Schedule: &model.AdminJobScheduleObj{
				Type: "None",
			},
		},
	}

	p.On("PutSystemGcSchedule", putGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PutSystemGcScheduleOK{}, nil)

	err := cl.ResetSystemGarbageCollection(ctx)

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_Health(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	healthParams := &products.GetHealthParams{
		Context: ctx,
	}

	p.On("GetHealth", healthParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetHealthOK{Payload: &model.OverallHealthStatus{}}, nil)

	_, err := cl.Health(ctx)

	assert.NoError(t, err)

	p.AssertExpectations(t)
}

func TestRESTClient_Health_ErrorReturn(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	healthParams := &products.GetHealthParams{
		Context: ctx,
	}

	p.On("GetHealth", healthParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetHealthOK{Payload: &model.OverallHealthStatus{}},
			errors.New("err"))

	_, err := cl.Health(ctx)

	if assert.Error(t, err) {
		assert.Equal(t, errors.New("err"), err)
	}

	p.AssertExpectations(t)
}

func TestErrSystemGcInProgress_Error(t *testing.T) {
	var e ErrSystemGcInProgress

	assert.Equal(t, ErrSystemGcInProgressMsg, e.Error())
}

func TestErrSystemGcScheduleIdentical_Error(t *testing.T) {
	var e ErrSystemGcScheduleIdentical

	assert.Equal(t, ErrSystemGcScheduleIdenticalMsg, e.Error())
}

func TestErrSystemGcScheduleNotProvided_Error(t *testing.T) {
	var e ErrSystemGcScheduleNotProvided

	assert.Equal(t, ErrSystemGcScheduleNotProvidedMsg, e.Error())
}

func TestErrSystemGcUndefined_Error(t *testing.T) {
	var e ErrSystemGcUndefined

	assert.Equal(t, ErrSystemGcUndefinedMsg, e.Error())
}

func TestErrSystemInternalErrors_Error(t *testing.T) {
	var e ErrSystemInternalErrors

	assert.Equal(t, ErrSystemInternalErrorsMsg, e.Error())
}

func TestErrSystemInvalidSchedule_Error(t *testing.T) {
	var e ErrSystemInvalidSchedule

	assert.Equal(t, ErrSystemInvalidScheduleMsg, e.Error())
}

func TestErrSystemNoPermission_Error(t *testing.T) {
	var e ErrSystemNoPermission

	assert.Equal(t, ErrSystemNoPermissionMsg, e.Error())
}

func TestErrSystemUnauthorized_Error(t *testing.T) {
	var e ErrSystemUnauthorized

	assert.Equal(t, ErrSystemUnauthorizedMsg, e.Error())
}
