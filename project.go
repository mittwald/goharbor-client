package goharborclient

import (
	"context"
	"errors"

	"github.com/go-openapi/runtime"
	"github.com/mittwald/goharbor-client/api/v1.10.0/client/products"
	"github.com/mittwald/goharbor-client/api/v1.10.0/model"
)

// ProjectRESTClient is a subclient for RESTClient handling project related
// actions.
type ProjectRESTClient struct {
	parent *RESTClient
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

// NewProject creates a new project with name as project name.
// CountLimit and StorageLimit limits space and access for this project.
// Returns the project as it is stored inside Harbor or an error,
// if the project could not be created.
func (c *ProjectRESTClient) NewProject(ctx context.Context, name string,
	countLimit int, storageLimit int) (*model.Project, error) {
	pReq := &model.ProjectReq{
		CveWhitelist: nil,
		Metadata:     nil,
		ProjectName:  name,
		CountLimit:   int64(countLimit),
		StorageLimit: int64(storageLimit) * 1024 * 1024,
	}

	_, err := c.parent.Client.Products.PostProjects(
		&products.PostProjectsParams{
			Project: pReq,
			Context: ctx,
		}, c.parent.AuthInfo)

	err = handleSwaggerProjectErrors(err)
	if err != nil {
		return nil, err
	}

	project, err := c.Get(ctx, name)
	if err != nil {
		return nil, err
	}

	return project, nil
}

// Delete deletes a project.
// Returns an error when no matching project is found or when
// having difficulties talking to the API.
func (c *ProjectRESTClient) Delete(ctx context.Context,
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

	_, err = c.parent.Client.Products.DeleteProjectsProjectID(
		&products.DeleteProjectsProjectIDParams{
			ProjectID: int64(p.ProjectID),
			Context:   ctx,
		}, c.parent.AuthInfo)

	return handleSwaggerProjectErrors(err)
}

// Get returns a project identified by name.
// Returns an error if it cannot find a matching project or when
// having difficulties talking to the API.
func (c *ProjectRESTClient) Get(ctx context.Context,
	name string) (*model.Project, error) {
	if name == "" {
		return nil, errors.New("no name provided")
	}
	resp, err := c.parent.Client.Products.GetProjects(
		&products.GetProjectsParams{
			Name:    &name,
			Context: ctx,
		}, c.parent.AuthInfo)

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

// List retrieves projects filtered by name (all if name is empty string).
// Returns an error if no projects were found.
func (c *ProjectRESTClient) List(ctx context.Context,
	nameFilter string) ([]*model.Project, error) {
	resp, err := c.parent.Client.Products.GetProjects(
		&products.GetProjectsParams{
			Name:    &nameFilter,
			Context: ctx,
		}, c.parent.AuthInfo)

	err = handleSwaggerProjectErrors(err)
	if err != nil {
		return nil, err
	}

	if len(resp.Payload) == 0 {
		return nil, &ErrProjectNotFound{}
	}

	return resp.Payload, nil
}

// Update overwrites properties of a stored project with properties of p.Update.
// Return an error if Name/ID pair of p does not match a stored project.
func (c *ProjectRESTClient) Update(ctx context.Context, p *model.Project,
	countLimit int, storageLimit int) error {
	project, err := c.Get(ctx, p.Name)
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

	_, err = c.parent.Client.Products.PutProjectsProjectID(
		&products.PutProjectsProjectIDParams{
			Project:   pReq,
			ProjectID: int64(p.ProjectID),
			Context:   ctx,
		}, c.parent.AuthInfo)

	return handleSwaggerProjectErrors(err)
}

// AddUserMember adds an existing user to a project with a role identified by roleID.
func (c *ProjectRESTClient) AddUserMember(ctx context.Context, p *model.Project, u *model.User, roleID int) error {
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

	userExists, err := c.parent.Users().userExists(ctx, u)
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

	_, err = c.parent.Client.Products.PostProjectsProjectIDMembers(&products.PostProjectsProjectIDMembersParams{
		ProjectID:     int64(p.ProjectID),
		ProjectMember: m,
		Context:       ctx,
	}, c.parent.AuthInfo)

	return handleSwaggerProjectErrors(err)
}

// List members lists all members (users and groups alike) of a project.
func (c *ProjectRESTClient) ListMembers(ctx context.Context, p *model.Project) ([]*model.ProjectMemberEntity, error) {
	if p == nil {
		return nil, &ErrProjectNotProvided{}
	}

	entityName := ""

	resp, err := c.parent.Client.Products.GetProjectsProjectIDMembers(&products.GetProjectsProjectIDMembersParams{
		Entityname: &entityName,
		ProjectID:  int64(p.ProjectID),
		Context:    ctx,
	}, c.parent.AuthInfo)

	err = handleSwaggerProjectErrors(err)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

// UpdateUserMemberRole updates the role of a member.
func (c *ProjectRESTClient) UpdateUserMemberRole(ctx context.Context, p *model.Project, u *model.User, roleID int) error {
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

	_, err = c.parent.Client.Products.PutProjectsProjectIDMembersMid(&products.PutProjectsProjectIDMembersMidParams{
		Mid:       mid,
		ProjectID: int64(p.ProjectID),
		Role:      roleRequest,
		Context:   ctx,
	}, c.parent.AuthInfo)

	return handleSwaggerProjectErrors(err)
}

// DeleteUserMember deletes the membership of a user on a project.
func (c *ProjectRESTClient) DeleteUserMember(ctx context.Context, p *model.Project, u *model.User) error {
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

	_, err = c.parent.Client.Products.DeleteProjectsProjectIDMembersMid(
		&products.DeleteProjectsProjectIDMembersMidParams{
			Mid:       mid,
			ProjectID: int64(p.ProjectID),
			Context:   ctx,
		}, c.parent.AuthInfo)

	return handleSwaggerProjectErrors(err)
}

// GetMetadataFromKV returns a ProjectMetadata object matching
// the provided key and containing the provided value.
func GetMetadataFromKV(key ProjectMetadataKey, value string) *model.ProjectMetadata {
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
func (c *ProjectRESTClient) AddMetadata(ctx context.Context, p *model.Project, key ProjectMetadataKey, value string) error {
	if p == nil {
		return &ErrProjectNotProvided{}
	}

	_, err := c.parent.Client.Products.PostProjectsProjectIDMetadatas(&products.PostProjectsProjectIDMetadatasParams{
		Metadata:  GetMetadataFromKV(key, value),
		ProjectID: int64(p.ProjectID),
		Context:   ctx,
	}, c.parent.AuthInfo)

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

// GetMetadataValue retrieves metadata with key of project p.
func (c *ProjectRESTClient) GetMetadataValue(ctx context.Context, p *model.Project, key ProjectMetadataKey) (string, error) {
	if p == nil {
		return "", &ErrProjectNotProvided{}
	}

	resp, err := c.parent.Client.Products.GetProjectsProjectIDMetadatasMetaName(&products.GetProjectsProjectIDMetadatasMetaNameParams{
		MetaName:  string(key),
		ProjectID: int64(p.ProjectID),
		Context:   ctx,
	}, c.parent.AuthInfo)

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
func (c *ProjectRESTClient) ListMetadata(ctx context.Context, p *model.Project) (*model.ProjectMetadata, error) {
	if p == nil {
		return nil, &ErrProjectNotProvided{}
	}

	resp, err := c.parent.Client.Products.GetProjectsProjectIDMetadatas(&products.GetProjectsProjectIDMetadatasParams{
		ProjectID: int64(p.ProjectID),
		Context:   ctx,
	}, c.parent.AuthInfo)

	err = handleSwaggerProjectErrors(err)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

// UpdateMetadata deletes the specified metadata key, if it exists and re-adds this metadata key with the given value.
// This function works around the faulty behaviour of the corresponding 'Update' endpoint of the Harbor API.
func (c *ProjectRESTClient) UpdateMetadata(ctx context.Context, p *model.Project, key ProjectMetadataKey, value string) error {
	if p == nil {
		return &ErrProjectNotProvided{}
	}

	pID := int64(p.ProjectID)
	metaKeyName := string(key)

	_, err := c.parent.Client.Products.GetProjectsProjectIDMetadatasMetaName(&products.GetProjectsProjectIDMetadatasMetaNameParams{
		MetaName:  metaKeyName,
		ProjectID: pID,
		Context:   ctx,
	}, c.parent.AuthInfo)

	if err != nil {
		return handleSwaggerProjectErrors(err)
	}

	_, err = c.parent.Client.Products.DeleteProjectsProjectIDMetadatasMetaName(&products.DeleteProjectsProjectIDMetadatasMetaNameParams{
		MetaName:  metaKeyName,
		ProjectID: pID,
		Context:   ctx,
	}, c.parent.AuthInfo)

	if err != nil {
		return handleSwaggerProjectErrors(err)
	}

	_, err = c.parent.Client.Products.PostProjectsProjectIDMetadatas(&products.PostProjectsProjectIDMetadatasParams{
		Metadata:  GetMetadataFromKV(key, value),
		ProjectID: pID,
		Context:   ctx,
	}, c.parent.AuthInfo)

	return handleSwaggerProjectErrors(err)
}

// DeleteMetadataValue deletes metadata of project p given by key.
func (c *ProjectRESTClient) DeleteMetadataValue(ctx context.Context, p *model.Project, key ProjectMetadataKey) error {
	if p == nil {
		return &ErrProjectNotProvided{}
	}

	_, err := c.parent.Client.Products.DeleteProjectsProjectIDMetadatasMetaName(&products.DeleteProjectsProjectIDMetadatasMetaNameParams{
		MetaName:  string(key),
		ProjectID: int64(p.ProjectID),
		Context:   ctx,
	}, c.parent.AuthInfo)

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
func (c *ProjectRESTClient) projectExists(ctx context.Context, p *model.Project) (bool, error) {
	_, err := c.Get(ctx, p.Name)

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
func (c *ProjectRESTClient) getMid(ctx context.Context, p *model.Project, u *model.User) (int64, error) {
	members, err := c.ListMembers(ctx, p)
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
