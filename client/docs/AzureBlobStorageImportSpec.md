# AzureBlobStorageImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the object storage repository. | 
**Description** | **string** | Description of the object storage repository. | 
**Tag** | **string** | Tag that identifies the object storage repository. | 
**Type** | [**ERepositoryType**](ERepositoryType.md) |  | 
**EnableTaskLimit** | Pointer to **bool** | If *true*, the maximum number of concurrent tasks is limited. | [optional] 
**MaxTaskCount** | Pointer to **int32** | Maximum number of concurrent tasks. | [optional] 
**Account** | [**AzureBlobStorageAccountImportModel**](AzureBlobStorageAccountImportModel.md) |  | 
**Container** | [**AzureBlobStorageContainerModel**](AzureBlobStorageContainerModel.md) |  | 
**ProxyAppliance** | Pointer to [**AzureStorageProxyModel**](AzureStorageProxyModel.md) |  | [optional] 
**MountServer** | [**MountServerSettingsImportSpec**](MountServerSettingsImportSpec.md) |  | 

## Methods

### NewAzureBlobStorageImportSpec

`func NewAzureBlobStorageImportSpec(name string, description string, tag string, type_ ERepositoryType, account AzureBlobStorageAccountImportModel, container AzureBlobStorageContainerModel, mountServer MountServerSettingsImportSpec, ) *AzureBlobStorageImportSpec`

NewAzureBlobStorageImportSpec instantiates a new AzureBlobStorageImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureBlobStorageImportSpecWithDefaults

`func NewAzureBlobStorageImportSpecWithDefaults() *AzureBlobStorageImportSpec`

NewAzureBlobStorageImportSpecWithDefaults instantiates a new AzureBlobStorageImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AzureBlobStorageImportSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureBlobStorageImportSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureBlobStorageImportSpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AzureBlobStorageImportSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AzureBlobStorageImportSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AzureBlobStorageImportSpec) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTag

`func (o *AzureBlobStorageImportSpec) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AzureBlobStorageImportSpec) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AzureBlobStorageImportSpec) SetTag(v string)`

SetTag sets Tag field to given value.


### GetType

`func (o *AzureBlobStorageImportSpec) GetType() ERepositoryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AzureBlobStorageImportSpec) GetTypeOk() (*ERepositoryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AzureBlobStorageImportSpec) SetType(v ERepositoryType)`

SetType sets Type field to given value.


### GetEnableTaskLimit

`func (o *AzureBlobStorageImportSpec) GetEnableTaskLimit() bool`

GetEnableTaskLimit returns the EnableTaskLimit field if non-nil, zero value otherwise.

### GetEnableTaskLimitOk

`func (o *AzureBlobStorageImportSpec) GetEnableTaskLimitOk() (*bool, bool)`

GetEnableTaskLimitOk returns a tuple with the EnableTaskLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTaskLimit

`func (o *AzureBlobStorageImportSpec) SetEnableTaskLimit(v bool)`

SetEnableTaskLimit sets EnableTaskLimit field to given value.

### HasEnableTaskLimit

`func (o *AzureBlobStorageImportSpec) HasEnableTaskLimit() bool`

HasEnableTaskLimit returns a boolean if a field has been set.

### GetMaxTaskCount

`func (o *AzureBlobStorageImportSpec) GetMaxTaskCount() int32`

GetMaxTaskCount returns the MaxTaskCount field if non-nil, zero value otherwise.

### GetMaxTaskCountOk

`func (o *AzureBlobStorageImportSpec) GetMaxTaskCountOk() (*int32, bool)`

GetMaxTaskCountOk returns a tuple with the MaxTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskCount

`func (o *AzureBlobStorageImportSpec) SetMaxTaskCount(v int32)`

SetMaxTaskCount sets MaxTaskCount field to given value.

### HasMaxTaskCount

`func (o *AzureBlobStorageImportSpec) HasMaxTaskCount() bool`

HasMaxTaskCount returns a boolean if a field has been set.

### GetAccount

`func (o *AzureBlobStorageImportSpec) GetAccount() AzureBlobStorageAccountImportModel`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AzureBlobStorageImportSpec) GetAccountOk() (*AzureBlobStorageAccountImportModel, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AzureBlobStorageImportSpec) SetAccount(v AzureBlobStorageAccountImportModel)`

SetAccount sets Account field to given value.


### GetContainer

`func (o *AzureBlobStorageImportSpec) GetContainer() AzureBlobStorageContainerModel`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *AzureBlobStorageImportSpec) GetContainerOk() (*AzureBlobStorageContainerModel, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *AzureBlobStorageImportSpec) SetContainer(v AzureBlobStorageContainerModel)`

SetContainer sets Container field to given value.


### GetProxyAppliance

`func (o *AzureBlobStorageImportSpec) GetProxyAppliance() AzureStorageProxyModel`

GetProxyAppliance returns the ProxyAppliance field if non-nil, zero value otherwise.

### GetProxyApplianceOk

`func (o *AzureBlobStorageImportSpec) GetProxyApplianceOk() (*AzureStorageProxyModel, bool)`

GetProxyApplianceOk returns a tuple with the ProxyAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAppliance

`func (o *AzureBlobStorageImportSpec) SetProxyAppliance(v AzureStorageProxyModel)`

SetProxyAppliance sets ProxyAppliance field to given value.

### HasProxyAppliance

`func (o *AzureBlobStorageImportSpec) HasProxyAppliance() bool`

HasProxyAppliance returns a boolean if a field has been set.

### GetMountServer

`func (o *AzureBlobStorageImportSpec) GetMountServer() MountServerSettingsImportSpec`

GetMountServer returns the MountServer field if non-nil, zero value otherwise.

### GetMountServerOk

`func (o *AzureBlobStorageImportSpec) GetMountServerOk() (*MountServerSettingsImportSpec, bool)`

GetMountServerOk returns a tuple with the MountServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServer

`func (o *AzureBlobStorageImportSpec) SetMountServer(v MountServerSettingsImportSpec)`

SetMountServer sets MountServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


