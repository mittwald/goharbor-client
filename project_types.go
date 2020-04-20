package harbor

import "time"

// ProjectMetadata holds the metadata of a project.
type ProjectMetadata struct {
	ID        int64  `json:"id"`
	ProjectID int64  `json:"project_id"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	Deleted   bool   `json:"deleted"`
}

// Project holds the details of a project.
type Project struct {
	ProjectID    int64             `json:"project_id"`
	OwnerID      int64             `json:"owner_id"`
	Name         string            `json:"name"`
	CreationTime time.Time         `json:"creation_time"`
	UpdateTime   time.Time         `json:"update_time"`
	Deleted      bool              `json:"deleted"`
	OwnerName    string            `json:"owner_name"`
	Toggleable   bool              `json:"toggleable"`
	Role         int               `json:"current_user_role_id"`
	RepoCount    int64             `json:"repo_count"`
	Metadata     map[string]string `json:"metadata"`
	CVEWhitelist CVEWhitelist      `json:"CVEWhitelist"`
	StorageLimit int64             `json:"storageLimit"`
}

// Role holds the details of a role.
type Role struct {
	RoleID   int    `json:"role_id"`
	RoleCode string `json:"role_code"`
	Name     string `json:"role_name"`
	RoleMask int    `json:"role_mask"`
}

type RoleRequest struct {
	Role int64 `json:"role"`
}

// CVEWhitelistItem holds the CVE ids of a whitelisted item
type CVEWhitelistItem struct {
	CVEID string `json:"CVEID"`
}

// CVEWhitelist holds project specific information next to the set CVEWhitelistItem's
type CVEWhitelist struct {
	ID        int64            `json:"id"`
	ProjectID int64            `json:"projectID"`
	Items     CVEWhitelistItem `json:"items,optional"`
}

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

// ProjectRequest holds the information needed to create a project
type ProjectRequest struct {
	Name     string            `url:"name,omitempty" json:"project_name"`
	Public   *int              `url:"public,omitempty" json:"public"` //deprecated, reserved for project creation in replication
	Metadata map[string]string `url:"-" json:"metadata"`
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

// MemberRequest holds the information needed to update a project member
type MemberRequest struct {
	UserName string `json:"username"`
	Roles    []int  `json:"roles"`
}

// ProjectMemberRequest holds the information needed to add a project member
type ProjectMemberRequest struct {
	RoleID     int        `json:"role_id"`
	MemberUser MemberUser `json:"member_user"`
}

// MemberUser holds the user information needed for a project member request
type MemberUser struct {
	Username string `json:"username"`
	UserID   int    `json:"user_id"`
}
