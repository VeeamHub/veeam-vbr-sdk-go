# PlacementPolicyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EPlacementPolicyType**](EPlacementPolicyType.md) |  | 
**Settings** | Pointer to [**[]BackupPlacementSettingsModel**](BackupPlacementSettingsModel.md) | Placement policy settings. | [optional] 
**EnforceStrictPlacementPolicy** | Pointer to **bool** | If *true*, the backup job fails in case the placement policy cannot be met. | [optional] 

## Methods

### NewPlacementPolicyModel

`func NewPlacementPolicyModel(type_ EPlacementPolicyType, ) *PlacementPolicyModel`

NewPlacementPolicyModel instantiates a new PlacementPolicyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlacementPolicyModelWithDefaults

`func NewPlacementPolicyModelWithDefaults() *PlacementPolicyModel`

NewPlacementPolicyModelWithDefaults instantiates a new PlacementPolicyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PlacementPolicyModel) GetType() EPlacementPolicyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlacementPolicyModel) GetTypeOk() (*EPlacementPolicyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlacementPolicyModel) SetType(v EPlacementPolicyType)`

SetType sets Type field to given value.


### GetSettings

`func (o *PlacementPolicyModel) GetSettings() []BackupPlacementSettingsModel`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *PlacementPolicyModel) GetSettingsOk() (*[]BackupPlacementSettingsModel, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *PlacementPolicyModel) SetSettings(v []BackupPlacementSettingsModel)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *PlacementPolicyModel) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetEnforceStrictPlacementPolicy

`func (o *PlacementPolicyModel) GetEnforceStrictPlacementPolicy() bool`

GetEnforceStrictPlacementPolicy returns the EnforceStrictPlacementPolicy field if non-nil, zero value otherwise.

### GetEnforceStrictPlacementPolicyOk

`func (o *PlacementPolicyModel) GetEnforceStrictPlacementPolicyOk() (*bool, bool)`

GetEnforceStrictPlacementPolicyOk returns a tuple with the EnforceStrictPlacementPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceStrictPlacementPolicy

`func (o *PlacementPolicyModel) SetEnforceStrictPlacementPolicy(v bool)`

SetEnforceStrictPlacementPolicy sets EnforceStrictPlacementPolicy field to given value.

### HasEnforceStrictPlacementPolicy

`func (o *PlacementPolicyModel) HasEnforceStrictPlacementPolicy() bool`

HasEnforceStrictPlacementPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


