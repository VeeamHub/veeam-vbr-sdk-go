# BackupLogShippingServersImportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoSelection** | **bool** | If *true*, Veeam Backup &amp; Replication chooses an optimal log shipping server automatically. | 
**ShippingServerNames** | Pointer to **[]string** | Array of servers that are explicitly selected for log shipping. | [optional] 

## Methods

### NewBackupLogShippingServersImportModel

`func NewBackupLogShippingServersImportModel(autoSelection bool, ) *BackupLogShippingServersImportModel`

NewBackupLogShippingServersImportModel instantiates a new BackupLogShippingServersImportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupLogShippingServersImportModelWithDefaults

`func NewBackupLogShippingServersImportModelWithDefaults() *BackupLogShippingServersImportModel`

NewBackupLogShippingServersImportModelWithDefaults instantiates a new BackupLogShippingServersImportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoSelection

`func (o *BackupLogShippingServersImportModel) GetAutoSelection() bool`

GetAutoSelection returns the AutoSelection field if non-nil, zero value otherwise.

### GetAutoSelectionOk

`func (o *BackupLogShippingServersImportModel) GetAutoSelectionOk() (*bool, bool)`

GetAutoSelectionOk returns a tuple with the AutoSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSelection

`func (o *BackupLogShippingServersImportModel) SetAutoSelection(v bool)`

SetAutoSelection sets AutoSelection field to given value.


### GetShippingServerNames

`func (o *BackupLogShippingServersImportModel) GetShippingServerNames() []string`

GetShippingServerNames returns the ShippingServerNames field if non-nil, zero value otherwise.

### GetShippingServerNamesOk

`func (o *BackupLogShippingServersImportModel) GetShippingServerNamesOk() (*[]string, bool)`

GetShippingServerNamesOk returns a tuple with the ShippingServerNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingServerNames

`func (o *BackupLogShippingServersImportModel) SetShippingServerNames(v []string)`

SetShippingServerNames sets ShippingServerNames field to given value.

### HasShippingServerNames

`func (o *BackupLogShippingServersImportModel) HasShippingServerNames() bool`

HasShippingServerNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


