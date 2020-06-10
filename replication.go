package goharborclient

import (
	"context"

	"github.com/mittwald/goharbor-client/api/v1.10.0/client/products"
	"github.com/mittwald/goharbor-client/api/v1.10.0/model"
)

// ReplicationRESTClient is a subclient for RESTClient handling user related actions.
type ReplicationRESTClient struct {
	parent *RESTClient
}

// NewReplication creates a new replication with the given arguments.
func (c *ReplicationRESTClient) NewReplication(ctx context.Context, destRegistry, srcRegistry *model.Registry,
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

	_, err := c.parent.Client.Products.PostReplicationPolicies(
		&products.PostReplicationPoliciesParams{
			Policy:  pReq,
			Context: ctx,
		}, c.parent.AuthInfo)

	err = handleSwaggerReplicationErrors(err)
	if err != nil {
		return nil, err
	}

	replication, err := c.Get(ctx, name)
	if err != nil {
		return nil, err
	}

	return replication, nil
}

// Get returns a replication identified by name.
// Returns an error if it cannot find a matching replication or when
// having difficulties talking to the API.
func (c *ReplicationRESTClient) Get(ctx context.Context, name string) (*model.ReplicationPolicy, error) {
	if name == "" {
		return nil, &ErrReplicationNotProvided{}
	}
	resp, err := c.parent.Client.Products.GetReplicationPolicies(
		&products.GetReplicationPoliciesParams{
			Name:    &name,
			Context: ctx,
		}, c.parent.AuthInfo)

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
func (c *ReplicationRESTClient) Delete(ctx context.Context,
	r *model.ReplicationPolicy) error {
	if r == nil {
		return &ErrReplicationNotProvided{}
	}

	replication, err := c.Get(ctx, r.Name)
	if err != nil {
		return err
	}

	if r.ID != replication.ID {
		return &ErrReplicationMismatch{}
	}

	_, err = c.parent.Client.Products.DeleteReplicationPoliciesID(
		&products.DeleteReplicationPoliciesIDParams{
			ID:      replication.ID,
			Context: ctx,
		}, c.parent.AuthInfo)

	return handleSwaggerReplicationErrors(err)
}

// Update updates a replication.
func (c *ReplicationRESTClient) Update(ctx context.Context,
	r *model.ReplicationPolicy) error {
	if r == nil {
		return &ErrReplicationNotProvided{}
	}

	replication, err := c.Get(ctx, r.Name)
	if err != nil {
		return err
	}

	if r.ID != replication.ID {
		return &ErrReplicationMismatch{}
	}

	_, err = c.parent.Client.Products.PutReplicationPoliciesID(
		&products.PutReplicationPoliciesIDParams{
			ID:      replication.ID,
			Policy:  r,
			Context: ctx,
		}, c.parent.AuthInfo)

	return handleSwaggerReplicationErrors(err)
}
