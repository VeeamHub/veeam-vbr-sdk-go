# InstantViVMCustomizedRecoverySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | [**InstantViVMCustomizedRecoveryDestinationSpec**](InstantViVMCustomizedRecoveryDestinationSpec.md) |  | 
**Datastore** | [**InstantViVMCustomizedRecoveryDatastoreSpec**](InstantViVMCustomizedRecoveryDatastoreSpec.md) |  | 
**Overwrite** | Pointer to **bool** | If *true*, Veeam Backup &amp; Replication overwrites the existing VM that has the same name. | [optional] 

## Methods

### NewInstantViVMCustomizedRecoverySpec

`func NewInstantViVMCustomizedRecoverySpec(destination InstantViVMCustomizedRecoveryDestinationSpec, datastore InstantViVMCustomizedRecoveryDatastoreSpec, ) *InstantViVMCustomizedRecoverySpec`

NewInstantViVMCustomizedRecoverySpec instantiates a new InstantViVMCustomizedRecoverySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstantViVMCustomizedRecoverySpecWithDefaults

`func NewInstantViVMCustomizedRecoverySpecWithDefaults() *InstantViVMCustomizedRecoverySpec`

NewInstantViVMCustomizedRecoverySpecWithDefaults instantiates a new InstantViVMCustomizedRecoverySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *InstantViVMCustomizedRecoverySpec) GetDestination() InstantViVMCustomizedRecoveryDestinationSpec`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *InstantViVMCustomizedRecoverySpec) GetDestinationOk() (*InstantViVMCustomizedRecoveryDestinationSpec, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *InstantViVMCustomizedRecoverySpec) SetDestination(v InstantViVMCustomizedRecoveryDestinationSpec)`

SetDestination sets Destination field to given value.


### GetDatastore

`func (o *InstantViVMCustomizedRecoverySpec) GetDatastore() InstantViVMCustomizedRecoveryDatastoreSpec`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *InstantViVMCustomizedRecoverySpec) GetDatastoreOk() (*InstantViVMCustomizedRecoveryDatastoreSpec, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *InstantViVMCustomizedRecoverySpec) SetDatastore(v InstantViVMCustomizedRecoveryDatastoreSpec)`

SetDatastore sets Datastore field to given value.


### GetOverwrite

`func (o *InstantViVMCustomizedRecoverySpec) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *InstantViVMCustomizedRecoverySpec) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *InstantViVMCustomizedRecoverySpec) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *InstantViVMCustomizedRecoverySpec) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


