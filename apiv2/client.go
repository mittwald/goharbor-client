package apiv2

import (
	"context"
	"net/url"
	"strings"

	"github.com/mittwald/goharbor-client/v4/apiv2/auditlog"
	"github.com/mittwald/goharbor-client/v4/apiv2/gc"
	modelv2 "github.com/mittwald/goharbor-client/v4/apiv2/model"
	"github.com/mittwald/goharbor-client/v4/apiv2/quota"
	"github.com/mittwald/goharbor-client/v4/apiv2/retention"
	"github.com/mittwald/goharbor-client/v4/apiv2/robot"

	"github.com/go-openapi/runtime"
	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	v2client "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client"

	"github.com/mittwald/goharbor-client/v4/apiv2/internal/legacyapi/client"
	legacymodel "github.com/mittwald/goharbor-client/v4/apiv2/model/legacy"
	"github.com/mittwald/goharbor-client/v4/apiv2/project"
	"github.com/mittwald/goharbor-client/v4/apiv2/registry"
	"github.com/mittwald/goharbor-client/v4/apiv2/replication"
	"github.com/mittwald/goharbor-client/v4/apiv2/system"
	"github.com/mittwald/goharbor-client/v4/apiv2/user"
)

const v2URLSuffix string = "/v2.0"

type Client interface {
	user.Client
	project.Client
	registry.Client
	replication.Client
	system.Client
	retention.Client
	quota.Client
	gc.Client
}

// RESTClient implements the Client interface as a REST client
type RESTClient struct {
	auditlog    *auditlog.RESTClient
	user        *user.RESTClient
	project     *project.RESTClient
	registry    *registry.RESTClient
	replication *replication.RESTClient
	system      *system.RESTClient
	retention   *retention.RESTClient
	quota       *quota.RESTClient
	gc          *gc.RESTClient
	robot       *robot.RESTClient
}

