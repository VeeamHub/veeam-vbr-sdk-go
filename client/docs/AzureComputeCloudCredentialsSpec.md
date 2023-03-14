# AzureComputeCloudCredentialsSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionName** | **string** | Name under which the cloud credentials record will be shown in Veeam Backup &amp; Replication. | 
**CreationMode** | [**EAzureComputeCredentialsCreationMode**](EAzureComputeCredentialsCreationMode.md) |  | 
**ExistingAccount** | Pointer to [**AzureComputeCredentialsExistingAccountSpec**](AzureComputeCredentialsExistingAccountSpec.md) |  | [optional] 
**NewAccount** | Pointer to [**AzureComputeCredentialsNewAccountSpec**](AzureComputeCredentialsNewAccountSpec.md) |  | [optional] 
**Tag** | Pointer to **string** | Tag used to identify the cloud credentials record. | [optional] 

## Methods

### NewAzureComputeCloudCredentialsSpec

`func NewAzureComputeCloudCredentialsSpec(connectionName string, creationMode EAzureComputeCredentialsCreationMode, ) *AzureComputeCloudCredentialsSpec`

NewAzureComputeCloudCredentialsSpec instantiates a new AzureComputeCloudCredentialsSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureComputeCloudCredentialsSpecWithDefaults

`func NewAzureComputeCloudCredentialsSpecWithDefaults() *AzureComputeCloudCredentialsSpec`

NewAzureComputeCloudCredentialsSpecWithDefaults instantiates a new AzureComputeCloudCredentialsSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionName

`func (o *AzureComputeCloudCredentialsSpec) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *AzureComputeCloudCredentialsSpec) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *AzureComputeCloudCredentialsSpec) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.


### GetCreationMode

`func (o *AzureComputeCloudCredentialsSpec) GetCreationMode() EAzureComputeCredentialsCreationMode`

GetCreationMode returns the CreationMode field if non-nil, zero value otherwise.

### GetCreationModeOk

`func (o *AzureComputeCloudCredentialsSpec) GetCreationModeOk() (*EAzureComputeCredentialsCreationMode, bool)`

GetCreationModeOk returns a tuple with the CreationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationMode

`func (o *AzureComputeCloudCredentialsSpec) SetCreationMode(v EAzureComputeCredentialsCreationMode)`

SetCreationMode sets CreationMode field to given value.


### GetExistingAccount

`func (o *AzureComputeCloudCredentialsSpec) GetExistingAccount() AzureComputeCredentialsExistingAccountSpec`

GetExistingAccount returns the ExistingAccount field if non-nil, zero value otherwise.

### GetExistingAccountOk

`func (o *AzureComputeCloudCredentialsSpec) GetExistingAccountOk() (*AzureComputeCredentialsExistingAccountSpec, bool)`

GetExistingAccountOk returns a tuple with the ExistingAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingAccount

`func (o *AzureComputeCloudCredentialsSpec) SetExistingAccount(v AzureComputeCredentialsExistingAccountSpec)`

SetExistingAccount sets ExistingAccount field to given value.

### HasExistingAccount

`func (o *AzureComputeCloudCredentialsSpec) HasExistingAccount() bool`

HasExistingAccount returns a boolean if a field has been set.

### GetNewAccount

`func (o *AzureComputeCloudCredentialsSpec) GetNewAccount() AzureComputeCredentialsNewAccountSpec`

GetNewAccount returns the NewAccount field if non-nil, zero value otherwise.

### GetNewAccountOk

`func (o *AzureComputeCloudCredentialsSpec) GetNewAccountOk() (*AzureComputeCredentialsNewAccountSpec, bool)`

GetNewAccountOk returns a tuple with the NewAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAccount

`func (o *AzureComputeCloudCredentialsSpec) SetNewAccount(v AzureComputeCredentialsNewAccountSpec)`

SetNewAccount sets NewAccount field to given value.

### HasNewAccount

`func (o *AzureComputeCloudCredentialsSpec) HasNewAccount() bool`

HasNewAccount returns a boolean if a field has been set.

### GetTag

`func (o *AzureComputeCloudCredentialsSpec) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AzureComputeCloudCredentialsSpec) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AzureComputeCloudCredentialsSpec) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *AzureComputeCloudCredentialsSpec) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


