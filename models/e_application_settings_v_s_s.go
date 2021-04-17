// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// EApplicationSettingsVSS Behavior scenario for application-aware processing.
//
// swagger:model EApplicationSettingsVSS
type EApplicationSettingsVSS string

func NewEApplicationSettingsVSS(value EApplicationSettingsVSS) *EApplicationSettingsVSS {
	v := value
	return &v
}

const (

	// EApplicationSettingsVSSRequireSuccess captures enum value "requireSuccess"
	EApplicationSettingsVSSRequireSuccess EApplicationSettingsVSS = "requireSuccess"

	// EApplicationSettingsVSSIgnoreFailures captures enum value "ignoreFailures"
	EApplicationSettingsVSSIgnoreFailures EApplicationSettingsVSS = "ignoreFailures"

	// EApplicationSettingsVSSDisabled captures enum value "disabled"
	EApplicationSettingsVSSDisabled EApplicationSettingsVSS = "disabled"
)

// for schema
var eApplicationSettingsVSSEnum []interface{}

func init() {
	var res []EApplicationSettingsVSS
	if err := json.Unmarshal([]byte(`["requireSuccess","ignoreFailures","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eApplicationSettingsVSSEnum = append(eApplicationSettingsVSSEnum, v)
	}
}

func (m EApplicationSettingsVSS) validateEApplicationSettingsVSSEnum(path, location string, value EApplicationSettingsVSS) error {
	if err := validate.EnumCase(path, location, value, eApplicationSettingsVSSEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this e application settings v s s
func (m EApplicationSettingsVSS) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEApplicationSettingsVSSEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this e application settings v s s based on context it is used
func (m EApplicationSettingsVSS) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}