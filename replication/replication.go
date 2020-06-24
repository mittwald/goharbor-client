package replication

import (
	"context"

	"github.com/go-openapi/runtime"
	"github.com/mittwald/goharbor-client/internal/api/v1.10.0/client"

	"github.com/mittwald/goharbor-client/internal/api/v1.10.0/client/products"
	"github.com/mittwald/goharbor-client/internal/api/v1.10.0/model"
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
	NewReplication(ctx context.Context, destRegistry, srcRegistry *model.Registry,
		replicateDeletion, override, enablePolicy bool,
		filters []*model.ReplicationFilter, trigger *model.ReplicationTrigger,
		destNamespace, description, name string) (*model.ReplicationPolicy, error)
	GetReplication(ctx context.Context, name string) (*model.ReplicationPolicy, error)
	DeleteReplication(ctx context.Context, r *model.ReplicationPolicy) error
	UpdateReplication(ctx context.Context, r *model.ReplicationPolicy) error
}

// NewReplication creates a new replication with the given arguments.
func (c *RESTClient) NewReplication(ctx context.Context, destRegistry, srcRegistry *model.Registry,
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

	err = handleSwaggerReplicationErrors(err)
	if err != nil {
		return nil, err
	}

	replication, err := c.GetReplication(ctx, name)
	if err != nil {
		return nil, err
	}

	return replication, nil
}

// Get returns a replication identified by name.
// Returns an error if it cannot find a matching replication or when
// having difficulties talking to the API.
func (c *RESTClient) GetReplication(ctx context.Context, name string) (*model.ReplicationPolicy, error) {
	if name == "" {
		return nil, &ErrReplicationNotProvided{}
	}
	resp, err := c.Client.Products.GetReplicationPolicies(
		&products.GetReplicationPoliciesParams{
			Name:    &name,
			Context: ctx,
		}, c.AuthInfo)

	err = handleSwaggerReplicationErrors(err)
	if err != nil {
		return nil, err
	}

	for _, p := range resp.Payload {
		if p.Name == name {
			return p, nil
		}
	}

	return nil, &ErrReplicationNotFound{}
}

// Delete deletes a replication.
// Returns an error when no matching replication is found or when
// having difficulties talking to the API.
func (c *RESTClient) DeleteReplication(ctx context.Context,
	r *model.ReplicationPolicy) error {
	if r == nil {
		return &ErrReplicationNotProvided{}
	}

	replication, err := c.GetReplication(ctx, r.Name)
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
func (c *RESTClient) UpdateReplication(ctx context.Context,
	r *model.ReplicationPolicy) error {
	if r == nil {
		return &ErrReplicationNotProvided{}
	}

	replication, err := c.GetReplication(ctx, r.Name)
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
