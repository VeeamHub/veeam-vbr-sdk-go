# AzureStorageCloudCredentialsImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ECloudCredentialsType**](ECloudCredentialsType.md) |  | 
**Description** | Pointer to **string** | Description of the cloud credentials record. | [optional] 
**Tag** | **string** | Tag used to identify the cloud credentials record. | 
**Account** | **string** | Name of the Azure storage account. | 
**SharedKey** | **string** | Shared key of the Azure storage account. | 

## Methods

### NewAzureStorageCloudCredentialsImportSpec

`func NewAzureStorageCloudCredentialsImportSpec(type_ ECloudCredentialsType, tag string, account string, sharedKey string, ) *AzureStorageCloudCredentialsImportSpec`

NewAzureStorageCloudCredentialsImportSpec instantiates a new AzureStorageCloudCredentialsImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureStorageCloudCredentialsImportSpecWithDefaults

`func NewAzureStorageCloudCredentialsImportSpecWithDefaults() *AzureStorageCloudCredentialsImportSpec`

NewAzureStorageCloudCredentialsImportSpecWithDefaults instantiates a new AzureStorageCloudCredentialsImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AzureStorageCloudCredentialsImportSpec) GetType() ECloudCredentialsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AzureStorageCloudCredentialsImportSpec) GetTypeOk() (*ECloudCredentialsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AzureStorageCloudCredentialsImportSpec) SetType(v ECloudCredentialsType)`

SetType sets Type field to given value.


### GetDescription

`func (o *AzureStorageCloudCredentialsImportSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AzureStorageCloudCredentialsImportSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AzureStorageCloudCredentialsImportSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AzureStorageCloudCredentialsImportSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTag

`func (o *AzureStorageCloudCredentialsImportSpec) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AzureStorageCloudCredentialsImportSpec) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AzureStorageCloudCredentialsImportSpec) SetTag(v string)`

SetTag sets Tag field to given value.


### GetAccount

`func (o *AzureStorageCloudCredentialsImportSpec) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AzureStorageCloudCredentialsImportSpec) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AzureStorageCloudCredentialsImportSpec) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetSharedKey

`func (o *AzureStorageCloudCredentialsImportSpec) GetSharedKey() string`

GetSharedKey returns the SharedKey field if non-nil, zero value otherwise.

### GetSharedKeyOk

`func (o *AzureStorageCloudCredentialsImportSpec) GetSharedKeyOk() (*string, bool)`

GetSharedKeyOk returns a tuple with the SharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedKey

`func (o *AzureStorageCloudCredentialsImportSpec) SetSharedKey(v string)`

SetSharedKey sets SharedKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


