// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserGroup user group
//
// swagger:model UserGroup
type UserGroup struct {

	// The name of the user group
	GroupName string `json:"group_name,omitempty"`

	// The group type, 1 for LDAP group, 2 for HTTP group, 3 for OIDC group.
	GroupType int64 `json:"group_type,omitempty"`

	// The ID of the user group
	ID int64 `json:"id,omitempty"`

	// The DN of the LDAP group if group type is 1 (LDAP group).
	LdapGroupDn string `json:"ldap_group_dn,omitempty"`
}

// Validate validates this user group
func (m *UserGroup) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserGroup) UnmarshalBinary(b []byte) error {
	var res UserGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
