package artifact

import (
	"context"
	"fmt"

	"github.com/go-openapi/runtime"
	v2client "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/artifact"
	"github.com/mittwald/goharbor-client/v5/apiv2/model"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/config"
)

// RESTClient is a subclient for handling artifact related actions.
type RESTClient struct {
	// Options contains optional configuration when making API calls.
	Options *config.Options

	// The new client of the harbor v2 API
	V2Client *v2client.Harbor

	// AuthInfo contains the auth information that is provided on API calls.
	AuthInfo runtime.ClientAuthInfoWriter
}

func NewClient(v2Client *v2client.Harbor, opts *config.Options, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		Options:  opts,
		V2Client: v2Client,
		AuthInfo: authInfo,
	}
}

type Client interface {
	AddArtifactLabel(ctx context.Context, projectName, repositoryName, reference string, label *model.Label) error
	CopyArtifact(ctx context.Context, from *CopyReference, projectName, repositoryName string) error
	CreateTag(ctx context.Context, projectName, repositoryName, reference string, tag *model.Tag) error
	DeleteTag(ctx context.Context, projectName, repositoryName, reference, tagName string) error
	GetArtifact(ctx context.Context, projectName, repositoryName, reference string) (*model.Artifact, error)
	DeleteArtifact(ctx context.Context, projectName, repositoryName, reference string) error
	ListArtifacts(ctx context.Context, projectName, repositoryName string) ([]*model.Artifact, error)
	ListTags(ctx context.Context, projectName, repositoryName, reference string) ([]*model.Tag, error)
	RemoveLabel(ctx context.Context, projectName, repositoryName, reference string, id int64) error
	// TODO: Introduce this, once https://github.com/goharbor/harbor/issues/13468 is resolved.
	// GetAddition(ctx context.Context, projectName, repositoryName, reference string, addition Addition) (string, error)
	// GetVulnerabilitiesAddition(ctx context.Context, projectName, repositoryName, reference string) (string, error)
}

// ToString returns a string representation of a CopyReference.
// Possible formats are "project/repository:tag" or "project/repository@digest".
// Returns an error if neither tag nor digest is set.
func (in CopyReference) toString() (string, error) {
	var suffix string

	if in.Digest == "" && in.Tag == "" {
		return "", fmt.Errorf("no tag or digest specified")
	}

	if in.Digest != "" {
		suffix = "@" + in.Digest
	}
	if in.Tag != "" {
		suffix = ":" + in.Tag
	}

	return in.ProjectName + "/" + in.RepositoryName + suffix, nil
}

// CopyReference defines the parameters needed for an artifact copy.
type CopyReference struct {
	ProjectName    string
	RepositoryName string
	Tag            string
	Digest         string
}

type Addition string

const (
	AdditionBuildHistory Addition = "build_history"
	AdditionValuesYAML   Addition = "values.yaml"
	AdditionReadme       Addition = "readme.md"
	AdditionDependencies Addition = "dependencies"
)

// TODO: Introduce this, once https://github.com/goharbor/harbor/issues/13468 is resolved.
//func (in Addition) string() string {
//	return string(in)
//}

