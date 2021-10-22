//go:build integration

package registry

import (
	"context"
	"fmt"
	"net/url"
	"testing"

	v2client "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client"
	modelv2 "github.com/mittwald/goharbor-client/v4/apiv2/model"
	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/config"
	integrationtest "github.com/mittwald/goharbor-client/v4/apiv2/pkg/testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/stretchr/testify/require"

	"github.com/mittwald/goharbor-client/v4/apiv2/internal/legacyapi/client"
)

var (
	u, _                = url.Parse(integrationtest.Host)
	legacySwaggerClient = client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	v2SwaggerClient     = v2client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	authInfo            = runtimeclient.BasicAuth(integrationtest.User, integrationtest.Password)
	opts                = config.Options{}
	defaultOpts         = opts.Defaults()
	credential          = modelv2.RegistryCredential{
		AccessKey:    integrationtest.User,
		AccessSecret: integrationtest.Password,
		Type:         "basic",
	}
	name         = "test-registry"
	registryType = "harbor"
	registryURL  = "http://registry-docker-registry:5000/"
)

func TestAPIRegistryNew(t *testing.T) {
	ctx := context.Background()

	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)

	err := c.NewRegistry(ctx, &modelv2.Registry{
		Credential: &credential,
		Insecure:   false,
		Name:       name,
		Type:       registryType,
		URL:        registryURL,
	})

	require.NoError(t, err)

	reg, err := c.GetRegistryByName(ctx, name)

	defer c.DeleteRegistryByID(ctx, reg.ID)

	require.NoError(t, err)
	require.Equal(t, name, reg.Name)
}

func TestAPIRegistryGet(t *testing.T) {
	ctx := context.Background()

	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)

	for i := 0; i < 10; i++ {
		err := c.NewRegistry(ctx, &modelv2.Registry{
			Credential: &credential,
			Insecure:   false,
			Name:       fmt.Sprintf(name+"%d", i),
			Type:       registryType,
			URL:        registryURL,
		})
		require.NoError(t, err)

		reg, _ := c.GetRegistryByName(ctx, fmt.Sprintf(name+"%d", i))

		defer c.DeleteRegistryByID(ctx, reg.ID)
	}
}

func TestAPIRegistryDelete(t *testing.T) {
	ctx := context.Background()

	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)

	err := c.NewRegistry(ctx, &modelv2.Registry{
		Credential: &credential,
		Insecure:   false,
		Name:       name,
		Type:       registryType,
		URL:        registryURL,
	})

	require.NoError(t, err)

	reg, err := c.GetRegistryByName(ctx, name)

	require.NoError(t, err)

	err = c.DeleteRegistryByID(ctx, reg.ID)

	require.NoError(t, err)
}

func TestAPIRegistryUpdate(t *testing.T) {
	ctx := context.Background()

	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)

	err := c.NewRegistry(ctx, &modelv2.Registry{
		Credential: &credential,
		Insecure:   false,
		Name:       name,
		Type:       registryType,
		URL:        registryURL,
	})

	require.NoError(t, err)

	reg, err := c.GetRegistryByName(ctx, name)

	require.NoError(t, err)

	desc := strPtr("test")
	ins := boolPtr(true)

	err = c.UpdateRegistry(ctx, &modelv2.RegistryUpdate{
		Description: desc,
		Insecure:    ins,
	}, reg.ID)

	require.NoError(t, err)

	reg, err = c.GetRegistryByName(ctx, name)

	defer c.DeleteRegistryByID(ctx, reg.ID)

	require.NoError(t, err)

	require.Equal(t, *desc, reg.Description)
	require.Equal(t, *ins, reg.Insecure)
}

func boolPtr(in bool) *bool {
	return &in
}

func strPtr(in string) *string {
	return &in
}
