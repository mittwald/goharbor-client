package harbor

import (
	"time"
)

// Label holds information used for a label
type Label struct {
	ID           int64     `orm:"pk;auto;column(id)" json:"id"`
	Name         string    `orm:"column(name)" json:"name"`
	Description  string    `orm:"column(description)" json:"description"`
	Color        string    `orm:"column(color)" json:"color"`
	Level        string    `orm:"column(level)" json:"-"`
	Scope        string    `orm:"column(scope)" json:"scope"`
	ProjectID    int64     `orm:"column(project_id)" json:"project_id"`
	CreationTime time.Time `orm:"column(creation_time);auto_now_add" json:"creation_time"`
	UpdateTime   time.Time `orm:"column(update_time);auto_now" json:"update_time"`
	Deleted      bool      `orm:"column(deleted)" json:"deleted"`
}

type ManifestResp struct {
	Manifest interface{} `json:"manifest"`
	Config   interface{} `json:"config,omitempty" `
}

// MemberReq -  Create Project Member Request
type MemberReq struct {
	ProjectID   int64     `json:"project_id"`
	Role        int       `json:"role_id,omitempty"`
	MemberUser  User      `json:"member_user,omitempty"`
	MemberGroup UserGroup `json:"member_group,omitempty"`
}

// UserGroup ...
type UserGroup struct {
	ID          int    `orm:"pk;auto;column(id)" json:"id,omitempty"`
	GroupName   string `orm:"column(group_name)" json:"group_name,omitempty"`
	GroupType   int    `orm:"column(group_type)" json:"group_type,omitempty"`
	LdapGroupDN string `orm:"column(ldap_group_dn)" json:"ldap_group_dn,omitempty"`
}

type RepoResp struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	ProjectID   int64  `json:"project_id"`
	Description string `json:"description"`
	PullCount   int64  `json:"pull_count"`
	StarCount   int64  `json:"star_count"`
	TagsCount   int64  `json:"tags_count"`
}

// RepositoryQuery : query parameters for repository
type RepositoryQuery struct {
	Name        string
	ProjectIDs  []int64
	ProjectName string
	LabelID     int64
	Pagination
	Sorting
}

// RepoRecord holds the record of an repository in DB, all the infors are from the registry notification event.
type RepoRecord struct {
	RepositoryID int64     `orm:"pk;auto;column(repository_id)" json:"repository_id"`
	Name         string    `orm:"column(name)" json:"name"`
	ProjectID    int64     `orm:"column(project_id)"  json:"project_id"`
	Description  string    `orm:"column(description)" json:"description"`
	PullCount    int64     `orm:"column(pull_count)" json:"pull_count"`
	StarCount    int64     `orm:"column(star_count)" json:"star_count"`
	CreationTime time.Time `orm:"column(creation_time);auto_now_add" json:"creation_time"`
	UpdateTime   time.Time `orm:"column(update_time);auto_now" json:"update_time"`
}

// Role holds the details of a role.
type Role struct {
	RoleID   int    `orm:"pk;auto;column(role_id)" json:"role_id"`
	RoleCode string `orm:"column(role_code)" json:"role_code"`
	Name     string `orm:"column(name)" json:"role_name"`
	RoleMask int    `orm:"column(role_mask)" json:"role_mask"`
}

// Pagination ...
type Pagination struct {
	Page int64
	Size int64
}

type Hashes map[string][]byte

// Signature ...
type Signature struct {
	Tag    string      `json:"tag"`
	Hashes Hashes `json:"hashes"`
}

// Sorting sort by given field, ascending or descending
type Sorting struct {
	Sort string // in format [+-]?<FIELD_NAME>, e.g. '+creation_time', '-creation_time'
}

// TagCfg ...
type TagCfg struct {
	Labels map[string]string `json:"labels"`
}

// TagDetail ...
type TagDetail struct {
	Digest        string    `json:"digest"`
	Name          string    `json:"name"`
	Size          int64     `json:"size"`
	Architecture  string    `json:"architecture"`
	OS            string    `json:"os"`
	OSVersion     string    `json:"os.version"`
	DockerVersion string    `json:"docker_version"`
	Author        string    `json:"author"`
	Created       time.Time `json:"created"`
	Config        *TagCfg   `json:"config"`
	Immutable     bool      `json:"immutable"`
}

// TagResp holds the information of one image tag
type TagResp struct {
	TagDetail
	Signature    *Target                `json:"signature"`
	ScanOverview map[string]interface{} `json:"scan_overview,omitempty"`
	Labels       []*Label               `json:"labels"`
	PushTime     time.Time              `json:"push_time"`
	PullTime     time.Time              `json:"pull_time"`
}

type Target struct {
	Tag    string      `json:"tag"`
	Hashes Hashes `json:"hashes"`
}

// VulnerabilityItem is an item in the vulnerability result returned by the vulnerability details API
type VulnerabilityItem struct {
	ID          string `json:"id"`
	Severity    int64  `json:"severity"`
	Pkg         string `json:"package"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Fixed       string `json:"fixedVersion,omitempty"`
}
