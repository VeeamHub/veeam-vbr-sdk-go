# BackupOracleSettingsImportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseGuestCredentials** | **bool** | If *true*, Veeam Backup &amp; Replication uses credentials specified in the guest processing settings. | 
**Credentials** | Pointer to [**CredentialsImportModel**](CredentialsImportModel.md) |  | [optional] 
**ArchiveLogs** | [**EBackupOracleLogsSettings**](EBackupOracleLogsSettings.md) |  | 
**DeleteHoursCount** | Pointer to **int32** | Time period in hours to keep archived logs. This parameter should be specified if the &#x60;EBackupOracleLogsSettings&#x60; value is *deleteExpiredHours*. | [optional] 
**DeleteGBsCount** | Pointer to **int32** | Maximum size for archived logs in GB. This parameter should be specified if the &#x60;EBackupOracleLogsSettings&#x60; value is *deleteExpiredGBs*. | [optional] 
**BackupLogs** | Pointer to **bool** | If *true*, archived logs are backed up. | [optional] 
**BackupMinsCount** | Pointer to **int32** | Frequency of archived log backup, in minutes. | [optional] 
**RetainLogBackups** | Pointer to [**ERetainLogBackupsType**](ERetainLogBackupsType.md) |  | [optional] 
**KeepDaysCount** | Pointer to **int32** | Number of days to keep archived logs. | [optional] 
**LogShippingServers** | Pointer to [**BackupLogShippingServersImportModel**](BackupLogShippingServersImportModel.md) |  | [optional] 

## Methods

### NewBackupOracleSettingsImportModel

`func NewBackupOracleSettingsImportModel(useGuestCredentials bool, archiveLogs EBackupOracleLogsSettings, ) *BackupOracleSettingsImportModel`

NewBackupOracleSettingsImportModel instantiates a new BackupOracleSettingsImportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupOracleSettingsImportModelWithDefaults

`func NewBackupOracleSettingsImportModelWithDefaults() *BackupOracleSettingsImportModel`

NewBackupOracleSettingsImportModelWithDefaults instantiates a new BackupOracleSettingsImportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseGuestCredentials

`func (o *BackupOracleSettingsImportModel) GetUseGuestCredentials() bool`

GetUseGuestCredentials returns the UseGuestCredentials field if non-nil, zero value otherwise.

### GetUseGuestCredentialsOk

`func (o *BackupOracleSettingsImportModel) GetUseGuestCredentialsOk() (*bool, bool)`

GetUseGuestCredentialsOk returns a tuple with the UseGuestCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGuestCredentials

`func (o *BackupOracleSettingsImportModel) SetUseGuestCredentials(v bool)`

SetUseGuestCredentials sets UseGuestCredentials field to given value.


### GetCredentials

`func (o *BackupOracleSettingsImportModel) GetCredentials() CredentialsImportModel`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *BackupOracleSettingsImportModel) GetCredentialsOk() (*CredentialsImportModel, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *BackupOracleSettingsImportModel) SetCredentials(v CredentialsImportModel)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *BackupOracleSettingsImportModel) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetArchiveLogs

`func (o *BackupOracleSettingsImportModel) GetArchiveLogs() EBackupOracleLogsSettings`

GetArchiveLogs returns the ArchiveLogs field if non-nil, zero value otherwise.

### GetArchiveLogsOk

`func (o *BackupOracleSettingsImportModel) GetArchiveLogsOk() (*EBackupOracleLogsSettings, bool)`

GetArchiveLogsOk returns a tuple with the ArchiveLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveLogs

`func (o *BackupOracleSettingsImportModel) SetArchiveLogs(v EBackupOracleLogsSettings)`

SetArchiveLogs sets ArchiveLogs field to given value.


### GetDeleteHoursCount

`func (o *BackupOracleSettingsImportModel) GetDeleteHoursCount() int32`

GetDeleteHoursCount returns the DeleteHoursCount field if non-nil, zero value otherwise.

### GetDeleteHoursCountOk

`func (o *BackupOracleSettingsImportModel) GetDeleteHoursCountOk() (*int32, bool)`

GetDeleteHoursCountOk returns a tuple with the DeleteHoursCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteHoursCount

`func (o *BackupOracleSettingsImportModel) SetDeleteHoursCount(v int32)`

SetDeleteHoursCount sets DeleteHoursCount field to given value.

### HasDeleteHoursCount

`func (o *BackupOracleSettingsImportModel) HasDeleteHoursCount() bool`

HasDeleteHoursCount returns a boolean if a field has been set.

### GetDeleteGBsCount

`func (o *BackupOracleSettingsImportModel) GetDeleteGBsCount() int32`

