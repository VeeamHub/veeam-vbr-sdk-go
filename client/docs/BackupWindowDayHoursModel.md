# BackupWindowDayHoursModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Day** | [**EDayOfWeek**](EDayOfWeek.md) |  | 
**Hours** | **string** | String of hours in the following format: *1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1* where *1* means enabled, *0* means disabled.  | 

## Methods

### NewBackupWindowDayHoursModel

`func NewBackupWindowDayHoursModel(day EDayOfWeek, hours string, ) *BackupWindowDayHoursModel`

NewBackupWindowDayHoursModel instantiates a new BackupWindowDayHoursModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupWindowDayHoursModelWithDefaults

`func NewBackupWindowDayHoursModelWithDefaults() *BackupWindowDayHoursModel`

NewBackupWindowDayHoursModelWithDefaults instantiates a new BackupWindowDayHoursModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDay

`func (o *BackupWindowDayHoursModel) GetDay() EDayOfWeek`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *BackupWindowDayHoursModel) GetDayOk() (*EDayOfWeek, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *BackupWindowDayHoursModel) SetDay(v EDayOfWeek)`

SetDay sets Day field to given value.


### GetHours

`func (o *BackupWindowDayHoursModel) GetHours() string`

GetHours returns the Hours field if non-nil, zero value otherwise.

### GetHoursOk

`func (o *BackupWindowDayHoursModel) GetHoursOk() (*string, bool)`

GetHoursOk returns a tuple with the Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHours

`func (o *BackupWindowDayHoursModel) SetHours(v string)`

SetHours sets Hours field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


