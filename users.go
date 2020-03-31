package harbor

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

// User holds the details of a user.
type User struct {
	UserID       int    `json:"user_id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	RealName     string `json:"realname"`
	Comment      string `json:"comment"`
	Deleted      bool   `json:"deleted"`
	Rolename     string `json:"role_name"`
	Role         int    `json:"role_id"`
	RoleList     []Role `json:"role_list"`
	HasAdminRole bool   `json:"has_admin_role"`
	ResetUUID    string `json:"reset_uuid"`
	Salt         string `json:"-"`
}

type UserRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	RealName     string `json:"realname"`
	Comment      string `json:"comment,omitempty"`
	Role         int    `json:"role_id"`
	HasAdminRole bool   `json:"has_admin_role"`
	UserID       int    `json:"user_id,omitempty"`
}

type ChangePassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type UserSearchResult []struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
}

// Add a user
func (s *ProjectsService) AddUser(usr UserRequest) (*gorequest.Response, []error) {
	resp, _, errs := s.client.
		NewRequest(gorequest.POST, fmt.Sprintf("users")).
		Send(usr).
		End()
	return &resp, errs
}

// Get a user's profile by ID
func (s *ProjectsService) GetUser(usr UserRequest) (*gorequest.Response, []error) {
	resp, _, errs := s.client.
		NewRequest(gorequest.GET, fmt.Sprintf("users/%d", usr.UserID)).
		End()
	return &resp, errs
}

// Search User searches for a user by name
func (s *ProjectsService) SearchUser(usr UserRequest) (*gorequest.Response, []error) {
	resp, _, errs := s.client.
		NewRequest(gorequest.GET, fmt.Sprintf("users/search?username=%s", usr.Username)).
		EndStruct(&UserSearchResult{})
	return &resp, errs
}

// Delete a user
func (s *ProjectsService) DeleteUser(usr UserRequest) (*gorequest.Response, []error) {
	resp, _, errs := s.client.
		NewRequest(gorequest.DELETE, fmt.Sprintf("users/%d", usr.UserID)).
		End()
	return &resp, errs
}

// Toggle administrator
func (s *ProjectsService) ToggleUserSysAdmin(usr UserRequest) (*gorequest.Response, []error) {
	resp, _, errs := s.client.
		NewRequest(gorequest.PUT, fmt.Sprintf("users/%d/sysadmin", usr.UserID)).
		Send(usr.HasAdminRole).
		End()
	return &resp, errs
}

// Update a user's password
func (s *ProjectsService) UpdateUserPassword(oldUsr, newUsr UserRequest) (*gorequest.Response, []error) {
	cp := ChangePassword{
		OldPassword: oldUsr.Password,
		NewPassword: newUsr.Password,
	}
	resp, _, errs := s.client.
		NewRequest(gorequest.PUT, fmt.Sprintf("users/%d/password", oldUsr.UserID)).
		Send(cp).
		End()
	return &resp, errs
}

// Update a user's profile
func (s *ProjectsService) UpdateUserProfile(usr UserRequest) (*gorequest.Response, []error) {
	resp, _, errs := s.client.
		NewRequest(gorequest.PUT, fmt.Sprintf("users/%d", usr.UserID)).
		Send(usr).
		End()
	return &resp, errs
}
