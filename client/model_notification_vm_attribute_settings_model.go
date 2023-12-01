/*
Veeam Backup & Replication REST API

This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br>Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Some request and response bodies refer to reusable schema objects that are defined in the `schemas` section of the REST API specification. Schemas, in turn, may inherit a part of their properties by referring to other schemas, and be polymorphic by referring to multiple alternate schemas.

API version: 1.1-rev0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NotificationVmAttributeSettingsModel VM attribute settings.
type NotificationVmAttributeSettingsModel struct {
	// If *true*, information about successfully performed backup is written to a VM attribute.
	IsEnabled bool `json:"isEnabled"`
	// Name of the VM attribute.
	Notes *string `json:"notes,omitempty"`
	// If *true*, information about successfully performed backup is appended to the existing value of the attribute added by the user.
	AppendToExisitingValue *bool `json:"appendToExisitingValue,omitempty"`
}

// NewNotificationVmAttributeSettingsModel instantiates a new NotificationVmAttributeSettingsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationVmAttributeSettingsModel(isEnabled bool) *NotificationVmAttributeSettingsModel {
	this := NotificationVmAttributeSettingsModel{}
	this.IsEnabled = isEnabled
	return &this
}

// NewNotificationVmAttributeSettingsModelWithDefaults instantiates a new NotificationVmAttributeSettingsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationVmAttributeSettingsModelWithDefaults() *NotificationVmAttributeSettingsModel {
	this := NotificationVmAttributeSettingsModel{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value
func (o *NotificationVmAttributeSettingsModel) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *NotificationVmAttributeSettingsModel) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *NotificationVmAttributeSettingsModel) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *NotificationVmAttributeSettingsModel) GetNotes() string {
	if o == nil || isNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationVmAttributeSettingsModel) GetNotesOk() (*string, bool) {
	if o == nil || isNil(o.Notes) {
    return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *NotificationVmAttributeSettingsModel) HasNotes() bool {
	if o != nil && !isNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *NotificationVmAttributeSettingsModel) SetNotes(v string) {
	o.Notes = &v
}

// GetAppendToExisitingValue returns the AppendToExisitingValue field value if set, zero value otherwise.
func (o *NotificationVmAttributeSettingsModel) GetAppendToExisitingValue() bool {
	if o == nil || isNil(o.AppendToExisitingValue) {
		var ret bool
		return ret
	}
	return *o.AppendToExisitingValue
}

// GetAppendToExisitingValueOk returns a tuple with the AppendToExisitingValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationVmAttributeSettingsModel) GetAppendToExisitingValueOk() (*bool, bool) {
	if o == nil || isNil(o.AppendToExisitingValue) {
    return nil, false
	}
	return o.AppendToExisitingValue, true
}

// HasAppendToExisitingValue returns a boolean if a field has been set.
func (o *NotificationVmAttributeSettingsModel) HasAppendToExisitingValue() bool {
	if o != nil && !isNil(o.AppendToExisitingValue) {
		return true
	}

	return false
}

// SetAppendToExisitingValue gets a reference to the given bool and assigns it to the AppendToExisitingValue field.
func (o *NotificationVmAttributeSettingsModel) SetAppendToExisitingValue(v bool) {
	o.AppendToExisitingValue = &v
}

func (o NotificationVmAttributeSettingsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if !isNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !isNil(o.AppendToExisitingValue) {
		toSerialize["appendToExisitingValue"] = o.AppendToExisitingValue
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationVmAttributeSettingsModel struct {
	value *NotificationVmAttributeSettingsModel
	isSet bool
}

func (v NullableNotificationVmAttributeSettingsModel) Get() *NotificationVmAttributeSettingsModel {
	return v.value
}

func (v *NullableNotificationVmAttributeSettingsModel) Set(val *NotificationVmAttributeSettingsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationVmAttributeSettingsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationVmAttributeSettingsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationVmAttributeSettingsModel(val *NotificationVmAttributeSettingsModel) *NullableNotificationVmAttributeSettingsModel {
	return &NullableNotificationVmAttributeSettingsModel{value: val, isSet: true}
}

func (v NullableNotificationVmAttributeSettingsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationVmAttributeSettingsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

