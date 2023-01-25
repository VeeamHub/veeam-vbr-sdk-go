# AzureLocationBrowserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | **string** | Location name. | 
**StorageAccounts** | [**[]AzureStorageAccountBrowserModel**](AzureStorageAccountBrowserModel.md) | Array of storage accounts associated with the location. | 
**ResourceGroups** | [**[]AzureResourceGroupBrowserModel**](AzureResourceGroupBrowserModel.md) | Array of Azure resource groups. | 

## Methods

### NewAzureLocationBrowserModel

`func NewAzureLocationBrowserModel(location string, storageAccounts []AzureStorageAccountBrowserModel, resourceGroups []AzureResourceGroupBrowserModel, ) *AzureLocationBrowserModel`

NewAzureLocationBrowserModel instantiates a new AzureLocationBrowserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureLocationBrowserModelWithDefaults

`func NewAzureLocationBrowserModelWithDefaults() *AzureLocationBrowserModel`

NewAzureLocationBrowserModelWithDefaults instantiates a new AzureLocationBrowserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *AzureLocationBrowserModel) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AzureLocationBrowserModel) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AzureLocationBrowserModel) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetStorageAccounts

`func (o *AzureLocationBrowserModel) GetStorageAccounts() []AzureStorageAccountBrowserModel`

GetStorageAccounts returns the StorageAccounts field if non-nil, zero value otherwise.

### GetStorageAccountsOk

`func (o *AzureLocationBrowserModel) GetStorageAccountsOk() (*[]AzureStorageAccountBrowserModel, bool)`

GetStorageAccountsOk returns a tuple with the StorageAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccounts

`func (o *AzureLocationBrowserModel) SetStorageAccounts(v []AzureStorageAccountBrowserModel)`

SetStorageAccounts sets StorageAccounts field to given value.


### GetResourceGroups

`func (o *AzureLocationBrowserModel) GetResourceGroups() []AzureResourceGroupBrowserModel`

GetResourceGroups returns the ResourceGroups field if non-nil, zero value otherwise.

### GetResourceGroupsOk

`func (o *AzureLocationBrowserModel) GetResourceGroupsOk() (*[]AzureResourceGroupBrowserModel, bool)`

GetResourceGroupsOk returns a tuple with the ResourceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroups

`func (o *AzureLocationBrowserModel) SetResourceGroups(v []AzureResourceGroupBrowserModel)`

SetResourceGroups sets ResourceGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


