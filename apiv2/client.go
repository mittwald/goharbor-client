package apiv2

import (
	"context"
	"net/url"
	"strings"

	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/clients/label"

	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/clients/artifact"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/clients/repository"

	"k8s.io/apimachinery/pkg/util/intstr"

	modelv2 "github.com/mittwald/goharbor-client/v5/apiv2/model"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/clients/auditlog"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/clients/gc"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/clients/health"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/clients/member"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/clients/projectmeta"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/clients/quota"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/clients/retention"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/clients/robot"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/clients/robotv1"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/clients/webhook"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/common"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/config"

	"github.com/go-openapi/runtime"
	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	v2client "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client"

	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/clients/project"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/clients/registry"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/clients/replication"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/clients/systeminfo"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/clients/user"
)

const v2URLSuffix string = "/v2.0"

type Client interface {
	auditlog.Client
	artifact.Client
	gc.Client
	health.Client
	label.Client
	member.Client
	project.Client
	projectmeta.Client
	quota.Client
	registry.Client
	replication.Client
	repository.Client
	retention.Client
	robot.Client
	robotv1.Client
	systeminfo.Client
	user.Client
	webhook.Client
}

// RESTClient implements the Client interface as a REST client
type RESTClient struct {
	auditlog    *auditlog.RESTClient
	artifact    *artifact.RESTClient
	gc          *gc.RESTClient
	health      *health.RESTClient
	label       *label.RESTClient
	member      *member.RESTClient
	project     *project.RESTClient
	projectmeta *projectmeta.RESTClient
	quota       *quota.RESTClient
	registry    *registry.RESTClient
	replication *replication.RESTClient
	repository  *repository.RESTClient
	retention   *retention.RESTClient
	robot       *robot.RESTClient
	robotv1     *robotv1.RESTClient
	systeminfo  *systeminfo.RESTClient
	user        *user.RESTClient
	webhook     *webhook.RESTClient
}

// NewRESTClient constructs a new REST client containing each sub client.
func NewRESTClient(v2Client *v2client.Harbor, opts *config.Options, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	if opts == nil {
		opts = config.Defaults()
	}

	return &RESTClient{
		auditlog:    auditlog.NewClient(v2Client, opts, authInfo),
		artifact:    artifact.NewClient(v2Client, opts, authInfo),
		gc:          gc.NewClient(v2Client, opts, authInfo),
		health:      health.NewClient(v2Client, opts, authInfo),
		label:       label.NewClient(v2Client, opts, authInfo),
		member:      member.NewClient(v2Client, opts, authInfo),
		project:     project.NewClient(v2Client, opts, authInfo),
		projectmeta: projectmeta.NewClient(v2Client, opts, authInfo),
		quota:       quota.NewClient(v2Client, opts, authInfo),
		registry:    registry.NewClient(v2Client, opts, authInfo),
		replication: replication.NewClient(v2Client, opts, authInfo),
		repository:  repository.NewClient(v2Client, opts, authInfo),
		retention:   retention.NewClient(v2Client, opts, authInfo),
		robot:       robot.NewClient(v2Client, opts, authInfo),
		robotv1:     robotv1.NewClient(v2Client, opts, authInfo),
		systeminfo:  systeminfo.NewClient(v2Client, opts, authInfo),
		user:        user.NewClient(v2Client, opts, authInfo),
		webhook:     webhook.NewClient(v2Client, opts, authInfo),
	}
}

// NewRESTClientForHost constructs a new REST client containing a swagger API client using the defined
// host string and basePath, the additional Harbor v2 API suffix as well as basic auth info.
func NewRESTClientForHost(u, username, password string, opts *config.Options, clientOpts runtimeclient.TLSClientOptions) (*RESTClient, error) {
	if !strings.HasSuffix(u, v2URLSuffix) {
		u += v2URLSuffix
	}

	harborURL, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	client, err := runtimeclient.TLSClient(clientOpts)
	if err != nil {
		return nil, err
	}

	v2SwaggerClient := v2client.New(runtimeclient.NewWithClient(harborURL.Host, harborURL.Path, []string{harborURL.Scheme}, client), strfmt.Default)
	authInfo := runtimeclient.BasicAuth(username, password)

	return NewRESTClient(v2SwaggerClient, opts, authInfo), nil
}

