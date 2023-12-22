//go:build !integration

package scanall

import (
	"context"
	"testing"

	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/scan_all"
	"github.com/mittwald/goharbor-client/v5/apiv2/mocks"
	"github.com/mittwald/goharbor-client/v5/apiv2/model"
	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var ctx = context.Background()

func APIandMockClientsForTests() (*RESTClient, *clienttesting.MockClients) {
	desiredMockClients := &clienttesting.MockClients{
		ScanAll:         mocks.MockScan_allClientService{},
		ProjectMetadata: mocks.MockProject_metadataClientService{},
	}

	v2Client := clienttesting.BuildV2ClientWithMocks(desiredMockClients)

	cl := NewClient(v2Client, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	return cl, desiredMockClients
}

func TestRESTClient_CreateScanAllSchedule(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	params := &scan_all.CreateScanAllScheduleParams{
		Context: ctx,
		Schedule: &model.Schedule{
			Schedule: &model.ScheduleObj{
				Type: "Daily",
			},
		},
	}

	params.WithTimeout(apiClient.Options.Timeout)

	mockClient.ScanAll.On("CreateScanAllSchedule", params, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&scan_all.CreateScanAllScheduleCreated{}, nil)

	schedule := &model.Schedule{
		Schedule: &model.ScheduleObj{
			Type: "Daily",
		},
	}

	err := apiClient.CreateScanAllSchedule(ctx, schedule)
	require.NoError(t, err)

	mockClient.Retention.AssertExpectations(t)
}

func TestRESTClient_GetScanAllSchedule(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	params := &scan_all.GetScanAllScheduleParams{
		Context: ctx,
	}

	params.WithTimeout(apiClient.Options.Timeout)

	mockClient.ScanAll.On("GetScanAllSchedule", params, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&scan_all.GetScanAllScheduleOK{
			Payload: &model.Schedule{
				Schedule: &model.ScheduleObj{
					Type: "Daily",
				},
			},
		}, nil)
	_, err := apiClient.GetScanAllSchedule(ctx)
	require.NoError(t, err)

	mockClient.Retention.AssertExpectations(t)
}

func TestRESTClient_UpdateScanAllSchedule(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	params := &scan_all.UpdateScanAllScheduleParams{
		Context: ctx,
		Schedule: &model.Schedule{
			Schedule: &model.ScheduleObj{
				Type: "Daily",
			},
		},
	}

	params.WithTimeout(apiClient.Options.Timeout)

	mockClient.ScanAll.On("UpdateScanAllSchedule", params, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&scan_all.UpdateScanAllScheduleOK{}, nil)

	schedule := &model.Schedule{
		Schedule: &model.ScheduleObj{
			Type: "Daily",
		},
	}

	err := apiClient.UpdateScanAllSchedule(ctx, schedule)
	require.NoError(t, err)

	mockClient.Retention.AssertExpectations(t)
}
