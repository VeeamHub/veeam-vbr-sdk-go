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

// ESennightOfMonth Sennight of the month.
//
// swagger:model ESennightOfMonth
type ESennightOfMonth string

func NewESennightOfMonth(value ESennightOfMonth) *ESennightOfMonth {
	v := value
	return &v
}

const (

	// ESennightOfMonthFirst captures enum value "First"
	ESennightOfMonthFirst ESennightOfMonth = "First"

	// ESennightOfMonthSecond captures enum value "Second"
	ESennightOfMonthSecond ESennightOfMonth = "Second"

	// ESennightOfMonthThird captures enum value "Third"
	ESennightOfMonthThird ESennightOfMonth = "Third"

	// ESennightOfMonthFourth captures enum value "Fourth"
	ESennightOfMonthFourth ESennightOfMonth = "Fourth"

	// ESennightOfMonthFifth captures enum value "Fifth"
	ESennightOfMonthFifth ESennightOfMonth = "Fifth"

	// ESennightOfMonthLast captures enum value "Last"
	ESennightOfMonthLast ESennightOfMonth = "Last"
)

// for schema
var eSennightOfMonthEnum []interface{}

func init() {
	var res []ESennightOfMonth
	if err := json.Unmarshal([]byte(`["First","Second","Third","Fourth","Fifth","Last"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eSennightOfMonthEnum = append(eSennightOfMonthEnum, v)
	}
}

func (m ESennightOfMonth) validateESennightOfMonthEnum(path, location string, value ESennightOfMonth) error {
	if err := validate.EnumCase(path, location, value, eSennightOfMonthEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this e sennight of month
func (m ESennightOfMonth) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateESennightOfMonthEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this e sennight of month based on context it is used
func (m ESennightOfMonth) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}