// AuditLog Client

func (c *RESTClient) ListAuditLogs(ctx context.Context) ([]*modelv2.AuditLog, error) {
	return c.auditlog.ListAuditLogs(ctx)
}

// Artifact Client

// TODO: Introduce this, once https://github.com/goharbor/harbor/issues/13468 is resolved.
//func (c *RESTClient) GetVulnerabilitiesAddition(ctx context.Context, projectName, repositoryName, reference string) (interface{}, error) {
//	return c.artifact.GetVulnerabilitiesAddition(ctx, projectName, repositoryName, reference)
//}
//
//func (c *RESTClient) GetAddition(ctx context.Context, projectName, repositoryName, reference string, addition artifact.Addition) (interface{}, error) {
//	return c.artifact.GetAddition(ctx, projectName, repositoryName, reference, addition)
//}

func (c *RESTClient) AddArtifactLabel(ctx context.Context, projectName, repositoryName, reference string, label *modelv2.Label) error {
	return c.artifact.AddArtifactLabel(ctx, projectName, repositoryName, reference, label)
}

func (c *RESTClient) CopyArtifact(ctx context.Context, from *artifact.CopyReference, projectName, repositoryName string) error {
	return c.artifact.CopyArtifact(ctx, from, projectName, repositoryName)
}

func (c *RESTClient) CreateTag(ctx context.Context, projectName, repositoryName, reference string, tag *modelv2.Tag) error {
	return c.artifact.CreateTag(ctx, projectName, repositoryName, reference, tag)
}

func (c *RESTClient) DeleteTag(ctx context.Context, projectName, repositoryName, reference, tagName string) error {
	return c.artifact.DeleteTag(ctx, projectName, repositoryName, reference, tagName)
}

func (c *RESTClient) GetArtifact(ctx context.Context, projectName, repositoryName, reference string) (*modelv2.Artifact, error) {
	return c.artifact.GetArtifact(ctx, projectName, repositoryName, reference)
}

func (c *RESTClient) ListArtifacts(ctx context.Context, projectName, repositoryName string) ([]*modelv2.Artifact, error) {
	return c.artifact.ListArtifacts(ctx, projectName, repositoryName)
}

func (c *RESTClient) ListTags(ctx context.Context, projectName, repositoryName, reference string) ([]*modelv2.Tag, error) {
	return c.artifact.ListTags(ctx, projectName, repositoryName, reference)
}

func (c *RESTClient) RemoveLabel(ctx context.Context, projectName, repositoryName, reference string, id int64) error {
	return c.artifact.RemoveLabel(ctx, projectName, repositoryName, reference, id)
}

// GC Client

func (c *RESTClient) NewGarbageCollection(ctx context.Context, gcSchedule *modelv2.Schedule) error {
	return c.gc.NewGarbageCollection(ctx, gcSchedule)
}

func (c *RESTClient) UpdateGarbageCollection(ctx context.Context, newGCSchedule *modelv2.Schedule) error {
	return c.gc.UpdateGarbageCollection(ctx, newGCSchedule)
}

func (c *RESTClient) GetGarbageCollectionExecution(ctx context.Context, id int64) (*modelv2.GCHistory, error) {
	return c.gc.GetGarbageCollectionExecution(ctx, id)
}

func (c *RESTClient) GetGarbageCollectionSchedule(ctx context.Context) (*modelv2.GCHistory, error) {
	return c.gc.GetGarbageCollectionSchedule(ctx)
}

func (c *RESTClient) ResetGarbageCollection(ctx context.Context) error {
	return c.gc.ResetGarbageCollection(ctx)
}

// Health Client

func (c *RESTClient) GetHealth(ctx context.Context) (*modelv2.OverallHealthStatus, error) {
	return c.health.GetHealth(ctx)
}

