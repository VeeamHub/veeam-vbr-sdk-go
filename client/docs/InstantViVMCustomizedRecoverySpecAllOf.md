# InstantViVMCustomizedRecoverySpecAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | [**InstantViVMCustomizedRecoveryDestinationSpec**](InstantViVMCustomizedRecoveryDestinationSpec.md) |  | 
**Datastore** | [**InstantViVMCustomizedRecoveryDatastoreSpec**](InstantViVMCustomizedRecoveryDatastoreSpec.md) |  | 
**Overwrite** | Pointer to **bool** | If *true*, Veeam Backup &amp; Replication overwrites the existing VM that has the same name. | [optional] 

## Methods

### NewInstantViVMCustomizedRecoverySpecAllOf

`func NewInstantViVMCustomizedRecoverySpecAllOf(destination InstantViVMCustomizedRecoveryDestinationSpec, datastore InstantViVMCustomizedRecoveryDatastoreSpec, ) *InstantViVMCustomizedRecoverySpecAllOf`

NewInstantViVMCustomizedRecoverySpecAllOf instantiates a new InstantViVMCustomizedRecoverySpecAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstantViVMCustomizedRecoverySpecAllOfWithDefaults

`func NewInstantViVMCustomizedRecoverySpecAllOfWithDefaults() *InstantViVMCustomizedRecoverySpecAllOf`

NewInstantViVMCustomizedRecoverySpecAllOfWithDefaults instantiates a new InstantViVMCustomizedRecoverySpecAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *InstantViVMCustomizedRecoverySpecAllOf) GetDestination() InstantViVMCustomizedRecoveryDestinationSpec`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *InstantViVMCustomizedRecoverySpecAllOf) GetDestinationOk() (*InstantViVMCustomizedRecoveryDestinationSpec, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *InstantViVMCustomizedRecoverySpecAllOf) SetDestination(v InstantViVMCustomizedRecoveryDestinationSpec)`

SetDestination sets Destination field to given value.


### GetDatastore

`func (o *InstantViVMCustomizedRecoverySpecAllOf) GetDatastore() InstantViVMCustomizedRecoveryDatastoreSpec`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *InstantViVMCustomizedRecoverySpecAllOf) GetDatastoreOk() (*InstantViVMCustomizedRecoveryDatastoreSpec, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *InstantViVMCustomizedRecoverySpecAllOf) SetDatastore(v InstantViVMCustomizedRecoveryDatastoreSpec)`

SetDatastore sets Datastore field to given value.


### GetOverwrite

`func (o *InstantViVMCustomizedRecoverySpecAllOf) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *InstantViVMCustomizedRecoverySpecAllOf) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *InstantViVMCustomizedRecoverySpecAllOf) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *InstantViVMCustomizedRecoverySpecAllOf) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


