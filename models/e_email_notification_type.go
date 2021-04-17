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

// EEmailNotificationType Type of email notification settings: global notification settings specified for the backup server, or custom notification settings specified for this job.
//
//
// swagger:model EEmailNotificationType
type EEmailNotificationType string

func NewEEmailNotificationType(value EEmailNotificationType) *EEmailNotificationType {
	v := value
	return &v
}

const (

	// EEmailNotificationTypeUseGlobalNotificationSettings captures enum value "UseGlobalNotificationSettings"
	EEmailNotificationTypeUseGlobalNotificationSettings EEmailNotificationType = "UseGlobalNotificationSettings"

	// EEmailNotificationTypeUseCustomNotificationSettings captures enum value "UseCustomNotificationSettings"
	EEmailNotificationTypeUseCustomNotificationSettings EEmailNotificationType = "UseCustomNotificationSettings"
)

// for schema
var eEmailNotificationTypeEnum []interface{}

func init() {
	var res []EEmailNotificationType
	if err := json.Unmarshal([]byte(`["UseGlobalNotificationSettings","UseCustomNotificationSettings"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eEmailNotificationTypeEnum = append(eEmailNotificationTypeEnum, v)
	}
}

func (m EEmailNotificationType) validateEEmailNotificationTypeEnum(path, location string, value EEmailNotificationType) error {
	if err := validate.EnumCase(path, location, value, eEmailNotificationTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this e email notification type
func (m EEmailNotificationType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEEmailNotificationTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this e email notification type based on context it is used
func (m EEmailNotificationType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}