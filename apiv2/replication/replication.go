package replication

import (
	"context"

	"github.com/go-openapi/runtime"
	v2client "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client"
	replicationapi "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/replication"
	"github.com/mittwald/goharbor-client/v4/apiv2/internal/legacyapi/client"
	modelv2 "github.com/mittwald/goharbor-client/v4/apiv2/model"

	"github.com/mittwald/goharbor-client/v4/apiv2/internal/legacyapi/client/products"
	legacymodel "github.com/mittwald/goharbor-client/v4/apiv2/model/legacy"
)

// RESTClient is a subclient for handling replication related actions.
type RESTClient struct {
	// The legacy swagger client
	LegacyClient *client.Harbor

	// The new client of the harbor v2 API
	V2Client *v2client.Harbor

	// AuthInfo contains the auth information that is provided on API calls.
	AuthInfo runtime.ClientAuthInfoWriter
}

func NewClient(legacyClient *client.Harbor, v2Client *v2client.Harbor, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		LegacyClient: legacyClient,
		V2Client:     v2Client,
		AuthInfo:     authInfo,
	}
}

type Client interface {
	NewReplicationPolicy(ctx context.Context, destRegistry, srcRegistry *legacymodel.Registry,
		replicateDeletion, override, enablePolicy bool,
		filters []*legacymodel.ReplicationFilter, trigger *legacymodel.ReplicationTrigger,
		destNamespace, description, name string) (*legacymodel.ReplicationPolicy, error)
	GetReplicationPolicy(ctx context.Context, name string) (*legacymodel.ReplicationPolicy, error)
	GetReplicationPolicyByID(ctx context.Context, id int64) (*legacymodel.ReplicationPolicy, error)
	DeleteReplicationPolicy(ctx context.Context, r *legacymodel.ReplicationPolicy) error
	UpdateReplicationPolicy(ctx context.Context, r *legacymodel.ReplicationPolicy) error
	TriggerReplicationExecution(ctx context.Context, r *modelv2.StartReplicationExecution) error
	GetReplicationExecutions(ctx context.Context, r *modelv2.ReplicationExecution) ([]*modelv2.ReplicationExecution, error)
	GetReplicationExecutionByID(ctx context.Context, id int64) (*modelv2.ReplicationExecution, error)
}

// NewReplication creates a new replication with the given arguments.
func (c *RESTClient) NewReplicationPolicy(ctx context.Context, destRegistry, srcRegistry *legacymodel.Registry,
	replicateDeletion, override, enablePolicy bool,
	filters []*legacymodel.ReplicationFilter, trigger *legacymodel.ReplicationTrigger,
	destNamespace, description, name string) (*legacymodel.ReplicationPolicy, error) {
	pReq := &legacymodel.ReplicationPolicy{
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

	_, err := c.LegacyClient.Products.PostReplicationPolicies(
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
func (c *RESTClient) GetReplicationPolicy(ctx context.Context, name string) (*legacymodel.ReplicationPolicy, error) {
	if name == "" {
		return nil, &ErrReplicationNotProvided{}
	}
	resp, err := c.LegacyClient.Products.GetReplicationPolicies(
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
func (c *RESTClient) GetReplicationPolicyByID(ctx context.Context, id int64) (*legacymodel.ReplicationPolicy, error) {
	resp, err := c.LegacyClient.Products.GetReplicationPoliciesID(
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
func (c *RESTClient) DeleteReplicationPolicy(ctx context.Context, r *legacymodel.ReplicationPolicy) error {
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

	_, err = c.LegacyClient.Products.DeleteReplicationPoliciesID(
		&products.DeleteReplicationPoliciesIDParams{
			ID:      replication.ID,
			Context: ctx,
		}, c.AuthInfo)

	return handleSwaggerReplicationErrors(err)
}

// Update updates a replication.
func (c *RESTClient) UpdateReplicationPolicy(ctx context.Context, r *legacymodel.ReplicationPolicy) error {
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

	_, err = c.LegacyClient.Products.PutReplicationPoliciesID(
		&products.PutReplicationPoliciesIDParams{
			ID:      replication.ID,
			Policy:  r,
			Context: ctx,
		}, c.AuthInfo)

	return handleSwaggerReplicationErrors(err)
}

// TriggerReplicationExecution triggers the execution of a replication 'r'.
func (c *RESTClient) TriggerReplicationExecution(ctx context.Context, r *modelv2.StartReplicationExecution) error {
	if r == nil {
		return &ErrReplicationExecutionNotProvided{}
	}

	_, err := c.V2Client.Replication.StartReplication(
		&replicationapi.StartReplicationParams{
			Execution: r,
			Context:   ctx,
		}, c.AuthInfo)

	return handleSwaggerReplicationErrors(err)
}

// GetReplicationExecutions lists replication executions specified by execution ID, status or trigger.
// Specifying the property "policy_id" will return executions of the specified policy.
func (c *RESTClient) GetReplicationExecutions(ctx context.Context, r *modelv2.ReplicationExecution) ([]*modelv2.ReplicationExecution, error) {
	resp, err := c.V2Client.Replication.ListReplicationExecutions(&replicationapi.ListReplicationExecutionsParams{
		PolicyID: &r.PolicyID,
		Status:   &r.Status,
		Trigger:  &r.Trigger,
		Context:  ctx,
	}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerReplicationErrors(err)
	}

	return resp.Payload, nil
}

// GetReplicationExecutionByID returns a replication execution specified by ID.
func (c *RESTClient) GetReplicationExecutionByID(ctx context.Context, id int64) (*modelv2.ReplicationExecution, error) {
	resp, err := c.V2Client.Replication.GetReplicationExecution(&replicationapi.GetReplicationExecutionParams{
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