// Label Client

func (c *RESTClient) CreateLabel(ctx context.Context, l *modelv2.Label) error {
	return c.label.CreateLabel(ctx, l)
}

func (c *RESTClient) GetLabelByID(ctx context.Context, id int64) (*modelv2.Label, error) {
	return c.label.GetLabelByID(ctx, id)
}

func (c *RESTClient) ListLabels(ctx context.Context, name string, projectID *int64) ([]*modelv2.Label, error) {
	return c.label.ListLabels(ctx, name, projectID)
}

func (c *RESTClient) DeleteLabel(ctx context.Context, id int64) error {
	return c.label.DeleteLabel(ctx, id)
}

func (c *RESTClient) UpdateLabel(ctx context.Context, id int64, l *modelv2.Label) error {
	return c.label.UpdateLabel(ctx, id, l)
}

// Member Client

func (c *RESTClient) AddProjectMember(ctx context.Context, projectNameOrID string, m *modelv2.ProjectMember) error {
	return c.member.AddProjectMember(ctx, projectNameOrID, m)
}

func (c *RESTClient) ListProjectMembers(ctx context.Context, projectNameOrID, memberQuery string) ([]*modelv2.ProjectMemberEntity, error) {
	return c.member.ListProjectMembers(ctx, projectNameOrID, memberQuery)
}

func (c *RESTClient) UpdateProjectMember(ctx context.Context, projectNameOrID string, m *modelv2.ProjectMember) error {
	return c.member.UpdateProjectMember(ctx, projectNameOrID, m)
}

func (c *RESTClient) DeleteProjectMember(ctx context.Context, projectNameOrID string, m *modelv2.ProjectMember) error {
	return c.member.DeleteProjectMember(ctx, projectNameOrID, m)
}

// Project Client

func (c *RESTClient) NewProject(ctx context.Context, projectRequest *modelv2.ProjectReq) error {
	return c.project.NewProject(ctx, projectRequest)
}

func (c *RESTClient) DeleteProject(ctx context.Context, nameOrID string) error {
	return c.project.DeleteProject(ctx, nameOrID)
}

func (c *RESTClient) GetProject(ctx context.Context, nameOrID string) (*modelv2.Project, error) {
	return c.project.GetProject(ctx, nameOrID)
}

func (c *RESTClient) ListProjects(ctx context.Context, nameFilter string) ([]*modelv2.Project, error) {
	return c.project.ListProjects(ctx, nameFilter)
}

func (c *RESTClient) UpdateProject(ctx context.Context, p *modelv2.Project, storageLimit *int64) error {
	return c.project.UpdateProject(ctx, p, storageLimit)
}

func (c *RESTClient) ProjectExists(ctx context.Context, nameOrID string) (bool, error) {
	return c.project.ProjectExists(ctx, nameOrID)
}

// Projectmeta Client

func (c *RESTClient) AddProjectMetadata(ctx context.Context, projectNameOrID string, key common.MetadataKey, value string) error {
	return c.projectmeta.AddProjectMetadata(ctx, projectNameOrID, key, value)
}

func (c *RESTClient) GetProjectMetadataValue(ctx context.Context, projectNameOrID string, key common.MetadataKey) (string, error) {
	return c.projectmeta.GetProjectMetadataValue(ctx, projectNameOrID, key)
}

func (c *RESTClient) ListProjectMetadata(ctx context.Context, projectNameOrID string) (map[string]string, error) {
	return c.projectmeta.ListProjectMetadata(ctx, projectNameOrID)
}

func (c *RESTClient) UpdateProjectMetadata(ctx context.Context, projectNameOrID string, key common.MetadataKey, value string) error {
	return c.projectmeta.UpdateProjectMetadata(ctx, projectNameOrID, key, value)
}

func (c *RESTClient) DeleteProjectMetadataValue(ctx context.Context, projectNameOrID string, key common.MetadataKey) error {
	return c.projectmeta.DeleteProjectMetadataValue(ctx, projectNameOrID, key)
}

