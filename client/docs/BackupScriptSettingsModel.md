# BackupScriptSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScriptProcessingMode** | [**EBackupScriptProcessingMode**](EBackupScriptProcessingMode.md) |  | 
**WindowsScripts** | Pointer to [**BackupWindowsScriptModel**](BackupWindowsScriptModel.md) |  | [optional] 
**LinuxScripts** | Pointer to [**BackupLinuxScriptModel**](BackupLinuxScriptModel.md) |  | [optional] 

## Methods

### NewBackupScriptSettingsModel

`func NewBackupScriptSettingsModel(scriptProcessingMode EBackupScriptProcessingMode, ) *BackupScriptSettingsModel`

NewBackupScriptSettingsModel instantiates a new BackupScriptSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupScriptSettingsModelWithDefaults

`func NewBackupScriptSettingsModelWithDefaults() *BackupScriptSettingsModel`

NewBackupScriptSettingsModelWithDefaults instantiates a new BackupScriptSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScriptProcessingMode

`func (o *BackupScriptSettingsModel) GetScriptProcessingMode() EBackupScriptProcessingMode`

GetScriptProcessingMode returns the ScriptProcessingMode field if non-nil, zero value otherwise.

### GetScriptProcessingModeOk

`func (o *BackupScriptSettingsModel) GetScriptProcessingModeOk() (*EBackupScriptProcessingMode, bool)`

GetScriptProcessingModeOk returns a tuple with the ScriptProcessingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptProcessingMode

`func (o *BackupScriptSettingsModel) SetScriptProcessingMode(v EBackupScriptProcessingMode)`

SetScriptProcessingMode sets ScriptProcessingMode field to given value.


### GetWindowsScripts

`func (o *BackupScriptSettingsModel) GetWindowsScripts() BackupWindowsScriptModel`

GetWindowsScripts returns the WindowsScripts field if non-nil, zero value otherwise.

### GetWindowsScriptsOk

`func (o *BackupScriptSettingsModel) GetWindowsScriptsOk() (*BackupWindowsScriptModel, bool)`

GetWindowsScriptsOk returns a tuple with the WindowsScripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsScripts

`func (o *BackupScriptSettingsModel) SetWindowsScripts(v BackupWindowsScriptModel)`

SetWindowsScripts sets WindowsScripts field to given value.

### HasWindowsScripts

`func (o *BackupScriptSettingsModel) HasWindowsScripts() bool`

HasWindowsScripts returns a boolean if a field has been set.

### GetLinuxScripts

`func (o *BackupScriptSettingsModel) GetLinuxScripts() BackupLinuxScriptModel`

GetLinuxScripts returns the LinuxScripts field if non-nil, zero value otherwise.

### GetLinuxScriptsOk

`func (o *BackupScriptSettingsModel) GetLinuxScriptsOk() (*BackupLinuxScriptModel, bool)`

GetLinuxScriptsOk returns a tuple with the LinuxScripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxScripts

`func (o *BackupScriptSettingsModel) SetLinuxScripts(v BackupLinuxScriptModel)`

SetLinuxScripts sets LinuxScripts field to given value.

### HasLinuxScripts

`func (o *BackupScriptSettingsModel) HasLinuxScripts() bool`

HasLinuxScripts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


