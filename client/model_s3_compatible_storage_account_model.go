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

// S3CompatibleStorageAccountModel Account used to access the S3 compatible storage.
type S3CompatibleStorageAccountModel struct {
	// Endpoint address and port number of the storage.
	ServicePoint string `json:"servicePoint"`
	// ID of a region where the storage is located.
	RegionId string `json:"regionId"`
	// ID of a cloud credentials record used to access the storage.
	CredentialsId string `json:"credentialsId"`
	ConnectionSettings *ObjectStorageConnectionModel `json:"connectionSettings,omitempty"`
}

// NewS3CompatibleStorageAccountModel instantiates a new S3CompatibleStorageAccountModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3CompatibleStorageAccountModel(servicePoint string, regionId string, credentialsId string) *S3CompatibleStorageAccountModel {
	this := S3CompatibleStorageAccountModel{}
	this.ServicePoint = servicePoint
	this.RegionId = regionId
	this.CredentialsId = credentialsId
	return &this
}

// NewS3CompatibleStorageAccountModelWithDefaults instantiates a new S3CompatibleStorageAccountModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3CompatibleStorageAccountModelWithDefaults() *S3CompatibleStorageAccountModel {
	this := S3CompatibleStorageAccountModel{}
	return &this
}

// GetServicePoint returns the ServicePoint field value
func (o *S3CompatibleStorageAccountModel) GetServicePoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServicePoint
}

// GetServicePointOk returns a tuple with the ServicePoint field value
// and a boolean to check if the value has been set.
func (o *S3CompatibleStorageAccountModel) GetServicePointOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ServicePoint, true
}

// SetServicePoint sets field value
func (o *S3CompatibleStorageAccountModel) SetServicePoint(v string) {
	o.ServicePoint = v
}

// GetRegionId returns the RegionId field value
func (o *S3CompatibleStorageAccountModel) GetRegionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionId
}

// GetRegionIdOk returns a tuple with the RegionId field value
// and a boolean to check if the value has been set.
func (o *S3CompatibleStorageAccountModel) GetRegionIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RegionId, true
}

// SetRegionId sets field value
func (o *S3CompatibleStorageAccountModel) SetRegionId(v string) {
	o.RegionId = v
}

// GetCredentialsId returns the CredentialsId field value
func (o *S3CompatibleStorageAccountModel) GetCredentialsId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CredentialsId
}

// GetCredentialsIdOk returns a tuple with the CredentialsId field value
// and a boolean to check if the value has been set.
func (o *S3CompatibleStorageAccountModel) GetCredentialsIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CredentialsId, true
}

// SetCredentialsId sets field value
func (o *S3CompatibleStorageAccountModel) SetCredentialsId(v string) {
	o.CredentialsId = v
}

// GetConnectionSettings returns the ConnectionSettings field value if set, zero value otherwise.
func (o *S3CompatibleStorageAccountModel) GetConnectionSettings() ObjectStorageConnectionModel {
	if o == nil || isNil(o.ConnectionSettings) {
		var ret ObjectStorageConnectionModel
		return ret
	}
	return *o.ConnectionSettings
}

// GetConnectionSettingsOk returns a tuple with the ConnectionSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3CompatibleStorageAccountModel) GetConnectionSettingsOk() (*ObjectStorageConnectionModel, bool) {
	if o == nil || isNil(o.ConnectionSettings) {
    return nil, false
	}
	return o.ConnectionSettings, true
}

// HasConnectionSettings returns a boolean if a field has been set.
func (o *S3CompatibleStorageAccountModel) HasConnectionSettings() bool {
	if o != nil && !isNil(o.ConnectionSettings) {
		return true
	}

	return false
}

// SetConnectionSettings gets a reference to the given ObjectStorageConnectionModel and assigns it to the ConnectionSettings field.
func (o *S3CompatibleStorageAccountModel) SetConnectionSettings(v ObjectStorageConnectionModel) {
	o.ConnectionSettings = &v
}

func (o S3CompatibleStorageAccountModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["servicePoint"] = o.ServicePoint
	}
	if true {
		toSerialize["regionId"] = o.RegionId
	}
	if true {
		toSerialize["credentialsId"] = o.CredentialsId
	}
	if !isNil(o.ConnectionSettings) {
		toSerialize["connectionSettings"] = o.ConnectionSettings
	}
	return json.Marshal(toSerialize)
}

type NullableS3CompatibleStorageAccountModel struct {
	value *S3CompatibleStorageAccountModel
	isSet bool
}

func (v NullableS3CompatibleStorageAccountModel) Get() *S3CompatibleStorageAccountModel {
	return v.value
}

func (v *NullableS3CompatibleStorageAccountModel) Set(val *S3CompatibleStorageAccountModel) {
	v.value = val
	v.isSet = true
}

func (v NullableS3CompatibleStorageAccountModel) IsSet() bool {
	return v.isSet
}

func (v *NullableS3CompatibleStorageAccountModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3CompatibleStorageAccountModel(val *S3CompatibleStorageAccountModel) *NullableS3CompatibleStorageAccountModel {
	return &NullableS3CompatibleStorageAccountModel{value: val, isSet: true}
}

func (v NullableS3CompatibleStorageAccountModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3CompatibleStorageAccountModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

