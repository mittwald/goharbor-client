package harbor

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

// UserClient abstracts away the communication implementation of
// user related methods of Harbor.
type UserClient interface {
	// SearchUser searches for a user by name.
	SearchUser(usr UserMember) (UserSearchResults, error)

	// ListUsers retrieves  a list of users.
	ListUsers() ([]User, error)

	// GetUser retrieves a user's profile by ID.
	GetUser(usr UserRequest) (User, error)

	// AddUser retrieves a user.
	AddUser(usr UserRequest) error

	// UpdateUserProfile updates a user's profile.
	UpdateUserProfile(usr UserRequest) error

	// DeleteUser deletes a user.
	DeleteUser(usr UserRequest) error

	// ToggleUserSysAdmin toggles administrator privileges of a user.
	ToggleUserSysAdmin(usr UserRequest) error

	// UpdateUserPassword updates a users password.
	// NOTE: when using an admin user, the usage of ChangePasswordAsAdmin is recommended.
	UpdateUserPassword(uid int64, oldPw, newPw string) error

	// UpdateUserPasswordAsAdmin updates a users password as admin (only usable by an admin user).
	UpdateUserPasswordAsAdmin(uid int64, newPw string) error
}

// RestUserClient implements the UserClient interface by communicating via Rest api.
type RestUserClient struct {
	*RestClient
}

// SearchUser satisfies the UserClient interface.
func (s *RestUserClient) SearchUser(usr UserMember) (UserSearchResults, error) {
	var u UserSearchResults
	resp, _, errs := s.NewRequest(gorequest.GET, "/search").
		Query(map[string]string{"username": usr.Username}).
		EndStruct(&u)
	return u, CheckResponse(errs, resp, 200)
}

// ListUsers satisfies the UserClient interface.
func (s *RestUserClient) ListUsers() ([]User, error) {
	var u []User
	resp, _, errs := s.NewRequest(gorequest.GET, "").
		EndStruct(&u)
	return u, CheckResponse(errs, resp, 200)
}

// GetUser satisfies the UserClient interface.
func (s *RestUserClient) GetUser(usr UserRequest) (User, error) {
	var u User
	resp, _, errs := s.NewRequest(gorequest.GET,"/"+I64toA(usr.UserID)).
		EndStruct(&u)
	return u, CheckResponse(errs, resp, 200)
}

// AddUser satisfies the UserClient interface.
func (s *RestUserClient) AddUser(usr UserRequest) error {
	resp, _, errs := s.NewRequest(gorequest.POST, "").
		Send(usr).
		End()
	return CheckResponse(errs, resp, 201)
}

// UpdateUserProfile satisfies the UserClient interface.
func (s *RestUserClient) UpdateUserProfile(usr UserRequest) error {
	resp, _, errs := s.NewRequest(gorequest.PUT,"/"+I64toA(usr.UserID)).
		Send(usr).
		End()
	return CheckResponse(errs, resp, 200)
}

// DeleteUser satisfies the UserClient interface.
func (s *RestUserClient) DeleteUser(usr UserRequest) error {
	resp, _, errs := s.NewRequest(gorequest.DELETE,"/"+I64toA(usr.UserID)).
		End()
	return CheckResponse(errs, resp, 200)
}

// ToggleUserSysAdmin satisfies the UserClient interface.
func (s *RestUserClient) ToggleUserSysAdmin(usr UserRequest) error {
	resp, _, errs := s.NewRequest(gorequest.PUT,fmt.Sprintf("/%d/sysadmin", usr.UserID)).
		Send(usr.HasAdminRole).
		End()
	return CheckResponse(errs, resp, 200)
}

// UpdateUserPassword satisfies the UserClient interface.
func (s *RestUserClient) UpdateUserPassword(uid int64, oldPw, newPw string) error {
	cp := ChangePassword{
		OldPassword: oldPw,
		NewPassword: newPw,
	}

	resp, _, errs := s.NewRequest(gorequest.PUT,
		fmt.Sprintf("/%d/password", uid)).
		Send(cp).
		End()
	return CheckResponse(errs, resp, 200)
}

// UpdateUserPasswordAsAdmin satisfies the UserClient interface.
func (s *RestUserClient) UpdateUserPasswordAsAdmin(uid int64, newPw string) error {
	cp := ChangePasswordAsAdmin{
		NewPassword: newPw,
	}

	resp, _, errs := s.NewRequest(gorequest.PUT,
		fmt.Sprintf("/%d/password", uid)).
		Send(cp).
		End()
	return CheckResponse(errs, resp, 200)
}
