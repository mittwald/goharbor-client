package goharborclient

import (
	"context"
	"testing"

	"github.com/mittwald/goharbor-client/api/v1.10.0/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAPIRegistryNew(t *testing.T) {
	if !*integrationTest {
		t.Skip()
	}

	name := "test-registry"
	registryType := "harbor"
	url := "http://registry-docker-registry:5000/"
	credential := model.RegistryCredential{
		AccessKey:    defaultUser,
		AccessSecret: defaultPassword,
		Type:         "basic",
	}

	ctx := context.Background()
	c := NewClient(host, defaultUser, defaultPassword)

	r, err := c.Registries().NewRegistry(ctx, name, registryType, url, &credential, false)
	defer c.Registries().Delete(ctx, r)

	require.NoError(t, err)
	assert.Equal(t, name, r.Name)
}

func TestAPIRegistryGet(t *testing.T) {
	if !*integrationTest {
		t.Skip()
	}

	name := "test-registry"
	registryType := "harbor"
	url := "http://registry-docker-registry:5000/"
	credential := model.RegistryCredential{
		AccessKey:    defaultUser,
		AccessSecret: defaultPassword,
		Type:         "basic",
	}

	ctx := context.Background()
	c := NewClient(host, defaultUser, defaultPassword)

	r, err := c.Registries().NewRegistry(ctx, name, registryType, url, &credential, false)
	require.NoError(t, err)
	defer c.Registries().Delete(ctx, r)

	p2, err := c.Registries().Get(ctx, name)
	require.NoError(t, err)
	assert.Equal(t, r, p2)
}

func TestAPIRegistryDelete(t *testing.T) {
	if !*integrationTest {
		t.Skip()
	}

	name := "test-registry"
	registryType := "harbor"
	url := "http://registry-docker-registry:5000/"
	credential := model.RegistryCredential{
		AccessKey:    defaultUser,
		AccessSecret: defaultPassword,
		Type:         "basic",
	}

	ctx := context.Background()
	c := NewClient(host, defaultUser, defaultPassword)

	r, err := c.Registries().NewRegistry(ctx, name, registryType, url, &credential, false)
	require.NoError(t, err)

	err = c.Registries().Delete(ctx, r)
	require.NoError(t, err)

	r, err = c.Registries().Get(ctx, name)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrRegistryNotFound{}, err)
	}
}
