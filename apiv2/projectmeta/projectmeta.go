package projectmeta

import (
	"context"

	"github.com/go-openapi/runtime"

	v2client "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client"
	projectmeta "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/project_metadata"
	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/common"
	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/config"
)

// RESTClient is a subclient for handling webhook related actions.
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
	AddProjectMetadata(ctx context.Context, projectNameOrID string, key common.MetadataKey, value string) error
	GetProjectMetadataValue(ctx context.Context, projectNameOrID string, key common.MetadataKey) (string, error)
	ListProjectMetadata(ctx context.Context, projectNameOrID string) (map[string]string, error)
	UpdateProjectMetadata(ctx context.Context, projectNameOrID string, key common.MetadataKey, value string) error
	DeleteProjectMetadataValue(ctx context.Context, projectNameOrID string, key common.MetadataKey) error
}

// AddProjectMetadata AddMetadata adds a metadata value using a specific key to the specified project.
func (c *RESTClient) AddProjectMetadata(ctx context.Context, projectNameOrID string, key common.MetadataKey, value string) error {
	params := &projectmeta.AddProjectMetadatasParams{
		Metadata: map[string]string{
			string(key): value,
		},
		ProjectNameOrID: projectNameOrID,
		Context:         ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.ProjectMetadata.AddProjectMetadatas(params, c.AuthInfo)

	return handleSwaggerProjectMetaErrors(err)
}

// GetProjectMetadataValue retrieves the corresponding metadata value to the key of the specified project.
func (c *RESTClient) GetProjectMetadataValue(ctx context.Context, projectNameOrID string, key common.MetadataKey) (string, error) {
	if key == "" {
		return "", &common.ErrProjectMetadataKeyUndefined{}
	}

	params := &projectmeta.GetProjectMetadataParams{
		MetaName:        key.String(),
		ProjectNameOrID: projectNameOrID,
		Context:         ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	meta, err := c.V2Client.ProjectMetadata.GetProjectMetadata(params, c.AuthInfo)
	if err != nil {
		return "", handleSwaggerProjectMetaErrors(err)
	}

	if meta == nil {
		return "", &common.ErrProjectMetadataUndefined{}
	}

	return retrieveMetadataValue(key, meta.Payload)
}

// ListProjectMetadata lists the metadata of project.
func (c *RESTClient) ListProjectMetadata(ctx context.Context, projectNameOrID string) (map[string]string, error) {
	params := &projectmeta.ListProjectMetadatasParams{
		ProjectNameOrID: projectNameOrID,
		Context:         ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	meta, err := c.V2Client.ProjectMetadata.ListProjectMetadatas(params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerProjectMetaErrors(err)
	}

	if meta.Payload != nil {
		return meta.Payload, nil
	}

	return nil, &common.ErrProjectMetadataUndefined{}
}

// UpdateProjectMetadata UpdateMetadata deletes the specified metadata key, if it exists and re-adds this metadata key with the given value.
// This function works around the faulty behaviour of the corresponding 'Update' endpoint of the Harbor API.
func (c *RESTClient) UpdateProjectMetadata(ctx context.Context, projectNameOrID string, key common.MetadataKey, value string) error {
	params := &projectmeta.UpdateProjectMetadataParams{
		MetaName: key.String(),
		Metadata: map[string]string{
			key.String(): value,
		},
		ProjectNameOrID: projectNameOrID,
		Context:         ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.ProjectMetadata.UpdateProjectMetadata(params, c.AuthInfo)

	return handleSwaggerProjectMetaErrors(err)
}

// DeleteProjectMetadataValue DeleteMetadataValue deletes metadata of project p given by key.
func (c *RESTClient) DeleteProjectMetadataValue(ctx context.Context, projectNameOrID string, key common.MetadataKey) error {
	params := &projectmeta.DeleteProjectMetadataParams{
		MetaName:        key.String(),
		ProjectNameOrID: projectNameOrID,
		Context:         ctx,
	}
	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.ProjectMetadata.DeleteProjectMetadata(params, c.AuthInfo)

	return handleSwaggerProjectMetaErrors(err)
}

// // retrieveMetadataValue returns the value of the metadata k that is contained in the project metadata m.
// // Returns an empty string plus an error when encountering a nil pointer, or if the requested key k is invalid.
func retrieveMetadataValue(k common.MetadataKey, m map[string]string) (string, error) {
	var r string

	switch k {
	case common.ProjectMetadataKeyEnableContentTrust:
		if m[k.String()] == "" {
			return "", &common.ErrProjectMetadataValueEnableContentTrustUndefined{}
		}
		r = m[k.String()]
	case common.ProjectMetadataKeyAutoScan:
		if m[k.String()] == "" {
			return "", &common.ErrProjectMetadataValueAutoScanUndefined{}
		}
		r = m[k.String()]
	case common.ProjectMetadataKeySeverity:
		if m[k.String()] == "" {
			return "", &common.ErrProjectMetadataValueSeverityUndefined{}
		}
		r = m[k.String()]
	case common.ProjectMetadataKeyReuseSysCVEAllowlist:
		if m[k.String()] == "" {
			return "", &common.ErrProjectMetadataValueReuseSysCveAllowlistUndefined{}
		}
		r = m[k.String()]
	case common.ProjectMetadataKeyPublic:
		if m[k.String()] == "" {
			return "", &common.ErrProjectMetadataValuePublicUndefined{}
		}
		r = m[k.String()]
	case common.ProjectMetadataKeyPreventVul:
		if m[k.String()] == "" {
			return "", &common.ErrProjectMetadataValuePreventVulUndefined{}
		}
		r = m[k.String()]
	case common.ProjectMetadataKeyRetentionID:
		if m[k.String()] == "" {
			return "", &common.ErrProjectMetadataValueRetentionIDUndefined{}
		}
		r = m[k.String()]
	default:
		return "", &common.ErrProjectInvalidRequest{}
	}

	return r, nil
}
