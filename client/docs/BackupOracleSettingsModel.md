# BackupOracleSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseGuestCredentials** | **bool** | If *true*, Veeam Backup &amp; Replication uses credentials specified in the guest processing settings. | 
**CredentialsId** | Pointer to **string** | ID of the credentials record that is used if &#x60;useGuestCredentials&#x60; is *false*. | [optional] 
**ArchiveLogs** | [**EBackupOracleLogsSettings**](EBackupOracleLogsSettings.md) |  | 
**DeleteHoursCount** | Pointer to **int32** | Time period in hours to keep archived logs. This parameter should be specified if the &#x60;EBackupOracleLogsSettings&#x60; value is *deleteExpiredHours*. | [optional] 
**DeleteGBsCount** | Pointer to **int32** | Maximum size for archived logs in GB. This parameter should be specified if the &#x60;EBackupOracleLogsSettings&#x60; value is *deleteExpiredGBs*. | [optional] 
**BackupLogs** | Pointer to **bool** | If *true*, archived logs are backed up. | [optional] 
**BackupMinsCount** | Pointer to **int32** | Frequency of archived log backup, in minutes. | [optional] 
**RetainLogBackups** | Pointer to [**ERetainLogBackupsType**](ERetainLogBackupsType.md) |  | [optional] 
**KeepDaysCount** | Pointer to **int32** | Number of days to keep archived logs. | [optional] 
**LogShippingServers** | Pointer to [**BackupLogShippingServersModel**](BackupLogShippingServersModel.md) |  | [optional] 

## Methods

### NewBackupOracleSettingsModel

`func NewBackupOracleSettingsModel(useGuestCredentials bool, archiveLogs EBackupOracleLogsSettings, ) *BackupOracleSettingsModel`

NewBackupOracleSettingsModel instantiates a new BackupOracleSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupOracleSettingsModelWithDefaults

`func NewBackupOracleSettingsModelWithDefaults() *BackupOracleSettingsModel`

NewBackupOracleSettingsModelWithDefaults instantiates a new BackupOracleSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseGuestCredentials

`func (o *BackupOracleSettingsModel) GetUseGuestCredentials() bool`

GetUseGuestCredentials returns the UseGuestCredentials field if non-nil, zero value otherwise.

### GetUseGuestCredentialsOk

`func (o *BackupOracleSettingsModel) GetUseGuestCredentialsOk() (*bool, bool)`

GetUseGuestCredentialsOk returns a tuple with the UseGuestCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGuestCredentials

`func (o *BackupOracleSettingsModel) SetUseGuestCredentials(v bool)`

SetUseGuestCredentials sets UseGuestCredentials field to given value.


### GetCredentialsId

`func (o *BackupOracleSettingsModel) GetCredentialsId() string`

GetCredentialsId returns the CredentialsId field if non-nil, zero value otherwise.

### GetCredentialsIdOk

`func (o *BackupOracleSettingsModel) GetCredentialsIdOk() (*string, bool)`

GetCredentialsIdOk returns a tuple with the CredentialsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsId

`func (o *BackupOracleSettingsModel) SetCredentialsId(v string)`

SetCredentialsId sets CredentialsId field to given value.

### HasCredentialsId

`func (o *BackupOracleSettingsModel) HasCredentialsId() bool`

HasCredentialsId returns a boolean if a field has been set.

### GetArchiveLogs

`func (o *BackupOracleSettingsModel) GetArchiveLogs() EBackupOracleLogsSettings`

GetArchiveLogs returns the ArchiveLogs field if non-nil, zero value otherwise.

### GetArchiveLogsOk

`func (o *BackupOracleSettingsModel) GetArchiveLogsOk() (*EBackupOracleLogsSettings, bool)`

GetArchiveLogsOk returns a tuple with the ArchiveLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveLogs

`func (o *BackupOracleSettingsModel) SetArchiveLogs(v EBackupOracleLogsSettings)`

SetArchiveLogs sets ArchiveLogs field to given value.


### GetDeleteHoursCount

`func (o *BackupOracleSettingsModel) GetDeleteHoursCount() int32`

GetDeleteHoursCount returns the DeleteHoursCount field if non-nil, zero value otherwise.

### GetDeleteHoursCountOk

`func (o *BackupOracleSettingsModel) GetDeleteHoursCountOk() (*int32, bool)`

GetDeleteHoursCountOk returns a tuple with the DeleteHoursCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteHoursCount

`func (o *BackupOracleSettingsModel) SetDeleteHoursCount(v int32)`

SetDeleteHoursCount sets DeleteHoursCount field to given value.

### HasDeleteHoursCount

`func (o *BackupOracleSettingsModel) HasDeleteHoursCount() bool`

