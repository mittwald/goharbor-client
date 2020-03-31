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
	Deleted      int    `json:"deleted"`
	Rolename     string `json:"role_name"`
	Role         int    `json:"role_id"`
	RoleList     []Role `json:"role_list"`
	HasAdminRole int    `json:"has_admin_role"`
	ResetUUID    string `json:"reset_uuid"`
	Salt         string `json:"-"`
}

type UserRequest struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	RealName     string `json:"realname"`
	Role         int    `json:"role_id"`
	HasAdminRole bool   `json:"has_admin_role"`
}

// Add a user
func (s *ProjectsService) AddUser(usr UserRequest) (*gorequest.Response, []error) {
	resp, _, errs := s.client.
		NewRequest(gorequest.POST, fmt.Sprintf("users")).
		Send(usr).
		End()
	return &resp, errs
}
