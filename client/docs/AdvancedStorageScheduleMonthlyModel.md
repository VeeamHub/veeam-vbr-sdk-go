# AdvancedStorageScheduleMonthlyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, the monthly schedule is enabled. | 
**DayOfWeek** | Pointer to [**EDayOfWeek**](EDayOfWeek.md) |  | [optional] 
**DayNumberInMonth** | Pointer to [**EDayNumberInMonth**](EDayNumberInMonth.md) |  | [optional] 
**DayOfMonths** | Pointer to **int32** | Day of the month when the operation is performed. | [optional] 
**Months** | Pointer to [**[]EMonth**](EMonth.md) | Months when the operation is performed. | [optional] 

## Methods

### NewAdvancedStorageScheduleMonthlyModel

`func NewAdvancedStorageScheduleMonthlyModel(isEnabled bool, ) *AdvancedStorageScheduleMonthlyModel`

NewAdvancedStorageScheduleMonthlyModel instantiates a new AdvancedStorageScheduleMonthlyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedStorageScheduleMonthlyModelWithDefaults

`func NewAdvancedStorageScheduleMonthlyModelWithDefaults() *AdvancedStorageScheduleMonthlyModel`

NewAdvancedStorageScheduleMonthlyModelWithDefaults instantiates a new AdvancedStorageScheduleMonthlyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *AdvancedStorageScheduleMonthlyModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *AdvancedStorageScheduleMonthlyModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *AdvancedStorageScheduleMonthlyModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetDayOfWeek

`func (o *AdvancedStorageScheduleMonthlyModel) GetDayOfWeek() EDayOfWeek`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *AdvancedStorageScheduleMonthlyModel) GetDayOfWeekOk() (*EDayOfWeek, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *AdvancedStorageScheduleMonthlyModel) SetDayOfWeek(v EDayOfWeek)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *AdvancedStorageScheduleMonthlyModel) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetDayNumberInMonth

`func (o *AdvancedStorageScheduleMonthlyModel) GetDayNumberInMonth() EDayNumberInMonth`

GetDayNumberInMonth returns the DayNumberInMonth field if non-nil, zero value otherwise.

### GetDayNumberInMonthOk

`func (o *AdvancedStorageScheduleMonthlyModel) GetDayNumberInMonthOk() (*EDayNumberInMonth, bool)`

GetDayNumberInMonthOk returns a tuple with the DayNumberInMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayNumberInMonth

`func (o *AdvancedStorageScheduleMonthlyModel) SetDayNumberInMonth(v EDayNumberInMonth)`

SetDayNumberInMonth sets DayNumberInMonth field to given value.

### HasDayNumberInMonth

`func (o *AdvancedStorageScheduleMonthlyModel) HasDayNumberInMonth() bool`

HasDayNumberInMonth returns a boolean if a field has been set.

### GetDayOfMonths

`func (o *AdvancedStorageScheduleMonthlyModel) GetDayOfMonths() int32`

GetDayOfMonths returns the DayOfMonths field if non-nil, zero value otherwise.

### GetDayOfMonthsOk

`func (o *AdvancedStorageScheduleMonthlyModel) GetDayOfMonthsOk() (*int32, bool)`

GetDayOfMonthsOk returns a tuple with the DayOfMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonths

`func (o *AdvancedStorageScheduleMonthlyModel) SetDayOfMonths(v int32)`

SetDayOfMonths sets DayOfMonths field to given value.

### HasDayOfMonths

`func (o *AdvancedStorageScheduleMonthlyModel) HasDayOfMonths() bool`

HasDayOfMonths returns a boolean if a field has been set.

### GetMonths

`func (o *AdvancedStorageScheduleMonthlyModel) GetMonths() []EMonth`

GetMonths returns the Months field if non-nil, zero value otherwise.

### GetMonthsOk

`func (o *AdvancedStorageScheduleMonthlyModel) GetMonthsOk() (*[]EMonth, bool)`

GetMonthsOk returns a tuple with the Months field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonths

`func (o *AdvancedStorageScheduleMonthlyModel) SetMonths(v []EMonth)`

SetMonths sets Months field to given value.

### HasMonths

`func (o *AdvancedStorageScheduleMonthlyModel) HasMonths() bool`

HasMonths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


