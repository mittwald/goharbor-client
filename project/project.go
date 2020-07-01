package project

import (
	"context"
	"errors"

	uc "github.com/mittwald/goharbor-client/user"

	"github.com/mittwald/goharbor-client/internal/api/v1.10.0/client"

	"github.com/go-openapi/runtime"
	"github.com/mittwald/goharbor-client/internal/api/v1.10.0/client/products"
	"github.com/mittwald/goharbor-client/internal/api/v1.10.0/model"
)

// RESTClient is a subclient forhandling project related actions.
type RESTClient struct {
	// The swagger client
	Client *client.Harbor

	// AuthInfo contain auth information, which are provided on API calls.
	AuthInfo runtime.ClientAuthInfoWriter
}

func NewClient(cl *client.Harbor, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		Client:   cl,
		AuthInfo: authInfo,
	}
}

type Client interface {
	NewProject(ctx context.Context, name string, countLimit int, storageLimit int)
	DeleteProject(ctx context.Context, p *model.Project) error
	GetProject(ctx context.Context, name string) (*model.Project, error)
	ListProjects(ctx context.Context, nameFilter string) ([]*model.Project, error)
	UpdateProject(ctx context.Context, p *model.Project, countLimit int, storageLimit int) error

	AddProjectMember(ctx context.Context, p *model.Project, u *model.User, roleID int) error
	ListProjectMembers(ctx context.Context, p *model.Project) ([]*model.ProjectMemberEntity, error)
	UpdateProjectMemberRole(ctx context.Context, p *model.Project, u *model.User, roleID int) error
	DeleteProjectMember(ctx context.Context, p *model.Project, u *model.User) error

	AddProjectMetadata(ctx context.Context, p *model.Project, key ProjectMetadataKey, value string) error
	ListProjectMetadata(ctx context.Context, p *model.Project) (*model.ProjectMetadata, error)
	GetProjectMetadataValue(ctx context.Context, p *model.Project, key ProjectMetadataKey) (string, error)
	UpdateProjectMetadata(ctx context.Context, p *model.Project, key ProjectMetadataKey, value string) error
	DeleteProjectMetadataValue(ctx context.Context, p *model.Project, key ProjectMetadataKey) error
}

type ProjectMetadataKey string

const (
	EnableContentTrustProjectMetadataKey   ProjectMetadataKey = "enable_content_trust"
	AutoScanProjectMetadataKey             ProjectMetadataKey = "auto_scan"
	SeverityProjectMetadataKey             ProjectMetadataKey = "severity"
	ReuseSysCVEWhitelistProjectMetadataKey ProjectMetadataKey = "reuse_sys_cve_whitelist"
	PublicProjectMetadataKey               ProjectMetadataKey = "public"
	PreventVulProjectMetadataKey           ProjectMetadataKey = "prevent_vul"
)

