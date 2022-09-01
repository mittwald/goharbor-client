//go:build !integration

package projectmeta

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	projectmeta "github.com/testwill/goharbor-client/v5/apiv2/internal/api/client/project_metadata"
	"github.com/testwill/goharbor-client/v5/apiv2/mocks"
	"github.com/testwill/goharbor-client/v5/apiv2/pkg/common"
	"github.com/testwill/goharbor-client/v5/apiv2/pkg/errors"
	clienttesting "github.com/testwill/goharbor-client/v5/apiv2/pkg/testing"
)

var (
	exampleProjectID = 1
	ctx              = context.Background()
)

func APIandMockClientsForTests() (*RESTClient, *clienttesting.MockClients) {
	desiredMockClients := &clienttesting.MockClients{
		ProjectMetadata: mocks.MockProject_metadataClientService{},
	}

	v2Client := clienttesting.BuildV2ClientWithMocks(desiredMockClients)

	cl := NewClient(v2Client, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	return cl, desiredMockClients
}

func TestRESTClient_AddProjectMetadata(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	addParams := &projectmeta.AddProjectMetadatasParams{
		Metadata: map[string]string{
			common.ProjectMetadataKeyEnableContentTrust.String(): "false",
		},
		ProjectNameOrID: strconv.Itoa(exampleProjectID),
		Context:         ctx,
	}

	addParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.ProjectMetadata.On("AddProjectMetadatas", addParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectmeta.AddProjectMetadatasOK{}, nil)

	err := apiClient.AddProjectMetadata(ctx, strconv.Itoa(exampleProjectID), common.ProjectMetadataKeyEnableContentTrust, "false")

	require.NoError(t, err)

	mockClient.ProjectMetadata.AssertExpectations(t)
}

func TestRESTClient_GetProjectMetadataValue_ValuesUndefined(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := &projectmeta.GetProjectMetadataParams{
		MetaName:        "",
		ProjectNameOrID: strconv.Itoa(exampleProjectID),
		Context:         ctx,
	}

	getParams.WithTimeout(apiClient.Options.Timeout)

	var k common.MetadataKey
	t.Run("ProjectMetadataValueEnableContentTrustUndefined", func(t *testing.T) {
		k = common.ProjectMetadataKeyEnableContentTrust
		getParams.MetaName = common.ProjectMetadataKeyEnableContentTrust.String()

		mockClient.ProjectMetadata.On("GetProjectMetadata", getParams,
			mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&projectmeta.GetProjectMetadataOK{}, nil)
		val, err := apiClient.GetProjectMetadataValue(ctx, strconv.Itoa(exampleProjectID), k)
		require.Error(t, err)
		require.Equal(t, val, "")
		require.IsType(t, &errors.ErrProjectMetadataValueEnableContentTrustUndefined{}, err)

		mockClient.ProjectMetadata.AssertExpectations(t)
	})
	t.Run("ProjectMetadataValueAutoScanUndefined", func(t *testing.T) {
		k = common.ProjectMetadataKeyAutoScan
		getParams.MetaName = common.ProjectMetadataKeyAutoScan.String()

		mockClient.ProjectMetadata.On("GetProjectMetadata", getParams,
			mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&projectmeta.GetProjectMetadataOK{}, nil)
		val, err := apiClient.GetProjectMetadataValue(ctx, strconv.Itoa(exampleProjectID), k)
		require.Error(t, err)
		require.Equal(t, val, "")
		require.IsType(t, &errors.ErrProjectMetadataValueAutoScanUndefined{}, err)

		mockClient.ProjectMetadata.AssertExpectations(t)
	})
	t.Run("ProjectMetadataValueSeverityUndefined", func(t *testing.T) {
		k = common.ProjectMetadataKeySeverity
		getParams.MetaName = common.ProjectMetadataKeySeverity.String()

		mockClient.ProjectMetadata.On("GetProjectMetadata", getParams,
			mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&projectmeta.GetProjectMetadataOK{}, nil)
		val, err := apiClient.GetProjectMetadataValue(ctx, strconv.Itoa(exampleProjectID), k)
		require.Error(t, err)
		require.Equal(t, val, "")
		require.IsType(t, &errors.ErrProjectMetadataValueSeverityUndefined{}, err)

		mockClient.ProjectMetadata.AssertExpectations(t)
	})
	t.Run("ProjectMetadataValueReuseSysCveAllowlistUndefined", func(t *testing.T) {
		k = common.ProjectMetadataKeyReuseSysCVEAllowlist
		getParams.MetaName = common.ProjectMetadataKeyReuseSysCVEAllowlist.String()

		mockClient.ProjectMetadata.On("GetProjectMetadata", getParams,
			mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&projectmeta.GetProjectMetadataOK{}, nil)
		val, err := apiClient.GetProjectMetadataValue(ctx, strconv.Itoa(exampleProjectID), k)
		require.Error(t, err)
		require.Equal(t, val, "")
		require.IsType(t, &errors.ErrProjectMetadataValueReuseSysCveAllowlistUndefined{}, err)

		mockClient.ProjectMetadata.AssertExpectations(t)
	})
	t.Run("ProjectMetadataValuePublicUndefined", func(t *testing.T) {
		k = common.ProjectMetadataKeyPublic
		getParams.MetaName = common.ProjectMetadataKeyPublic.String()

		mockClient.ProjectMetadata.On("GetProjectMetadata", getParams,
			mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&projectmeta.GetProjectMetadataOK{}, nil)
		val, err := apiClient.GetProjectMetadataValue(ctx, strconv.Itoa(exampleProjectID), k)
		require.Error(t, err)
		require.Equal(t, val, "")
		require.IsType(t, &errors.ErrProjectMetadataValuePublicUndefined{}, err)

		mockClient.ProjectMetadata.AssertExpectations(t)
	})
	t.Run("ProjectMetadataValuePreventVulUndefined", func(t *testing.T) {
		k = common.ProjectMetadataKeyPreventVul
		getParams.MetaName = common.ProjectMetadataKeyPreventVul.String()

		mockClient.ProjectMetadata.On("GetProjectMetadata", getParams,
			mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&projectmeta.GetProjectMetadataOK{}, nil)
		val, err := apiClient.GetProjectMetadataValue(ctx, strconv.Itoa(exampleProjectID), k)
		require.Error(t, err)
		require.Equal(t, val, "")
		require.IsType(t, &errors.ErrProjectMetadataValuePreventVulUndefined{}, err)

		mockClient.ProjectMetadata.AssertExpectations(t)
	})
}

func TestRESTClient_GetProjectMetadataValue_MetadataKeyUndefined(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	_, err := apiClient.GetProjectMetadataValue(ctx, strconv.Itoa(exampleProjectID), "")

	require.Error(t, err)
	require.IsType(t, &errors.ErrProjectMetadataKeyUndefined{}, err)

	mockClient.ProjectMetadata.AssertExpectations(t)
}

func TestRESTClient_ListProjectMetadata(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	listParams := &projectmeta.ListProjectMetadatasParams{
		ProjectNameOrID: strconv.Itoa(exampleProjectID),
		Context:         ctx,
	}

	listParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.ProjectMetadata.On("ListProjectMetadatas", listParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectmeta.ListProjectMetadatasOK{Payload: map[string]string{}}, nil)

	_, err := apiClient.ListProjectMetadata(ctx, strconv.Itoa(exampleProjectID))

	require.NoError(t, err)

	mockClient.ProjectMetadata.AssertExpectations(t)
}

func TestRESTClient_ListProjectMetadata_Undefined(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	listParams := &projectmeta.ListProjectMetadatasParams{
		ProjectNameOrID: strconv.Itoa(exampleProjectID),
		Context:         ctx,
	}

	listParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.ProjectMetadata.On("ListProjectMetadatas", listParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectmeta.ListProjectMetadatasOK{}, nil)

	_, err := apiClient.ListProjectMetadata(ctx, strconv.Itoa(exampleProjectID))

	require.Error(t, err)
	require.ErrorIs(t, err, &errors.ErrProjectMetadataUndefined{})

	mockClient.ProjectMetadata.AssertExpectations(t)
}

func TestRESTClient_UpdateProjectMetadata(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	updateParams := &projectmeta.UpdateProjectMetadataParams{
		MetaName: common.ProjectMetadataKeyAutoScan.String(),
		Metadata: map[string]string{
			common.ProjectMetadataKeyAutoScan.String(): "true",
		},
		ProjectNameOrID: strconv.Itoa(exampleProjectID),
		Context:         ctx,
	}

	updateParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.ProjectMetadata.On("UpdateProjectMetadata", updateParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectmeta.UpdateProjectMetadataOK{}, nil)

	err := apiClient.UpdateProjectMetadata(ctx, strconv.Itoa(exampleProjectID), common.ProjectMetadataKeyAutoScan, "true")

	require.NoError(t, err)

	mockClient.ProjectMetadata.AssertExpectations(t)
}

func TestRESTClient_DeleteProjectMetadataValue(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	deleteParams := &projectmeta.DeleteProjectMetadataParams{
		MetaName:        common.ProjectMetadataKeyAutoScan.String(),
		ProjectNameOrID: strconv.Itoa(exampleProjectID),
		Context:         ctx,
	}

	deleteParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.ProjectMetadata.On("DeleteProjectMetadata", deleteParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectmeta.DeleteProjectMetadataOK{}, nil)

	err := apiClient.DeleteProjectMetadataValue(ctx, strconv.Itoa(exampleProjectID), common.ProjectMetadataKeyAutoScan)
	require.NoError(t, err)

	mockClient.ProjectMetadata.AssertExpectations(t)
}
