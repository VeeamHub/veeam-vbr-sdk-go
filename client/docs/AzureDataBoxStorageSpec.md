# AzureDataBoxStorageSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableTaskLimit** | Pointer to **bool** | If *true*, the maximum number of concurrent tasks is limited. | [optional] 
**MaxTaskCount** | Pointer to **int32** | Maximum number of concurrent tasks. | [optional] 
**Account** | [**AzureDataBoxStorageAccountModel**](AzureDataBoxStorageAccountModel.md) |  | 
**Container** | [**AzureDataBoxStorageContainerModel**](AzureDataBoxStorageContainerModel.md) |  | 

## Methods

### NewAzureDataBoxStorageSpec

`func NewAzureDataBoxStorageSpec(account AzureDataBoxStorageAccountModel, container AzureDataBoxStorageContainerModel, ) *AzureDataBoxStorageSpec`

NewAzureDataBoxStorageSpec instantiates a new AzureDataBoxStorageSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureDataBoxStorageSpecWithDefaults

`func NewAzureDataBoxStorageSpecWithDefaults() *AzureDataBoxStorageSpec`

NewAzureDataBoxStorageSpecWithDefaults instantiates a new AzureDataBoxStorageSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableTaskLimit

`func (o *AzureDataBoxStorageSpec) GetEnableTaskLimit() bool`

GetEnableTaskLimit returns the EnableTaskLimit field if non-nil, zero value otherwise.

### GetEnableTaskLimitOk

`func (o *AzureDataBoxStorageSpec) GetEnableTaskLimitOk() (*bool, bool)`

GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTaskLimit

`func (o *AzureDataBoxStorageSpec) SetEnableTaskLimit(v bool)`

SetEnableTaskLimit sets EnableTaskLimit field to given value.

### HasEnableTaskLimit

`func (o *AzureDataBoxStorageSpec) HasEnableTaskLimit() bool`

HasEnableTaskLimit returns a boolean if a field has been set.

### GetMaxTaskCount

`func (o *AzureDataBoxStorageSpec) GetMaxTaskCount() int32`

GetMaxTaskCount returns the MaxTaskCount field if non-nil, zero value otherwise.

### GetMaxTaskCountOk

`func (o *AzureDataBoxStorageSpec) GetMaxTaskCountOk() (*int32, bool)`

GetMaxTaskCountOk returns a tuple with the MaxTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskCount

`func (o *AzureDataBoxStorageSpec) SetMaxTaskCount(v int32)`

SetMaxTaskCount sets MaxTaskCount field to given value.

### HasMaxTaskCount

`func (o *AzureDataBoxStorageSpec) HasMaxTaskCount() bool`

HasMaxTaskCount returns a boolean if a field has been set.

### GetAccount

`func (o *AzureDataBoxStorageSpec) GetAccount() AzureDataBoxStorageAccountModel`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AzureDataBoxStorageSpec) GetAccountOk() (*AzureDataBoxStorageAccountModel, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AzureDataBoxStorageSpec) SetAccount(v AzureDataBoxStorageAccountModel)`

SetAccount sets Account field to given value.


### GetContainer

`func (o *AzureDataBoxStorageSpec) GetContainer() AzureDataBoxStorageContainerModel`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *AzureDataBoxStorageSpec) GetContainerOk() (*AzureDataBoxStorageContainerModel, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *AzureDataBoxStorageSpec) SetContainer(v AzureDataBoxStorageContainerModel)`

SetContainer sets Container field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


