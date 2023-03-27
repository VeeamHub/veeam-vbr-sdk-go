# SyntheticFullSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, active full backups are enabled. | 
**Weekly** | Pointer to [**AdvancedStorageScheduleWeeklyModel**](AdvancedStorageScheduleWeeklyModel.md) |  | [optional] 
**Monthly** | Pointer to [**AdvancedStorageScheduleMonthlyModel**](AdvancedStorageScheduleMonthlyModel.md) |  | [optional] 

## Methods

### NewSyntheticFullSettingsModel

`func NewSyntheticFullSettingsModel(isEnabled bool, ) *SyntheticFullSettingsModel`

NewSyntheticFullSettingsModel instantiates a new SyntheticFullSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticFullSettingsModelWithDefaults

`func NewSyntheticFullSettingsModelWithDefaults() *SyntheticFullSettingsModel`

NewSyntheticFullSettingsModelWithDefaults instantiates a new SyntheticFullSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *SyntheticFullSettingsModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *SyntheticFullSettingsModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *SyntheticFullSettingsModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetWeekly

`func (o *SyntheticFullSettingsModel) GetWeekly() AdvancedStorageScheduleWeeklyModel`

GetWeekly returns the Weekly field if non-nil, zero value otherwise.

### GetWeeklyOk

`func (o *SyntheticFullSettingsModel) GetWeeklyOk() (*AdvancedStorageScheduleWeeklyModel, bool)`

GetWeeklyOk returns a tuple with the Weekly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekly

`func (o *SyntheticFullSettingsModel) SetWeekly(v AdvancedStorageScheduleWeeklyModel)`

SetWeekly sets Weekly field to given value.

### HasWeekly

`func (o *SyntheticFullSettingsModel) HasWeekly() bool`

HasWeekly returns a boolean if a field has been set.

### GetMonthly

`func (o *SyntheticFullSettingsModel) GetMonthly() AdvancedStorageScheduleMonthlyModel`

GetMonthly returns the Monthly field if non-nil, zero value otherwise.

### GetMonthlyOk

`func (o *SyntheticFullSettingsModel) GetMonthlyOk() (*AdvancedStorageScheduleMonthlyModel, bool)`

GetMonthlyOk returns a tuple with the Monthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthly

`func (o *SyntheticFullSettingsModel) SetMonthly(v AdvancedStorageScheduleMonthlyModel)`

SetMonthly sets Monthly field to given value.

### HasMonthly

`func (o *SyntheticFullSettingsModel) HasMonthly() bool`

HasMonthly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


