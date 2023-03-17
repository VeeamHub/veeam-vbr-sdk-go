# AzureComputeBrowserFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShowAllStorageAccounts** | Pointer to **bool** | If *true*, the result contains compute resorces for all storage accounts. If *false*, the result contains compute resorces available for the specified storage account only. | [optional] [default to false]
**SubscriptionId** | Pointer to **string** | Filters compute resorces by ID that Veeam Backup &amp; Replication assigned to the Azure subscription. | [optional] 
**Location** | Pointer to **string** | Filters compute resorces by Azure location name. | [optional] 
**HasNetworks** | Pointer to **bool** | If *true*, the result contains Azure resource groups with virtual networks only. | [optional] [default to false]

## Methods

### NewAzureComputeBrowserFilters

`func NewAzureComputeBrowserFilters() *AzureComputeBrowserFilters`

NewAzureComputeBrowserFilters instantiates a new AzureComputeBrowserFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureComputeBrowserFiltersWithDefaults

`func NewAzureComputeBrowserFiltersWithDefaults() *AzureComputeBrowserFilters`

NewAzureComputeBrowserFiltersWithDefaults instantiates a new AzureComputeBrowserFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShowAllStorageAccounts

`func (o *AzureComputeBrowserFilters) GetShowAllStorageAccounts() bool`

GetShowAllStorageAccounts returns the ShowAllStorageAccounts field if non-nil, zero value otherwise.

### GetShowAllStorageAccountsOk

`func (o *AzureComputeBrowserFilters) GetShowAllStorageAccountsOk() (*bool, bool)`

GetShowAllStorageAccountsOk returns a tuple with the ShowAllStorageAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAllStorageAccounts

`func (o *AzureComputeBrowserFilters) SetShowAllStorageAccounts(v bool)`

SetShowAllStorageAccounts sets ShowAllStorageAccounts field to given value.

### HasShowAllStorageAccounts

`func (o *AzureComputeBrowserFilters) HasShowAllStorageAccounts() bool`

HasShowAllStorageAccounts returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *AzureComputeBrowserFilters) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AzureComputeBrowserFilters) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AzureComputeBrowserFilters) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *AzureComputeBrowserFilters) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetLocation

`func (o *AzureComputeBrowserFilters) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AzureComputeBrowserFilters) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AzureComputeBrowserFilters) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AzureComputeBrowserFilters) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetHasNetworks

`func (o *AzureComputeBrowserFilters) GetHasNetworks() bool`

GetHasNetworks returns the HasNetworks field if non-nil, zero value otherwise.

### GetHasNetworksOk

`func (o *AzureComputeBrowserFilters) GetHasNetworksOk() (*bool, bool)`

GetHasNetworksOk returns a tuple with the HasNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNetworks

`func (o *AzureComputeBrowserFilters) SetHasNetworks(v bool)`

SetHasNetworks sets HasNetworks field to given value.

### HasHasNetworks

`func (o *AzureComputeBrowserFilters) HasHasNetworks() bool`

HasHasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


