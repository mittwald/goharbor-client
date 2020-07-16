package registry

import (
	"context"

	"github.com/go-openapi/runtime"
	"github.com/mittwald/goharbor-client/internal/api/v1_10_0/client"

	"github.com/mittwald/goharbor-client/internal/api/v1_10_0/client/products"
	model "github.com/mittwald/goharbor-client/model/v1_10_0"
)

// RESTClient is a subclient for handling registry related actions.
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
	NewRegistry(ctx context.Context, name, registryType, url string,
		credential *model.RegistryCredential, insecure bool) (*model.Registry, error)
	GetRegistry(ctx context.Context, name string) (*model.Registry, error)
	DeleteRegistry(ctx context.Context, r *model.Registry) error
	UpdateRegistry(ctx context.Context, r *model.Registry) error
}

// NewRegistry creates a new project with name as project name.
// CountLimit and StorageLimit limits space and access for this project.
// Returns the registry as it is stored inside Harbor or an error,
// if it cannot be created.
func (c *RESTClient) NewRegistry(ctx context.Context, name, registryType, url string,
	credential *model.RegistryCredential, insecure bool) (*model.Registry, error) {
	rReq := &model.Registry{
		Credential: credential,
		Insecure:   insecure,
		Name:       name,
		Type:       registryType,
		URL:        url,
	}

	_, err := c.Client.Products.PostRegistries(
		&products.PostRegistriesParams{
			Registry: rReq,
			Context:  ctx,
		}, c.AuthInfo)

	err = handleSwaggerRegistryErrors(err)
	if err != nil {
		return nil, err
	}

	registry, err := c.GetRegistry(ctx, name)
	if err != nil {
		return nil, err
	}

	return registry, nil
}

// Get returns a registry identified by name.
// Returns an error if it cannot find a matching registry or when
// having difficulties talking to the API.
func (c *RESTClient) GetRegistry(ctx context.Context, name string) (*model.Registry, error) {
	if name == "" {
		return nil, &ErrRegistryNotProvided{}
	}
	resp, err := c.Client.Products.GetRegistries(
		&products.GetRegistriesParams{
			Name:    &name,
			Context: ctx,
		}, c.AuthInfo)

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
func (c *RESTClient) DeleteRegistry(ctx context.Context,
	r *model.Registry) error {
	if r == nil {
		return &ErrRegistryNotProvided{}
	}

	registry, err := c.GetRegistry(ctx, r.Name)
	if err != nil {
		return err
	}

	if r.ID != registry.ID {
		return &ErrRegistryMismatch{}
	}

	_, err = c.Client.Products.DeleteRegistriesID(
		&products.DeleteRegistriesIDParams{
			ID:      registry.ID,
			Context: ctx,
		}, c.AuthInfo)

	return handleSwaggerRegistryErrors(err)
}

func (c *RESTClient) UpdateRegistry(ctx context.Context, r *model.Registry) error {
	if r == nil {
		return &ErrRegistryNotProvided{}
	}

	rReq := &model.PutRegistry{
		AccessKey:      r.Credential.AccessKey,
		AccessSecret:   r.Credential.AccessSecret,
		CredentialType: r.Credential.Type,
		Description:    r.Description,
		Insecure:       r.Insecure,
		Name:           r.Name,
		URL:            r.URL,
	}

	registry, err := c.GetRegistry(ctx, r.Name)
	if err != nil {
		return err
	}

	if r.ID != registry.ID {
		return &ErrRegistryMismatch{}
	}

	_, err = c.Client.Products.PutRegistriesID(
		&products.PutRegistriesIDParams{
			ID:         registry.ID,
			RepoTarget: rReq,
			Context:    ctx,
		}, c.AuthInfo)

	return handleSwaggerRegistryErrors(err)
}
