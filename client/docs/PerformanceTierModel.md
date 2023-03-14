# PerformanceTierModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PerformanceExtents** | [**[]PerformanceExtentModel**](PerformanceExtentModel.md) | Array of performance extents. | 
**AdvancedSettings** | Pointer to [**PerformanceTierAdvancedSettingsModel**](PerformanceTierAdvancedSettingsModel.md) |  | [optional] 

## Methods

### NewPerformanceTierModel

`func NewPerformanceTierModel(performanceExtents []PerformanceExtentModel, ) *PerformanceTierModel`

NewPerformanceTierModel instantiates a new PerformanceTierModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformanceTierModelWithDefaults

`func NewPerformanceTierModelWithDefaults() *PerformanceTierModel`

NewPerformanceTierModelWithDefaults instantiates a new PerformanceTierModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPerformanceExtents

`func (o *PerformanceTierModel) GetPerformanceExtents() []PerformanceExtentModel`

GetPerformanceExtents returns the PerformanceExtents field if non-nil, zero value otherwise.

### GetPerformanceExtentsOk

`func (o *PerformanceTierModel) GetPerformanceExtentsOk() (*[]PerformanceExtentModel, bool)`

GetPerformanceExtentsOk returns a tuple with the PerformanceExtents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceExtents

`func (o *PerformanceTierModel) SetPerformanceExtents(v []PerformanceExtentModel)`

SetPerformanceExtents sets PerformanceExtents field to given value.


### GetAdvancedSettings

`func (o *PerformanceTierModel) GetAdvancedSettings() PerformanceTierAdvancedSettingsModel`

GetAdvancedSettings returns the AdvancedSettings field if non-nil, zero value otherwise.

### GetAdvancedSettingsOk

`func (o *PerformanceTierModel) GetAdvancedSettingsOk() (*PerformanceTierAdvancedSettingsModel, bool)`

GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSettings

`func (o *PerformanceTierModel) SetAdvancedSettings(v PerformanceTierAdvancedSettingsModel)`

SetAdvancedSettings sets AdvancedSettings field to given value.

### HasAdvancedSettings

`func (o *PerformanceTierModel) HasAdvancedSettings() bool`

HasAdvancedSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


