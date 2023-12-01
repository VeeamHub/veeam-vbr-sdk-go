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

// CloudBrowserSpec struct for CloudBrowserSpec
type CloudBrowserSpec struct {
	// ID of the object storage account (for browsing either storage or compute infrastructure).
	CredentialsId string `json:"credentialsId"`
	ServiceType ECloudServiceType `json:"serviceType"`
}

// NewCloudBrowserSpec instantiates a new CloudBrowserSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudBrowserSpec(credentialsId string, serviceType ECloudServiceType) *CloudBrowserSpec {
	this := CloudBrowserSpec{}
	this.CredentialsId = credentialsId
	this.ServiceType = serviceType
	return &this
}

// NewCloudBrowserSpecWithDefaults instantiates a new CloudBrowserSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudBrowserSpecWithDefaults() *CloudBrowserSpec {
	this := CloudBrowserSpec{}
	return &this
}

// GetCredentialsId returns the CredentialsId field value
func (o *CloudBrowserSpec) GetCredentialsId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CredentialsId
}

// GetCredentialsIdOk returns a tuple with the CredentialsId field value
// and a boolean to check if the value has been set.
func (o *CloudBrowserSpec) GetCredentialsIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CredentialsId, true
}

// SetCredentialsId sets field value
func (o *CloudBrowserSpec) SetCredentialsId(v string) {
	o.CredentialsId = v
}

// GetServiceType returns the ServiceType field value
func (o *CloudBrowserSpec) GetServiceType() ECloudServiceType {
	if o == nil {
		var ret ECloudServiceType
		return ret
	}

	return o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value
// and a boolean to check if the value has been set.
func (o *CloudBrowserSpec) GetServiceTypeOk() (*ECloudServiceType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ServiceType, true
}

// SetServiceType sets field value
func (o *CloudBrowserSpec) SetServiceType(v ECloudServiceType) {
	o.ServiceType = v
}

func (o CloudBrowserSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["credentialsId"] = o.CredentialsId
	}
	if true {
		toSerialize["serviceType"] = o.ServiceType
	}
	return json.Marshal(toSerialize)
}

type NullableCloudBrowserSpec struct {
	value *CloudBrowserSpec
	isSet bool
}

func (v NullableCloudBrowserSpec) Get() *CloudBrowserSpec {
	return v.value
}

func (v *NullableCloudBrowserSpec) Set(val *CloudBrowserSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudBrowserSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudBrowserSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudBrowserSpec(val *CloudBrowserSpec) *NullableCloudBrowserSpec {
	return &NullableCloudBrowserSpec{value: val, isSet: true}
}

func (v NullableCloudBrowserSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudBrowserSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

