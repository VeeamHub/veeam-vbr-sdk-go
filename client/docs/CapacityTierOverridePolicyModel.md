# CapacityTierOverridePolicyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, Veeam Backup &amp; Replication moves oldest backup files sooner if the scale-out backup repository is reaching the threshold. | 
**OverrideSpaceThresholdPercents** | Pointer to **int32** | Space threshold of the scale-out backup repository, in percent. | [optional] 

## Methods

### NewCapacityTierOverridePolicyModel

`func NewCapacityTierOverridePolicyModel(isEnabled bool, ) *CapacityTierOverridePolicyModel`

NewCapacityTierOverridePolicyModel instantiates a new CapacityTierOverridePolicyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapacityTierOverridePolicyModelWithDefaults

`func NewCapacityTierOverridePolicyModelWithDefaults() *CapacityTierOverridePolicyModel`

NewCapacityTierOverridePolicyModelWithDefaults instantiates a new CapacityTierOverridePolicyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *CapacityTierOverridePolicyModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *CapacityTierOverridePolicyModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *CapacityTierOverridePolicyModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetOverrideSpaceThresholdPercents

`func (o *CapacityTierOverridePolicyModel) GetOverrideSpaceThresholdPercents() int32`

GetOverrideSpaceThresholdPercents returns the OverrideSpaceThresholdPercents field if non-nil, zero value otherwise.

### GetOverrideSpaceThresholdPercentsOk

`func (o *CapacityTierOverridePolicyModel) GetOverrideSpaceThresholdPercentsOk() (*int32, bool)`

GetOverrideSpaceThresholdPercentsOk returns a tuple with the OverrideSpaceThresholdPercents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSpaceThresholdPercents

`func (o *CapacityTierOverridePolicyModel) SetOverrideSpaceThresholdPercents(v int32)`

SetOverrideSpaceThresholdPercents sets OverrideSpaceThresholdPercents field to given value.

### HasOverrideSpaceThresholdPercents

`func (o *CapacityTierOverridePolicyModel) HasOverrideSpaceThresholdPercents() bool`

HasOverrideSpaceThresholdPercents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


