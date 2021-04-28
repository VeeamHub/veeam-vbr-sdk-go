# BackupSQLSettingsImportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogsProcessing** | [**ESQLLogsProcessing**](ESQLLogsProcessing.md) |  | 
**BackupMinsCount** | Pointer to **int32** | Frequency of transaction log backup, in minutes. | [optional] 
**RetainLogBackups** | Pointer to [**ERetainLogBackupsType**](ERetainLogBackupsType.md) |  | [optional] 
**KeepDaysCount** | Pointer to **int32** | Number of days to keep transaction logs in the backup repository. | [optional] 
**LogShippingServers** | Pointer to [**BackupLogShippingServersImportModel**](BackupLogShippingServersImportModel.md) |  | [optional] 

## Methods

### NewBackupSQLSettingsImportModel

`func NewBackupSQLSettingsImportModel(logsProcessing ESQLLogsProcessing, ) *BackupSQLSettingsImportModel`

NewBackupSQLSettingsImportModel instantiates a new BackupSQLSettingsImportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupSQLSettingsImportModelWithDefaults

`func NewBackupSQLSettingsImportModelWithDefaults() *BackupSQLSettingsImportModel`

NewBackupSQLSettingsImportModelWithDefaults instantiates a new BackupSQLSettingsImportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogsProcessing

`func (o *BackupSQLSettingsImportModel) GetLogsProcessing() ESQLLogsProcessing`

GetLogsProcessing returns the LogsProcessing field if non-nil, zero value otherwise.

### GetLogsProcessingOk

`func (o *BackupSQLSettingsImportModel) GetLogsProcessingOk() (*ESQLLogsProcessing, bool)`

GetLogsProcessingOk returns a tuple with the LogsProcessing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsProcessing

`func (o *BackupSQLSettingsImportModel) SetLogsProcessing(v ESQLLogsProcessing)`

SetLogsProcessing sets LogsProcessing field to given value.


### GetBackupMinsCount

`func (o *BackupSQLSettingsImportModel) GetBackupMinsCount() int32`

GetBackupMinsCount returns the BackupMinsCount field if non-nil, zero value otherwise.

### GetBackupMinsCountOk

`func (o *BackupSQLSettingsImportModel) GetBackupMinsCountOk() (*int32, bool)`

GetBackupMinsCountOk returns a tuple with the BackupMinsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupMinsCount

`func (o *BackupSQLSettingsImportModel) SetBackupMinsCount(v int32)`

SetBackupMinsCount sets BackupMinsCount field to given value.

### HasBackupMinsCount

`func (o *BackupSQLSettingsImportModel) HasBackupMinsCount() bool`

HasBackupMinsCount returns a boolean if a field has been set.

### GetRetainLogBackups

`func (o *BackupSQLSettingsImportModel) GetRetainLogBackups() ERetainLogBackupsType`

GetRetainLogBackups returns the RetainLogBackups field if non-nil, zero value otherwise.

### GetRetainLogBackupsOk

`func (o *BackupSQLSettingsImportModel) GetRetainLogBackupsOk() (*ERetainLogBackupsType, bool)`

GetRetainLogBackupsOk returns a tuple with the RetainLogBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainLogBackups

`func (o *BackupSQLSettingsImportModel) SetRetainLogBackups(v ERetainLogBackupsType)`

SetRetainLogBackups sets RetainLogBackups field to given value.

### HasRetainLogBackups

`func (o *BackupSQLSettingsImportModel) HasRetainLogBackups() bool`

HasRetainLogBackups returns a boolean if a field has been set.

### GetKeepDaysCount

`func (o *BackupSQLSettingsImportModel) GetKeepDaysCount() int32`

GetKeepDaysCount returns the KeepDaysCount field if non-nil, zero value otherwise.

### GetKeepDaysCountOk

`func (o *BackupSQLSettingsImportModel) GetKeepDaysCountOk() (*int32, bool)`

GetKeepDaysCountOk returns a tuple with the KeepDaysCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepDaysCount

`func (o *BackupSQLSettingsImportModel) SetKeepDaysCount(v int32)`

SetKeepDaysCount sets KeepDaysCount field to given value.

### HasKeepDaysCount

`func (o *BackupSQLSettingsImportModel) HasKeepDaysCount() bool`

HasKeepDaysCount returns a boolean if a field has been set.

### GetLogShippingServers

`func (o *BackupSQLSettingsImportModel) GetLogShippingServers() BackupLogShippingServersImportModel`

GetLogShippingServers returns the LogShippingServers field if non-nil, zero value otherwise.

### GetLogShippingServersOk

`func (o *BackupSQLSettingsImportModel) GetLogShippingServersOk() (*BackupLogShippingServersImportModel, bool)`

GetLogShippingServersOk returns a tuple with the LogShippingServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogShippingServers

`func (o *BackupSQLSettingsImportModel) SetLogShippingServers(v BackupLogShippingServersImportModel)`

SetLogShippingServers sets LogShippingServers field to given value.

### HasLogShippingServers

`func (o *BackupSQLSettingsImportModel) HasLogShippingServers() bool`

HasLogShippingServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


