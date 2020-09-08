// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InstanceCreatedResp instance created resp
//
// swagger:model InstanceCreatedResp
type InstanceCreatedResp struct {

	// ID of instance created
	ID int64 `json:"id,omitempty"`
}

// Validate validates this instance created resp
func (m *InstanceCreatedResp) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InstanceCreatedResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceCreatedResp) UnmarshalBinary(b []byte) error {
	var res InstanceCreatedResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
