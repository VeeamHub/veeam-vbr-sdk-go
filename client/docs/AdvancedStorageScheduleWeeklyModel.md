# AdvancedStorageScheduleWeeklyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, the weekly schedule is enabled. | [default to false]
**Days** | Pointer to [**[]EDayOfWeek**](EDayOfWeek.md) | Days of the week when the operation is performed. | [optional] 

## Methods

### NewAdvancedStorageScheduleWeeklyModel

`func NewAdvancedStorageScheduleWeeklyModel(isEnabled bool, ) *AdvancedStorageScheduleWeeklyModel`

NewAdvancedStorageScheduleWeeklyModel instantiates a new AdvancedStorageScheduleWeeklyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedStorageScheduleWeeklyModelWithDefaults

`func NewAdvancedStorageScheduleWeeklyModelWithDefaults() *AdvancedStorageScheduleWeeklyModel`

NewAdvancedStorageScheduleWeeklyModelWithDefaults instantiates a new AdvancedStorageScheduleWeeklyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *AdvancedStorageScheduleWeeklyModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *AdvancedStorageScheduleWeeklyModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *AdvancedStorageScheduleWeeklyModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetDays

`func (o *AdvancedStorageScheduleWeeklyModel) GetDays() []EDayOfWeek`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *AdvancedStorageScheduleWeeklyModel) GetDaysOk() (*[]EDayOfWeek, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *AdvancedStorageScheduleWeeklyModel) SetDays(v []EDayOfWeek)`

SetDays sets Days field to given value.

### HasDays

`func (o *AdvancedStorageScheduleWeeklyModel) HasDays() bool`

HasDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


