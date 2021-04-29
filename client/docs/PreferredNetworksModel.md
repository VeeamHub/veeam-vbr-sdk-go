# PreferredNetworksModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, prefered networks are enabled. | 
**Networks** | Pointer to [**[]PreferredNetworkModel**](PreferredNetworkModel.md) | Array of networks. | [optional] 

## Methods

### NewPreferredNetworksModel

`func NewPreferredNetworksModel(isEnabled bool, ) *PreferredNetworksModel`

NewPreferredNetworksModel instantiates a new PreferredNetworksModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreferredNetworksModelWithDefaults

`func NewPreferredNetworksModelWithDefaults() *PreferredNetworksModel`

NewPreferredNetworksModelWithDefaults instantiates a new PreferredNetworksModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *PreferredNetworksModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *PreferredNetworksModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *PreferredNetworksModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetNetworks

`func (o *PreferredNetworksModel) GetNetworks() []PreferredNetworkModel`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *PreferredNetworksModel) GetNetworksOk() (*[]PreferredNetworkModel, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *PreferredNetworksModel) SetNetworks(v []PreferredNetworkModel)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *PreferredNetworksModel) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


