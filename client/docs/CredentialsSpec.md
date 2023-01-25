# CredentialsSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | User name. | 
**Password** | Pointer to **string** | Password. | [optional] 
**Description** | Pointer to **string** | Description of the credentials record. | [optional] 
**Type** | [**ECredentialsType**](ECredentialsType.md) |  | 
**Tag** | Pointer to **string** | Tag used to identify the credentials record. | [optional] 
**SSHPort** | Pointer to **int32** | SSH port used to connect to a Linux server. | [optional] 
**AutoElevated** | Pointer to **bool** | If *true*, the permissions of the account are automatically elevated to the root user. | [optional] 
**AddToSudoers** | Pointer to **bool** | If *true*, the account is automatically added to the sudoers file. | [optional] 
**UseSu** | Pointer to **bool** | If *true*, the &#x60;su&#x60; command is used for Linux distributions where the &#x60;sudo&#x60; command is not available. | [optional] 
**PrivateKey** | Pointer to **string** | Private key. | [optional] 
**Passphrase** | Pointer to **string** | Passphrase that protects the private key. | [optional] 
**RootPassword** | Pointer to **string** | Password for the root account. | [optional] 

## Methods

### NewCredentialsSpec

`func NewCredentialsSpec(username string, type_ ECredentialsType, ) *CredentialsSpec`

NewCredentialsSpec instantiates a new CredentialsSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsSpecWithDefaults

`func NewCredentialsSpecWithDefaults() *CredentialsSpec`

NewCredentialsSpecWithDefaults instantiates a new CredentialsSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *CredentialsSpec) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CredentialsSpec) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CredentialsSpec) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *CredentialsSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CredentialsSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CredentialsSpec) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CredentialsSpec) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetDescription

`func (o *CredentialsSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CredentialsSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CredentialsSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CredentialsSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *CredentialsSpec) GetType() ECredentialsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CredentialsSpec) GetTypeOk() (*ECredentialsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CredentialsSpec) SetType(v ECredentialsType)`

SetType sets Type field to given value.


### GetTag

`func (o *CredentialsSpec) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CredentialsSpec) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CredentialsSpec) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CredentialsSpec) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetSSHPort

`func (o *CredentialsSpec) GetSSHPort() int32`

GetSSHPort returns the SSHPort field if non-nil, zero value otherwise.

### GetSSHPortOk

`func (o *CredentialsSpec) GetSSHPortOk() (*int32, bool)`

GetSSHPortOk returns a tuple with the SSHPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSHPort

`func (o *CredentialsSpec) SetSSHPort(v int32)`

SetSSHPort sets SSHPort field to given value.

### HasSSHPort

`func (o *CredentialsSpec) HasSSHPort() bool`

HasSSHPort returns a boolean if a field has been set.

### GetAutoElevated

`func (o *CredentialsSpec) GetAutoElevated() bool`

GetAutoElevated returns the AutoElevated field if non-nil, zero value otherwise.

### GetAutoElevatedOk

`func (o *CredentialsSpec) GetAutoElevatedOk() (*bool, bool)`

GetAutoElevatedOk returns a tuple with the AutoElevated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoElevated

`func (o *CredentialsSpec) SetAutoElevated(v bool)`

SetAutoElevated sets AutoElevated field to given value.

### HasAutoElevated

`func (o *CredentialsSpec) HasAutoElevated() bool`

HasAutoElevated returns a boolean if a field has been set.

### GetAddToSudoers

`func (o *CredentialsSpec) GetAddToSudoers() bool`

GetAddToSudoers returns the AddToSudoers field if non-nil, zero value otherwise.

### GetAddToSudoersOk

`func (o *CredentialsSpec) GetAddToSudoersOk() (*bool, bool)`

GetAddToSudoersOk returns a tuple with the AddToSudoers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddToSudoers

`func (o *CredentialsSpec) SetAddToSudoers(v bool)`

SetAddToSudoers sets AddToSudoers field to given value.

### HasAddToSudoers

`func (o *CredentialsSpec) HasAddToSudoers() bool`

HasAddToSudoers returns a boolean if a field has been set.

### GetUseSu

`func (o *CredentialsSpec) GetUseSu() bool`

GetUseSu returns the UseSu field if non-nil, zero value otherwise.

### GetUseSuOk

`func (o *CredentialsSpec) GetUseSuOk() (*bool, bool)`

GetUseSuOk returns a tuple with the UseSu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSu

`func (o *CredentialsSpec) SetUseSu(v bool)`

SetUseSu sets UseSu field to given value.

### HasUseSu

`func (o *CredentialsSpec) HasUseSu() bool`

HasUseSu returns a boolean if a field has been set.

### GetPrivateKey

`func (o *CredentialsSpec) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CredentialsSpec) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CredentialsSpec) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *CredentialsSpec) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPassphrase

`func (o *CredentialsSpec) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *CredentialsSpec) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *CredentialsSpec) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *CredentialsSpec) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetRootPassword

`func (o *CredentialsSpec) GetRootPassword() string`

GetRootPassword returns the RootPassword field if non-nil, zero value otherwise.

### GetRootPasswordOk

`func (o *CredentialsSpec) GetRootPasswordOk() (*string, bool)`

GetRootPasswordOk returns a tuple with the RootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPassword

`func (o *CredentialsSpec) SetRootPassword(v string)`

SetRootPassword sets RootPassword field to given value.

### HasRootPassword

`func (o *CredentialsSpec) HasRootPassword() bool`

HasRootPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