// Quota Client

func (c *RESTClient) ListQuotas(ctx context.Context, referenceType, referenceID *string) ([]*modelv2.Quota, error) {
	return c.quota.ListQuotas(ctx, referenceType, referenceID)
}

func (c *RESTClient) GetQuotaByProjectID(ctx context.Context, projectID int64) (*modelv2.Quota, error) {
	return c.quota.GetQuotaByProjectID(ctx, projectID)
}

func (c *RESTClient) UpdateStorageQuotaByProjectID(ctx context.Context, projectID int64, storageLimit int64) error {
	return c.quota.UpdateStorageQuotaByProjectID(ctx, projectID, storageLimit)
}

// Registry Client

func (c *RESTClient) NewRegistry(ctx context.Context, reg *modelv2.Registry) error {
	return c.registry.NewRegistry(ctx, reg)
}

func (c *RESTClient) GetRegistryByID(ctx context.Context, id int64) (*modelv2.Registry, error) {
	return c.registry.GetRegistryByID(ctx, id)
}

func (c *RESTClient) GetRegistryByName(ctx context.Context, name string) (*modelv2.Registry, error) {
	return c.registry.GetRegistryByName(ctx, name)
}

func (c *RESTClient) ListRegistries(ctx context.Context) ([]*modelv2.Registry, error) {
	return c.registry.ListRegistries(ctx)
}

func (c *RESTClient) DeleteRegistryByID(ctx context.Context, id int64) error {
	return c.registry.DeleteRegistryByID(ctx, id)
}

func (c *RESTClient) UpdateRegistry(ctx context.Context, u *modelv2.RegistryUpdate, id int64) error {
	return c.registry.UpdateRegistry(ctx, u, id)
}

// Replication Client

func (c *RESTClient) NewReplicationPolicy(ctx context.Context, destRegistry, srcRegistry *modelv2.Registry,
	replicateDeletion, override, enablePolicy bool,
	filters []*modelv2.ReplicationFilter, trigger *modelv2.ReplicationTrigger,
	destNamespace, description, name string) error {
	return c.replication.NewReplicationPolicy(ctx, destRegistry, srcRegistry,
		replicateDeletion, override, enablePolicy,
		filters, trigger, destNamespace, description, name)
}

func (c *RESTClient) GetReplicationPolicyByName(ctx context.Context, name string) (*modelv2.ReplicationPolicy, error) {
	return c.replication.GetReplicationPolicyByName(ctx, name)
}

func (c *RESTClient) ListReplicationPolicies(ctx context.Context) ([]*modelv2.ReplicationPolicy, error) {
	return c.replication.ListReplicationPolicies(ctx)
}

func (c *RESTClient) GetReplicationPolicyByID(ctx context.Context, id int64) (*modelv2.ReplicationPolicy, error) {
	return c.replication.GetReplicationPolicyByID(ctx, id)
}

func (c *RESTClient) DeleteReplicationPolicyByID(ctx context.Context, id int64) error {
	return c.replication.DeleteReplicationPolicyByID(ctx, id)
}

func (c *RESTClient) UpdateReplicationPolicy(ctx context.Context, r *modelv2.ReplicationPolicy, id int64) error {
	return c.replication.UpdateReplicationPolicy(ctx, r, id)
}

func (c *RESTClient) TriggerReplicationExecution(ctx context.Context, r *modelv2.StartReplicationExecution) error {
	return c.replication.TriggerReplicationExecution(ctx, r)
}

func (c *RESTClient) ListReplicationExecutions(ctx context.Context, policyID *int64, status, trigger *string) ([]*modelv2.ReplicationExecution, error) {
	return c.replication.ListReplicationExecutions(ctx, policyID, status, trigger)
}

func (c *RESTClient) GetReplicationExecutionByID(ctx context.Context, id int64) (*modelv2.ReplicationExecution, error) {
	return c.replication.GetReplicationExecutionByID(ctx, id)
}

// Repository Client

