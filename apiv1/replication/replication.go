package replication

import (
	"context"

	"github.com/go-openapi/runtime"
	"github.com/mittwald/goharbor-client/v2/apiv1/internal/api/client"

	"github.com/mittwald/goharbor-client/v2/apiv1/internal/api/client/products"
	model "github.com/mittwald/goharbor-client/v2/apiv1/model"
)

// RESTClient is a subclient for handling replication related actions.
type RESTClient struct {
	// The swagger client
	Client *client.Harbor

	// AuthInfo contain auth information, which are provided on API calls.
	AuthInfo runtime.ClientAuthInfoWriter
}

func NewClient(cl *client.Harbor, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		Client:   cl,
		AuthInfo: authInfo,
	}
}

type Client interface {
	NewReplicationPolicy(ctx context.Context, destRegistry, srcRegistry *model.Registry,
		replicateDeletion, override, enablePolicy bool,
		filters []*model.ReplicationFilter, trigger *model.ReplicationTrigger,
		destNamespace, description, name string) (*model.ReplicationPolicy, error)
	GetReplicationPolicy(ctx context.Context, name string) (*model.ReplicationPolicy, error)
	GetReplicationPolicyByID(ctx context.Context, id int64) (*model.ReplicationPolicy, error)
	DeleteReplicationPolicy(ctx context.Context, r *model.ReplicationPolicy) error
	UpdateReplicationPolicy(ctx context.Context, r *model.ReplicationPolicy) error

	TriggerReplicationExecution(ctx context.Context, r *model.ReplicationExecution) error
	GetReplicationExecutions(ctx context.Context, r *model.ReplicationExecution) ([]*model.ReplicationExecution, error)
	GetReplicationExecutionsByID(ctx context.Context,
		r *model.ReplicationExecution) (*model.ReplicationExecution, error)
}

// NewReplication creates a new replication with the given arguments.
func (c *RESTClient) NewReplicationPolicy(ctx context.Context, destRegistry, srcRegistry *model.Registry,
	replicateDeletion, override, enablePolicy bool,
	filters []*model.ReplicationFilter, trigger *model.ReplicationTrigger,
	destNamespace, description, name string) (*model.ReplicationPolicy, error) {
	pReq := &model.ReplicationPolicy{
		Deletion:      replicateDeletion,
		Description:   description,
		DestNamespace: destNamespace,
		DestRegistry:  destRegistry,
		Enabled:       enablePolicy,
		Filters:       filters,
		Name:          name,
		Override:      override,
		SrcRegistry:   srcRegistry,
		Trigger:       trigger,
	}

	_, err := c.Client.Products.PostReplicationPolicies(
		&products.PostReplicationPoliciesParams{
			Policy:  pReq,
			Context: ctx,
		}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerReplicationErrors(err)
	}

	replication, err := c.GetReplicationPolicy(ctx, name)
	if err != nil {
		return nil, err
	}

	return replication, nil
}

// GetReplicationPolicy returns a replication identified by name.
// Returns an error if it cannot find a matching replication or when
// having difficulties talking to the API.
func (c *RESTClient) GetReplicationPolicy(ctx context.Context, name string) (*model.ReplicationPolicy, error) {
	if name == "" {
		return nil, &ErrReplicationNotProvided{}
	}
	resp, err := c.Client.Products.GetReplicationPolicies(
		&products.GetReplicationPoliciesParams{
			Name:    &name,
			Context: ctx,
		}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerReplicationErrors(err)
	}

	for _, p := range resp.Payload {
		if p.Name == name {
			return p, nil
		}
	}

	return nil, &ErrReplicationNotFound{}
}

// GetReplicationPolicyByID returns a replication identified by id.
// Returns an error if it cannot find a matching replication or when
// having difficulties talking to the API.
func (c *RESTClient) GetReplicationPolicyByID(ctx context.Context, id int64) (*model.ReplicationPolicy, error) {
	resp, err := c.Client.Products.GetReplicationPoliciesID(
		&products.GetReplicationPoliciesIDParams{
			ID:      id,
			Context: ctx,
		}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerReplicationErrors(err)
	}

	if resp.Payload.ID == id {
		return resp.Payload, nil
	}

	return nil, &ErrReplicationNotFound{}
}

