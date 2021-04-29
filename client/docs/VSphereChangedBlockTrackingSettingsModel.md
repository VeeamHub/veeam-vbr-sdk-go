# VSphereChangedBlockTrackingSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, CBT data is used. | 
**EnableCBTautomatically** | Pointer to **bool** | If *true*, CBT is enabled for all processed VMs even if CBT is disabled in VM configuration. CBT is used for VMs with virtual hardware version 7 or later. These VMs must not have existing snapshots. | [optional] 
**ResetCBTonActiveFull** | Pointer to **bool** | If *true*, CBT is reset before creating active full backups. | [optional] 

## Methods

### NewVSphereChangedBlockTrackingSettingsModel

`func NewVSphereChangedBlockTrackingSettingsModel(isEnabled bool, ) *VSphereChangedBlockTrackingSettingsModel`

NewVSphereChangedBlockTrackingSettingsModel instantiates a new VSphereChangedBlockTrackingSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVSphereChangedBlockTrackingSettingsModelWithDefaults

`func NewVSphereChangedBlockTrackingSettingsModelWithDefaults() *VSphereChangedBlockTrackingSettingsModel`

NewVSphereChangedBlockTrackingSettingsModelWithDefaults instantiates a new VSphereChangedBlockTrackingSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *VSphereChangedBlockTrackingSettingsModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *VSphereChangedBlockTrackingSettingsModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *VSphereChangedBlockTrackingSettingsModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetEnableCBTautomatically

`func (o *VSphereChangedBlockTrackingSettingsModel) GetEnableCBTautomatically() bool`

GetEnableCBTautomatically returns the EnableCBTautomatically field if non-nil, zero value otherwise.

### GetEnableCBTautomaticallyOk

`func (o *VSphereChangedBlockTrackingSettingsModel) GetEnableCBTautomaticallyOk() (*bool, bool)`

GetEnableCBTautomaticallyOk returns a tuple with the EnableCBTautomatically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCBTautomatically

`func (o *VSphereChangedBlockTrackingSettingsModel) SetEnableCBTautomatically(v bool)`

SetEnableCBTautomatically sets EnableCBTautomatically field to given value.

### HasEnableCBTautomatically

`func (o *VSphereChangedBlockTrackingSettingsModel) HasEnableCBTautomatically() bool`

HasEnableCBTautomatically returns a boolean if a field has been set.

### GetResetCBTonActiveFull

`func (o *VSphereChangedBlockTrackingSettingsModel) GetResetCBTonActiveFull() bool`

GetResetCBTonActiveFull returns the ResetCBTonActiveFull field if non-nil, zero value otherwise.

### GetResetCBTonActiveFullOk

`func (o *VSphereChangedBlockTrackingSettingsModel) GetResetCBTonActiveFullOk() (*bool, bool)`

GetResetCBTonActiveFullOk returns a tuple with the ResetCBTonActiveFull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetCBTonActiveFull

`func (o *VSphereChangedBlockTrackingSettingsModel) SetResetCBTonActiveFull(v bool)`

SetResetCBTonActiveFull sets ResetCBTonActiveFull field to given value.

### HasResetCBTonActiveFull

`func (o *VSphereChangedBlockTrackingSettingsModel) HasResetCBTonActiveFull() bool`

HasResetCBTonActiveFull returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


