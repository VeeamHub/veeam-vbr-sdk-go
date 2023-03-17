# RepositoryAccessPermissionsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessPolicy** | [**ERepositoryAccessType**](ERepositoryAccessType.md) |  | 
**Accounts** | Pointer to **[]string** | (For *AllowExplicit* access policy) Array of accounts that have access to the backup repository. | [optional] 
**EncryptBackups** | **bool** | If *true*, Veeam Backup &amp; Replication encrypts Veeam Agent backup files stored in the backup repository. | 
**PasswordId** | Pointer to **string** | ID of the password used for encryption. | [optional] 

## Methods

### NewRepositoryAccessPermissionsModel

`func NewRepositoryAccessPermissionsModel(accessPolicy ERepositoryAccessType, encryptBackups bool, ) *RepositoryAccessPermissionsModel`

NewRepositoryAccessPermissionsModel instantiates a new RepositoryAccessPermissionsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryAccessPermissionsModelWithDefaults

`func NewRepositoryAccessPermissionsModelWithDefaults() *RepositoryAccessPermissionsModel`

NewRepositoryAccessPermissionsModelWithDefaults instantiates a new RepositoryAccessPermissionsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessPolicy

`func (o *RepositoryAccessPermissionsModel) GetAccessPolicy() ERepositoryAccessType`

GetAccessPolicy returns the AccessPolicy field if non-nil, zero value otherwise.

### GetAccessPolicyOk

`func (o *RepositoryAccessPermissionsModel) GetAccessPolicyOk() (*ERepositoryAccessType, bool)`

GetAccessPolicyOk returns a tuple with the AccessPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicy

`func (o *RepositoryAccessPermissionsModel) SetAccessPolicy(v ERepositoryAccessType)`

SetAccessPolicy sets AccessPolicy field to given value.


### GetAccounts

`func (o *RepositoryAccessPermissionsModel) GetAccounts() []string`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *RepositoryAccessPermissionsModel) GetAccountsOk() (*[]string, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *RepositoryAccessPermissionsModel) SetAccounts(v []string)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *RepositoryAccessPermissionsModel) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetEncryptBackups

`func (o *RepositoryAccessPermissionsModel) GetEncryptBackups() bool`

GetEncryptBackups returns the EncryptBackups field if non-nil, zero value otherwise.

### GetEncryptBackupsOk

`func (o *RepositoryAccessPermissionsModel) GetEncryptBackupsOk() (*bool, bool)`

GetEncryptBackupsOk returns a tuple with the EncryptBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptBackups

`func (o *RepositoryAccessPermissionsModel) SetEncryptBackups(v bool)`

SetEncryptBackups sets EncryptBackups field to given value.


### GetPasswordId

`func (o *RepositoryAccessPermissionsModel) GetPasswordId() string`

GetPasswordId returns the PasswordId field if non-nil, zero value otherwise.

### GetPasswordIdOk

`func (o *RepositoryAccessPermissionsModel) GetPasswordIdOk() (*string, bool)`

GetPasswordIdOk returns a tuple with the PasswordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordId

`func (o *RepositoryAccessPermissionsModel) SetPasswordId(v string)`

SetPasswordId sets PasswordId field to given value.

### HasPasswordId

`func (o *RepositoryAccessPermissionsModel) HasPasswordId() bool`

HasPasswordId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


