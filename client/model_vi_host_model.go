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

// ViHostModel struct for ViHostModel
type ViHostModel struct {
	ManagedServerModel
	ViHostType *EViHostType `json:"viHostType,omitempty"`
	// Port used to communicate with the server.
	Port int32 `json:"port"`
}

// NewViHostModel instantiates a new ViHostModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViHostModel(port int32, id string, name string, description string, type_ EManagedServerType, credentialsId string) *ViHostModel {
	this := ViHostModel{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.Type = type_
	this.CredentialsId = credentialsId
	this.Port = port
	return &this
}

// NewViHostModelWithDefaults instantiates a new ViHostModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViHostModelWithDefaults() *ViHostModel {
	this := ViHostModel{}
	return &this
}

// GetViHostType returns the ViHostType field value if set, zero value otherwise.
func (o *ViHostModel) GetViHostType() EViHostType {
	if o == nil || isNil(o.ViHostType) {
		var ret EViHostType
		return ret
	}
	return *o.ViHostType
}

// GetViHostTypeOk returns a tuple with the ViHostType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViHostModel) GetViHostTypeOk() (*EViHostType, bool) {
	if o == nil || isNil(o.ViHostType) {
    return nil, false
	}
	return o.ViHostType, true
}

// HasViHostType returns a boolean if a field has been set.
func (o *ViHostModel) HasViHostType() bool {
	if o != nil && !isNil(o.ViHostType) {
		return true
	}

	return false
}

// SetViHostType gets a reference to the given EViHostType and assigns it to the ViHostType field.
func (o *ViHostModel) SetViHostType(v EViHostType) {
	o.ViHostType = &v
}

// GetPort returns the Port field value
func (o *ViHostModel) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *ViHostModel) GetPortOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *ViHostModel) SetPort(v int32) {
	o.Port = v
}

func (o ViHostModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedManagedServerModel, errManagedServerModel := json.Marshal(o.ManagedServerModel)
	if errManagedServerModel != nil {
		return []byte{}, errManagedServerModel
	}
	errManagedServerModel = json.Unmarshal([]byte(serializedManagedServerModel), &toSerialize)
	if errManagedServerModel != nil {
		return []byte{}, errManagedServerModel
	}
	if !isNil(o.ViHostType) {
		toSerialize["viHostType"] = o.ViHostType
	}
	if true {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableViHostModel struct {
	value *ViHostModel
	isSet bool
}

func (v NullableViHostModel) Get() *ViHostModel {
	return v.value
}

func (v *NullableViHostModel) Set(val *ViHostModel) {
	v.value = val
	v.isSet = true
}

func (v NullableViHostModel) IsSet() bool {
	return v.isSet
}

func (v *NullableViHostModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViHostModel(val *ViHostModel) *NullableViHostModel {
	return &NullableViHostModel{value: val, isSet: true}
}

func (v NullableViHostModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViHostModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

