// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DetailedTag detailed tag
//
// swagger:model DetailedTag
type DetailedTag struct {

	// The architecture of the image.
	Architecture string `json:"architecture,omitempty"`

	// The author of the image.
	Author string `json:"author,omitempty"`

	// The build time of the image.
	Created string `json:"created,omitempty"`

	// The digest of the tag.
	Digest string `json:"digest,omitempty"`

	// The version of docker which builds the image.
	DockerVersion string `json:"docker_version,omitempty"`

	// The label list.
	Labels []*Label `json:"labels"`

	// The name of the tag.
	Name string `json:"name,omitempty"`

	// The os of the image.
	Os string `json:"os,omitempty"`

	// The overview of the scan result.
	ScanOverview ScanOverview `json:"scan_overview,omitempty"`

	// The signature of image, defined by RepoSignature. If it is null, the image is unsigned.
	Signature interface{} `json:"signature,omitempty"`

	// The size of the image.
	Size int64 `json:"size,omitempty"`
}

// Validate validates this detailed tag
func (m *DetailedTag) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScanOverview(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DetailedTag) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	for i := 0; i < len(m.Labels); i++ {
		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {
			if err := m.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DetailedTag) validateScanOverview(formats strfmt.Registry) error {
	if swag.IsZero(m.ScanOverview) { // not required
		return nil
	}

	if m.ScanOverview != nil {
		if err := m.ScanOverview.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scan_overview")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this detailed tag based on the context it is used
func (m *DetailedTag) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScanOverview(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DetailedTag) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Labels); i++ {

		if m.Labels[i] != nil {
			if err := m.Labels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DetailedTag) contextValidateScanOverview(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ScanOverview.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("scan_overview")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DetailedTag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DetailedTag) UnmarshalBinary(b []byte) error {
	var res DetailedTag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
