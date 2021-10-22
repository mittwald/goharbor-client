//go:build !integration

package gc

import (
	"context"
	"net/http"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/gc"
	modelv2 "github.com/mittwald/goharbor-client/v4/apiv2/model"
	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/common"
	unittesting "github.com/mittwald/goharbor-client/v4/apiv2/pkg/testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/mittwald/goharbor-client/v4/apiv2/mocks"
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
	ctx = context.Background()
)

func APIandMockClientsForTests() (*RESTClient, *unittesting.MockClients) {
	desiredMockClients := &unittesting.MockClients{
		Project: mocks.MockProjectClientService{},
	}

	v2Client := unittesting.BuildV2ClientWithMocks(desiredMockClients)

	cl := NewClient(v2Client, &unittesting.DefaultOpts, authInfo)

	return cl, desiredMockClients
}

func TestRESTClient_NewGarbageCollection(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	mockClient.GC.On("CreateGCSchedule", postGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&gc.CreateGCScheduleCreated{}, nil)

	err := apiClient.NewGarbageCollection(ctx, schedule)

	assert.NoError(t, err)

	mockClient.GC.AssertExpectations(t)
}

func TestRESTClient_NewGarbageCollection_ErrSystemInvalidSchedule(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	mockClient.GC.On("CreateGCSchedule", postGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &runtime.APIError{Code: http.StatusBadRequest})

	err := apiClient.NewGarbageCollection(ctx, schedule)

	if assert.Error(t, err) {
		assert.IsType(t, &common.ErrSystemInvalidSchedule{}, err)
	}

	mockClient.GC.AssertExpectations(t)
}

func TestRESTClient_NewGarbageCollection_StatusUnauthorized(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	mockClient.GC.On("CreateGCSchedule", postGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &runtime.APIError{Code: http.StatusUnauthorized})

	err := apiClient.NewGarbageCollection(ctx, schedule)

	if assert.Error(t, err) {
		assert.IsType(t, &common.ErrSystemUnauthorized{}, err)
	}

	mockClient.GC.AssertExpectations(t)
}

func TestRESTClient_NewGarbageCollection_ErrSystemGcInProgress(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	mockClient.GC.On("CreateGCSchedule", postGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &runtime.APIError{Code: http.StatusConflict})

	err := apiClient.NewGarbageCollection(ctx, schedule)

	if assert.Error(t, err) {
		assert.IsType(t, &common.ErrSystemGcInProgress{}, err)
	}

	mockClient.GC.AssertExpectations(t)
}

func TestRESTClient_NewGarbageCollection_ErrSystemInternalErrors(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	mockClient.GC.On("CreateGCSchedule", postGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &runtime.APIError{Code: http.StatusInternalServerError})

	err := apiClient.NewGarbageCollection(ctx, schedule)

	if assert.Error(t, err) {
		assert.IsType(t, &common.ErrSystemInternalErrors{}, err)
	}

	mockClient.GC.AssertExpectations(t)
}

func TestRESTClient_NewGarbageCollection_ErrSystemGcInProgress_2(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	mockClient.GC.On("CreateGCSchedule", postGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &runtime.APIError{Code: http.StatusConflict})

	err := apiClient.NewGarbageCollection(ctx, schedule)

	if assert.Error(t, err) {
		assert.IsType(t, &common.ErrSystemGcInProgress{}, err)
	}

	mockClient.GC.AssertExpectations(t)
}

func TestRESTClient_NewGarbageCollection_ErrSystemInvalidSchedule_2(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	mockClient.GC.On("CreateGCSchedule", postGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(nil, &runtime.APIError{Code: http.StatusBadRequest})

	err := apiClient.NewGarbageCollection(ctx, schedule)

	if assert.Error(t, err) {
		assert.IsType(t, &common.ErrSystemInvalidSchedule{}, err)
	}

	mockClient.GC.AssertExpectations(t)
}

func TestRESTClient_GetGarbageCollection(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	heldGC := &modelv2.GCHistory{
		Schedule: &modelv2.ScheduleObj{
			Cron: schedule.Schedule.Cron,
			Type: schedule.Schedule.Type,
		},
	}

	mockClient.GC.On("GetGCSchedule", getGcScheduleParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&gc.GetGCScheduleOK{Payload: heldGC}, nil)

	gc, err := apiClient.GetGarbageCollectionSchedule(ctx)

	assert.NoError(t, err)

	assert.IsType(t, &modelv2.GCHistory{}, gc)
	assert.Equal(t, gc.Schedule, schedule.Schedule)

	mockClient.GC.AssertExpectations(t)
}

func TestRESTClient_GetGarbageCollectionSchedule_ScheduleNil(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	nilSchedulegcReq := &modelv2.GCHistory{
		Schedule: nil,
	}

	mockClient.GC.On("GetGCSchedule", getGcScheduleParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&gc.GetGCScheduleOK{Payload: nilSchedulegcReq}, nil)

	_, err := apiClient.GetGarbageCollectionSchedule(ctx)

	if assert.Error(t, err) {
		assert.IsType(t, &common.ErrSystemGcScheduleUndefined{}, err)
	}

	mockClient.GC.AssertExpectations(t)
}

func TestRESTClient_UpdateGarbageCollection(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

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

	mockClient.GC.On("UpdateGCSchedule", &putGCParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&gc.UpdateGCScheduleOK{}, nil)

	err := apiClient.UpdateGarbageCollection(ctx, newGcReq)

	assert.NoError(t, err)

	mockClient.GC.AssertExpectations(t)
}

func TestRESTClient_UpdateGarbageCollection_ScheduleNil(t *testing.T) {
	apiClient, _ := APIandMockClientsForTests()

	err := apiClient.UpdateGarbageCollection(ctx, nil)

	if assert.Error(t, err) {
		assert.IsType(t, &common.ErrSystemGcScheduleNotProvided{}, err)
	}
}

func TestRESTClient_ResetGarbageCollection(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

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

	mockClient.GC.On("UpdateGCSchedule", putGcParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&gc.UpdateGCScheduleOK{}, nil)

	err := apiClient.ResetGarbageCollection(ctx)

	assert.NoError(t, err)

	mockClient.GC.AssertExpectations(t)
}
