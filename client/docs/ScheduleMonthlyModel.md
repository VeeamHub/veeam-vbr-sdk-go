# ScheduleMonthlyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, monthly schedule is enabled. | [default to false]
**LocalTime** | Pointer to **string** | Local time when the job must start. | [optional] 
**DayOfWeek** | Pointer to [**EDayOfWeek**](EDayOfWeek.md) |  | [optional] 
**DayNumberInMonth** | Pointer to [**EDayNumberInMonth**](EDayNumberInMonth.md) |  | [optional] 
**DayOfMonth** | Pointer to **int32** | Day of the month when the job must start. | [optional] 
**Months** | Pointer to [**[]EMonth**](EMonth.md) | Months when the job must start. | [optional] 

## Methods

### NewScheduleMonthlyModel

`func NewScheduleMonthlyModel(isEnabled bool, ) *ScheduleMonthlyModel`

NewScheduleMonthlyModel instantiates a new ScheduleMonthlyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleMonthlyModelWithDefaults

`func NewScheduleMonthlyModelWithDefaults() *ScheduleMonthlyModel`

NewScheduleMonthlyModelWithDefaults instantiates a new ScheduleMonthlyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *ScheduleMonthlyModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ScheduleMonthlyModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ScheduleMonthlyModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetLocalTime

`func (o *ScheduleMonthlyModel) GetLocalTime() string`

GetLocalTime returns the LocalTime field if non-nil, zero value otherwise.

### GetLocalTimeOk

`func (o *ScheduleMonthlyModel) GetLocalTimeOk() (*string, bool)`

GetLocalTimeOk returns a tuple with the LocalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalTime

`func (o *ScheduleMonthlyModel) SetLocalTime(v string)`

SetLocalTime sets LocalTime field to given value.

### HasLocalTime

`func (o *ScheduleMonthlyModel) HasLocalTime() bool`

HasLocalTime returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *ScheduleMonthlyModel) GetDayOfWeek() EDayOfWeek`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *ScheduleMonthlyModel) GetDayOfWeekOk() (*EDayOfWeek, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *ScheduleMonthlyModel) SetDayOfWeek(v EDayOfWeek)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *ScheduleMonthlyModel) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetDayNumberInMonth

`func (o *ScheduleMonthlyModel) GetDayNumberInMonth() EDayNumberInMonth`

GetDayNumberInMonth returns the DayNumberInMonth field if non-nil, zero value otherwise.

### GetDayNumberInMonthOk

`func (o *ScheduleMonthlyModel) GetDayNumberInMonthOk() (*EDayNumberInMonth, bool)`

GetDayNumberInMonthOk returns a tuple with the DayNumberInMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayNumberInMonth

`func (o *ScheduleMonthlyModel) SetDayNumberInMonth(v EDayNumberInMonth)`

SetDayNumberInMonth sets DayNumberInMonth field to given value.

### HasDayNumberInMonth

`func (o *ScheduleMonthlyModel) HasDayNumberInMonth() bool`

HasDayNumberInMonth returns a boolean if a field has been set.

### GetDayOfMonth

`func (o *ScheduleMonthlyModel) GetDayOfMonth() int32`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *ScheduleMonthlyModel) GetDayOfMonthOk() (*int32, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *ScheduleMonthlyModel) SetDayOfMonth(v int32)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *ScheduleMonthlyModel) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### GetMonths

`func (o *ScheduleMonthlyModel) GetMonths() []EMonth`

GetMonths returns the Months field if non-nil, zero value otherwise.

### GetMonthsOk

`func (o *ScheduleMonthlyModel) GetMonthsOk() (*[]EMonth, bool)`

GetMonthsOk returns a tuple with the Months field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonths

`func (o *ScheduleMonthlyModel) SetMonths(v []EMonth)`

SetMonths sets Months field to given value.

### HasMonths

`func (o *ScheduleMonthlyModel) HasMonths() bool`

HasMonths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


