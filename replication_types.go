package harbor

import "time"

// To ensure type-safe queries to the harbor API,
// the following typings include typings from the upstream sources:
// https://github.com/goharbor/harbor/src/replication/dao/models/

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

// ReplicationPolicy is the model for a ng replication policy.
type ReplicationPolicy struct {
	ID                int64     `orm:"pk;auto;column(id)" json:"id"`
	Name              string    `orm:"column(name)" json:"name"`
	Description       string    `orm:"column(description)" json:"description"`
	Creator           string    `orm:"column(creator)" json:"creator"`
	SrcRegistryID     int64     `orm:"column(src_registry_id)" json:"src_registry_id"`
	DestRegistryID    int64     `orm:"column(dest_registry_id)" json:"dest_registry_id"`
	DestNamespace     string    `orm:"column(dest_namespace)" json:"dest_namespace"`
	Override          bool      `orm:"column(override)" json:"override"`
	Enabled           bool      `orm:"column(enabled)" json:"enabled"`
	Trigger           string    `orm:"column(trigger)" json:"trigger"`
	Filters           string    `orm:"column(filters)" json:"filters"`
	ReplicateDeletion bool      `orm:"column(replicate_deletion)" json:"replicate_deletion"`
	CreationTime      time.Time `orm:"column(creation_time);auto_now_add" json:"creation_time"`
	UpdateTime        time.Time `orm:"column(update_time);auto_now" json:"update_time"`
}
