package replication

import (
	"context"

	"github.com/go-openapi/runtime"

	v2client "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client"
	replicationapi "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/replication"
	"github.com/mittwald/goharbor-client/v5/apiv2/model"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/config"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/errors"
)

// RESTClient is a subclient for handling replication related actions.
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
	NewReplicationPolicy(ctx context.Context, destRegistry, srcRegistry *model.Registry,
		replicateDeletion, override, enablePolicy bool,
		filters []*model.ReplicationFilter, trigger *model.ReplicationTrigger,
		destNamespace, description, name string) error
	GetReplicationPolicyByName(ctx context.Context, name string) (*model.ReplicationPolicy, error)
	ListReplicationPolicies(ctx context.Context) ([]*model.ReplicationPolicy, error)
	GetReplicationPolicyByID(ctx context.Context, id int64) (*model.ReplicationPolicy, error)
	DeleteReplicationPolicyByID(ctx context.Context, id int64) error
	UpdateReplicationPolicy(ctx context.Context, r *model.ReplicationPolicy, id int64) error
	TriggerReplicationExecution(ctx context.Context, r *model.StartReplicationExecution) error
	ListReplicationExecutions(ctx context.Context, policyID *int64, status, trigger *string) ([]*model.ReplicationExecution, error)
	GetReplicationExecutionByID(ctx context.Context, id int64) (*model.ReplicationExecution, error)
}

// NewReplicationPolicy creates a new replication policy with the given arguments.
func (c *RESTClient) NewReplicationPolicy(ctx context.Context, destRegistry, srcRegistry *model.Registry,
	replicateDeletion, override, enablePolicy bool,
	filters []*model.ReplicationFilter, trigger *model.ReplicationTrigger,
	destNamespace, description, name string,
) error {
	params := &replicationapi.CreateReplicationPolicyParams{
		Policy: &model.ReplicationPolicy{
			Description:               description,
			DestNamespace:             destNamespace,
			DestNamespaceReplaceCount: nil,
			DestRegistry:              destRegistry,
			Enabled:                   enablePolicy,
			Filters:                   filters,
			Name:                      name,
			Override:                  override,
			ReplicateDeletion:         replicateDeletion,
			Deletion:                  replicateDeletion,
			SrcRegistry:               srcRegistry,
			Trigger:                   trigger,
		},
		Context: ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Replication.CreateReplicationPolicy(params, c.AuthInfo)
	if err != nil {
		return handleSwaggerReplicationErrors(err)
	}

	return nil
}

// GetReplicationPolicyByName returns a replication identified by name.
func (c *RESTClient) GetReplicationPolicyByName(ctx context.Context, name string) (*model.ReplicationPolicy, error) {
	if name == "" {
		return nil, &ErrReplicationNotProvided{}
	}

	c.Options.Query = "name=" + name

	policies, err := c.ListReplicationPolicies(ctx)
	if err != nil {
		return nil, err
	}

	switch {
	case len(policies) > 1:
		return nil, &errors.ErrMultipleResults{}
	case len(policies) == 0:
		return nil, &errors.ErrNotFound{}
	}

	return policies[0], nil
}

func (c *RESTClient) ListReplicationPolicies(ctx context.Context) ([]*model.ReplicationPolicy, error) {
	var replicationPolicies []*model.ReplicationPolicy
	page := c.Options.Page

	params := &replicationapi.ListReplicationPoliciesParams{
		Page:     &page,
		PageSize: &c.Options.PageSize,
		Q:        &c.Options.Query,
		Sort:     &c.Options.Sort,
		Context:  ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	for {
		resp, err := c.V2Client.Replication.ListReplicationPolicies(params, c.AuthInfo)
		if err != nil {
			return nil, handleSwaggerReplicationErrors(err)
		}

		totalCount := resp.XTotalCount

		replicationPolicies = append(replicationPolicies, resp.Payload...)

		if int64(len(replicationPolicies)) >= totalCount {
			break
		}

		page++
	}

	if len(replicationPolicies) > 0 {
		return replicationPolicies, nil
	}

	return nil, &errors.ErrNotFound{}
}

// GetReplicationPolicyByID returns a replication identified by id.
func (c *RESTClient) GetReplicationPolicyByID(ctx context.Context, id int64) (*model.ReplicationPolicy, error) {
	params := &replicationapi.GetReplicationPolicyParams{
		ID:      id,
		Context: ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.Replication.GetReplicationPolicy(params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerReplicationErrors(err)
	}

	if resp.Payload.ID == id {
		return resp.Payload, nil
	}

	return nil, &errors.ErrNotFound{}
}

// DeleteReplicationPolicyByID deletes a replication policy identified by id.
func (c *RESTClient) DeleteReplicationPolicyByID(ctx context.Context, id int64) error {
	params := &replicationapi.DeleteReplicationPolicyParams{
		ID:      id,
		Context: ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Replication.DeleteReplicationPolicy(params, c.AuthInfo)

	return handleSwaggerReplicationErrors(err)
}

// UpdateReplicationPolicy updates the replication policy identified by id with the provided policy 'r'.
func (c *RESTClient) UpdateReplicationPolicy(ctx context.Context, r *model.ReplicationPolicy, id int64) error {
	if r == nil {
		return &ErrReplicationNotProvided{}
	}

	params := &replicationapi.UpdateReplicationPolicyParams{
		ID:      id,
		Policy:  r,
		Context: ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Replication.UpdateReplicationPolicy(params, c.AuthInfo)

	return handleSwaggerReplicationErrors(err)
}

// TriggerReplicationExecution triggers the execution of a replication 'r'.
func (c *RESTClient) TriggerReplicationExecution(ctx context.Context, r *model.StartReplicationExecution) error {
	if r == nil {
		return &ErrReplicationExecutionNotProvided{}
	}

	params := &replicationapi.StartReplicationParams{
		Execution: r,
		Context:   ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Replication.StartReplication(params, c.AuthInfo)

	return handleSwaggerReplicationErrors(err)
}

// ListReplicationExecutions lists replication executions specified by execution ID, status or trigger.
// Specifying the property "policy_id" will return executions of the specified policy.
func (c *RESTClient) ListReplicationExecutions(ctx context.Context, policyID *int64, status, trigger *string) ([]*model.ReplicationExecution, error) {
	var replicationExecutions []*model.ReplicationExecution
	page := c.Options.Page

	params := &replicationapi.ListReplicationExecutionsParams{
		Page:     &page,
		PageSize: &c.Options.PageSize,
		PolicyID: policyID,
		Sort:     &c.Options.Sort,
		Status:   status,
		Trigger:  trigger,
		Context:  ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	for {
		resp, err := c.V2Client.Replication.ListReplicationExecutions(params, c.AuthInfo)
		if err != nil {
			return nil, handleSwaggerReplicationErrors(err)
		}

		if len(resp.Payload) == 0 {
			break
		}

		totalCount := resp.XTotalCount

		replicationExecutions = append(replicationExecutions, resp.Payload...)

		if int64(len(replicationExecutions)) >= totalCount {
			break
		}

		page++
	}

	return replicationExecutions, nil
}

// GetReplicationExecutionByID returns a replication execution specified by ID.
func (c *RESTClient) GetReplicationExecutionByID(ctx context.Context, id int64) (*model.ReplicationExecution, error) {
	params := &replicationapi.GetReplicationExecutionParams{
		ID:      id,
		Context: ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.Replication.GetReplicationExecution(params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerReplicationErrors(err)
	}

	if resp.Payload.ID != id {
		return nil, &ErrReplicationExecutionReplicationIDMismatch{}
	}

	return resp.Payload, nil
}
