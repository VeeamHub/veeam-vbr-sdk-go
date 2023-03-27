# AzureArchiveStorageImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the object storage repository. | 
**Description** | **string** | Description of the object storage repository. | 
**Tag** | **string** | Tag that identifies the object storage repository. | 
**Type** | [**ERepositoryType**](ERepositoryType.md) |  | 
**Account** | [**AzureArchiveStorageAccountImportModel**](AzureArchiveStorageAccountImportModel.md) |  | 
**Container** | [**AzureArchiveStorageContainerModel**](AzureArchiveStorageContainerModel.md) |  | 
**ProxyAppliance** | Pointer to [**AzureStorageProxyModel**](AzureStorageProxyModel.md) |  | [optional] 

## Methods

### NewAzureArchiveStorageImportSpec

`func NewAzureArchiveStorageImportSpec(name string, description string, tag string, type_ ERepositoryType, account AzureArchiveStorageAccountImportModel, container AzureArchiveStorageContainerModel, ) *AzureArchiveStorageImportSpec`

NewAzureArchiveStorageImportSpec instantiates a new AzureArchiveStorageImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureArchiveStorageImportSpecWithDefaults

`func NewAzureArchiveStorageImportSpecWithDefaults() *AzureArchiveStorageImportSpec`

NewAzureArchiveStorageImportSpecWithDefaults instantiates a new AzureArchiveStorageImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AzureArchiveStorageImportSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureArchiveStorageImportSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureArchiveStorageImportSpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AzureArchiveStorageImportSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AzureArchiveStorageImportSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AzureArchiveStorageImportSpec) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTag

`func (o *AzureArchiveStorageImportSpec) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AzureArchiveStorageImportSpec) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AzureArchiveStorageImportSpec) SetTag(v string)`

SetTag sets Tag field to given value.


### GetType

`func (o *AzureArchiveStorageImportSpec) GetType() ERepositoryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AzureArchiveStorageImportSpec) GetTypeOk() (*ERepositoryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AzureArchiveStorageImportSpec) SetType(v ERepositoryType)`

SetType sets Type field to given value.


### GetAccount

`func (o *AzureArchiveStorageImportSpec) GetAccount() AzureArchiveStorageAccountImportModel`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AzureArchiveStorageImportSpec) GetAccountOk() (*AzureArchiveStorageAccountImportModel, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AzureArchiveStorageImportSpec) SetAccount(v AzureArchiveStorageAccountImportModel)`

SetAccount sets Account field to given value.


### GetContainer

`func (o *AzureArchiveStorageImportSpec) GetContainer() AzureArchiveStorageContainerModel`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *AzureArchiveStorageImportSpec) GetContainerOk() (*AzureArchiveStorageContainerModel, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *AzureArchiveStorageImportSpec) SetContainer(v AzureArchiveStorageContainerModel)`

SetContainer sets Container field to given value.


### GetProxyAppliance

`func (o *AzureArchiveStorageImportSpec) GetProxyAppliance() AzureStorageProxyModel`

GetProxyAppliance returns the ProxyAppliance field if non-nil, zero value otherwise.

### GetProxyApplianceOk

`func (o *AzureArchiveStorageImportSpec) GetProxyApplianceOk() (*AzureStorageProxyModel, bool)`

GetProxyApplianceOk returns a tuple with the ProxyAppliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAppliance

`func (o *AzureArchiveStorageImportSpec) SetProxyAppliance(v AzureStorageProxyModel)`

SetProxyAppliance sets ProxyAppliance field to given value.

### HasProxyAppliance

`func (o *AzureArchiveStorageImportSpec) HasProxyAppliance() bool`

HasProxyAppliance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


