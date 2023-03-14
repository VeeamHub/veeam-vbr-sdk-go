# AzureSubscriptionBrowserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID that Veeam Backup &amp; Replication assigned to the Azure subscription. | 
**AzureSubscriptionId** | Pointer to **string** | Original Azure subscription ID. | [optional] 
**Locations** | Pointer to [**[]AzureLocationBrowserModel**](AzureLocationBrowserModel.md) | Array of Azure geographic regions. | [optional] 

## Methods

### NewAzureSubscriptionBrowserModel

`func NewAzureSubscriptionBrowserModel(id string, ) *AzureSubscriptionBrowserModel`

NewAzureSubscriptionBrowserModel instantiates a new AzureSubscriptionBrowserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureSubscriptionBrowserModelWithDefaults

`func NewAzureSubscriptionBrowserModelWithDefaults() *AzureSubscriptionBrowserModel`

NewAzureSubscriptionBrowserModelWithDefaults instantiates a new AzureSubscriptionBrowserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AzureSubscriptionBrowserModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureSubscriptionBrowserModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureSubscriptionBrowserModel) SetId(v string)`

SetId sets Id field to given value.


### GetAzureSubscriptionId

`func (o *AzureSubscriptionBrowserModel) GetAzureSubscriptionId() string`

GetAzureSubscriptionId returns the AzureSubscriptionId field if non-nil, zero value otherwise.

### GetAzureSubscriptionIdOk

`func (o *AzureSubscriptionBrowserModel) GetAzureSubscriptionIdOk() (*string, bool)`

GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionId

`func (o *AzureSubscriptionBrowserModel) SetAzureSubscriptionId(v string)`

SetAzureSubscriptionId sets AzureSubscriptionId field to given value.

### HasAzureSubscriptionId

`func (o *AzureSubscriptionBrowserModel) HasAzureSubscriptionId() bool`

HasAzureSubscriptionId returns a boolean if a field has been set.

### GetLocations

`func (o *AzureSubscriptionBrowserModel) GetLocations() []AzureLocationBrowserModel`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *AzureSubscriptionBrowserModel) GetLocationsOk() (*[]AzureLocationBrowserModel, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *AzureSubscriptionBrowserModel) SetLocations(v []AzureLocationBrowserModel)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *AzureSubscriptionBrowserModel) HasLocations() bool`

HasLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


