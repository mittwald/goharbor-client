package goharborclient

import (
	"context"
	"github.com/mittwald/goharbor-client/user"

	"github.com/mittwald/goharbor-client/api/v1.10.0/client/products"
	"github.com/mittwald/goharbor-client/api/v1.10.0/model"
)

// RegistryRESTClient is a subclient for RESTClient handling registry related
// actions.
type RegistryRESTClient struct {
	parent *user.RESTClient
}

// NewRegistry creates a new project with name as project name.
// CountLimit and StorageLimit limits space and access for this project.
// Returns the project as it is stored inside Harbor or an error,
// if the project could not be created.
func (c *RegistryRESTClient) NewRegistry(ctx context.Context, name, registryType, url string,
	credential *model.RegistryCredential, insecure bool) (*model.Registry, error) {
	rReq := &model.Registry{
		Credential: credential,
		Insecure:   insecure,
		Name:       name,
		Type:       registryType,
		URL:        url,
	}

	_, err := c.parent.Client.Products.PostRegistries(
		&products.PostRegistriesParams{
			Registry: rReq,
			Context:  ctx,
		}, c.parent.AuthInfo)

	err = handleSwaggerRegistryErrors(err)
	if err != nil {
		return nil, err
	}

	registry, err := c.Get(ctx, name)
	if err != nil {
		return nil, err
	}

	return registry, nil
}

// Get returns a registry identified by name.
// Returns an error if it cannot find a matching registry or when
// having difficulties talking to the API.
func (c *RegistryRESTClient) Get(ctx context.Context, name string) (*model.Registry, error) {
	if name == "" {
		return nil, &ErrRegistryNotProvided{}
	}
	resp, err := c.parent.Client.Products.GetRegistries(
		&products.GetRegistriesParams{
			Name:    &name,
			Context: ctx,
		}, c.parent.AuthInfo)

	err = handleSwaggerRegistryErrors(err)
	if err != nil {
		return nil, err
	}

	for _, r := range resp.Payload {
		if r.Name == name {
			return r, nil
		}
	}

	return nil, &ErrRegistryNotFound{}
}

// Delete deletes a registry.
// Returns an error when no matching registry is found or when
// having difficulties talking to the API.
func (c *RegistryRESTClient) Delete(ctx context.Context,
	r *model.Registry) error {
	if r == nil {
		return &ErrRegistryNotProvided{}
	}

	registry, err := c.Get(ctx, r.Name)
	if err != nil {
		return err
	}

	if r.ID != registry.ID {
		return &ErrRegistryMismatch{}
	}

	_, err = c.parent.Client.Products.DeleteRegistriesID(
		&products.DeleteRegistriesIDParams{
			ID:      registry.ID,
			Context: ctx,
		}, c.parent.AuthInfo)

	return handleSwaggerRegistryErrors(err)
}
