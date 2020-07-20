// +build !integration

package system

import (
	"context"
	"testing"

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
)

func TestRESTClient_NewSystemGarbageCollection(t *testing.T) {
	p := &mocks.MockClientService{}

	c := &client.Harbor{
		Products:  p,
		Transport: nil,
	}

	cl := NewClient(c, authInfo)

	ctx := context.Background()

	gcReq := &model.AdminJobSchedule{
		Schedule: &model.AdminJobScheduleObj{
			Cron: exampleCron,
			Type: exampleScheduleType,
		},
	}

	getGcParams := &products.GetSystemGcScheduleParams{
		Context: ctx,
	}

	postGcParams := &products.PostSystemGcScheduleParams{
		Schedule: gcReq,
		Context:  ctx,
	}

	p.On("PostSystemGcSchedule", postGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.PostSystemGcScheduleOK{}, nil)

	p.On("GetSystemGcSchedule", getGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetSystemGcScheduleOK{Payload: gcReq}, nil)

	_, err := cl.NewSystemGarbageCollection(ctx, exampleCron, exampleScheduleType)

	assert.NoError(t, err)

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

	gcReq := &model.AdminJobSchedule{
		Schedule: &model.AdminJobScheduleObj{
			Cron: exampleCron,
			Type: exampleScheduleType,
		},
	}

	getGcParams := &products.GetSystemGcScheduleParams{
		Context: ctx,
	}

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

	gcReq := &model.AdminJobSchedule{
		Schedule: nil,
	}

	getGcParams := &products.GetSystemGcScheduleParams{
		Context: ctx,
	}

	p.On("GetSystemGcSchedule", getGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetSystemGcScheduleOK{Payload: gcReq}, nil)

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

	gcReq := &model.AdminJobSchedule{
		Schedule: &model.AdminJobScheduleObj{
			Cron: exampleCron,
			Type: exampleScheduleType,
		},
	}

	newGcReq := model.AdminJobSchedule{
		Schedule: &model.AdminJobScheduleObj{
			Cron: "1 * * * *",
			Type: "Hourly",
		},
	}

	getGcParams := &products.GetSystemGcScheduleParams{
		Context: ctx,
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

	newGcReq := model.AdminJobSchedule{
		Schedule: nil,
	}

	err := cl.UpdateSystemGarbageCollection(ctx, newGcReq.Schedule)

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

	gcReq := &model.AdminJobSchedule{
		Schedule: &model.AdminJobScheduleObj{
			Cron: exampleCron,
			Type: exampleScheduleType,
		},
	}

	getGcParams := &products.GetSystemGcScheduleParams{
		Context: ctx,
	}

	p.On("GetSystemGcSchedule", getGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&products.GetSystemGcScheduleOK{Payload: gcReq}, nil)

	err := cl.UpdateSystemGarbageCollection(ctx, gcReq.Schedule)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrSystemGcScheduleIdentical{}, err)
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
