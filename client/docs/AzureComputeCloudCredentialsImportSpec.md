# AzureComputeCloudCredentialsImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ECloudCredentialsType**](ECloudCredentialsType.md) |  | 
**Description** | Pointer to **string** | Description of the cloud credentials record. | [optional] 
**Tag** | **string** | Tag used to identify the cloud credentials record. | 
**ConnectionName** | **string** | Name under which the cloud credentials record is shown in Veeam Backup &amp; Replication. | 
**ExistingAccount** | [**AzureComputeCredentialsExistingAccountSpec**](AzureComputeCredentialsExistingAccountSpec.md) |  | 

## Methods

### NewAzureComputeCloudCredentialsImportSpec

`func NewAzureComputeCloudCredentialsImportSpec(type_ ECloudCredentialsType, tag string, connectionName string, existingAccount AzureComputeCredentialsExistingAccountSpec, ) *AzureComputeCloudCredentialsImportSpec`

NewAzureComputeCloudCredentialsImportSpec instantiates a new AzureComputeCloudCredentialsImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureComputeCloudCredentialsImportSpecWithDefaults

`func NewAzureComputeCloudCredentialsImportSpecWithDefaults() *AzureComputeCloudCredentialsImportSpec`

NewAzureComputeCloudCredentialsImportSpecWithDefaults instantiates a new AzureComputeCloudCredentialsImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AzureComputeCloudCredentialsImportSpec) GetType() ECloudCredentialsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AzureComputeCloudCredentialsImportSpec) GetTypeOk() (*ECloudCredentialsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AzureComputeCloudCredentialsImportSpec) SetType(v ECloudCredentialsType)`

SetType sets Type field to given value.


### GetDescription

`func (o *AzureComputeCloudCredentialsImportSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AzureComputeCloudCredentialsImportSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AzureComputeCloudCredentialsImportSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AzureComputeCloudCredentialsImportSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTag

`func (o *AzureComputeCloudCredentialsImportSpec) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AzureComputeCloudCredentialsImportSpec) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AzureComputeCloudCredentialsImportSpec) SetTag(v string)`

SetTag sets Tag field to given value.


### GetConnectionName

`func (o *AzureComputeCloudCredentialsImportSpec) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *AzureComputeCloudCredentialsImportSpec) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *AzureComputeCloudCredentialsImportSpec) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.


### GetExistingAccount

`func (o *AzureComputeCloudCredentialsImportSpec) GetExistingAccount() AzureComputeCredentialsExistingAccountSpec`

GetExistingAccount returns the ExistingAccount field if non-nil, zero value otherwise.

### GetExistingAccountOk

`func (o *AzureComputeCloudCredentialsImportSpec) GetExistingAccountOk() (*AzureComputeCredentialsExistingAccountSpec, bool)`

GetExistingAccountOk returns a tuple with the ExistingAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingAccount

`func (o *AzureComputeCloudCredentialsImportSpec) SetExistingAccount(v AzureComputeCredentialsExistingAccountSpec)`

SetExistingAccount sets ExistingAccount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


