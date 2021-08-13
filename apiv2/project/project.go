package project

import (
	"context"
	"errors"
	"strconv"

	projectapi "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/project"
	"github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client/robotv1"

	modelv2 "github.com/mittwald/goharbor-client/v4/apiv2/model"
	uc "github.com/mittwald/goharbor-client/v4/apiv2/user"

	v2client "github.com/mittwald/goharbor-client/v4/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v4/apiv2/internal/legacyapi/client"

	"github.com/go-openapi/runtime"

	"github.com/mittwald/goharbor-client/v4/apiv2/internal/legacyapi/client/products"
	legacymodel "github.com/mittwald/goharbor-client/v4/apiv2/model/legacy"
)

const (
	ProjectMetadataKeyEnableContentTrust   MetadataKey = "enable_content_trust"
	ProjectMetadataKeyAutoScan             MetadataKey = "auto_scan"
	ProjectMetadataKeySeverity             MetadataKey = "severity"
	ProjectMetadataKeyReuseSysCveAllowlist MetadataKey = "reuse_sys_cve_whitelist"
	ProjectMetadataKeyPublic               MetadataKey = "public"
	ProjectMetadataKeyPreventVul           MetadataKey = "prevent_vul"
	ProjectMetadataKeyRetentionID          MetadataKey = "retention_id"
)

// RESTClient is a subclient for handling project related actions.
type RESTClient struct {
	// The legacy swagger client
	LegacyClient *client.Harbor

	// The new client of the harbor v2 API
	V2Client *v2client.Harbor

	// AuthInfo contains the auth information that is provided on API calls.
	AuthInfo runtime.ClientAuthInfoWriter
}

func NewClient(legacyClient *client.Harbor, v2Client *v2client.Harbor, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		LegacyClient: legacyClient,
		V2Client:     v2Client,
		AuthInfo:     authInfo,
	}
}

type Client interface {
	NewProject(ctx context.Context, name string, storageLimit *int64, public *bool) (*modelv2.Project, error)
	DeleteProject(ctx context.Context, p *modelv2.Project) error
	GetProject(ctx context.Context, nameOrID string) (*modelv2.Project, error)
	ListProjects(ctx context.Context, nameFilter string) ([]*modelv2.Project, error)
	UpdateProject(ctx context.Context, p *modelv2.Project, storageLimit *int64) error

	AddProjectMember(ctx context.Context, p *modelv2.Project, u *legacymodel.User, roleID int) error
	ListProjectMembers(ctx context.Context, p *modelv2.Project) ([]*legacymodel.ProjectMemberEntity, error)
	UpdateProjectMemberRole(ctx context.Context, p *modelv2.Project, u *legacymodel.User, roleID int) error
	DeleteProjectMember(ctx context.Context, p *modelv2.Project, u *legacymodel.User) error

	AddProjectMetadata(ctx context.Context, p *modelv2.Project, key MetadataKey, value string) error
	GetProjectMetadataValue(ctx context.Context, projectNameOrID string, key MetadataKey) (string, error)
	ListProjectMetadata(ctx context.Context, p *modelv2.Project) (*modelv2.ProjectMetadata, error)
	UpdateProjectMetadata(ctx context.Context, p *modelv2.Project, key MetadataKey, value string) error
	DeleteProjectMetadataValue(ctx context.Context, p *modelv2.Project, key MetadataKey) error

	ListProjectRobots(ctx context.Context, p *modelv2.Project) ([]*modelv2.Robot, error)
	AddProjectRobot(ctx context.Context, p *modelv2.Project, r *modelv2.RobotCreateV1) (*modelv2.RobotCreated, error)
	UpdateProjectRobot(ctx context.Context, p *modelv2.Project, robotID int64, r *modelv2.Robot) error
	DeleteProjectRobot(ctx context.Context, p *modelv2.Project, robotID int64) error

	ListProjectWebhookPolicies(ctx context.Context, p *modelv2.Project) ([]*legacymodel.WebhookPolicy, error)
	AddProjectWebhookPolicy(ctx context.Context, p *modelv2.Project, policy *legacymodel.WebhookPolicy) error
	UpdateProjectWebhookPolicy(ctx context.Context, p *modelv2.Project, policyID int64, policy *legacymodel.WebhookPolicy) error
	DeleteProjectWebhookPolicy(ctx context.Context, p *modelv2.Project, policyID int64) error
}

