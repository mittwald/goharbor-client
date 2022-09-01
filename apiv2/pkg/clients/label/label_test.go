//go:build !integration

package label

import (
	"context"
	"testing"

	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/label"
	"github.com/mittwald/goharbor-client/v5/apiv2/mocks"
	"github.com/mittwald/goharbor-client/v5/apiv2/model"
	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/util"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	ctx       = context.Background()
	testLabel = model.Label{
		Color:       "#000000",
		Description: "test",
		Name:        "test-label",
	}
)

func APIandMockClientsForTests() (*RESTClient, *clienttesting.MockClients) {
	desiredMockClients := &clienttesting.MockClients{
		Label: mocks.MockLabelClientService{},
	}

	v2Client := clienttesting.BuildV2ClientWithMocks(desiredMockClients)

	cl := NewClient(v2Client, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	return cl, desiredMockClients
}

func TestRESTClient_CreateLabel(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	createParams := &label.CreateLabelParams{
		Label:   &testLabel,
		Context: ctx,
	}

	createParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Label.On("CreateLabel", createParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&label.CreateLabelCreated{}, nil)

	err := apiClient.CreateLabel(ctx, &testLabel)
	require.NoError(t, err)

	mockClient.Label.AssertExpectations(t)
}

func TestRESTClient_DeleteLabel(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	deleteParams := &label.DeleteLabelParams{
		LabelID: 1,
		Context: ctx,
	}

	deleteParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Label.On("DeleteLabel", deleteParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&label.DeleteLabelOK{}, nil)

	err := apiClient.DeleteLabel(ctx, 1)
	require.NoError(t, err)

	mockClient.Label.AssertExpectations(t)
}

func TestRESTClient_GetLabelByID(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := &label.GetLabelByIDParams{
		LabelID: 1,
		Context: ctx,
	}

	getParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Label.On("GetLabelByID", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&label.GetLabelByIDOK{
			Payload: &testLabel,
		}, nil)

	l, err := apiClient.GetLabelByID(ctx, 1)

	require.NoError(t, err)
	require.Equal(t, testLabel, *l)

	mockClient.Label.AssertExpectations(t)
}

func TestRESTClient_ListLabels_ScopeGlobal(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	listParams := &label.ListLabelsParams{
		Name:      util.StringPtr("test"),
		Page:      &apiClient.Options.Page,
		PageSize:  &apiClient.Options.PageSize,
		ProjectID: nil,
		Q:         &apiClient.Options.Query,
		Scope:     util.StringPtr(ScopeGlobal.String()),
		Sort:      &apiClient.Options.Sort,
		Context:   ctx,
	}

	listParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Label.On("ListLabels", listParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&label.ListLabelsOK{
			Payload: []*model.Label{&testLabel},
		}, nil)


	t.Run("ErrProjectIDProvided", func(t *testing.T) {
		_, err := apiClient.ListLabels(ctx, "test", util.Int64Ptr(1), ScopeGlobal)
		require.Error(t, err)
	})

	labels, err := apiClient.ListLabels(ctx, "test", nil, ScopeGlobal)

	require.NoError(t, err)
	require.Equal(t, 1, len(labels))

	mockClient.Label.AssertExpectations(t)
}

func TestRESTClient_ListLabels_ScopeProject(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	listParams := &label.ListLabelsParams{
		Name:      util.StringPtr("test"),
		Page:      &apiClient.Options.Page,
		PageSize:  &apiClient.Options.PageSize,
		ProjectID: util.Int64Ptr(1),
		Q:         &apiClient.Options.Query,
		Scope:     util.StringPtr(ScopeProject.String()),
		Sort:      &apiClient.Options.Sort,
		Context:   ctx,
	}

	listParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Label.On("ListLabels", listParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&label.ListLabelsOK{
			Payload: []*model.Label{&testLabel},
		}, nil)

	t.Run("ErrNoProjectIDProvided", func(t *testing.T) {
		_, err := apiClient.ListLabels(ctx, "test", nil, ScopeProject)
		require.Error(t, err)
	})

	labels, err := apiClient.ListLabels(ctx, "test", util.Int64Ptr(1), ScopeProject)

	require.NoError(t, err)
	require.Equal(t, 1, len(labels))

	mockClient.Label.AssertExpectations(t)
}

func TestRESTClient_UpdateLabel(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	updateParams := &label.UpdateLabelParams{
		Label:   &testLabel,
		LabelID: 1,
		Context: ctx,
	}

	updateParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Label.On("UpdateLabel", updateParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&label.UpdateLabelOK{}, nil)

	t.Run("ErrNoLabelNameProvided", func(t *testing.T) {
		invalidLabel := testLabel
		invalidLabel.Name = ""
		err := apiClient.UpdateLabel(ctx, 1, &invalidLabel)
		require.Error(t, err)
	})

	err := apiClient.UpdateLabel(ctx, 1, &testLabel)
	require.NoError(t, err)

	mockClient.Label.AssertExpectations(t)
}
