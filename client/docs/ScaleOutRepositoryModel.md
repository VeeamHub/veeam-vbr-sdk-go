# ScaleOutRepositoryModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the scale-out backup repository. | 
**Name** | **string** | Name of the scale-out backup repository. | 
**Description** | **string** | Description of the scale-out backup repository. | 
**Tag** | Pointer to **string** | Tag assigned to the scale-out backup repository. | [optional] 
**PerformanceTier** | [**PerformanceTierModel**](PerformanceTierModel.md) |  | 
**PlacementPolicy** | Pointer to [**PlacementPolicyModel**](PlacementPolicyModel.md) |  | [optional] 
**CapacityTier** | Pointer to [**CapacityTierModel**](CapacityTierModel.md) |  | [optional] 
**ArchiveTier** | Pointer to [**ArchiveTierModel**](ArchiveTierModel.md) |  | [optional] 

## Methods

### NewScaleOutRepositoryModel

`func NewScaleOutRepositoryModel(id string, name string, description string, performanceTier PerformanceTierModel, ) *ScaleOutRepositoryModel`

NewScaleOutRepositoryModel instantiates a new ScaleOutRepositoryModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScaleOutRepositoryModelWithDefaults

`func NewScaleOutRepositoryModelWithDefaults() *ScaleOutRepositoryModel`

NewScaleOutRepositoryModelWithDefaults instantiates a new ScaleOutRepositoryModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScaleOutRepositoryModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScaleOutRepositoryModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScaleOutRepositoryModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ScaleOutRepositoryModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScaleOutRepositoryModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScaleOutRepositoryModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ScaleOutRepositoryModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScaleOutRepositoryModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScaleOutRepositoryModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTag

`func (o *ScaleOutRepositoryModel) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ScaleOutRepositoryModel) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ScaleOutRepositoryModel) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ScaleOutRepositoryModel) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetPerformanceTier

`func (o *ScaleOutRepositoryModel) GetPerformanceTier() PerformanceTierModel`

GetPerformanceTier returns the PerformanceTier field if non-nil, zero value otherwise.

### GetPerformanceTierOk

`func (o *ScaleOutRepositoryModel) GetPerformanceTierOk() (*PerformanceTierModel, bool)`

GetPerformanceTierOk returns a tuple with the PerformanceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceTier

`func (o *ScaleOutRepositoryModel) SetPerformanceTier(v PerformanceTierModel)`

SetPerformanceTier sets PerformanceTier field to given value.


### GetPlacementPolicy

`func (o *ScaleOutRepositoryModel) GetPlacementPolicy() PlacementPolicyModel`

GetPlacementPolicy returns the PlacementPolicy field if non-nil, zero value otherwise.

### GetPlacementPolicyOk

`func (o *ScaleOutRepositoryModel) GetPlacementPolicyOk() (*PlacementPolicyModel, bool)`

GetPlacementPolicyOk returns a tuple with the PlacementPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementPolicy

`func (o *ScaleOutRepositoryModel) SetPlacementPolicy(v PlacementPolicyModel)`

SetPlacementPolicy sets PlacementPolicy field to given value.

### HasPlacementPolicy

`func (o *ScaleOutRepositoryModel) HasPlacementPolicy() bool`

HasPlacementPolicy returns a boolean if a field has been set.

### GetCapacityTier

`func (o *ScaleOutRepositoryModel) GetCapacityTier() CapacityTierModel`

GetCapacityTier returns the CapacityTier field if non-nil, zero value otherwise.

### GetCapacityTierOk

`func (o *ScaleOutRepositoryModel) GetCapacityTierOk() (*CapacityTierModel, bool)`

GetCapacityTierOk returns a tuple with the CapacityTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityTier

`func (o *ScaleOutRepositoryModel) SetCapacityTier(v CapacityTierModel)`

SetCapacityTier sets CapacityTier field to given value.

### HasCapacityTier

`func (o *ScaleOutRepositoryModel) HasCapacityTier() bool`

HasCapacityTier returns a boolean if a field has been set.

### GetArchiveTier

`func (o *ScaleOutRepositoryModel) GetArchiveTier() ArchiveTierModel`

GetArchiveTier returns the ArchiveTier field if non-nil, zero value otherwise.

### GetArchiveTierOk

`func (o *ScaleOutRepositoryModel) GetArchiveTierOk() (*ArchiveTierModel, bool)`

GetArchiveTierOk returns a tuple with the ArchiveTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveTier

`func (o *ScaleOutRepositoryModel) SetArchiveTier(v ArchiveTierModel)`

SetArchiveTier sets ArchiveTier field to given value.

### HasArchiveTier

`func (o *ScaleOutRepositoryModel) HasArchiveTier() bool`

HasArchiveTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


