# InstantViVMRecoveryMount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Mount ID. | 
**SessionId** | **string** | ID of the restore session. Use the ID to track the progress. For details, see [Get Session](#tag/Sessions/operation/GetSession). | 
**State** | [**EInstantRecoveryMountState**](EInstantRecoveryMountState.md) |  | 
**Spec** | [**InstantViVMRecoverySpec**](InstantViVMRecoverySpec.md) |  | 
**VmName** | **string** | Name of the recovered VM. | 
**ErrorMessage** | Pointer to **string** | Error message. | [optional] 

## Methods

### NewInstantViVMRecoveryMount

`func NewInstantViVMRecoveryMount(id string, sessionId string, state EInstantRecoveryMountState, spec InstantViVMRecoverySpec, vmName string, ) *InstantViVMRecoveryMount`

NewInstantViVMRecoveryMount instantiates a new InstantViVMRecoveryMount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstantViVMRecoveryMountWithDefaults

`func NewInstantViVMRecoveryMountWithDefaults() *InstantViVMRecoveryMount`

NewInstantViVMRecoveryMountWithDefaults instantiates a new InstantViVMRecoveryMount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstantViVMRecoveryMount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstantViVMRecoveryMount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstantViVMRecoveryMount) SetId(v string)`

SetId sets Id field to given value.


### GetSessionId

`func (o *InstantViVMRecoveryMount) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *InstantViVMRecoveryMount) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *InstantViVMRecoveryMount) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetState

`func (o *InstantViVMRecoveryMount) GetState() EInstantRecoveryMountState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InstantViVMRecoveryMount) GetStateOk() (*EInstantRecoveryMountState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *InstantViVMRecoveryMount) SetState(v EInstantRecoveryMountState)`

SetState sets State field to given value.


### GetSpec

`func (o *InstantViVMRecoveryMount) GetSpec() InstantViVMRecoverySpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *InstantViVMRecoveryMount) GetSpecOk() (*InstantViVMRecoverySpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *InstantViVMRecoveryMount) SetSpec(v InstantViVMRecoverySpec)`

SetSpec sets Spec field to given value.


### GetVmName

`func (o *InstantViVMRecoveryMount) GetVmName() string`

GetVmName returns the VmName field if non-nil, zero value otherwise.

### GetVmNameOk

`func (o *InstantViVMRecoveryMount) GetVmNameOk() (*string, bool)`

GetVmNameOk returns a tuple with the VmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmName

`func (o *InstantViVMRecoveryMount) SetVmName(v string)`

SetVmName sets VmName field to given value.


### GetErrorMessage

`func (o *InstantViVMRecoveryMount) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *InstantViVMRecoveryMount) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *InstantViVMRecoveryMount) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *InstantViVMRecoveryMount) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