type MetadataKey string

// NewProject creates a new project with name as the project's name.
// Returns the project as it is stored inside Harbor or an error,
// if the project could not be created.
// CountLimit limits the number of repositories for this project.
// StorageLimit limits the allocatable space for this project.
func (c *RESTClient) NewProject(ctx context.Context, name string, storageLimit *int64, public *bool) (*modelv2.Project, error) {
	pReq := &modelv2.ProjectReq{
		ProjectName:  name,
		StorageLimit: storageLimit,
		Public:       public,
	}

	_, err := c.V2Client.Project.CreateProject(
		&projectapi.CreateProjectParams{
			Project: pReq,
			Context: ctx,
		}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerProjectErrors(err)
	}

	project, err := c.GetProject(ctx, name)
	if err != nil {
		return nil, err
	}

	return project, nil
}

// DeleteProject deletes the specified project.
// Returns an error when no matching project is found or when
// having difficulties talking to the API.
func (c *RESTClient) DeleteProject(ctx context.Context, p *modelv2.Project) error {
	if p == nil {
		return &ErrProjectNotProvided{}
	}

	projectExists, err := c.projectExists(ctx, p)
	if err != nil {
		return err
	}

	if !projectExists {
		return &ErrProjectMismatch{}
	}

	_, err = c.V2Client.Project.DeleteProject(
		&projectapi.DeleteProjectParams{
			ProjectNameOrID: ProjectIDAsString(p.ProjectID),
			Context:         ctx,
		}, c.AuthInfo)

	return handleSwaggerProjectErrors(err)
}

// GetProject returns an existing project identified by nameOrID.
// nameOrID may contain a unique project name or its unique ID.
// Returns an error if it cannot find a matching project or when
// having difficulties talking to the API.
func (c *RESTClient) GetProject(ctx context.Context, nameOrID string) (*modelv2.Project, error) {
	if nameOrID == "" {
		return nil, &ErrProjectNameNotProvided{}
	}

	resp, err := c.V2Client.Project.GetProject(&projectapi.GetProjectParams{
		ProjectNameOrID: nameOrID,
		Context:         ctx,
	}, c.AuthInfo)
	if err != nil {
		if resp == nil {
			return nil, &ErrProjectNotFound{}
		}
		return nil, handleSwaggerProjectErrors(err)
	}

	return resp.Payload, nil
}

// ListProjects returns a list of projects based on a name filter.
// Returns all projects if name is an empty string.
// Returns an error if no projects were found.
func (c *RESTClient) ListProjects(ctx context.Context, nameFilter string) ([]*modelv2.Project, error) {
	resp, err := c.V2Client.Project.ListProjects(&projectapi.ListProjectsParams{
		Name:    &nameFilter,
		Context: ctx,
	}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerProjectErrors(err)
	}

	if len(resp.Payload) == 0 {
		return nil, &ErrProjectNotFound{}
	}

	return resp.Payload, nil
}

// UpdateProject updates a project with the specified data.
// Returns an error if name/ID pair of p does not match a stored project.
// Note: Only positive values of storageLimit are supported through this method.
// If you want to set an infinite storageLimit (-1),
// please refer to the quota client's 'UpdateStorageQuotaByProjectID' method.
func (c *RESTClient) UpdateProject(ctx context.Context, p *modelv2.Project, storageLimit *int64) error {
	project, err := c.GetProject(ctx, p.Name)
	if err != nil {
		return err
	}

	if p.ProjectID != project.ProjectID {
		return &ErrProjectMismatch{}
	}

	pReq := &modelv2.ProjectReq{
		CveAllowlist: p.CveAllowlist,
		Metadata:     p.Metadata,
		ProjectName:  p.Name,
		StorageLimit: storageLimit,
	}

	_, err = c.V2Client.Project.UpdateProject(&projectapi.UpdateProjectParams{
		Project:         pReq,
		ProjectNameOrID: ProjectIDAsString(p.ProjectID),
		Context:         ctx,
	}, c.AuthInfo)

	return handleSwaggerProjectErrors(err)
}

