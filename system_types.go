package harbor

// To ensure type-safe queries to the harbor API,
// the following typings include typings from the upstream sources:
// https://github.com/goharbor/harbor/tree/v1.10.2/src/core/api/models

type ScheduleType string

const (
	ScheduleTypeHourly ScheduleType = "Hourly"
	ScheduleTypeDaily               = "Daily"
	ScheduleTypeWeekly              = "Weekly"
	ScheduleTypeCustom              = "Custom"
	ScheduleTypeManual              = "Manual"
	ScheduleTypeNone                = "None"
)

// AdminJobReq holds request information for admin job
type AdminJobReq struct {
	AdminJobSchedule
	Name       string                 `json:"name"`
	Status     string                 `json:"status"`
	ID         int64                  `json:"id"`
	Parameters map[string]interface{} `json:"parameters"`
}

// AdminJobSchedule holds the information of an admin job schedule
type AdminJobSchedule struct {
	Schedule *ScheduleParam `json:"schedule"`
}

// ScheduleParam defines the parameters of a schedule trigger
type ScheduleParam struct {
	// Daily, Hourly, Weekly, Custom, Manual, None
	// Note: When creating pre-defined schedule types (e.g. 'Hourly'), the Cron string has to be provided.
	Type ScheduleType `json:"type"`
	// The cron string of scheduled job
	Cron string `json:"cron"`
}
