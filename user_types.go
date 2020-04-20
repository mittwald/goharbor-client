package harbor

// User holds the details of a user
type User struct {
	UserID       int64  `json:"user_id"`
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
	UserID       int64  `json:"user_id,omitempty"`
}

// ChangePassword holds the information needed to change a users password
type ChangePassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

// ChangePasswordAsAdmin holds the information needed to change a users password as administrator
type ChangePasswordAsAdmin struct {
	NewPassword string `json:"new_password"`
}

type UserSearchResults []UserSearchResult

// UserSearchResult holds the information returned by the API when querying for a user name
type UserSearchResult struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
}
