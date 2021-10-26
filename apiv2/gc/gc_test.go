// +build !integration

package gc

import (
	"context"
	"net/http"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	v2client "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/gc"
	modelv2 "github.com/mittwald/goharbor-client/v5/apiv2/model"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/mittwald/goharbor-client/v5/apiv2/internal/legacyapi/client"
	"github.com/mittwald/goharbor-client/v5/apiv2/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	authInfo            = runtimeclient.BasicAuth("foo", "bar")
	getGcScheduleParams = &gc.GetGCScheduleParams{
		Context: context.Background(),
	}
	postGcParams = &gc.CreateGCScheduleParams{
		Context:  context.Background(),
		Schedule: schedule,
	}
	schedule = &modelv2.Schedule{
		Schedule: &modelv2.ScheduleObj{
			Cron: "0 * * * *",
			Type: "Hourly",
		},
	}
)

func BuildLegacyClientWithMock(service *mocks.MockProductsClientService) *client.Harbor {
	return &client.Harbor{
		Products: service,
	}
}

func BuildGCClientWithMocks(gc *mocks.MockGcClientService) *v2client.Harbor {
	return &v2client.Harbor{
		Gc: gc,
	}
}

func TestRESTClient_NewGarbageCollection(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	g := &mocks.MockGcClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildGCClientWithMocks(g)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	g.On("CreateGCSchedule", postGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&gc.CreateGCScheduleCreated{}, nil)

	err := cl.NewGarbageCollection(ctx, schedule)

	assert.NoError(t, err)

	g.AssertExpectations(t)
}

func TestRESTClient_NewGarbageCollection_ErrSystemInvalidSchedule(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	g := &mocks.MockGcClientService{}
	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildGCClientWithMocks(g)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	g.On("CreateGCSchedule", postGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &runtime.APIError{Code: http.StatusBadRequest})

	err := cl.NewGarbageCollection(ctx, schedule)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrSystemInvalidSchedule{}, err)
	}

	g.AssertExpectations(t)
}

func TestRESTClient_NewGarbageCollection_StatusUnauthorized(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	g := &mocks.MockGcClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildGCClientWithMocks(g)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	g.On("CreateGCSchedule", postGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &runtime.APIError{Code: http.StatusUnauthorized})

	err := cl.NewGarbageCollection(ctx, schedule)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrSystemUnauthorized{}, err)
	}

	g.AssertExpectations(t)
}

func TestRESTClient_NewGarbageCollection_ErrSystemGcInProgress(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	g := &mocks.MockGcClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildGCClientWithMocks(g)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	g.On("CreateGCSchedule", postGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &runtime.APIError{Code: http.StatusConflict})

	err := cl.NewGarbageCollection(ctx, schedule)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrSystemGcInProgress{}, err)
	}

	g.AssertExpectations(t)
}

func TestRESTClient_NewGarbageCollection_ErrSystemInternalErrors(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	g := &mocks.MockGcClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildGCClientWithMocks(g)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	g.On("CreateGCSchedule", postGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &runtime.APIError{Code: http.StatusInternalServerError})

	err := cl.NewGarbageCollection(ctx, schedule)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrSystemInternalErrors{}, err)
	}

	g.AssertExpectations(t)
}

func TestRESTClient_NewGarbageCollection_ErrSystemGcInProgress_2(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	g := &mocks.MockGcClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildGCClientWithMocks(g)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	g.On("CreateGCSchedule", postGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &runtime.APIError{Code: http.StatusConflict})

	err := cl.NewGarbageCollection(ctx, schedule)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrSystemGcInProgress{}, err)
	}

	g.AssertExpectations(t)
}

func TestRESTClient_NewGarbageCollection_ErrSystemInvalidSchedule_2(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	g := &mocks.MockGcClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildGCClientWithMocks(g)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	g.On("CreateGCSchedule", postGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &runtime.APIError{Code: http.StatusBadRequest})

	err := cl.NewGarbageCollection(ctx, schedule)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrSystemInvalidSchedule{}, err)
	}

	g.AssertExpectations(t)
}

func TestRESTClient_GetGarbageCollection(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	g := &mocks.MockGcClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildGCClientWithMocks(g)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	heldGC := &modelv2.GCHistory{
		Schedule: &modelv2.ScheduleObj{
			Cron: schedule.Schedule.Cron,
			Type: schedule.Schedule.Type,
		},
	}

	g.On("GetGCSchedule", getGcScheduleParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&gc.GetGCScheduleOK{Payload: heldGC}, nil)

	gc, err := cl.GetGarbageCollectionSchedule(ctx)

	assert.NoError(t, err)

	assert.IsType(t, &modelv2.GCHistory{}, gc)
	assert.Equal(t, gc.Schedule, schedule.Schedule)

	g.AssertExpectations(t)
}

func TestRESTClient_GetGarbageCollectionSchedule_ScheduleNil(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	g := &mocks.MockGcClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildGCClientWithMocks(g)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	nilSchedulegcReq := &modelv2.GCHistory{
		Schedule: nil,
	}

	g.On("GetGCSchedule", getGcScheduleParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&gc.GetGCScheduleOK{Payload: nilSchedulegcReq}, nil)

	_, err := cl.GetGarbageCollectionSchedule(ctx)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrSystemGcScheduleUndefined{}, err)
	}

	g.AssertExpectations(t)
}

func TestRESTClient_UpdateGarbageCollection(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	g := &mocks.MockGcClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildGCClientWithMocks(g)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	newGcReq := &modelv2.Schedule{
		Parameters: map[string]interface{}{
			"delete_untagged": false,
		},
		Schedule: &modelv2.ScheduleObj{
			Cron: "1 * * * *",
			Type: "Hourly",
		},
		UpdateTime: strfmt.DateTime{},
	}

	putGCParams := gc.UpdateGCScheduleParams{
		Schedule: &modelv2.Schedule{
			Schedule:   newGcReq.Schedule,
			Parameters: newGcReq.Parameters,
		},
		Context: ctx,
	}

	g.On("UpdateGCSchedule", &putGCParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&gc.UpdateGCScheduleOK{}, nil)

	err := cl.UpdateGarbageCollection(ctx, newGcReq)

	assert.NoError(t, err)

	g.AssertExpectations(t)
}

func TestRESTClient_UpdateGarbageCollection_ScheduleNil(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	g := &mocks.MockGcClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildGCClientWithMocks(g)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	err := cl.UpdateGarbageCollection(ctx, nil)

	if assert.Error(t, err) {
		assert.IsType(t, &ErrSystemGcScheduleNotProvided{}, err)
	}
}

func TestRESTClient_ResetGarbageCollection(t *testing.T) {
	p := &mocks.MockProductsClientService{}
	g := &mocks.MockGcClientService{}

	legacyClient := BuildLegacyClientWithMock(p)
	v2Client := BuildGCClientWithMocks(g)

	cl := NewClient(legacyClient, v2Client, authInfo)

	ctx := context.Background()

	putGcParams := &gc.UpdateGCScheduleParams{
		Context: ctx,
		Schedule: &modelv2.Schedule{
			Parameters: map[string]interface{}{
				"delete_untagged": false,
			},
			Schedule: &modelv2.ScheduleObj{
				Type: "None",
			},
		},
	}

	g.On("UpdateGCSchedule", putGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&gc.UpdateGCScheduleOK{}, nil)

	err := cl.ResetGarbageCollection(ctx)

	assert.NoError(t, err)

	g.AssertExpectations(t)
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
