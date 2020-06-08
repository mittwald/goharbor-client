package goharborclient

import (
	"context"
	"github.com/mittwald/goharbor-client/api/v1.10.0/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
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

