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

// ViProxyModel vi proxy model
//
// swagger:model ViProxyModel
type ViProxyModel struct {
	ProxyModel

	// server
	// Required: true
	Server *ProxyServerSettingsModel `json:"server"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ViProxyModel) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ProxyModel
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ProxyModel = aO0

	// AO1
	var dataAO1 struct {
		Server *ProxyServerSettingsModel `json:"server"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Server = dataAO1.Server

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ViProxyModel) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ProxyModel)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Server *ProxyServerSettingsModel `json:"server"`
	}

	dataAO1.Server = m.Server

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this vi proxy model
func (m *ViProxyModel) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ProxyModel
	if err := m.ProxyModel.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ViProxyModel) validateServer(formats strfmt.Registry) error {

	if err := validate.Required("server", "body", m.Server); err != nil {
		return err
	}

	if m.Server != nil {
		if err := m.Server.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("server")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this vi proxy model based on the context it is used
func (m *ViProxyModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ProxyModel
	if err := m.ProxyModel.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ViProxyModel) contextValidateServer(ctx context.Context, formats strfmt.Registry) error {

	if m.Server != nil {
		if err := m.Server.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("server")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ViProxyModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ViProxyModel) UnmarshalBinary(b []byte) error {
	var res ViProxyModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}