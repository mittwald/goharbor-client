package label

import (
	"context"
	"fmt"

	"github.com/go-openapi/runtime"
	v2client "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/label"
	"github.com/mittwald/goharbor-client/v5/apiv2/model"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/config"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/util"
)

// RESTClient is a subclient for handling label related actions.
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
	CreateLabel(ctx context.Context, l *model.Label) error
	GetLabelByID(ctx context.Context, id int64) (*model.Label, error)
	ListLabels(ctx context.Context, name string, projectID *int64, scope Scope) ([]*model.Label, error)
	DeleteLabel(ctx context.Context, id int64) error
	UpdateLabel(ctx context.Context, id int64, l *model.Label) error
}

type Scope string

const (
	ScopeGlobal  Scope = "g"
	ScopeProject Scope = "p"
)

func (in Scope) String() string {
	return string(in)
}

func (c *RESTClient) CreateLabel(ctx context.Context, l *model.Label) error {
	params := &label.CreateLabelParams{
		Label:   l,
		Context: ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Label.CreateLabel(params, c.AuthInfo)
	if err != nil {
		return handleSwaggerLabelErrors(err)
	}

	return nil
}

func (c *RESTClient) GetLabelByID(ctx context.Context, id int64) (*model.Label, error) {
	params := &label.GetLabelByIDParams{
		LabelID: id,
		Context: ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.Label.GetLabelByID(params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerLabelErrors(err)
	}

	return resp.Payload, nil
}

func (c *RESTClient) ListLabels(ctx context.Context, name string, projectID *int64, scope Scope) ([]*model.Label, error) {
	switch scope {
	default:
		return nil, fmt.Errorf("invalid scope: %s", scope)
	case ScopeGlobal:
		if projectID != nil {
			return nil, fmt.Errorf("projectID must be nil for global scope")
		}
	case ScopeProject:
		if projectID == nil {
			return nil, fmt.Errorf("projectID must be set for project scope")
		}

	}

	params := &label.ListLabelsParams{
		Name:      &name,
		Page:      &c.Options.Page,
		PageSize:  &c.Options.PageSize,
		ProjectID: projectID,
		Q:         &c.Options.Query,
		Scope:     util.StringPtr(scope.String()),
		Sort:      &c.Options.Sort,
		Context:   ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.Label.ListLabels(params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerLabelErrors(err)
	}

	return resp.Payload, nil
}

func (c *RESTClient) DeleteLabel(ctx context.Context, id int64) error {
	params := &label.DeleteLabelParams{
		LabelID: id,
		Context: ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Label.DeleteLabel(params, c.AuthInfo)
	if err != nil {
		return handleSwaggerLabelErrors(err)
	}

	return nil
}

func (c *RESTClient) UpdateLabel(ctx context.Context, id int64, l *model.Label) error {
	// Name is the only required field for a label update
	if l.Name == "" {
		return fmt.Errorf("label name must be set")
	}

	params := &label.UpdateLabelParams{
		Label:   l,
		LabelID: id,
		Context: ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Label.UpdateLabel(params, c.AuthInfo)
	if err != nil {
		return handleSwaggerLabelErrors(err)
	}

	return nil
}
