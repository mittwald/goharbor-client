package project

import (
	"context"

	"github.com/go-openapi/runtime"
	"github.com/mittwald/goharbor-client/apiv2/internal/api/client/project"
	"github.com/mittwald/goharbor-client/apiv2/internal/legacyapi/client/products"
	modelv2 "github.com/mittwald/goharbor-client/apiv2/model"
)

// AddProjectMetadataV2 adds metadata with a specific key and value to project p.
func (c *RESTClient) AddProjectMetadataV2(ctx context.Context, p *modelv2.Project, key MetadataKey, value string) error {
	if p == nil {
		return &ErrProjectNotProvided{}
	}

	_, err := c.LegacyClient.Products.PostProjectsProjectIDMetadatas(
		&products.PostProjectsProjectIDMetadatasParams{
			Metadata:  getProjectMetadataByKey(key, value),
			ProjectID: int64(p.ProjectID),
			Context:   ctx,
		}, c.AuthInfo)

	err = handleSwaggerProjectErrors(err)
	if err != nil {
		t, ok := err.(*runtime.APIError)
		if ok && t.Code == 409 {
			// Unspecified error that returns when a metadata key is already defined.
			return &ErrProjectMetadataAlreadyExists{}
		}

		return err
	}

	return handleSwaggerProjectErrors(err)
}

// GetProjectMetadataValueV2 returns a ProjectMetadata value.
func (c *RESTClient) GetProjectMetadataValueV2(ctx context.Context, projectID int64, key MetadataKey) (string, error) {
	resp, err := c.GetProjectV2(ctx, projectID)
	if err != nil {
		return "", handleSwaggerProjectErrors(err)
	}

	var result string

	switch key {
	case ProjectMetadataKeyEnableContentTrust:
		result = *resp.Metadata.EnableContentTrust
	case ProjectMetadataKeyAutoScan:
		result = *resp.Metadata.AutoScan
	case ProjectMetadataKeySeverity:
		result = *resp.Metadata.Severity
	case ProjectMetadataKeyReuseSysCVEWhitelist:
		result = *resp.Metadata.ReuseSysCveAllowlist
	case ProjectMetadataKeyPublic:
		result = resp.Metadata.Public
	case ProjectMetadataKeyPreventVul:
		result = *resp.Metadata.PreventVul
	case ProjectMetadataKeyRetentionID:
		result = *resp.Metadata.RetentionID
	default:
		return "", &ErrProjectInvalidRequest{}
	}

	return result, nil
}

// GetProjectV2 returns a Harbor project.
func (c *RESTClient) GetProjectV2(ctx context.Context, projectID int64) (*modelv2.Project, error) {
	resp, err := c.V2Client.Project.GetProject(&project.GetProjectParams{
		ProjectID: projectID,
		Context:   ctx,
	}, c.AuthInfo)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
