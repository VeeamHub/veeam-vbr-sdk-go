# ScaleOutRepositorySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the scale-out backup repository. | 
**Description** | **string** | Description of the scale-out backup repository. | 
**Tag** | Pointer to **string** | Tag assigned to the scale-out backup repository. | [optional] 
**PerformanceTier** | [**PerformanceTierSpec**](PerformanceTierSpec.md) |  | 
**PlacementPolicy** | Pointer to [**PlacementPolicyModel**](PlacementPolicyModel.md) |  | [optional] 
**CapacityTier** | Pointer to [**CapacityTierModel**](CapacityTierModel.md) |  | [optional] 
**ArchiveTier** | Pointer to [**ArchiveTierModel**](ArchiveTierModel.md) |  | [optional] 

## Methods

### NewScaleOutRepositorySpec

`func NewScaleOutRepositorySpec(name string, description string, performanceTier PerformanceTierSpec, ) *ScaleOutRepositorySpec`

NewScaleOutRepositorySpec instantiates a new ScaleOutRepositorySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScaleOutRepositorySpecWithDefaults

`func NewScaleOutRepositorySpecWithDefaults() *ScaleOutRepositorySpec`

NewScaleOutRepositorySpecWithDefaults instantiates a new ScaleOutRepositorySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ScaleOutRepositorySpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScaleOutRepositorySpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScaleOutRepositorySpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ScaleOutRepositorySpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScaleOutRepositorySpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScaleOutRepositorySpec) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTag

`func (o *ScaleOutRepositorySpec) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ScaleOutRepositorySpec) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ScaleOutRepositorySpec) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ScaleOutRepositorySpec) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetPerformanceTier

`func (o *ScaleOutRepositorySpec) GetPerformanceTier() PerformanceTierSpec`

GetPerformanceTier returns the PerformanceTier field if non-nil, zero value otherwise.

### GetPerformanceTierOk

`func (o *ScaleOutRepositorySpec) GetPerformanceTierOk() (*PerformanceTierSpec, bool)`

GetPerformanceTierOk returns a tuple with the PerformanceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceTier

`func (o *ScaleOutRepositorySpec) SetPerformanceTier(v PerformanceTierSpec)`

SetPerformanceTier sets PerformanceTier field to given value.


### GetPlacementPolicy

`func (o *ScaleOutRepositorySpec) GetPlacementPolicy() PlacementPolicyModel`

GetPlacementPolicy returns the PlacementPolicy field if non-nil, zero value otherwise.

### GetPlacementPolicyOk

`func (o *ScaleOutRepositorySpec) GetPlacementPolicyOk() (*PlacementPolicyModel, bool)`

GetPlacementPolicyOk returns a tuple with the PlacementPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementPolicy

`func (o *ScaleOutRepositorySpec) SetPlacementPolicy(v PlacementPolicyModel)`

SetPlacementPolicy sets PlacementPolicy field to given value.

### HasPlacementPolicy

`func (o *ScaleOutRepositorySpec) HasPlacementPolicy() bool`

HasPlacementPolicy returns a boolean if a field has been set.

### GetCapacityTier

`func (o *ScaleOutRepositorySpec) GetCapacityTier() CapacityTierModel`

GetCapacityTier returns the CapacityTier field if non-nil, zero value otherwise.

### GetCapacityTierOk

`func (o *ScaleOutRepositorySpec) GetCapacityTierOk() (*CapacityTierModel, bool)`

GetCapacityTierOk returns a tuple with the CapacityTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityTier

`func (o *ScaleOutRepositorySpec) SetCapacityTier(v CapacityTierModel)`

SetCapacityTier sets CapacityTier field to given value.

### HasCapacityTier

`func (o *ScaleOutRepositorySpec) HasCapacityTier() bool`

HasCapacityTier returns a boolean if a field has been set.

### GetArchiveTier

`func (o *ScaleOutRepositorySpec) GetArchiveTier() ArchiveTierModel`

GetArchiveTier returns the ArchiveTier field if non-nil, zero value otherwise.

### GetArchiveTierOk

`func (o *ScaleOutRepositorySpec) GetArchiveTierOk() (*ArchiveTierModel, bool)`

GetArchiveTierOk returns a tuple with the ArchiveTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveTier

`func (o *ScaleOutRepositorySpec) SetArchiveTier(v ArchiveTierModel)`

SetArchiveTier sets ArchiveTier field to given value.

### HasArchiveTier

`func (o *ScaleOutRepositorySpec) HasArchiveTier() bool`

HasArchiveTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


