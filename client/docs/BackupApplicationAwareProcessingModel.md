# BackupApplicationAwareProcessingModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, application-aware processing is enabled. | 
**AppSettings** | Pointer to [**[]BackupApplicationSettingsModel**](BackupApplicationSettingsModel.md) | Array of VMware vSphere objects and their application settings. | [optional] 

## Methods

### NewBackupApplicationAwareProcessingModel

`func NewBackupApplicationAwareProcessingModel(isEnabled bool, ) *BackupApplicationAwareProcessingModel`

NewBackupApplicationAwareProcessingModel instantiates a new BackupApplicationAwareProcessingModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupApplicationAwareProcessingModelWithDefaults

`func NewBackupApplicationAwareProcessingModelWithDefaults() *BackupApplicationAwareProcessingModel`

NewBackupApplicationAwareProcessingModelWithDefaults instantiates a new BackupApplicationAwareProcessingModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *BackupApplicationAwareProcessingModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *BackupApplicationAwareProcessingModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *BackupApplicationAwareProcessingModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetAppSettings

`func (o *BackupApplicationAwareProcessingModel) GetAppSettings() []BackupApplicationSettingsModel`

GetAppSettings returns the AppSettings field if non-nil, zero value otherwise.

### GetAppSettingsOk

`func (o *BackupApplicationAwareProcessingModel) GetAppSettingsOk() (*[]BackupApplicationSettingsModel, bool)`

GetAppSettingsOk returns a tuple with the AppSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSettings

`func (o *BackupApplicationAwareProcessingModel) SetAppSettings(v []BackupApplicationSettingsModel)`

SetAppSettings sets AppSettings field to given value.

### HasAppSettings

`func (o *BackupApplicationAwareProcessingModel) HasAppSettings() bool`

HasAppSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


