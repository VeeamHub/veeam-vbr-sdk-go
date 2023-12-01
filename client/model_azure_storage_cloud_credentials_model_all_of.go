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

// AzureStorageCloudCredentialsModelAllOf struct for AzureStorageCloudCredentialsModelAllOf
type AzureStorageCloudCredentialsModelAllOf struct {
	// Name of the Azure storage account.
	Account string `json:"account"`
	// Tag used to identify the cloud credentials record.
	Tag *string `json:"tag,omitempty"`
}

// NewAzureStorageCloudCredentialsModelAllOf instantiates a new AzureStorageCloudCredentialsModelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureStorageCloudCredentialsModelAllOf(account string) *AzureStorageCloudCredentialsModelAllOf {
	this := AzureStorageCloudCredentialsModelAllOf{}
	this.Account = account
	return &this
}

// NewAzureStorageCloudCredentialsModelAllOfWithDefaults instantiates a new AzureStorageCloudCredentialsModelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureStorageCloudCredentialsModelAllOfWithDefaults() *AzureStorageCloudCredentialsModelAllOf {
	this := AzureStorageCloudCredentialsModelAllOf{}
	return &this
}

// GetAccount returns the Account field value
func (o *AzureStorageCloudCredentialsModelAllOf) GetAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *AzureStorageCloudCredentialsModelAllOf) GetAccountOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *AzureStorageCloudCredentialsModelAllOf) SetAccount(v string) {
	o.Account = v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *AzureStorageCloudCredentialsModelAllOf) GetTag() string {
	if o == nil || isNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureStorageCloudCredentialsModelAllOf) GetTagOk() (*string, bool) {
	if o == nil || isNil(o.Tag) {
    return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *AzureStorageCloudCredentialsModelAllOf) HasTag() bool {
	if o != nil && !isNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *AzureStorageCloudCredentialsModelAllOf) SetTag(v string) {
	o.Tag = &v
}

func (o AzureStorageCloudCredentialsModelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account"] = o.Account
	}
	if !isNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableAzureStorageCloudCredentialsModelAllOf struct {
	value *AzureStorageCloudCredentialsModelAllOf
	isSet bool
}

func (v NullableAzureStorageCloudCredentialsModelAllOf) Get() *AzureStorageCloudCredentialsModelAllOf {
	return v.value
}

func (v *NullableAzureStorageCloudCredentialsModelAllOf) Set(val *AzureStorageCloudCredentialsModelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureStorageCloudCredentialsModelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureStorageCloudCredentialsModelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureStorageCloudCredentialsModelAllOf(val *AzureStorageCloudCredentialsModelAllOf) *NullableAzureStorageCloudCredentialsModelAllOf {
	return &NullableAzureStorageCloudCredentialsModelAllOf{value: val, isSet: true}
}

func (v NullableAzureStorageCloudCredentialsModelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureStorageCloudCredentialsModelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

