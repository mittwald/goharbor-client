//go:build integration

package replication

import (
	"context"
	"net/url"
	"testing"

	"github.com/go-openapi/strfmt"

	v2client "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v4/apiv2/internal/legacyapi/client"
	modelv2 "github.com/mittwald/goharbor-client/v4/apiv2/model"
	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/common"
	"github.com/mittwald/goharbor-client/v4/apiv2/pkg/config"
	integrationtest "github.com/mittwald/goharbor-client/v4/apiv2/pkg/testing"

	runtimeclient "github.com/go-openapi/runtime/client"

	"github.com/mittwald/goharbor-client/v4/apiv2/registry"

	"github.com/stretchr/testify/require"
)

var (
	u, _                = url.Parse(integrationtest.Host)
	legacySwaggerClient = client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	v2SwaggerClient     = v2client.New(runtimeclient.New(u.Host, u.Path, []string{u.Scheme}), strfmt.Default)
	authInfo            = runtimeclient.BasicAuth(integrationtest.User, integrationtest.Password)
	opts                = config.Options{}
	defaultOpts         = opts.Defaults()
)

func TestAPIReplicationNewDestRegistry(t *testing.T) {
	name := "test-project"

	ctx := context.Background()
	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)

	regName := "test-registry"
	registryType := "harbor"
	url := "http://registry-docker-registry:5000/"
	credential := modelv2.RegistryCredential{
		AccessKey:    integrationtest.User,
		AccessSecret: integrationtest.Password,
		Type:         "basic",
	}

	var filters []*modelv2.ReplicationFilter
	trigger := &modelv2.ReplicationTrigger{
		TriggerSettings: &modelv2.ReplicationTriggerSettings{
			Cron: "",
		},
		Type: "manual",
	}

	rc := registry.NewClient(c.V2Client, defaultOpts, c.AuthInfo)

	registry := &modelv2.Registry{
		Credential: &credential,
		Insecure:   false,
		Name:       regName,
		Type:       registryType,
		UpdateTime: strfmt.DateTime{},
		URL:        url,
	}

	err := rc.NewRegistry(ctx, registry)
	require.NoError(t, err)

	reg, err := rc.GetRegistryByName(ctx, regName)
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
	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)

	regName := "test-registry"
	registryType := "harbor"
	url := "http://registry-docker-registry:5000/"
	credential := modelv2.RegistryCredential{
		AccessKey:    integrationtest.User,
		AccessSecret: integrationtest.Password,
		Type:         "basic",
	}

	var filters []*modelv2.ReplicationFilter
	trigger := &modelv2.ReplicationTrigger{
		TriggerSettings: &modelv2.ReplicationTriggerSettings{
			Cron: "",
		},
		Type: "manual",
	}

	rc := registry.NewClient(c.V2Client, defaultOpts, c.AuthInfo)

	err := rc.NewRegistry(ctx, &modelv2.Registry{
		Credential: &credential,
		Insecure:   false,
		Name:       regName,
		Type:       registryType,
		URL:        url,
	})
	require.NoError(t, err)

	reg, err := rc.GetRegistryByName(ctx, regName)

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
	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)

	regName := "test-registry"
	registryType := "harbor"
	url := "http://registry-docker-registry:5000/"
	credential := modelv2.RegistryCredential{
		AccessKey:    integrationtest.User,
		AccessSecret: integrationtest.Password,
		Type:         "basic",
	}

	var filters []*modelv2.ReplicationFilter
	trigger := &modelv2.ReplicationTrigger{
		TriggerSettings: &modelv2.ReplicationTriggerSettings{
			Cron: "",
		},
		Type: "manual",
	}

	rc := registry.NewClient(c.V2Client, defaultOpts, c.AuthInfo)

	err := rc.NewRegistry(ctx, &modelv2.Registry{
		Credential: &credential,
		Insecure:   false,
		Name:       regName,
		Type:       registryType,
		URL:        url,
	})
	require.NoError(t, err)

	reg, err := rc.GetRegistryByName(ctx, regName)
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
	require.ErrorIs(t, err, &common.ErrNotFound{})
}

func TestAPIReplicationUpdate(t *testing.T) {
	name := "test-project"

	ctx := context.Background()
	c := NewClient(v2SwaggerClient, defaultOpts, authInfo)

	regName := "test-registry"
	registryType := "harbor"
	url := "http://registry-docker-registry:5000/"
	credential := modelv2.RegistryCredential{
		AccessKey:    integrationtest.User,
		AccessSecret: integrationtest.Password,
		Type:         "basic",
	}

	var filters []*modelv2.ReplicationFilter
	trigger := &modelv2.ReplicationTrigger{
		TriggerSettings: &modelv2.ReplicationTriggerSettings{
			Cron: "",
		},
		Type: "manual",
	}

	rc := registry.NewClient(c.V2Client, defaultOpts, c.AuthInfo)

	err := rc.NewRegistry(ctx, &modelv2.Registry{
		Credential: &credential,
		Insecure:   false,
		Name:       regName,
		Type:       registryType,
		URL:        url,
	})

	require.NoError(t, err)

	reg, err := rc.GetRegistryByName(ctx, regName)
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
