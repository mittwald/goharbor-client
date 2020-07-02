// +build integration

package replication

import (
	"context"
	"flag"
	"testing"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/mittwald/goharbor-client/internal/api/v1_10_0/client"
	"github.com/mittwald/goharbor-client/registry"

	model "github.com/mittwald/goharbor-client/model/v1_10_0"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	host     = "localhost:30002"
	user     = "admin"
	password = "Harbor12345"
)

var (
	swaggerClient = client.New(runtimeclient.New(host, "/api", []string{"http"}), strfmt.Default)
	authInfo      = runtimeclient.BasicAuth(user, password)
	harborVersion = flag.String("version", "1.10.2",
		"Harbor version, used in conjunction with -integration, "+
			"defaults to 1.10.2")
	skipSpinUp = flag.Bool("skip-spinup", false,
		"Skip kind cluster creation")
)

func TestAPIReplicationNewDestRegistry(t *testing.T) {
	name := "test-project"

	ctx := context.Background()
	c := NewClient(swaggerClient, authInfo)

	regName := "test-registry"
	registryType := "harbor"
	url := "http://registry-docker-registry:5000/"
	credential := model.RegistryCredential{
		AccessKey:    user,
		AccessSecret: password,
		Type:         "basic",
	}

	var filters []*model.ReplicationFilter
	trigger := &model.ReplicationTrigger{
		TriggerSettings: &model.TriggerSettings{
			Cron: "",
		},
		Type: "manual",
	}

	rc := registry.NewClient(c.Client, c.AuthInfo)

	reg, err := rc.NewRegistry(ctx, regName, registryType, url, &credential, false)
	require.NoError(t, err)
	defer rc.DeleteRegistry(ctx, reg)

	r, err := c.NewReplication(
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
	defer c.DeleteReplication(ctx, r)

	assert.Equal(t, name, r.Name)
}

func TestAPIReplicationNewSrcRegistry(t *testing.T) {

	name := "test-project"

	ctx := context.Background()
	c := NewClient(swaggerClient, authInfo)

	regName := "test-registry"
	registryType := "harbor"
	url := "http://registry-docker-registry:5000/"
	credential := model.RegistryCredential{
		AccessKey:    user,
		AccessSecret: password,
		Type:         "basic",
	}

	var filters []*model.ReplicationFilter
	trigger := &model.ReplicationTrigger{
		TriggerSettings: &model.TriggerSettings{
			Cron: "",
		},
		Type: "manual",
	}

	rc := registry.NewClient(c.Client, c.AuthInfo)

	reg, err := rc.NewRegistry(ctx, regName, registryType, url, &credential, false)
	require.NoError(t, err)
	defer rc.DeleteRegistry(ctx, reg)

	r, err := c.NewReplication(
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
	defer c.DeleteReplication(ctx, r)

	assert.Equal(t, name, r.Name)
}

func TestAPIReplicationDelete(t *testing.T) {

	name := "test-project"

	ctx := context.Background()
	c := NewClient(swaggerClient, authInfo)

	regName := "test-registry"
	registryType := "harbor"
	url := "http://registry-docker-registry:5000/"
	credential := model.RegistryCredential{
		AccessKey:    user,
		AccessSecret: password,
		Type:         "basic",
	}

	var filters []*model.ReplicationFilter
	trigger := &model.ReplicationTrigger{
		TriggerSettings: &model.TriggerSettings{
			Cron: "",
		},
		Type: "manual",
	}

	rc := registry.NewClient(c.Client, c.AuthInfo)

	reg, err := rc.NewRegistry(ctx, regName, registryType, url, &credential, false)
	require.NoError(t, err)

	r, err := c.NewReplication(
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

	err = c.DeleteReplication(ctx, r)
	require.NoError(t, err)

	r, err = c.GetReplication(ctx, name)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationNotFound{}, err)
	}
}

func TestAPIReplicationUpdate(t *testing.T) {

	name := "test-project"

	ctx := context.Background()
	c := NewClient(swaggerClient, authInfo)

	regName := "test-registry"
	registryType := "harbor"
	url := "http://registry-docker-registry:5000/"
	credential := model.RegistryCredential{
		AccessKey:    user,
		AccessSecret: password,
		Type:         "basic",
	}

	var filters []*model.ReplicationFilter
	trigger := &model.ReplicationTrigger{
		TriggerSettings: &model.TriggerSettings{
			Cron: "",
		},
		Type: "manual",
	}

	rc := registry.NewClient(c.Client, c.AuthInfo)

	reg, err := rc.NewRegistry(ctx, regName, registryType, url, &credential, false)
	require.NoError(t, err)
	defer rc.DeleteRegistry(ctx, reg)

	r, err := c.NewReplication(
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
	defer c.DeleteReplication(ctx, r)

	descBefore := r.Description

	r.Description = "b"

	err = c.UpdateReplication(ctx, r)
	require.NoError(t, err)

	r, err = c.GetReplication(ctx, name)
	assert.NoError(t, err)

	assert.NotEqual(t, descBefore, r.Description)
}
