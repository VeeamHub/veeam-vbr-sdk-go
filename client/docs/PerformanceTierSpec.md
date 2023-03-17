# PerformanceTierSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PerformanceExtents** | [**[]PerformanceExtentSpec**](PerformanceExtentSpec.md) | Array of performance extents. | 
**AdvancedSettings** | Pointer to [**PerformanceTierAdvancedSettingsModel**](PerformanceTierAdvancedSettingsModel.md) |  | [optional] 

## Methods

### NewPerformanceTierSpec

`func NewPerformanceTierSpec(performanceExtents []PerformanceExtentSpec, ) *PerformanceTierSpec`

NewPerformanceTierSpec instantiates a new PerformanceTierSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformanceTierSpecWithDefaults

`func NewPerformanceTierSpecWithDefaults() *PerformanceTierSpec`

NewPerformanceTierSpecWithDefaults instantiates a new PerformanceTierSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPerformanceExtents

`func (o *PerformanceTierSpec) GetPerformanceExtents() []PerformanceExtentSpec`

GetPerformanceExtents returns the PerformanceExtents field if non-nil, zero value otherwise.

### GetPerformanceExtentsOk

`func (o *PerformanceTierSpec) GetPerformanceExtentsOk() (*[]PerformanceExtentSpec, bool)`

GetPerformanceExtentsOk returns a tuple with the PerformanceExtents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceExtents

`func (o *PerformanceTierSpec) SetPerformanceExtents(v []PerformanceExtentSpec)`

SetPerformanceExtents sets PerformanceExtents field to given value.


### GetAdvancedSettings

`func (o *PerformanceTierSpec) GetAdvancedSettings() PerformanceTierAdvancedSettingsModel`

GetAdvancedSettings returns the AdvancedSettings field if non-nil, zero value otherwise.

### GetAdvancedSettingsOk

`func (o *PerformanceTierSpec) GetAdvancedSettingsOk() (*PerformanceTierAdvancedSettingsModel, bool)`

GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSettings

`func (o *PerformanceTierSpec) SetAdvancedSettings(v PerformanceTierAdvancedSettingsModel)`

SetAdvancedSettings sets AdvancedSettings field to given value.

### HasAdvancedSettings

`func (o *PerformanceTierSpec) HasAdvancedSettings() bool`

HasAdvancedSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


