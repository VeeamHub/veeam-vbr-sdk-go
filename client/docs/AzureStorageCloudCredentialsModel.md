# AzureStorageCloudCredentialsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | Name of the Azure storage account. | 
**Tag** | Pointer to **string** | Tag used to identify the cloud credentials record. | [optional] 

## Methods

### NewAzureStorageCloudCredentialsModel

`func NewAzureStorageCloudCredentialsModel(account string, ) *AzureStorageCloudCredentialsModel`

NewAzureStorageCloudCredentialsModel instantiates a new AzureStorageCloudCredentialsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureStorageCloudCredentialsModelWithDefaults

`func NewAzureStorageCloudCredentialsModelWithDefaults() *AzureStorageCloudCredentialsModel`

NewAzureStorageCloudCredentialsModelWithDefaults instantiates a new AzureStorageCloudCredentialsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AzureStorageCloudCredentialsModel) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AzureStorageCloudCredentialsModel) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AzureStorageCloudCredentialsModel) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetTag

`func (o *AzureStorageCloudCredentialsModel) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AzureStorageCloudCredentialsModel) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AzureStorageCloudCredentialsModel) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *AzureStorageCloudCredentialsModel) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


