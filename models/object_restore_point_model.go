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

// ObjectRestorePointModel object restore point model
//
// swagger:model ObjectRestorePointModel
type ObjectRestorePointModel struct {

	// Array of operations allowed for the restore point.
	// Required: true
	AllowedOperations []EObjectRestorePointOperation `json:"allowedOperations"`

	// ID of a backup that contains the restore point.
	// Required: true
	// Format: uuid
	BackupID *strfmt.UUID `json:"backupId"`

	// Date and time when the restore point was created.
	// Required: true
	// Format: date-time
	CreationTime *strfmt.DateTime `json:"creationTime"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// ID of a platform on which the object was created.
	// Required: true
	// Format: uuid
	PlatformID *strfmt.UUID `json:"platformId"`

	// platform name
	// Required: true
	PlatformName *EPlatformType `json:"platformName"`
}

// Validate validates this object restore point model
func (m *ObjectRestorePointModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowedOperations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatformID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatformName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectRestorePointModel) validateAllowedOperations(formats strfmt.Registry) error {

	if err := validate.Required("allowedOperations", "body", m.AllowedOperations); err != nil {
		return err
	}

	for i := 0; i < len(m.AllowedOperations); i++ {

		if err := m.AllowedOperations[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allowedOperations" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ObjectRestorePointModel) validateBackupID(formats strfmt.Registry) error {

	if err := validate.Required("backupId", "body", m.BackupID); err != nil {
		return err
	}

	if err := validate.FormatOf("backupId", "body", "uuid", m.BackupID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ObjectRestorePointModel) validateCreationTime(formats strfmt.Registry) error {

	if err := validate.Required("creationTime", "body", m.CreationTime); err != nil {
		return err
	}

	if err := validate.FormatOf("creationTime", "body", "date-time", m.CreationTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ObjectRestorePointModel) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ObjectRestorePointModel) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ObjectRestorePointModel) validatePlatformID(formats strfmt.Registry) error {

	if err := validate.Required("platformId", "body", m.PlatformID); err != nil {
		return err
	}

	if err := validate.FormatOf("platformId", "body", "uuid", m.PlatformID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ObjectRestorePointModel) validatePlatformName(formats strfmt.Registry) error {

	if err := validate.Required("platformName", "body", m.PlatformName); err != nil {
		return err
	}

	if err := validate.Required("platformName", "body", m.PlatformName); err != nil {
		return err
	}

	if m.PlatformName != nil {
		if err := m.PlatformName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platformName")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this object restore point model based on the context it is used
func (m *ObjectRestorePointModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllowedOperations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePlatformName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectRestorePointModel) contextValidateAllowedOperations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AllowedOperations); i++ {

		if err := m.AllowedOperations[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allowedOperations" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ObjectRestorePointModel) contextValidatePlatformName(ctx context.Context, formats strfmt.Registry) error {

	if m.PlatformName != nil {
		if err := m.PlatformName.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platformName")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ObjectRestorePointModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectRestorePointModel) UnmarshalBinary(b []byte) error {
	var res ObjectRestorePointModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}