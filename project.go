package harbor

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"time"
)

// ProjectClient handles communication with the project related methods of the Harbor API.
type ProjectClient struct {
	client *Client
}
// ProjectMetadata holds the metadata of a project.
type ProjectMetadata struct {
	ID        int64  `json:"id"`
	ProjectID int64  `json:"project_id"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	Deleted   bool   `json:"deleted"`
}

// Project holds the details of a project.
type Project struct {
	ProjectID    int64             `json:"project_id"`
	OwnerID      int64             `json:"owner_id"`
	Name         string            `json:"name"`
	CreationTime time.Time         `json:"creation_time"`
	UpdateTime   time.Time         `json:"update_time"`
	Deleted      bool              `json:"deleted"`
	OwnerName    string            `json:"owner_name"`
	Toggleable   bool              `json:"toggleable"`
	Role         int               `json:"current_user_role_id"`
	RepoCount    int64             `json:"repo_count"`
	Metadata     map[string]string `json:"metadata"`
	CVEWhitelist CVEWhitelist      `json:"CVEWhitelist"`
	StorageLimit int64             `json:"storageLimit"`
}

// Role holds the details of a role.
type Role struct {
	RoleID   int    `json:"role_id"`
	RoleCode string `json:"role_code"`
	Name     string `json:"role_name"`
	RoleMask int    `json:"role_mask"`
}

type RoleRequest struct {
	Role int64 `json:"role"`
}

// CVEWhitelistItem holds the CVE ids of a whitelisted item
type CVEWhitelistItem struct {
	CVEID string `json:"CVEID"`
}

// CVEWhitelist holds project specific information next to the set CVEWhitelistItem's
type CVEWhitelist struct {
	ID        int64            `json:"id"`
	ProjectID int64            `json:"projectID"`
	Items     CVEWhitelistItem `json:"items,optional"`
}

// AccessLog holds the information of log entries
type AccessLog struct {
	LogID     int       `json:"log_id"`
	Username  string    `json:"username"`
	ProjectID int64     `json:"project_id"`
	RepoName  string    `json:"repo_name"`
	RepoTag   string    `json:"repo_tag"`
	GUID      string    `json:"guid"`
	Operation string    `json:"operation"`
	OpTime    time.Time `json:"op_time"`
}

// ProjectRequest holds the information needed to create a project
type ProjectRequest struct {
	Name     string            `url:"name,omitempty" json:"project_name"`
	Public   *int              `url:"public,omitempty" json:"public"` //deprecated, reserved for project creation in replication
	Metadata map[string]string `url:"-" json:"metadata"`
}

// ListProjectsOptions holds the information needed to list a project
type ListProjectsOptions struct {
	ListOptions
	Name   string `url:"name,omitempty" json:"name,omitempty"`
	Public bool   `url:"public,omitempty" json:"public,omitempty"`
	Owner  string `url:"owner,omitempty" json:"owner,omitempty"`
}

// LogQueryParam is used to set query conditions when listing access logs
type ListLogOptions struct {
	ListOptions
	Username   string     `url:"username,omitempty"`        // the operator's username of the log
	Repository string     `url:"repository,omitempty"`      // repository name
	Tag        string     `url:"tag,omitempty"`             // tag name
	Operations []string   `url:"operation,omitempty"`       // operations
	BeginTime  *time.Time `url:"begin_timestamp,omitempty"` // the time after which the operation is done
	EndTime    *time.Time `url:"end_timestamp,omitempty"`   // the time before which the operation is doen
}

// MemberRequest holds the information needed to update a project member
type MemberRequest struct {
	UserName string `json:"username"`
	Roles    []int  `json:"roles"`
}

// ProjectMemberRequest holds the information needed to add a project member
type ProjectMemberRequest struct {
	RoleID     int        `json:"role_id"`
	MemberUser MemberUser `json:"member_user"`
}

// MemberUser holds the user information needed for a project member request
type MemberUser struct {
	Username string `json:"username"`
	UserID   int    `json:"user_id"`
}

// CreateProject
// Creates a new project
func (s *ProjectClient) CreateProject(p ProjectRequest) error {
	resp, _, err := s.client.
		NewRequest(gorequest.POST, "projects").
		Send(p).
		End()

	if err != nil {
		return err[len(err)-1]
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("API returned %d when creating project", resp.StatusCode)
	}
	return nil
}

// List projects
//
// This endpoint returns all projects created by Harbor,
// and can be filtered by project name.
func (s *ProjectClient) ListProject(opt *ListProjectsOptions) ([]Project, gorequest.Response, []error) {
	var projects []Project
	resp, _, errs := s.client.
		NewRequest(gorequest.GET, "projects").
		Query(*opt).
		EndStruct(&projects)
	return projects, resp, errs
}

// CheckProject
// Check if the project name provided already exist
func (s *ProjectClient) CheckProject(projectName string) error {
	resp, _, err := s.client.
		NewRequest(gorequest.HEAD, "projects").
		Query(fmt.Sprintf("project_name=%s", projectName)).
		End()
	if err != nil {
		return err[len(err)-1]
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("API returned %d when checking project", resp.StatusCode)
	}

	return  nil
}

// GetProjectByID
// Return specific project details
func (s *ProjectClient) GetProjectByID(pid int64) (Project, gorequest.Response, []error) {
	var project Project
	resp, _, errs := s.client.
		NewRequest(gorequest.GET, fmt.Sprintf("projects/%d", pid)).
		EndStruct(&project)
	return project, resp, errs
}

// UpdateProject
// Update the properties of a project.
func (s *ProjectClient) UpdateProject(pid int64, p Project) (gorequest.Response, []error) {
	resp, _, errs := s.client.
		NewRequest(gorequest.PUT, fmt.Sprintf("projects/%d", pid)).
		Send(p).
		End()
	return resp, errs
}

// DeleteProject
// Delete a project by project ID.
func (s *ProjectClient) DeleteProject(pid int64) error {
	resp, _, err := s.client.
		NewRequest(gorequest.DELETE, fmt.Sprintf("projects/%d", pid)).
		End()
	if err != nil {
		return  err[len(err)-1]
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("API returned %d when deleting project", resp.StatusCode)
	}
	return  nil
}

// GetProjectLogByID
// Get access logs of a project with user-specified filter operations and date time ranges
func (s *ProjectClient) GetProjectLogByID(pid int64, opt ListLogOptions) ([]AccessLog, gorequest.Response, []error) {
	var accessLog []AccessLog
	resp, _, errs := s.client.
		NewRequest(gorequest.GET, fmt.Sprintf("projects/%d", pid)).
		Query(opt).
		EndStruct(&accessLog)
	return accessLog, resp, errs
}

// GetProjectMetadataById
// Get the metadata of a project
func (s *ProjectClient) GetProjectMetadataById(pid int64) (map[string]string, gorequest.Response, []error) {
	var metadata map[string]string
	resp, _, errs := s.client.
		NewRequest(gorequest.GET, fmt.Sprintf("projects/%d", pid)).
		EndStruct(&metadata)
	return metadata, resp, errs
}

// AddProjectMetadata
// Add metadata to a project
func (s *ProjectClient) AddProjectMetadata(pid int64, metadata map[string]string) (gorequest.Response, []error) {
	resp, _, errs := s.client.
		NewRequest(gorequest.POST, fmt.Sprintf("projects/%d/metadatas", pid)).
		Send(metadata).
		End()
	return resp, errs
}

// GetProjectMetadata
// Get the specified metadata value of a project
func (s *ProjectClient) GetProjectMetadata(pid int64, specified string) (map[string]string, gorequest.Response, []error) {
	var metadata map[string]string
	resp, _, errs := s.client.
		NewRequest(gorequest.GET, fmt.Sprintf("projects/%d/metadatas/%s", pid, specified)).
		EndStruct(&metadata)
	return metadata, resp, errs
}

// UpdateProjectMetadata
// Update the metadata of a project.
func (s *ProjectClient) UpdateProjectMetadata(pid int64, metadataName string) (gorequest.Response, []error) {
	resp, _, errs := s.client.
		NewRequest(gorequest.PUT, fmt.Sprintf("projects/%d/%s", pid, metadataName)).
		End()
	return resp, errs
}

// DeleteProjectMetadata
// Delete a specified metadata value of a project.
func (s *ProjectClient) DeleteProjectMetadata(pid int64, metadataName string) (gorequest.Response, []error) {
	resp, _, errs := s.client.
		NewRequest(gorequest.DELETE, fmt.Sprintf("projects/%d/%s", pid, metadataName)).
		End()
	return resp, errs
}

// GetProjectMembers
// Get the specified project’s members
func (s *ProjectClient) GetProjectMembers(pid int64) ([]User, gorequest.Response, []error) {
	var users []User
	resp, _, errs := s.client.
		NewRequest(gorequest.GET, fmt.Sprintf("projects/%d/members", pid)).
		EndStruct(&users)
	return users, resp, errs
}

// UpdateProjectMember
// Update a project member
func (s *ProjectClient) UpdateProjectMember(pid, mid int64, role RoleRequest) (gorequest.Response, []error) {
	resp, _, errs := s.client.
		NewRequest(gorequest.PUT, fmt.Sprintf("projects/%d/members/%d", pid, mid)).
		Send(role).
		End()
	return resp, errs
}

// AddProjectMember
//
// This endpoint is for user to add project role member accompany with relevant project and user.
//
func (s *ProjectClient) AddProjectMember(pid int, member ProjectMemberRequest) (gorequest.Response, []error) {
	resp, _, errs := s.client.
		NewRequest(gorequest.POST, fmt.Sprintf("projects/%d/members", pid)).
		Send(member).
		End()
	return resp, errs
}

// GetProjectMemberRole
// Get the role of a project member
//
// Harbor API docs: https://github.com/vmware/harbor/blob/release-1.4.0/docs/swagger.yaml#L522
func (s *ProjectClient) GetProjectMemberRole(pid, mid int) (Role, gorequest.Response, []error) {
	var role Role
	resp, _, errs := s.client.
		NewRequest(gorequest.GET, fmt.Sprintf("projects/%d/members/%d", pid, mid)).
		EndStruct(&role)
	return role, resp, errs
}



// DeleteProjectMember
// Delete a project member
func (s *ProjectClient) DeleteProjectMember(pid, mid int64) (gorequest.Response, []error) {
	resp, _, errs := s.client.
		NewRequest(gorequest.DELETE, fmt.Sprintf("projects/%d/members/%d", pid, mid)).
		End()
	return resp, errs
}
