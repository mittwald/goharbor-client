//go:build examples

package apiv2

import (
	"context"
	"github.com/go-openapi/runtime"
	"net/url"

	"github.com/mittwald/goharbor-client/v5/apiv2/model"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	v2client "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/config"
)

var (
	harborClient, _ = NewRESTClientForHost("", "", "", nil)
	ctx             = context.Background()
)

func ExampleNewRESTClientForHost() {
	// This example constructs a new goharbor-client for a goharbor instance
	// and create an example project.
	ctx := context.Background()
	apiURL := "harbor.mydomain.com/api"
	username := "user"
	password := "password"

	harborClient, err := NewRESTClientForHost(apiURL, username, password, nil)
	if err != nil {
		panic(err)
	}

	err = harborClient.NewProject(ctx, &model.ProjectReq{
		ProjectName: "my-project",
	})

	if err != nil {
		panic(err)
	}
}

func ExampleNewRESTClient() {
	// This example constructs a new (goharbor) REST client
	// and create an example project.
	ctx := context.Background()
	apiURL := "harbor.mydomain.com/api"
	username := "user"
	password := "password"

	harborURL, err := url.Parse(apiURL)
	if err != nil {
		panic(err)
	}

	v2SwaggerClient := v2client.New(runtimeclient.New(harborURL.Host, harborURL.Path, []string{harborURL.Scheme}), strfmt.Default)
	authInfo := runtimeclient.BasicAuth(username, password)

	harborClient := NewRESTClient(v2SwaggerClient, nil, authInfo)

	err = harborClient.NewProject(ctx, &model.ProjectReq{
		ProjectName: "my-project",
	})

	if err != nil {
		panic(err)
	}
}

func ExampleNewRESTClient_withOptions() {
	// This example constructs a new (goharbor) REST client using the provided 'options',
	// and lists all projects matching the 'options' configuration.
	ctx := context.Background()
	apiURL := "harbor.mydomain.com/api"
	username := "user"
	password := "password"

	harborURL, err := url.Parse(apiURL)
	if err != nil {
		panic(err)
	}

	v2SwaggerClient := v2client.New(runtimeclient.New(harborURL.Host, harborURL.Path, []string{harborURL.Scheme}), strfmt.Default)
	authInfo := runtimeclient.BasicAuth(username, password)

	options := config.Defaults().WithSort("-name")

	harborClient := NewRESTClient(v2SwaggerClient, options, authInfo)

	// List all projects containing 'test-' in their name.
	// options.Query with a value of 'name=~test-' might be used instead.
	projects, err := harborClient.ListProjects(ctx, "test-")
	if err != nil {
		panic(err)
	}

	for project := range projects {
		// your logic here
		_ = project
	}
}

func ExampleRESTClient_NewUser() {
	err := harborClient.NewUser(ctx, "test-user", "foo@example.com", "test user", "password", "a test user")
	if err != nil {
		panic(err)
	}
}

func ExampleNewRESTClientWithAuthFunc() {
	// This example constructs a new goharbor-client for a goharbor instance
	// and create an example project.
	ctx := context.Background()
	apiURL := "harbor.mydomain.com/api"
	tokenAuthFunc := func() runtime.ClientAuthInfoWriterFunc {
		return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
			return r.SetHeaderParam("Authorization", "Bearer 123456")
		})
	}
	harborClient, err := NewRESTClientWithAuthFunc(apiURL, tokenAuthFunc(), nil)
	if err != nil {
		panic(err)
	}

	err = harborClient.NewProject(ctx, &model.ProjectReq{
		ProjectName: "my-project",
	})

	if err != nil {
		panic(err)
	}
}
