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

// BackupJobModel struct for BackupJobModel
type BackupJobModel struct {
	JobModel
	// If *true*, the job has a high priority in getting backup infrastructure resources.
	IsHighPriority bool `json:"isHighPriority"`
	VirtualMachines BackupJobVirtualMachinesModel `json:"virtualMachines"`
	Storage BackupJobStorageModel `json:"storage"`
	GuestProcessing BackupJobGuestProcessingModel `json:"guestProcessing"`
	Schedule BackupScheduleModel `json:"schedule"`
}

// NewBackupJobModel instantiates a new BackupJobModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupJobModel(isHighPriority bool, virtualMachines BackupJobVirtualMachinesModel, storage BackupJobStorageModel, guestProcessing BackupJobGuestProcessingModel, schedule BackupScheduleModel, id string, name string, description string, type_ EJobType, isDisabled bool) *BackupJobModel {
	this := BackupJobModel{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.Type = type_
	this.IsDisabled = isDisabled
	this.IsHighPriority = isHighPriority
	this.VirtualMachines = virtualMachines
	this.Storage = storage
	this.GuestProcessing = guestProcessing
	this.Schedule = schedule
	return &this
}

// NewBackupJobModelWithDefaults instantiates a new BackupJobModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupJobModelWithDefaults() *BackupJobModel {
	this := BackupJobModel{}
	return &this
}

// GetIsHighPriority returns the IsHighPriority field value
func (o *BackupJobModel) GetIsHighPriority() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsHighPriority
}

// GetIsHighPriorityOk returns a tuple with the IsHighPriority field value
// and a boolean to check if the value has been set.
func (o *BackupJobModel) GetIsHighPriorityOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IsHighPriority, true
}

// SetIsHighPriority sets field value
func (o *BackupJobModel) SetIsHighPriority(v bool) {
	o.IsHighPriority = v
}

// GetVirtualMachines returns the VirtualMachines field value
func (o *BackupJobModel) GetVirtualMachines() BackupJobVirtualMachinesModel {
	if o == nil {
		var ret BackupJobVirtualMachinesModel
		return ret
	}

	return o.VirtualMachines
}

// GetVirtualMachinesOk returns a tuple with the VirtualMachines field value
// and a boolean to check if the value has been set.
func (o *BackupJobModel) GetVirtualMachinesOk() (*BackupJobVirtualMachinesModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.VirtualMachines, true
}

// SetVirtualMachines sets field value
func (o *BackupJobModel) SetVirtualMachines(v BackupJobVirtualMachinesModel) {
	o.VirtualMachines = v
}

// GetStorage returns the Storage field value
func (o *BackupJobModel) GetStorage() BackupJobStorageModel {
	if o == nil {
		var ret BackupJobStorageModel
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *BackupJobModel) GetStorageOk() (*BackupJobStorageModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *BackupJobModel) SetStorage(v BackupJobStorageModel) {
	o.Storage = v
}

// GetGuestProcessing returns the GuestProcessing field value
func (o *BackupJobModel) GetGuestProcessing() BackupJobGuestProcessingModel {
	if o == nil {
		var ret BackupJobGuestProcessingModel
		return ret
	}

	return o.GuestProcessing
}

// GetGuestProcessingOk returns a tuple with the GuestProcessing field value
// and a boolean to check if the value has been set.
func (o *BackupJobModel) GetGuestProcessingOk() (*BackupJobGuestProcessingModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.GuestProcessing, true
}

// SetGuestProcessing sets field value
func (o *BackupJobModel) SetGuestProcessing(v BackupJobGuestProcessingModel) {
	o.GuestProcessing = v
}

// GetSchedule returns the Schedule field value
func (o *BackupJobModel) GetSchedule() BackupScheduleModel {
	if o == nil {
		var ret BackupScheduleModel
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *BackupJobModel) GetScheduleOk() (*BackupScheduleModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *BackupJobModel) SetSchedule(v BackupScheduleModel) {
	o.Schedule = v
}

func (o BackupJobModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedJobModel, errJobModel := json.Marshal(o.JobModel)
	if errJobModel != nil {
		return []byte{}, errJobModel
	}
	errJobModel = json.Unmarshal([]byte(serializedJobModel), &toSerialize)
	if errJobModel != nil {
		return []byte{}, errJobModel
	}
	if true {
		toSerialize["isHighPriority"] = o.IsHighPriority
	}
	if true {
		toSerialize["virtualMachines"] = o.VirtualMachines
	}
	if true {
		toSerialize["storage"] = o.Storage
	}
	if true {
		toSerialize["guestProcessing"] = o.GuestProcessing
	}
	if true {
		toSerialize["schedule"] = o.Schedule
	}
	return json.Marshal(toSerialize)
}

type NullableBackupJobModel struct {
	value *BackupJobModel
	isSet bool
}

func (v NullableBackupJobModel) Get() *BackupJobModel {
	return v.value
}

func (v *NullableBackupJobModel) Set(val *BackupJobModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupJobModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupJobModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupJobModel(val *BackupJobModel) *NullableBackupJobModel {
	return &NullableBackupJobModel{value: val, isSet: true}
}

func (v NullableBackupJobModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupJobModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