GetDeleteGBsCount returns the DeleteGBsCount field if non-nil, zero value otherwise.

### GetDeleteGBsCountOk

`func (o *BackupOracleSettingsImportModel) GetDeleteGBsCountOk() (*int32, bool)`

GetDeleteGBsCountOk returns a tuple with the DeleteGBsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteGBsCount

`func (o *BackupOracleSettingsImportModel) SetDeleteGBsCount(v int32)`

SetDeleteGBsCount sets DeleteGBsCount field to given value.

### HasDeleteGBsCount

`func (o *BackupOracleSettingsImportModel) HasDeleteGBsCount() bool`

HasDeleteGBsCount returns a boolean if a field has been set.

### GetBackupLogs

`func (o *BackupOracleSettingsImportModel) GetBackupLogs() bool`

GetBackupLogs returns the BackupLogs field if non-nil, zero value otherwise.

### GetBackupLogsOk

`func (o *BackupOracleSettingsImportModel) GetBackupLogsOk() (*bool, bool)`

GetBackupLogsOk returns a tuple with the BackupLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupLogs

`func (o *BackupOracleSettingsImportModel) SetBackupLogs(v bool)`

SetBackupLogs sets BackupLogs field to given value.

### HasBackupLogs

`func (o *BackupOracleSettingsImportModel) HasBackupLogs() bool`

HasBackupLogs returns a boolean if a field has been set.

### GetBackupMinsCount

`func (o *BackupOracleSettingsImportModel) GetBackupMinsCount() int32`

GetBackupMinsCount returns the BackupMinsCount field if non-nil, zero value otherwise.

### GetBackupMinsCountOk

`func (o *BackupOracleSettingsImportModel) GetBackupMinsCountOk() (*int32, bool)`

GetBackupMinsCountOk returns a tuple with the BackupMinsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupMinsCount

`func (o *BackupOracleSettingsImportModel) SetBackupMinsCount(v int32)`

SetBackupMinsCount sets BackupMinsCount field to given value.

### HasBackupMinsCount

`func (o *BackupOracleSettingsImportModel) HasBackupMinsCount() bool`

HasBackupMinsCount returns a boolean if a field has been set.

### GetRetainLogBackups

`func (o *BackupOracleSettingsImportModel) GetRetainLogBackups() ERetainLogBackupsType`

GetRetainLogBackups returns the RetainLogBackups field if non-nil, zero value otherwise.

### GetRetainLogBackupsOk

`func (o *BackupOracleSettingsImportModel) GetRetainLogBackupsOk() (*ERetainLogBackupsType, bool)`

GetRetainLogBackupsOk returns a tuple with the RetainLogBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainLogBackups

`func (o *BackupOracleSettingsImportModel) SetRetainLogBackups(v ERetainLogBackupsType)`

SetRetainLogBackups sets RetainLogBackups field to given value.

### HasRetainLogBackups

`func (o *BackupOracleSettingsImportModel) HasRetainLogBackups() bool`

HasRetainLogBackups returns a boolean if a field has been set.

### GetKeepDaysCount

`func (o *BackupOracleSettingsImportModel) GetKeepDaysCount() int32`

GetKeepDaysCount returns the KeepDaysCount field if non-nil, zero value otherwise.

### GetKeepDaysCountOk

`func (o *BackupOracleSettingsImportModel) GetKeepDaysCountOk() (*int32, bool)`

GetKeepDaysCountOk returns a tuple with the KeepDaysCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepDaysCount

`func (o *BackupOracleSettingsImportModel) SetKeepDaysCount(v int32)`

SetKeepDaysCount sets KeepDaysCount field to given value.

### HasKeepDaysCount

`func (o *BackupOracleSettingsImportModel) HasKeepDaysCount() bool`

HasKeepDaysCount returns a boolean if a field has been set.

### GetLogShippingServers

`func (o *BackupOracleSettingsImportModel) GetLogShippingServers() BackupLogShippingServersImportModel`

GetLogShippingServers returns the LogShippingServers field if non-nil, zero value otherwise.

### GetLogShippingServersOk

`func (o *BackupOracleSettingsImportModel) GetLogShippingServersOk() (*BackupLogShippingServersImportModel, bool)`

GetLogShippingServersOk returns a tuple with the LogShippingServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogShippingServers

`func (o *BackupOracleSettingsImportModel) SetLogShippingServers(v BackupLogShippingServersImportModel)`

SetLogShippingServers sets LogShippingServers field to given value.

### HasLogShippingServers

`func (o *BackupOracleSettingsImportModel) HasLogShippingServers() bool`

HasLogShippingServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


