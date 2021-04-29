# GFSPolicySettingsMonthlyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, the monthly GFS retention policy is enabled. | 
**KeepForNumberOfMonths** | Pointer to **int32** | Number of months to keep full backups for archival purposes. | [optional] 
**DesiredTime** | Pointer to [**ESennightOfMonth**](ESennightOfMonth.md) |  | [optional] 

## Methods

### NewGFSPolicySettingsMonthlyModel

`func NewGFSPolicySettingsMonthlyModel(isEnabled bool, ) *GFSPolicySettingsMonthlyModel`

NewGFSPolicySettingsMonthlyModel instantiates a new GFSPolicySettingsMonthlyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGFSPolicySettingsMonthlyModelWithDefaults

`func NewGFSPolicySettingsMonthlyModelWithDefaults() *GFSPolicySettingsMonthlyModel`

NewGFSPolicySettingsMonthlyModelWithDefaults instantiates a new GFSPolicySettingsMonthlyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *GFSPolicySettingsMonthlyModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *GFSPolicySettingsMonthlyModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *GFSPolicySettingsMonthlyModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetKeepForNumberOfMonths

`func (o *GFSPolicySettingsMonthlyModel) GetKeepForNumberOfMonths() int32`

GetKeepForNumberOfMonths returns the KeepForNumberOfMonths field if non-nil, zero value otherwise.

### GetKeepForNumberOfMonthsOk

`func (o *GFSPolicySettingsMonthlyModel) GetKeepForNumberOfMonthsOk() (*int32, bool)`

GetKeepForNumberOfMonthsOk returns a tuple with the KeepForNumberOfMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepForNumberOfMonths

`func (o *GFSPolicySettingsMonthlyModel) SetKeepForNumberOfMonths(v int32)`

SetKeepForNumberOfMonths sets KeepForNumberOfMonths field to given value.

### HasKeepForNumberOfMonths

`func (o *GFSPolicySettingsMonthlyModel) HasKeepForNumberOfMonths() bool`

HasKeepForNumberOfMonths returns a boolean if a field has been set.

### GetDesiredTime

`func (o *GFSPolicySettingsMonthlyModel) GetDesiredTime() ESennightOfMonth`

GetDesiredTime returns the DesiredTime field if non-nil, zero value otherwise.

### GetDesiredTimeOk

`func (o *GFSPolicySettingsMonthlyModel) GetDesiredTimeOk() (*ESennightOfMonth, bool)`

GetDesiredTimeOk returns a tuple with the DesiredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredTime

`func (o *GFSPolicySettingsMonthlyModel) SetDesiredTime(v ESennightOfMonth)`

SetDesiredTime sets DesiredTime field to given value.

### HasDesiredTime

`func (o *GFSPolicySettingsMonthlyModel) HasDesiredTime() bool`

HasDesiredTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


