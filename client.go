package goharborclient

import (
	"context"
	"github.com/mittwald/goharbor-client/project"
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
}

type RESTClient struct {
	user    *user.RESTClient
	project *project.RESTClient
}

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
func (c *RESTClient) Registries() *RegistryRESTClient {
	return &RegistryRESTClient{
		parent: c,
	}
}

// Replications returns a project subclient for handling replication related actions.
func (c *RESTClient) Replications() *ReplicationRESTClient {
	return &ReplicationRESTClient{
		parent: c,
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
