//go:build !integration

package purge

import (
	"context"
	"testing"

	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/purge"
	"github.com/mittwald/goharbor-client/v5/apiv2/mocks"
	"github.com/mittwald/goharbor-client/v5/apiv2/model"
	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	ctx = context.Background()
)

func APIandMockClientsForTests() (*RESTClient, *clienttesting.MockClients) {
	desiredMockClients := &clienttesting.MockClients{
		Purge: mocks.MockPurgeClientService{},
	}

	v2Client := clienttesting.BuildV2ClientWithMocks(desiredMockClients)

	cl := NewClient(v2Client, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	return cl, desiredMockClients
}

func TestRESTClient_CreatePurgeSchedule(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	schedule := &model.Schedule{
		Parameters: map[string]interface{}{
			"audit_retention_hour": 168,
			"dry_run":              true,
			"include_operations":   "create,delete,pull",
		},
		Schedule: &model.ScheduleObj{
			Cron: "0 0 * * * *",
			Type: "Hourly",
		},
	}
	createParams := &purge.CreatePurgeScheduleParams{
		Schedule: schedule,
		Context:  ctx,
	}
	createParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Purge.On("CreatePurgeSchedule", createParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&purge.CreatePurgeScheduleCreated{}, nil)

	err := apiClient.CreatePurgeSchedule(ctx, schedule)
	require.NoError(t, err)

	mockClient.Purge.AssertExpectations(t)
}

func TestRESTClient_ListPurgeHistory(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	listParams := &purge.GetPurgeHistoryParams{
		Page:     &apiClient.Options.Page,
		PageSize: &apiClient.Options.PageSize,
		Q:        &apiClient.Options.Query,
		Sort:     &apiClient.Options.Sort,
		Context:  ctx,
	}
	listParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Purge.On("GetPurgeHistory", listParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&purge.GetPurgeHistoryOK{}, nil)

	_, err := apiClient.ListPurgeHistory(ctx)
	require.NoError(t, err)

	mockClient.Purge.AssertExpectations(t)
}

func TestRESTClient_GetPurgeJob(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := &purge.GetPurgeJobParams{
		PurgeID: 1,
		Context: ctx,
	}
	getParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Purge.On("GetPurgeJob", getParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&purge.GetPurgeJobOK{}, nil)
	_, err := apiClient.GetPurgeJob(ctx, 1)
	require.NoError(t, err)

	mockClient.Purge.AssertExpectations(t)
}

func TestRESTClient_GetPurgeJobLog(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := &purge.GetPurgeJobLogParams{
		PurgeID: 1,
		Context: ctx,
	}
	getParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Purge.On("GetPurgeJobLog", getParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&purge.GetPurgeJobLogOK{}, nil)
	_, err := apiClient.GetPurgeJobLog(ctx, 1)
	require.NoError(t, err)

	mockClient.Purge.AssertExpectations(t)
}

func TestRESTClient_GetPurgeSchedule(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := &purge.GetPurgeScheduleParams{
		Context: ctx,
	}
	getParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Purge.On("GetPurgeSchedule", getParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&purge.GetPurgeScheduleOK{}, nil)
	_, err := apiClient.GetPurgeSchedule(ctx)
	require.NoError(t, err)

	mockClient.Purge.AssertExpectations(t)
}

func TestRESTClient_StopPurge(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	putParams := &purge.StopPurgeParams{
		PurgeID: 1,
		Context: ctx,
	}
	putParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Purge.On("StopPurge", putParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&purge.StopPurgeOK{}, nil)

	err := apiClient.StopPurge(ctx, 1)
	require.NoError(t, err)

	mockClient.Purge.AssertExpectations(t)
}

func TestRESTClient_UpdatePurgeSchedule(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	schedule := &model.Schedule{
		ID: 1,
		Parameters: map[string]interface{}{
			"audit_retention_hour": 168,
			"dry_run":              true,
			"include_operations":   "create,delete,pull",
		},
		Schedule: &model.ScheduleObj{
			Cron: "0 0 * * * *",
			Type: "Hourly",
		},
	}

	putParams := &purge.UpdatePurgeScheduleParams{
		Schedule: schedule,
		Context:  ctx,
	}
	putParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Purge.On("UpdatePurgeSchedule", putParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&purge.UpdatePurgeScheduleOK{}, nil)

	err := apiClient.UpdatePurgeSchedule(ctx, schedule)
	require.NoError(t, err)

	mockClient.Purge.AssertExpectations(t)
}
