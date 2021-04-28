# GFSPolicySettingsWeeklyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, the weekly GFS retention policy is enabled. | 
**KeepForNumberOfWeeks** | Pointer to **int32** | Number of weeks to keep full backups for archival purposes. | [optional] 
**DesiredTime** | Pointer to [**EDayOfWeek**](EDayOfWeek.md) |  | [optional] 

## Methods

### NewGFSPolicySettingsWeeklyModel

`func NewGFSPolicySettingsWeeklyModel(isEnabled bool, ) *GFSPolicySettingsWeeklyModel`

NewGFSPolicySettingsWeeklyModel instantiates a new GFSPolicySettingsWeeklyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGFSPolicySettingsWeeklyModelWithDefaults

`func NewGFSPolicySettingsWeeklyModelWithDefaults() *GFSPolicySettingsWeeklyModel`

NewGFSPolicySettingsWeeklyModelWithDefaults instantiates a new GFSPolicySettingsWeeklyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *GFSPolicySettingsWeeklyModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *GFSPolicySettingsWeeklyModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *GFSPolicySettingsWeeklyModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetKeepForNumberOfWeeks

`func (o *GFSPolicySettingsWeeklyModel) GetKeepForNumberOfWeeks() int32`

GetKeepForNumberOfWeeks returns the KeepForNumberOfWeeks field if non-nil, zero value otherwise.

### GetKeepForNumberOfWeeksOk

`func (o *GFSPolicySettingsWeeklyModel) GetKeepForNumberOfWeeksOk() (*int32, bool)`

GetKeepForNumberOfWeeksOk returns a tuple with the KeepForNumberOfWeeks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepForNumberOfWeeks

`func (o *GFSPolicySettingsWeeklyModel) SetKeepForNumberOfWeeks(v int32)`

SetKeepForNumberOfWeeks sets KeepForNumberOfWeeks field to given value.

### HasKeepForNumberOfWeeks

`func (o *GFSPolicySettingsWeeklyModel) HasKeepForNumberOfWeeks() bool`

HasKeepForNumberOfWeeks returns a boolean if a field has been set.

### GetDesiredTime

`func (o *GFSPolicySettingsWeeklyModel) GetDesiredTime() EDayOfWeek`

GetDesiredTime returns the DesiredTime field if non-nil, zero value otherwise.

### GetDesiredTimeOk

`func (o *GFSPolicySettingsWeeklyModel) GetDesiredTimeOk() (*EDayOfWeek, bool)`

GetDesiredTimeOk returns a tuple with the DesiredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredTime

`func (o *GFSPolicySettingsWeeklyModel) SetDesiredTime(v EDayOfWeek)`

SetDesiredTime sets DesiredTime field to given value.

### HasDesiredTime

`func (o *GFSPolicySettingsWeeklyModel) HasDesiredTime() bool`

HasDesiredTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


