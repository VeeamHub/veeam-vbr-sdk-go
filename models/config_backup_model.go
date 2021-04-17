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

// ConfigBackupModel config backup model
//
// swagger:model ConfigBackupModel
type ConfigBackupModel struct {

	// ID of the backup repository on which the configuration backup is stored.
	// Required: true
	// Format: uuid
	BackupRepositoryID *strfmt.UUID `json:"backupRepositoryId"`

	// encryption
	// Required: true
	Encryption *ConfigBackupEncryptionModel `json:"encryption"`

	// If *true*, configuration backup is enabled.
	// Required: true
	IsEnabled *bool `json:"isEnabled"`

	// last successful backup
	// Required: true
	LastSuccessfulBackup *ConfigBackupLastSuccessfulModel `json:"lastSuccessfulBackup"`

	// notifications
	// Required: true
	Notifications *ConfigBackupNotificationsModel `json:"notifications"`

	// Number of restore points to keep in the backup repository.
	// Required: true
	RestorePointsToKeep *int32 `json:"restorePointsToKeep"`

	// schedule
	// Required: true
	Schedule *ConfigBackupScheduleModel `json:"schedule"`
}

// Validate validates this config backup model
func (m *ConfigBackupModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupRepositoryID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncryption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastSuccessfulBackup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotifications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestorePointsToKeep(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigBackupModel) validateBackupRepositoryID(formats strfmt.Registry) error {

	if err := validate.Required("backupRepositoryId", "body", m.BackupRepositoryID); err != nil {
		return err
	}

	if err := validate.FormatOf("backupRepositoryId", "body", "uuid", m.BackupRepositoryID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConfigBackupModel) validateEncryption(formats strfmt.Registry) error {

	if err := validate.Required("encryption", "body", m.Encryption); err != nil {
		return err
	}

	if m.Encryption != nil {
		if err := m.Encryption.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryption")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigBackupModel) validateIsEnabled(formats strfmt.Registry) error {

	if err := validate.Required("isEnabled", "body", m.IsEnabled); err != nil {
		return err
	}

	return nil
}

func (m *ConfigBackupModel) validateLastSuccessfulBackup(formats strfmt.Registry) error {

	if err := validate.Required("lastSuccessfulBackup", "body", m.LastSuccessfulBackup); err != nil {
		return err
	}

	if m.LastSuccessfulBackup != nil {
		if err := m.LastSuccessfulBackup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastSuccessfulBackup")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigBackupModel) validateNotifications(formats strfmt.Registry) error {

	if err := validate.Required("notifications", "body", m.Notifications); err != nil {
		return err
	}

	if m.Notifications != nil {
		if err := m.Notifications.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notifications")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigBackupModel) validateRestorePointsToKeep(formats strfmt.Registry) error {

	if err := validate.Required("restorePointsToKeep", "body", m.RestorePointsToKeep); err != nil {
		return err
	}

	return nil
}

func (m *ConfigBackupModel) validateSchedule(formats strfmt.Registry) error {

	if err := validate.Required("schedule", "body", m.Schedule); err != nil {
		return err
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config backup model based on the context it is used
func (m *ConfigBackupModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEncryption(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastSuccessfulBackup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNotifications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigBackupModel) contextValidateEncryption(ctx context.Context, formats strfmt.Registry) error {

	if m.Encryption != nil {
		if err := m.Encryption.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryption")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigBackupModel) contextValidateLastSuccessfulBackup(ctx context.Context, formats strfmt.Registry) error {

	if m.LastSuccessfulBackup != nil {
		if err := m.LastSuccessfulBackup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastSuccessfulBackup")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigBackupModel) contextValidateNotifications(ctx context.Context, formats strfmt.Registry) error {

	if m.Notifications != nil {
		if err := m.Notifications.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notifications")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigBackupModel) contextValidateSchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.Schedule != nil {
		if err := m.Schedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigBackupModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigBackupModel) UnmarshalBinary(b []byte) error {
	var res ConfigBackupModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}