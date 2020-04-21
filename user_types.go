package harbor

import "time"

// User holds the details of a user.
type User struct {
	UserID          int    `orm:"pk;auto;column(user_id)" json:"user_id"`
	Username        string `orm:"column(username)" json:"username"`
	Email           string `orm:"column(email)" json:"email"`
	Password        string `orm:"column(password)" json:"password"`
	PasswordVersion string `orm:"column(password_version)" json:"password_version"`
	Realname        string `orm:"column(realname)" json:"realname"`
	Comment         string `orm:"column(comment)" json:"comment"`
	Deleted         bool   `orm:"column(deleted)" json:"deleted"`
	Rolename        string `orm:"-" json:"role_name"`
	// if this field is named as "RoleID", beego orm can not map role_id
	// to it.
	Role         int  `orm:"-" json:"role_id"`
	SysAdminFlag bool `orm:"column(sysadmin_flag)" json:"sysadmin_flag"`
	// AdminRoleInAuth to store the admin privilege granted by external authentication provider
	AdminRoleInAuth bool      `orm:"-" json:"admin_role_in_auth"`
	ResetUUID       string    `orm:"column(reset_uuid)" json:"reset_uuid"`
	Salt            string    `orm:"column(salt)" json:"-"`
	CreationTime    time.Time `orm:"column(creation_time);auto_now_add" json:"creation_time"`
	UpdateTime      time.Time `orm:"column(update_time);auto_now" json:"update_time"`
	GroupIDs        []int     `orm:"-" json:"-"`
	OIDCUserMeta    *OIDCUser `orm:"-" json:"oidc_user_meta,omitempty"`
}

// OIDCUser ...
type OIDCUser struct {
	ID     int64 `orm:"pk;auto;column(id)" json:"id"`
	UserID int   `orm:"column(user_id)" json:"user_id"`
	// encrypted secret
	Secret string `orm:"column(secret)" json:"-"`
	// secret in plain text
	PlainSecret  string    `orm:"-" json:"secret"`
	SubIss       string    `orm:"column(subiss)" json:"subiss"`
	Token        string    `orm:"column(token)" json:"-"`
	CreationTime time.Time `orm:"column(creation_time);auto_now_add" json:"creation_time"`
	UpdateTime   time.Time `orm:"column(update_time);auto_now" json:"update_time"`
}

// UserMember ...
type UserMember struct {
	ID       int    `orm:"pk;column(user_id)" json:"user_id"`
	Username string `json:"username"`
	Rolename string `json:"role_name"`
	Role     int    `json:"role_id"`
}

type UserSearchResults []UserSearchResult

// UserSearchResult holds the information returned by the API when querying for a user name
type UserSearchResult struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
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
