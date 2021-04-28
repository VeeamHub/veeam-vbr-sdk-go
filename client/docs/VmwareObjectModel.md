# VmwareObjectModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostName** | **string** | Name of the host. | 
**Name** | **string** | Name of the VMware vSphere object. | 
**Type** | [**EVmwareInventoryType**](EVmwareInventoryType.md) |  | 
**ObjectId** | Pointer to **string** | ID of the VMware vSphere object. | [optional] 

## Methods

### NewVmwareObjectModel

`func NewVmwareObjectModel(hostName string, name string, type_ EVmwareInventoryType, ) *VmwareObjectModel`

NewVmwareObjectModel instantiates a new VmwareObjectModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareObjectModelWithDefaults

`func NewVmwareObjectModelWithDefaults() *VmwareObjectModel`

NewVmwareObjectModelWithDefaults instantiates a new VmwareObjectModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostName

`func (o *VmwareObjectModel) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *VmwareObjectModel) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *VmwareObjectModel) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetName

`func (o *VmwareObjectModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmwareObjectModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmwareObjectModel) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *VmwareObjectModel) GetType() EVmwareInventoryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VmwareObjectModel) GetTypeOk() (*EVmwareInventoryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VmwareObjectModel) SetType(v EVmwareInventoryType)`

SetType sets Type field to given value.


### GetObjectId

`func (o *VmwareObjectModel) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *VmwareObjectModel) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *VmwareObjectModel) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *VmwareObjectModel) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


