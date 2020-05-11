package harbor

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"net/url"
)

// ProjectClient abstracts away the communication implementation of
// project related methods of Harbor.
type ProjectClient interface {
	// ListProjects returns all projects created by Harbor,
	// and can be filtered by project name.
	ListProjects(opt ListProjectsOptions) ([]Project, error)

	// CheckProject checks if the project name provided already exist.
	CheckProject(projectName string) error

	// GetProjectByID returns specific project details.
	GetProjectByID(pid int64) (Project, error)

	// CreateProject creates a new project.
	CreateProject(p ProjectRequest) error

	// UpdateProject updates the properties of a project.
	UpdateProject(pid int64, p Project) error

	// DeleteProject deletes a project by project ID.
	DeleteProject(pid int64) error

	// GetProjectLogByID retrieves access logs of a project
	// with user-specified filter operations and date time ranges.
	GetProjectLogByID(pid int64, opt ListLogOptions) ([]AccessLog, error)

	// GetProjectMetadata retrieves the metadata of a project.
	GetProjectMetadata(pid int64) (map[string]string, error)

	// AddProjectMetadata adds metadata to a project.
	AddProjectMetadata(pid int64, metadata map[string]string) error

	// GetProjectMetadataSingle retrieves the specified metadata value of a project.
	GetProjectMetadataSingle(pid int64, specified string) (map[string]string, error)

	// UpdateProjectMetadataSingle updates the metadata of a project.
	UpdateProjectMetadataSingle(pid int64, metadataName string) error

	// DeleteProjectMetadataSingle deletes a specified metadata value of a project.
	DeleteProjectMetadataSingle(pid int64, metadataName string) error

	// GetProjectMembers retrieves members of the specified project.
	GetProjectMembers(pid int64) ([]Member, error)

	// AddProjectMember adds a project member to a project.
	AddProjectMember(pid int64, member MemberReq) error

	// GetProjectMember retrieves the role of a project member.
	GetProjectMember(pid, mid int64) (Role, error)

	// UpdateProjectMember updates a project member.
	UpdateProjectMember(pid, mid int64, role RoleRequest) error

	// DeleteProjectMember deletes a project member.
	DeleteProjectMember(pid, mid int64) error
}

// RestProjectClient implements the ProjectClient interface by communicating via Rest api.
type RestProjectClient struct {
	*RestClient
}

// ListProjects satisfies the ProjectClient interface.
func (s *RestProjectClient) ListProjects(opt ListProjectsOptions) ([]Project, error) {
	var projects []Project
	resp, _, errs := s.NewRequest(gorequest.GET, "").
		Query(opt).
		EndStruct(&projects)

	return projects, CheckResponse(errs, resp, 200)
}

// CheckProject satisfies the ProjectClient interface.
func (s *RestProjectClient) CheckProject(projectName string) error {
	resp, _, errs := s.NewRequest(gorequest.HEAD, "").
		Query(map[string]string{"project_name": projectName}).
		End()

	return CheckResponse(errs, resp, 200)
}

// GetProjectByID satisfies the ProjectClient interface.
func (s *RestProjectClient) GetProjectByID(pid int64) (Project, error) {
	var project Project
	resp, _, errs := s.NewRequest(gorequest.GET,"/"+I64toA(pid)).
		EndStruct(&project)
	return project, CheckResponse(errs, resp, 200)
}

// CreateProject satisfies the ProjectClient interface.
func (s *RestProjectClient) CreateProject(p ProjectRequest) error {
	resp, _, errs := s.NewRequest(gorequest.POST, "").
		Send(p).
		End()

	return CheckResponse(errs, resp, 201)
}

// UpdateProject satisfies the ProjectClient interface.
func (s *RestProjectClient) UpdateProject(pid int64, p Project) error {
	resp, _, errs := s.NewRequest(gorequest.PUT,"/"+I64toA(pid)).
		Send(p).
		End()
	return CheckResponse(errs, resp, 200)
}

// DeleteProject satisfies the ProjectClient interface.
func (s *RestProjectClient) DeleteProject(pid int64) error {
	resp, _, errs := s.NewRequest(gorequest.DELETE, "/"+I64toA(pid)).
		End()
	return CheckResponse(errs, resp, 200)
}

