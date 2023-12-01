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

// AdvancedSmtpOptionsModel Advanced global email notification settings.
type AdvancedSmtpOptionsModel struct {
	// Port number for the SMTP server.
	Port int32 `json:"port"`
	// Connection timeout for the SMTP server.
	TimeoutMs int32 `json:"timeoutMs"`
	// If *true*, secure connection for email operations is used.
	SSLEnabled bool `json:"SSLEnabled"`
	// If *true*, the `credentialsId` credentials are used to connect to the SMTP server.
	AuthRequired bool `json:"authRequired"`
	// ID of the credentials used to connect to the server.
	CredentialsId *string `json:"credentialsId,omitempty"`
}

// NewAdvancedSmtpOptionsModel instantiates a new AdvancedSmtpOptionsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvancedSmtpOptionsModel(port int32, timeoutMs int32, sSLEnabled bool, authRequired bool) *AdvancedSmtpOptionsModel {
	this := AdvancedSmtpOptionsModel{}
	this.Port = port
	this.TimeoutMs = timeoutMs
	this.SSLEnabled = sSLEnabled
	this.AuthRequired = authRequired
	return &this
}

// NewAdvancedSmtpOptionsModelWithDefaults instantiates a new AdvancedSmtpOptionsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvancedSmtpOptionsModelWithDefaults() *AdvancedSmtpOptionsModel {
	this := AdvancedSmtpOptionsModel{}
	return &this
}

// GetPort returns the Port field value
func (o *AdvancedSmtpOptionsModel) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *AdvancedSmtpOptionsModel) GetPortOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *AdvancedSmtpOptionsModel) SetPort(v int32) {
	o.Port = v
}

// GetTimeoutMs returns the TimeoutMs field value
func (o *AdvancedSmtpOptionsModel) GetTimeoutMs() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TimeoutMs
}

// GetTimeoutMsOk returns a tuple with the TimeoutMs field value
// and a boolean to check if the value has been set.
func (o *AdvancedSmtpOptionsModel) GetTimeoutMsOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TimeoutMs, true
}

// SetTimeoutMs sets field value
func (o *AdvancedSmtpOptionsModel) SetTimeoutMs(v int32) {
	o.TimeoutMs = v
}

// GetSSLEnabled returns the SSLEnabled field value
func (o *AdvancedSmtpOptionsModel) GetSSLEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SSLEnabled
}

// GetSSLEnabledOk returns a tuple with the SSLEnabled field value
// and a boolean to check if the value has been set.
func (o *AdvancedSmtpOptionsModel) GetSSLEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SSLEnabled, true
}

// SetSSLEnabled sets field value
func (o *AdvancedSmtpOptionsModel) SetSSLEnabled(v bool) {
	o.SSLEnabled = v
}

// GetAuthRequired returns the AuthRequired field value
func (o *AdvancedSmtpOptionsModel) GetAuthRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AuthRequired
}

// GetAuthRequiredOk returns a tuple with the AuthRequired field value
// and a boolean to check if the value has been set.
func (o *AdvancedSmtpOptionsModel) GetAuthRequiredOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AuthRequired, true
}

// SetAuthRequired sets field value
func (o *AdvancedSmtpOptionsModel) SetAuthRequired(v bool) {
	o.AuthRequired = v
}

// GetCredentialsId returns the CredentialsId field value if set, zero value otherwise.
func (o *AdvancedSmtpOptionsModel) GetCredentialsId() string {
	if o == nil || isNil(o.CredentialsId) {
		var ret string
		return ret
	}
	return *o.CredentialsId
}

// GetCredentialsIdOk returns a tuple with the CredentialsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedSmtpOptionsModel) GetCredentialsIdOk() (*string, bool) {
	if o == nil || isNil(o.CredentialsId) {
    return nil, false
	}
	return o.CredentialsId, true
}

// HasCredentialsId returns a boolean if a field has been set.
func (o *AdvancedSmtpOptionsModel) HasCredentialsId() bool {
	if o != nil && !isNil(o.CredentialsId) {
		return true
	}

	return false
}

// SetCredentialsId gets a reference to the given string and assigns it to the CredentialsId field.
func (o *AdvancedSmtpOptionsModel) SetCredentialsId(v string) {
	o.CredentialsId = &v
}

func (o AdvancedSmtpOptionsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["port"] = o.Port
	}
	if true {
		toSerialize["timeoutMs"] = o.TimeoutMs
	}
	if true {
		toSerialize["SSLEnabled"] = o.SSLEnabled
	}
	if true {
		toSerialize["authRequired"] = o.AuthRequired
	}
	if !isNil(o.CredentialsId) {
		toSerialize["credentialsId"] = o.CredentialsId
	}
	return json.Marshal(toSerialize)
}

type NullableAdvancedSmtpOptionsModel struct {
	value *AdvancedSmtpOptionsModel
	isSet bool
}

func (v NullableAdvancedSmtpOptionsModel) Get() *AdvancedSmtpOptionsModel {
	return v.value
}

func (v *NullableAdvancedSmtpOptionsModel) Set(val *AdvancedSmtpOptionsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvancedSmtpOptionsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvancedSmtpOptionsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvancedSmtpOptionsModel(val *AdvancedSmtpOptionsModel) *NullableAdvancedSmtpOptionsModel {
	return &NullableAdvancedSmtpOptionsModel{value: val, isSet: true}
}

func (v NullableAdvancedSmtpOptionsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvancedSmtpOptionsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

