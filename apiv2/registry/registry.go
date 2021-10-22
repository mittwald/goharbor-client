package registry

import "C"
import (
	"context"

	"github.com/go-openapi/runtime"

	v2client "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/registry"
	modelv2 "github.com/mittwald/goharbor-client/v4/apiv2/model"
	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/common"
	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/config"
)

// RESTClient is a subclient for handling registry related actions.
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

type Client interface{}

// NewRegistry creates a new registry.
func (c *RESTClient) NewRegistry(ctx context.Context, reg *modelv2.Registry) error {
	params := &registry.CreateRegistryParams{
		Registry: reg,
		Context:  ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Registry.CreateRegistry(params, c.AuthInfo)

	return handleSwaggerRegistryErrors(err)
}

// GetRegistryByID returns a registry identified by ID.
// Returns an error if it cannot find a matching registry or when
// having difficulties talking to the API.
func (c *RESTClient) GetRegistryByID(ctx context.Context, id int64) (*modelv2.Registry, error) {
	params := &registry.GetRegistryParams{
		ID:      id,
		Context: ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.Registry.GetRegistry(params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerRegistryErrors(err)
	}

	if resp.Payload != nil {
		return nil, &common.ErrRegistryNotFound{}
	}

	return resp.Payload, nil
}

func (c *RESTClient) GetRegistryByName(ctx context.Context, name string) (*modelv2.Registry, error) {
	c.Options.Query = "name=" + name

	registries, err := c.ListRegistries(ctx)
	if err != nil {
		return nil, handleSwaggerRegistryErrors(err)
	}

	if len(registries) > 1 {
		return nil, &common.ErrMultipleResults{}
	}
	return registries[0], nil
}

func (c *RESTClient) ListRegistries(ctx context.Context) ([]*modelv2.Registry, error) {
	params := &registry.ListRegistriesParams{
		PageSize: &c.Options.PageSize,
		Q:        &c.Options.Query,
		Sort:     &c.Options.Sort,
		Context:  ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	resp, err := c.V2Client.Registry.ListRegistries(params, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerRegistryErrors(err)
	}

	if len(resp.Payload) == 0 {
		return nil, &common.ErrRegistryNotFound{}
	}

	return resp.Payload, nil
}

// DeleteRegistryByID deletes a registry identified by ID.
// Returns an error when no matching registry is found or when
// having difficulties talking to the API.
func (c *RESTClient) DeleteRegistryByID(ctx context.Context,
	id int64) error {
	params := &registry.DeleteRegistryParams{
		ID:      id,
		Context: ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Registry.DeleteRegistry(params, c.AuthInfo)

	return handleSwaggerRegistryErrors(err)
}

// UpdateRegistry updates a registry identified by ID with the provided RegistryUpdate 'r'.
func (c *RESTClient) UpdateRegistry(ctx context.Context, r *modelv2.RegistryUpdate, id int64) error {
	if r == nil {
		return &common.ErrRegistryNotProvided{}
	}

	params := &registry.UpdateRegistryParams{
		ID:       id,
		Registry: r,
		Context:  ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Registry.UpdateRegistry(params, c.AuthInfo)
	if err != nil {
		return handleSwaggerRegistryErrors(err)
	}

	return handleSwaggerRegistryErrors(err)
}
