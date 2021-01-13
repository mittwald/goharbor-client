package apiv2

import (
	"context"
	"net/url"
	"strings"

	modelv2 "github.com/mittwald/goharbor-client/v3/apiv2/model"
	"github.com/mittwald/goharbor-client/v3/apiv2/retention"

	"github.com/go-openapi/runtime"
	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	v2client "github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/legacyapi/client"
	model "github.com/mittwald/goharbor-client/v3/apiv2/model/legacy"
	"github.com/mittwald/goharbor-client/v3/apiv2/project"
	"github.com/mittwald/goharbor-client/v3/apiv2/registry"
	"github.com/mittwald/goharbor-client/v3/apiv2/replication"
	"github.com/mittwald/goharbor-client/v3/apiv2/system"
	"github.com/mittwald/goharbor-client/v3/apiv2/user"
)

const v2URLSuffix string = "/v2.0"

type Client interface {
	user.Client
	project.Client
	registry.Client
	replication.Client
	system.Client
	retention.Client
}

// RESTClient implements the Client interface as a REST client
type RESTClient struct {
	user        *user.RESTClient
	project     *project.RESTClient
	registry    *registry.RESTClient
	replication *replication.RESTClient
	system      *system.RESTClient
	retention   *retention.RESTClient
}

// NewRESTClient constructs a new REST client containing each sub client.
func NewRESTClient(legacyClient *client.Harbor, v2Client *v2client.Harbor, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		user:        user.NewClient(legacyClient, v2Client, authInfo),
		project:     project.NewClient(legacyClient, v2Client, authInfo),
		registry:    registry.NewClient(legacyClient, v2Client, authInfo),
		replication: replication.NewClient(legacyClient, v2Client, authInfo),
		system:      system.NewClient(legacyClient, v2Client, authInfo),
		retention:   retention.NewClient(legacyClient, v2Client, authInfo),
	}
}

