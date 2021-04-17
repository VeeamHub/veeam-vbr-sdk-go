// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GuestFileSystemIndexingModel VM guest OS file indexing.
//
// swagger:model GuestFileSystemIndexingModel
type GuestFileSystemIndexingModel struct {

	// Array of VMs with guest OS file indexing options.
	IndexingSettings []*BackupIndexingSettingsModel `json:"indexingSettings"`

	// is enabled
	// Required: true
	IsEnabled *bool `json:"isEnabled"`
}

// Validate validates this guest file system indexing model
func (m *GuestFileSystemIndexingModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIndexingSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsEnabled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GuestFileSystemIndexingModel) validateIndexingSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.IndexingSettings) { // not required
		return nil
	}

	for i := 0; i < len(m.IndexingSettings); i++ {
		if swag.IsZero(m.IndexingSettings[i]) { // not required
			continue
		}

		if m.IndexingSettings[i] != nil {
			if err := m.IndexingSettings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("indexingSettings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GuestFileSystemIndexingModel) validateIsEnabled(formats strfmt.Registry) error {

	if err := validate.Required("isEnabled", "body", m.IsEnabled); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this guest file system indexing model based on the context it is used
func (m *GuestFileSystemIndexingModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIndexingSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GuestFileSystemIndexingModel) contextValidateIndexingSettings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IndexingSettings); i++ {

		if m.IndexingSettings[i] != nil {
			if err := m.IndexingSettings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("indexingSettings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GuestFileSystemIndexingModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GuestFileSystemIndexingModel) UnmarshalBinary(b []byte) error {
	var res GuestFileSystemIndexingModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}