# SyntheticFullSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, synthetic full backups are enabled. | 
**Days** | Pointer to [**[]EDayOfWeek**](EDayOfWeek.md) | Days of the week when Veeam Backup &amp; Replication creates a synthetic full backup. | [optional] 

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


### GetDays

`func (o *SyntheticFullSettingsModel) GetDays() []EDayOfWeek`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *SyntheticFullSettingsModel) GetDaysOk() (*[]EDayOfWeek, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *SyntheticFullSettingsModel) SetDays(v []EDayOfWeek)`

SetDays sets Days field to given value.

### HasDays

`func (o *SyntheticFullSettingsModel) HasDays() bool`

HasDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


