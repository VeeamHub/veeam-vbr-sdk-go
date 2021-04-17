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

// BackupStorageSettingsEncryptionModel Encription of backup files.
//
// swagger:model BackupStorageSettingsEncryptionModel
type BackupStorageSettingsEncryptionModel struct {

	// ID of the password used for encryption. The value is *null* for exported objects.
	// Format: uuid
	EncryptionPasswordIDOrNull strfmt.UUID `json:"encryptionPasswordIdOrNull,omitempty"`

	// Tag used to identify the password.
	EncryptionPasswordTag string `json:"encryptionPasswordTag,omitempty"`

	// If *true*, the content of backup files is encrypted.
	// Required: true
	IsEnabled *bool `json:"isEnabled"`
}

// Validate validates this backup storage settings encryption model
func (m *BackupStorageSettingsEncryptionModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEncryptionPasswordIDOrNull(formats); err != nil {
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

func (m *BackupStorageSettingsEncryptionModel) validateEncryptionPasswordIDOrNull(formats strfmt.Registry) error {
	if swag.IsZero(m.EncryptionPasswordIDOrNull) { // not required
		return nil
	}

	if err := validate.FormatOf("encryptionPasswordIdOrNull", "body", "uuid", m.EncryptionPasswordIDOrNull.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BackupStorageSettingsEncryptionModel) validateIsEnabled(formats strfmt.Registry) error {

	if err := validate.Required("isEnabled", "body", m.IsEnabled); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this backup storage settings encryption model based on context it is used
func (m *BackupStorageSettingsEncryptionModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BackupStorageSettingsEncryptionModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupStorageSettingsEncryptionModel) UnmarshalBinary(b []byte) error {
	var res BackupStorageSettingsEncryptionModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}