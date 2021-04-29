# VmwareFcdWriteCacheSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RedirectIsEnabled** | **bool** | If *true*, cache redirection is enabled. In this case, all changes made to the recovered disks while the Instant FCD Recovery is active are redirected to the specified &#x60;cacheDatastore&#x60; associated with the &#x60;storagePolicy&#x60;. | 
**CacheDatastore** | Pointer to [**VmwareObjectModel**](VmwareObjectModel.md) |  | [optional] 
**StoragePolicy** | Pointer to [**VmwareObjectModel**](VmwareObjectModel.md) |  | [optional] 

## Methods

### NewVmwareFcdWriteCacheSpec

`func NewVmwareFcdWriteCacheSpec(redirectIsEnabled bool, ) *VmwareFcdWriteCacheSpec`

NewVmwareFcdWriteCacheSpec instantiates a new VmwareFcdWriteCacheSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareFcdWriteCacheSpecWithDefaults

`func NewVmwareFcdWriteCacheSpecWithDefaults() *VmwareFcdWriteCacheSpec`

NewVmwareFcdWriteCacheSpecWithDefaults instantiates a new VmwareFcdWriteCacheSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedirectIsEnabled

`func (o *VmwareFcdWriteCacheSpec) GetRedirectIsEnabled() bool`

GetRedirectIsEnabled returns the RedirectIsEnabled field if non-nil, zero value otherwise.

### GetRedirectIsEnabledOk

`func (o *VmwareFcdWriteCacheSpec) GetRedirectIsEnabledOk() (*bool, bool)`

GetRedirectIsEnabledOk returns a tuple with the RedirectIsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectIsEnabled

`func (o *VmwareFcdWriteCacheSpec) SetRedirectIsEnabled(v bool)`

SetRedirectIsEnabled sets RedirectIsEnabled field to given value.


### GetCacheDatastore

`func (o *VmwareFcdWriteCacheSpec) GetCacheDatastore() VmwareObjectModel`

GetCacheDatastore returns the CacheDatastore field if non-nil, zero value otherwise.

### GetCacheDatastoreOk

`func (o *VmwareFcdWriteCacheSpec) GetCacheDatastoreOk() (*VmwareObjectModel, bool)`

GetCacheDatastoreOk returns a tuple with the CacheDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheDatastore

`func (o *VmwareFcdWriteCacheSpec) SetCacheDatastore(v VmwareObjectModel)`

SetCacheDatastore sets CacheDatastore field to given value.

### HasCacheDatastore

`func (o *VmwareFcdWriteCacheSpec) HasCacheDatastore() bool`

HasCacheDatastore returns a boolean if a field has been set.

### GetStoragePolicy

`func (o *VmwareFcdWriteCacheSpec) GetStoragePolicy() VmwareObjectModel`

GetStoragePolicy returns the StoragePolicy field if non-nil, zero value otherwise.

### GetStoragePolicyOk

`func (o *VmwareFcdWriteCacheSpec) GetStoragePolicyOk() (*VmwareObjectModel, bool)`

GetStoragePolicyOk returns a tuple with the StoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicy

`func (o *VmwareFcdWriteCacheSpec) SetStoragePolicy(v VmwareObjectModel)`

SetStoragePolicy sets StoragePolicy field to given value.

### HasStoragePolicy

`func (o *VmwareFcdWriteCacheSpec) HasStoragePolicy() bool`

HasStoragePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


