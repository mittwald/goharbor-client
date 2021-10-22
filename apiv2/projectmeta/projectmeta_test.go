//go:build !integration

package projectmeta

import (
	"context"
	"strconv"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	projectmeta "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/project_metadata"
	"github.com/mittwald/goharbor-client/v4/apiv2/mocks"
	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/common"
	unittesting "github.com/mittwald/goharbor-client/v4/apiv2/pkg/testing"
)

var (
	authInfo         = runtimeclient.BasicAuth("foo", "bar")
	exampleProjectID = 1
	ctx              = context.Background()
)

func APIandMockClientsForTests() (*RESTClient, *unittesting.MockClients) {
	desiredMockClients := &unittesting.MockClients{
		ProjectMetadata: mocks.MockProject_metadataClientService{},
	}

	v2Client := unittesting.BuildV2ClientWithMocks(desiredMockClients)

	cl := NewClient(v2Client, &unittesting.DefaultOpts, authInfo)

	return cl, desiredMockClients
}

func TestRESTClient_AddProjectMetadata(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	addParams := &projectmeta.AddProjectMetadatasParams{
		Metadata: map[string]string{
			common.MetadataKeyEnableContentTrust.String(): "false",
		},
		ProjectNameOrID: strconv.Itoa(exampleProjectID),
		Context:         ctx,
	}

	mockClient.ProjectMetadata.On("AddProjectMetadatas", addParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectmeta.AddProjectMetadatasOK{}, nil)

	err := apiClient.AddProjectMetadata(ctx, strconv.Itoa(exampleProjectID), common.MetadataKeyEnableContentTrust, "false")

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

	var k common.MetadataKey
	t.Run("ProjectMetadataValueEnableContentTrustUndefined", func(t *testing.T) {
		k = common.MetadataKeyEnableContentTrust
		getParams.MetaName = common.MetadataKeyEnableContentTrust.String()

		mockClient.ProjectMetadata.On("GetProjectMetadata", getParams,
			mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&projectmeta.GetProjectMetadataOK{}, nil)
		val, err := apiClient.GetProjectMetadataValue(ctx, strconv.Itoa(exampleProjectID), k)
		require.Error(t, err)
		require.Equal(t, val, "")
		require.IsType(t, &common.ErrProjectMetadataValueEnableContentTrustUndefined{}, err)

		mockClient.ProjectMetadata.AssertExpectations(t)
	})
	t.Run("ProjectMetadataValueAutoScanUndefined", func(t *testing.T) {
		k = common.MetadataKeyAutoScan
		getParams.MetaName = common.MetadataKeyAutoScan.String()

		mockClient.ProjectMetadata.On("GetProjectMetadata", getParams,
			mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&projectmeta.GetProjectMetadataOK{}, nil)
		val, err := apiClient.GetProjectMetadataValue(ctx, strconv.Itoa(exampleProjectID), k)
		require.Error(t, err)
		require.Equal(t, val, "")
		require.IsType(t, &common.ErrProjectMetadataValueAutoScanUndefined{}, err)

		mockClient.ProjectMetadata.AssertExpectations(t)
	})
	t.Run("ProjectMetadataValueSeverityUndefined", func(t *testing.T) {
		k = common.MetadataKeySeverity
		getParams.MetaName = common.MetadataKeySeverity.String()

		mockClient.ProjectMetadata.On("GetProjectMetadata", getParams,
			mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&projectmeta.GetProjectMetadataOK{}, nil)
		val, err := apiClient.GetProjectMetadataValue(ctx, strconv.Itoa(exampleProjectID), k)
		require.Error(t, err)
		require.Equal(t, val, "")
		require.IsType(t, &common.ErrProjectMetadataValueSeverityUndefined{}, err)

		mockClient.ProjectMetadata.AssertExpectations(t)
	})
	t.Run("ProjectMetadataValueReuseSysCveAllowlistUndefined", func(t *testing.T) {
		k = common.MetadataKeyReuseSysCVEAllowlist
		getParams.MetaName = common.MetadataKeyReuseSysCVEAllowlist.String()

		mockClient.ProjectMetadata.On("GetProjectMetadata", getParams,
			mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&projectmeta.GetProjectMetadataOK{}, nil)
		val, err := apiClient.GetProjectMetadataValue(ctx, strconv.Itoa(exampleProjectID), k)
		require.Error(t, err)
		require.Equal(t, val, "")
		require.IsType(t, &common.ErrProjectMetadataValueReuseSysCveAllowlistUndefined{}, err)

		mockClient.ProjectMetadata.AssertExpectations(t)
	})
	t.Run("ProjectMetadataValuePublicUndefined", func(t *testing.T) {
		k = common.MetadataKeyPublic
		getParams.MetaName = common.MetadataKeyPublic.String()

		mockClient.ProjectMetadata.On("GetProjectMetadata", getParams,
			mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&projectmeta.GetProjectMetadataOK{}, nil)
		val, err := apiClient.GetProjectMetadataValue(ctx, strconv.Itoa(exampleProjectID), k)
		require.Error(t, err)
		require.Equal(t, val, "")
		require.IsType(t, &common.ErrProjectMetadataValuePublicUndefined{}, err)

		mockClient.ProjectMetadata.AssertExpectations(t)
	})
	t.Run("ProjectMetadataValuePreventVulUndefined", func(t *testing.T) {
		k = common.MetadataKeyPreventVul
		getParams.MetaName = common.MetadataKeyPreventVul.String()

		mockClient.ProjectMetadata.On("GetProjectMetadata", getParams,
			mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
			Return(&projectmeta.GetProjectMetadataOK{}, nil)
		val, err := apiClient.GetProjectMetadataValue(ctx, strconv.Itoa(exampleProjectID), k)
		require.Error(t, err)
		require.Equal(t, val, "")
		require.IsType(t, &common.ErrProjectMetadataValuePreventVulUndefined{}, err)

		mockClient.ProjectMetadata.AssertExpectations(t)
	})
}

func TestRESTClient_GetProjectMetadataValue_MetadataKeyUndefined(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	_, err := apiClient.GetProjectMetadataValue(ctx, strconv.Itoa(exampleProjectID), "")

	require.Error(t, err)
	require.IsType(t, &common.ErrProjectMetadataKeyUndefined{}, err)

	mockClient.ProjectMetadata.AssertExpectations(t)
}

func TestRESTClient_ListProjectMetadata(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	listParams := &projectmeta.ListProjectMetadatasParams{
		ProjectNameOrID: strconv.Itoa(exampleProjectID),
		Context:         ctx,
	}

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

	mockClient.ProjectMetadata.On("ListProjectMetadatas", listParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectmeta.ListProjectMetadatasOK{}, nil)

	_, err := apiClient.ListProjectMetadata(ctx, strconv.Itoa(exampleProjectID))

	require.Error(t, err)
	require.ErrorIs(t, err, &common.ErrProjectMetadataUndefined{})

	mockClient.ProjectMetadata.AssertExpectations(t)
}

func TestRESTClient_UpdateProjectMetadata(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	updateParams := &projectmeta.UpdateProjectMetadataParams{
		MetaName: common.MetadataKeyAutoScan.String(),
		Metadata: map[string]string{
			common.MetadataKeyAutoScan.String(): "true",
		},
		ProjectNameOrID: strconv.Itoa(exampleProjectID),
		Context:         ctx,
	}

	mockClient.ProjectMetadata.On("UpdateProjectMetadata", updateParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectmeta.UpdateProjectMetadataOK{}, nil)

	err := apiClient.UpdateProjectMetadata(ctx, strconv.Itoa(exampleProjectID), common.MetadataKeyAutoScan, "true")

	require.NoError(t, err)

	mockClient.ProjectMetadata.AssertExpectations(t)
}

func TestRESTClient_DeleteProjectMetadataValue(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	deleteParams := &projectmeta.DeleteProjectMetadataParams{
		MetaName:        common.MetadataKeyAutoScan.String(),
		ProjectNameOrID: strconv.Itoa(exampleProjectID),
		Context:         ctx,
	}

	mockClient.ProjectMetadata.On("DeleteProjectMetadata", deleteParams,
		mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&projectmeta.DeleteProjectMetadataOK{}, nil)

	err := apiClient.DeleteProjectMetadataValue(ctx, strconv.Itoa(exampleProjectID), common.MetadataKeyAutoScan)
	require.NoError(t, err)

	mockClient.ProjectMetadata.AssertExpectations(t)
}