// NewRESTClientForHost constructs a new REST client containing a swagger API client using the defined
// host string and basePath, the additional Harbor v2 API suffix as well as basic auth info.
func NewRESTClientForHost(u, username, password string) (*RESTClient, error) {
	if !strings.HasSuffix(u, v2URLSuffix) {
		u += v2URLSuffix
	}

	harborURL, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	legacySwaggerClient := client.New(runtimeclient.New(harborURL.Host, harborURL.Path, []string{harborURL.Scheme}), strfmt.Default)
	v2SwaggerClient := v2client.New(runtimeclient.New(harborURL.Host, harborURL.Path, []string{harborURL.Scheme}), strfmt.Default)
	authInfo := runtimeclient.BasicAuth(username, password)

	return NewRESTClient(legacySwaggerClient, v2SwaggerClient, authInfo), nil
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

// UpdateUserPassword wraps the UpdateUserPassword method of the user sub-package.
func (c *RESTClient) UpdateUserPassword(ctx context.Context, id int64, password *model.Password) error {
	return c.user.UpdateUserPassword(ctx, id, password)
}

// Project Client

// NewProject wraps the NewProject method of the project sub-package.
func (c *RESTClient) NewProject(ctx context.Context, name string, storageLimit int) (*modelv2.Project, error) {
	return c.project.NewProject(ctx, name, storageLimit)
}

// DeleteProject wraps the DeleteProject method of the registry sub-package.
func (c *RESTClient) DeleteProject(ctx context.Context, p *modelv2.Project) error {
	return c.project.DeleteProject(ctx, p)
}

// GetProjectByName wraps the GetProjectByName method of the registry sub-package.
func (c *RESTClient) GetProjectByName(ctx context.Context, name string) (*modelv2.Project, error) {
	return c.project.GetProjectByName(ctx, name)
}

// GetProjectByID wraps the GetProjectByID method of the registry sub-package.
func (c *RESTClient) GetProjectByID(ctx context.Context, projectID int64) (*modelv2.Project, error) {
	return c.project.GetProjectByID(ctx, projectID)
}

// ListProjects wraps the ListProjects method of the registry sub-package.
func (c *RESTClient) ListProjects(ctx context.Context, nameFilter string) ([]*modelv2.Project, error) {
	return c.project.ListProjects(ctx, nameFilter)
}

// UpdateProject wraps the UpdateProject method of the registry sub-package.
func (c *RESTClient) UpdateProject(ctx context.Context, p *modelv2.Project, storageLimit int) error {
	return c.project.UpdateProject(ctx, p, storageLimit)
}

// AddProjectMember wraps the AddProjectMember method of the registry sub-package.
func (c *RESTClient) AddProjectMember(ctx context.Context, p *modelv2.Project, u *model.User, roleID int) error {
	return c.project.AddProjectMember(ctx, p, u, roleID)
}

// ListProjectMembers wraps the ListProjectMembers method of the registry sub-package.
func (c *RESTClient) ListProjectMembers(ctx context.Context, p *modelv2.Project) ([]*model.ProjectMemberEntity, error) {
	return c.project.ListProjectMembers(ctx, p)
}

// UpdateProjectMemberRole wraps the UpdateProjectMemberRole method of the registry sub-package.
func (c *RESTClient) UpdateProjectMemberRole(ctx context.Context, p *modelv2.Project, u *model.User, roleID int) error {
	return c.project.UpdateProjectMemberRole(ctx, p, u, roleID)
}

// DeleteProjectMember wraps the DeleteProjectMember method of the registry sub-package.
func (c *RESTClient) DeleteProjectMember(ctx context.Context, p *modelv2.Project, u *model.User) error {
	return c.project.DeleteProjectMember(ctx, p, u)
}

// AddProjectMetadata wraps the AddProjectMetadata method of the registry sub-package.
func (c *RESTClient) AddProjectMetadata(ctx context.Context, p *modelv2.Project, key project.MetadataKey, value string) error {
	return c.project.AddProjectMetadata(ctx, p, key, value)
}

// ListProjectMetadata wraps the ListProjectMetadata method of the registry sub-package.
func (c *RESTClient) ListProjectMetadata(ctx context.Context, p *modelv2.Project) (*modelv2.ProjectMetadata, error) {
	return c.project.ListProjectMetadata(ctx, p)
}

// GetProjectMetadataValue wraps the GetProjectMetadataValue method of the registry sub-package.
func (c *RESTClient) GetProjectMetadataValue(ctx context.Context, projectID int64, key project.MetadataKey) (string, error) {
	return c.project.GetProjectMetadataValue(ctx, projectID, key)
}

// UpdateProjectMetadata wraps the UpdateProjectMetadata method of the registry sub-package.
func (c *RESTClient) UpdateProjectMetadata(ctx context.Context, p *modelv2.Project, key project.MetadataKey, value string) error {
	return c.project.UpdateProjectMetadata(ctx, p, key, value)
}

// DeleteProjectMetadataValue wraps the DeleteProjectMetadataValue method of the registry sub-package.
func (c *RESTClient) DeleteProjectMetadataValue(ctx context.Context, p *modelv2.Project, key project.MetadataKey) error {
	return c.project.DeleteProjectMetadataValue(ctx, p, key)
}

// ListProjectRobots wraps the ListProjectRobots method of the project sub-package.
func (c *RESTClient) ListProjectRobots(ctx context.Context, p *modelv2.Project) ([]*model.RobotAccount, error) {
	return c.project.ListProjectRobots(ctx, p)
}

// AddProjectRobot wraps the AddProjectRobot method of the project sub-package.
func (c *RESTClient) AddProjectRobot(ctx context.Context, p *modelv2.Project, robot *model.RobotAccountCreate) (string, error) {
	return c.project.AddProjectRobot(ctx, p, robot)
}

// UpdateProjectRobot wraps the UpdateProjectRobot method of the project sub-package.
func (c *RESTClient) UpdateProjectRobot(ctx context.Context, p *modelv2.Project, robotID int, robot *model.RobotAccountUpdate) error {
	return c.project.UpdateProjectRobot(ctx, p, robotID, robot)
}

// DeleteProjectRobot wraps the DeleteProjectRobot method of the project sub-package.
func (c *RESTClient) DeleteProjectRobot(ctx context.Context, p *modelv2.Project, robotID int) error {
	return c.project.DeleteProjectRobot(ctx, p, robotID)
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

// UpdateRegistry wraps the UpdateRegistry method of the registry sub-package.
func (c *RESTClient) UpdateRegistry(ctx context.Context, r *model.Registry) error {
	return c.registry.UpdateRegistry(ctx, r)
}

// Replication Client

// NewReplicationPolicy wraps the NewReplicationPolicy method of the replication sub-package.
func (c *RESTClient) NewReplicationPolicy(ctx context.Context, destRegistry, srcRegistry *model.Registry,
	replicateDeletion, override, enablePolicy bool, filters []*model.ReplicationFilter,
	trigger *model.ReplicationTrigger, destNamespace, description, name string) (*model.ReplicationPolicy, error) {
	return c.replication.NewReplicationPolicy(ctx, destRegistry, srcRegistry, replicateDeletion,
		override, enablePolicy, filters, trigger, destNamespace, description, name)
}

// GetReplicationPolicy wraps the GetReplicationPolicy method of the replication sub-package.
func (c *RESTClient) GetReplicationPolicy(ctx context.Context, name string) (*model.ReplicationPolicy, error) {
	return c.replication.GetReplicationPolicy(ctx, name)
}

// GetReplicationPolicyByID wraps the GetReplicationPolicyByID method of the replication sub-package.
func (c *RESTClient) GetReplicationPolicyByID(ctx context.Context, id int64) (*model.ReplicationPolicy, error) {
	return c.replication.GetReplicationPolicyByID(ctx, id)
}

// DeleteReplicationPolicy wraps the DeleteReplicationPolicy method of the replication sub-package.
func (c *RESTClient) DeleteReplicationPolicy(ctx context.Context, r *model.ReplicationPolicy) error {
	return c.replication.DeleteReplicationPolicy(ctx, r)
}

// UpdateReplicationPolicy wraps the UpdateReplicationPolicy method of the replication sub-package.
func (c *RESTClient) UpdateReplicationPolicy(ctx context.Context, r *model.ReplicationPolicy) error {
	return c.replication.UpdateReplicationPolicy(ctx, r)
}

// TriggerReplicationExecution wraps the TriggerReplicationExecution method of the replication sub-package.
func (c *RESTClient) TriggerReplicationExecution(ctx context.Context, r *model.ReplicationExecution) error {
	return c.replication.TriggerReplicationExecution(ctx, r)
}

// GetReplicationExecutions wraps the GetReplicationExecutions method of the replication sub-package.
func (c *RESTClient) GetReplicationExecutions(ctx context.Context,
	r *model.ReplicationExecution) ([]*model.ReplicationExecution, error) {
	return c.replication.GetReplicationExecutions(ctx, r)
}

// GetReplicationExecutionsByID wraps the GetReplicationExecutionsByID method of the replication sub-package.
func (c *RESTClient) GetReplicationExecutionByID(ctx context.Context, id int64) (*model.ReplicationExecution, error) {
	return c.replication.GetReplicationExecutionByID(ctx, id)
}

// System Client

// NewSystemGarbageCollection wraps the NewSystemGarbageCollection method of the system sub-package.
func (c *RESTClient) NewSystemGarbageCollection(ctx context.Context,
	cron, scheduleType string) (*model.AdminJobSchedule, error) {
	return c.system.NewSystemGarbageCollection(ctx, cron, scheduleType)
}

// UpdateSystemGarbageCollection wraps the UpdateSystemGarbageCollection method of the system sub-package.
func (c *RESTClient) UpdateSystemGarbageCollection(ctx context.Context,
	newGcSchedule *model.AdminJobScheduleObj) error {
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

// Retention Client

// NewRetentionPolicy wraps the NewRetentionPolicy method of the retention sub-package.
func (c *RESTClient) NewRetentionPolicy(ctx context.Context, rep *model.RetentionPolicy) error {
	return c.retention.NewRetentionPolicy(ctx, rep)
}

// GetRetentionPolicyByProjectID wraps the GetRetentionPolicyByProject method of the retention sub-package.
func (c *RESTClient) GetRetentionPolicyByProject(ctx context.Context, project *modelv2.Project) (*model.RetentionPolicy, error) {
	return c.retention.GetRetentionPolicyByProject(ctx, project)
}

// UpdateRetentionPolicy wraps the UpdateRetentionPolicy method of the retention sub-package.
func (c *RESTClient) UpdateRetentionPolicy(ctx context.Context, ret *model.RetentionPolicy) error {
	return c.retention.UpdateRetentionPolicy(ctx, ret)
}

// DisableRetentionPolicy wraps the DisableRetentionPolicy method of the retention sub-package.
func (c *RESTClient) DisableRetentionPolicy(ctx context.Context, ret *model.RetentionPolicy) error {
	return c.retention.DisableRetentionPolicy(ctx, ret)
}