// Delete deletes a replication.
// Returns an error when no matching replication is found or when
// having difficulties talking to the API.
func (c *RESTClient) DeleteReplicationPolicy(ctx context.Context,
	r *model.ReplicationPolicy) error {
	if r == nil {
		return &ErrReplicationNotProvided{}
	}

	replication, err := c.GetReplicationPolicy(ctx, r.Name)
	if err != nil {
		return err
	}

	if r.ID != replication.ID {
		return &ErrReplicationMismatch{}
	}

	_, err = c.Client.Products.DeleteReplicationPoliciesID(
		&products.DeleteReplicationPoliciesIDParams{
			ID:      replication.ID,
			Context: ctx,
		}, c.AuthInfo)

	return handleSwaggerReplicationErrors(err)
}

// Update updates a replication.
func (c *RESTClient) UpdateReplicationPolicy(ctx context.Context,
	r *model.ReplicationPolicy) error {
	if r == nil {
		return &ErrReplicationNotProvided{}
	}

	replication, err := c.GetReplicationPolicy(ctx, r.Name)
	if err != nil {
		return err
	}

	if r.ID != replication.ID {
		return &ErrReplicationMismatch{}
	}

	_, err = c.Client.Products.PutReplicationPoliciesID(
		&products.PutReplicationPoliciesIDParams{
			ID:      replication.ID,
			Policy:  r,
			Context: ctx,
		}, c.AuthInfo)

	return handleSwaggerReplicationErrors(err)
}

// TriggerReplicationExecution triggers the execution of a replication where only the property "policy_id" is needed.
func (c *RESTClient) TriggerReplicationExecution(ctx context.Context, r *model.ReplicationExecution) error {
	if r == nil {
		return &ErrReplicationExecutionNotProvided{}
	}

	if _, err := c.GetReplicationPolicyByID(ctx, r.PolicyID); err != nil {
		return &ErrReplicationExecutionReplicationPolicyIDNotFound{}
	}

	_, err := c.Client.Products.PostReplicationExecutions(
		&products.PostReplicationExecutionsParams{
			Execution: r,
			Context:   ctx,
		}, c.AuthInfo)

	return handleSwaggerReplicationErrors(err)
}

// GetReplicationExecutions lists replication executions specified by ID, status or trigger.
// Specifying the property "policy_id" will return executions of the specified policy.
func (c *RESTClient) GetReplicationExecutions(ctx context.Context,
	r *model.ReplicationExecution) ([]*model.ReplicationExecution, error) {
	if _, err := c.GetReplicationPolicyByID(ctx, r.PolicyID); err != nil {
		return nil, &ErrReplicationExecutionReplicationPolicyIDNotFound{}
	}

	resp, err := c.Client.Products.GetReplicationExecutions(
		&products.GetReplicationExecutionsParams{
			PolicyID: &r.ID,
			Status:   &r.Status,
			Trigger:  &r.Trigger,
			Context:  ctx,
		}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerReplicationErrors(err)
	}

	return resp.Payload, nil
}

func (c *RESTClient) GetReplicationExecutionsByID(ctx context.Context,
	id int64) (*model.ReplicationExecution, error) {
	if _, err := c.GetReplicationPolicyByID(ctx, id); err != nil {
		return nil, &ErrReplicationExecutionReplicationPolicyIDNotFound{}
	}

	resp, err := c.Client.Products.GetReplicationExecutionsID(
		&products.GetReplicationExecutionsIDParams{
			ID:      id,
			Context: ctx,
		}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerReplicationErrors(err)
	}

	if resp.Payload.ID != id {
		return nil, &ErrReplicationExecutionReplicationIDMismatch{}
	}

	return resp.Payload, nil
}
