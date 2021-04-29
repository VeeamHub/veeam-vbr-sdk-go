# ObjectRestorePointDiskModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | ID of the disk. | 
**Type** | [**EDiskInfoType**](EDiskInfoType.md) |  | 
**Name** | **string** | Name of the disk. | 
**Capacity** | **int64** | Capacity of the disk. | 
**State** | [**EDiskInfoProcessState**](EDiskInfoProcessState.md) |  | 

## Methods

### NewObjectRestorePointDiskModel

`func NewObjectRestorePointDiskModel(uid string, type_ EDiskInfoType, name string, capacity int64, state EDiskInfoProcessState, ) *ObjectRestorePointDiskModel`

NewObjectRestorePointDiskModel instantiates a new ObjectRestorePointDiskModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectRestorePointDiskModelWithDefaults

`func NewObjectRestorePointDiskModelWithDefaults() *ObjectRestorePointDiskModel`

NewObjectRestorePointDiskModelWithDefaults instantiates a new ObjectRestorePointDiskModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *ObjectRestorePointDiskModel) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ObjectRestorePointDiskModel) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ObjectRestorePointDiskModel) SetUid(v string)`

SetUid sets Uid field to given value.


### GetType

`func (o *ObjectRestorePointDiskModel) GetType() EDiskInfoType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ObjectRestorePointDiskModel) GetTypeOk() (*EDiskInfoType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ObjectRestorePointDiskModel) SetType(v EDiskInfoType)`

SetType sets Type field to given value.


### GetName

`func (o *ObjectRestorePointDiskModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectRestorePointDiskModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectRestorePointDiskModel) SetName(v string)`

SetName sets Name field to given value.


### GetCapacity

`func (o *ObjectRestorePointDiskModel) GetCapacity() int64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *ObjectRestorePointDiskModel) GetCapacityOk() (*int64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *ObjectRestorePointDiskModel) SetCapacity(v int64)`

SetCapacity sets Capacity field to given value.


### GetState

`func (o *ObjectRestorePointDiskModel) GetState() EDiskInfoProcessState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ObjectRestorePointDiskModel) GetStateOk() (*EDiskInfoProcessState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ObjectRestorePointDiskModel) SetState(v EDiskInfoProcessState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


