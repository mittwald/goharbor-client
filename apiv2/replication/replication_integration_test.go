// +build integration

package replication

import (
	"context"
	"flag"
	"net/url"
	"testing"

	"github.com/go-openapi/strfmt"
	v2client "github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/legacyapi/client"
	integrationtest "github.com/mittwald/goharbor-client/v3/apiv2/testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/mittwald/goharbor-client/v3/apiv2/registry"

	model "github.com/mittwald/goharbor-client/v3/apiv2/model/legacy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	u, _                = url.Parse(integrationtest.Host)
	legacySwaggerClient = client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	v2SwaggerClient     = v2client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	authInfo            = runtimeclient.BasicAuth(integrationtest.User, integrationtest.Password)
	harborVersion       = flag.String("version", "2.1.1",
		"Harbor version, used in conjunction with -integration, "+
			"defaults to 2.1.1")
	skipSpinUp = flag.Bool("skip-spinup", false,
		"Skip kind cluster creation")
)

func TestAPIReplicationNewDestRegistry(t *testing.T) {
	name := "test-project"

	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	regName := "test-registry"
	registryType := "harbor"
	url := "http://registry-docker-registry:5000/"
	credential := model.RegistryCredential{
		AccessKey:    integrationtest.User,
		AccessSecret: integrationtest.Password,
		Type:         "basic",
	}

	var filters []*model.ReplicationFilter
	trigger := &model.ReplicationTrigger{
		TriggerSettings: &model.TriggerSettings{
			Cron: "",
		},
		Type: "manual",
	}

	rc := registry.NewClient(c.LegacyClient, c.V2Client, c.AuthInfo)

	reg, err := rc.NewRegistry(ctx, regName, registryType, url, &credential, false)
	require.NoError(t, err)
	defer rc.DeleteRegistry(ctx, reg)

	r, err := c.NewReplicationPolicy(

		ctx,
		reg,
		nil,
		true,
		true,
		true,
		filters,
		trigger,
		"", "", name,
	)
	require.NoError(t, err)
	defer c.DeleteReplicationPolicy(ctx, r)

	assert.Equal(t, name, r.Name)
}

func TestAPIReplicationNewSrcRegistry(t *testing.T) {
	name := "test-project"

	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	regName := "test-registry"
	registryType := "harbor"
	url := "http://registry-docker-registry:5000/"
	credential := model.RegistryCredential{
		AccessKey:    integrationtest.User,
		AccessSecret: integrationtest.Password,
		Type:         "basic",
	}

	var filters []*model.ReplicationFilter
	trigger := &model.ReplicationTrigger{
		TriggerSettings: &model.TriggerSettings{
			Cron: "",
		},
		Type: "manual",
	}

	rc := registry.NewClient(c.LegacyClient, c.V2Client, c.AuthInfo)

	reg, err := rc.NewRegistry(ctx, regName, registryType, url, &credential, false)
	require.NoError(t, err)
	defer rc.DeleteRegistry(ctx, reg)

	r, err := c.NewReplicationPolicy(
		ctx,
		nil,
		reg,
		true,
		true,
		true,
		filters,
		trigger,
		"", "", name,
	)
	require.NoError(t, err)
	defer c.DeleteReplicationPolicy(ctx, r)

	assert.Equal(t, name, r.Name)
}

func TestAPIReplicationDelete(t *testing.T) {
	name := "test-project"

	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	regName := "test-registry"
	registryType := "harbor"
	url := "http://registry-docker-registry:5000/"
	credential := model.RegistryCredential{
		AccessKey:    integrationtest.User,
		AccessSecret: integrationtest.Password,
		Type:         "basic",
	}

	var filters []*model.ReplicationFilter
	trigger := &model.ReplicationTrigger{
		TriggerSettings: &model.TriggerSettings{
			Cron: "",
		},
		Type: "manual",
	}

	rc := registry.NewClient(c.LegacyClient, c.V2Client, c.AuthInfo)

	reg, err := rc.NewRegistry(ctx, regName, registryType, url, &credential, false)
	require.NoError(t, err)

	r, err := c.NewReplicationPolicy(
		ctx,
		nil,
		reg,
		true,
		true,
		true,
		filters,
		trigger,
		"", "", name,
	)
	require.NoError(t, err)
	defer rc.DeleteRegistry(ctx, reg)

	err = c.DeleteReplicationPolicy(ctx, r)
	require.NoError(t, err)

	r, err = c.GetReplicationPolicy(ctx, name)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationNotFound{}, err)
	}
}

func TestAPIReplicationUpdate(t *testing.T) {
	name := "test-project"

	ctx := context.Background()
	c := NewClient(legacySwaggerClient, v2SwaggerClient, authInfo)

	regName := "test-registry"
	registryType := "harbor"
	url := "http://registry-docker-registry:5000/"
	credential := model.RegistryCredential{
		AccessKey:    integrationtest.User,
		AccessSecret: integrationtest.Password,
		Type:         "basic",
	}

	var filters []*model.ReplicationFilter
	trigger := &model.ReplicationTrigger{
		TriggerSettings: &model.TriggerSettings{
			Cron: "",
		},
		Type: "manual",
	}

	rc := registry.NewClient(c.LegacyClient, c.V2Client, c.AuthInfo)

	reg, err := rc.NewRegistry(ctx, regName, registryType, url, &credential, false)
	require.NoError(t, err)
	defer rc.DeleteRegistry(ctx, reg)

	r, err := c.NewReplicationPolicy(
		ctx,
		nil,
		reg,
		true,
		true,
		true,
		filters,
		trigger,
		"", "a", name,
	)

	require.NoError(t, err)
	defer c.DeleteReplicationPolicy(ctx, r)

	descBefore := r.Description

	r.Description = "b"

	err = c.UpdateReplicationPolicy(ctx, r)
	require.NoError(t, err)

	r, err = c.GetReplicationPolicy(ctx, name)
	assert.NoError(t, err)

	assert.NotEqual(t, descBefore, r.Description)
}
