// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LdapImportUsers ldap import users
//
// swagger:model LdapImportUsers
type LdapImportUsers struct {

	// selected uid list
	LdapUIDList []string `json:"ldap_uid_list"`
}

// Validate validates this ldap import users
func (m *LdapImportUsers) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LdapImportUsers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapImportUsers) UnmarshalBinary(b []byte) error {
	var res LdapImportUsers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
