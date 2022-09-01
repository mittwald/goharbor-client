//go:build integration

package registry

import (
	"context"
	"fmt"
	"testing"

	modelv2 "github.com/testwill/goharbor-client/v5/apiv2/model"

	clienttesting "github.com/testwill/goharbor-client/v5/apiv2/pkg/testing"

	"github.com/stretchr/testify/require"
)

var (
	credential = modelv2.RegistryCredential{
		AccessKey:    clienttesting.User,
		AccessSecret: clienttesting.Password,
		Type:         "basic",
	}
	name         = "test-registry"
	registryType = "harbor"
	registryURL  = "http://registry-docker-registry:5000/"
)

func TestAPIRegistryNew(t *testing.T) {
	ctx := context.Background()

	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

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

	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

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

	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

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

	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

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
