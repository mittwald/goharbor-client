package harbor

import "time"

// To ensure type-safe queries to the harbor API,
// the following typings include typings from the upstream sources:
// https://github.com/goharbor/harbor/tree/v1.10.2/src/replication/dao/models/

// ReplicationExecution holds information about one replication execution.
type ReplicationExecution struct {
	ID         int64       `orm:"pk;auto;column(id)" json:"id"`
	PolicyID   int64       `orm:"column(policy_id)" json:"policy_id"`
	Status     string      `orm:"column(status)" json:"status"`
	StatusText string      `orm:"column(status_text)" json:"status_text"`
	Total      int         `orm:"column(total)" json:"total"`
	Failed     int         `orm:"column(failed)" json:"failed"`
	Succeed    int         `orm:"column(succeed)" json:"succeed"`
	InProgress int         `orm:"column(in_progress)" json:"in_progress"`
	Stopped    int         `orm:"column(stopped)" json:"stopped"`
	Trigger    TriggerType `orm:"column(trigger)" json:"trigger"`
	StartTime  time.Time   `orm:"column(start_time)" json:"start_time"`
	EndTime    time.Time   `orm:"column(end_time)" json:"end_time"`
}

// ReplicationPolicy defines the structure of a replication policy
type ReplicationPolicy struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Creator     string `json:"creator"`
	// source
	SrcRegistry *Registry `json:"src_registry"`
	// destination
	DestRegistry *Registry `json:"dest_registry"`
	// Only support two dest namespace modes:
	// Put all the src resources to the one single dest namespace
	// or keep namespaces same with the source ones (under this case,
	// the DestNamespace should be set to empty)
	DestNamespace string `json:"dest_namespace"`
	// Filters
	Filters []*Filter `json:"filters"`
	// Trigger
	Trigger *Trigger `json:"trigger"`
	// Settings
	Deletion bool `json:"deletion"`
	// If override the image tag
	Override bool `json:"override"`
	// Operations
	Enabled      bool      `json:"enabled"`
	CreationTime time.Time `json:"creation_time"`
	UpdateTime   time.Time `json:"update_time"`
}

// Filter holds the info of the filter
type Filter struct {
	Type  FilterType  `json:"type"`
	Value interface{} `json:"value"`
}

// Trigger holds info for a trigger
type Trigger struct {
	Type     TriggerType      `json:"type"`
	Settings *TriggerSettings `json:"trigger_settings"`
}

// TriggerSettings holds the settings of a trigger
type TriggerSettings struct {
	Cron string `json:"cron"`
}
