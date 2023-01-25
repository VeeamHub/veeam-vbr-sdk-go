# AzureComputeBrowserModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionType** | Pointer to **string** | Azure region type. | [optional] 
**Subscriptions** | Pointer to [**[]AzureSubscriptionBrowserModel**](AzureSubscriptionBrowserModel.md) | Array of Azure subscriptions associated with the account. | [optional] 

## Methods

### NewAzureComputeBrowserModelAllOf

`func NewAzureComputeBrowserModelAllOf() *AzureComputeBrowserModelAllOf`

NewAzureComputeBrowserModelAllOf instantiates a new AzureComputeBrowserModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureComputeBrowserModelAllOfWithDefaults

`func NewAzureComputeBrowserModelAllOfWithDefaults() *AzureComputeBrowserModelAllOf`

NewAzureComputeBrowserModelAllOfWithDefaults instantiates a new AzureComputeBrowserModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionType

`func (o *AzureComputeBrowserModelAllOf) GetRegionType() string`

GetRegionType returns the RegionType field if non-nil, zero value otherwise.

### GetRegionTypeOk

`func (o *AzureComputeBrowserModelAllOf) GetRegionTypeOk() (*string, bool)`

GetRegionTypeOk returns a tuple with the RegionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionType

`func (o *AzureComputeBrowserModelAllOf) SetRegionType(v string)`

SetRegionType sets RegionType field to given value.

### HasRegionType

`func (o *AzureComputeBrowserModelAllOf) HasRegionType() bool`

HasRegionType returns a boolean if a field has been set.

### GetSubscriptions

`func (o *AzureComputeBrowserModelAllOf) GetSubscriptions() []AzureSubscriptionBrowserModel`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *AzureComputeBrowserModelAllOf) GetSubscriptionsOk() (*[]AzureSubscriptionBrowserModel, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *AzureComputeBrowserModelAllOf) SetSubscriptions(v []AzureSubscriptionBrowserModel)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *AzureComputeBrowserModelAllOf) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


