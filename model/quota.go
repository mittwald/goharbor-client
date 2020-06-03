// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Quota The quota object
//
// swagger:model Quota
type Quota struct {

	// the creation time of the quota
	CreationTime string `json:"creation_time,omitempty"`

	// The hard limits of the quota
	Hard ResourceList `json:"hard,omitempty"`

	// ID of the quota
	ID int64 `json:"id,omitempty"`

	// The reference object of the quota
	Ref QuotaRefObject `json:"ref,omitempty"`

	// the update time of the quota
	UpdateTime string `json:"update_time,omitempty"`

	// The used status of the quota
	Used ResourceList `json:"used,omitempty"`
}

// Validate validates this quota
func (m *Quota) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHard(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Quota) validateHard(formats strfmt.Registry) error {

	if swag.IsZero(m.Hard) { // not required
		return nil
	}

	if err := m.Hard.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("hard")
		}
		return err
	}

	return nil
}

func (m *Quota) validateUsed(formats strfmt.Registry) error {

	if swag.IsZero(m.Used) { // not required
		return nil
	}

	if err := m.Used.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("used")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Quota) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Quota) UnmarshalBinary(b []byte) error {
	var res Quota
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
