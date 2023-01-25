# AzureStorageAccountBrowserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageAccountName** | Pointer to **string** | Account name. | [optional] 
**InstanceSizes** | Pointer to **[]string** | Array of available instance sizes. | [optional] 

## Methods

### NewAzureStorageAccountBrowserModel

`func NewAzureStorageAccountBrowserModel() *AzureStorageAccountBrowserModel`

NewAzureStorageAccountBrowserModel instantiates a new AzureStorageAccountBrowserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureStorageAccountBrowserModelWithDefaults

`func NewAzureStorageAccountBrowserModelWithDefaults() *AzureStorageAccountBrowserModel`

NewAzureStorageAccountBrowserModelWithDefaults instantiates a new AzureStorageAccountBrowserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageAccountName

`func (o *AzureStorageAccountBrowserModel) GetStorageAccountName() string`

GetStorageAccountName returns the StorageAccountName field if non-nil, zero value otherwise.

### GetStorageAccountNameOk

`func (o *AzureStorageAccountBrowserModel) GetStorageAccountNameOk() (*string, bool)`

GetStorageAccountNameOk returns a tuple with the StorageAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccountName

`func (o *AzureStorageAccountBrowserModel) SetStorageAccountName(v string)`

SetStorageAccountName sets StorageAccountName field to given value.

### HasStorageAccountName

`func (o *AzureStorageAccountBrowserModel) HasStorageAccountName() bool`

HasStorageAccountName returns a boolean if a field has been set.

### GetInstanceSizes

`func (o *AzureStorageAccountBrowserModel) GetInstanceSizes() []string`

GetInstanceSizes returns the InstanceSizes field if non-nil, zero value otherwise.

### GetInstanceSizesOk

`func (o *AzureStorageAccountBrowserModel) GetInstanceSizesOk() (*[]string, bool)`

GetInstanceSizesOk returns a tuple with the InstanceSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSizes

`func (o *AzureStorageAccountBrowserModel) SetInstanceSizes(v []string)`

SetInstanceSizes sets InstanceSizes field to given value.

### HasInstanceSizes

`func (o *AzureStorageAccountBrowserModel) HasInstanceSizes() bool`

HasInstanceSizes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


