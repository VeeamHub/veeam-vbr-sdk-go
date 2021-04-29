# BackupApplicationSettingsImportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmObject** | [**VmwareObjectModel**](VmwareObjectModel.md) |  | 
**Vss** | [**EApplicationSettingsVSS**](EApplicationSettingsVSS.md) |  | 
**UsePersistentGuestAgent** | Pointer to **bool** | If *true*, persistent guest agent is used. | [optional] 
**TransactionLogs** | Pointer to [**ETransactionLogsSettings**](ETransactionLogsSettings.md) |  | [optional] 
**Sql** | Pointer to [**BackupSQLSettingsImportModel**](BackupSQLSettingsImportModel.md) |  | [optional] 
**Oracle** | Pointer to [**BackupOracleSettingsImportModel**](BackupOracleSettingsImportModel.md) |  | [optional] 
**Exclusions** | Pointer to [**BackupFSExclusionsModel**](BackupFSExclusionsModel.md) |  | [optional] 
**Scripts** | Pointer to [**BackupScriptSettingsModel**](BackupScriptSettingsModel.md) |  | [optional] 

## Methods

### NewBackupApplicationSettingsImportModel

`func NewBackupApplicationSettingsImportModel(vmObject VmwareObjectModel, vss EApplicationSettingsVSS, ) *BackupApplicationSettingsImportModel`

NewBackupApplicationSettingsImportModel instantiates a new BackupApplicationSettingsImportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupApplicationSettingsImportModelWithDefaults

`func NewBackupApplicationSettingsImportModelWithDefaults() *BackupApplicationSettingsImportModel`

NewBackupApplicationSettingsImportModelWithDefaults instantiates a new BackupApplicationSettingsImportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmObject

`func (o *BackupApplicationSettingsImportModel) GetVmObject() VmwareObjectModel`

GetVmObject returns the VmObject field if non-nil, zero value otherwise.

### GetVmObjectOk

`func (o *BackupApplicationSettingsImportModel) GetVmObjectOk() (*VmwareObjectModel, bool)`

GetVmObjectOk returns a tuple with the VmObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmObject

`func (o *BackupApplicationSettingsImportModel) SetVmObject(v VmwareObjectModel)`

SetVmObject sets VmObject field to given value.


### GetVss

`func (o *BackupApplicationSettingsImportModel) GetVss() EApplicationSettingsVSS`

GetVss returns the Vss field if non-nil, zero value otherwise.

### GetVssOk

`func (o *BackupApplicationSettingsImportModel) GetVssOk() (*EApplicationSettingsVSS, bool)`

GetVssOk returns a tuple with the Vss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVss

`func (o *BackupApplicationSettingsImportModel) SetVss(v EApplicationSettingsVSS)`

SetVss sets Vss field to given value.


### GetUsePersistentGuestAgent

`func (o *BackupApplicationSettingsImportModel) GetUsePersistentGuestAgent() bool`

GetUsePersistentGuestAgent returns the UsePersistentGuestAgent field if non-nil, zero value otherwise.

### GetUsePersistentGuestAgentOk

`func (o *BackupApplicationSettingsImportModel) GetUsePersistentGuestAgentOk() (*bool, bool)`

GetUsePersistentGuestAgentOk returns a tuple with the UsePersistentGuestAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePersistentGuestAgent

`func (o *BackupApplicationSettingsImportModel) SetUsePersistentGuestAgent(v bool)`

SetUsePersistentGuestAgent sets UsePersistentGuestAgent field to given value.

### HasUsePersistentGuestAgent

`func (o *BackupApplicationSettingsImportModel) HasUsePersistentGuestAgent() bool`

HasUsePersistentGuestAgent returns a boolean if a field has been set.

### GetTransactionLogs

`func (o *BackupApplicationSettingsImportModel) GetTransactionLogs() ETransactionLogsSettings`

GetTransactionLogs returns the TransactionLogs field if non-nil, zero value otherwise.

### GetTransactionLogsOk

`func (o *BackupApplicationSettingsImportModel) GetTransactionLogsOk() (*ETransactionLogsSettings, bool)`

GetTransactionLogsOk returns a tuple with the TransactionLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionLogs

`func (o *BackupApplicationSettingsImportModel) SetTransactionLogs(v ETransactionLogsSettings)`

SetTransactionLogs sets TransactionLogs field to given value.

### HasTransactionLogs

`func (o *BackupApplicationSettingsImportModel) HasTransactionLogs() bool`

HasTransactionLogs returns a boolean if a field has been set.

### GetSql

`func (o *BackupApplicationSettingsImportModel) GetSql() BackupSQLSettingsImportModel`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *BackupApplicationSettingsImportModel) GetSqlOk() (*BackupSQLSettingsImportModel, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *BackupApplicationSettingsImportModel) SetSql(v BackupSQLSettingsImportModel)`

SetSql sets Sql field to given value.

### HasSql

`func (o *BackupApplicationSettingsImportModel) HasSql() bool`

HasSql returns a boolean if a field has been set.

### GetOracle

`func (o *BackupApplicationSettingsImportModel) GetOracle() BackupOracleSettingsImportModel`

GetOracle returns the Oracle field if non-nil, zero value otherwise.

### GetOracleOk

`func (o *BackupApplicationSettingsImportModel) GetOracleOk() (*BackupOracleSettingsImportModel, bool)`

GetOracleOk returns a tuple with the Oracle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracle

`func (o *BackupApplicationSettingsImportModel) SetOracle(v BackupOracleSettingsImportModel)`

SetOracle sets Oracle field to given value.

### HasOracle

`func (o *BackupApplicationSettingsImportModel) HasOracle() bool`

HasOracle returns a boolean if a field has been set.

### GetExclusions

`func (o *BackupApplicationSettingsImportModel) GetExclusions() BackupFSExclusionsModel`

GetExclusions returns the Exclusions field if non-nil, zero value otherwise.

### GetExclusionsOk

`func (o *BackupApplicationSettingsImportModel) GetExclusionsOk() (*BackupFSExclusionsModel, bool)`

GetExclusionsOk returns a tuple with the Exclusions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusions

`func (o *BackupApplicationSettingsImportModel) SetExclusions(v BackupFSExclusionsModel)`

SetExclusions sets Exclusions field to given value.

### HasExclusions

`func (o *BackupApplicationSettingsImportModel) HasExclusions() bool`

HasExclusions returns a boolean if a field has been set.

### GetScripts

`func (o *BackupApplicationSettingsImportModel) GetScripts() BackupScriptSettingsModel`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *BackupApplicationSettingsImportModel) GetScriptsOk() (*BackupScriptSettingsModel, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *BackupApplicationSettingsImportModel) SetScripts(v BackupScriptSettingsModel)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *BackupApplicationSettingsImportModel) HasScripts() bool`

HasScripts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


