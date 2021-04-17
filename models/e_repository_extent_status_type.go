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

// ERepositoryExtentStatusType e repository extent status type
//
// swagger:model ERepositoryExtentStatusType
type ERepositoryExtentStatusType string

func NewERepositoryExtentStatusType(value ERepositoryExtentStatusType) *ERepositoryExtentStatusType {
	v := value
	return &v
}

const (

	// ERepositoryExtentStatusTypeNormal captures enum value "Normal"
	ERepositoryExtentStatusTypeNormal ERepositoryExtentStatusType = "Normal"

	// ERepositoryExtentStatusTypeEvacuate captures enum value "Evacuate"
	ERepositoryExtentStatusTypeEvacuate ERepositoryExtentStatusType = "Evacuate"

	// ERepositoryExtentStatusTypePending captures enum value "Pending"
	ERepositoryExtentStatusTypePending ERepositoryExtentStatusType = "Pending"

	// ERepositoryExtentStatusTypeSealed captures enum value "Sealed"
	ERepositoryExtentStatusTypeSealed ERepositoryExtentStatusType = "Sealed"

	// ERepositoryExtentStatusTypeMaintenance captures enum value "Maintenance"
	ERepositoryExtentStatusTypeMaintenance ERepositoryExtentStatusType = "Maintenance"
)

// for schema
var eRepositoryExtentStatusTypeEnum []interface{}

func init() {
	var res []ERepositoryExtentStatusType
	if err := json.Unmarshal([]byte(`["Normal","Evacuate","Pending","Sealed","Maintenance"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eRepositoryExtentStatusTypeEnum = append(eRepositoryExtentStatusTypeEnum, v)
	}
}

func (m ERepositoryExtentStatusType) validateERepositoryExtentStatusTypeEnum(path, location string, value ERepositoryExtentStatusType) error {
	if err := validate.EnumCase(path, location, value, eRepositoryExtentStatusTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this e repository extent status type
func (m ERepositoryExtentStatusType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateERepositoryExtentStatusTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this e repository extent status type based on context it is used
func (m ERepositoryExtentStatusType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}