HasDeleteHoursCount returns a boolean if a field has been set.

### GetDeleteGBsCount

`func (o *BackupOracleSettingsModel) GetDeleteGBsCount() int32`

GetDeleteGBsCount returns the DeleteGBsCount field if non-nil, zero value otherwise.

### GetDeleteGBsCountOk

`func (o *BackupOracleSettingsModel) GetDeleteGBsCountOk() (*int32, bool)`

GetDeleteGBsCountOk returns a tuple with the DeleteGBsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteGBsCount

`func (o *BackupOracleSettingsModel) SetDeleteGBsCount(v int32)`

SetDeleteGBsCount sets DeleteGBsCount field to given value.

### HasDeleteGBsCount

`func (o *BackupOracleSettingsModel) HasDeleteGBsCount() bool`

HasDeleteGBsCount returns a boolean if a field has been set.

### GetBackupLogs

`func (o *BackupOracleSettingsModel) GetBackupLogs() bool`

GetBackupLogs returns the BackupLogs field if non-nil, zero value otherwise.

### GetBackupLogsOk

`func (o *BackupOracleSettingsModel) GetBackupLogsOk() (*bool, bool)`

GetBackupLogsOk returns a tuple with the BackupLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupLogs

`func (o *BackupOracleSettingsModel) SetBackupLogs(v bool)`

SetBackupLogs sets BackupLogs field to given value.

### HasBackupLogs

`func (o *BackupOracleSettingsModel) HasBackupLogs() bool`

HasBackupLogs returns a boolean if a field has been set.

### GetBackupMinsCount

`func (o *BackupOracleSettingsModel) GetBackupMinsCount() int32`

GetBackupMinsCount returns the BackupMinsCount field if non-nil, zero value otherwise.

### GetBackupMinsCountOk

`func (o *BackupOracleSettingsModel) GetBackupMinsCountOk() (*int32, bool)`

GetBackupMinsCountOk returns a tuple with the BackupMinsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupMinsCount

`func (o *BackupOracleSettingsModel) SetBackupMinsCount(v int32)`

SetBackupMinsCount sets BackupMinsCount field to given value.

### HasBackupMinsCount

`func (o *BackupOracleSettingsModel) HasBackupMinsCount() bool`

HasBackupMinsCount returns a boolean if a field has been set.

### GetRetainLogBackups

`func (o *BackupOracleSettingsModel) GetRetainLogBackups() ERetainLogBackupsType`

GetRetainLogBackups returns the RetainLogBackups field if non-nil, zero value otherwise.

### GetRetainLogBackupsOk

`func (o *BackupOracleSettingsModel) GetRetainLogBackupsOk() (*ERetainLogBackupsType, bool)`

GetRetainLogBackupsOk returns a tuple with the RetainLogBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainLogBackups

`func (o *BackupOracleSettingsModel) SetRetainLogBackups(v ERetainLogBackupsType)`

SetRetainLogBackups sets RetainLogBackups field to given value.

### HasRetainLogBackups

`func (o *BackupOracleSettingsModel) HasRetainLogBackups() bool`

HasRetainLogBackups returns a boolean if a field has been set.

### GetKeepDaysCount

`func (o *BackupOracleSettingsModel) GetKeepDaysCount() int32`

GetKeepDaysCount returns the KeepDaysCount field if non-nil, zero value otherwise.

### GetKeepDaysCountOk

`func (o *BackupOracleSettingsModel) GetKeepDaysCountOk() (*int32, bool)`

GetKeepDaysCountOk returns a tuple with the KeepDaysCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepDaysCount

`func (o *BackupOracleSettingsModel) SetKeepDaysCount(v int32)`

SetKeepDaysCount sets KeepDaysCount field to given value.

### HasKeepDaysCount

`func (o *BackupOracleSettingsModel) HasKeepDaysCount() bool`

HasKeepDaysCount returns a boolean if a field has been set.

### GetLogShippingServers

`func (o *BackupOracleSettingsModel) GetLogShippingServers() BackupLogShippingServersModel`

GetLogShippingServers returns the LogShippingServers field if non-nil, zero value otherwise.

### GetLogShippingServersOk

`func (o *BackupOracleSettingsModel) GetLogShippingServersOk() (*BackupLogShippingServersModel, bool)`

GetLogShippingServersOk returns a tuple with the LogShippingServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogShippingServers

`func (o *BackupOracleSettingsModel) SetLogShippingServers(v BackupLogShippingServersModel)`

SetLogShippingServers sets LogShippingServers field to given value.

### HasLogShippingServers

`func (o *BackupOracleSettingsModel) HasLogShippingServers() bool`

HasLogShippingServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


