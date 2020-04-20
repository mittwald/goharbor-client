package harbor

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

// UserClient handles communication with the user related methods of the Harbor API
type UserClient struct {
	*Client
}

// SearchUser
// Search User searches for a user by name
func (s *UserClient) SearchUser(usr UserRequest) (UserSearchResults, error) {
	var u UserSearchResults
	resp, _, errs := s.NewRequest(gorequest.GET, "/search").
		Query(fmt.Sprintf("username=%s", usr.Username)).
		EndStruct(&u)
	return u, CheckResponse(errs, resp, 200)
}

// ListUsers
// Get a list of users
func (s *UserClient) ListUsers() ([]User, error) {
	var u []User
	resp, _, errs := s.NewRequest(gorequest.GET, "").
		EndStruct(&u)
	return u, CheckResponse(errs, resp, 200)
}

// GetUser
// Get a user's profile by ID
func (s *UserClient) GetUser(usr UserRequest) (User, error) {
	var u User
	resp, _, errs := s.NewRequest(gorequest.GET, "/"+I64toA(usr.UserID)).
		EndStruct(&u)
	return u, CheckResponse(errs, resp, 200)
}

// AddUser
// Add a user
func (s *UserClient) AddUser(usr UserRequest) error {
	resp, _, errs := s.NewRequest(gorequest.POST, "").
		Send(usr).
		End()
	return CheckResponse(errs, resp, 201)
}

// UpdateUserProfile
// Update a user's profile
func (s *UserClient) UpdateUserProfile(usr UserRequest) error {
	resp, _, errs := s.NewRequest(gorequest.PUT, "/"+I64toA(usr.UserID)).
		Send(usr).
		End()
	return CheckResponse(errs, resp, 200)
}

// DeleteUser
// Delete a user
func (s *UserClient) DeleteUser(usr UserRequest) error {
	resp, _, errs := s.NewRequest(gorequest.DELETE, "/"+I64toA(usr.UserID)).
		End()
	return CheckResponse(errs, resp, 200)
}

// ToggleUserSysAdmin
// Toggle administrator privileges of a user
func (s *UserClient) ToggleUserSysAdmin(usr UserRequest) error {
	resp, _, errs := s.NewRequest(gorequest.PUT, fmt.Sprintf("/%d/sysadmin", usr.UserID)).
		Send(usr.HasAdminRole).
		End()
	return CheckResponse(errs, resp, 200)
}

// UpdateUserPassword
// Update a users password
// NOTE: when using an admin user, the usage of ChangePasswordAsAdmin is recommended
func (s *UserClient) UpdateUserPassword(uid int64, oldPw, newPw string) error {
	cp := ChangePassword{
		OldPassword: oldPw,
		NewPassword: newPw,
	}

	resp, _, errs := s.NewRequest(gorequest.PUT, fmt.Sprintf("/%d/password", uid)).
		Send(cp).
		End()
	return CheckResponse(errs, resp, 200)
}

// UpdateUserPasswordAsAdmin
// Update a users password as admin (only usable by an admin user)
func (s *UserClient) UpdateUserPasswordAsAdmin(uid int64, newPw string) error {
	cp := ChangePasswordAsAdmin{
		NewPassword: newPw,
	}

	resp, _, errs := s.NewRequest(gorequest.PUT, fmt.Sprintf("/%d/password", uid)).
		Send(cp).
		End()
	return CheckResponse(errs, resp, 200)
}
