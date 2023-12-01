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

// AzureBlobStorageModel Microsoft Azure Blob storage.
type AzureBlobStorageModel struct {
	RepositoryModel
	// If *true*, the maximum number of concurrent tasks is limited.
	EnableTaskLimit *bool `json:"enableTaskLimit,omitempty"`
	// Maximum number of concurrent tasks.
	MaxTaskCount *int32 `json:"maxTaskCount,omitempty"`
	Account AzureBlobStorageAccountModel `json:"account"`
	Container AzureBlobStorageContainerModel `json:"container"`
	ProxyAppliance *AzureStorageProxyModel `json:"proxyAppliance,omitempty"`
	MountServer MountServerSettingsModel `json:"mountServer"`
}

// NewAzureBlobStorageModel instantiates a new AzureBlobStorageModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureBlobStorageModel(account AzureBlobStorageAccountModel, container AzureBlobStorageContainerModel, mountServer MountServerSettingsModel, id string, name string, description string, type_ ERepositoryType) *AzureBlobStorageModel {
	this := AzureBlobStorageModel{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.Type = type_
	this.Account = account
	this.Container = container
	this.MountServer = mountServer
	return &this
}

// NewAzureBlobStorageModelWithDefaults instantiates a new AzureBlobStorageModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureBlobStorageModelWithDefaults() *AzureBlobStorageModel {
	this := AzureBlobStorageModel{}
	return &this
}

// GetEnableTaskLimit returns the EnableTaskLimit field value if set, zero value otherwise.
func (o *AzureBlobStorageModel) GetEnableTaskLimit() bool {
	if o == nil || isNil(o.EnableTaskLimit) {
		var ret bool
		return ret
	}
	return *o.EnableTaskLimit
}

// GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageModel) GetEnableTaskLimitOk() (*bool, bool) {
	if o == nil || isNil(o.EnableTaskLimit) {
    return nil, false
	}
	return o.EnableTaskLimit, true
}

// HasEnableTaskLimit returns a boolean if a field has been set.
func (o *AzureBlobStorageModel) HasEnableTaskLimit() bool {
	if o != nil && !isNil(o.EnableTaskLimit) {
		return true
	}

	return false
}

// SetEnableTaskLimit gets a reference to the given bool and assigns it to the EnableTaskLimit field.
func (o *AzureBlobStorageModel) SetEnableTaskLimit(v bool) {
	o.EnableTaskLimit = &v
}

// GetMaxTaskCount returns the MaxTaskCount field value if set, zero value otherwise.
func (o *AzureBlobStorageModel) GetMaxTaskCount() int32 {
	if o == nil || isNil(o.MaxTaskCount) {
		var ret int32
		return ret
	}
	return *o.MaxTaskCount
}

// GetMaxTaskCountOk returns a tuple with the MaxTaskCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageModel) GetMaxTaskCountOk() (*int32, bool) {
	if o == nil || isNil(o.MaxTaskCount) {
    return nil, false
	}
	return o.MaxTaskCount, true
}

// HasMaxTaskCount returns a boolean if a field has been set.
func (o *AzureBlobStorageModel) HasMaxTaskCount() bool {
	if o != nil && !isNil(o.MaxTaskCount) {
		return true
	}

	return false
}

// SetMaxTaskCount gets a reference to the given int32 and assigns it to the MaxTaskCount field.
func (o *AzureBlobStorageModel) SetMaxTaskCount(v int32) {
	o.MaxTaskCount = &v
}

// GetAccount returns the Account field value
func (o *AzureBlobStorageModel) GetAccount() AzureBlobStorageAccountModel {
	if o == nil {
		var ret AzureBlobStorageAccountModel
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageModel) GetAccountOk() (*AzureBlobStorageAccountModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *AzureBlobStorageModel) SetAccount(v AzureBlobStorageAccountModel) {
	o.Account = v
}

// GetContainer returns the Container field value
func (o *AzureBlobStorageModel) GetContainer() AzureBlobStorageContainerModel {
	if o == nil {
		var ret AzureBlobStorageContainerModel
		return ret
	}

	return o.Container
}

// GetContainerOk returns a tuple with the Container field value
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageModel) GetContainerOk() (*AzureBlobStorageContainerModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Container, true
}

