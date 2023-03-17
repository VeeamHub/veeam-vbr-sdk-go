# RestoreTargetNetworkSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**VmwareObjectModel**](VmwareObjectModel.md) |  | 
**Disconnected** | Pointer to **bool** | If *true*, the restored VMs is not connected to any virtual network. | [optional] 

## Methods

### NewRestoreTargetNetworkSpec

`func NewRestoreTargetNetworkSpec(network VmwareObjectModel, ) *RestoreTargetNetworkSpec`

NewRestoreTargetNetworkSpec instantiates a new RestoreTargetNetworkSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreTargetNetworkSpecWithDefaults

`func NewRestoreTargetNetworkSpecWithDefaults() *RestoreTargetNetworkSpec`

NewRestoreTargetNetworkSpecWithDefaults instantiates a new RestoreTargetNetworkSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *RestoreTargetNetworkSpec) GetNetwork() VmwareObjectModel`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *RestoreTargetNetworkSpec) GetNetworkOk() (*VmwareObjectModel, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *RestoreTargetNetworkSpec) SetNetwork(v VmwareObjectModel)`

SetNetwork sets Network field to given value.


### GetDisconnected

`func (o *RestoreTargetNetworkSpec) GetDisconnected() bool`

GetDisconnected returns the Disconnected field if non-nil, zero value otherwise.

### GetDisconnectedOk

`func (o *RestoreTargetNetworkSpec) GetDisconnectedOk() (*bool, bool)`

GetDisconnectedOk returns a tuple with the Disconnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnected

`func (o *RestoreTargetNetworkSpec) SetDisconnected(v bool)`

SetDisconnected sets Disconnected field to given value.

### HasDisconnected

`func (o *RestoreTargetNetworkSpec) HasDisconnected() bool`

HasDisconnected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


