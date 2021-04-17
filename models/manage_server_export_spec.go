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

// ManageServerExportSpec manage server export spec
//
// swagger:model ManageServerExportSpec
type ManageServerExportSpec struct {

	// Array of server IDs.
	// Unique: true
	Ids []strfmt.UUID `json:"ids"`

	// Array of server names. Wildcard characters are supported.
	// Unique: true
	Names []string `json:"names"`

	// Array of server types.
	// Unique: true
	Types []EManagedServerType `json:"types"`
}

// Validate validates this manage server export spec
func (m *ManageServerExportSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManageServerExportSpec) validateIds(formats strfmt.Registry) error {
	if swag.IsZero(m.Ids) { // not required
		return nil
	}

	if err := validate.UniqueItems("ids", "body", m.Ids); err != nil {
		return err
	}

	for i := 0; i < len(m.Ids); i++ {

		if err := validate.FormatOf("ids"+"."+strconv.Itoa(i), "body", "uuid", m.Ids[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *ManageServerExportSpec) validateNames(formats strfmt.Registry) error {
	if swag.IsZero(m.Names) { // not required
		return nil
	}

	if err := validate.UniqueItems("names", "body", m.Names); err != nil {
		return err
	}

	return nil
}

func (m *ManageServerExportSpec) validateTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.Types) { // not required
		return nil
	}

	if err := validate.UniqueItems("types", "body", m.Types); err != nil {
		return err
	}

	for i := 0; i < len(m.Types); i++ {

		if err := m.Types[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("types" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this manage server export spec based on the context it is used
func (m *ManageServerExportSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManageServerExportSpec) contextValidateTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Types); i++ {

		if err := m.Types[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("types" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ManageServerExportSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManageServerExportSpec) UnmarshalBinary(b []byte) error {
	var res ManageServerExportSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}