// AddProjectMember creates a membership between a user and a project.
func (c *RESTClient) AddProjectMember(ctx context.Context, p *modelv2.Project, u *legacymodel.User, roleID int) error {
	if p == nil {
		return &ErrProjectNotProvided{}
	}

	if u == nil {
		return &ErrProjectNoMemberProvided{}
	}
	projectExists, err := c.projectExists(ctx, p)
	if err != nil {
		return err
	}

	if !projectExists {
		return &ErrProjectMismatch{}
	}

	userClient := uc.NewClient(c.LegacyClient, c.V2Client, c.AuthInfo)

	userExists, err := userClient.UserExists(ctx, u)
	if err != nil {
		return err
	}
	if !userExists {
		return &ErrProjectMemberMismatch{}
	}

	m := &legacymodel.ProjectMember{
		RoleID: int64(roleID),
		MemberUser: &legacymodel.UserEntity{
			UserID:   u.UserID,
			Username: u.Username,
		},
		MemberGroup: &legacymodel.UserGroup{},
	}

	_, err = c.LegacyClient.Products.PostProjectsProjectIDMembers(
		&products.PostProjectsProjectIDMembersParams{
			ProjectID:     int64(p.ProjectID),
			ProjectMember: m,
			Context:       ctx,
		}, c.AuthInfo)

	return handleSwaggerProjectErrors(err)
}

// ListProjectMembers returns a list of project members.
func (c *RESTClient) ListProjectMembers(ctx context.Context, p *modelv2.Project) ([]*legacymodel.ProjectMemberEntity, error) {
	if p == nil {
		return nil, &ErrProjectNotProvided{}
	}

	entityName := ""

	resp, err := c.LegacyClient.Products.GetProjectsProjectIDMembers(
		&products.GetProjectsProjectIDMembersParams{
			Entityname: &entityName,
			ProjectID:  int64(p.ProjectID),
			Context:    ctx,
		}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerProjectErrors(err)
	}

	return resp.Payload, nil
}

// UpdateProjectMemberRole updates the role of a project member.
func (c *RESTClient) UpdateProjectMemberRole(ctx context.Context, p *modelv2.Project, u *legacymodel.User, roleID int) error {
	if p == nil {
		return &ErrProjectNotProvided{}
	}
	if u == nil {
		return &ErrProjectNoMemberProvided{}
	}

	projectExists, err := c.projectExists(ctx, p)
	if err != nil {
		return err
	}
	if !projectExists {
		return &ErrProjectMismatch{}
	}

	mid, err := c.getMid(ctx, p, u)
	if err != nil {
		return err
	}

	roleRequest := &legacymodel.RoleRequest{RoleID: int64(roleID)}

	_, err = c.LegacyClient.Products.PutProjectsProjectIDMembersMid(
		&products.PutProjectsProjectIDMembersMidParams{
			Mid:       mid,
			ProjectID: int64(p.ProjectID),
			Role:      roleRequest,
			Context:   ctx,
		}, c.AuthInfo)

	return handleSwaggerProjectErrors(err)
}

// DeleteProjectMember deletes the membership between a user and a project.
func (c *RESTClient) DeleteProjectMember(ctx context.Context, p *modelv2.Project, u *legacymodel.User) error {
	if p == nil {
		return &ErrProjectNotProvided{}
	}
	if u == nil {
		return &ErrProjectNoMemberProvided{}
	}

	projectExists, err := c.projectExists(ctx, p)
	if err != nil {
		return err
	}
	if !projectExists {
		return &ErrProjectMismatch{}
	}

	mid, err := c.getMid(ctx, p, u)
	if err != nil {
		return err
	}

	_, err = c.LegacyClient.Products.DeleteProjectsProjectIDMembersMid(
		&products.DeleteProjectsProjectIDMembersMidParams{
			Mid:       mid,
			ProjectID: int64(p.ProjectID),
			Context:   ctx,
		}, c.AuthInfo)

	return handleSwaggerProjectErrors(err)
}

