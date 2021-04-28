# BackupSQLSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogsProcessing** | [**ESQLLogsProcessing**](ESQLLogsProcessing.md) |  | 
**BackupMinsCount** | Pointer to **int32** | Frequency of transaction log backup, in minutes. | [optional] 
**RetainLogBackups** | Pointer to [**ERetainLogBackupsType**](ERetainLogBackupsType.md) |  | [optional] 
**KeepDaysCount** | Pointer to **int32** | Number of days to keep transaction logs in the backup repository. | [optional] 
**LogShippingServers** | Pointer to [**BackupLogShippingServersModel**](BackupLogShippingServersModel.md) |  | [optional] 

## Methods

### NewBackupSQLSettingsModel

`func NewBackupSQLSettingsModel(logsProcessing ESQLLogsProcessing, ) *BackupSQLSettingsModel`

NewBackupSQLSettingsModel instantiates a new BackupSQLSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupSQLSettingsModelWithDefaults

`func NewBackupSQLSettingsModelWithDefaults() *BackupSQLSettingsModel`

NewBackupSQLSettingsModelWithDefaults instantiates a new BackupSQLSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogsProcessing

`func (o *BackupSQLSettingsModel) GetLogsProcessing() ESQLLogsProcessing`

GetLogsProcessing returns the LogsProcessing field if non-nil, zero value otherwise.

### GetLogsProcessingOk

`func (o *BackupSQLSettingsModel) GetLogsProcessingOk() (*ESQLLogsProcessing, bool)`

GetLogsProcessingOk returns a tuple with the LogsProcessing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsProcessing

`func (o *BackupSQLSettingsModel) SetLogsProcessing(v ESQLLogsProcessing)`

SetLogsProcessing sets LogsProcessing field to given value.


### GetBackupMinsCount

`func (o *BackupSQLSettingsModel) GetBackupMinsCount() int32`

GetBackupMinsCount returns the BackupMinsCount field if non-nil, zero value otherwise.

### GetBackupMinsCountOk

`func (o *BackupSQLSettingsModel) GetBackupMinsCountOk() (*int32, bool)`

GetBackupMinsCountOk returns a tuple with the BackupMinsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupMinsCount

`func (o *BackupSQLSettingsModel) SetBackupMinsCount(v int32)`

SetBackupMinsCount sets BackupMinsCount field to given value.

### HasBackupMinsCount

`func (o *BackupSQLSettingsModel) HasBackupMinsCount() bool`

HasBackupMinsCount returns a boolean if a field has been set.

### GetRetainLogBackups

`func (o *BackupSQLSettingsModel) GetRetainLogBackups() ERetainLogBackupsType`

GetRetainLogBackups returns the RetainLogBackups field if non-nil, zero value otherwise.

### GetRetainLogBackupsOk

`func (o *BackupSQLSettingsModel) GetRetainLogBackupsOk() (*ERetainLogBackupsType, bool)`

GetRetainLogBackupsOk returns a tuple with the RetainLogBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainLogBackups

`func (o *BackupSQLSettingsModel) SetRetainLogBackups(v ERetainLogBackupsType)`

SetRetainLogBackups sets RetainLogBackups field to given value.

### HasRetainLogBackups

`func (o *BackupSQLSettingsModel) HasRetainLogBackups() bool`

HasRetainLogBackups returns a boolean if a field has been set.

### GetKeepDaysCount

`func (o *BackupSQLSettingsModel) GetKeepDaysCount() int32`

GetKeepDaysCount returns the KeepDaysCount field if non-nil, zero value otherwise.

### GetKeepDaysCountOk

`func (o *BackupSQLSettingsModel) GetKeepDaysCountOk() (*int32, bool)`

GetKeepDaysCountOk returns a tuple with the KeepDaysCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepDaysCount

`func (o *BackupSQLSettingsModel) SetKeepDaysCount(v int32)`

SetKeepDaysCount sets KeepDaysCount field to given value.

### HasKeepDaysCount

`func (o *BackupSQLSettingsModel) HasKeepDaysCount() bool`

HasKeepDaysCount returns a boolean if a field has been set.

### GetLogShippingServers

`func (o *BackupSQLSettingsModel) GetLogShippingServers() BackupLogShippingServersModel`

GetLogShippingServers returns the LogShippingServers field if non-nil, zero value otherwise.

### GetLogShippingServersOk

`func (o *BackupSQLSettingsModel) GetLogShippingServersOk() (*BackupLogShippingServersModel, bool)`

GetLogShippingServersOk returns a tuple with the LogShippingServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogShippingServers

`func (o *BackupSQLSettingsModel) SetLogShippingServers(v BackupLogShippingServersModel)`

SetLogShippingServers sets LogShippingServers field to given value.

### HasLogShippingServers

`func (o *BackupSQLSettingsModel) HasLogShippingServers() bool`

HasLogShippingServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


