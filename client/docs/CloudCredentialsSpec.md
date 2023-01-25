# CloudCredentialsSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the cloud credentials record. | [optional] 
**Type** | [**ECloudCredentialsType**](ECloudCredentialsType.md) |  | 
**AccessKey** | **string** | Access ID of a Google HMAC key. | 
**SecretKey** | **string** | Secret linked to the access ID. | 
**Tag** | Pointer to **string** | Tag used to identify the cloud credentials record. | [optional] 
**ConnectionName** | **string** | Name under which the cloud credentials record will be shown in Veeam Backup &amp; Replication. | 
**CreationMode** | [**EAzureComputeCredentialsCreationMode**](EAzureComputeCredentialsCreationMode.md) |  | 
**ExistingAccount** | Pointer to [**AzureComputeCredentialsExistingAccountSpec**](AzureComputeCredentialsExistingAccountSpec.md) |  | [optional] 
**NewAccount** | Pointer to [**AzureComputeCredentialsNewAccountSpec**](AzureComputeCredentialsNewAccountSpec.md) |  | [optional] 
**Account** | **string** | Name of the Azure storage account. | 
**SharedKey** | **string** | Shared key of the Azure storage account. | 

## Methods

### NewCloudCredentialsSpec

`func NewCloudCredentialsSpec(type_ ECloudCredentialsType, accessKey string, secretKey string, connectionName string, creationMode EAzureComputeCredentialsCreationMode, account string, sharedKey string, ) *CloudCredentialsSpec`

NewCloudCredentialsSpec instantiates a new CloudCredentialsSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCredentialsSpecWithDefaults

`func NewCloudCredentialsSpecWithDefaults() *CloudCredentialsSpec`

NewCloudCredentialsSpecWithDefaults instantiates a new CloudCredentialsSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CloudCredentialsSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudCredentialsSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudCredentialsSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudCredentialsSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *CloudCredentialsSpec) GetType() ECloudCredentialsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudCredentialsSpec) GetTypeOk() (*ECloudCredentialsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudCredentialsSpec) SetType(v ECloudCredentialsType)`

SetType sets Type field to given value.


### GetAccessKey

`func (o *CloudCredentialsSpec) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *CloudCredentialsSpec) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *CloudCredentialsSpec) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetSecretKey

`func (o *CloudCredentialsSpec) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *CloudCredentialsSpec) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *CloudCredentialsSpec) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### GetTag

`func (o *CloudCredentialsSpec) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CloudCredentialsSpec) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CloudCredentialsSpec) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CloudCredentialsSpec) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetConnectionName

`func (o *CloudCredentialsSpec) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *CloudCredentialsSpec) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *CloudCredentialsSpec) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.


### GetCreationMode

`func (o *CloudCredentialsSpec) GetCreationMode() EAzureComputeCredentialsCreationMode`

GetCreationMode returns the CreationMode field if non-nil, zero value otherwise.

### GetCreationModeOk

`func (o *CloudCredentialsSpec) GetCreationModeOk() (*EAzureComputeCredentialsCreationMode, bool)`

GetCreationModeOk returns a tuple with the CreationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationMode

`func (o *CloudCredentialsSpec) SetCreationMode(v EAzureComputeCredentialsCreationMode)`

SetCreationMode sets CreationMode field to given value.


### GetExistingAccount

`func (o *CloudCredentialsSpec) GetExistingAccount() AzureComputeCredentialsExistingAccountSpec`

GetExistingAccount returns the ExistingAccount field if non-nil, zero value otherwise.

### GetExistingAccountOk

`func (o *CloudCredentialsSpec) GetExistingAccountOk() (*AzureComputeCredentialsExistingAccountSpec, bool)`

GetExistingAccountOk returns a tuple with the ExistingAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingAccount

`func (o *CloudCredentialsSpec) SetExistingAccount(v AzureComputeCredentialsExistingAccountSpec)`

SetExistingAccount sets ExistingAccount field to given value.

### HasExistingAccount

`func (o *CloudCredentialsSpec) HasExistingAccount() bool`

HasExistingAccount returns a boolean if a field has been set.

### GetNewAccount

`func (o *CloudCredentialsSpec) GetNewAccount() AzureComputeCredentialsNewAccountSpec`

GetNewAccount returns the NewAccount field if non-nil, zero value otherwise.

### GetNewAccountOk

`func (o *CloudCredentialsSpec) GetNewAccountOk() (*AzureComputeCredentialsNewAccountSpec, bool)`

GetNewAccountOk returns a tuple with the NewAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAccount

`func (o *CloudCredentialsSpec) SetNewAccount(v AzureComputeCredentialsNewAccountSpec)`

SetNewAccount sets NewAccount field to given value.

### HasNewAccount

`func (o *CloudCredentialsSpec) HasNewAccount() bool`

HasNewAccount returns a boolean if a field has been set.

### GetAccount

`func (o *CloudCredentialsSpec) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CloudCredentialsSpec) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CloudCredentialsSpec) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetSharedKey

`func (o *CloudCredentialsSpec) GetSharedKey() string`

GetSharedKey returns the SharedKey field if non-nil, zero value otherwise.

### GetSharedKeyOk

`func (o *CloudCredentialsSpec) GetSharedKeyOk() (*string, bool)`

GetSharedKeyOk returns a tuple with the SharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedKey

`func (o *CloudCredentialsSpec) SetSharedKey(v string)`

SetSharedKey sets SharedKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