func (c *RESTClient) AddArtifactLabel(ctx context.Context, projectName, repositoryName, reference string, label *model.Label) error {
	params := &artifact.AddLabelParams{
		Label:          label,
		ProjectName:    projectName,
		Reference:      reference,
		RepositoryName: repositoryName,
		Context:        ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Artifact.AddLabel(params, c.AuthInfo)

	return handleSwaggerArtifactErrors(err)
}

func (c *RESTClient) CopyArtifact(ctx context.Context, from *CopyReference, projectName, repositoryName string) error {
	f, err := from.toString()
	if err != nil {
		return err
	}

	params := &artifact.CopyArtifactParams{
		From:           f,
		ProjectName:    projectName,
		RepositoryName: repositoryName,
		Context:        ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err = c.V2Client.Artifact.CopyArtifact(params, c.AuthInfo)

	return handleSwaggerArtifactErrors(err)
}

func (c *RESTClient) CreateTag(ctx context.Context, projectName, repositoryName, reference string, tag *model.Tag) error {
	params := &artifact.CreateTagParams{
		ProjectName:    projectName,
		Reference:      reference,
		RepositoryName: repositoryName,
		Tag:            tag,
		Context:        ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Artifact.CreateTag(params, c.AuthInfo)

	return handleSwaggerArtifactErrors(err)
}

func (c *RESTClient) DeleteTag(ctx context.Context, projectName, repositoryName, reference, tagName string) error {
	params := &artifact.DeleteTagParams{
		ProjectName:    projectName,
		Reference:      reference,
		RepositoryName: repositoryName,
		TagName:        tagName,
		Context:        ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Artifact.DeleteTag(params, c.AuthInfo)
	return handleSwaggerArtifactErrors(err)
}

func (c *RESTClient) GetArtifact(ctx context.Context, projectName, repositoryName, reference string) (*model.Artifact, error) {
	params := artifact.NewGetArtifactParams()
	params.WithTimeout(c.Options.Timeout)
	params.WithPage(&c.Options.Page)
	params.WithPageSize(&c.Options.PageSize)
	params.WithProjectName(projectName)
	params.WithRepositoryName(repositoryName)
	params.WithReference(reference)
	params.WithContext(ctx)

	resp, err := c.V2Client.Artifact.GetArtifact(params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerArtifactErrors(err)
	}

	return resp.Payload, nil
}

func (c *RESTClient) DeleteArtifact(ctx context.Context, projectName, repositoryName, reference string) error {
	params := artifact.NewDeleteArtifactParams()
	params.WithTimeout(c.Options.Timeout)
	params.WithProjectName(projectName)
	params.WithRepositoryName(repositoryName)
	params.WithReference(reference)
	params.WithContext(ctx)

	_, err := c.V2Client.Artifact.DeleteArtifact(params, c.AuthInfo)
	if err != nil {
		return handleSwaggerArtifactErrors(err)
	}

	return nil
}

func (c *RESTClient) ListArtifacts(ctx context.Context, projectName, repositoryName string) ([]*model.Artifact, error) {
	var artifacts []*model.Artifact
	page := c.Options.Page

	params := artifact.NewListArtifactsParams()
	params.WithContext(ctx)
	params.WithTimeout(c.Options.Timeout)
	params.Page = &page
	params.PageSize = &c.Options.PageSize
	params.Q = &c.Options.Query
	params.Sort = &c.Options.Sort
	params.WithProjectName(projectName)
	params.WithRepositoryName(repositoryName)

	for {
		resp, err := c.V2Client.Artifact.ListArtifacts(params, c.AuthInfo)
		if err != nil {
			return nil, handleSwaggerArtifactErrors(err)
		}

		if len(resp.Payload) == 0 {
			break
		}

		totalCount := resp.XTotalCount

		artifacts = append(artifacts, resp.Payload...)

		if int64(len(artifacts)) >= totalCount {
			break
		}

		page++
	}

	return artifacts, nil
}

func (c *RESTClient) ListTags(ctx context.Context, projectName, repositoryName, reference string) ([]*model.Tag, error) {
	var tags []*model.Tag
	page := c.Options.Page

	params := artifact.NewListTagsParams()
	params.Page = &page
	params.PageSize = &c.Options.PageSize
	params.WithProjectName(projectName)
	params.WithRepositoryName(repositoryName)
	params.WithReference(reference)
	params.Q = &c.Options.Query
	params.Sort = &c.Options.Sort
	params.WithContext(ctx)
	params.WithTimeout(c.Options.Timeout)

	for {
		resp, err := c.V2Client.Artifact.ListTags(params, c.AuthInfo)
		if err != nil {
			return nil, handleSwaggerArtifactErrors(err)
		}

		if len(resp.Payload) == 0 {
			break
		}

		totalCount := resp.XTotalCount

		tags = append(tags, resp.Payload...)

		if int64(len(tags)) >= totalCount {
			break
		}

		page++
	}

	return tags, nil
}

func (c *RESTClient) RemoveLabel(ctx context.Context, projectName, repositoryName, reference string, id int64) error {
	params := &artifact.RemoveLabelParams{
		LabelID:        id,
		ProjectName:    projectName,
		RepositoryName: repositoryName,
		Reference:      reference,
		Context:        ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Artifact.RemoveLabel(params, c.AuthInfo)
	if err != nil {
		return handleSwaggerArtifactErrors(err)
	}

	return nil
}

// TODO: Introduce this, once https://github.com/goharbor/harbor/issues/13468 is resolved.
//func (c *RESTClient) GetAddition(ctx context.Context, projectName, repositoryName, reference string, addition Addition) (interface{}, error) {
//	params := &artifact.GetAdditionParams{
//		ProjectName:    projectName,
//		RepositoryName: repositoryName,
//		Reference:      reference,
//		Addition:       addition.String(),
//		Context:        ctx,
//	}
//
//	params.WithTimeout(c.Options.Timeout)
//
//	resp, err := c.V2Client.Artifact.GetAddition(params, c.AuthInfo)
//	if err != nil {
//		return nil, err
//	}
//
//	return resp.Payload, nil
//}

//func (c *RESTClient) GetVulnerabilitiesAddition(ctx context.Context, projectName, repositoryName, reference string) (interface{}, error) {
//	params := artifact.NewGetVulnerabilitiesAdditionParams()
//	params.ProjectName = projectName
//	params.RepositoryName = repositoryName
//	params.Reference = reference
//	params.Context = ctx
//	params.WithTimeout(c.Options.Timeout)
//	xAcceptVulnerabilities := "application/vnd.security.vulnerability.report; version=1.1, application/vnd.scanner.adapter.vuln.report.harbor+json; version=1.0"
//	params.WithXAcceptVulnerabilities(&xAcceptVulnerabilities)
//
//	resp, err := c.V2Client.Artifact.GetVulnerabilitiesAddition(params, c.AuthInfo)
//	if err != nil {
//		return nil, handleSwaggerArtifactErrors(err)
//	}
//
//	return resp.Payload, nil
//}
