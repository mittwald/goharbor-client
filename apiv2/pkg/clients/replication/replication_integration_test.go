//go:build integration

package replication

import (
	"context"
	"testing"

	modelv2 "github.com/testwill/goharbor-client/v5/apiv2/model"
	"github.com/testwill/goharbor-client/v5/apiv2/pkg/errors"
	clienttesting "github.com/testwill/goharbor-client/v5/apiv2/pkg/testing"

	"github.com/testwill/goharbor-client/v5/apiv2/pkg/clients/registry"

	"github.com/stretchr/testify/require"
)

var exampleRegistry = &modelv2.Registry{
	Credential: &modelv2.RegistryCredential{
		AccessKey:    clienttesting.User,
		AccessSecret: clienttesting.Password,
		Type:         "basic",
	},
	Insecure: false,
	Name:     "test-registry",
	Type:     "harbor",
	URL:      "http://registry-docker-registry:5000/",
}

func TestAPIReplicationNewDestRegistry(t *testing.T) {
	name := "test-project"

	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	var filters []*modelv2.ReplicationFilter
	trigger := &modelv2.ReplicationTrigger{
		TriggerSettings: &modelv2.ReplicationTriggerSettings{
			Cron: "",
		},
		Type: "manual",
	}

	rc := registry.NewClient(c.V2Client, clienttesting.DefaultOpts, c.AuthInfo)

	err := rc.NewRegistry(ctx, exampleRegistry)
	require.NoError(t, err)

	reg, err := rc.GetRegistryByName(ctx, exampleRegistry.Name)
	require.NoError(t, err)

	defer rc.DeleteRegistryByID(ctx, reg.ID)

	err = c.NewReplicationPolicy(
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

	rep, err := c.GetReplicationPolicyByName(ctx, name)

	defer c.DeleteReplicationPolicyByID(ctx, rep.ID)

	require.Equal(t, name, rep.Name)
}

func TestAPIReplicationNewSrcRegistry(t *testing.T) {
	name := "test-project"

	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	var filters []*modelv2.ReplicationFilter
	trigger := &modelv2.ReplicationTrigger{
		TriggerSettings: &modelv2.ReplicationTriggerSettings{
			Cron: "",
		},
		Type: "manual",
	}

	rc := registry.NewClient(c.V2Client, clienttesting.DefaultOpts, c.AuthInfo)

	err := rc.NewRegistry(ctx, exampleRegistry)
	require.NoError(t, err)

	reg, err := rc.GetRegistryByName(ctx, exampleRegistry.Name)

	require.NoError(t, err)

	defer rc.DeleteRegistryByID(ctx, reg.ID)

	err = c.NewReplicationPolicy(
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

	rep, err := c.GetReplicationPolicyByName(ctx, name)
	require.NoError(t, err)

	defer c.DeleteReplicationPolicyByID(ctx, rep.ID)

	require.Equal(t, name, rep.Name)
}

func TestAPIReplicationDelete(t *testing.T) {
	name := "test-project"

	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	var filters []*modelv2.ReplicationFilter
	trigger := &modelv2.ReplicationTrigger{
		TriggerSettings: &modelv2.ReplicationTriggerSettings{
			Cron: "",
		},
		Type: "manual",
	}

	rc := registry.NewClient(c.V2Client, clienttesting.DefaultOpts, c.AuthInfo)

	err := rc.NewRegistry(ctx, exampleRegistry)
	require.NoError(t, err)

	reg, err := rc.GetRegistryByName(ctx, exampleRegistry.Name)
	require.NoError(t, err)

	defer rc.DeleteRegistryByID(ctx, reg.ID)

	err = c.NewReplicationPolicy(
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

	rep, err := c.GetReplicationPolicyByName(ctx, name)

	err = c.DeleteReplicationPolicyByID(ctx, rep.ID)
	require.NoError(t, err)

	_, err = c.GetReplicationPolicyByID(ctx, rep.ID)
	require.Error(t, err)
	require.ErrorIs(t, err, &errors.ErrNotFound{})
}

func TestAPIReplicationUpdate(t *testing.T) {
	name := "test-project"

	ctx := context.Background()
	c := NewClient(clienttesting.V2SwaggerClient, clienttesting.DefaultOpts, clienttesting.AuthInfo)

	var filters []*modelv2.ReplicationFilter
	trigger := &modelv2.ReplicationTrigger{
		TriggerSettings: &modelv2.ReplicationTriggerSettings{
			Cron: "",
		},
		Type: "manual",
	}

	rc := registry.NewClient(c.V2Client, clienttesting.DefaultOpts, c.AuthInfo)

	err := rc.NewRegistry(ctx, exampleRegistry)

	require.NoError(t, err)

	reg, err := rc.GetRegistryByName(ctx, exampleRegistry.Name)
	require.NoError(t, err)

	defer rc.DeleteRegistryByID(ctx, reg.ID)

	err = c.NewReplicationPolicy(
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

	rep, err := c.GetReplicationPolicyByName(ctx, name)

	defer c.DeleteReplicationPolicyByID(ctx, rep.ID)

	descBefore := rep.Description

	rep.Description = "b"

	err = c.UpdateReplicationPolicy(ctx, rep, rep.ID)
	require.NoError(t, err)

	rep, err = c.GetReplicationPolicyByName(ctx, name)
	require.NoError(t, err)

	require.NotEqual(t, descBefore, rep.Description)
}
