# BackupProxiesSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoSelection** | **bool** | If *true*, backup proxies are detected and assigned automatically. | [default to true]
**ProxyIds** | Pointer to **[]string** | Array of proxy IDs. | [optional] 

## Methods

### NewBackupProxiesSettingsModel

`func NewBackupProxiesSettingsModel(autoSelection bool, ) *BackupProxiesSettingsModel`

NewBackupProxiesSettingsModel instantiates a new BackupProxiesSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupProxiesSettingsModelWithDefaults

`func NewBackupProxiesSettingsModelWithDefaults() *BackupProxiesSettingsModel`

NewBackupProxiesSettingsModelWithDefaults instantiates a new BackupProxiesSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoSelection

`func (o *BackupProxiesSettingsModel) GetAutoSelection() bool`

GetAutoSelection returns the AutoSelection field if non-nil, zero value otherwise.

### GetAutoSelectionOk

`func (o *BackupProxiesSettingsModel) GetAutoSelectionOk() (*bool, bool)`

GetAutoSelectionOk returns a tuple with the AutoSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSelection

`func (o *BackupProxiesSettingsModel) SetAutoSelection(v bool)`

SetAutoSelection sets AutoSelection field to given value.


### GetProxyIds

`func (o *BackupProxiesSettingsModel) GetProxyIds() []string`

GetProxyIds returns the ProxyIds field if non-nil, zero value otherwise.

### GetProxyIdsOk

`func (o *BackupProxiesSettingsModel) GetProxyIdsOk() (*[]string, bool)`

GetProxyIdsOk returns a tuple with the ProxyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyIds

`func (o *BackupProxiesSettingsModel) SetProxyIds(v []string)`

SetProxyIds sets ProxyIds field to given value.

### HasProxyIds

`func (o *BackupProxiesSettingsModel) HasProxyIds() bool`

HasProxyIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


