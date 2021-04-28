# ConfigBackupScheduleModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, backup scheduling is enabled. | 
**Daily** | Pointer to [**ScheduleDailyModel**](ScheduleDailyModel.md) |  | [optional] 
**Monthly** | Pointer to [**ScheduleMonthlyModel**](ScheduleMonthlyModel.md) |  | [optional] 

## Methods

### NewConfigBackupScheduleModel

`func NewConfigBackupScheduleModel(isEnabled bool, ) *ConfigBackupScheduleModel`

NewConfigBackupScheduleModel instantiates a new ConfigBackupScheduleModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigBackupScheduleModelWithDefaults

`func NewConfigBackupScheduleModelWithDefaults() *ConfigBackupScheduleModel`

NewConfigBackupScheduleModelWithDefaults instantiates a new ConfigBackupScheduleModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *ConfigBackupScheduleModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ConfigBackupScheduleModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ConfigBackupScheduleModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetDaily

`func (o *ConfigBackupScheduleModel) GetDaily() ScheduleDailyModel`

GetDaily returns the Daily field if non-nil, zero value otherwise.

### GetDailyOk

`func (o *ConfigBackupScheduleModel) GetDailyOk() (*ScheduleDailyModel, bool)`

GetDailyOk returns a tuple with the Daily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaily

`func (o *ConfigBackupScheduleModel) SetDaily(v ScheduleDailyModel)`

SetDaily sets Daily field to given value.

### HasDaily

`func (o *ConfigBackupScheduleModel) HasDaily() bool`

HasDaily returns a boolean if a field has been set.

### GetMonthly

`func (o *ConfigBackupScheduleModel) GetMonthly() ScheduleMonthlyModel`

GetMonthly returns the Monthly field if non-nil, zero value otherwise.

### GetMonthlyOk

`func (o *ConfigBackupScheduleModel) GetMonthlyOk() (*ScheduleMonthlyModel, bool)`

GetMonthlyOk returns a tuple with the Monthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthly

`func (o *ConfigBackupScheduleModel) SetMonthly(v ScheduleMonthlyModel)`

SetMonthly sets Monthly field to given value.

### HasMonthly

`func (o *ConfigBackupScheduleModel) HasMonthly() bool`

HasMonthly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


