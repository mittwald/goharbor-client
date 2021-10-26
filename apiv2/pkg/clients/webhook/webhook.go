package webhook

import (
	"context"
	"strconv"

	"github.com/go-openapi/runtime"

	v2client "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/webhook"
	modelv2 "github.com/mittwald/goharbor-client/v5/apiv2/model"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/config"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/errors"
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
	ListProjectWebhookPolicies(ctx context.Context, projectID int) ([]*modelv2.WebhookPolicy, error)
	AddProjectWebhookPolicy(ctx context.Context, projectID int, policy *modelv2.WebhookPolicy) error
	UpdateProjectWebhookPolicy(ctx context.Context, projectID int, policy *modelv2.WebhookPolicy) error
	DeleteProjectWebhookPolicy(ctx context.Context, projectID int, policyID int64) error
}

// ListProjectWebhookPolicies returns a list of all webhook policies in project p.
func (c *RESTClient) ListProjectWebhookPolicies(ctx context.Context, projectID int) ([]*modelv2.WebhookPolicy, error) {
	params := &webhook.ListWebhookPoliciesOfProjectParams{
		PageSize:        &c.Options.PageSize,
		ProjectNameOrID: strconv.Itoa(projectID),
		Q:               &c.Options.Query,
		Sort:            &c.Options.Sort,
		Context:         ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.Webhook.ListWebhookPoliciesOfProject(params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerWebhookErrors(err)
	}

	return resp.Payload, nil
}

// AddProjectWebhookPolicy adds a webhook policy to project p.
func (c *RESTClient) AddProjectWebhookPolicy(ctx context.Context, projectID int, policy *modelv2.WebhookPolicy) error {
	if policy == nil {
		return &errors.ErrProjectNoWebhookPolicyProvided{}
	}

	params := &webhook.CreateWebhookPolicyOfProjectParams{
		Policy:          policy,
		ProjectNameOrID: strconv.Itoa(projectID),
		Context:         ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Webhook.CreateWebhookPolicyOfProject(params, c.AuthInfo)

	return handleSwaggerWebhookErrors(err)
}

// UpdateProjectWebhookPolicy updates the WebhookPolicy 'policy' in the project identified by 'projectID'.
func (c *RESTClient) UpdateProjectWebhookPolicy(ctx context.Context, projectID int, policy *modelv2.WebhookPolicy) error {
	if policy == nil {
		return &errors.ErrProjectNoWebhookPolicyProvided{}
	}

	params := &webhook.UpdateWebhookPolicyOfProjectParams{
		Policy:          policy,
		ProjectNameOrID: strconv.Itoa(projectID),
		WebhookPolicyID: policy.ID,
		Context:         ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Webhook.UpdateWebhookPolicyOfProject(params, c.AuthInfo)

	return handleSwaggerWebhookErrors(err)
}

// DeleteProjectWebhookPolicy deletes the webhook policy identified
// by 'policyID' from the project identified by 'projectID'.
func (c *RESTClient) DeleteProjectWebhookPolicy(ctx context.Context, projectID int, policyID int64) error {
	params := &webhook.DeleteWebhookPolicyOfProjectParams{
		ProjectNameOrID: strconv.Itoa(projectID),
		WebhookPolicyID: policyID,
		Context:         ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Webhook.DeleteWebhookPolicyOfProject(params, c.AuthInfo)

	return handleSwaggerWebhookErrors(err)
}
