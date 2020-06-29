package goharborclient

import (
	"context"
	"github.com/go-openapi/runtime"
	"github.com/mittwald/goharbor-client/internal/api/v1.10.0/client"
	"github.com/mittwald/goharbor-client/internal/api/v1.10.0/model"
)

func ExampleNewRESTClient() {
	var h *client.Harbor
	var a runtime.ClientAuthInfoWriter

	// Construct an example client
	cl := NewRESTClient(h, a)

	ctx := context.Background()

	// Create an example project
	project, err := cl.project.NewProject(ctx, "example-project", 100, 0)
	if err != nil {
		panic(err)
	}

	// Create an example user
	usr, err := cl.user.NewUser(ctx, "example-user", "test@example.com", "Test User", "Sup3rS3cr3t!", "example comment")
	if err != nil {
		panic(err)
	}

	err = cl.project.AddProjectMember(ctx, project, usr, 1)
	if err != nil {
		panic(err)
	}

	// Registry Credentials using basic auth information
	registryCredentials := &model.RegistryCredential{
		AccessKey:    "admin",
		AccessSecret: "Sup3rS3cr3t!",
		Type:         "basic",
	}

	// Create an example registry
	reg, err := cl.registry.NewRegistry(ctx, "example-registry", "harbor", "demo.goharbor.io", registryCredentials, false)
	if err != nil {
		panic(err)
	}

	// Replication filter values
	replicationFilters := []*model.ReplicationFilter{
		{
			Type:  "name",
			Value: "alpine",
		},
		{
			Type:  "tag",
			Value: "latest",
		},
	}

	// Replication trigger defining an hourly interval
	replicationTrigger := &model.ReplicationTrigger{
		TriggerSettings: &model.TriggerSettings{Cron: "0 * * * *"},
		Type:            "scheduled",
	}

	// Create an example replication using the registry as source registry
	_, err = cl.replication.NewReplication(ctx, nil, reg, true, true, true,
		replicationFilters, replicationTrigger, "", "", "example-replication")
	if err != nil {
		panic(err)
	}
}
