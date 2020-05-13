package harbor

// To ensure type-safe queries to the harbor API,
// the following typings include typings from the upstream sources:
// https://github.com/goharbor/harbor/src/core/api/models

const (
	// ScheduleHourly : 'Hourly'
	ScheduleHourly = "Hourly"
	// ScheduleDaily : 'Daily'
	ScheduleDaily = "Daily"
	// ScheduleWeekly : 'Weekly'
	ScheduleWeekly = "Weekly"
	// ScheduleCustom : 'Custom'
	ScheduleCustom = "Custom"
	// ScheduleManual : 'Manual'
	ScheduleManual = "Manual"
	// ScheduleNone : 'None'
	ScheduleNone = "None"
)

// AdminJobReq holds request information for admin job
type AdminJobReq struct {
	AdminJobSchedule
	Name       string                 `json:"name"`
	Status     string                 `json:"status"`
	ID         int64                  `json:"id"`
	Parameters map[string]interface{} `json:"parameters"`
}

// AdminJobSchedule ...
type AdminJobSchedule struct {
	Schedule *ScheduleParam `json:"schedule"`
}

// ScheduleParam defines the parameter of schedule trigger
type ScheduleParam struct {
	// Daily, Weekly, Custom, Manual, None
	Type string `json:"type"`
	// The cron string of scheduled job
	Cron string `json:"cron"`
}
