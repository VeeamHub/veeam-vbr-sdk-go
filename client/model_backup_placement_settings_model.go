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

// BackupPlacementSettingsModel struct for BackupPlacementSettingsModel
type BackupPlacementSettingsModel struct {
	// ID of a performance extent.
	ExtentId string `json:"extentId"`
	AllowedBackups EAllowedBackupsType `json:"allowedBackups"`
}

// NewBackupPlacementSettingsModel instantiates a new BackupPlacementSettingsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupPlacementSettingsModel(extentId string, allowedBackups EAllowedBackupsType) *BackupPlacementSettingsModel {
	this := BackupPlacementSettingsModel{}
	this.ExtentId = extentId
	this.AllowedBackups = allowedBackups
	return &this
}

// NewBackupPlacementSettingsModelWithDefaults instantiates a new BackupPlacementSettingsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupPlacementSettingsModelWithDefaults() *BackupPlacementSettingsModel {
	this := BackupPlacementSettingsModel{}
	return &this
}

// GetExtentId returns the ExtentId field value
func (o *BackupPlacementSettingsModel) GetExtentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtentId
}

// GetExtentIdOk returns a tuple with the ExtentId field value
// and a boolean to check if the value has been set.
func (o *BackupPlacementSettingsModel) GetExtentIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExtentId, true
}

// SetExtentId sets field value
func (o *BackupPlacementSettingsModel) SetExtentId(v string) {
	o.ExtentId = v
}

// GetAllowedBackups returns the AllowedBackups field value
func (o *BackupPlacementSettingsModel) GetAllowedBackups() EAllowedBackupsType {
	if o == nil {
		var ret EAllowedBackupsType
		return ret
	}

	return o.AllowedBackups
}

// GetAllowedBackupsOk returns a tuple with the AllowedBackups field value
// and a boolean to check if the value has been set.
func (o *BackupPlacementSettingsModel) GetAllowedBackupsOk() (*EAllowedBackupsType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AllowedBackups, true
}

// SetAllowedBackups sets field value
func (o *BackupPlacementSettingsModel) SetAllowedBackups(v EAllowedBackupsType) {
	o.AllowedBackups = v
}

func (o BackupPlacementSettingsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["extentId"] = o.ExtentId
	}
	if true {
		toSerialize["allowedBackups"] = o.AllowedBackups
	}
	return json.Marshal(toSerialize)
}

type NullableBackupPlacementSettingsModel struct {
	value *BackupPlacementSettingsModel
	isSet bool
}

func (v NullableBackupPlacementSettingsModel) Get() *BackupPlacementSettingsModel {
	return v.value
}

func (v *NullableBackupPlacementSettingsModel) Set(val *BackupPlacementSettingsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupPlacementSettingsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupPlacementSettingsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupPlacementSettingsModel(val *BackupPlacementSettingsModel) *NullableBackupPlacementSettingsModel {
	return &NullableBackupPlacementSettingsModel{value: val, isSet: true}
}

func (v NullableBackupPlacementSettingsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupPlacementSettingsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

