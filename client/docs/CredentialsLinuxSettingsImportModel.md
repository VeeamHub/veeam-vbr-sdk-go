# CredentialsLinuxSettingsImportModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SSHPort** | Pointer to **int32** | SSH port used to connect to a Linux server. | [optional] 
**AutoElevated** | Pointer to **bool** | If *true*, the permissions of the account are automatically elevated to the root user. | [optional] 
**AddToSudoers** | Pointer to **bool** | If *true*, the account is automatically added to the sudoers file. | [optional] 
**UseSu** | Pointer to **bool** | If *true*, the &#x60;su&#x60; command is used for Linux distributions where the &#x60;sudo&#x60; command is not available. | [optional] 
**PrivateKey** | Pointer to **string** | Private key. | [optional] 
**Passphrase** | Pointer to **string** | Passphrase that protects the private key. | [optional] 
**RootPassword** | Pointer to **string** | Password for the root account. | [optional] 

## Methods

### NewCredentialsLinuxSettingsImportModel

`func NewCredentialsLinuxSettingsImportModel() *CredentialsLinuxSettingsImportModel`

NewCredentialsLinuxSettingsImportModel instantiates a new CredentialsLinuxSettingsImportModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsLinuxSettingsImportModelWithDefaults

`func NewCredentialsLinuxSettingsImportModelWithDefaults() *CredentialsLinuxSettingsImportModel`

NewCredentialsLinuxSettingsImportModelWithDefaults instantiates a new CredentialsLinuxSettingsImportModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSSHPort

`func (o *CredentialsLinuxSettingsImportModel) GetSSHPort() int32`

GetSSHPort returns the SSHPort field if non-nil, zero value otherwise.

### GetSSHPortOk

`func (o *CredentialsLinuxSettingsImportModel) GetSSHPortOk() (*int32, bool)`

GetSSHPortOk returns a tuple with the SSHPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSHPort

`func (o *CredentialsLinuxSettingsImportModel) SetSSHPort(v int32)`

SetSSHPort sets SSHPort field to given value.

### HasSSHPort

`func (o *CredentialsLinuxSettingsImportModel) HasSSHPort() bool`

HasSSHPort returns a boolean if a field has been set.

### GetAutoElevated

`func (o *CredentialsLinuxSettingsImportModel) GetAutoElevated() bool`

GetAutoElevated returns the AutoElevated field if non-nil, zero value otherwise.

### GetAutoElevatedOk

`func (o *CredentialsLinuxSettingsImportModel) GetAutoElevatedOk() (*bool, bool)`

GetAutoElevatedOk returns a tuple with the AutoElevated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoElevated

`func (o *CredentialsLinuxSettingsImportModel) SetAutoElevated(v bool)`

SetAutoElevated sets AutoElevated field to given value.

### HasAutoElevated

`func (o *CredentialsLinuxSettingsImportModel) HasAutoElevated() bool`

HasAutoElevated returns a boolean if a field has been set.

### GetAddToSudoers

`func (o *CredentialsLinuxSettingsImportModel) GetAddToSudoers() bool`

GetAddToSudoers returns the AddToSudoers field if non-nil, zero value otherwise.

### GetAddToSudoersOk

`func (o *CredentialsLinuxSettingsImportModel) GetAddToSudoersOk() (*bool, bool)`

GetAddToSudoersOk returns a tuple with the AddToSudoers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddToSudoers

`func (o *CredentialsLinuxSettingsImportModel) SetAddToSudoers(v bool)`

SetAddToSudoers sets AddToSudoers field to given value.

### HasAddToSudoers

`func (o *CredentialsLinuxSettingsImportModel) HasAddToSudoers() bool`

HasAddToSudoers returns a boolean if a field has been set.

### GetUseSu

`func (o *CredentialsLinuxSettingsImportModel) GetUseSu() bool`

GetUseSu returns the UseSu field if non-nil, zero value otherwise.

### GetUseSuOk

`func (o *CredentialsLinuxSettingsImportModel) GetUseSuOk() (*bool, bool)`

GetUseSuOk returns a tuple with the UseSu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSu

`func (o *CredentialsLinuxSettingsImportModel) SetUseSu(v bool)`

SetUseSu sets UseSu field to given value.

### HasUseSu

`func (o *CredentialsLinuxSettingsImportModel) HasUseSu() bool`

HasUseSu returns a boolean if a field has been set.

### GetPrivateKey

`func (o *CredentialsLinuxSettingsImportModel) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CredentialsLinuxSettingsImportModel) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CredentialsLinuxSettingsImportModel) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *CredentialsLinuxSettingsImportModel) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPassphrase

`func (o *CredentialsLinuxSettingsImportModel) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *CredentialsLinuxSettingsImportModel) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *CredentialsLinuxSettingsImportModel) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *CredentialsLinuxSettingsImportModel) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetRootPassword

`func (o *CredentialsLinuxSettingsImportModel) GetRootPassword() string`

GetRootPassword returns the RootPassword field if non-nil, zero value otherwise.

### GetRootPasswordOk

`func (o *CredentialsLinuxSettingsImportModel) GetRootPasswordOk() (*string, bool)`

GetRootPasswordOk returns a tuple with the RootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPassword

`func (o *CredentialsLinuxSettingsImportModel) SetRootPassword(v string)`

SetRootPassword sets RootPassword field to given value.

### HasRootPassword

`func (o *CredentialsLinuxSettingsImportModel) HasRootPassword() bool`

HasRootPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


