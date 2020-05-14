package harbor

import "time"

// To ensure type-safe queries to the harbor API,
// the following typings include typings from the upstream sources:
// https://github.com/goharbor/harbor/tree/v1.10.2/src/common/models

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

// CVEWhitelist defines the data model for a CVE whitelist
type CVEWhitelist struct {
	ID           int64              `orm:"pk;auto;column(id)" json:"id"`
	ProjectID    int64              `orm:"column(project_id)" json:"project_id"`
	ExpiresAt    *int64             `orm:"column(expires_at)" json:"expires_at,omitempty"`
	Items        []CVEWhitelistItem `orm:"-" json:"items"`
	ItemsText    string             `orm:"column(items)" json:"-"`
	CreationTime time.Time          `orm:"column(creation_time);auto_now_add" json:"creation_time"`
	UpdateTime   time.Time          `orm:"column(update_time);auto_now" json:"update_time"`
}

// CVEWhitelistItem defines one item in the CVE whitelist
type CVEWhitelistItem struct {
	CVEID string `json:"cve_id"`
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

// Member holds the details of a member.
type Member struct {
	ID         int    `orm:"pk;column(id)" json:"id"`
	ProjectID  int64  `orm:"column(project_id)" json:"project_id"`
	Entityname string `orm:"column(entity_name)" json:"entity_name"`
	Rolename   string `json:"role_name"`
	Role       int    `json:"role_id"`
	EntityID   int    `orm:"column(entity_id)" json:"entity_id"`
	EntityType string `orm:"column(entity_type)" json:"entity_type"`
}

// UserMember ...
type UserMember struct {
	ID       int    `orm:"pk;column(user_id)" json:"user_id"`
	Username string `json:"username"`
	Rolename string `json:"role_name"`
	Role     int    `json:"role_id"`
}

// MemberReq -  Create Project Member Request
type MemberReq struct {
	ProjectID   int64     `json:"project_id"`
	Role        int       `json:"role_id,omitempty"`
	MemberUser  User      `json:"member_user,omitempty"`
	MemberGroup UserGroup `json:"member_group,omitempty"`
}

// Project holds the details of a project.
type Project struct {
	ProjectID    int64             `orm:"pk;auto;column(project_id)" json:"project_id"`
	OwnerID      int               `orm:"column(owner_id)" json:"owner_id"`
	Name         string            `orm:"column(name)" json:"name"`
	CreationTime time.Time         `orm:"column(creation_time);auto_now_add" json:"creation_time"`
	UpdateTime   time.Time         `orm:"column(update_time);auto_now" json:"update_time"`
	Deleted      bool              `orm:"column(deleted)" json:"deleted"`
	OwnerName    string            `orm:"-" json:"owner_name"`
	Role         int               `orm:"-" json:"current_user_role_id"`
	RoleList     []int             `orm:"-" json:"current_user_role_ids"`
	RepoCount    int64             `orm:"-" json:"repo_count"`
	ChartCount   uint64            `orm:"-" json:"chart_count"`
	Metadata     map[string]string `orm:"-" json:"metadata"`
	CVEWhitelist CVEWhitelist      `orm:"-" json:"cve_whitelist"`
}

// ProjectRequest holds informations that need for creating project API
type ProjectRequest struct {
	Name         string            `json:"project_name"`
	Public       *int              `json:"public"` // deprecated, reserved for project creation in replication
	Metadata     map[string]string `json:"metadata"`
	CVEWhitelist CVEWhitelist      `json:"cve_whitelist"`

	StorageLimit *int64 `json:"storage_limit,omitempty"`
}

// RoleRequest holds the information of a user's role
type RoleRequest struct {
	Role int `json:"role"`
}
