# BackupJobImportProxiesModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutomaticSelection** | **bool** | If *true*, backup proxies are detected and assigned automatically. | [default to true]
**Proxies** | Pointer to [**[]BackupProxyImportModel**](BackupProxyImportModel.md) | Array of backup proxies. | [optional] 

## Methods

### NewBackupJobImportProxiesModel

`func NewBackupJobImportProxiesModel(automaticSelection bool, ) *BackupJobImportProxiesModel`

NewBackupJobImportProxiesModel instantiates a new BackupJobImportProxiesModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupJobImportProxiesModelWithDefaults

`func NewBackupJobImportProxiesModelWithDefaults() *BackupJobImportProxiesModel`

NewBackupJobImportProxiesModelWithDefaults instantiates a new BackupJobImportProxiesModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomaticSelection

`func (o *BackupJobImportProxiesModel) GetAutomaticSelection() bool`

GetAutomaticSelection returns the AutomaticSelection field if non-nil, zero value otherwise.

### GetAutomaticSelectionOk

`func (o *BackupJobImportProxiesModel) GetAutomaticSelectionOk() (*bool, bool)`

GetAutomaticSelectionOk returns a tuple with the AutomaticSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticSelection

`func (o *BackupJobImportProxiesModel) SetAutomaticSelection(v bool)`

SetAutomaticSelection sets AutomaticSelection field to given value.


### GetProxies

`func (o *BackupJobImportProxiesModel) GetProxies() []BackupProxyImportModel`

GetProxies returns the Proxies field if non-nil, zero value otherwise.

### GetProxiesOk

`func (o *BackupJobImportProxiesModel) GetProxiesOk() (*[]BackupProxyImportModel, bool)`

GetProxiesOk returns a tuple with the Proxies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxies

`func (o *BackupJobImportProxiesModel) SetProxies(v []BackupProxyImportModel)`

SetProxies sets Proxies field to given value.

### HasProxies

`func (o *BackupJobImportProxiesModel) HasProxies() bool`

HasProxies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


