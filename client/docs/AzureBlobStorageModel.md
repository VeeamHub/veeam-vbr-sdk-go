# AzureBlobStorageModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableTaskLimit** | Pointer to **bool** | If *true*, the maximum number of concurrent tasks is limited. | [optional] 
**MaxTaskCount** | Pointer to **int32** | Maximum number of concurrent tasks. | [optional] 
**Account** | [**AzureBlobStorageAccountModel**](AzureBlobStorageAccountModel.md) |  | 
**Container** | [**AzureBlobStorageContainerModel**](AzureBlobStorageContainerModel.md) |  | 
**ProxyAppliance** | Pointer to [**AzureStorageProxyModel**](AzureStorageProxyModel.md) |  | [optional] 
**MountServer** | [**MountServerSettingsModel**](MountServerSettingsModel.md) |  | 

## Methods

### NewAzureBlobStorageModel

`func NewAzureBlobStorageModel(account AzureBlobStorageAccountModel, container AzureBlobStorageContainerModel, mountServer MountServerSettingsModel, ) *AzureBlobStorageModel`

NewAzureBlobStorageModel instantiates a new AzureBlobStorageModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureBlobStorageModelWithDefaults

`func NewAzureBlobStorageModelWithDefaults() *AzureBlobStorageModel`

NewAzureBlobStorageModelWithDefaults instantiates a new AzureBlobStorageModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableTaskLimit

`func (o *AzureBlobStorageModel) GetEnableTaskLimit() bool`

GetEnableTaskLimit returns the EnableTaskLimit field if non-nil, zero value otherwise.

### GetEnableTaskLimitOk

`func (o *AzureBlobStorageModel) GetEnableTaskLimitOk() (*bool, bool)`

GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTaskLimit

`func (o *AzureBlobStorageModel) SetEnableTaskLimit(v bool)`

SetEnableTaskLimit sets EnableTaskLimit field to given value.

### HasEnableTaskLimit

`func (o *AzureBlobStorageModel) HasEnableTaskLimit() bool`

HasEnableTaskLimit returns a boolean if a field has been set.

### GetMaxTaskCount

`func (o *AzureBlobStorageModel) GetMaxTaskCount() int32`

GetMaxTaskCount returns the MaxTaskCount field if non-nil, zero value otherwise.

### GetMaxTaskCountOk

`func (o *AzureBlobStorageModel) GetMaxTaskCountOk() (*int32, bool)`

GetMaxTaskCountOk returns a tuple with the MaxTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskCount

`func (o *AzureBlobStorageModel) SetMaxTaskCount(v int32)`

SetMaxTaskCount sets MaxTaskCount field to given value.

### HasMaxTaskCount

`func (o *AzureBlobStorageModel) HasMaxTaskCount() bool`

HasMaxTaskCount returns a boolean if a field has been set.

### GetAccount

`func (o *AzureBlobStorageModel) GetAccount() AzureBlobStorageAccountModel`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AzureBlobStorageModel) GetAccountOk() (*AzureBlobStorageAccountModel, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AzureBlobStorageModel) SetAccount(v AzureBlobStorageAccountModel)`

SetAccount sets Account field to given value.


### GetContainer

`func (o *AzureBlobStorageModel) GetContainer() AzureBlobStorageContainerModel`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *AzureBlobStorageModel) GetContainerOk() (*AzureBlobStorageContainerModel, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *AzureBlobStorageModel) SetContainer(v AzureBlobStorageContainerModel)`

SetContainer sets Container field to given value.


### GetProxyAppliance

`func (o *AzureBlobStorageModel) GetProxyAppliance() AzureStorageProxyModel`

GetProxyAppliance returns the ProxyAppliance field if non-nil, zero value otherwise.

### GetProxyApplianceOk

`func (o *AzureBlobStorageModel) GetProxyApplianceOk() (*AzureStorageProxyModel, bool)`

GetProxyApplianceOk returns a tuple with the ProxyAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAppliance

`func (o *AzureBlobStorageModel) SetProxyAppliance(v AzureStorageProxyModel)`

SetProxyAppliance sets ProxyAppliance field to given value.

### HasProxyAppliance

`func (o *AzureBlobStorageModel) HasProxyAppliance() bool`

HasProxyAppliance returns a boolean if a field has been set.

### GetMountServer

`func (o *AzureBlobStorageModel) GetMountServer() MountServerSettingsModel`

GetMountServer returns the MountServer field if non-nil, zero value otherwise.

### GetMountServerOk

`func (o *AzureBlobStorageModel) GetMountServerOk() (*MountServerSettingsModel, bool)`

GetMountServerOk returns a tuple with the MountServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServer

`func (o *AzureBlobStorageModel) SetMountServer(v MountServerSettingsModel)`

SetMountServer sets MountServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


