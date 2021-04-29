# BackupWindowSettingModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | [**[]BackupWindowDayHoursModel**](BackupWindowDayHoursModel.md) | Array of per-day schemes. | 

## Methods

### NewBackupWindowSettingModel

`func NewBackupWindowSettingModel(days []BackupWindowDayHoursModel, ) *BackupWindowSettingModel`

NewBackupWindowSettingModel instantiates a new BackupWindowSettingModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupWindowSettingModelWithDefaults

`func NewBackupWindowSettingModelWithDefaults() *BackupWindowSettingModel`

NewBackupWindowSettingModelWithDefaults instantiates a new BackupWindowSettingModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *BackupWindowSettingModel) GetDays() []BackupWindowDayHoursModel`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *BackupWindowSettingModel) GetDaysOk() (*[]BackupWindowDayHoursModel, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *BackupWindowSettingModel) SetDays(v []BackupWindowDayHoursModel)`

SetDays sets Days field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


