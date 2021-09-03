package artifact

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/go-openapi/runtime"

	v2client "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/artifact"
	modelv2 "github.com/mittwald/goharbor-client/v4/apiv2/model"
)

type (
	MIMEType       string
	RepositoryName string
)

const (
	MIMETypeV10 MIMEType = "application/vnd.scanner.adapter.vuln.report.harbor+json; version=1.0"
	MIMETypeV11 MIMEType = "application/vnd.security.vulnerability.report; version=1.1"
)

// RESTClient is a subclient for handling user related actions.
type RESTClient struct {
	// The new client of the harbor v2 API
	V2Client *v2client.Harbor

	// AuthInfo contains the auth information that is provided on API calls.
	AuthInfo runtime.ClientAuthInfoWriter
}

func NewClient(v2Client *v2client.Harbor, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		V2Client: v2Client,
		AuthInfo: authInfo,
	}
}

type Client interface {
	GetArtifact(ctx context.Context, options ParamOptions) (*modelv2.Artifact, error)
	ListArtifacts(ctx context.Context, options ParamOptions) ([]*modelv2.Artifact, error)
	DeleteArtifact(ctx context.Context, options ParamOptions) error
	CreateTag(ctx context.Context, options ParamOptions) error
	DeleteTag(ctx context.Context, options ParamOptions) error
	GetRepositoryVulnerabilities(ctx context.Context, projectName, repositoryName, reference string,
		MIMETypes []MIMEType) (string, error)
}

// ParamOptions holds the necessary options for artifact API requests.
type ParamOptions struct {
	PageSize               int64
	ProjectName            string
	RepositoryName         string
	Reference              string
	WithLabel              bool
	WithScanOverview       bool
	WithSignature          bool
	WithTag                bool
	WithImmutableStatus    bool
	XAcceptVulnerabilities string
	Query                  string
	Tag                    *modelv2.Tag
}

// setParamOptions sets the default params for 'in', assuming 'in' is a parameter type used by the artifact API.
// This comes in handy as the different request param types share most of their fields.
func setParamOptions(in interface{}, opts ParamOptions) error {
	truePtr := true

	switch in := in.(type) {
	default:
		return fmt.Errorf("could not assert %T type", in)
	case *artifact.GetArtifactParams:
		in.ProjectName = opts.ProjectName
		in.Reference = opts.Reference
		in.RepositoryName = url.QueryEscape(opts.RepositoryName)
		in.WithScanOverview = &opts.WithScanOverview
		in.PageSize = &opts.PageSize
		in.WithLabel = &opts.WithLabel
		in.WithSignature = &opts.WithSignature
		in.WithTag = &opts.WithTag
		// Implicitly set 'WithImmutableStatus' when 'WithTag' is set.
		if opts.WithTag {
			in.WithImmutableStatus = &truePtr
		}
		in.XAcceptVulnerabilities = &opts.XAcceptVulnerabilities
	case *artifact.ListArtifactsParams:
		in.ProjectName = opts.ProjectName
		in.RepositoryName = url.QueryEscape(opts.RepositoryName)
		in.WithScanOverview = &opts.WithScanOverview
		in.PageSize = &opts.PageSize
		in.WithLabel = &opts.WithLabel
		// Implicitly set 'WithImmutableStatus' when 'WithTag' is set.
		if opts.WithTag {
			in.WithImmutableStatus = &truePtr
		}
		in.WithSignature = &opts.WithSignature
		in.WithTag = &opts.WithTag
		in.XAcceptVulnerabilities = &opts.XAcceptVulnerabilities
		in.Q = &opts.Query
	case *artifact.DeleteArtifactParams:
		in.ProjectName = opts.ProjectName
		in.Reference = opts.Reference
		in.RepositoryName = url.QueryEscape(opts.RepositoryName)
	case *artifact.CreateTagParams:
		in.ProjectName = opts.ProjectName
		in.Reference = opts.Reference
		in.RepositoryName = url.QueryEscape(opts.RepositoryName)
		in.Tag = opts.Tag
	case *artifact.DeleteTagParams:
		in.ProjectName = opts.ProjectName
		in.Reference = opts.Reference
		in.RepositoryName = url.QueryEscape(opts.RepositoryName)
		in.TagName = opts.Tag.Name
	}

	return nil
}