// getProjectMetadataByKey returns a ProjectMetadata object matching
// the provided key and containing the provided value.
func getProjectMetadataByKey(key MetadataKey, value string) *modelv2.ProjectMetadata {
	var m modelv2.ProjectMetadata

	switch key {
	case ProjectMetadataKeyEnableContentTrust:
		m.EnableContentTrust = &value
	case ProjectMetadataKeyAutoScan:
		m.AutoScan = &value
	case ProjectMetadataKeySeverity:
		m.Severity = &value
	case ProjectMetadataKeyReuseSysCveAllowlist:
		m.ReuseSysCveAllowlist = &value
	case ProjectMetadataKeyPublic:
		m.Public = value
	case ProjectMetadataKeyPreventVul:
		m.PreventVul = &value
	case ProjectMetadataKeyRetentionID:
		m.RetentionID = &value
	}

	return &m
}

// AddProjectMetadata AddMetadata adds metadata with a specific key and value to project p.
// See this for more explanation of possible keys and values:
// https://github.com/goharbor/harbor/blob/v1.10.2/api/harbor/swagger.yaml#L4894
func (c *RESTClient) AddProjectMetadata(ctx context.Context, p *modelv2.Project, key MetadataKey, value string) error {
	if p == nil {
		return &ErrProjectNotProvided{}
	}

	meta := getProjectMetadataByKey(key, value)

	_, err := c.V2Client.Project.UpdateProject(&projectapi.UpdateProjectParams{
		ProjectNameOrID: ProjectIDAsString(p.ProjectID),
		Project: &modelv2.ProjectReq{
			Metadata:    meta,
			ProjectName: p.Name,
		},
		Context: ctx,
	}, c.AuthInfo)

	err = handleSwaggerProjectErrors(err)
	if err != nil {
		t, ok := err.(*runtime.APIError)
		if ok && t.Code == 409 {
			// Unspecified error that returns when a metadata key is already defined.
			return &ErrProjectMetadataAlreadyExists{}
		}

		return err
	}

	return handleSwaggerProjectErrors(err)
}

// GetProjectMetadataValue retrieves metadata with key of project p.
func (c *RESTClient) GetProjectMetadataValue(ctx context.Context, projectNameOrID string, key MetadataKey) (string, error) {
	project, err := c.GetProject(ctx, projectNameOrID)
	if err != nil {
		return "", handleSwaggerProjectErrors(err)
	}

	if project.Metadata == nil {
		return "", &ErrProjectMetadataUndefined{}
	}

	return retrieveMetadataValue(key, project.Metadata)
}

// ListProjectMetadata ListMetadata lists all metadata of a project
func (c *RESTClient) ListProjectMetadata(ctx context.Context, p *modelv2.Project) (*modelv2.ProjectMetadata, error) {
	if p == nil {
		return nil, &ErrProjectNotProvided{}
	}

	resp, err := c.GetProject(ctx, ProjectIDAsString(p.ProjectID))
	if err != nil {
		return nil, handleSwaggerProjectErrors(err)
	}

	if resp.Metadata != nil {
		return resp.Metadata, nil
	}

	return nil, &ErrProjectMetadataAlreadyExists{}
}

