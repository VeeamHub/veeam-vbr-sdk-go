# RestoreTargetDatastoreSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastore** | Pointer to [**VmwareObjectModel**](VmwareObjectModel.md) |  | [optional] 
**DiskType** | Pointer to [**EDiskCreationMode**](EDiskCreationMode.md) |  | [optional] 

## Methods

### NewRestoreTargetDatastoreSpec

`func NewRestoreTargetDatastoreSpec() *RestoreTargetDatastoreSpec`

NewRestoreTargetDatastoreSpec instantiates a new RestoreTargetDatastoreSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreTargetDatastoreSpecWithDefaults

`func NewRestoreTargetDatastoreSpecWithDefaults() *RestoreTargetDatastoreSpec`

NewRestoreTargetDatastoreSpecWithDefaults instantiates a new RestoreTargetDatastoreSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastore

`func (o *RestoreTargetDatastoreSpec) GetDatastore() VmwareObjectModel`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *RestoreTargetDatastoreSpec) GetDatastoreOk() (*VmwareObjectModel, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *RestoreTargetDatastoreSpec) SetDatastore(v VmwareObjectModel)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *RestoreTargetDatastoreSpec) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetDiskType

`func (o *RestoreTargetDatastoreSpec) GetDiskType() EDiskCreationMode`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *RestoreTargetDatastoreSpec) GetDiskTypeOk() (*EDiskCreationMode, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *RestoreTargetDatastoreSpec) SetDiskType(v EDiskCreationMode)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *RestoreTargetDatastoreSpec) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


