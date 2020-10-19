package registry

import (
	"context"

	"github.com/go-openapi/runtime"
	v2client "github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/legacyapi/client"

	"github.com/mittwald/goharbor-client/v3/apiv2/internal/legacyapi/client/products"
	model "github.com/mittwald/goharbor-client/v3/apiv2/model/legacy"
)

// RESTClient is a subclient for handling registry related actions.
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

	_, err := c.LegacyClient.Products.PostRegistries(
		&products.PostRegistriesParams{
			Registry: rReq,
			Context:  ctx,
		}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerRegistryErrors(err)
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
	resp, err := c.LegacyClient.Products.GetRegistries(
		&products.GetRegistriesParams{
			Name:    &name,
			Context: ctx,
		}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerRegistryErrors(err)
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

	_, err = c.LegacyClient.Products.DeleteRegistriesID(
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

	_, err = c.LegacyClient.Products.PutRegistriesID(
		&products.PutRegistriesIDParams{
			ID:         registry.ID,
			RepoTarget: rReq,
			Context:    ctx,
		}, c.AuthInfo)

	return handleSwaggerRegistryErrors(err)
}
