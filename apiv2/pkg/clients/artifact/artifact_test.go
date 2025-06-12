//go:build !integration

package artifact

import (
	"context"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/util"
	"testing"

	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/artifact"
	"github.com/mittwald/goharbor-client/v5/apiv2/mocks"
	"github.com/mittwald/goharbor-client/v5/apiv2/model"
	clienttesting "github.com/mittwald/goharbor-client/v5/apiv2/pkg/testing"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	ctx            = context.Background()
	projectName    = "test-project"
	repositoryName = "test-repository"
	reference      = "test-artifact"
	label          = model.Label{
		Color:       "#000000",
		Description: "test",
		Name:        "test-label",
	}
)

func APIandMockClientsForTests() (*RESTClient, *clienttesting.MockClients) {
	desiredMockClients := &clienttesting.MockClients{
		Artifact: mocks.MockArtifactClientService{},
	}

	v2Client := clienttesting.BuildV2ClientWithMocks(desiredMockClients)

	cl := NewClient(v2Client, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	return cl, desiredMockClients
}

func TestRESTClient_AddArtifactLabel(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	addParams := &artifact.AddLabelParams{
		Label:          &label,
		ProjectName:    projectName,
		Reference:      reference,
		RepositoryName: repositoryName,
		Context:        ctx,
	}

	addParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Artifact.On("AddLabel", addParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&artifact.AddLabelOK{}, nil)

	err := apiClient.AddArtifactLabel(ctx, projectName, repositoryName, reference, &label)

	require.NoError(t, err)

	mockClient.Artifact.AssertExpectations(t)
}

func TestRESTClient_CopyArtifact(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	from := CopyReference{
		ProjectName:    "some-other-project",
		RepositoryName: "some-repo",
		Tag:            "v1.0.0",
		Digest:         "sha256:1234567890",
	}

	copyParams := &artifact.CopyArtifactParams{
		From:           "some-other-project/some-repo:v1.0.0",
		ProjectName:    projectName,
		RepositoryName: repositoryName,
		Context:        ctx,
	}

	copyParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Artifact.On("CopyArtifact", copyParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&artifact.CopyArtifactCreated{}, nil)

	err := apiClient.CopyArtifact(ctx, &from, projectName, repositoryName)

	require.NoError(t, err)

	mockClient.Artifact.AssertExpectations(t)
}

func TestRESTClient_CopyArtifact_WithDigest(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	from := CopyReference{
		ProjectName:    "some-other-project",
		RepositoryName: "some-repo",
	}

	from.Digest = "sha256:1234567890"

	copyParams := &artifact.CopyArtifactParams{
		From:           "some-other-project/some-repo@sha256:1234567890",
		ProjectName:    projectName,
		RepositoryName: repositoryName,
		Context:        ctx,
	}

	copyParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Artifact.On("CopyArtifact", copyParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&artifact.CopyArtifactCreated{}, nil)

	err := apiClient.CopyArtifact(ctx, &from, projectName, repositoryName)
	require.NoError(t, err)

	mockClient.Artifact.AssertExpectations(t)
}

func TestRESTClient_CopyArtifact_WithoutTagOrDigest(t *testing.T) {
	apiClient, _ := APIandMockClientsForTests()

	from := CopyReference{
		ProjectName:    "some-other-project",
		RepositoryName: "some-repo",
	}

	from.Tag = ""
	from.Digest = ""
	err := apiClient.CopyArtifact(ctx, &from, projectName, repositoryName)
	require.Error(t, err)
	require.Contains(t, err.Error(), "no tag or digest specified")
}

func TestRESTClient_CopyArtifact_WithTag(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	from := CopyReference{
		ProjectName:    "some-other-project",
		RepositoryName: "some-repo",
	}

	from.Tag = "v1.0.0"

	copyParams := &artifact.CopyArtifactParams{
		From:           "some-other-project/some-repo:v1.0.0",
		ProjectName:    projectName,
		RepositoryName: repositoryName,
		Context:        ctx,
	}

	copyParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Artifact.On("CopyArtifact", copyParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&artifact.CopyArtifactCreated{}, nil)

	err := apiClient.CopyArtifact(ctx, &from, projectName, repositoryName)
	require.NoError(t, err)

	mockClient.Artifact.AssertExpectations(t)
}

func TestRESTClient_CreateTag(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	tag := &model.Tag{
		Name: "v1.0.0",
	}

	createParams := &artifact.CreateTagParams{
		ProjectName:    projectName,
		Reference:      reference,
		RepositoryName: repositoryName,
		Tag:            tag,
		Context:        ctx,
	}

	createParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Artifact.On("CreateTag", createParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&artifact.CreateTagCreated{}, nil)

	err := apiClient.CreateTag(ctx, projectName, repositoryName, reference, tag)

	require.NoError(t, err)

	mockClient.Artifact.AssertExpectations(t)
}

func TestRESTClient_DeleteTag(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	deleteParams := &artifact.DeleteTagParams{
		ProjectName:    projectName,
		Reference:      reference,
		RepositoryName: repositoryName,
		TagName:        "v1.0.0",
		Context:        ctx,
	}

	deleteParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Artifact.On("DeleteTag", deleteParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&artifact.DeleteTagOK{}, nil)

	err := apiClient.DeleteTag(ctx, projectName, repositoryName, reference, "v1.0.0")

	require.NoError(t, err)

	mockClient.Artifact.AssertExpectations(t)
}

func TestRESTClient_GetArtifact(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	getParams := artifact.NewGetArtifactParams()

	getParams.WithTimeout(apiClient.Options.Timeout)
	getParams.WithPage(&apiClient.Options.Page)
	getParams.WithPageSize(&apiClient.Options.PageSize)
	getParams.WithProjectName(projectName)
	getParams.WithRepositoryName(repositoryName)
	getParams.WithReference(reference)
	getParams.WithContext(ctx)
	getParams.WithWithLabel(util.BoolPtr(true))
	getParams.WithWithAccessory(util.BoolPtr(true))

	getParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Artifact.On("GetArtifact", getParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&artifact.GetArtifactOK{Payload: &model.Artifact{}}, nil)

	resp, err := apiClient.GetArtifact(ctx, projectName, repositoryName, reference)
	require.NoError(t, err)
	require.NotNil(t, resp)

	mockClient.Artifact.AssertExpectations(t)
}

func TestRESTClient_ListArtifacts(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	listParams := artifact.NewListArtifactsParams()
	listParams.WithProjectName(projectName)
	listParams.WithRepositoryName(repositoryName)
	listParams.WithContext(ctx)
	listParams.WithPage(&apiClient.Options.Page)
	listParams.WithPageSize(&apiClient.Options.PageSize)
	listParams.WithSort(&apiClient.Options.Sort)
	listParams.WithQ(&apiClient.Options.Query)
	listParams.WithTimeout(apiClient.Options.Timeout)
	listParams.WithWithLabel(util.BoolPtr(true))

	mockClient.Artifact.On("ListArtifacts", listParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&artifact.ListArtifactsOK{Payload: []*model.Artifact{}}, nil)

	resp, err := apiClient.ListArtifacts(ctx, projectName, repositoryName)
	require.NoError(t, err)
	require.Len(t, resp, 0)

	mockClient.Artifact.AssertExpectations(t)
}

func TestRESTClient_ListTags(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	listParams := artifact.NewListTagsParams()
	listParams.WithProjectName(projectName)
	listParams.WithRepositoryName(repositoryName)
	listParams.WithReference(reference)
	listParams.WithTimeout(apiClient.Options.Timeout)
	listParams.WithContext(ctx)
	listParams.WithPage(&apiClient.Options.Page)
	listParams.WithPageSize(&apiClient.Options.PageSize)
	listParams.WithSort(&apiClient.Options.Sort)
	listParams.WithQ(&apiClient.Options.Query)

	mockClient.Artifact.On("ListTags", listParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&artifact.ListTagsOK{Payload: []*model.Tag{}}, nil)

	resp, err := apiClient.ListTags(ctx, projectName, repositoryName, reference)
	require.NoError(t, err)
	require.Len(t, resp, 0)

	mockClient.Artifact.AssertExpectations(t)
}

func TestRESTClient_RemoveLabel(t *testing.T) {
	apiClient, mockClient := APIandMockClientsForTests()

	removeParams := &artifact.RemoveLabelParams{
		LabelID:        1,
		ProjectName:    projectName,
		Reference:      reference,
		RepositoryName: repositoryName,
		Context:        ctx,
	}

	removeParams.WithTimeout(apiClient.Options.Timeout)

	mockClient.Artifact.On("RemoveLabel", removeParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&artifact.RemoveLabelOK{}, nil)

	err := apiClient.RemoveLabel(ctx, projectName, repositoryName, reference, 1)
	require.NoError(t, err)

	mockClient.Artifact.AssertExpectations(t)
}
