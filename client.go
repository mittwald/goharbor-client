package goharborclient

import (
	"context"
	"github.com/mittwald/goharbor-client/project"
	"github.com/mittwald/goharbor-client/registry"
	"github.com/mittwald/goharbor-client/replication"
	"github.com/mittwald/goharbor-client/user"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/mittwald/goharbor-client/api/v1.10.0/client"
	"github.com/mittwald/goharbor-client/api/v1.10.0/client/products"
	"github.com/mittwald/goharbor-client/api/v1.10.0/model"
)

type Client interface {
	user.Client
	project.Client
	registry.Client
	replication.Client
}

type RESTClient struct {
	user        *user.RESTClient
	project     *project.RESTClient
	registry    *registry.RESTClient
	replication *replication.RESTClient
}

// User Client
func (c *RESTClient) NewUser(ctx context.Context, username, email, realname, password, comments string) (*model.User, error) {
	return c.user.NewUser(ctx, username, email, realname, password, comments)
}

func (c *RESTClient) GetUser(ctx context.Context, username string) (*model.User, error) {
	return c.user.GetUser(ctx, username)
}
func (c *RESTClient) DeleteUser(ctx context.Context, u *model.User) error {
	return c.user.DeleteUser(ctx, u)
}
func (c *RESTClient) UpdateUser(ctx context.Context, u *model.User) error {
	return c.user.UpdateUser(ctx, u)
}

// Project Client
func (c *RESTClient) NewProject(ctx context.Context, name string, countLimit int, storageLimit int) (*model.Project, error) {
	return c.project.NewProject(ctx, name, countLimit, storageLimit)
}
func (c *RESTClient) DeleteProject(ctx context.Context, p *model.Project) error {
	return c.project.DeleteProject(ctx, p)
}
func (c *RESTClient) GetProject(ctx context.Context, name string) (*model.Project, error) {
	return c.project.GetProject(ctx, name)
}
func (c *RESTClient) ListProjects(ctx context.Context, nameFilter string) ([]*model.Project, error) {
	return c.project.ListProjects(ctx, nameFilter)
}
func (c *RESTClient) UpdateProject(ctx context.Context, p *model.Project, countLimit int, storageLimit int) error {
	return c.project.UpdateProject(ctx, p, countLimit, storageLimit)
}

func (c *RESTClient) AddProjectMember(ctx context.Context, p *model.Project, u *model.User, roleID int) error {
	return c.project.AddProjectMember(ctx, p, u, roleID)
}
func (c *RESTClient) ListProjectMembers(ctx context.Context, p *model.Project) ([]*model.ProjectMemberEntity, error) {
	return c.project.ListProjectMembers(ctx, p)
}
func (c *RESTClient) UpdateProjectMemberRole(ctx context.Context, p *model.Project, u *model.User, roleID int) error {
	return c.project.UpdateProjectMemberRole(ctx, p, u, roleID)
}
func (c *RESTClient) DeleteProjectMember(ctx context.Context, p *model.Project, u *model.User) error {
	return c.project.DeleteProjectMember(ctx, p, u)
}

func (c *RESTClient) AddProjectMetadata(ctx context.Context, p *model.Project, key project.ProjectMetadataKey, value string) error {
	return c.project.AddProjectMetadata(ctx, p, key, value)
}

func (c *RESTClient) ListProjectMetadata(ctx context.Context, p *model.Project) (*model.ProjectMetadata, error) {
	return c.project.ListProjectMetadata(ctx, p)

}
func (c *RESTClient) UpdateProjectMetadata(ctx context.Context, p *model.Project, key project.ProjectMetadataKey, value string) error {
	return c.project.UpdateProjectMetadata(ctx, p, key, value)
}
func (c *RESTClient) DeleteProjectMetadataValue(ctx context.Context, p *model.Project, key project.ProjectMetadataKey) error {
	return c.project.DeleteProjectMetadataValue(ctx, p, key)
}

// Registry Client
func (c *RESTClient) NewRegistry(ctx context.Context, name, registryType, url string,
	credential *model.RegistryCredential, insecure bool) (*model.Registry, error) {

	return c.registry.NewRegistry(ctx, name, registryType, url,
		credential, insecure)
}

func (c *RESTClient) GetRegistry(ctx context.Context, name string) (*model.Registry, error) {
	return c.registry.GetRegistry(ctx, name)

}

func (c *RESTClient) DeleteRegistry(ctx context.Context, r *model.Registry) error {
	return c.registry.DeleteRegistry(ctx, r)
}

// Replication Client
func (c *RESTClient) NewReplication(ctx context.Context, destRegistry, srcRegistry *model.Registry,
	replicateDeletion, override, enablePolicy bool, filters []*model.ReplicationFilter,
	trigger *model.ReplicationTrigger, destNamespace, description, name string) (*model.ReplicationPolicy, error) {

	return c.replication.NewReplication(ctx, destRegistry, srcRegistry, replicateDeletion,
		override, enablePolicy, filters, trigger, destNamespace, description, name)
}

func (c *RESTClient) GetReplication(ctx context.Context, name string) (*model.ReplicationPolicy, error) {
	return c.replication.GetReplication(ctx, name)
}

func (c *RESTClient) DeleteReplication(ctx context.Context, r *model.ReplicationPolicy) error {
	return c.replication.DeleteReplication(ctx, r)
}

func (c *RESTClient) UpdateReplication(ctx context.Context, r *model.ReplicationPolicy) error {
	return c.replication.UpdateReplication(ctx, r)
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

// Registries returns a project subclient for handling project related actions.
func (c *RESTClient) System() *SystemRESTClient {
	return &SystemRESTClient{
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
