# EntireViVMRestoreSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectRestorePointId** | **string** | ID of the restore point. | 
**Type** | [**EEntireVMRestoreModeType**](EEntireVMRestoreModeType.md) |  | 
**RestoreProxies** | Pointer to [**RestoreProxySpec**](RestoreProxySpec.md) |  | [optional] 
**SecureRestore** | Pointer to [**SecureRestoreSpec**](SecureRestoreSpec.md) |  | [optional] 
**PowerUp** | Pointer to **bool** | If *true*, Veeam Backup &amp; Replication starts the restored VM on the target host. | [optional] 
**Reason** | Pointer to **string** | Reason for restoring the VM. | [optional] 

## Methods

### NewEntireViVMRestoreSpec

`func NewEntireViVMRestoreSpec(objectRestorePointId string, type_ EEntireVMRestoreModeType, ) *EntireViVMRestoreSpec`

NewEntireViVMRestoreSpec instantiates a new EntireViVMRestoreSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntireViVMRestoreSpecWithDefaults

`func NewEntireViVMRestoreSpecWithDefaults() *EntireViVMRestoreSpec`

NewEntireViVMRestoreSpecWithDefaults instantiates a new EntireViVMRestoreSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectRestorePointId

`func (o *EntireViVMRestoreSpec) GetObjectRestorePointId() string`

GetObjectRestorePointId returns the ObjectRestorePointId field if non-nil, zero value otherwise.

### GetObjectRestorePointIdOk

`func (o *EntireViVMRestoreSpec) GetObjectRestorePointIdOk() (*string, bool)`

GetObjectRestorePointIdOk returns a tuple with the ObjectRestorePointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectRestorePointId

`func (o *EntireViVMRestoreSpec) SetObjectRestorePointId(v string)`

SetObjectRestorePointId sets ObjectRestorePointId field to given value.


### GetType

`func (o *EntireViVMRestoreSpec) GetType() EEntireVMRestoreModeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntireViVMRestoreSpec) GetTypeOk() (*EEntireVMRestoreModeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntireViVMRestoreSpec) SetType(v EEntireVMRestoreModeType)`

SetType sets Type field to given value.


### GetRestoreProxies

`func (o *EntireViVMRestoreSpec) GetRestoreProxies() RestoreProxySpec`

GetRestoreProxies returns the RestoreProxies field if non-nil, zero value otherwise.

### GetRestoreProxiesOk

`func (o *EntireViVMRestoreSpec) GetRestoreProxiesOk() (*RestoreProxySpec, bool)`

GetRestoreProxiesOk returns a tuple with the RestoreProxies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreProxies

`func (o *EntireViVMRestoreSpec) SetRestoreProxies(v RestoreProxySpec)`

SetRestoreProxies sets RestoreProxies field to given value.

### HasRestoreProxies

`func (o *EntireViVMRestoreSpec) HasRestoreProxies() bool`

HasRestoreProxies returns a boolean if a field has been set.

### GetSecureRestore

`func (o *EntireViVMRestoreSpec) GetSecureRestore() SecureRestoreSpec`

GetSecureRestore returns the SecureRestore field if non-nil, zero value otherwise.

### GetSecureRestoreOk

`func (o *EntireViVMRestoreSpec) GetSecureRestoreOk() (*SecureRestoreSpec, bool)`

GetSecureRestoreOk returns a tuple with the SecureRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureRestore

`func (o *EntireViVMRestoreSpec) SetSecureRestore(v SecureRestoreSpec)`

SetSecureRestore sets SecureRestore field to given value.

### HasSecureRestore

`func (o *EntireViVMRestoreSpec) HasSecureRestore() bool`

HasSecureRestore returns a boolean if a field has been set.

### GetPowerUp

`func (o *EntireViVMRestoreSpec) GetPowerUp() bool`

GetPowerUp returns the PowerUp field if non-nil, zero value otherwise.

### GetPowerUpOk

`func (o *EntireViVMRestoreSpec) GetPowerUpOk() (*bool, bool)`

GetPowerUpOk returns a tuple with the PowerUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerUp

`func (o *EntireViVMRestoreSpec) SetPowerUp(v bool)`

SetPowerUp sets PowerUp field to given value.

### HasPowerUp

`func (o *EntireViVMRestoreSpec) HasPowerUp() bool`

HasPowerUp returns a boolean if a field has been set.

### GetReason

`func (o *EntireViVMRestoreSpec) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EntireViVMRestoreSpec) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EntireViVMRestoreSpec) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *EntireViVMRestoreSpec) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