// NewRESTClient constructs a new REST client containing each sub client.
func NewRESTClient(legacyClient *client.Harbor, v2Client *v2client.Harbor, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		auditlog:    auditlog.NewClient(v2Client, authInfo),
		user:        user.NewClient(legacyClient, v2Client, authInfo),
		project:     project.NewClient(legacyClient, v2Client, authInfo),
		registry:    registry.NewClient(legacyClient, v2Client, authInfo),
		replication: replication.NewClient(legacyClient, v2Client, authInfo),
		system:      system.NewClient(legacyClient, v2Client, authInfo),
		retention:   retention.NewClient(legacyClient, v2Client, authInfo),
		quota:       quota.NewClient(legacyClient, v2Client, authInfo),
		gc:          gc.NewClient(legacyClient, v2Client, authInfo),
		robot:       robot.NewClient(v2Client, authInfo),
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
func (c *RESTClient) NewUser(ctx context.Context, username, email, realname, password, comments string) (*legacymodel.User, error) {
	return c.user.NewUser(ctx, username, email, realname, password, comments)
}

// GetUser wraps the GetUser method of the user sub-package.
func (c *RESTClient) GetUser(ctx context.Context, username string) (*legacymodel.User, error) {
	return c.user.GetUser(ctx, username)
}

// GetUserByID wraps the GetUserByID method of the user sub-package.
func (c *RESTClient) GetUserByID(ctx context.Context, id int64) (*legacymodel.User, error) {
	return c.user.GetUserByID(ctx, id)
}

// ListUsers wraps the ListUsers method of the user sub-package.
func (c *RESTClient) ListUsers(ctx context.Context) ([]*legacymodel.User, error) {
	return c.user.ListUsers(ctx)
}

// DeleteUser wraps the DeleteUser method of the user sub-package.
func (c *RESTClient) DeleteUser(ctx context.Context, u *legacymodel.User) error {
	return c.user.DeleteUser(ctx, u)
}

// UpdateUser wraps the UpdateUser method of the user sub-package.
func (c *RESTClient) UpdateUser(ctx context.Context, u *legacymodel.User) error {
	return c.user.UpdateUser(ctx, u)
}

// UpdateUserPassword wraps the UpdateUserPassword method of the user sub-package.
func (c *RESTClient) UpdateUserPassword(ctx context.Context, id int64, password *legacymodel.Password) error {
	return c.user.UpdateUserPassword(ctx, id, password)
}

// UserExists wraps the UserExists method of the user sub-package.
func (c *RESTClient) UserExists(ctx context.Context, u *legacymodel.User) (bool, error) {
	return c.user.UserExists(ctx, u)
}

// Project Client

// NewProject wraps the NewProject method of the project sub-package.
func (c *RESTClient) NewProject(ctx context.Context, name string, storageLimit *int64) (*modelv2.Project, error) {
	return c.project.NewProject(ctx, name, storageLimit)
}

// DeleteProject wraps the DeleteProject method of the project sub-package.
func (c *RESTClient) DeleteProject(ctx context.Context, p *modelv2.Project) error {
	return c.project.DeleteProject(ctx, p)
}

// GetProject wraps the GetProject method of the project sub-package.
func (c *RESTClient) GetProject(ctx context.Context, nameOrID string) (*modelv2.Project, error) {
	return c.project.GetProject(ctx, nameOrID)
}

// ListProjects wraps the ListProjects method of the project sub-package.
func (c *RESTClient) ListProjects(ctx context.Context, nameFilter string) ([]*modelv2.Project, error) {
	return c.project.ListProjects(ctx, nameFilter)
}

// UpdateProject wraps the ListProjects method of the project sub-package.
func (c *RESTClient) UpdateProject(ctx context.Context, p *modelv2.Project, storageLimit *int64) error {
	return c.project.UpdateProject(ctx, p, storageLimit)
}

// AddProjectMember wraps the AddProjectMember method of the project sub-package.
func (c *RESTClient) AddProjectMember(ctx context.Context, p *modelv2.Project, u *legacymodel.User, roleID int) error {
	return c.project.AddProjectMember(ctx, p, u, roleID)
}

// ListProjectMembers wraps the ListProjectMembers method of the project sub-package.
func (c *RESTClient) ListProjectMembers(ctx context.Context, p *modelv2.Project) ([]*legacymodel.ProjectMemberEntity, error) {
	return c.project.ListProjectMembers(ctx, p)
}

// UpdateProjectMemberRole wraps the UpdateProjectMemberRole method of the project sub-package.
func (c *RESTClient) UpdateProjectMemberRole(ctx context.Context, p *modelv2.Project, u *legacymodel.User, roleID int) error {
	return c.project.UpdateProjectMemberRole(ctx, p, u, roleID)
}

// DeleteProjectMember wraps the DeleteProjectMember method of the project sub-package.
func (c *RESTClient) DeleteProjectMember(ctx context.Context, p *modelv2.Project, u *legacymodel.User) error {
	return c.project.DeleteProjectMember(ctx, p, u)
}

// AddProjectMetadata wraps the AddProjectMetadata method of the project sub-package.
func (c *RESTClient) AddProjectMetadata(ctx context.Context, p *modelv2.Project, key project.MetadataKey, value string) error {
	return c.project.AddProjectMetadata(ctx, p, key, value)
}

// GetProjectMetadataValue wraps the GetProjectMetadataValue method of the project sub-package.
func (c *RESTClient) GetProjectMetadataValue(ctx context.Context, projectNameOrID string, key project.MetadataKey) (string, error) {
	return c.project.GetProjectMetadataValue(ctx, projectNameOrID, key)
}

// ListProjectMetadata wraps the ListProjectMetadata method of the project sub-package.
func (c *RESTClient) ListProjectMetadata(ctx context.Context, p *modelv2.Project) (*modelv2.ProjectMetadata, error) {
	return c.project.ListProjectMetadata(ctx, p)
}

// UpdateProjectMetadata wraps the UpdateProjectMetadata method of the project sub-package.
func (c *RESTClient) UpdateProjectMetadata(ctx context.Context, p *modelv2.Project, key project.MetadataKey, value string) error {
	return c.project.UpdateProjectMetadata(ctx, p, key, value)
}

// DeleteProjectMetadataValue wraps the DeleteProjectMetadataValue method of the project sub-package.
func (c *RESTClient) DeleteProjectMetadataValue(ctx context.Context, p *modelv2.Project, key project.MetadataKey) error {
	return c.project.DeleteProjectMetadataValue(ctx, p, key)
}

// ListProjectRobots wraps the ListProjectRobots method of the project sub-package.
func (c *RESTClient) ListProjectRobots(ctx context.Context, p *modelv2.Project) ([]*modelv2.Robot, error) {
	return c.project.ListProjectRobots(ctx, p)
}

// AddProjectRobot wraps the AddProjectRobot method of the project sub-package.
func (c *RESTClient) AddProjectRobot(ctx context.Context, p *modelv2.Project, r *modelv2.RobotCreateV1) (*modelv2.RobotCreated, error) {
	return c.project.AddProjectRobot(ctx, p, r)
}

// UpdateProjectRobot wraps the UpdateProjectRobot method of the project sub-package.
func (c *RESTClient) UpdateProjectRobot(ctx context.Context, p *modelv2.Project, robotID int64, r *modelv2.Robot) error {
	return c.project.UpdateProjectRobot(ctx, p, robotID, r)
}

// DeleteProjectRobot wraps the DeleteProjectRobot method of the project sub-package.
func (c *RESTClient) DeleteProjectRobot(ctx context.Context, p *modelv2.Project, robotID int64) error {
	return c.project.DeleteProjectRobot(ctx, p, robotID)
}

// ListProjectWebhookPolicies wraps the ListProjectWebhookPolicies method of the project sub-package.
func (c *RESTClient) ListProjectWebhookPolicies(ctx context.Context, p *modelv2.Project) ([]*legacymodel.WebhookPolicy, error) {
	return c.project.ListProjectWebhookPolicies(ctx, p)
}

// UpdateProjectWebhookPolicy wraps the UpdateProjectWebhookPolicy method of the project sub-package.
func (c *RESTClient) UpdateProjectWebhookPolicy(ctx context.Context, p *modelv2.Project, policyID int64, policy *legacymodel.WebhookPolicy) error {
	return c.project.UpdateProjectWebhookPolicy(ctx, p, policyID, policy)
}

// AddProjectWebhookPolicy wraps the AddProjectWebhookPolicy method of the project sub-package.
func (c *RESTClient) AddProjectWebhookPolicy(ctx context.Context, p *modelv2.Project, policy *legacymodel.WebhookPolicy) error {
	return c.project.AddProjectWebhookPolicy(ctx, p, policy)
}

// DeleteProjectWebhookPolicy wraps the DeleteProjectWebhookPolicy method of the project sub-package.
func (c *RESTClient) DeleteProjectWebhookPolicy(ctx context.Context, p *modelv2.Project, policyID int64) error {
	return c.project.DeleteProjectWebhookPolicy(ctx, p, policyID)
}

// Registry Client

// NewRegistry wraps the NewRegistry method of the registry sub-package.
func (c *RESTClient) NewRegistry(ctx context.Context, name, registryType, url string,
	credential *legacymodel.RegistryCredential, insecure bool) (*legacymodel.Registry, error) {
	return c.registry.NewRegistry(ctx, name, registryType, url,
		credential, insecure)
}

// GetRegistry wraps the GetRegistry method of the registry sub-package.
func (c *RESTClient) GetRegistry(ctx context.Context, name string) (*legacymodel.Registry, error) {
	return c.registry.GetRegistry(ctx, name)
}

// DeleteRegistry wraps the DeleteRegistry method of the registry sub-package.
func (c *RESTClient) DeleteRegistry(ctx context.Context, r *legacymodel.Registry) error {
	return c.registry.DeleteRegistry(ctx, r)
}

// UpdateRegistry wraps the UpdateRegistry method of the registry sub-package.
func (c *RESTClient) UpdateRegistry(ctx context.Context, r *legacymodel.Registry) error {
	return c.registry.UpdateRegistry(ctx, r)
}

// Replication Client

// NewReplicationPolicy wraps the NewReplicationPolicy method of the replication sub-package.
func (c *RESTClient) NewReplicationPolicy(ctx context.Context, destRegistry, srcRegistry *legacymodel.Registry,
	replicateDeletion, override, enablePolicy bool, filters []*legacymodel.ReplicationFilter,
	trigger *legacymodel.ReplicationTrigger, destNamespace, description, name string) (*legacymodel.ReplicationPolicy, error) {
	return c.replication.NewReplicationPolicy(ctx, destRegistry, srcRegistry, replicateDeletion,
		override, enablePolicy, filters, trigger, destNamespace, description, name)
}

// GetReplicationPolicy wraps the GetReplicationPolicy method of the replication sub-package.
func (c *RESTClient) GetReplicationPolicy(ctx context.Context, name string) (*legacymodel.ReplicationPolicy, error) {
	return c.replication.GetReplicationPolicy(ctx, name)
}

// GetReplicationPolicyByID wraps the GetReplicationPolicyByID method of the replication sub-package.
func (c *RESTClient) GetReplicationPolicyByID(ctx context.Context, id int64) (*legacymodel.ReplicationPolicy, error) {
	return c.replication.GetReplicationPolicyByID(ctx, id)
}

// DeleteReplicationPolicy wraps the DeleteReplicationPolicy method of the replication sub-package.
func (c *RESTClient) DeleteReplicationPolicy(ctx context.Context, r *legacymodel.ReplicationPolicy) error {
	return c.replication.DeleteReplicationPolicy(ctx, r)
}

// UpdateReplicationPolicy wraps the UpdateReplicationPolicy method of the replication sub-package.
func (c *RESTClient) UpdateReplicationPolicy(ctx context.Context, r *legacymodel.ReplicationPolicy) error {
	return c.replication.UpdateReplicationPolicy(ctx, r)
}

// TriggerReplicationExecution wraps the TriggerReplicationExecution method of the replication sub-package.
func (c *RESTClient) TriggerReplicationExecution(ctx context.Context, r *modelv2.StartReplicationExecution) error {
	return c.replication.TriggerReplicationExecution(ctx, r)
}

// GetReplicationExecutions wraps the GetReplicationExecutions method of the replication sub-package.
func (c *RESTClient) GetReplicationExecutions(ctx context.Context, r *modelv2.ReplicationExecution) ([]*modelv2.ReplicationExecution, error) {
	return c.replication.GetReplicationExecutions(ctx, r)
}

// GetReplicationExecutionByID GetReplicationExecutionsByID wraps the GetReplicationExecutionsByID method of the replication sub-package.
func (c *RESTClient) GetReplicationExecutionByID(ctx context.Context, id int64) (*modelv2.ReplicationExecution, error) {
	return c.replication.GetReplicationExecutionByID(ctx, id)
}

// GarbageCollection Client

// NewGarbageCollection wraps the NewSystemGarbageCollection method of the system sub-package.
func (c *RESTClient) NewGarbageCollection(ctx context.Context, gcSchedule *modelv2.Schedule) error {
	return c.gc.NewGarbageCollection(ctx, gcSchedule)
}

// UpdateGarbageCollection wraps the UpdateSystemGarbageCollection method of the system sub-package.
func (c *RESTClient) UpdateGarbageCollection(ctx context.Context, newGCSchedule *modelv2.Schedule) error {
	return c.gc.UpdateGarbageCollection(ctx, newGCSchedule)
}

// GetGarbageCollectionSchedule wraps the GetSystemGarbageCollection method of the system sub-package.
func (c *RESTClient) GetGarbageCollectionSchedule(ctx context.Context) (*modelv2.GCHistory, error) {
	return c.gc.GetGarbageCollectionSchedule(ctx)
}

// ResetGarbageCollection wraps the ResetSystemGarbageCollection method of the system sub-package.
func (c *RESTClient) ResetGarbageCollection(ctx context.Context) error {
	return c.gc.ResetGarbageCollection(ctx)
}

// System Client

// Health wraps the Health method of the system sub-package.
func (c *RESTClient) Health(ctx context.Context) (*legacymodel.OverallHealthStatus, error) {
	return c.system.Health(ctx)
}

// Retention Client

// NewRetentionPolicy wraps the NewRetentionPolicy method of the retention sub-package.
func (c *RESTClient) NewRetentionPolicy(ctx context.Context, ret *modelv2.RetentionPolicy) error {
	return c.retention.NewRetentionPolicy(ctx, ret)
}

// GetRetentionPolicyByProject GetRetentionPolicyByProjectID wraps the GetRetentionPolicyByProject method of the retention sub-package.
func (c *RESTClient) GetRetentionPolicyByProject(ctx context.Context, project *modelv2.Project) (*modelv2.RetentionPolicy, error) {
	return c.retention.GetRetentionPolicyByProject(ctx, project)
}

// GetRetentionPolicyByID wraps the GetRetentionPolicyByID method of the retention sub-package.
func (c *RESTClient) GetRetentionPolicyByID(ctx context.Context, id int64) (*modelv2.RetentionPolicy, error) {
	return c.retention.GetRetentionPolicyByID(ctx, id)
}

// UpdateRetentionPolicy wraps the UpdateRetentionPolicy method of the retention sub-package.
func (c *RESTClient) UpdateRetentionPolicy(ctx context.Context, ret *modelv2.RetentionPolicy) error {
	return c.retention.UpdateRetentionPolicy(ctx, ret)
}

// DisableRetentionPolicy wraps the DisableRetentionPolicy method of the retention sub-package.
func (c *RESTClient) DisableRetentionPolicy(ctx context.Context, ret *modelv2.RetentionPolicy) error {
	return c.retention.DisableRetentionPolicy(ctx, ret)
}

// Quota Client

// GetQuotaByProjectID wraps the GetQuotaByProjectID method of the quota sub-package.
func (c *RESTClient) GetQuotaByProjectID(ctx context.Context, projectID int64) (*legacymodel.Quota, error) {
	return c.quota.GetQuotaByProjectID(ctx, projectID)
}

// UpdateStorageQuotaByProjectID wraps the UpdateStorageQuotaByProjectID method of the quota sub-package.
func (c *RESTClient) UpdateStorageQuotaByProjectID(ctx context.Context, projectID int64, storageLimit int64) error {
	return c.quota.UpdateStorageQuotaByProjectID(ctx, projectID, storageLimit)
}

// Robot Client

// ListRobotAccounts wraps the ListRobotAccounts method of the robot sub-package.
func (c *RESTClient) ListRobotAccounts(ctx context.Context) ([]*modelv2.Robot, error) {
	return c.robot.ListRobotAccounts(ctx)
}

// GetRobotAccountByName wraps the GetRobotAccountByName method of the robot sub-package.
func (c *RESTClient) GetRobotAccountByName(ctx context.Context, name string) (*modelv2.Robot, error) {
	return c.robot.GetRobotAccountByName(ctx, name)
}

// GetRobotAccountByID wraps the GetRobotAccountByID method of the robot sub-package.
func (c *RESTClient) GetRobotAccountByID(ctx context.Context, id int64) (*modelv2.Robot, error) {
	return c.robot.GetRobotAccountByID(ctx, id)
}

// NewRobotAccount wraps the NewRobotAccount method of the robot sub-package.
func (c *RESTClient) NewRobotAccount(ctx context.Context, r *modelv2.RobotCreate) (*modelv2.RobotCreated, error) {
	return c.robot.NewRobotAccount(ctx, r)
}

// DeleteRobotAccountByName wraps the DeleteRobotAccountByName method of the robot sub-package.
func (c *RESTClient) DeleteRobotAccountByName(ctx context.Context, name string) error {
	return c.robot.DeleteRobotAccountByName(ctx, name)
}

// DeleteRobotAccountByID wraps the DeleteRobotAccountByID method of the robot sub-package.
func (c *RESTClient) DeleteRobotAccountByID(ctx context.Context, id int64) error {
	return c.robot.DeleteRobotAccountByID(ctx, id)
}

// UpdateRobotAccount wraps the UpdateRobotAccount method of the robot sub-package.
func (c *RESTClient) UpdateRobotAccount(ctx context.Context, r *modelv2.Robot) error {
	return c.robot.UpdateRobotAccount(ctx, r)
}

// RefreshRobotAccountSecretByID wraps the RefreshRobotAccountSecretByID method of the robot sub-package.
func (c *RESTClient) RefreshRobotAccountSecretByID(ctx context.Context, id int64, sec string) (*modelv2.RobotSec, error) {
	return c.robot.RefreshRobotAccountSecretByID(ctx, id, sec)
}

// RefreshRobotAccountSecretByName wraps the RefreshRobotAccountSecretByName method of the robot sub-package.
func (c *RESTClient) RefreshRobotAccountSecretByName(ctx context.Context, name string, sec string) (*modelv2.RobotSec, error) {
	return c.robot.RefreshRobotAccountSecretByName(ctx, name, sec)
}

// AuditLog Client

// ListAuditLogs wraps the ListAuditLogs method of the auditlog sub-package.
func (c *RESTClient) ListAuditLogs(ctx context.Context, pageSize *int64, query *string) ([]*modelv2.AuditLog, error) {
	return c.auditlog.ListAuditLogs(ctx, pageSize, query)
}
