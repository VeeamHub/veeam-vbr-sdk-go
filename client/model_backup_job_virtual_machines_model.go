/*
 * Veeam Backup & Replication REST API
 *
 * This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br> Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br> Parameters, request bodies, and response bodies are defined inline or refer to schemas defined globally. Some schemas are polymorphic. 
 *
 * API version: 1.0-rev1
 * Contact: support@veeam.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// BackupJobVirtualMachinesModel struct for BackupJobVirtualMachinesModel
type BackupJobVirtualMachinesModel struct {
	// Array of VM and VM containers processed by the job.
	Includes []VmwareObjectSizeModel `json:"includes"`
	Excludes *BackupJobExclusions `json:"excludes,omitempty"`
}

// NewBackupJobVirtualMachinesModel instantiates a new BackupJobVirtualMachinesModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupJobVirtualMachinesModel(includes []VmwareObjectSizeModel) *BackupJobVirtualMachinesModel {
	this := BackupJobVirtualMachinesModel{}
	this.Includes = includes
	return &this
}

// NewBackupJobVirtualMachinesModelWithDefaults instantiates a new BackupJobVirtualMachinesModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupJobVirtualMachinesModelWithDefaults() *BackupJobVirtualMachinesModel {
	this := BackupJobVirtualMachinesModel{}
	return &this
}

// GetIncludes returns the Includes field value
func (o *BackupJobVirtualMachinesModel) GetIncludes() []VmwareObjectSizeModel {
	if o == nil {
		var ret []VmwareObjectSizeModel
		return ret
	}

	return o.Includes
}

// GetIncludesOk returns a tuple with the Includes field value
// and a boolean to check if the value has been set.
func (o *BackupJobVirtualMachinesModel) GetIncludesOk() (*[]VmwareObjectSizeModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Includes, true
}

// SetIncludes sets field value
func (o *BackupJobVirtualMachinesModel) SetIncludes(v []VmwareObjectSizeModel) {
	o.Includes = v
}

// GetExcludes returns the Excludes field value if set, zero value otherwise.
func (o *BackupJobVirtualMachinesModel) GetExcludes() BackupJobExclusions {
	if o == nil || o.Excludes == nil {
		var ret BackupJobExclusions
		return ret
	}
	return *o.Excludes
}

// GetExcludesOk returns a tuple with the Excludes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupJobVirtualMachinesModel) GetExcludesOk() (*BackupJobExclusions, bool) {
	if o == nil || o.Excludes == nil {
		return nil, false
	}
	return o.Excludes, true
}

// HasExcludes returns a boolean if a field has been set.
func (o *BackupJobVirtualMachinesModel) HasExcludes() bool {
	if o != nil && o.Excludes != nil {
		return true
	}

	return false
}

// SetExcludes gets a reference to the given BackupJobExclusions and assigns it to the Excludes field.
func (o *BackupJobVirtualMachinesModel) SetExcludes(v BackupJobExclusions) {
	o.Excludes = &v
}

func (o BackupJobVirtualMachinesModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["includes"] = o.Includes
	}
	if o.Excludes != nil {
		toSerialize["excludes"] = o.Excludes
	}
	return json.Marshal(toSerialize)
}

type NullableBackupJobVirtualMachinesModel struct {
	value *BackupJobVirtualMachinesModel
	isSet bool
}

func (v NullableBackupJobVirtualMachinesModel) Get() *BackupJobVirtualMachinesModel {
	return v.value
}

func (v *NullableBackupJobVirtualMachinesModel) Set(val *BackupJobVirtualMachinesModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupJobVirtualMachinesModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupJobVirtualMachinesModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupJobVirtualMachinesModel(val *BackupJobVirtualMachinesModel) *NullableBackupJobVirtualMachinesModel {
	return &NullableBackupJobVirtualMachinesModel{value: val, isSet: true}
}

func (v NullableBackupJobVirtualMachinesModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupJobVirtualMachinesModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