// UpdateProjectMetadata UpdateMetadata deletes the specified metadata key, if it exists and re-adds this metadata key with the given value.
// This function works around the faulty behaviour of the corresponding 'Update' endpoint of the Harbor API.
func (c *RESTClient) UpdateProjectMetadata(ctx context.Context, p *modelv2.Project, key MetadataKey, value string) error {
	if p == nil {
		return &ErrProjectNotProvided{}
	}

	metaKeyName := string(key)

	_, err := c.LegacyClient.Products.GetProjectsProjectIDMetadatasMetaName(
		&products.GetProjectsProjectIDMetadatasMetaNameParams{
			MetaName:  metaKeyName,
			ProjectID: int64(p.ProjectID),
			Context:   ctx,
		}, c.AuthInfo)
	if err != nil {
		return handleSwaggerProjectErrors(err)
	}

	_, err = c.LegacyClient.Products.DeleteProjectsProjectIDMetadatasMetaName(
		&products.DeleteProjectsProjectIDMetadatasMetaNameParams{
			MetaName:  metaKeyName,
			ProjectID: int64(p.ProjectID),
			Context:   ctx,
		}, c.AuthInfo)

	if err != nil {
		return handleSwaggerProjectErrors(err)
	}

	meta := getProjectMetadataByKey(key, value)

	_, err = c.V2Client.Project.UpdateProject(&projectapi.UpdateProjectParams{
		Project: &modelv2.ProjectReq{
			Metadata:    meta,
			ProjectName: p.Name,
		},
		ProjectNameOrID: ProjectIDAsString(p.ProjectID),
		Context:         ctx,
	}, c.AuthInfo)

	return handleSwaggerProjectErrors(err)
}

// DeleteProjectMetadataValue DeleteMetadataValue deletes metadata of project p given by key.
func (c *RESTClient) DeleteProjectMetadataValue(ctx context.Context, p *modelv2.Project, key MetadataKey) error {
	if p == nil {
		return &ErrProjectNotProvided{}
	}

	_, err := c.LegacyClient.Products.DeleteProjectsProjectIDMetadatasMetaName(
		&products.DeleteProjectsProjectIDMetadatasMetaNameParams{
			MetaName:  string(key),
			ProjectID: int64(p.ProjectID),
			Context:   ctx,
		}, c.AuthInfo)

	return handleSwaggerProjectErrors(err)
}

// ListProjectRobots returns a list of all robot accounts in project p.
func (c *RESTClient) ListProjectRobots(ctx context.Context, p *modelv2.Project) ([]*modelv2.Robot, error) {
	if p == nil {
		return nil, &ErrProjectNotProvided{}
	}

	resp, err := c.V2Client.Robotv1.ListRobotV1(&robotv1.ListRobotV1Params{
		ProjectNameOrID: ProjectIDAsString(p.ProjectID),
		Context:         ctx,
	}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerProjectErrors(err)
	}

	return resp.Payload, nil
}

// AddProjectRobot creates the robot account 'r' and adds it to the project 'p'.
// and returns a 'RobotCreated' response.
func (c *RESTClient) AddProjectRobot(ctx context.Context, p *modelv2.Project, r *modelv2.RobotCreateV1) (*modelv2.RobotCreated, error) {
	if p == nil {
		return nil, &ErrProjectNotProvided{}
	}

	resp, err := c.V2Client.Robotv1.CreateRobotV1(&robotv1.CreateRobotV1Params{
		ProjectNameOrID: ProjectIDAsString(p.ProjectID),
		Robot:           r,
		Context:         ctx,
	}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerProjectErrors(err)
	}

	return resp.Payload, nil
}

// UpdateProjectRobot updates a robot account 'r' in project 'p' using the 'robotID'.
func (c *RESTClient) UpdateProjectRobot(ctx context.Context, p *modelv2.Project, robotID int64, r *modelv2.Robot) error {
	if p == nil {
		return &ErrProjectNotProvided{}
	}

	_, err := c.V2Client.Robotv1.UpdateRobotV1(&robotv1.UpdateRobotV1Params{
		ProjectNameOrID: ProjectIDAsString(p.ProjectID),
		Robot:           r,
		RobotID:         robotID,
		Context:         ctx,
	}, c.AuthInfo)
	if err != nil {
		return handleSwaggerProjectErrors(err)
	}

	return nil
}

// DeleteProjectRobot deletes a robot account from project p.
func (c *RESTClient) DeleteProjectRobot(ctx context.Context, p *modelv2.Project, robotID int64) error {
	if p == nil {
		return &ErrProjectNotProvided{}
	}

	_, err := c.V2Client.Robotv1.DeleteRobotV1(&robotv1.DeleteRobotV1Params{
		ProjectNameOrID: ProjectIDAsString(p.ProjectID),
		RobotID:         robotID,
		Context:         ctx,
	}, c.AuthInfo)
	if err != nil {
		return handleSwaggerProjectErrors(err)
	}

	return nil
}

