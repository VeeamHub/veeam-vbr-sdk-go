# ConfigBackupEncryptionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | If *true*, backup encryption is enabled. | 
**PasswordId** | **string** | ID of the password used for encryption. | 

## Methods

### NewConfigBackupEncryptionModel

`func NewConfigBackupEncryptionModel(isEnabled bool, passwordId string, ) *ConfigBackupEncryptionModel`

NewConfigBackupEncryptionModel instantiates a new ConfigBackupEncryptionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigBackupEncryptionModelWithDefaults

`func NewConfigBackupEncryptionModelWithDefaults() *ConfigBackupEncryptionModel`

NewConfigBackupEncryptionModelWithDefaults instantiates a new ConfigBackupEncryptionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *ConfigBackupEncryptionModel) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ConfigBackupEncryptionModel) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ConfigBackupEncryptionModel) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetPasswordId

`func (o *ConfigBackupEncryptionModel) GetPasswordId() string`

GetPasswordId returns the PasswordId field if non-nil, zero value otherwise.

### GetPasswordIdOk

`func (o *ConfigBackupEncryptionModel) GetPasswordIdOk() (*string, bool)`

GetPasswordIdOk returns a tuple with the PasswordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordId

`func (o *ConfigBackupEncryptionModel) SetPasswordId(v string)`

SetPasswordId sets PasswordId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


