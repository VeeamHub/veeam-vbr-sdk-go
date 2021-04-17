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

// LinuxHostImportSpec linux host import spec
//
// swagger:model LinuxHostImportSpec
type LinuxHostImportSpec struct {

	// credentials
	// Required: true
	Credentials *CredentialsImportModel `json:"credentials"`

	// Description of the server.
	// Required: true
	Description *string `json:"description"`

	// Full DNS name or IP address of the server.
	// Required: true
	Name *string `json:"name"`

	// SSH key fingerprint used to verify the server identity.
	// Required: true
	SSHFingerprint *string `json:"sshFingerprint"`

	// ssh settings
	SSHSettings *LinuxHostSSHSettingsModel `json:"sshSettings,omitempty"`

	// type
	// Required: true
	Type *EManagedServerType `json:"type"`
}

// Validate validates this linux host import spec
func (m *LinuxHostImportSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSSHFingerprint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSSHSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LinuxHostImportSpec) validateCredentials(formats strfmt.Registry) error {

	if err := validate.Required("credentials", "body", m.Credentials); err != nil {
		return err
	}

	if m.Credentials != nil {
		if err := m.Credentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

func (m *LinuxHostImportSpec) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *LinuxHostImportSpec) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *LinuxHostImportSpec) validateSSHFingerprint(formats strfmt.Registry) error {

	if err := validate.Required("sshFingerprint", "body", m.SSHFingerprint); err != nil {
		return err
	}

	return nil
}

func (m *LinuxHostImportSpec) validateSSHSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.SSHSettings) { // not required
		return nil
	}

	if m.SSHSettings != nil {
		if err := m.SSHSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sshSettings")
			}
			return err
		}
	}

	return nil
}

func (m *LinuxHostImportSpec) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this linux host import spec based on the context it is used
func (m *LinuxHostImportSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSSHSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LinuxHostImportSpec) contextValidateCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.Credentials != nil {
		if err := m.Credentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

func (m *LinuxHostImportSpec) contextValidateSSHSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.SSHSettings != nil {
		if err := m.SSHSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sshSettings")
			}
			return err
		}
	}

	return nil
}

func (m *LinuxHostImportSpec) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LinuxHostImportSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinuxHostImportSpec) UnmarshalBinary(b []byte) error {
	var res LinuxHostImportSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}