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

// ScheduleAfterThisJobModel Job chaining options.
//
// swagger:model ScheduleAfterThisJobModel
type ScheduleAfterThisJobModel struct {

	// If *true*, job chaining is enabled.
	// Required: true
	IsEnabled *bool `json:"isEnabled"`

	// Name of the preceding job.
	JobName string `json:"jobName,omitempty"`
}

// Validate validates this schedule after this job model
func (m *ScheduleAfterThisJobModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsEnabled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScheduleAfterThisJobModel) validateIsEnabled(formats strfmt.Registry) error {

	if err := validate.Required("isEnabled", "body", m.IsEnabled); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this schedule after this job model based on context it is used
func (m *ScheduleAfterThisJobModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ScheduleAfterThisJobModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScheduleAfterThisJobModel) UnmarshalBinary(b []byte) error {
	var res ScheduleAfterThisJobModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}