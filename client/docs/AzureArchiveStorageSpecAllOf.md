# AzureArchiveStorageSpecAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**AzureArchiveStorageAccountModel**](AzureArchiveStorageAccountModel.md) |  | 
**Container** | [**AzureArchiveStorageContainerModel**](AzureArchiveStorageContainerModel.md) |  | 
**ProxyAppliance** | [**AzureStorageProxyModel**](AzureStorageProxyModel.md) |  | 

## Methods

### NewAzureArchiveStorageSpecAllOf

`func NewAzureArchiveStorageSpecAllOf(account AzureArchiveStorageAccountModel, container AzureArchiveStorageContainerModel, proxyAppliance AzureStorageProxyModel, ) *AzureArchiveStorageSpecAllOf`

NewAzureArchiveStorageSpecAllOf instantiates a new AzureArchiveStorageSpecAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureArchiveStorageSpecAllOfWithDefaults

`func NewAzureArchiveStorageSpecAllOfWithDefaults() *AzureArchiveStorageSpecAllOf`

NewAzureArchiveStorageSpecAllOfWithDefaults instantiates a new AzureArchiveStorageSpecAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AzureArchiveStorageSpecAllOf) GetAccount() AzureArchiveStorageAccountModel`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AzureArchiveStorageSpecAllOf) GetAccountOk() (*AzureArchiveStorageAccountModel, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AzureArchiveStorageSpecAllOf) SetAccount(v AzureArchiveStorageAccountModel)`

SetAccount sets Account field to given value.


### GetContainer

`func (o *AzureArchiveStorageSpecAllOf) GetContainer() AzureArchiveStorageContainerModel`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *AzureArchiveStorageSpecAllOf) GetContainerOk() (*AzureArchiveStorageContainerModel, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *AzureArchiveStorageSpecAllOf) SetContainer(v AzureArchiveStorageContainerModel)`

SetContainer sets Container field to given value.


### GetProxyAppliance

`func (o *AzureArchiveStorageSpecAllOf) GetProxyAppliance() AzureStorageProxyModel`

GetProxyAppliance returns the ProxyAppliance field if non-nil, zero value otherwise.

### GetProxyApplianceOk

`func (o *AzureArchiveStorageSpecAllOf) GetProxyApplianceOk() (*AzureStorageProxyModel, bool)`

GetProxyApplianceOk returns a tuple with the ProxyAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAppliance

`func (o *AzureArchiveStorageSpecAllOf) SetProxyAppliance(v AzureStorageProxyModel)`

SetProxyAppliance sets ProxyAppliance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


