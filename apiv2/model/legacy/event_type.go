// Code generated by go-swagger; DO NOT EDIT.

package legacy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// EventType Webhook supportted event type.
// Example: pullImage
//
// swagger:model EventType
type EventType string

// Validate validates this event type
func (m EventType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this event type based on context it is used
func (m EventType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
