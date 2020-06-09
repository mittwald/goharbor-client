package goharborclient

import (
	"context"
	"testing"

	"github.com/mittwald/goharbor-client/api/v1.10.0/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAPIReplicationNewDestRegistry(t *testing.T) {
	if !*integrationTest {
		t.Skip()
	}

	name := "test-project"

	ctx := context.Background()
	c := NewClient(host, defaultUser, defaultPassword)

	regName := "test-registry"
	registryType := "harbor"
	url := "http://registry-docker-registry:5000/"
	credential := model.RegistryCredential{
		AccessKey:    defaultUser,
		AccessSecret: defaultPassword,
		Type:         "basic",
	}

	var filters []*model.ReplicationFilter
	trigger := &model.ReplicationTrigger{
		TriggerSettings: &model.TriggerSettings{
			Cron: "",
		},
		Type: "manual",
	}

	reg, err := c.Registries().NewRegistry(ctx, regName, registryType, url, &credential, false)
	require.NoError(t, err)

	r, err := c.Replications().NewReplication(
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
	defer c.Registries().Delete(ctx, reg)
	defer c.Replications().Delete(ctx, r)

	require.NoError(t, err)
	assert.Equal(t, name, r.Name)
}

func TestAPIReplicationNewSrcRegistry(t *testing.T) {
	if !*integrationTest {
		t.Skip()
	}

	name := "test-project"

	ctx := context.Background()
	c := NewClient(host, defaultUser, defaultPassword)

	regName := "test-registry"
	registryType := "harbor"
	url := "http://registry-docker-registry:5000/"
	credential := model.RegistryCredential{
		AccessKey:    defaultUser,
		AccessSecret: defaultPassword,
		Type:         "basic",
	}

	var filters []*model.ReplicationFilter
	trigger := &model.ReplicationTrigger{
		TriggerSettings: &model.TriggerSettings{
			Cron: "",
		},
		Type: "manual",
	}

	reg, err := c.Registries().NewRegistry(ctx, regName, registryType, url, &credential, false)
	require.NoError(t, err)

	r, err := c.Replications().NewReplication(
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
	defer c.Registries().Delete(ctx, reg)
	defer c.Replications().Delete(ctx, r)

	require.NoError(t, err)
	assert.Equal(t, name, r.Name)
}

func TestAPIReplicationDelete(t *testing.T) {
	if !*integrationTest {
		t.Skip()
	}

	name := "test-project"

	ctx := context.Background()
	c := NewClient(host, defaultUser, defaultPassword)

	regName := "test-registry"
	registryType := "harbor"
	url := "http://registry-docker-registry:5000/"
	credential := model.RegistryCredential{
		AccessKey:    defaultUser,
		AccessSecret: defaultPassword,
		Type:         "basic",
	}

	var filters []*model.ReplicationFilter
	trigger := &model.ReplicationTrigger{
		TriggerSettings: &model.TriggerSettings{
			Cron: "",
		},
		Type: "manual",
	}

	reg, err := c.Registries().NewRegistry(ctx, regName, registryType, url, &credential, false)
	require.NoError(t, err)

	r, err := c.Replications().NewReplication(
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

	err = c.Replications().Delete(ctx, r)
	require.NoError(t, err)

	r, err = c.Replications().Get(ctx, name)
	if assert.Error(t, err) {
		assert.IsType(t, &ErrReplicationNotFound{}, err)
	}

	defer c.Registries().Delete(ctx, reg)
}

func TestAPIReplicationUpdate(t *testing.T) {
	if !*integrationTest {
		t.Skip()
	}

	name := "test-project"

	ctx := context.Background()
	c := NewClient(host, defaultUser, defaultPassword)

	regName := "test-registry"
	registryType := "harbor"
	url := "http://registry-docker-registry:5000/"
	credential := model.RegistryCredential{
		AccessKey:    defaultUser,
		AccessSecret: defaultPassword,
		Type:         "basic",
	}

	var filters []*model.ReplicationFilter
	trigger := &model.ReplicationTrigger{
		TriggerSettings: &model.TriggerSettings{
			Cron: "",
		},
		Type: "manual",
	}

	reg, err := c.Registries().NewRegistry(ctx, regName, registryType, url, &credential, false)
	require.NoError(t, err)

	r, err := c.Replications().NewReplication(
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

	descBefore := r.Description

	r.Description = "b"

	err = c.Replications().Update(ctx, r)
	require.NoError(t, err)

	r, err = c.Replications().Get(ctx, name)
	assert.NoError(t, err)

	assert.NotEqual(t, descBefore, r.Description)

	defer c.Registries().Delete(ctx, reg)
	defer c.Replications().Delete(ctx, r)
}
