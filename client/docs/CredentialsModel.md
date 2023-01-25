# CredentialsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the credentials record. | 
**Username** | **string** | User name. | 
**Description** | **string** | Description of the credentials record. | 
**Type** | [**ECredentialsType**](ECredentialsType.md) |  | 
**CreationTime** | **time.Time** | Date and time when the credentials were created. | 
**Tag** | Pointer to **string** | Tag used to identify the credentials record. | [optional] 
**SSHPort** | Pointer to **int32** | SSH port used to connect to a Linux server. | [optional] 
**AutoElevated** | Pointer to **bool** | If *true*, the permissions of the account are automatically elevated to the root user. | [optional] 
**AddToSudoers** | Pointer to **bool** | If *true*, the account is automatically added to the sudoers file. | [optional] 
**UseSu** | Pointer to **bool** | If *true*, the &#x60;su&#x60; command is used for Linux distributions where the &#x60;sudo&#x60; command is not available. | [optional] 
**PrivateKey** | Pointer to **string** | Private key. | [optional] 
**Passphrase** | Pointer to **string** | Passphrase that protects the private key. | [optional] 

## Methods

### NewCredentialsModel

`func NewCredentialsModel(id string, username string, description string, type_ ECredentialsType, creationTime time.Time, ) *CredentialsModel`

NewCredentialsModel instantiates a new CredentialsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsModelWithDefaults

`func NewCredentialsModelWithDefaults() *CredentialsModel`

NewCredentialsModelWithDefaults instantiates a new CredentialsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CredentialsModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CredentialsModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CredentialsModel) SetId(v string)`

SetId sets Id field to given value.


### GetUsername

`func (o *CredentialsModel) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CredentialsModel) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CredentialsModel) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetDescription

`func (o *CredentialsModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CredentialsModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CredentialsModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *CredentialsModel) GetType() ECredentialsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CredentialsModel) GetTypeOk() (*ECredentialsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CredentialsModel) SetType(v ECredentialsType)`

SetType sets Type field to given value.


### GetCreationTime

`func (o *CredentialsModel) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CredentialsModel) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CredentialsModel) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetTag

`func (o *CredentialsModel) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CredentialsModel) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CredentialsModel) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CredentialsModel) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetSSHPort

`func (o *CredentialsModel) GetSSHPort() int32`

GetSSHPort returns the SSHPort field if non-nil, zero value otherwise.

### GetSSHPortOk

`func (o *CredentialsModel) GetSSHPortOk() (*int32, bool)`

GetSSHPortOk returns a tuple with the SSHPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSHPort

`func (o *CredentialsModel) SetSSHPort(v int32)`

SetSSHPort sets SSHPort field to given value.

### HasSSHPort

`func (o *CredentialsModel) HasSSHPort() bool`

HasSSHPort returns a boolean if a field has been set.

### GetAutoElevated

`func (o *CredentialsModel) GetAutoElevated() bool`

GetAutoElevated returns the AutoElevated field if non-nil, zero value otherwise.

### GetAutoElevatedOk

`func (o *CredentialsModel) GetAutoElevatedOk() (*bool, bool)`

GetAutoElevatedOk returns a tuple with the AutoElevated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoElevated

`func (o *CredentialsModel) SetAutoElevated(v bool)`

SetAutoElevated sets AutoElevated field to given value.

### HasAutoElevated

`func (o *CredentialsModel) HasAutoElevated() bool`

HasAutoElevated returns a boolean if a field has been set.

### GetAddToSudoers

`func (o *CredentialsModel) GetAddToSudoers() bool`

GetAddToSudoers returns the AddToSudoers field if non-nil, zero value otherwise.

### GetAddToSudoersOk

`func (o *CredentialsModel) GetAddToSudoersOk() (*bool, bool)`

GetAddToSudoersOk returns a tuple with the AddToSudoers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddToSudoers

`func (o *CredentialsModel) SetAddToSudoers(v bool)`

SetAddToSudoers sets AddToSudoers field to given value.

### HasAddToSudoers

`func (o *CredentialsModel) HasAddToSudoers() bool`

HasAddToSudoers returns a boolean if a field has been set.

### GetUseSu

`func (o *CredentialsModel) GetUseSu() bool`

GetUseSu returns the UseSu field if non-nil, zero value otherwise.

### GetUseSuOk

`func (o *CredentialsModel) GetUseSuOk() (*bool, bool)`

GetUseSuOk returns a tuple with the UseSu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSu

`func (o *CredentialsModel) SetUseSu(v bool)`

SetUseSu sets UseSu field to given value.

### HasUseSu

`func (o *CredentialsModel) HasUseSu() bool`

HasUseSu returns a boolean if a field has been set.

### GetPrivateKey

`func (o *CredentialsModel) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CredentialsModel) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CredentialsModel) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *CredentialsModel) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPassphrase

`func (o *CredentialsModel) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *CredentialsModel) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *CredentialsModel) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *CredentialsModel) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


