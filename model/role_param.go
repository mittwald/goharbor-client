// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RoleParam role param
//
// swagger:model RoleParam
type RoleParam struct {

	// Role ID for updating project role member.
	Roles []int32 `json:"roles"`

	// Username relevant to a project role member.
	Username string `json:"username,omitempty"`
}

// Validate validates this role param
func (m *RoleParam) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RoleParam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoleParam) UnmarshalBinary(b []byte) error {
	var res RoleParam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
