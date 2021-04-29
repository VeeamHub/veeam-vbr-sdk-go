# CapacityTierModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If *true*, the capacity tier is enabled. | [optional] 
**ExtentId** | Pointer to **string** | ID of an object storage repository added as a capacity extent. | [optional] 
**OffloadWindow** | Pointer to [**BackupWindowSettingModel**](BackupWindowSettingModel.md) |  | [optional] 
**CopyPolicyEnabled** | Pointer to **bool** | If *true*, Veeam Backup &amp; Replication copies backups from the performance extents to the capacity extent as soon as the backups are created. | [optional] 
**MovePolicyEnabled** | Pointer to **bool** | If *true*, Veeam Backup &amp; Replication moves backup files that belong to inactive backup chains from the performance extents to the capacity extent. | [optional] 
**OperationalRestorePeriodDays** | Pointer to **int32** | Number of days after which inactive backup chains on the performance extents are moved to the capacity extent. Specify *0* to offload inactive backup chains on the same day they are created. | [optional] 
**OverridePolicy** | Pointer to [**CapacityTierOverridePolicyModel**](CapacityTierOverridePolicyModel.md) |  | [optional] 
**Encryption** | Pointer to [**BackupStorageSettingsEncryptionModel**](BackupStorageSettingsEncryptionModel.md) |  | [optional] 

## Methods

### NewCapacityTierModel

`func NewCapacityTierModel() *CapacityTierModel`

NewCapacityTierModel instantiates a new CapacityTierModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapacityTierModelWithDefaults

`func NewCapacityTierModelWithDefaults() *CapacityTierModel`

NewCapacityTierModelWithDefaults instantiates a new CapacityTierModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *CapacityTierModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CapacityTierModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CapacityTierModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CapacityTierModel) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExtentId

`func (o *CapacityTierModel) GetExtentId() string`

GetExtentId returns the ExtentId field if non-nil, zero value otherwise.

### GetExtentIdOk

`func (o *CapacityTierModel) GetExtentIdOk() (*string, bool)`

GetExtentIdOk returns a tuple with the ExtentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtentId

`func (o *CapacityTierModel) SetExtentId(v string)`

SetExtentId sets ExtentId field to given value.

### HasExtentId

`func (o *CapacityTierModel) HasExtentId() bool`

HasExtentId returns a boolean if a field has been set.

### GetOffloadWindow

`func (o *CapacityTierModel) GetOffloadWindow() BackupWindowSettingModel`

GetOffloadWindow returns the OffloadWindow field if non-nil, zero value otherwise.

### GetOffloadWindowOk

`func (o *CapacityTierModel) GetOffloadWindowOk() (*BackupWindowSettingModel, bool)`

GetOffloadWindowOk returns a tuple with the OffloadWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffloadWindow

`func (o *CapacityTierModel) SetOffloadWindow(v BackupWindowSettingModel)`

SetOffloadWindow sets OffloadWindow field to given value.

### HasOffloadWindow

`func (o *CapacityTierModel) HasOffloadWindow() bool`

HasOffloadWindow returns a boolean if a field has been set.

### GetCopyPolicyEnabled

`func (o *CapacityTierModel) GetCopyPolicyEnabled() bool`

GetCopyPolicyEnabled returns the CopyPolicyEnabled field if non-nil, zero value otherwise.

### GetCopyPolicyEnabledOk

`func (o *CapacityTierModel) GetCopyPolicyEnabledOk() (*bool, bool)`

GetCopyPolicyEnabledOk returns a tuple with the CopyPolicyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyPolicyEnabled

`func (o *CapacityTierModel) SetCopyPolicyEnabled(v bool)`

SetCopyPolicyEnabled sets CopyPolicyEnabled field to given value.

### HasCopyPolicyEnabled

`func (o *CapacityTierModel) HasCopyPolicyEnabled() bool`

HasCopyPolicyEnabled returns a boolean if a field has been set.

### GetMovePolicyEnabled

`func (o *CapacityTierModel) GetMovePolicyEnabled() bool`

GetMovePolicyEnabled returns the MovePolicyEnabled field if non-nil, zero value otherwise.

### GetMovePolicyEnabledOk

`func (o *CapacityTierModel) GetMovePolicyEnabledOk() (*bool, bool)`

GetMovePolicyEnabledOk returns a tuple with the MovePolicyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovePolicyEnabled

`func (o *CapacityTierModel) SetMovePolicyEnabled(v bool)`

SetMovePolicyEnabled sets MovePolicyEnabled field to given value.

### HasMovePolicyEnabled

`func (o *CapacityTierModel) HasMovePolicyEnabled() bool`

HasMovePolicyEnabled returns a boolean if a field has been set.

### GetOperationalRestorePeriodDays

`func (o *CapacityTierModel) GetOperationalRestorePeriodDays() int32`

GetOperationalRestorePeriodDays returns the OperationalRestorePeriodDays field if non-nil, zero value otherwise.

### GetOperationalRestorePeriodDaysOk

`func (o *CapacityTierModel) GetOperationalRestorePeriodDaysOk() (*int32, bool)`

GetOperationalRestorePeriodDaysOk returns a tuple with the OperationalRestorePeriodDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalRestorePeriodDays

`func (o *CapacityTierModel) SetOperationalRestorePeriodDays(v int32)`

SetOperationalRestorePeriodDays sets OperationalRestorePeriodDays field to given value.

### HasOperationalRestorePeriodDays

`func (o *CapacityTierModel) HasOperationalRestorePeriodDays() bool`

HasOperationalRestorePeriodDays returns a boolean if a field has been set.

### GetOverridePolicy

`func (o *CapacityTierModel) GetOverridePolicy() CapacityTierOverridePolicyModel`

GetOverridePolicy returns the OverridePolicy field if non-nil, zero value otherwise.

### GetOverridePolicyOk

`func (o *CapacityTierModel) GetOverridePolicyOk() (*CapacityTierOverridePolicyModel, bool)`

GetOverridePolicyOk returns a tuple with the OverridePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridePolicy

`func (o *CapacityTierModel) SetOverridePolicy(v CapacityTierOverridePolicyModel)`

SetOverridePolicy sets OverridePolicy field to given value.

### HasOverridePolicy

`func (o *CapacityTierModel) HasOverridePolicy() bool`

HasOverridePolicy returns a boolean if a field has been set.

### GetEncryption

`func (o *CapacityTierModel) GetEncryption() BackupStorageSettingsEncryptionModel`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *CapacityTierModel) GetEncryptionOk() (*BackupStorageSettingsEncryptionModel, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *CapacityTierModel) SetEncryption(v BackupStorageSettingsEncryptionModel)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *CapacityTierModel) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


