# VmwareFcdInstantRecoveryMount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Mount ID. | 
**SessionId** | **string** | ID of the restore session. Use the ID to track the progress. For details, see [Get Session](#tag/Sessions/operation/GetSession). | 
**State** | [**EInstantRecoveryMountState**](EInstantRecoveryMountState.md) |  | 
**Spec** | [**VmwareFcdInstantRecoverySpec**](VmwareFcdInstantRecoverySpec.md) |  | 
**ErrorMessage** | Pointer to **string** | Error message. | [optional] 
**MountedDisks** | Pointer to [**[]VmwareFcdInstantRecoveryDiskInfo**](VmwareFcdInstantRecoveryDiskInfo.md) | Array of mounted disks. | [optional] 

## Methods

### NewVmwareFcdInstantRecoveryMount

`func NewVmwareFcdInstantRecoveryMount(id string, sessionId string, state EInstantRecoveryMountState, spec VmwareFcdInstantRecoverySpec, ) *VmwareFcdInstantRecoveryMount`

NewVmwareFcdInstantRecoveryMount instantiates a new VmwareFcdInstantRecoveryMount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareFcdInstantRecoveryMountWithDefaults

`func NewVmwareFcdInstantRecoveryMountWithDefaults() *VmwareFcdInstantRecoveryMount`

NewVmwareFcdInstantRecoveryMountWithDefaults instantiates a new VmwareFcdInstantRecoveryMount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VmwareFcdInstantRecoveryMount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VmwareFcdInstantRecoveryMount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VmwareFcdInstantRecoveryMount) SetId(v string)`

SetId sets Id field to given value.


### GetSessionId

`func (o *VmwareFcdInstantRecoveryMount) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *VmwareFcdInstantRecoveryMount) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *VmwareFcdInstantRecoveryMount) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetState

`func (o *VmwareFcdInstantRecoveryMount) GetState() EInstantRecoveryMountState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VmwareFcdInstantRecoveryMount) GetStateOk() (*EInstantRecoveryMountState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VmwareFcdInstantRecoveryMount) SetState(v EInstantRecoveryMountState)`

SetState sets State field to given value.


### GetSpec

`func (o *VmwareFcdInstantRecoveryMount) GetSpec() VmwareFcdInstantRecoverySpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VmwareFcdInstantRecoveryMount) GetSpecOk() (*VmwareFcdInstantRecoverySpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VmwareFcdInstantRecoveryMount) SetSpec(v VmwareFcdInstantRecoverySpec)`

SetSpec sets Spec field to given value.


### GetErrorMessage

`func (o *VmwareFcdInstantRecoveryMount) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *VmwareFcdInstantRecoveryMount) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *VmwareFcdInstantRecoveryMount) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *VmwareFcdInstantRecoveryMount) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetMountedDisks

`func (o *VmwareFcdInstantRecoveryMount) GetMountedDisks() []VmwareFcdInstantRecoveryDiskInfo`

GetMountedDisks returns the MountedDisks field if non-nil, zero value otherwise.

### GetMountedDisksOk

`func (o *VmwareFcdInstantRecoveryMount) GetMountedDisksOk() (*[]VmwareFcdInstantRecoveryDiskInfo, bool)`

GetMountedDisksOk returns a tuple with the MountedDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountedDisks

`func (o *VmwareFcdInstantRecoveryMount) SetMountedDisks(v []VmwareFcdInstantRecoveryDiskInfo)`

SetMountedDisks sets MountedDisks field to given value.

### HasMountedDisks

`func (o *VmwareFcdInstantRecoveryMount) HasMountedDisks() bool`

HasMountedDisks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


