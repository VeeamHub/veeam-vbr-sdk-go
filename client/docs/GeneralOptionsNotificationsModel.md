# GeneralOptionsNotificationsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageSpaceThresholdEnabled** | **bool** | If *true*, notifications about critical amount of free space in backup storage are enabled. | 
**StorageSpaceThreshold** | **int32** | Space threshold of backup storage, in percent. | 
**DatastoreSpaceThresholdEnabled** | **bool** | If *true*, notifications about critical amount of free space in production datastore are enabled. | 
**DatastoreSpaceThreshold** | **int32** | Space threshold of production datastore, in percent. | 
**SkipVMSpaceThresholdEnabled** | **bool** | If *true* and the &#x60;skipVMSpaceThreshold&#x60; threshold is reached, Veeam Backup &amp; Replication terminates backup and replication jobs working with production datastores before VM snapshots are taken. | 
**SkipVMSpaceThreshold** | **int32** | Space threshold of production datastore, in percent. | 
**NotifyOnSupportExpiration** | **bool** | If *true*, notifications about support contract expiration are enabled. | 
**NotifyOnUpdates** | **bool** | If *true*, notifications about updates are enabled. | 

## Methods

### NewGeneralOptionsNotificationsModel

`func NewGeneralOptionsNotificationsModel(storageSpaceThresholdEnabled bool, storageSpaceThreshold int32, datastoreSpaceThresholdEnabled bool, datastoreSpaceThreshold int32, skipVMSpaceThresholdEnabled bool, skipVMSpaceThreshold int32, notifyOnSupportExpiration bool, notifyOnUpdates bool, ) *GeneralOptionsNotificationsModel`

NewGeneralOptionsNotificationsModel instantiates a new GeneralOptionsNotificationsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneralOptionsNotificationsModelWithDefaults

`func NewGeneralOptionsNotificationsModelWithDefaults() *GeneralOptionsNotificationsModel`

NewGeneralOptionsNotificationsModelWithDefaults instantiates a new GeneralOptionsNotificationsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageSpaceThresholdEnabled

`func (o *GeneralOptionsNotificationsModel) GetStorageSpaceThresholdEnabled() bool`

GetStorageSpaceThresholdEnabled returns the StorageSpaceThresholdEnabled field if non-nil, zero value otherwise.

### GetStorageSpaceThresholdEnabledOk

`func (o *GeneralOptionsNotificationsModel) GetStorageSpaceThresholdEnabledOk() (*bool, bool)`

GetStorageSpaceThresholdEnabledOk returns a tuple with the StorageSpaceThresholdEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSpaceThresholdEnabled

`func (o *GeneralOptionsNotificationsModel) SetStorageSpaceThresholdEnabled(v bool)`

SetStorageSpaceThresholdEnabled sets StorageSpaceThresholdEnabled field to given value.


### GetStorageSpaceThreshold

`func (o *GeneralOptionsNotificationsModel) GetStorageSpaceThreshold() int32`

GetStorageSpaceThreshold returns the StorageSpaceThreshold field if non-nil, zero value otherwise.

### GetStorageSpaceThresholdOk

`func (o *GeneralOptionsNotificationsModel) GetStorageSpaceThresholdOk() (*int32, bool)`

GetStorageSpaceThresholdOk returns a tuple with the StorageSpaceThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSpaceThreshold

`func (o *GeneralOptionsNotificationsModel) SetStorageSpaceThreshold(v int32)`

SetStorageSpaceThreshold sets StorageSpaceThreshold field to given value.


### GetDatastoreSpaceThresholdEnabled

`func (o *GeneralOptionsNotificationsModel) GetDatastoreSpaceThresholdEnabled() bool`

GetDatastoreSpaceThresholdEnabled returns the DatastoreSpaceThresholdEnabled field if non-nil, zero value otherwise.

### GetDatastoreSpaceThresholdEnabledOk

`func (o *GeneralOptionsNotificationsModel) GetDatastoreSpaceThresholdEnabledOk() (*bool, bool)`

GetDatastoreSpaceThresholdEnabledOk returns a tuple with the DatastoreSpaceThresholdEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreSpaceThresholdEnabled

