# LinuxCredentialsModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | Pointer to **string** | Tag used to identify the credentials record. | [optional] 
**SSHPort** | Pointer to **int32** | SSH port used to connect to a Linux server. | [optional] 
**AutoElevated** | Pointer to **bool** | If *true*, the permissions of the account are automatically elevated to the root user. | [optional] 
**AddToSudoers** | Pointer to **bool** | If *true*, the account is automatically added to the sudoers file. | [optional] 
**UseSu** | Pointer to **bool** | If *true*, the &#x60;su&#x60; command is used for Linux distributions where the &#x60;sudo&#x60; command is not available. | [optional] 
**PrivateKey** | Pointer to **string** | Private key. | [optional] 
**Passphrase** | Pointer to **string** | Passphrase that protects the private key. | [optional] 

## Methods

### NewLinuxCredentialsModelAllOf

`func NewLinuxCredentialsModelAllOf() *LinuxCredentialsModelAllOf`

NewLinuxCredentialsModelAllOf instantiates a new LinuxCredentialsModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinuxCredentialsModelAllOfWithDefaults

`func NewLinuxCredentialsModelAllOfWithDefaults() *LinuxCredentialsModelAllOf`

NewLinuxCredentialsModelAllOfWithDefaults instantiates a new LinuxCredentialsModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *LinuxCredentialsModelAllOf) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *LinuxCredentialsModelAllOf) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *LinuxCredentialsModelAllOf) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *LinuxCredentialsModelAllOf) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetSSHPort

`func (o *LinuxCredentialsModelAllOf) GetSSHPort() int32`

GetSSHPort returns the SSHPort field if non-nil, zero value otherwise.

### GetSSHPortOk

`func (o *LinuxCredentialsModelAllOf) GetSSHPortOk() (*int32, bool)`

GetSSHPortOk returns a tuple with the SSHPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSHPort

`func (o *LinuxCredentialsModelAllOf) SetSSHPort(v int32)`

SetSSHPort sets SSHPort field to given value.

### HasSSHPort

`func (o *LinuxCredentialsModelAllOf) HasSSHPort() bool`

HasSSHPort returns a boolean if a field has been set.

### GetAutoElevated

`func (o *LinuxCredentialsModelAllOf) GetAutoElevated() bool`

GetAutoElevated returns the AutoElevated field if non-nil, zero value otherwise.

### GetAutoElevatedOk

`func (o *LinuxCredentialsModelAllOf) GetAutoElevatedOk() (*bool, bool)`

GetAutoElevatedOk returns a tuple with the AutoElevated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoElevated

`func (o *LinuxCredentialsModelAllOf) SetAutoElevated(v bool)`

SetAutoElevated sets AutoElevated field to given value.

### HasAutoElevated

`func (o *LinuxCredentialsModelAllOf) HasAutoElevated() bool`

HasAutoElevated returns a boolean if a field has been set.

### GetAddToSudoers

`func (o *LinuxCredentialsModelAllOf) GetAddToSudoers() bool`

GetAddToSudoers returns the AddToSudoers field if non-nil, zero value otherwise.

### GetAddToSudoersOk

`func (o *LinuxCredentialsModelAllOf) GetAddToSudoersOk() (*bool, bool)`

GetAddToSudoersOk returns a tuple with the AddToSudoers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddToSudoers

`func (o *LinuxCredentialsModelAllOf) SetAddToSudoers(v bool)`

SetAddToSudoers sets AddToSudoers field to given value.

### HasAddToSudoers

`func (o *LinuxCredentialsModelAllOf) HasAddToSudoers() bool`

HasAddToSudoers returns a boolean if a field has been set.

### GetUseSu

`func (o *LinuxCredentialsModelAllOf) GetUseSu() bool`

GetUseSu returns the UseSu field if non-nil, zero value otherwise.

### GetUseSuOk

`func (o *LinuxCredentialsModelAllOf) GetUseSuOk() (*bool, bool)`

GetUseSuOk returns a tuple with the UseSu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSu

`func (o *LinuxCredentialsModelAllOf) SetUseSu(v bool)`

SetUseSu sets UseSu field to given value.

### HasUseSu

`func (o *LinuxCredentialsModelAllOf) HasUseSu() bool`

HasUseSu returns a boolean if a field has been set.

### GetPrivateKey

`func (o *LinuxCredentialsModelAllOf) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *LinuxCredentialsModelAllOf) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *LinuxCredentialsModelAllOf) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *LinuxCredentialsModelAllOf) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPassphrase

`func (o *LinuxCredentialsModelAllOf) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *LinuxCredentialsModelAllOf) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *LinuxCredentialsModelAllOf) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *LinuxCredentialsModelAllOf) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


