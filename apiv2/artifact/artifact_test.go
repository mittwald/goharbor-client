// +build !integration

package artifact

import (
	"context"
	"net/url"
	"strings"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	v2client "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/artifact"
	"github.com/mittwald/goharbor-client/v4/apiv2/mocks"
	modelv2 "github.com/mittwald/goharbor-client/v4/apiv2/model"
)

var (
	authInfo             = runtimeclient.BasicAuth("foo", "bar")
	projectName          = "project"
	repositoryName       = "repository/foo-bar"
	int64One       int64 = 1
	emptyStr             = ""
	falsePtr             = false
)

func BuildV2ClientWithMocks(artifact *mocks.MockArtifactClientService) *v2client.Harbor {
	return &v2client.Harbor{
		Artifact: artifact,
	}
}

func TestRESTClient_GetArtifact(t *testing.T) {
	a := &mocks.MockArtifactClientService{}

	v2Client := BuildV2ClientWithMocks(a)

	cl := NewClient(v2Client, authInfo)

	ctx := context.Background()

	getArtifactParams := &artifact.GetArtifactParams{
		XAcceptVulnerabilities: &emptyStr,
		Page:                   &int64One,
		PageSize:               &int64One,
		ProjectName:            projectName,
		Reference:              "",
		RepositoryName:         url.QueryEscape(repositoryName),
		WithImmutableStatus:    &falsePtr,
		WithLabel:              &falsePtr,
		WithScanOverview:       &falsePtr,
		WithSignature:          &falsePtr,
		WithTag:                &falsePtr,
		Context:                ctx,
	}

	getOpts := ParamOptions{
		PageSize:               1,
		ProjectName:            projectName,
		RepositoryName:         repositoryName,
		Reference:              "",
		WithScanOverview:       false,
		WithLabel:              false,
		WithSignature:          false,
		WithTag:                false,
		WithImmutableStatus:    false,
		XAcceptVulnerabilities: "",
	}

	a.On("GetArtifact", getArtifactParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&artifact.GetArtifactOK{}, nil)

	_, err := cl.GetArtifact(ctx, getOpts)

	require.NoError(t, err)

	a.AssertExpectations(t)
}

func TestRESTClient_ListArtifacts(t *testing.T) {
	a := &mocks.MockArtifactClientService{}

	v2Client := BuildV2ClientWithMocks(a)

	cl := NewClient(v2Client, authInfo)

	ctx := context.Background()

	listArtifactParams := &artifact.ListArtifactsParams{
		XAcceptVulnerabilities: &emptyStr,
		Page:                   &int64One,
		PageSize:               &int64One,
		ProjectName:            projectName,
		RepositoryName:         url.QueryEscape(repositoryName),
		WithImmutableStatus:    &falsePtr,
		WithLabel:              &falsePtr,
		WithScanOverview:       &falsePtr,
		WithSignature:          &falsePtr,
		WithTag:                &falsePtr,
		Context:                ctx,
		Q:                      &emptyStr,
	}

	listOpts := ParamOptions{
		PageSize:               1,
		ProjectName:            projectName,
		RepositoryName:         repositoryName,
		Reference:              "",
		WithScanOverview:       false,
		WithLabel:              false,
		WithSignature:          false,
		WithTag:                false,
		WithImmutableStatus:    false,
		XAcceptVulnerabilities: "",
	}

	a.On("ListArtifacts", listArtifactParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&artifact.ListArtifactsOK{}, nil)

	_, err := cl.ListArtifacts(ctx, listOpts)

	require.NoError(t, err)

	a.AssertExpectations(t)
}

func TestRESTClient_DeleteArtifact(t *testing.T) {
	a := &mocks.MockArtifactClientService{}

	v2Client := BuildV2ClientWithMocks(a)

	cl := NewClient(v2Client, authInfo)

	ctx := context.Background()

	deleteParams := &artifact.DeleteArtifactParams{
		ProjectName:    projectName,
		RepositoryName: url.QueryEscape(repositoryName),
		Context:        ctx,
	}

	deleteOpts := ParamOptions{
		ProjectName:    projectName,
		RepositoryName: repositoryName,
	}

	a.On("DeleteArtifact", deleteParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&artifact.DeleteArtifactOK{}, nil)

	err := cl.DeleteArtifact(ctx, deleteOpts)

	require.NoError(t, err)

	a.AssertExpectations(t)
}

func TestRESTClient_CreateTag(t *testing.T) {
	a := &mocks.MockArtifactClientService{}

	v2Client := BuildV2ClientWithMocks(a)

	cl := NewClient(v2Client, authInfo)

	ctx := context.Background()

	createTagParams := &artifact.CreateTagParams{
		ProjectName:    projectName,
		Reference:      "",
		RepositoryName: url.QueryEscape(repositoryName),
		Tag: &modelv2.Tag{
			ArtifactID:   1,
			Immutable:    true,
			Name:         "v1.0.0",
			RepositoryID: 1,
			Signed:       false,
		},
		Context: ctx,
	}

	createOpts := ParamOptions{
		ProjectName:    projectName,
		RepositoryName: repositoryName,
		Tag: &modelv2.Tag{
			ArtifactID:   1,
			Immutable:    true,
			Name:         "v1.0.0",
			RepositoryID: 1,
			Signed:       false,
		},
	}

	a.On("CreateTag", createTagParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&artifact.CreateTagCreated{}, nil)

	err := cl.CreateTag(ctx, createOpts)

	require.NoError(t, err)

	a.AssertExpectations(t)
}

func TestRESTClient_DeleteTag(t *testing.T) {
	a := &mocks.MockArtifactClientService{}

	v2Client := BuildV2ClientWithMocks(a)

	cl := NewClient(v2Client, authInfo)

	ctx := context.Background()

	deleteTagParams := &artifact.DeleteTagParams{
		ProjectName:    projectName,
		Reference:      "",
		RepositoryName: url.QueryEscape(repositoryName),
		TagName:        "v1.0.0",
		Context:        ctx,
	}

	deleteOpts := ParamOptions{
		ProjectName:    projectName,
		RepositoryName: repositoryName,
		Tag:            &modelv2.Tag{Name: "v1.0.0"},
	}

	a.On("DeleteTag", deleteTagParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&artifact.DeleteTagOK{}, nil)

	err := cl.DeleteTag(ctx, deleteOpts)

	require.NoError(t, err)

	a.AssertExpectations(t)
}

func TestRESTClient_GetVulnerabilitiesAddition(t *testing.T) {
	a := &mocks.MockArtifactClientService{}

	v2Client := BuildV2ClientWithMocks(a)

	cl := NewClient(v2Client, authInfo)

	ctx := context.Background()

	v := string(MIMETypeV10) + "," + string(MIMETypeV11)

	getVulnerabilitiesAdditionParams := &artifact.GetVulnerabilitiesAdditionParams{
		XAcceptVulnerabilities: &v,
		ProjectName:            projectName,
		Reference:              "",
		RepositoryName:         url.QueryEscape(repositoryName),
		Context:                ctx,
	}

	strings.Join([]string{string(MIMETypeV10), string(MIMETypeV11)}, ",")
	m := []MIMEType{
		MIMETypeV10,
		MIMETypeV11,
	}

	a.On("GetVulnerabilitiesAddition", getVulnerabilitiesAdditionParams, mock.AnythingOfType("runtime.ClientAuthInfoWriterFunc")).
		Return(&artifact.GetVulnerabilitiesAdditionOK{}, nil)

	_, err := cl.GetRepositoryVulnerabilities(ctx, projectName, repositoryName, "", m)

	require.NoError(t, err)

	a.AssertExpectations(t)
}
