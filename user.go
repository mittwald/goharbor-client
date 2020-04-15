package harbor

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

// UserClient handles communication with the user related methods of the Harbor API
type UserClient struct {
	client *Client
}

// User holds the details of a user
type User struct {
	UserID       int    `json:"user_id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	RealName     string `json:"realname"`
	Comment      string `json:"comment"`
	Deleted      bool   `json:"deleted"`
	RoleName     string `json:"role_name"`
	Role         int    `json:"role_id"`
	RoleList     []Role `json:"role_list"`
	HasAdminRole bool   `json:"has_admin_role"`
	ResetUUID    string `json:"reset_uuid"`
	Salt         string `json:"-"`
}

// UserRequest holds the information needed for basic operations on users
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

// ChangePassword holds the information needed to change a users password
// NOTE: when using an admin user, the usage of ChangePasswordAsAdmin is recommended
type ChangePassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

// ChangePasswordAsAdmin holds the information needed to change a users password as administrator
type ChangePasswordAsAdmin struct {
	NewPassword string `json:"new_password"`
}

// UserSearchResult holds the information returned by the API when querying for a user name

type UserSearchResults []UserSearchResult

type UserSearchResult struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
}

// AddUser
// Add a user
func (s *UserClient) AddUser(usr UserRequest) (gorequest.Response, []error) {
	resp, _, errs := s.client.
		NewRequest(gorequest.POST, fmt.Sprintf("users")).
		Send(usr).
		End()
	return resp, errs
}

// GetUser
// Get a user's profile by ID
func (s *UserClient) GetUser(usr UserRequest) (User, gorequest.Response, []error) {
	var u User
	resp, _, errs := s.client.
		NewRequest(gorequest.GET, fmt.Sprintf("users/%d", usr.UserID)).
		EndStruct(&u)
	return u, resp, errs
}

// SearchUser
// Search User searches for a user by name
func (s *UserClient) SearchUser(usr UserRequest) (UserSearchResults, gorequest.Response, []error) {
	var u UserSearchResults
	resp, _, errs := s.client.
		NewRequest(gorequest.GET, fmt.Sprintf("users/search?username=%s", usr.Username)).
		EndStruct(&u)
	return u, resp, errs
}

// DeleteUser
// Delete a user
func (s *UserClient) DeleteUser(usr UserRequest) (gorequest.Response, []error) {
	resp, _, errs := s.client.
		NewRequest(gorequest.DELETE, fmt.Sprintf("users/%d", usr.UserID)).
		End()
	return resp, errs
}

// ToggleUserSysAdmin
// Toggle administrator privileges of a user
func (s *UserClient) ToggleUserSysAdmin(usr UserRequest) (gorequest.Response, []error) {
	resp, _, errs := s.client.
		NewRequest(gorequest.PUT, fmt.Sprintf("users/%d/sysadmin", usr.UserID)).
		Send(usr.HasAdminRole).
		End()
	return resp, errs
}

// UpdateUserPassword
// Update a users password
func (s *UserClient) UpdateUserPassword(oldUsr, newUsr UserRequest) (gorequest.Response, []error) {
	cp := ChangePassword{
		OldPassword: oldUsr.Password,
		NewPassword: newUsr.Password,
	}
	resp, _, errs := s.client.
		NewRequest(gorequest.PUT, fmt.Sprintf("users/%d/password", oldUsr.UserID)).
		Send(cp).
		End()
	return resp, errs
}

// UpdateUserPasswordAsAdmin
// Update a user's password as admin (only usable by an admin user)
func (s *UserClient) UpdateUserPasswordAsAdmin(newUsr UserRequest) (gorequest.Response, []error) {
	cp := ChangePasswordAsAdmin{
		NewPassword: newUsr.Password,
	}
	resp, _, errs := s.client.
		NewRequest(gorequest.PUT, fmt.Sprintf("users/%d/password", newUsr.UserID)).
		Send(cp).
		End()
	return resp, errs
}

// UpdateUserProfile
// Update a user's profile
func (s *UserClient) UpdateUserProfile(usr UserRequest) (gorequest.Response, []error) {
	resp, _, errs := s.client.
		NewRequest(gorequest.PUT, fmt.Sprintf("users/%d", usr.UserID)).
		Send(usr).
		End()
	return resp, errs
}
