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

// BackupJobGuestProcessingModel Guest processing settings.
type BackupJobGuestProcessingModel struct {
	AppAwareProcessing BackupApplicationAwareProcessingModel `json:"appAwareProcessing"`
	GuestFSIndexing GuestFileSystemIndexingModel `json:"guestFSIndexing"`
	GuestInteractionProxies *GuestInteractionProxiesSettingsModel `json:"guestInteractionProxies,omitempty"`
	GuestCredentials *GuestOsCredentialsModel `json:"guestCredentials,omitempty"`
}

// NewBackupJobGuestProcessingModel instantiates a new BackupJobGuestProcessingModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupJobGuestProcessingModel(appAwareProcessing BackupApplicationAwareProcessingModel, guestFSIndexing GuestFileSystemIndexingModel) *BackupJobGuestProcessingModel {
	this := BackupJobGuestProcessingModel{}
	this.AppAwareProcessing = appAwareProcessing
	this.GuestFSIndexing = guestFSIndexing
	return &this
}

// NewBackupJobGuestProcessingModelWithDefaults instantiates a new BackupJobGuestProcessingModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupJobGuestProcessingModelWithDefaults() *BackupJobGuestProcessingModel {
	this := BackupJobGuestProcessingModel{}
	return &this
}

// GetAppAwareProcessing returns the AppAwareProcessing field value
func (o *BackupJobGuestProcessingModel) GetAppAwareProcessing() BackupApplicationAwareProcessingModel {
	if o == nil {
		var ret BackupApplicationAwareProcessingModel
		return ret
	}

	return o.AppAwareProcessing
}

// GetAppAwareProcessingOk returns a tuple with the AppAwareProcessing field value
// and a boolean to check if the value has been set.
func (o *BackupJobGuestProcessingModel) GetAppAwareProcessingOk() (*BackupApplicationAwareProcessingModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AppAwareProcessing, true
}

// SetAppAwareProcessing sets field value
func (o *BackupJobGuestProcessingModel) SetAppAwareProcessing(v BackupApplicationAwareProcessingModel) {
	o.AppAwareProcessing = v
}

// GetGuestFSIndexing returns the GuestFSIndexing field value
func (o *BackupJobGuestProcessingModel) GetGuestFSIndexing() GuestFileSystemIndexingModel {
	if o == nil {
		var ret GuestFileSystemIndexingModel
		return ret
	}

	return o.GuestFSIndexing
}

// GetGuestFSIndexingOk returns a tuple with the GuestFSIndexing field value
// and a boolean to check if the value has been set.
func (o *BackupJobGuestProcessingModel) GetGuestFSIndexingOk() (*GuestFileSystemIndexingModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.GuestFSIndexing, true
}

// SetGuestFSIndexing sets field value
func (o *BackupJobGuestProcessingModel) SetGuestFSIndexing(v GuestFileSystemIndexingModel) {
	o.GuestFSIndexing = v
}

// GetGuestInteractionProxies returns the GuestInteractionProxies field value if set, zero value otherwise.
func (o *BackupJobGuestProcessingModel) GetGuestInteractionProxies() GuestInteractionProxiesSettingsModel {
	if o == nil || isNil(o.GuestInteractionProxies) {
		var ret GuestInteractionProxiesSettingsModel
		return ret
	}
	return *o.GuestInteractionProxies
}

// GetGuestInteractionProxiesOk returns a tuple with the GuestInteractionProxies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupJobGuestProcessingModel) GetGuestInteractionProxiesOk() (*GuestInteractionProxiesSettingsModel, bool) {
	if o == nil || isNil(o.GuestInteractionProxies) {
    return nil, false
	}
	return o.GuestInteractionProxies, true
}

// HasGuestInteractionProxies returns a boolean if a field has been set.
func (o *BackupJobGuestProcessingModel) HasGuestInteractionProxies() bool {
	if o != nil && !isNil(o.GuestInteractionProxies) {
		return true
	}

	return false
}

// SetGuestInteractionProxies gets a reference to the given GuestInteractionProxiesSettingsModel and assigns it to the GuestInteractionProxies field.
func (o *BackupJobGuestProcessingModel) SetGuestInteractionProxies(v GuestInteractionProxiesSettingsModel) {
	o.GuestInteractionProxies = &v
}

// GetGuestCredentials returns the GuestCredentials field value if set, zero value otherwise.
func (o *BackupJobGuestProcessingModel) GetGuestCredentials() GuestOsCredentialsModel {
	if o == nil || isNil(o.GuestCredentials) {
		var ret GuestOsCredentialsModel
		return ret
	}
	return *o.GuestCredentials
}

// GetGuestCredentialsOk returns a tuple with the GuestCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupJobGuestProcessingModel) GetGuestCredentialsOk() (*GuestOsCredentialsModel, bool) {
	if o == nil || isNil(o.GuestCredentials) {
    return nil, false
	}
	return o.GuestCredentials, true
}

// HasGuestCredentials returns a boolean if a field has been set.
func (o *BackupJobGuestProcessingModel) HasGuestCredentials() bool {
	if o != nil && !isNil(o.GuestCredentials) {
		return true
	}

	return false
}

// SetGuestCredentials gets a reference to the given GuestOsCredentialsModel and assigns it to the GuestCredentials field.
func (o *BackupJobGuestProcessingModel) SetGuestCredentials(v GuestOsCredentialsModel) {
	o.GuestCredentials = &v
}

func (o BackupJobGuestProcessingModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["appAwareProcessing"] = o.AppAwareProcessing
	}
	if true {
		toSerialize["guestFSIndexing"] = o.GuestFSIndexing
	}
	if !isNil(o.GuestInteractionProxies) {
		toSerialize["guestInteractionProxies"] = o.GuestInteractionProxies
	}
	if !isNil(o.GuestCredentials) {
		toSerialize["guestCredentials"] = o.GuestCredentials
	}
	return json.Marshal(toSerialize)
}

type NullableBackupJobGuestProcessingModel struct {
	value *BackupJobGuestProcessingModel
	isSet bool
}

func (v NullableBackupJobGuestProcessingModel) Get() *BackupJobGuestProcessingModel {
	return v.value
}

func (v *NullableBackupJobGuestProcessingModel) Set(val *BackupJobGuestProcessingModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupJobGuestProcessingModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupJobGuestProcessingModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupJobGuestProcessingModel(val *BackupJobGuestProcessingModel) *NullableBackupJobGuestProcessingModel {
	return &NullableBackupJobGuestProcessingModel{value: val, isSet: true}
}

func (v NullableBackupJobGuestProcessingModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupJobGuestProcessingModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

