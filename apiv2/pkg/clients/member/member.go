package member

import (
	"context"

	"github.com/go-openapi/runtime"

	v2client "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client"
	"github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/member"
	"github.com/mittwald/goharbor-client/v5/apiv2/model"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/config"
	"github.com/mittwald/goharbor-client/v5/apiv2/pkg/errors"
)

type EntityType string

const (
	EntityTypeUser  EntityType = "u"
	EntityTypeGroup EntityType = "g"
)

func (t EntityType) String() string {
	return string(t)
}

// RESTClient is a subclient for handling system related actions.
type RESTClient struct {
	// Options contains optional configuration when making API calls.
	Options *config.Options

	// The new client of the harbor v2 API
	V2Client *v2client.Harbor

	// AuthInfo contains the auth information that is provided on API calls.
	AuthInfo runtime.ClientAuthInfoWriter
}

func NewClient(v2Client *v2client.Harbor, opts *config.Options, authInfo runtime.ClientAuthInfoWriter) *RESTClient {
	return &RESTClient{
		Options:  opts,
		V2Client: v2Client,
		AuthInfo: authInfo,
	}
}

type Client interface {
	AddProjectMember(ctx context.Context, projectNameOrID string, m *model.ProjectMember) error
	ListProjectMembers(ctx context.Context, projectNameOrID, memberQuery string) ([]*model.ProjectMemberEntity, error)
	UpdateProjectMember(ctx context.Context, projectNameOrID string, m *model.ProjectMember) error
	DeleteProjectMember(ctx context.Context, projectNameOrID string, m *model.ProjectMember) error
}

// AddProjectMember adds the project member 'm' to the corresponding project.
func (c *RESTClient) AddProjectMember(ctx context.Context, projectNameOrID string, m *model.ProjectMember) error {
	if m == nil {
		return &errors.ErrProjectNoMemberProvided{}
	}

	params := &member.CreateProjectMemberParams{
		ProjectMember:   m,
		ProjectNameOrID: projectNameOrID,
		Context:         ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err := c.V2Client.Member.CreateProjectMember(params, c.AuthInfo)

	return handleSwaggerMemberErrors(err)
}

// ListProjectMembers returns a list of project members.
func (c *RESTClient) ListProjectMembers(ctx context.Context, projectNameOrID, memberQuery string) ([]*model.ProjectMemberEntity, error) {
	var members []*model.ProjectMemberEntity
	page := c.Options.Page

	params := &member.ListProjectMembersParams{
		Page:            &page,
		PageSize:        &c.Options.PageSize,
		Entityname:      &memberQuery,
		ProjectNameOrID: projectNameOrID,
		Context:         ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	for {
		resp, err := c.V2Client.Member.ListProjectMembers(params, c.AuthInfo)
		if err != nil {
			return nil, handleSwaggerMemberErrors(err)
		}

		if len(resp.Payload) == 0 {
			break
		}

		totalCount := resp.XTotalCount

		members = append(members, resp.Payload...)

		if int64(len(members)) >= totalCount {
			break
		}

		page++
	}

	return members, nil
}

// UpdateProjectMember updates a project member.
func (c *RESTClient) UpdateProjectMember(ctx context.Context, projectNameOrID string, m *model.ProjectMember) error {
	mid, err := c.getMemberID(ctx, projectNameOrID, m)
	if err != nil {
		return err
	}

	params := &member.UpdateProjectMemberParams{
		Mid:             mid,
		ProjectNameOrID: projectNameOrID,
		Role:            &model.RoleRequest{RoleID: m.RoleID},
		Context:         ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err = c.V2Client.Member.UpdateProjectMember(params, c.AuthInfo)

	return handleSwaggerMemberErrors(err)
}

// DeleteProjectMember deletes the membership between a user and a project.
func (c *RESTClient) DeleteProjectMember(ctx context.Context, projectNameOrID string, m *model.ProjectMember) error {
	if m == nil {
		return &errors.ErrProjectNoMemberProvided{}
	}

	mid, err := c.getMemberID(ctx, projectNameOrID, m)
	if err != nil {
		return err
	}

	params := &member.DeleteProjectMemberParams{
		Mid:             mid,
		ProjectNameOrID: projectNameOrID,
		Context:         ctx,
	}

	params.WithTimeout(c.Options.Timeout)

	_, err = c.V2Client.Member.DeleteProjectMember(params, c.AuthInfo)

	return handleSwaggerMemberErrors(err)
}

// getMemberID returns the member ID of a user or usergroup in project p.
func (c *RESTClient) getMemberID(ctx context.Context, projectNameOrID string, m *model.ProjectMember) (int64, error) {
	members, err := c.ListProjectMembers(ctx, projectNameOrID, "")
	if err != nil {
		return 0, err
	}

	for _, v := range members {
		switch v.EntityType {
		default:
			return 0, &errors.ErrProjectMemberMismatch{}
		case EntityTypeGroup.String():
			if v.EntityName == m.MemberGroup.GroupName {
				return v.ID, nil
			}
		case EntityTypeUser.String():
			if v.EntityName == m.MemberUser.Username {
				return v.ID, nil
			}
		}

		if v.EntityType == EntityTypeUser.String() && v.EntityName == m.MemberUser.Username {
			return v.ID, nil
		}
		if v.EntityType == EntityTypeGroup.String() && v.EntityName == m.MemberGroup.GroupName {
			return v.ID, nil
		}
	}

	return 0, &errors.ErrNoMemberFound{}
}