// NewProject creates a new project with name as the project's name.
// Returns the project as it is stored inside Harbor or an error,
// if the project could not be created.
// CountLimit limits the number of repositories for this project.
// StorageLimit limits the allocatable space for this project.
func (c *RESTClient) NewProject(ctx context.Context, name string,
	countLimit int, storageLimit int) (*model.Project, error) {
	pReq := &model.ProjectReq{
		CveWhitelist: nil,
		Metadata:     nil,
		ProjectName:  name,
		CountLimit:   int64(countLimit),
		StorageLimit: int64(storageLimit) * 1024 * 1024,
	}

	_, err := c.Client.Products.PostProjects(
		&products.PostProjectsParams{
			Project: pReq,
			Context: ctx,
		}, c.AuthInfo)

	err = handleSwaggerProjectErrors(err)
	if err != nil {
		return nil, err
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
func (c *RESTClient) DeleteProject(ctx context.Context,
	p *model.Project) error {
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

	_, err = c.Client.Products.DeleteProjectsProjectID(
		&products.DeleteProjectsProjectIDParams{
			ProjectID: int64(p.ProjectID),
			Context:   ctx,
		}, c.AuthInfo)

	return handleSwaggerProjectErrors(err)
}

// GetProject returns an existing project identified by name.
// Returns an error if it cannot find a matching project or when
// having difficulties talking to the API.
func (c *RESTClient) GetProject(ctx context.Context,
	name string) (*model.Project, error) {
	if name == "" {
		return nil, errors.New("no name provided")
	}
	resp, err := c.Client.Products.GetProjects(
		&products.GetProjectsParams{
			Name:    &name,
			Context: ctx,
		}, c.AuthInfo)

	err = handleSwaggerProjectErrors(err)
	if err != nil {
		return nil, err
	}

	for _, p := range resp.Payload {
		if p.Name == name {
			return p, nil
		}
	}

	return nil, &ErrProjectNotFound{}
}

// ListProjects returns a list of projects based on a name filter.
// Returns all projects if name is an empty string.
// Returns an error if no projects were found.
func (c *RESTClient) ListProjects(ctx context.Context,
	nameFilter string) ([]*model.Project, error) {
	resp, err := c.Client.Products.GetProjects(
		&products.GetProjectsParams{
			Name:    &nameFilter,
			Context: ctx,
		}, c.AuthInfo)

	err = handleSwaggerProjectErrors(err)
	if err != nil {
		return nil, err
	}

	if len(resp.Payload) == 0 {
		return nil, &ErrProjectNotFound{}
	}

	return resp.Payload, nil
}

// UpdateProject updates a project with the specified data.
// Returns an error if name/ID pair of p does not match a stored project.
func (c *RESTClient) UpdateProject(ctx context.Context, p *model.Project,
	countLimit int, storageLimit int) error {
	project, err := c.GetProject(ctx, p.Name)
	if err != nil {
		return err
	}

	if p.ProjectID != project.ProjectID {
		return &ErrProjectMismatch{}
	}

	pReq := &model.ProjectReq{
		CveWhitelist: p.CveWhitelist,
		Metadata:     p.Metadata,
		ProjectName:  p.Name,
		CountLimit:   int64(countLimit),
		StorageLimit: int64(storageLimit) * 1024 * 1024,
	}

	_, err = c.Client.Products.PutProjectsProjectID(
		&products.PutProjectsProjectIDParams{
			Project:   pReq,
			ProjectID: int64(p.ProjectID),
			Context:   ctx,
		}, c.AuthInfo)

	return handleSwaggerProjectErrors(err)
}

// AddProjectMember creates a membership between a user and a project.
func (c *RESTClient) AddProjectMember(ctx context.Context, p *model.Project, u *model.User, roleID int) error {
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

	userClient := uc.NewClient(c.Client, c.AuthInfo)

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

	_, err = c.Client.Products.PostProjectsProjectIDMembers(&products.PostProjectsProjectIDMembersParams{
		ProjectID:     int64(p.ProjectID),
		ProjectMember: m,
		Context:       ctx,
	}, c.AuthInfo)

	return handleSwaggerProjectErrors(err)
}

// ListProjectMembers returns a list of project members.
func (c *RESTClient) ListProjectMembers(ctx context.Context, p *model.Project) ([]*model.ProjectMemberEntity, error) {
	if p == nil {
		return nil, &ErrProjectNotProvided{}
	}

	entityName := ""

	resp, err := c.Client.Products.GetProjectsProjectIDMembers(&products.GetProjectsProjectIDMembersParams{
		Entityname: &entityName,
		ProjectID:  int64(p.ProjectID),
		Context:    ctx,
	}, c.AuthInfo)

	err = handleSwaggerProjectErrors(err)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

// UpdateProjectMemberRole updates the role of a project member.
func (c *RESTClient) UpdateProjectMemberRole(ctx context.Context, p *model.Project, u *model.User, roleID int) error {
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

	_, err = c.Client.Products.PutProjectsProjectIDMembersMid(&products.PutProjectsProjectIDMembersMidParams{
		Mid:       mid,
		ProjectID: int64(p.ProjectID),
		Role:      roleRequest,
		Context:   ctx,
	}, c.AuthInfo)

	return handleSwaggerProjectErrors(err)
}

// DeleteProjectMember deletes the membership between a user and a project.
func (c *RESTClient) DeleteProjectMember(ctx context.Context, p *model.Project, u *model.User) error {
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

	_, err = c.Client.Products.DeleteProjectsProjectIDMembersMid(
		&products.DeleteProjectsProjectIDMembersMidParams{
			Mid:       mid,
			ProjectID: int64(p.ProjectID),
			Context:   ctx,
		}, c.AuthInfo)

	return handleSwaggerProjectErrors(err)
}

// GetMetadataByKey returns a ProjectMetadata object matching
// the provided key and containing the provided value.
func getProjectMetadataByKey(key ProjectMetadataKey, value string) *model.ProjectMetadata {
	var m model.ProjectMetadata
	switch key {
	case EnableContentTrustProjectMetadataKey:
		m.EnableContentTrust = value
	case AutoScanProjectMetadataKey:
		m.AutoScan = value
	case SeverityProjectMetadataKey:
		m.Severity = value
	case ReuseSysCVEWhitelistProjectMetadataKey:
		m.ReuseSysCveWhitelist = value
	case PublicProjectMetadataKey:
		m.Public = value
	case PreventVulProjectMetadataKey:
		m.PreventVul = value
	}
	return &m
}

// AddMetadata adds metadata with a specific key and value to project p.
// See this for more explanation of possible keys and values:
// https://github.com/goharbor/harbor/blob/v1.10.2/api/harbor/swagger.yaml#L4894
func (c *RESTClient) AddProjectMetadata(ctx context.Context, p *model.Project, key ProjectMetadataKey, value string) error {
	if p == nil {
		return &ErrProjectNotProvided{}
	}

	_, err := c.Client.Products.PostProjectsProjectIDMetadatas(&products.PostProjectsProjectIDMetadatasParams{
		Metadata:  getProjectMetadataByKey(key, value),
		ProjectID: int64(p.ProjectID),
		Context:   ctx,
	}, c.AuthInfo)

	err = handleSwaggerProjectErrors(err)
	if err != nil {
		t, ok := err.(*runtime.APIError)
		if ok && t.Code == 409 {
			// Unspecified error, which is returned when an existing metadata key
			// is tried to create a second time.
			return &ErrProjectMetadataAlreadyExists{}
		}
		return err
	}

	return handleSwaggerProjectErrors(err)
}

// GetProjectMetadataValue retrieves metadata with key of project p.
func (c *RESTClient) GetProjectMetadataValue(ctx context.Context, p *model.Project, key ProjectMetadataKey) (string, error) {
	if p == nil {
		return "", &ErrProjectNotProvided{}
	}

	resp, err := c.Client.Products.GetProjectsProjectIDMetadatasMetaName(&products.GetProjectsProjectIDMetadatasMetaNameParams{
		MetaName:  string(key),
		ProjectID: int64(p.ProjectID),
		Context:   ctx,
	}, c.AuthInfo)

	err = handleSwaggerProjectErrors(err)
	if err != nil {
		return "", err
	}

	var result string
	switch key {
	case EnableContentTrustProjectMetadataKey:
		result = resp.Payload.EnableContentTrust
	case AutoScanProjectMetadataKey:
		result = resp.Payload.AutoScan
	case SeverityProjectMetadataKey:
		result = resp.Payload.Severity
	case ReuseSysCVEWhitelistProjectMetadataKey:
		result = resp.Payload.ReuseSysCveWhitelist
	case PublicProjectMetadataKey:
		result = resp.Payload.Public
	case PreventVulProjectMetadataKey:
		result = resp.Payload.PreventVul
	default:
		return "", &ErrProjectInvalidRequest{}
	}

	return result, nil
}

// ListMetadata lists all metadata of a project
func (c *RESTClient) ListProjectMetadata(ctx context.Context, p *model.Project) (*model.ProjectMetadata, error) {
	if p == nil {
		return nil, &ErrProjectNotProvided{}
	}

	resp, err := c.Client.Products.GetProjectsProjectIDMetadatas(&products.GetProjectsProjectIDMetadatasParams{
		ProjectID: int64(p.ProjectID),
		Context:   ctx,
	}, c.AuthInfo)

	err = handleSwaggerProjectErrors(err)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

// UpdateMetadata deletes the specified metadata key, if it exists and re-adds this metadata key with the given value.
// This function works around the faulty behaviour of the corresponding 'Update' endpoint of the Harbor API.
func (c *RESTClient) UpdateProjectMetadata(ctx context.Context, p *model.Project, key ProjectMetadataKey, value string) error {
	if p == nil {
		return &ErrProjectNotProvided{}
	}

	pID := int64(p.ProjectID)
	metaKeyName := string(key)

	_, err := c.Client.Products.GetProjectsProjectIDMetadatasMetaName(&products.GetProjectsProjectIDMetadatasMetaNameParams{
		MetaName:  metaKeyName,
		ProjectID: pID,
		Context:   ctx,
	}, c.AuthInfo)

	if err != nil {
		return handleSwaggerProjectErrors(err)
	}

	_, err = c.Client.Products.DeleteProjectsProjectIDMetadatasMetaName(&products.DeleteProjectsProjectIDMetadatasMetaNameParams{
		MetaName:  metaKeyName,
		ProjectID: pID,
		Context:   ctx,
	}, c.AuthInfo)

	if err != nil {
		return handleSwaggerProjectErrors(err)
	}

	_, err = c.Client.Products.PostProjectsProjectIDMetadatas(&products.PostProjectsProjectIDMetadatasParams{
		Metadata:  getProjectMetadataByKey(key, value),
		ProjectID: pID,
		Context:   ctx,
	}, c.AuthInfo)

	return handleSwaggerProjectErrors(err)
}

// DeleteMetadataValue deletes metadata of project p given by key.
func (c *RESTClient) DeleteProjectMetadataValue(ctx context.Context, p *model.Project, key ProjectMetadataKey) error {
	if p == nil {
		return &ErrProjectNotProvided{}
	}

	_, err := c.Client.Products.DeleteProjectsProjectIDMetadatasMetaName(&products.DeleteProjectsProjectIDMetadatasMetaNameParams{
		MetaName:  string(key),
		ProjectID: int64(p.ProjectID),
		Context:   ctx,
	}, c.AuthInfo)

	return handleSwaggerProjectErrors(err)
}

// handleProjectErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with proper message.
func handleSwaggerProjectErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case 201:
			// Harbor sometimes return 201 instead of 200 despite the swagger spec
			// not declaring it.
			return nil
		case 400:
			return &ErrProjectIllegalIDFormat{}
		case 401:
			return &ErrProjectUnauthorized{}
		case 403:
			return &ErrProjectNoPermission{}
		case 404:
			return &ErrProjectUnknownResource{}
		case 500:
			return &ErrProjectInternalErrors{}
		}
	}

	switch in.(type) {
	case *products.DeleteProjectsProjectIDNotFound:
		return &ErrProjectIDNotExists{}
	case *products.PutProjectsProjectIDNotFound:
		return &ErrProjectIDNotExists{}
	case *products.PostProjectsConflict:
		return &ErrProjectNameAlreadyExists{}
	case *products.PostProjectsProjectIDMembersBadRequest:
		return &ErrProjectInvalidRequest{}
	case *products.PostProjectsProjectIDMetadatasBadRequest:
		return &ErrProjectInvalidRequest{}
	default:
		return in
	}
}

// projectExists returns true, if p matches a project on server side.
// Returns false, if not found.
// Returns an error in case of communication problems.
func (c *RESTClient) projectExists(ctx context.Context, p *model.Project) (bool, error) {
	_, err := c.GetProject(ctx, p.Name)

	if err != nil {
		if _, ok := err.(*ErrProjectNotFound); ok {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}

// getMid returns the member ID of a user u in project p.
// Returns an error, if user is not a member in project or
// in case a communication error has occured.
func (c *RESTClient) getMid(ctx context.Context, p *model.Project, u *model.User) (int64, error) {
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