func (c *RESTClient) GetRepository(ctx context.Context, projectName, repositoryName string) (*modelv2.Repository, error) {
	return c.repository.GetRepository(ctx, projectName, repositoryName)
}

func (c *RESTClient) UpdateRepository(ctx context.Context, projectName, repositoryName string, update *modelv2.Repository) error {
	return c.repository.UpdateRepository(ctx, projectName, repositoryName, update)
}

func (c *RESTClient) ListAllRepositories(ctx context.Context) ([]*modelv2.Repository, error) {
	return c.repository.ListAllRepositories(ctx)
}

func (c *RESTClient) ListRepositories(ctx context.Context, projectName string) ([]*modelv2.Repository, error) {
	return c.repository.ListRepositories(ctx, projectName)
}

func (c *RESTClient) DeleteRepository(ctx context.Context, projectName, repositoryName string) error {
	return c.repository.DeleteRepository(ctx, projectName, repositoryName)
}

// Retention Client

func (c *RESTClient) NewRetentionPolicy(ctx context.Context, ret *modelv2.RetentionPolicy) error {
	return c.retention.NewRetentionPolicy(ctx, ret)
}

func (c *RESTClient) GetRetentionPolicyByProject(ctx context.Context, projectNameOrID string) (*modelv2.RetentionPolicy, error) {
	return c.retention.GetRetentionPolicyByProject(ctx, projectNameOrID)
}

func (c *RESTClient) GetRetentionPolicyByID(ctx context.Context, id int64) (*modelv2.RetentionPolicy, error) {
	return c.retention.GetRetentionPolicyByID(ctx, id)
}

func (c *RESTClient) DeleteRetentionPolicyByID(ctx context.Context, id int64) error {
	return c.retention.DeleteRetentionPolicyByID(ctx, id)
}

func (c *RESTClient) UpdateRetentionPolicy(ctx context.Context, ret *modelv2.RetentionPolicy) error {
	return c.retention.UpdateRetentionPolicy(ctx, ret)
}

// Robot Client

func (c *RESTClient) ListRobotAccounts(ctx context.Context) ([]*modelv2.Robot, error) {
	return c.robot.ListRobotAccounts(ctx)
}

func (c *RESTClient) GetRobotAccountByName(ctx context.Context, name string) (*modelv2.Robot, error) {
	return c.robot.GetRobotAccountByName(ctx, name)
}

func (c *RESTClient) GetRobotAccountByID(ctx context.Context, id int64) (*modelv2.Robot, error) {
	return c.robot.GetRobotAccountByID(ctx, id)
}

func (c *RESTClient) NewRobotAccount(ctx context.Context, r *modelv2.RobotCreate) (*modelv2.RobotCreated, error) {
	return c.robot.NewRobotAccount(ctx, r)
}

func (c *RESTClient) DeleteRobotAccountByName(ctx context.Context, name string) error {
	return c.robot.DeleteRobotAccountByName(ctx, name)
}

func (c *RESTClient) DeleteRobotAccountByID(ctx context.Context, id int64) error {
	return c.robot.DeleteRobotAccountByID(ctx, id)
}

func (c *RESTClient) UpdateRobotAccount(ctx context.Context, r *modelv2.Robot) error {
	return c.robot.UpdateRobotAccount(ctx, r)
}

func (c *RESTClient) RefreshRobotAccountSecretByID(ctx context.Context, id int64, sec string) (*modelv2.RobotSec, error) {
	return c.robot.RefreshRobotAccountSecretByID(ctx, id, sec)
}

func (c *RESTClient) RefreshRobotAccountSecretByName(ctx context.Context, name string, sec string) (*modelv2.RobotSec, error) {
	return c.robot.RefreshRobotAccountSecretByName(ctx, name, sec)
}

// RobotV1 Client

func (c *RESTClient) ListProjectRobotsV1(ctx context.Context, projectNameOrID string) ([]*modelv2.Robot, error) {
	return c.robotv1.ListProjectRobotsV1(ctx, projectNameOrID)
}

