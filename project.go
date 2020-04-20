package harbor

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

// ProjectClient handles communication with the project related methods of the Harbor API.
type ProjectClient struct {
	*Client
}

// List projects
// This endpoint returns all projects created by Harbor,
// and can be filtered by project name
func (s *ProjectClient) ListProjects(opt ListProjectsOptions) ([]Project, error) {
	var projects []Project
	resp, _, errs := s.NewRequest(gorequest.GET, "").
		Query(opt).
		EndStruct(&projects)

	return projects, CheckResponse(errs, resp, 200)
}

// CheckProject
// Check if the project name provided already exist
func (s *ProjectClient) CheckProject(projectName string) error {
	resp, _, errs := s.NewRequest(gorequest.HEAD, "").
		Query(fmt.Sprintf("project_name=%s", projectName)).
		End()

	return CheckResponse(errs, resp, 200)
}

// GetProjectByID
// Return specific project details
func (s *ProjectClient) GetProjectByID(pid int64) (Project, error) {
	var project Project
	resp, _, errs := s.NewRequest(gorequest.GET, "/"+I64toA(pid)).
		EndStruct(&project)
	return project, CheckResponse(errs, resp, 200)
}

// CreateProject
// Creates a new project
func (s *ProjectClient) CreateProject(p ProjectRequest) error {
	resp, _, errs := s.NewRequest(gorequest.POST, "").
		Send(p).
		End()

	return CheckResponse(errs, resp, 200)
}

// UpdateProject
// Update the properties of a project
func (s *ProjectClient) UpdateProject(pid int64, p Project) error {
	resp, _, errs := s.NewRequest(gorequest.PUT, "/"+I64toA(pid)).
		Send(p).
		End()
	return CheckResponse(errs, resp, 200)
}

// DeleteProject
// Delete a project by project ID
func (s *ProjectClient) DeleteProject(pid int64) error {
	resp, _, errs := s.NewRequest(gorequest.DELETE, "/"+I64toA(pid)).
		End()
	return CheckResponse(errs, resp, 200)
}

// GetProjectLogByID
// Get access logs of a project with user-specified filter operations and date time ranges
func (s *ProjectClient) GetProjectLogByID(pid int64, opt ListLogOptions) ([]AccessLog, error) {
	var accessLog []AccessLog
	resp, _, errs := s.NewRequest(gorequest.GET, fmt.Sprintf("/%d/logs", pid)).
		Query(opt).
		EndStruct(&accessLog)
	return accessLog, CheckResponse(errs, resp, 200)
}

// GetProjectMetadataById
// Get the metadata of a project
func (s *ProjectClient) GetProjectMetadata(pid int64) (map[string]string, error) {
	var metadata map[string]string
	resp, _, errs := s.NewRequest(gorequest.GET, fmt.Sprintf("/%d/metedatas", pid)).
		EndStruct(&metadata)
	return metadata, CheckResponse(errs, resp, 200)
}

// AddProjectMetadata
// Add metadata to a project
func (s *ProjectClient) AddProjectMetadata(pid int64, metadata map[string]string) error {
	resp, _, errs := s.NewRequest(gorequest.POST, fmt.Sprintf("/%d/metadatas", pid)).
		Send(metadata).
		End()
	return CheckResponse(errs, resp, 200)
}

// GetProjectMetadata
// Get the specified metadata value of a project
func (s *ProjectClient) GetProjectMetadataSingle(pid int64, specified string) (map[string]string, error) {
	var metadata map[string]string
	resp, _, errs := s.NewRequest(gorequest.GET, fmt.Sprintf("/%d/metadatas/%s", pid, specified)).
		EndStruct(&metadata)
	return metadata, CheckResponse(errs, resp, 200)
}

// UpdateProjectMetadata
// Update the metadata of a project
func (s *ProjectClient) UpdateProjectMetadataSingle(pid int64, metadataName string) error {
	resp, _, errs := s.NewRequest(gorequest.PUT, fmt.Sprintf("/%d/metadatas/%s", pid, metadataName)).
		End()
	return CheckResponse(errs, resp, 200)
}

// DeleteProjectMetadata
// Delete a specified metadata value of a project
func (s *ProjectClient) DeleteProjectMetadataSingle(pid int64, metadataName string) error {
	resp, _, errs := s.NewRequest(gorequest.DELETE, fmt.Sprintf("/%d/metadatas/%s", pid, metadataName)).
		End()
	return CheckResponse(errs, resp, 200)
}

// GetProjectMembers
// Get members of the specified project
func (s *ProjectClient) GetProjectMembers(pid int64) ([]User, error) {
	var users []User
	resp, _, errs := s.NewRequest(gorequest.GET, fmt.Sprintf("/%d/members", pid)).
		EndStruct(&users)
	return users, CheckResponse(errs, resp, 200)
}

// AddProjectMember
// Add a project member to a project
func (s *ProjectClient) AddProjectMember(pid int, member ProjectMemberRequest) error {
	resp, _, errs := s.NewRequest(gorequest.POST, fmt.Sprintf("/%d/members", pid)).
		Send(member).
		End()
	return CheckResponse(errs, resp, 200)
}

// GetProjectMemberRole
// Get the role of a project member
func (s *ProjectClient) GetProjectMember(pid, mid int) (Role, error) {
	var role Role
	resp, _, errs := s.NewRequest(gorequest.GET, fmt.Sprintf("/%d/members/%d", pid, mid)).
		EndStruct(&role)
	return role, CheckResponse(errs, resp, 200)
}

// UpdateProjectMember
// Update a project member
func (s *ProjectClient) UpdateProjectMember(pid, mid int64, role RoleRequest) error {
	resp, _, errs := s.NewRequest(gorequest.PUT, fmt.Sprintf("/%d/members/%d", pid, mid)).
		Send(role).
		End()
	return CheckResponse(errs, resp, 200)
}

// DeleteProjectMember
// Delete a project member
func (s *ProjectClient) DeleteProjectMember(pid, mid int64) error {
	resp, _, errs := s.NewRequest(gorequest.DELETE, fmt.Sprintf("/%d/members/%d", pid, mid)).
		End()
	return CheckResponse(errs, resp, 200)
}
