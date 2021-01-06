package project

import (
	"context"
	"errors"

	projectapi "github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client/project"

	modelv2 "github.com/mittwald/goharbor-client/v3/apiv2/model"
	uc "github.com/mittwald/goharbor-client/v3/apiv2/user"

	v2client "github.com/mittwald/goharbor-client/v3/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/legacyapi/client"

	"github.com/go-openapi/runtime"
	"github.com/mittwald/goharbor-client/v3/apiv2/internal/legacyapi/client/products"
	model "github.com/mittwald/goharbor-client/v3/apiv2/model/legacy"
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
	NewProject(ctx context.Context, name string, storageLimit int) (*modelv2.Project, error)
	DeleteProject(ctx context.Context, p *modelv2.Project) error
	GetProjectByName(ctx context.Context, name string) (*modelv2.Project, error)
	GetProjectByID(ctx context.Context, projectID int64) (*modelv2.Project, error)
	ListProjects(ctx context.Context, nameFilter string) ([]*modelv2.Project, error)
	UpdateProject(ctx context.Context, p *modelv2.Project, storageLimit int) error

	AddProjectMember(ctx context.Context, p *modelv2.Project, u *model.User, roleID int) error
	ListProjectMembers(ctx context.Context, p *modelv2.Project) ([]*model.ProjectMemberEntity, error)
	UpdateProjectMemberRole(ctx context.Context, p *modelv2.Project, u *model.User, roleID int) error
	DeleteProjectMember(ctx context.Context, p *modelv2.Project, u *model.User) error

	AddProjectMetadata(ctx context.Context, p *modelv2.Project, key MetadataKey, value string) error
	ListProjectMetadata(ctx context.Context, p *modelv2.Project) (*model.ProjectMetadata, error)
	GetProjectMetadataValue(ctx context.Context, projectID int64, key MetadataKey) (string, error)
	UpdateProjectMetadata(ctx context.Context, p *modelv2.Project, key MetadataKey, value string) error
	DeleteProjectMetadataValue(ctx context.Context, p *modelv2.Project, key MetadataKey) error
}

type MetadataKey string

// NewProject creates a new project with name as the project's name.
// Returns the project as it is stored inside Harbor or an error,
// if the project could not be created.
// CountLimit limits the number of repositories for this project.
// StorageLimit limits the allocatable space for this project.
func (c *RESTClient) NewProject(ctx context.Context, name string, storageLimit int) (*modelv2.Project, error) {
	var sPtr = int64(storageLimit) * 1024 * 1024

	pReq := &modelv2.ProjectReq{
		ProjectName:  name,
		StorageLimit: &sPtr,
	}

	_, err := c.V2Client.Project.CreateProject(
		&projectapi.CreateProjectParams{
			Project: pReq,
			Context: ctx,
		}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerProjectErrors(err)
	}

	project, err := c.GetProjectByName(ctx, name)
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
			ProjectID: int64(p.ProjectID),
			Context:   ctx,
		}, c.AuthInfo)

	return handleSwaggerProjectErrors(err)
}

// GetProjectByName returns an existing project identified by name.
// Returns an error if it cannot find a matching project or when
// having difficulties talking to the API.
func (c *RESTClient) GetProjectByName(ctx context.Context, name string) (*modelv2.Project, error) {
	if name == "" {
		return nil, &ErrProjectNameNotProvided{}
	}

	projectList, err := c.ListProjects(ctx, name)

	if err != nil {
		return nil, handleSwaggerProjectErrors(err)
	}

	var projectID int64

	for _, p := range projectList {
		if p.Name == name {
			projectID = int64(p.ProjectID)
		}
	}

	resp, err := c.V2Client.Project.GetProject(&projectapi.GetProjectParams{
		ProjectID: projectID,
		Context:   ctx,
	}, c.AuthInfo)

	if err != nil {
		return nil, handleSwaggerProjectErrors(err)
	}

	if resp.Payload != nil {
		return resp.Payload, nil
	}

	return nil, &ErrProjectNotFound{}
}

