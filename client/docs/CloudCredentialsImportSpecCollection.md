# CloudCredentialsImportSpecCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureStorage** | Pointer to [**[]AzureStorageCloudCredentialsImportSpec**](AzureStorageCloudCredentialsImportSpec.md) | Array of Azure storage accounts. | [optional] 
**AzureCompute** | Pointer to [**[]AzureComputeCloudCredentialsImportSpec**](AzureComputeCloudCredentialsImportSpec.md) | Array of Azure compute accounts. | [optional] 
**Amazon** | Pointer to [**[]AmazonCloudCredentialsImportSpec**](AmazonCloudCredentialsImportSpec.md) | Array of AWS accounts. | [optional] 
**Google** | Pointer to [**[]GoogleCloudCredentialsImportSpec**](GoogleCloudCredentialsImportSpec.md) | Array of Google accounts. | [optional] 

## Methods

### NewCloudCredentialsImportSpecCollection

`func NewCloudCredentialsImportSpecCollection() *CloudCredentialsImportSpecCollection`

NewCloudCredentialsImportSpecCollection instantiates a new CloudCredentialsImportSpecCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCredentialsImportSpecCollectionWithDefaults

`func NewCloudCredentialsImportSpecCollectionWithDefaults() *CloudCredentialsImportSpecCollection`

NewCloudCredentialsImportSpecCollectionWithDefaults instantiates a new CloudCredentialsImportSpecCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureStorage

`func (o *CloudCredentialsImportSpecCollection) GetAzureStorage() []AzureStorageCloudCredentialsImportSpec`

GetAzureStorage returns the AzureStorage field if non-nil, zero value otherwise.

### GetAzureStorageOk

`func (o *CloudCredentialsImportSpecCollection) GetAzureStorageOk() (*[]AzureStorageCloudCredentialsImportSpec, bool)`

GetAzureStorageOk returns a tuple with the AzureStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureStorage

`func (o *CloudCredentialsImportSpecCollection) SetAzureStorage(v []AzureStorageCloudCredentialsImportSpec)`

SetAzureStorage sets AzureStorage field to given value.

### HasAzureStorage

`func (o *CloudCredentialsImportSpecCollection) HasAzureStorage() bool`

HasAzureStorage returns a boolean if a field has been set.

### GetAzureCompute

`func (o *CloudCredentialsImportSpecCollection) GetAzureCompute() []AzureComputeCloudCredentialsImportSpec`

GetAzureCompute returns the AzureCompute field if non-nil, zero value otherwise.

### GetAzureComputeOk

`func (o *CloudCredentialsImportSpecCollection) GetAzureComputeOk() (*[]AzureComputeCloudCredentialsImportSpec, bool)`

GetAzureComputeOk returns a tuple with the AzureCompute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureCompute

`func (o *CloudCredentialsImportSpecCollection) SetAzureCompute(v []AzureComputeCloudCredentialsImportSpec)`

SetAzureCompute sets AzureCompute field to given value.

### HasAzureCompute

`func (o *CloudCredentialsImportSpecCollection) HasAzureCompute() bool`

HasAzureCompute returns a boolean if a field has been set.

### GetAmazon

`func (o *CloudCredentialsImportSpecCollection) GetAmazon() []AmazonCloudCredentialsImportSpec`

GetAmazon returns the Amazon field if non-nil, zero value otherwise.

### GetAmazonOk

`func (o *CloudCredentialsImportSpecCollection) GetAmazonOk() (*[]AmazonCloudCredentialsImportSpec, bool)`

GetAmazonOk returns a tuple with the Amazon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazon

`func (o *CloudCredentialsImportSpecCollection) SetAmazon(v []AmazonCloudCredentialsImportSpec)`

SetAmazon sets Amazon field to given value.

### HasAmazon

`func (o *CloudCredentialsImportSpecCollection) HasAmazon() bool`

HasAmazon returns a boolean if a field has been set.

### GetGoogle

`func (o *CloudCredentialsImportSpecCollection) GetGoogle() []GoogleCloudCredentialsImportSpec`

GetGoogle returns the Google field if non-nil, zero value otherwise.

### GetGoogleOk

`func (o *CloudCredentialsImportSpecCollection) GetGoogleOk() (*[]GoogleCloudCredentialsImportSpec, bool)`

GetGoogleOk returns a tuple with the Google field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogle

`func (o *CloudCredentialsImportSpecCollection) SetGoogle(v []GoogleCloudCredentialsImportSpec)`

SetGoogle sets Google field to given value.

### HasGoogle

`func (o *CloudCredentialsImportSpecCollection) HasGoogle() bool`

HasGoogle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


