// Code generated by go-swagger; DO NOT EDIT.

package legacy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Password password
//
// swagger:model Password
type Password struct {

	// New password for marking as to be updated.
	NewPassword string `json:"new_password,omitempty"`

	// The user's existing password.
	OldPassword string `json:"old_password,omitempty"`
}

// Validate validates this password
func (m *Password) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this password based on context it is used
func (m *Password) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Password) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Password) UnmarshalBinary(b []byte) error {
	var res Password
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
