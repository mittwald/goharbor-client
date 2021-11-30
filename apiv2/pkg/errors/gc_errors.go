package errors

const (
	// ErrSystemInvalidScheduleMsg describes an invalid schedule type request
	ErrSystemInvalidScheduleMsg = "invalid schedule type"

	// ErrSystemGcInProgressMsg describes that a gc progress is already running
	ErrSystemGcInProgressMsg = "the system is already running a gc job"

	// ErrSystemUnauthorizedMsg describes an unauthorized request
	ErrSystemUnauthorizedMsg = "unauthorized"

	// ErrSystemInternalErrorsMsg describes server-side internal errors
	ErrSystemInternalErrorsMsg = "unexpected internal errors"

	// ErrSystemNoPermissionMsg describes a request error without permission
	ErrSystemNoPermissionMsg = "user does not have permission to the System"

	// ErrSystemGcUndefinedMsg describes a server-side response returning an empty GC schedule
	ErrSystemGcUndefinedMsg = "no schedule defined"

	// ErrSystemGcScheduleIdenticalMsg describes equality between two GC schedules
	ErrSystemGcScheduleIdenticalMsg = "the provided schedule is identical to the existing schedule"

	// ErrSystemGcScheduleNotProvidedMsg describes the absence of a required schedule
	ErrSystemGcScheduleNotProvidedMsg = "no schedule provided"

	// ErrSystemGcScheduleUndefinedMsg describes an error when the GC schedule is undefined
	ErrSystemGcScheduleUndefinedMsg = "the garbage collection schedule is undefined"

	// ErrSystemGcScheduleParametersUndefinedMsg describes an error when a GC schedule's parameters are undefined
	ErrSystemGcScheduleParametersUndefinedMsg = "garbage collection schedule parameters are undefined"
)

// ErrSystemInvalidSchedule describes an invalid schedule type request.
type ErrSystemInvalidSchedule struct{}

// Error returns the error message.
func (e *ErrSystemInvalidSchedule) Error() string {
	return ErrSystemInvalidScheduleMsg
}

// ErrSystemGcInProgress describes that a gc progress is already running.
type ErrSystemGcInProgress struct{}

// Error returns the error message.
func (e *ErrSystemGcInProgress) Error() string {
	return ErrSystemGcInProgressMsg
}

// ErrSystemUnauthorized describes an unauthorized request.
type ErrSystemUnauthorized struct{}

// Error returns the error message.
func (e *ErrSystemUnauthorized) Error() string {
	return ErrSystemUnauthorizedMsg
}

// ErrSystemInternalErrors describes server-side internal errors.
type ErrSystemInternalErrors struct{}

// Error returns the error message.
func (e *ErrSystemInternalErrors) Error() string {
	return ErrSystemInternalErrorsMsg
}

// ErrSystemNoPermission describes a request error without permission.
type ErrSystemNoPermission struct{}

// Error returns the error message.
func (e *ErrSystemNoPermission) Error() string {
	return ErrSystemNoPermissionMsg
}

// ErrSystemGcUndefined describes a server-side response returning an empty GC schedule.
type ErrSystemGcUndefined struct{}

// Error returns the error message.
func (e *ErrSystemGcUndefined) Error() string {
	return ErrSystemGcUndefinedMsg
}

// ErrSystemGcScheduleIdentical describes equality between two GC schedules.
type ErrSystemGcScheduleIdentical struct{}

// Error returns the error message.
func (e *ErrSystemGcScheduleIdentical) Error() string {
	return ErrSystemGcScheduleIdenticalMsg
}

// ErrSystemGcScheduleNotProvided describes the absence of a required schedule.
type ErrSystemGcScheduleNotProvided struct{}

// Error returns the error message.
func (e *ErrSystemGcScheduleNotProvided) Error() string {
	return ErrSystemGcScheduleNotProvidedMsg
}

// ErrSystemGcScheduleUndefined describes an error when the fetched gc schedule is undefined.
type ErrSystemGcScheduleUndefined struct{}

// Error returns the error message.
func (e *ErrSystemGcScheduleUndefined) Error() string {
	return ErrSystemGcScheduleUndefinedMsg
}

// ErrSystemGcScheduleParametersUndefined describes an error when a GC schedule's parameters are undefined
type ErrSystemGcScheduleParametersUndefined struct{}

// Error returns the error message.
func (e *ErrSystemGcScheduleParametersUndefined) Error() string {
	return ErrSystemGcScheduleParametersUndefinedMsg
}
