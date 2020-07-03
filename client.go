package goharborclient

import (
	"context"

	"github.com/go-openapi/runtime"
	"github.com/mittwald/goharbor-client/internal/api/v1_10_0/client"

	"github.com/mittwald/goharbor-client/project"
	"github.com/mittwald/goharbor-client/registry"
	"github.com/mittwald/goharbor-client/replication"
	"github.com/mittwald/goharbor-client/system"
	"github.com/mittwald/goharbor-client/user"

	model "github.com/mittwald/goharbor-client/model/v1_10_0"
)

// Client is an interface that groups all sub-package methods.
type Client interface {
	user.Client
	project.Client
	registry.Client
	replication.Client
	system.Client
}

// RESTClient implements the Client interface as a REST client
type RESTClient struct {
	user        *user.RESTClient
	project     *project.RESTClient
	registry    *registry.RESTClient
	replication *replication.RESTClient
	system      *system.RESTClient
}

// NewRESTClient constructs a new REST client containing each sub client.
func NewRESTClient(cl *client.Harbor, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	c := &RESTClient{
		user:        user.NewClient(cl, authInfo),
		project:     project.NewClient(cl, authInfo),
		registry:    registry.NewClient(cl, authInfo),
		replication: replication.NewClient(cl, authInfo),
		system:      system.NewClient(cl, authInfo),
	}
	return c
}

// User Client

// NewUser wraps the NewUser method of the user sub-package.
func (c *RESTClient) NewUser(ctx context.Context, username, email, realname, password, comments string) (*model.User, error) {
	return c.user.NewUser(ctx, username, email, realname, password, comments)
}

// GetUser wraps the GetUser method of the user sub-package.
func (c *RESTClient) GetUser(ctx context.Context, username string) (*model.User, error) {
	return c.user.GetUser(ctx, username)
}

// DeleteUser wraps the DeleteUser method of the user sub-package.
func (c *RESTClient) DeleteUser(ctx context.Context, u *model.User) error {
	return c.user.DeleteUser(ctx, u)
}

// UpdateUser wraps the UpdateUser method of the user sub-package.
func (c *RESTClient) UpdateUser(ctx context.Context, u *model.User) error {
	return c.user.UpdateUser(ctx, u)
}

// Project Client

// NewProject wraps the NewProject method of the project sub-package.
func (c *RESTClient) NewProject(ctx context.Context, name string, countLimit int, storageLimit int) (*model.Project, error) {
	return c.project.NewProject(ctx, name, countLimit, storageLimit)
}

// DeleteProject wraps the DeleteProject method of the project sub-package.

func (c *RESTClient) DeleteProject(ctx context.Context, p *model.Project) error {
	return c.project.DeleteProject(ctx, p)
}

// GetProject wraps the GetProject method of the project sub-package.
func (c *RESTClient) GetProject(ctx context.Context, name string) (*model.Project, error) {
	return c.project.GetProject(ctx, name)
}

// ListProjects wraps the ListProjects method of the project sub-package.
func (c *RESTClient) ListProjects(ctx context.Context, nameFilter string) ([]*model.Project, error) {
	return c.project.ListProjects(ctx, nameFilter)
}

// UpdateProject wraps the UpdateProject method of the project sub-package.
func (c *RESTClient) UpdateProject(ctx context.Context, p *model.Project, countLimit int, storageLimit int) error {
	return c.project.UpdateProject(ctx, p, countLimit, storageLimit)
}

// AddProjectMember wraps the AddProjectMember method of the project sub-package.
func (c *RESTClient) AddProjectMember(ctx context.Context, p *model.Project, u *model.User, roleID int) error {
	return c.project.AddProjectMember(ctx, p, u, roleID)
}

// ListProjectMembers wraps the ListProjectMembers method of the project sub-package.
func (c *RESTClient) ListProjectMembers(ctx context.Context, p *model.Project) ([]*model.ProjectMemberEntity, error) {
	return c.project.ListProjectMembers(ctx, p)
}

// UpdateProjectMemberRole wraps the UpdateProjectMemberRole method of the project sub-package.
func (c *RESTClient) UpdateProjectMemberRole(ctx context.Context, p *model.Project, u *model.User, roleID int) error {
	return c.project.UpdateProjectMemberRole(ctx, p, u, roleID)
}

// DeleteProjectMember wraps the DeleteProjectMember method of the project sub-package.
func (c *RESTClient) DeleteProjectMember(ctx context.Context, p *model.Project, u *model.User) error {
	return c.project.DeleteProjectMember(ctx, p, u)
}

