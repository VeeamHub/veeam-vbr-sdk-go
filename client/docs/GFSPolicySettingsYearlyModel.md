# GFSPolicySettingsYearlyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, the yearly GFS retention policy is enabled. | 
**KeepForNumberOfYears** | Pointer to **int32** | Number of years to keep full backups for archival purposes. Possible values are from 1 through 999. | [optional] 
**DesiredTime** | Pointer to [**EMonth**](EMonth.md) |  | [optional] 

## Methods

### NewGFSPolicySettingsYearlyModel

`func NewGFSPolicySettingsYearlyModel(isEnabled bool, ) *GFSPolicySettingsYearlyModel`

NewGFSPolicySettingsYearlyModel instantiates a new GFSPolicySettingsYearlyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGFSPolicySettingsYearlyModelWithDefaults

`func NewGFSPolicySettingsYearlyModelWithDefaults() *GFSPolicySettingsYearlyModel`

NewGFSPolicySettingsYearlyModelWithDefaults instantiates a new GFSPolicySettingsYearlyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *GFSPolicySettingsYearlyModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *GFSPolicySettingsYearlyModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *GFSPolicySettingsYearlyModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetKeepForNumberOfYears

`func (o *GFSPolicySettingsYearlyModel) GetKeepForNumberOfYears() int32`

GetKeepForNumberOfYears returns the KeepForNumberOfYears field if non-nil, zero value otherwise.

### GetKeepForNumberOfYearsOk

`func (o *GFSPolicySettingsYearlyModel) GetKeepForNumberOfYearsOk() (*int32, bool)`

GetKeepForNumberOfYearsOk returns a tuple with the KeepForNumberOfYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepForNumberOfYears

`func (o *GFSPolicySettingsYearlyModel) SetKeepForNumberOfYears(v int32)`

SetKeepForNumberOfYears sets KeepForNumberOfYears field to given value.

### HasKeepForNumberOfYears

`func (o *GFSPolicySettingsYearlyModel) HasKeepForNumberOfYears() bool`

HasKeepForNumberOfYears returns a boolean if a field has been set.

### GetDesiredTime

`func (o *GFSPolicySettingsYearlyModel) GetDesiredTime() EMonth`

GetDesiredTime returns the DesiredTime field if non-nil, zero value otherwise.

### GetDesiredTimeOk

`func (o *GFSPolicySettingsYearlyModel) GetDesiredTimeOk() (*EMonth, bool)`

GetDesiredTimeOk returns a tuple with the DesiredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredTime

`func (o *GFSPolicySettingsYearlyModel) SetDesiredTime(v EMonth)`

SetDesiredTime sets DesiredTime field to given value.

### HasDesiredTime

`func (o *GFSPolicySettingsYearlyModel) HasDesiredTime() bool`

HasDesiredTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