// ListProjectWebhookPolicies returns a list of all webhook policies in project p.
func (c *RESTClient) ListProjectWebhookPolicies(ctx context.Context, p *modelv2.Project) ([]*legacymodel.WebhookPolicy, error) {
	if p == nil {
		return nil, &ErrProjectNotProvided{}
	}

	resp, err := c.LegacyClient.Products.GetProjectsProjectIDWebhookPolicies(
		&products.GetProjectsProjectIDWebhookPoliciesParams{
			ProjectID: int64(p.ProjectID),
			Context:   ctx,
		}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerProjectErrors(err)
	}

	return resp.Payload, nil
}

// AddProjectWebhookPolicy adds a webhook policy to project p.
func (c *RESTClient) AddProjectWebhookPolicy(ctx context.Context, p *modelv2.Project, policy *legacymodel.WebhookPolicy) error {
	if p == nil {
		return &ErrProjectNotProvided{}
	}

	if policy == nil {
		return &ErrProjectNoWebhookPolicyProvided{}
	}

	_, err := c.LegacyClient.Products.PostProjectsProjectIDWebhookPolicies(
		&products.PostProjectsProjectIDWebhookPoliciesParams{
			Policy:    policy,
			ProjectID: int64(p.ProjectID),
			Context:   ctx,
		}, c.AuthInfo)
	return handleSwaggerProjectErrors(err)
}

// UpdateProjectWebhookPolicy updates a webhook policy in project p.
func (c *RESTClient) UpdateProjectWebhookPolicy(ctx context.Context, p *modelv2.Project, policyID int64, policy *legacymodel.WebhookPolicy) error {
	if p == nil {
		return &ErrProjectNotProvided{}
	}

	if policy == nil {
		return &ErrProjectNoWebhookPolicyProvided{}
	}

	_, err := c.LegacyClient.Products.PutProjectsProjectIDWebhookPoliciesPolicyID(
		&products.PutProjectsProjectIDWebhookPoliciesPolicyIDParams{
			ProjectID: int64(p.ProjectID),
			Policy:    policy,
			PolicyID:  policyID,
			Context:   ctx,
		}, c.AuthInfo)
	if err != nil {
		return handleSwaggerProjectErrors(err)
	}

	return nil
}

// DeleteProjectWebhookPolicy deletes a webhook policy from project p.
func (c *RESTClient) DeleteProjectWebhookPolicy(ctx context.Context, p *modelv2.Project, policyID int64) error {
	if p == nil {
		return &ErrProjectNotProvided{}
	}

	_, err := c.LegacyClient.Products.DeleteProjectsProjectIDWebhookPoliciesPolicyID(
		&products.DeleteProjectsProjectIDWebhookPoliciesPolicyIDParams{
			ProjectID: int64(p.ProjectID),
			PolicyID:  policyID,
			Context:   ctx,
		}, c.AuthInfo)
	if err != nil {
		return handleSwaggerProjectErrors(err)
	}

	return nil
}

// projectExists returns true, if p matches a project on server side.
// Returns false, if not found.
// Returns an error in case of communication problems.
func (c *RESTClient) projectExists(ctx context.Context, p *modelv2.Project) (bool, error) {
	_, err := c.GetProject(ctx, p.Name)
	if err != nil {
		if errors.Is(err, &ErrProjectNotFound{}) {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}

// getMid returns the member ID of a user u in project p.
// Returns an error, if user is not a member in project or
// in case a communication error has occurred.
func (c *RESTClient) getMid(ctx context.Context, p *modelv2.Project, u *legacymodel.User) (int64, error) {
	members, err := c.ListProjectMembers(ctx, p)
	if err != nil {
		return 0, err
	}

	for _, v := range members {
		if v.EntityType == "u" && v.EntityName == u.Username {
			return v.ID, nil
		}
	}

	return 0, &ErrProjectUserIsNoMember{}
}

func ProjectIDAsString(projectID int32) string {
	return strconv.Itoa(int(projectID))
}