// SetContainer sets field value
func (o *AzureBlobStorageModel) SetContainer(v AzureBlobStorageContainerModel) {
	o.Container = v
}

// GetProxyAppliance returns the ProxyAppliance field value if set, zero value otherwise.
func (o *AzureBlobStorageModel) GetProxyAppliance() AzureStorageProxyModel {
	if o == nil || isNil(o.ProxyAppliance) {
		var ret AzureStorageProxyModel
		return ret
	}
	return *o.ProxyAppliance
}

// GetProxyApplianceOk returns a tuple with the ProxyAppliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageModel) GetProxyApplianceOk() (*AzureStorageProxyModel, bool) {
	if o == nil || isNil(o.ProxyAppliance) {
    return nil, false
	}
	return o.ProxyAppliance, true
}

// HasProxyAppliance returns a boolean if a field has been set.
func (o *AzureBlobStorageModel) HasProxyAppliance() bool {
	if o != nil && !isNil(o.ProxyAppliance) {
		return true
	}

	return false
}

// SetProxyAppliance gets a reference to the given AzureStorageProxyModel and assigns it to the ProxyAppliance field.
func (o *AzureBlobStorageModel) SetProxyAppliance(v AzureStorageProxyModel) {
	o.ProxyAppliance = &v
}

// GetMountServer returns the MountServer field value
func (o *AzureBlobStorageModel) GetMountServer() MountServerSettingsModel {
	if o == nil {
		var ret MountServerSettingsModel
		return ret
	}

	return o.MountServer
}

// GetMountServerOk returns a tuple with the MountServer field value
// and a boolean to check if the value has been set.
func (o *AzureBlobStorageModel) GetMountServerOk() (*MountServerSettingsModel, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MountServer, true
}

// SetMountServer sets field value
func (o *AzureBlobStorageModel) SetMountServer(v MountServerSettingsModel) {
	o.MountServer = v
}

func (o AzureBlobStorageModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedRepositoryModel, errRepositoryModel := json.Marshal(o.RepositoryModel)
	if errRepositoryModel != nil {
		return []byte{}, errRepositoryModel
	}
	errRepositoryModel = json.Unmarshal([]byte(serializedRepositoryModel), &toSerialize)
	if errRepositoryModel != nil {
		return []byte{}, errRepositoryModel
	}
	if !isNil(o.EnableTaskLimit) {
		toSerialize["enableTaskLimit"] = o.EnableTaskLimit
	}
	if !isNil(o.MaxTaskCount) {
		toSerialize["maxTaskCount"] = o.MaxTaskCount
	}
	if true {
		toSerialize["account"] = o.Account
	}
	if true {
		toSerialize["container"] = o.Container
	}
	if !isNil(o.ProxyAppliance) {
		toSerialize["proxyAppliance"] = o.ProxyAppliance
	}
	if true {
		toSerialize["mountServer"] = o.MountServer
	}
	return json.Marshal(toSerialize)
}

type NullableAzureBlobStorageModel struct {
	value *AzureBlobStorageModel
	isSet bool
}

func (v NullableAzureBlobStorageModel) Get() *AzureBlobStorageModel {
	return v.value
}

func (v *NullableAzureBlobStorageModel) Set(val *AzureBlobStorageModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureBlobStorageModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureBlobStorageModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureBlobStorageModel(val *AzureBlobStorageModel) *NullableAzureBlobStorageModel {
	return &NullableAzureBlobStorageModel{value: val, isSet: true}
}

func (v NullableAzureBlobStorageModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureBlobStorageModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

