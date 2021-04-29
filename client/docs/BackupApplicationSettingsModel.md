# BackupApplicationSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmObject** | [**VmwareObjectModel**](VmwareObjectModel.md) |  | 
**Vss** | [**EApplicationSettingsVSS**](EApplicationSettingsVSS.md) |  | 
**UsePersistentGuestAgent** | Pointer to **bool** | If *true*, persistent guest agent is used. | [optional] 
**TransactionLogs** | Pointer to [**ETransactionLogsSettings**](ETransactionLogsSettings.md) |  | [optional] 
**Sql** | Pointer to [**BackupSQLSettingsModel**](BackupSQLSettingsModel.md) |  | [optional] 
**Oracle** | Pointer to [**BackupOracleSettingsModel**](BackupOracleSettingsModel.md) |  | [optional] 
**Exclusions** | Pointer to [**BackupFSExclusionsModel**](BackupFSExclusionsModel.md) |  | [optional] 
**Scripts** | Pointer to [**BackupScriptSettingsModel**](BackupScriptSettingsModel.md) |  | [optional] 

## Methods

### NewBackupApplicationSettingsModel

`func NewBackupApplicationSettingsModel(vmObject VmwareObjectModel, vss EApplicationSettingsVSS, ) *BackupApplicationSettingsModel`

NewBackupApplicationSettingsModel instantiates a new BackupApplicationSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupApplicationSettingsModelWithDefaults

`func NewBackupApplicationSettingsModelWithDefaults() *BackupApplicationSettingsModel`

NewBackupApplicationSettingsModelWithDefaults instantiates a new BackupApplicationSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmObject

`func (o *BackupApplicationSettingsModel) GetVmObject() VmwareObjectModel`

GetVmObject returns the VmObject field if non-nil, zero value otherwise.

### GetVmObjectOk

`func (o *BackupApplicationSettingsModel) GetVmObjectOk() (*VmwareObjectModel, bool)`

GetVmObjectOk returns a tuple with the VmObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmObject

`func (o *BackupApplicationSettingsModel) SetVmObject(v VmwareObjectModel)`

SetVmObject sets VmObject field to given value.


### GetVss

`func (o *BackupApplicationSettingsModel) GetVss() EApplicationSettingsVSS`

GetVss returns the Vss field if non-nil, zero value otherwise.

### GetVssOk

`func (o *BackupApplicationSettingsModel) GetVssOk() (*EApplicationSettingsVSS, bool)`

GetVssOk returns a tuple with the Vss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVss

`func (o *BackupApplicationSettingsModel) SetVss(v EApplicationSettingsVSS)`

SetVss sets Vss field to given value.


### GetUsePersistentGuestAgent

`func (o *BackupApplicationSettingsModel) GetUsePersistentGuestAgent() bool`

GetUsePersistentGuestAgent returns the UsePersistentGuestAgent field if non-nil, zero value otherwise.

### GetUsePersistentGuestAgentOk

`func (o *BackupApplicationSettingsModel) GetUsePersistentGuestAgentOk() (*bool, bool)`

GetUsePersistentGuestAgentOk returns a tuple with the UsePersistentGuestAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePersistentGuestAgent

`func (o *BackupApplicationSettingsModel) SetUsePersistentGuestAgent(v bool)`

SetUsePersistentGuestAgent sets UsePersistentGuestAgent field to given value.

### HasUsePersistentGuestAgent

`func (o *BackupApplicationSettingsModel) HasUsePersistentGuestAgent() bool`

HasUsePersistentGuestAgent returns a boolean if a field has been set.

### GetTransactionLogs

`func (o *BackupApplicationSettingsModel) GetTransactionLogs() ETransactionLogsSettings`

GetTransactionLogs returns the TransactionLogs field if non-nil, zero value otherwise.

### GetTransactionLogsOk

`func (o *BackupApplicationSettingsModel) GetTransactionLogsOk() (*ETransactionLogsSettings, bool)`

GetTransactionLogsOk returns a tuple with the TransactionLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionLogs

`func (o *BackupApplicationSettingsModel) SetTransactionLogs(v ETransactionLogsSettings)`

SetTransactionLogs sets TransactionLogs field to given value.

### HasTransactionLogs

`func (o *BackupApplicationSettingsModel) HasTransactionLogs() bool`

HasTransactionLogs returns a boolean if a field has been set.

### GetSql

`func (o *BackupApplicationSettingsModel) GetSql() BackupSQLSettingsModel`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *BackupApplicationSettingsModel) GetSqlOk() (*BackupSQLSettingsModel, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *BackupApplicationSettingsModel) SetSql(v BackupSQLSettingsModel)`

SetSql sets Sql field to given value.

### HasSql

`func (o *BackupApplicationSettingsModel) HasSql() bool`

HasSql returns a boolean if a field has been set.

### GetOracle

`func (o *BackupApplicationSettingsModel) GetOracle() BackupOracleSettingsModel`

GetOracle returns the Oracle field if non-nil, zero value otherwise.

### GetOracleOk

`func (o *BackupApplicationSettingsModel) GetOracleOk() (*BackupOracleSettingsModel, bool)`

GetOracleOk returns a tuple with the Oracle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracle

`func (o *BackupApplicationSettingsModel) SetOracle(v BackupOracleSettingsModel)`

SetOracle sets Oracle field to given value.

### HasOracle

`func (o *BackupApplicationSettingsModel) HasOracle() bool`

HasOracle returns a boolean if a field has been set.

### GetExclusions

`func (o *BackupApplicationSettingsModel) GetExclusions() BackupFSExclusionsModel`

GetExclusions returns the Exclusions field if non-nil, zero value otherwise.

### GetExclusionsOk

`func (o *BackupApplicationSettingsModel) GetExclusionsOk() (*BackupFSExclusionsModel, bool)`

GetExclusionsOk returns a tuple with the Exclusions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusions

`func (o *BackupApplicationSettingsModel) SetExclusions(v BackupFSExclusionsModel)`

SetExclusions sets Exclusions field to given value.

### HasExclusions

`func (o *BackupApplicationSettingsModel) HasExclusions() bool`

HasExclusions returns a boolean if a field has been set.

### GetScripts

`func (o *BackupApplicationSettingsModel) GetScripts() BackupScriptSettingsModel`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *BackupApplicationSettingsModel) GetScriptsOk() (*BackupScriptSettingsModel, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *BackupApplicationSettingsModel) SetScripts(v BackupScriptSettingsModel)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *BackupApplicationSettingsModel) HasScripts() bool`

HasScripts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