func (c *RESTClient) AddProjectRobotV1(ctx context.Context, projectNameOrID string, r *modelv2.RobotCreateV1) error {
	return c.robotv1.AddProjectRobotV1(ctx, projectNameOrID, r)
}

func (c *RESTClient) UpdateProjectRobotV1(ctx context.Context, projectNameOrID string, robotID int64, r *modelv2.Robot) error {
	return c.robotv1.UpdateProjectRobotV1(ctx, projectNameOrID, robotID, r)
}

func (c *RESTClient) DeleteProjectRobotV1(ctx context.Context, projectNameOrID string, robotID int64) error {
	return c.robotv1.DeleteProjectRobotV1(ctx, projectNameOrID, robotID)
}

// Systeminfo Client

func (c *RESTClient) GetSystemInfo(ctx context.Context) (*modelv2.GeneralInfo, error) {
	return c.systeminfo.GetSystemInfo(ctx)
}

// User Client

func (c *RESTClient) NewUser(ctx context.Context, username, email, realname, password, comments string) error {
	return c.user.NewUser(ctx, username, email, realname, password, comments)
}

func (c *RESTClient) GetUserByName(ctx context.Context, username string) (*modelv2.UserResp, error) {
	return c.user.GetUserByName(ctx, username)
}

func (c *RESTClient) GetUserByID(ctx context.Context, id int64) (*modelv2.UserResp, error) {
	return c.user.GetUserByID(ctx, id)
}

func (c *RESTClient) ListUsers(ctx context.Context) ([]*modelv2.UserResp, error) {
	return c.user.ListUsers(ctx)
}

func (c *RESTClient) SearchUsers(ctx context.Context, name string) ([]*modelv2.UserSearchRespItem, error) {
	return c.user.SearchUsers(ctx, name)
}

func (c *RESTClient) GetCurrentUserInfo(ctx context.Context) (*modelv2.UserResp, error) {
	return c.user.GetCurrentUserInfo(ctx)
}

func (c *RESTClient) GetCurrentUserPermisisons(ctx context.Context, relative bool, scope string) ([]*modelv2.Permission, error) {
	return c.user.GetCurrentUserPermisisons(ctx, relative, scope)
}

func (c *RESTClient) SetUserSysAdmin(ctx context.Context, id int64, admin bool) error {
	return c.user.SetUserSysAdmin(ctx, id, admin)
}

func (c *RESTClient) DeleteUser(ctx context.Context, id int64) error {
	return c.user.DeleteUser(ctx, id)
}

func (c *RESTClient) UpdateUserProfile(ctx context.Context, id int64, profile *modelv2.UserProfile) error {
	return c.user.UpdateUserProfile(ctx, id, profile)
}

func (c *RESTClient) UpdateUserPassword(ctx context.Context, userID int64, passwordRequest *modelv2.PasswordReq) error {
	return c.user.UpdateUserPassword(ctx, userID, passwordRequest)
}

func (c *RESTClient) UserExists(ctx context.Context, idOrName intstr.IntOrString) (bool, error) {
	return c.user.UserExists(ctx, idOrName)
}

// Webhook Client

func (c *RESTClient) ListProjectWebhookPolicies(ctx context.Context, projectID int) ([]*modelv2.WebhookPolicy, error) {
	return c.webhook.ListProjectWebhookPolicies(ctx, projectID)
}

func (c *RESTClient) AddProjectWebhookPolicy(ctx context.Context, projectID int, policy *modelv2.WebhookPolicy) error {
	return c.webhook.AddProjectWebhookPolicy(ctx, projectID, policy)
}

func (c *RESTClient) UpdateProjectWebhookPolicy(ctx context.Context, projectID int, policy *modelv2.WebhookPolicy) error {
	return c.webhook.UpdateProjectWebhookPolicy(ctx, projectID, policy)
}

func (c *RESTClient) DeleteProjectWebhookPolicy(ctx context.Context, projectID int, policyID int64) error {
	return c.webhook.DeleteProjectWebhookPolicy(ctx, projectID, policyID)
}
