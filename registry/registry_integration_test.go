// +build integration

package registry

import (
	"context"
	"flag"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	model "github.com/mittwald/goharbor-client/model/v1_10_0"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	host     = "http://localhost:30002/api"
	user     = "admin"
	password = "Harbor12345"
)

var (
	swaggerClient = goharborclient.NewRESTClientForHost(host, user, password)
	authInfo      = runtimeclient.BasicAuth(user, password)
	harborVersion = flag.String("version", "1.10.2",
		"Harbor version, used in conjunction with -integration, "+
			"defaults to 1.10.2")
	skipSpinUp = flag.Bool("skip-spinup", false,
		"Skip kind cluster creation")
)

func TestAPIRegistryNew(t *testing.T) {
	name := "test-registry"
	registryType := "harbor"
	url := "http://registry-docker-registry:5000/"
	credential := model.RegistryCredential{
		AccessKey:    user,
		AccessSecret: password,
		Type:         "basic",
	}

	ctx := context.Background()
	c := NewClient(swaggerClient, authInfo)

	r, err := c.NewRegistry(ctx, name, registryType, url, &credential, false)
	defer c.DeleteRegistry(ctx, r)

	require.NoError(t, err)
	assert.Equal(t, name, r.Name)
}

func TestAPIRegistryGet(t *testing.T) {
	name := "test-registry"
	registryType := "harbor"
	url := "http://registry-docker-registry:5000/"
	credential := model.RegistryCredential{
		AccessKey:    user,
		AccessSecret: password,
		Type:         "basic",
	}

	ctx := context.Background()
	c := NewClient(swaggerClient, authInfo)

	r, err := c.NewRegistry(ctx, name, registryType, url, &credential, false)
	require.NoError(t, err)
	defer c.DeleteRegistry(ctx, r)

	p2, err := c.GetRegistry(ctx, name)
	require.NoError(t, err)
	assert.Equal(t, r, p2)
}

func TestAPIRegistryDelete(t *testing.T) {
	name := "test-registry"
	registryType := "harbor"
	url := "http://registry-docker-registry:5000/"
	credential := model.RegistryCredential{
		AccessKey:    user,
		AccessSecret: password,
		Type:         "basic",
	}

	ctx := context.Background()
	c := NewClient(swaggerClient, authInfo)

	r, err := c.NewRegistry(ctx, name, registryType, url, &credential, false)
	require.NoError(t, err)

	err = c.DeleteRegistry(ctx, r)
	require.NoError(t, err)

	r, err = c.GetRegistry(ctx, name)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrRegistryNotFound{}, err)
	}
}
