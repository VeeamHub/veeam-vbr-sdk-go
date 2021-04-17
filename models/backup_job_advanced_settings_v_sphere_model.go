// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BackupJobAdvancedSettingsVSphereModel VMware vSphere settings for the backup job.
//
// swagger:model BackupJobAdvancedSettingsVSphereModel
type BackupJobAdvancedSettingsVSphereModel struct {

	// changed block tracking
	ChangedBlockTracking *VSphereChangedBlockTrackingSettingsModel `json:"changedBlockTracking,omitempty"`

	// If *true*, VMware Tools quiescence is enabled for freezing the VM file system and application data.
	EnableVMWareToolsQuiescence bool `json:"enableVMWareToolsQuiescence,omitempty"`
}

// Validate validates this backup job advanced settings v sphere model
func (m *BackupJobAdvancedSettingsVSphereModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChangedBlockTracking(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupJobAdvancedSettingsVSphereModel) validateChangedBlockTracking(formats strfmt.Registry) error {
	if swag.IsZero(m.ChangedBlockTracking) { // not required
		return nil
	}

	if m.ChangedBlockTracking != nil {
		if err := m.ChangedBlockTracking.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changedBlockTracking")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this backup job advanced settings v sphere model based on the context it is used
func (m *BackupJobAdvancedSettingsVSphereModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChangedBlockTracking(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupJobAdvancedSettingsVSphereModel) contextValidateChangedBlockTracking(ctx context.Context, formats strfmt.Registry) error {

	if m.ChangedBlockTracking != nil {
		if err := m.ChangedBlockTracking.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changedBlockTracking")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupJobAdvancedSettingsVSphereModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupJobAdvancedSettingsVSphereModel) UnmarshalBinary(b []byte) error {
	var res BackupJobAdvancedSettingsVSphereModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}