// GetArtifact returns a specific artifact of a project's repository based on the provided 'options'.
func (c *RESTClient) GetArtifact(ctx context.Context, options ParamOptions) (*modelv2.Artifact, error) {
	params := &artifact.GetArtifactParams{}

	if err := setParamOptions(params.WithDefaults().WithContext(ctx), options); err != nil {
		return nil, fmt.Errorf("error setting parameter options: %w", err)
	}

	resp, err := c.V2Client.Artifact.GetArtifact(params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerArtifactErrors(err)
	}

	return resp.Payload, nil
}

// ListArtifacts lists the artifacts of a project's repository based on the provided 'options'.
func (c *RESTClient) ListArtifacts(ctx context.Context, options ParamOptions) ([]*modelv2.Artifact, error) {
	params := &artifact.ListArtifactsParams{}

	if err := setParamOptions(params.WithDefaults().WithContext(ctx), options); err != nil {
		return nil, fmt.Errorf("error setting parameter options: %w", err)
	}

	resp, err := c.V2Client.Artifact.ListArtifacts(params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerArtifactErrors(err)
	}

	return resp.Payload, nil
}

// DeleteArtifact deletes a specific artifact of a repository based on the provided 'options'.
func (c *RESTClient) DeleteArtifact(ctx context.Context, options ParamOptions) error {
	params := &artifact.DeleteArtifactParams{}

	if err := setParamOptions(params.WithDefaults().WithContext(ctx), options); err != nil {
		return fmt.Errorf("error setting parameter options: %w", err)
	}

	if _, err := c.V2Client.Artifact.DeleteArtifact(params, c.AuthInfo); err != nil {
		return handleSwaggerArtifactErrors(err)
	}

	return nil
}

// CreateTag creates a tag (options.Tag) for the repository specified in 'options'.
func (c *RESTClient) CreateTag(ctx context.Context, options ParamOptions) error {
	params := &artifact.CreateTagParams{}

	if err := setParamOptions(params.WithDefaults().WithContext(ctx), options); err != nil {
		return fmt.Errorf("error setting parameter options: %w", err)
	}

	if _, err := c.V2Client.Artifact.CreateTag(params, c.AuthInfo); err != nil {
		return handleSwaggerArtifactErrors(err)
	}

	return nil
}

// DeleteTag deletes the tag (options.Tag.Name) of the repository specified in 'options'.
func (c *RESTClient) DeleteTag(ctx context.Context, options ParamOptions) error {
	params := &artifact.DeleteTagParams{}

	if err := setParamOptions(params.WithDefaults().WithContext(ctx), options); err != nil {
		return fmt.Errorf("error setting parameter options: %w", err)
	}

	if _, err := c.V2Client.Artifact.DeleteTag(params, c.AuthInfo); err != nil {
		return err
	}

	return nil
}

// GetRepositoryVulnerabilities returns a JSON formatted string containing vulnerability reports for the specified artifact.
func (c *RESTClient) GetRepositoryVulnerabilities(ctx context.Context, projectName, repositoryName, reference string,
	MIMETypes []MIMEType) (string, error) {
	var m []string

	for i := range MIMETypes {
		m = append(m, string(MIMETypes[i]))
	}

	requestedMIMETypes := strings.Join(m, ",")

	v, err := c.V2Client.Artifact.GetVulnerabilitiesAddition(&artifact.GetVulnerabilitiesAdditionParams{
		XAcceptVulnerabilities: &requestedMIMETypes,
		ProjectName:            projectName,
		Reference:              reference,
		RepositoryName:         url.QueryEscape(repositoryName),
		Context:                ctx,
	}, c.AuthInfo)
	if err != nil {
		return "", handleSwaggerArtifactErrors(err)
	}

	return v.Payload, nil
}
