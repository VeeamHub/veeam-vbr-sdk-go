// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BackupFSExclusionsModel VM guest OS file exclusion.
//
// swagger:model BackupFSExclusionsModel
type BackupFSExclusionsModel struct {

	// exclusion policy
	// Required: true
	ExclusionPolicy *EBackupExclusionPolicy `json:"exclusionPolicy"`

	// Array of files and folders. Full paths to files and folders, environmental variables and file masks with the asterisk (*) and question mark (?) characters can be used.
	ItemsList []string `json:"itemsList"`
}

// Validate validates this backup f s exclusions model
func (m *BackupFSExclusionsModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExclusionPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupFSExclusionsModel) validateExclusionPolicy(formats strfmt.Registry) error {

	if err := validate.Required("exclusionPolicy", "body", m.ExclusionPolicy); err != nil {
		return err
	}

	if err := validate.Required("exclusionPolicy", "body", m.ExclusionPolicy); err != nil {
		return err
	}

	if m.ExclusionPolicy != nil {
		if err := m.ExclusionPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exclusionPolicy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this backup f s exclusions model based on the context it is used
func (m *BackupFSExclusionsModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExclusionPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupFSExclusionsModel) contextValidateExclusionPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.ExclusionPolicy != nil {
		if err := m.ExclusionPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exclusionPolicy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupFSExclusionsModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupFSExclusionsModel) UnmarshalBinary(b []byte) error {
	var res BackupFSExclusionsModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}