// GetProjectLogByID satisfies the ProjectClient interface.
func (s *RestProjectClient) GetProjectLogByID(pid int64, opt ListLogOptions) ([]AccessLog, error) {
	var accessLog []AccessLog
	resp, _, errs := s.NewRequest(gorequest.GET,
		fmt.Sprintf("/%d/logs", pid)).
		Query(opt).
		EndStruct(&accessLog)
	return accessLog, CheckResponse(errs, resp, 200)
}

// GetProjectMetadata satisfies the ProjectClient interface.
func (s *RestProjectClient) GetProjectMetadata(pid int64) (map[string]string, error) {
	var metadata map[string]string
	resp, _, errs := s.NewRequest(gorequest.GET,
		fmt.Sprintf("/%d/metedatas", pid)).
		EndStruct(&metadata)
	return metadata, CheckResponse(errs, resp, 200)
}

// AddProjectMetadata satisfies the ProjectClient interface.
func (s *RestProjectClient) AddProjectMetadata(pid int64, metadata map[string]string) error {
	resp, _, errs := s.NewRequest(gorequest.POST,
		fmt.Sprintf("/%d/metadatas", pid)).
		Send(metadata).
		End()
	return CheckResponse(errs, resp, 200)
}

// GetProjectMetadataSingle satisfies the ProjectClient interface.
func (s *RestProjectClient) GetProjectMetadataSingle(pid int64, specified string) (map[string]string, error) {
	var metadata map[string]string
	resp, _, errs := s.NewRequest(gorequest.GET,
		fmt.Sprintf("/%d/metadatas/%s", pid, url.PathEscape(specified))).
		EndStruct(&metadata)
	return metadata, CheckResponse(errs, resp, 200)
}

// UpdateProjectMetadataSingle satisfies the ProjectClient interface.
func (s *RestProjectClient) UpdateProjectMetadataSingle(pid int64, metadataName string) error {
	resp, _, errs := s.NewRequest(gorequest.PUT,
		fmt.Sprintf("/%d/metadatas/%s", pid, url.PathEscape(metadataName))).
		End()
	return CheckResponse(errs, resp, 200)
}

// DeleteProjectMetadataSingle satisfies the ProjectClient interface.
func (s *RestProjectClient) DeleteProjectMetadataSingle(pid int64, metadataName string) error {
	resp, _, errs := s.NewRequest(gorequest.DELETE,
		fmt.Sprintf("/%d/metadatas/%s", pid, url.PathEscape(metadataName))).
		End()
	return CheckResponse(errs, resp, 200)
}

// GetProjectMembers satisfies the ProjectClient interface.
func (s *RestProjectClient) GetProjectMembers(pid int64) ([]Member, error) {
	var mem []Member
	resp, _, errs := s.NewRequest(gorequest.GET,
		fmt.Sprintf("/%d/members", pid)).
		EndStruct(&mem)
	return mem, CheckResponse(errs, resp, 200)
}

// AddProjectMember satisfies the ProjectClient interface.
func (s *RestProjectClient) AddProjectMember(pid int64, member MemberReq) error {
	resp, _, errs := s.NewRequest(gorequest.POST,
		fmt.Sprintf("/%d/members", pid)).
		Send(member).
		End()
	return CheckResponse(errs, resp, 201)
}

// GetProjectMember satisfies the ProjectClient interface.
func (s *RestProjectClient) GetProjectMember(pid, mid int64) (Role, error) {
	var role Role
	resp, _, errs := s.NewRequest(gorequest.GET,
		fmt.Sprintf("/%d/members/%d", pid, mid)).
		EndStruct(&role)
	return role, CheckResponse(errs, resp, 200)
}

// UpdateProjectMember satisfies the ProjectClient interface.
func (s *RestProjectClient) UpdateProjectMember(pid, mid int64, role RoleRequest) error {
	resp, _, errs := s.NewRequest(gorequest.PUT,
		fmt.Sprintf("/%d/members/%d", pid, mid)).
		Send(role).
		End()
	return CheckResponse(errs, resp, 200)
}

// DeleteProjectMember satisfies the ProjectClient interface.
func (s *RestProjectClient) DeleteProjectMember(pid, mid int64) error {
	resp, _, errs := s.NewRequest(gorequest.DELETE,
		fmt.Sprintf("/%d/members/%d", pid, mid)).
		End()
	return CheckResponse(errs, resp, 200)
}
