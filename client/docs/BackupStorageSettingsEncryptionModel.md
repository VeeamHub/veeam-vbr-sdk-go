# BackupStorageSettingsEncryptionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, the content of backup files is encrypted. | 
**EncryptionPasswordIdOrNull** | Pointer to **string** | ID of the password used for encryption. The value is *null* for exported objects. | [optional] 
**EncryptionPasswordTag** | Pointer to **string** | Tag used to identify the password. | [optional] 

## Methods

### NewBackupStorageSettingsEncryptionModel

`func NewBackupStorageSettingsEncryptionModel(isEnabled bool, ) *BackupStorageSettingsEncryptionModel`

NewBackupStorageSettingsEncryptionModel instantiates a new BackupStorageSettingsEncryptionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupStorageSettingsEncryptionModelWithDefaults

`func NewBackupStorageSettingsEncryptionModelWithDefaults() *BackupStorageSettingsEncryptionModel`

NewBackupStorageSettingsEncryptionModelWithDefaults instantiates a new BackupStorageSettingsEncryptionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *BackupStorageSettingsEncryptionModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *BackupStorageSettingsEncryptionModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *BackupStorageSettingsEncryptionModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetEncryptionPasswordIdOrNull

`func (o *BackupStorageSettingsEncryptionModel) GetEncryptionPasswordIdOrNull() string`

GetEncryptionPasswordIdOrNull returns the EncryptionPasswordIdOrNull field if non-nil, zero value otherwise.

### GetEncryptionPasswordIdOrNullOk

`func (o *BackupStorageSettingsEncryptionModel) GetEncryptionPasswordIdOrNullOk() (*string, bool)`

GetEncryptionPasswordIdOrNullOk returns a tuple with the EncryptionPasswordIdOrNull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionPasswordIdOrNull

`func (o *BackupStorageSettingsEncryptionModel) SetEncryptionPasswordIdOrNull(v string)`

SetEncryptionPasswordIdOrNull sets EncryptionPasswordIdOrNull field to given value.

### HasEncryptionPasswordIdOrNull

`func (o *BackupStorageSettingsEncryptionModel) HasEncryptionPasswordIdOrNull() bool`

HasEncryptionPasswordIdOrNull returns a boolean if a field has been set.

### GetEncryptionPasswordTag

`func (o *BackupStorageSettingsEncryptionModel) GetEncryptionPasswordTag() string`

GetEncryptionPasswordTag returns the EncryptionPasswordTag field if non-nil, zero value otherwise.

### GetEncryptionPasswordTagOk

`func (o *BackupStorageSettingsEncryptionModel) GetEncryptionPasswordTagOk() (*string, bool)`

GetEncryptionPasswordTagOk returns a tuple with the EncryptionPasswordTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionPasswordTag

`func (o *BackupStorageSettingsEncryptionModel) SetEncryptionPasswordTag(v string)`

SetEncryptionPasswordTag sets EncryptionPasswordTag field to given value.

### HasEncryptionPasswordTag

`func (o *BackupStorageSettingsEncryptionModel) HasEncryptionPasswordTag() bool`

HasEncryptionPasswordTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