`func (o *GeneralOptionsNotificationsModel) SetDatastoreSpaceThresholdEnabled(v bool)`

SetDatastoreSpaceThresholdEnabled sets DatastoreSpaceThresholdEnabled field to given value.


### GetDatastoreSpaceThreshold

`func (o *GeneralOptionsNotificationsModel) GetDatastoreSpaceThreshold() int32`

GetDatastoreSpaceThreshold returns the DatastoreSpaceThreshold field if non-nil, zero value otherwise.

### GetDatastoreSpaceThresholdOk

`func (o *GeneralOptionsNotificationsModel) GetDatastoreSpaceThresholdOk() (*int32, bool)`

GetDatastoreSpaceThresholdOk returns a tuple with the DatastoreSpaceThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreSpaceThreshold

`func (o *GeneralOptionsNotificationsModel) SetDatastoreSpaceThreshold(v int32)`

SetDatastoreSpaceThreshold sets DatastoreSpaceThreshold field to given value.


### GetSkipVMSpaceThresholdEnabled

`func (o *GeneralOptionsNotificationsModel) GetSkipVMSpaceThresholdEnabled() bool`

GetSkipVMSpaceThresholdEnabled returns the SkipVMSpaceThresholdEnabled field if non-nil, zero value otherwise.

### GetSkipVMSpaceThresholdEnabledOk

`func (o *GeneralOptionsNotificationsModel) GetSkipVMSpaceThresholdEnabledOk() (*bool, bool)`

GetSkipVMSpaceThresholdEnabledOk returns a tuple with the SkipVMSpaceThresholdEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipVMSpaceThresholdEnabled

`func (o *GeneralOptionsNotificationsModel) SetSkipVMSpaceThresholdEnabled(v bool)`

SetSkipVMSpaceThresholdEnabled sets SkipVMSpaceThresholdEnabled field to given value.


### GetSkipVMSpaceThreshold

`func (o *GeneralOptionsNotificationsModel) GetSkipVMSpaceThreshold() int32`

GetSkipVMSpaceThreshold returns the SkipVMSpaceThreshold field if non-nil, zero value otherwise.

### GetSkipVMSpaceThresholdOk

`func (o *GeneralOptionsNotificationsModel) GetSkipVMSpaceThresholdOk() (*int32, bool)`

GetSkipVMSpaceThresholdOk returns a tuple with the SkipVMSpaceThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipVMSpaceThreshold

`func (o *GeneralOptionsNotificationsModel) SetSkipVMSpaceThreshold(v int32)`

SetSkipVMSpaceThreshold sets SkipVMSpaceThreshold field to given value.


### GetNotifyOnSupportExpiration

`func (o *GeneralOptionsNotificationsModel) GetNotifyOnSupportExpiration() bool`

GetNotifyOnSupportExpiration returns the NotifyOnSupportExpiration field if non-nil, zero value otherwise.

### GetNotifyOnSupportExpirationOk

`func (o *GeneralOptionsNotificationsModel) GetNotifyOnSupportExpirationOk() (*bool, bool)`

GetNotifyOnSupportExpirationOk returns a tuple with the NotifyOnSupportExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOnSupportExpiration

`func (o *GeneralOptionsNotificationsModel) SetNotifyOnSupportExpiration(v bool)`

SetNotifyOnSupportExpiration sets NotifyOnSupportExpiration field to given value.


### GetNotifyOnUpdates

`func (o *GeneralOptionsNotificationsModel) GetNotifyOnUpdates() bool`

GetNotifyOnUpdates returns the NotifyOnUpdates field if non-nil, zero value otherwise.

### GetNotifyOnUpdatesOk

`func (o *GeneralOptionsNotificationsModel) GetNotifyOnUpdatesOk() (*bool, bool)`

GetNotifyOnUpdatesOk returns a tuple with the NotifyOnUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOnUpdates

`func (o *GeneralOptionsNotificationsModel) SetNotifyOnUpdates(v bool)`

SetNotifyOnUpdates sets NotifyOnUpdates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


