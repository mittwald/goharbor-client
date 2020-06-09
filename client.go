package goharborclient

import (
	"context"

	"github.com/go-openapi/runtime"
	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/mittwald/goharbor-client/api/v1.10.0/client"
	"github.com/mittwald/goharbor-client/api/v1.10.0/client/products"
	"github.com/mittwald/goharbor-client/api/v1.10.0/model"
)

// RESTClient implements a Harbor client capable of performing Harbor API
// calls using a swagger generated REST client under the hood.
type RESTClient struct {
	// The swagger client
	Client *client.Harbor

	// AuthInfo contain auth information, which are provided on API calls.
	AuthInfo runtime.ClientAuthInfoWriter
}

// NewClient creates a new Harbor client.
// host is the harbor hostname including the protocol scheme and port,
// i.e. "https://harbor.example.com" or "http://harbor.example.com:30002".
// user is the authentication username.
// password is the authentication password.
func NewClient(host, user, password string) *RESTClient {
	return &RESTClient{
		Client: client.New(runtimeclient.New(host,
			"/api", []string{"http"}), strfmt.Default),
		AuthInfo: runtimeclient.BasicAuth(user, password),
	}
}

// Projects returns a project subclient for handling project related actions.
func (c *RESTClient) Projects() *ProjectRESTClient {
	return &ProjectRESTClient{
		parent: c,
	}
}

// Users returns a user subclient for handling user related actions.
func (c *RESTClient) Users() *UserRESTClient {
	return &UserRESTClient{
		parent: c,
	}
}

// Registries returns a project subclient for handling project related actions.
func (c *RESTClient) Registries() *RegistryRESTClient {
	return &RegistryRESTClient{
		parent: c,
	}
}

// Health reports Harbor system health information.
func (c *RESTClient) Health(ctx context.Context) (*model.OverallHealthStatus, error) {
	resp, err := c.Client.Products.GetHealth(&products.GetHealthParams{
		Context: ctx,
	}, c.AuthInfo)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
