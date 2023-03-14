# ArchiveTierAdvancedSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostOptimizedArchiveEnabled** | Pointer to **bool** | If *true*, backups are archived as soon as the remaining retention time is above minimum storage period for the repository. | [optional] 
**ArchiveDeduplicationEnabled** | Pointer to **bool** | If *true*, each backup is stored as a delta to the previous one. | [optional] 

## Methods

### NewArchiveTierAdvancedSettingsModel

`func NewArchiveTierAdvancedSettingsModel() *ArchiveTierAdvancedSettingsModel`

NewArchiveTierAdvancedSettingsModel instantiates a new ArchiveTierAdvancedSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchiveTierAdvancedSettingsModelWithDefaults

`func NewArchiveTierAdvancedSettingsModelWithDefaults() *ArchiveTierAdvancedSettingsModel`

NewArchiveTierAdvancedSettingsModelWithDefaults instantiates a new ArchiveTierAdvancedSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostOptimizedArchiveEnabled

`func (o *ArchiveTierAdvancedSettingsModel) GetCostOptimizedArchiveEnabled() bool`

GetCostOptimizedArchiveEnabled returns the CostOptimizedArchiveEnabled field if non-nil, zero value otherwise.

### GetCostOptimizedArchiveEnabledOk

`func (o *ArchiveTierAdvancedSettingsModel) GetCostOptimizedArchiveEnabledOk() (*bool, bool)`

GetCostOptimizedArchiveEnabledOk returns a tuple with the CostOptimizedArchiveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostOptimizedArchiveEnabled

`func (o *ArchiveTierAdvancedSettingsModel) SetCostOptimizedArchiveEnabled(v bool)`

SetCostOptimizedArchiveEnabled sets CostOptimizedArchiveEnabled field to given value.

### HasCostOptimizedArchiveEnabled

`func (o *ArchiveTierAdvancedSettingsModel) HasCostOptimizedArchiveEnabled() bool`

HasCostOptimizedArchiveEnabled returns a boolean if a field has been set.

### GetArchiveDeduplicationEnabled

`func (o *ArchiveTierAdvancedSettingsModel) GetArchiveDeduplicationEnabled() bool`

GetArchiveDeduplicationEnabled returns the ArchiveDeduplicationEnabled field if non-nil, zero value otherwise.

### GetArchiveDeduplicationEnabledOk

`func (o *ArchiveTierAdvancedSettingsModel) GetArchiveDeduplicationEnabledOk() (*bool, bool)`

GetArchiveDeduplicationEnabledOk returns a tuple with the ArchiveDeduplicationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveDeduplicationEnabled

`func (o *ArchiveTierAdvancedSettingsModel) SetArchiveDeduplicationEnabled(v bool)`

SetArchiveDeduplicationEnabled sets ArchiveDeduplicationEnabled field to given value.

### HasArchiveDeduplicationEnabled

`func (o *ArchiveTierAdvancedSettingsModel) HasArchiveDeduplicationEnabled() bool`

HasArchiveDeduplicationEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