// AddProjectMetadata wraps the AddProjectMetadata method of the project sub-package.
func (c *RESTClient) AddProjectMetadata(ctx context.Context, p *model.Project, key project.ProjectMetadataKey, value string) error {
	return c.project.AddProjectMetadata(ctx, p, key, value)
}

// ListProjectMetadata wraps the ListProjectMetadata method of the project sub-package.
func (c *RESTClient) ListProjectMetadata(ctx context.Context, p *model.Project) (*model.ProjectMetadata, error) {
	return c.project.ListProjectMetadata(ctx, p)

}

// UpdateProjectMetadata wraps the UpdateProjectMetadata method of the project sub-package.
func (c *RESTClient) UpdateProjectMetadata(ctx context.Context, p *model.Project, key project.ProjectMetadataKey, value string) error {
	return c.project.UpdateProjectMetadata(ctx, p, key, value)
}

// DeleteProjectMetadataValue wraps the DeleteProjectMetadataValue method of the project sub-package.
func (c *RESTClient) DeleteProjectMetadataValue(ctx context.Context, p *model.Project, key project.ProjectMetadataKey) error {
	return c.project.DeleteProjectMetadataValue(ctx, p, key)
}

// Registry Client

// NewRegistry wraps the NewRegistry method of the registry sub-package.
func (c *RESTClient) NewRegistry(ctx context.Context, name, registryType, url string,
	credential *model.RegistryCredential, insecure bool) (*model.Registry, error) {
	return c.registry.NewRegistry(ctx, name, registryType, url,
		credential, insecure)
}

// GetRegistry wraps the GetRegistry method of the registry sub-package.
func (c *RESTClient) GetRegistry(ctx context.Context, name string) (*model.Registry, error) {
	return c.registry.GetRegistry(ctx, name)

}

// DeleteRegistry wraps the DeleteRegistry method of the registry sub-package.
func (c *RESTClient) DeleteRegistry(ctx context.Context, r *model.Registry) error {
	return c.registry.DeleteRegistry(ctx, r)
}

// Replication Client

// NewReplication wraps the NewReplication method of the replication sub-package.
func (c *RESTClient) NewReplication(ctx context.Context, destRegistry, srcRegistry *model.Registry,
	replicateDeletion, override, enablePolicy bool, filters []*model.ReplicationFilter,
	trigger *model.ReplicationTrigger, destNamespace, description, name string) (*model.ReplicationPolicy, error) {

	return c.replication.NewReplication(ctx, destRegistry, srcRegistry, replicateDeletion,
		override, enablePolicy, filters, trigger, destNamespace, description, name)
}

// GetReplication wraps the GetReplication method of the replication sub-package.
func (c *RESTClient) GetReplication(ctx context.Context, name string) (*model.ReplicationPolicy, error) {
	return c.replication.GetReplication(ctx, name)
}

// DeleteReplication wraps the DeleteReplication method of the replication sub-package.
func (c *RESTClient) DeleteReplication(ctx context.Context, r *model.ReplicationPolicy) error {
	return c.replication.DeleteReplication(ctx, r)
}

// UpdateReplication wraps the UpdateReplication method of the replication sub-package.
func (c *RESTClient) UpdateReplication(ctx context.Context, r *model.ReplicationPolicy) error {
	return c.replication.UpdateReplication(ctx, r)
}

// System Client

// NewSystemGarbageCollection wraps the NewSystemGarbageCollection method of the system sub-package.
func (c *RESTClient) NewSystemGarbageCollection(ctx context.Context, cron, scheduleType string) (*model.AdminJobSchedule, error) {
	return c.system.NewSystemGarbageCollection(ctx, cron, scheduleType)
}

// UpdateSystemGarbageCollection wraps the UpdateSystemGarbageCollection method of the system sub-package.
func (c *RESTClient) UpdateSystemGarbageCollection(ctx context.Context, newGcSchedule *model.AdminJobScheduleObj) error {
	return c.system.UpdateSystemGarbageCollection(ctx, newGcSchedule)
}

// GetSystemGarbageCollection wraps the GetSystemGarbageCollection method of the system sub-package.
func (c *RESTClient) GetSystemGarbageCollection(ctx context.Context) (*model.AdminJobSchedule, error) {
	return c.system.GetSystemGarbageCollection(ctx)
}

// ResetSystemGarbageCollection wraps the ResetSystemGarbageCollection method of the system sub-package.
func (c *RESTClient) ResetSystemGarbageCollection(ctx context.Context) error {
	return c.system.ResetSystemGarbageCollection(ctx)
}
