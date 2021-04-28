# ProxyDatastoreModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastore** | Pointer to [**VmwareObjectModel**](VmwareObjectModel.md) |  | [optional] 
**VmCount** | Pointer to **int32** | Number of VMs. | [optional] 

## Methods

### NewProxyDatastoreModel

`func NewProxyDatastoreModel() *ProxyDatastoreModel`

NewProxyDatastoreModel instantiates a new ProxyDatastoreModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyDatastoreModelWithDefaults

`func NewProxyDatastoreModelWithDefaults() *ProxyDatastoreModel`

NewProxyDatastoreModelWithDefaults instantiates a new ProxyDatastoreModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastore

`func (o *ProxyDatastoreModel) GetDatastore() VmwareObjectModel`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *ProxyDatastoreModel) GetDatastoreOk() (*VmwareObjectModel, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *ProxyDatastoreModel) SetDatastore(v VmwareObjectModel)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *ProxyDatastoreModel) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetVmCount

`func (o *ProxyDatastoreModel) GetVmCount() int32`

GetVmCount returns the VmCount field if non-nil, zero value otherwise.

### GetVmCountOk

`func (o *ProxyDatastoreModel) GetVmCountOk() (*int32, bool)`

GetVmCountOk returns a tuple with the VmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCount

`func (o *ProxyDatastoreModel) SetVmCount(v int32)`

SetVmCount sets VmCount field to given value.

### HasVmCount

`func (o *ProxyDatastoreModel) HasVmCount() bool`

HasVmCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


