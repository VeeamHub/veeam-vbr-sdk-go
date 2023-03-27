# InstantViVMRecoverySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectRestorePointId** | **string** | ID of the restore point. | 
**Type** | [**EInstantVMRecoveryModeType**](EInstantVMRecoveryModeType.md) |  | 
**VmTagsRestoreEnabled** | Pointer to **bool** | If *true*, Veeam Backup &amp; Replication restores tags that were assigned to the original VM, and assign them to the restored VM. | [optional] 
**SecureRestore** | [**SecureRestoreSpec**](SecureRestoreSpec.md) |  | 
**NicsEnabled** | Pointer to **bool** | If *true*, the restored VM is connected to the network. | [optional] 
**PowerUp** | Pointer to **bool** | If *true*, Veeam Backup &amp; Replication starts the restored VM on the target host. | [optional] 
**Reason** | Pointer to **string** | Reason for restoring the VM. | [optional] 

## Methods

### NewInstantViVMRecoverySpec

`func NewInstantViVMRecoverySpec(objectRestorePointId string, type_ EInstantVMRecoveryModeType, secureRestore SecureRestoreSpec, ) *InstantViVMRecoverySpec`

NewInstantViVMRecoverySpec instantiates a new InstantViVMRecoverySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstantViVMRecoverySpecWithDefaults

`func NewInstantViVMRecoverySpecWithDefaults() *InstantViVMRecoverySpec`

NewInstantViVMRecoverySpecWithDefaults instantiates a new InstantViVMRecoverySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectRestorePointId

`func (o *InstantViVMRecoverySpec) GetObjectRestorePointId() string`

GetObjectRestorePointId returns the ObjectRestorePointId field if non-nil, zero value otherwise.

### GetObjectRestorePointIdOk

`func (o *InstantViVMRecoverySpec) GetObjectRestorePointIdOk() (*string, bool)`

GetObjectRestorePointIdOk returns a tuple with the ObjectRestorePointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectRestorePointId

`func (o *InstantViVMRecoverySpec) SetObjectRestorePointId(v string)`

SetObjectRestorePointId sets ObjectRestorePointId field to given value.


### GetType

`func (o *InstantViVMRecoverySpec) GetType() EInstantVMRecoveryModeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstantViVMRecoverySpec) GetTypeOk() (*EInstantVMRecoveryModeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstantViVMRecoverySpec) SetType(v EInstantVMRecoveryModeType)`

SetType sets Type field to given value.


### GetVmTagsRestoreEnabled

`func (o *InstantViVMRecoverySpec) GetVmTagsRestoreEnabled() bool`

GetVmTagsRestoreEnabled returns the VmTagsRestoreEnabled field if non-nil, zero value otherwise.

### GetVmTagsRestoreEnabledOk

`func (o *InstantViVMRecoverySpec) GetVmTagsRestoreEnabledOk() (*bool, bool)`

GetVmTagsRestoreEnabledOk returns a tuple with the VmTagsRestoreEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTagsRestoreEnabled

`func (o *InstantViVMRecoverySpec) SetVmTagsRestoreEnabled(v bool)`

SetVmTagsRestoreEnabled sets VmTagsRestoreEnabled field to given value.

### HasVmTagsRestoreEnabled

`func (o *InstantViVMRecoverySpec) HasVmTagsRestoreEnabled() bool`

HasVmTagsRestoreEnabled returns a boolean if a field has been set.

### GetSecureRestore

`func (o *InstantViVMRecoverySpec) GetSecureRestore() SecureRestoreSpec`

GetSecureRestore returns the SecureRestore field if non-nil, zero value otherwise.

### GetSecureRestoreOk

`func (o *InstantViVMRecoverySpec) GetSecureRestoreOk() (*SecureRestoreSpec, bool)`

GetSecureRestoreOk returns a tuple with the SecureRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureRestore

`func (o *InstantViVMRecoverySpec) SetSecureRestore(v SecureRestoreSpec)`

SetSecureRestore sets SecureRestore field to given value.


### GetNicsEnabled

`func (o *InstantViVMRecoverySpec) GetNicsEnabled() bool`

GetNicsEnabled returns the NicsEnabled field if non-nil, zero value otherwise.

### GetNicsEnabledOk

`func (o *InstantViVMRecoverySpec) GetNicsEnabledOk() (*bool, bool)`

GetNicsEnabledOk returns a tuple with the NicsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicsEnabled

`func (o *InstantViVMRecoverySpec) SetNicsEnabled(v bool)`

SetNicsEnabled sets NicsEnabled field to given value.

### HasNicsEnabled

`func (o *InstantViVMRecoverySpec) HasNicsEnabled() bool`

HasNicsEnabled returns a boolean if a field has been set.

### GetPowerUp

`func (o *InstantViVMRecoverySpec) GetPowerUp() bool`

GetPowerUp returns the PowerUp field if non-nil, zero value otherwise.

### GetPowerUpOk

`func (o *InstantViVMRecoverySpec) GetPowerUpOk() (*bool, bool)`

GetPowerUpOk returns a tuple with the PowerUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerUp

`func (o *InstantViVMRecoverySpec) SetPowerUp(v bool)`

SetPowerUp sets PowerUp field to given value.

### HasPowerUp

`func (o *InstantViVMRecoverySpec) HasPowerUp() bool`

HasPowerUp returns a boolean if a field has been set.

### GetReason

`func (o *InstantViVMRecoverySpec) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *InstantViVMRecoverySpec) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *InstantViVMRecoverySpec) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *InstantViVMRecoverySpec) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


