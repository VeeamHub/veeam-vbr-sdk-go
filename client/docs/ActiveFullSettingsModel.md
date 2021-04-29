# ActiveFullSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, active full backups are enabled. | 
**Weekly** | Pointer to [**AdvancedStorageScheduleWeeklyModel**](AdvancedStorageScheduleWeeklyModel.md) |  | [optional] 
**Monthly** | Pointer to [**AdvancedStorageScheduleMonthlyModel**](AdvancedStorageScheduleMonthlyModel.md) |  | [optional] 

## Methods

### NewActiveFullSettingsModel

`func NewActiveFullSettingsModel(isEnabled bool, ) *ActiveFullSettingsModel`

NewActiveFullSettingsModel instantiates a new ActiveFullSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveFullSettingsModelWithDefaults

`func NewActiveFullSettingsModelWithDefaults() *ActiveFullSettingsModel`

NewActiveFullSettingsModelWithDefaults instantiates a new ActiveFullSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *ActiveFullSettingsModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ActiveFullSettingsModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ActiveFullSettingsModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetWeekly

`func (o *ActiveFullSettingsModel) GetWeekly() AdvancedStorageScheduleWeeklyModel`

GetWeekly returns the Weekly field if non-nil, zero value otherwise.

### GetWeeklyOk

`func (o *ActiveFullSettingsModel) GetWeeklyOk() (*AdvancedStorageScheduleWeeklyModel, bool)`

GetWeeklyOk returns a tuple with the Weekly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekly

`func (o *ActiveFullSettingsModel) SetWeekly(v AdvancedStorageScheduleWeeklyModel)`

SetWeekly sets Weekly field to given value.

### HasWeekly

`func (o *ActiveFullSettingsModel) HasWeekly() bool`

HasWeekly returns a boolean if a field has been set.

### GetMonthly

`func (o *ActiveFullSettingsModel) GetMonthly() AdvancedStorageScheduleMonthlyModel`

GetMonthly returns the Monthly field if non-nil, zero value otherwise.

### GetMonthlyOk

`func (o *ActiveFullSettingsModel) GetMonthlyOk() (*AdvancedStorageScheduleMonthlyModel, bool)`

GetMonthlyOk returns a tuple with the Monthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthly

`func (o *ActiveFullSettingsModel) SetMonthly(v AdvancedStorageScheduleMonthlyModel)`

SetMonthly sets Monthly field to given value.

### HasMonthly

`func (o *ActiveFullSettingsModel) HasMonthly() bool`

HasMonthly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