// GetProjectByID returns a project identified by its ID.
func (c *RESTClient) GetProjectByID(ctx context.Context, projectID int64) (*modelv2.Project, error) {
	resp, err := c.V2Client.Project.GetProject(&projectapi.GetProjectParams{
		ProjectID: projectID,
		Context:   ctx,
	}, c.AuthInfo)

	if err != nil {
		return nil, handleSwaggerProjectErrors(err)
	}

	if resp != nil {
		return resp.Payload, nil
	}

	return nil, &ErrProjectNotFound{}
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
func (c *RESTClient) UpdateProject(ctx context.Context, p *modelv2.Project, storageLimit int) error {
	project, err := c.GetProjectByName(ctx, p.Name)
	if err != nil {
		return err
	}

	if p.ProjectID != project.ProjectID {
		return &ErrProjectMismatch{}
	}

	var sPtr = int64(storageLimit) * 1024 * 1024

	pReq := &modelv2.ProjectReq{
		CveAllowlist: p.CveAllowlist,
		Metadata:     p.Metadata,
		ProjectName:  p.Name,
		StorageLimit: &sPtr,
	}

	_, err = c.V2Client.Project.UpdateProject(&projectapi.UpdateProjectParams{
		Project:   pReq,
		ProjectID: int64(p.ProjectID),
		Context:   ctx,
	}, c.AuthInfo)

	return handleSwaggerProjectErrors(err)
}

// AddProjectMember creates a membership between a user and a project.
func (c *RESTClient) AddProjectMember(ctx context.Context, p *modelv2.Project, u *model.User, roleID int) error {
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

	m := &model.ProjectMember{
		RoleID: int64(roleID),
		MemberUser: &model.UserEntity{
			UserID:   u.UserID,
			Username: u.Username,
		},
		MemberGroup: &model.UserGroup{},
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
func (c *RESTClient) ListProjectMembers(ctx context.Context, p *modelv2.Project) ([]*model.ProjectMemberEntity, error) {
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
func (c *RESTClient) UpdateProjectMemberRole(ctx context.Context, p *modelv2.Project, u *model.User, roleID int) error {
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

	roleRequest := &model.RoleRequest{RoleID: int64(roleID)}

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
func (c *RESTClient) DeleteProjectMember(ctx context.Context, p *modelv2.Project, u *model.User) error {
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

// AddMetadata adds metadata with a specific key and value to project p.
// See this for more explanation of possible keys and values:
// https://github.com/goharbor/harbor/blob/v1.10.2/api/harbor/swagger.yaml#L4894
func (c *RESTClient) AddProjectMetadata(ctx context.Context, p *modelv2.Project, key MetadataKey, value string) error {
	if p == nil {
		return &ErrProjectNotProvided{}
	}

	meta := getProjectMetadataByKey(key, value)

	_, err := c.V2Client.Project.UpdateProject(&projectapi.UpdateProjectParams{
		ProjectID: int64(p.ProjectID),
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
func (c *RESTClient) GetProjectMetadataValue(ctx context.Context, projectID int64, key MetadataKey) (string, error) {
	project, err := c.GetProjectByID(ctx, projectID)

	if err != nil {
		return "", handleSwaggerProjectErrors(err)
	}

	var result string

	switch key {
	case ProjectMetadataKeyEnableContentTrust:
		result = *project.Metadata.EnableContentTrust
	case ProjectMetadataKeyAutoScan:
		result = *project.Metadata.AutoScan
	case ProjectMetadataKeySeverity:
		result = *project.Metadata.Severity
	case ProjectMetadataKeyReuseSysCveAllowlist:
		result = *project.Metadata.ReuseSysCveAllowlist
	case ProjectMetadataKeyPublic:
		result = project.Metadata.Public
	case ProjectMetadataKeyPreventVul:
		result = *project.Metadata.PreventVul
	case ProjectMetadataKeyRetentionID:
		result = *project.Metadata.RetentionID
	default:
		return "", &ErrProjectInvalidRequest{}
	}

	return result, nil
}

// ListMetadata lists all metadata of a project
func (c *RESTClient) ListProjectMetadata(ctx context.Context, p *modelv2.Project) (*modelv2.ProjectMetadata, error) {
	if p == nil {
		return nil, &ErrProjectNotProvided{}
	}

	resp, err := c.GetProjectByID(ctx, int64(p.ProjectID))
	if err != nil {
		return nil, handleSwaggerProjectErrors(err)
	}

	if resp.Metadata != nil {
		return resp.Metadata, nil
	}

	return nil, &ErrProjectMetadataAlreadyExists{}
}

// UpdateMetadata deletes the specified metadata key, if it exists and re-adds this metadata key with the given value.
// This function works around the faulty behaviour of the corresponding 'Update' endpoint of the Harbor API.
func (c *RESTClient) UpdateProjectMetadata(ctx context.Context, p *modelv2.Project, key MetadataKey, value string) error {
	if p == nil {
		return &ErrProjectNotProvided{}
	}

	pID := int64(p.ProjectID)
	metaKeyName := string(key)

	_, err := c.LegacyClient.Products.GetProjectsProjectIDMetadatasMetaName(
		&products.GetProjectsProjectIDMetadatasMetaNameParams{
			MetaName:  metaKeyName,
			ProjectID: pID,
			Context:   ctx,
		}, c.AuthInfo)
	if err != nil {
		return handleSwaggerProjectErrors(err)
	}

	_, err = c.LegacyClient.Products.DeleteProjectsProjectIDMetadatasMetaName(
		&products.DeleteProjectsProjectIDMetadatasMetaNameParams{
			MetaName:  metaKeyName,
			ProjectID: pID,
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
		ProjectID: pID,
		Context:   ctx,
	}, c.AuthInfo)

	return handleSwaggerProjectErrors(err)
}

// DeleteMetadataValue deletes metadata of project p given by key.
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
func (c *RESTClient) ListProjectRobots(ctx context.Context, p *modelv2.Project) ([]*model.RobotAccount, error) {
	if p == nil {
		return nil, &ErrProjectNotProvided{}
	}

	resp, err := c.LegacyClient.Products.GetProjectsProjectIDRobots(
		&products.GetProjectsProjectIDRobotsParams{
			ProjectID: int64(p.ProjectID),
			Context:   ctx,
		}, c.AuthInfo)
	if err != nil {
		return nil, handleSwaggerProjectErrors(err)
	}

	return resp.Payload, nil
}

// AddProjectRobot adds a robot account to project p and returns the token.
func (c *RESTClient) AddProjectRobot(ctx context.Context, p *modelv2.Project, robot *model.RobotAccountCreate) (string, error) {
	if p == nil {
		return "", &ErrProjectNotProvided{}
	}

	resp, err := c.LegacyClient.Products.PostProjectsProjectIDRobots(
		&products.PostProjectsProjectIDRobotsParams{
			Robot:     robot,
			ProjectID: int64(p.ProjectID),
			Context:   ctx,
		}, c.AuthInfo)
	if err != nil {
		return "", handleSwaggerProjectErrors(err)
	}

	return resp.Payload.Token, nil
}

// UpdateProjectRobot updates a robot account in project p.
func (c *RESTClient) UpdateProjectRobot(ctx context.Context, p *modelv2.Project, robotID int, robot *model.RobotAccountUpdate) error {
	if p == nil {
		return &ErrProjectNotProvided{}
	}

	_, err := c.LegacyClient.Products.PutProjectsProjectIDRobotsRobotID(
		&products.PutProjectsProjectIDRobotsRobotIDParams{
			ProjectID: int64(p.ProjectID),
			Robot:     robot,
			RobotID:   int64(robotID),
			Context:   ctx,
		}, c.AuthInfo)
	if err != nil {
		return handleSwaggerProjectErrors(err)
	}

	return nil
}

// DeleteProjectRobot deletes a robot account from project p.
func (c *RESTClient) DeleteProjectRobot(ctx context.Context, p *modelv2.Project, robotID int) error {
	if p == nil {
		return &ErrProjectNotProvided{}
	}

	_, err := c.LegacyClient.Products.DeleteProjectsProjectIDRobotsRobotID(
		&products.DeleteProjectsProjectIDRobotsRobotIDParams{
			ProjectID: int64(p.ProjectID),
			RobotID:   int64(robotID),
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
	_, err := c.GetProjectByName(ctx, p.Name)
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
func (c *RESTClient) getMid(ctx context.Context, p *modelv2.Project, u *model.User) (int64, error) {
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
