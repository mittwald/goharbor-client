package goharborclient

import (
	"context"
	"errors"
	"fmt"

	"github.com/mittwald/goharbor-client/api/v1.10.0/client/products"
	"github.com/mittwald/goharbor-client/api/v1.10.0/model"
)

// RegistryRESTClient is a subclient for RESTClient handling registry related
// actions.
type RegistryRESTClient struct {
	parent *RESTClient
}

// RegistryError is an error describing a errors related to registry operations
// and implements the error interface.
type RegistryError struct {
	// ID of the related registry. -1 means undefined.
	RegistryID int64

	// Name of the related registry. Empty string means undefined.
	RegistryName string

	// Error message of the related registry.
	errorMessage string
}

// Error implements the Error interface.
func (r *RegistryError) Error() string {
	return fmt.Sprintf("%s (registry: %s, id: %d)",
		r.errorMessage, r.RegistryName, r.RegistryID)
}

// NewRegistryError creates a new RegistryError.
func NewRegistryError(msg string, id int64, name string) error {
	return &RegistryError{
		RegistryID:   id,
		RegistryName: name,
		errorMessage: msg,
	}
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

// Get returns a project identified by name.
// Returns an error if it cannot find a matching project or when
// having difficulties talking to the API.
func (c *RegistryRESTClient) Get(ctx context.Context, name string) (*model.Registry, error) {
	if name == "" {
		return nil, errors.New("no name provided")
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

	return nil, NewRegistryError(ErrRegistryNotFoundMsg, -1, name)
}

// Delete deletes a registry.
// Returns an error when no matching registry is found or when
// having difficulties talking to the API.
func (c *RegistryRESTClient) Delete(ctx context.Context,
	r *model.Registry) error {
	if r == nil {
		return errors.New("no registry provided")
	}

	registry, err := c.Get(ctx, r.Name)
	if err != nil {
		return err
	}

	if r.ID != registry.ID {
		return NewRegistryError(ErrRegistryMismatchMsg, r.ID, r.Name)
	}

	_, err = c.parent.Client.Products.DeleteRegistriesID(
		&products.DeleteRegistriesIDParams{
			ID:      registry.ID,
			Context: ctx,
		}, c.parent.AuthInfo)

	return handleSwaggerRegistryErrors(err)
}
