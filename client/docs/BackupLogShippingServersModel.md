# BackupLogShippingServersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoSelection** | **bool** | If *true*, Veeam Backup &amp; Replication chooses an optimal log shipping server automatically. | 
**ShippingServerIds** | Pointer to **[]string** | Array of servers that are explicitly selected for log shipping. | [optional] 

## Methods

### NewBackupLogShippingServersModel

`func NewBackupLogShippingServersModel(autoSelection bool, ) *BackupLogShippingServersModel`

NewBackupLogShippingServersModel instantiates a new BackupLogShippingServersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupLogShippingServersModelWithDefaults

`func NewBackupLogShippingServersModelWithDefaults() *BackupLogShippingServersModel`

NewBackupLogShippingServersModelWithDefaults instantiates a new BackupLogShippingServersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoSelection

`func (o *BackupLogShippingServersModel) GetAutoSelection() bool`

GetAutoSelection returns the AutoSelection field if non-nil, zero value otherwise.

### GetAutoSelectionOk

`func (o *BackupLogShippingServersModel) GetAutoSelectionOk() (*bool, bool)`

GetAutoSelectionOk returns a tuple with the AutoSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSelection

`func (o *BackupLogShippingServersModel) SetAutoSelection(v bool)`

SetAutoSelection sets AutoSelection field to given value.


### GetShippingServerIds

`func (o *BackupLogShippingServersModel) GetShippingServerIds() []string`

GetShippingServerIds returns the ShippingServerIds field if non-nil, zero value otherwise.

### GetShippingServerIdsOk

`func (o *BackupLogShippingServersModel) GetShippingServerIdsOk() (*[]string, bool)`

GetShippingServerIdsOk returns a tuple with the ShippingServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingServerIds

`func (o *BackupLogShippingServersModel) SetShippingServerIds(v []string)`

SetShippingServerIds sets ShippingServerIds field to given value.

### HasShippingServerIds

`func (o *BackupLogShippingServersModel) HasShippingServerIds() bool`

HasShippingServerIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


