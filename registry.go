package goharborclient

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-openapi/runtime"
	"github.com/mittwald/goharbor-client/api/v1.10.0/client/products"
	"github.com/mittwald/goharbor-client/api/v1.10.0/model"
)

const (
	// ErrRegistryIllegalIDFormat describes an illegal request format
	ErrRegistryIllegalIDFormat = "unsatisfied with constraints of the registry creation"

	// ErrRegistryUnauthorized describes an unauthorized request
	ErrRegistryUnauthorized = "unauthorized"

	// ErrRegistryInternalErrors describes server-side internal errors
	ErrRegistryInternalErrors = "unexpected internal errors"

	// ErrRegistryNoPermission describes a request error without permission
	ErrRegistryNoPermission = "user does not have permission to the registry"

	// ErrRegistryIDNotExists describes an error
	// when no proper registry ID is found
	ErrRegistryIDNotExists = "registry ID does not exist"

	// ErrRegistryNameAlreadyExists describes a duplicate project name error
	ErrRegistryNameAlreadyExists = "registry name already exists"

	// ErrRegistryMismatch describes a failed lookup
	// of a registry with name/id pair
	ErrRegistryMismatch = "id/name pair not found on server side"

	// ErrRegistryNotFound describes an error
	// when a specific project is not found
	ErrRegistryNotFound = "registry not found on server side"
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

	err = handleSwaggerRegistryErrors(err, -1, name)
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

	err = handleSwaggerRegistryErrors(err, -1, name)
	if err != nil {
		return nil, err
	}

	for _, r := range resp.Payload {
		if r.Name == name {
			return r, nil
		}
	}

	return nil, NewRegistryError(ErrRegistryNotFound, -1, name)
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
		return NewRegistryError(ErrRegistryMismatch, r.ID, r.Name)
	}

	_, err = c.parent.Client.Products.DeleteRegistriesID(
		&products.DeleteRegistriesIDParams{
			ID:      registry.ID,
			Context: ctx,
		}, c.parent.AuthInfo)

	return handleSwaggerRegistryErrors(err, r.ID, r.Name)
}

// handleSwaggerRegistryErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with a proper message.
func handleSwaggerRegistryErrors(in error, id int64, name string) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case 400:
			return NewRegistryError(ErrRegistryIllegalIDFormat, id, name)
		case 401:
			return NewRegistryError(ErrRegistryUnauthorized, id, name)
		case 403:
			return NewRegistryError(ErrRegistryNoPermission, id, name)
		case 500:
			return NewRegistryError(ErrRegistryInternalErrors, id, name)
		}
	}

	switch in.(type) {
	case *products.DeleteRegistriesIDNotFound:
		return NewRegistryError(ErrRegistryIDNotExists, id, name)
	case *products.PutRegistriesIDNotFound:
		return NewRegistryError(ErrRegistryIDNotExists, id, name)
	case *products.PostRegistriesConflict:
		return NewRegistryError(ErrRegistryNameAlreadyExists, id, name)
	default:
		return in
	}
}
