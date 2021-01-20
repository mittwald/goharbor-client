// Code generated by go-swagger; DO NOT EDIT.

package legacy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AdminJobScheduleObj admin job schedule obj
//
// swagger:model AdminJobScheduleObj
type AdminJobScheduleObj struct {

	// A cron expression, a time-based job scheduler.
	Cron string `json:"cron,omitempty"`

	// The schedule type. The valid values are 'Hourly', 'Daily', 'Weekly', 'Custom', 'Manually' and 'None'.
	// 'Manually' means to trigger it right away and 'None' means to cancel the schedule.
	//
	Type string `json:"type,omitempty"`
}

// Validate validates this admin job schedule obj
func (m *AdminJobScheduleObj) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this admin job schedule obj based on context it is used
func (m *AdminJobScheduleObj) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AdminJobScheduleObj) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdminJobScheduleObj) UnmarshalBinary(b []byte) error {
	var res AdminJobScheduleObj
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
