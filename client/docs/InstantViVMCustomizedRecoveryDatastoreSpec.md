# InstantViVMCustomizedRecoveryDatastoreSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RedirectIsEnabled** | **bool** | If *true*, redo logs are redirected to &#x60;cacheDatastore&#x60;. | 
**CacheDatastore** | Pointer to [**VmwareObjectModel**](VmwareObjectModel.md) |  | [optional] 

## Methods

### NewInstantViVMCustomizedRecoveryDatastoreSpec

`func NewInstantViVMCustomizedRecoveryDatastoreSpec(redirectIsEnabled bool, ) *InstantViVMCustomizedRecoveryDatastoreSpec`

NewInstantViVMCustomizedRecoveryDatastoreSpec instantiates a new InstantViVMCustomizedRecoveryDatastoreSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstantViVMCustomizedRecoveryDatastoreSpecWithDefaults

`func NewInstantViVMCustomizedRecoveryDatastoreSpecWithDefaults() *InstantViVMCustomizedRecoveryDatastoreSpec`

NewInstantViVMCustomizedRecoveryDatastoreSpecWithDefaults instantiates a new InstantViVMCustomizedRecoveryDatastoreSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedirectIsEnabled

`func (o *InstantViVMCustomizedRecoveryDatastoreSpec) GetRedirectIsEnabled() bool`

GetRedirectIsEnabled returns the RedirectIsEnabled field if non-nil, zero value otherwise.

### GetRedirectIsEnabledOk

`func (o *InstantViVMCustomizedRecoveryDatastoreSpec) GetRedirectIsEnabledOk() (*bool, bool)`

GetRedirectIsEnabledOk returns a tuple with the RedirectIsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectIsEnabled

`func (o *InstantViVMCustomizedRecoveryDatastoreSpec) SetRedirectIsEnabled(v bool)`

SetRedirectIsEnabled sets RedirectIsEnabled field to given value.


### GetCacheDatastore

`func (o *InstantViVMCustomizedRecoveryDatastoreSpec) GetCacheDatastore() VmwareObjectModel`

GetCacheDatastore returns the CacheDatastore field if non-nil, zero value otherwise.

### GetCacheDatastoreOk

`func (o *InstantViVMCustomizedRecoveryDatastoreSpec) GetCacheDatastoreOk() (*VmwareObjectModel, bool)`

GetCacheDatastoreOk returns a tuple with the CacheDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheDatastore

`func (o *InstantViVMCustomizedRecoveryDatastoreSpec) SetCacheDatastore(v VmwareObjectModel)`

SetCacheDatastore sets CacheDatastore field to given value.

### HasCacheDatastore

`func (o *InstantViVMCustomizedRecoveryDatastoreSpec) HasCacheDatastore() bool`

HasCacheDatastore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


