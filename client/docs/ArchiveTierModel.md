# ArchiveTierModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, the archive tier is enabled. | 
**ExtentId** | Pointer to **string** | ID of an object storage repository added as an archive extent. | [optional] 
**ArchivePeriodDays** | Pointer to **int32** | Number of days after which backup chains on the capacity extent are moved to the archive extent. Specify *0* to offload inactive backup chains on the same day they are created. | [optional] 
**AdvancedSettings** | Pointer to [**ArchiveTierAdvancedSettingsModel**](ArchiveTierAdvancedSettingsModel.md) |  | [optional] 

## Methods

### NewArchiveTierModel

`func NewArchiveTierModel(isEnabled bool, ) *ArchiveTierModel`

NewArchiveTierModel instantiates a new ArchiveTierModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchiveTierModelWithDefaults

`func NewArchiveTierModelWithDefaults() *ArchiveTierModel`

NewArchiveTierModelWithDefaults instantiates a new ArchiveTierModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *ArchiveTierModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ArchiveTierModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ArchiveTierModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetExtentId

`func (o *ArchiveTierModel) GetExtentId() string`

GetExtentId returns the ExtentId field if non-nil, zero value otherwise.

### GetExtentIdOk

`func (o *ArchiveTierModel) GetExtentIdOk() (*string, bool)`

GetExtentIdOk returns a tuple with the ExtentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtentId

`func (o *ArchiveTierModel) SetExtentId(v string)`

SetExtentId sets ExtentId field to given value.

### HasExtentId

`func (o *ArchiveTierModel) HasExtentId() bool`

HasExtentId returns a boolean if a field has been set.

### GetArchivePeriodDays

`func (o *ArchiveTierModel) GetArchivePeriodDays() int32`

GetArchivePeriodDays returns the ArchivePeriodDays field if non-nil, zero value otherwise.

### GetArchivePeriodDaysOk

`func (o *ArchiveTierModel) GetArchivePeriodDaysOk() (*int32, bool)`

GetArchivePeriodDaysOk returns a tuple with the ArchivePeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivePeriodDays

`func (o *ArchiveTierModel) SetArchivePeriodDays(v int32)`

SetArchivePeriodDays sets ArchivePeriodDays field to given value.

### HasArchivePeriodDays

`func (o *ArchiveTierModel) HasArchivePeriodDays() bool`

HasArchivePeriodDays returns a boolean if a field has been set.

### GetAdvancedSettings

`func (o *ArchiveTierModel) GetAdvancedSettings() ArchiveTierAdvancedSettingsModel`

GetAdvancedSettings returns the AdvancedSettings field if non-nil, zero value otherwise.

### GetAdvancedSettingsOk

`func (o *ArchiveTierModel) GetAdvancedSettingsOk() (*ArchiveTierAdvancedSettingsModel, bool)`

GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSettings

`func (o *ArchiveTierModel) SetAdvancedSettings(v ArchiveTierAdvancedSettingsModel)`

SetAdvancedSettings sets AdvancedSettings field to given value.

### HasAdvancedSettings

`func (o *ArchiveTierModel) HasAdvancedSettings() bool`

HasAdvancedSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


