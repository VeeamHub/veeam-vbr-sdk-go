# AzureComputeCloudNativeApplicationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **string** | Client ID assigned to the Azure Active Directory application. | 
**Secret** | **string** | Client secret assigned to the Azure Active Directory application. | 
**TenantId** | **string** | ID of a tenant where the Azure Active Directory application is registered in. | 

## Methods

### NewAzureComputeCloudNativeApplicationModel

`func NewAzureComputeCloudNativeApplicationModel(applicationId string, secret string, tenantId string, ) *AzureComputeCloudNativeApplicationModel`

NewAzureComputeCloudNativeApplicationModel instantiates a new AzureComputeCloudNativeApplicationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureComputeCloudNativeApplicationModelWithDefaults

`func NewAzureComputeCloudNativeApplicationModelWithDefaults() *AzureComputeCloudNativeApplicationModel`

NewAzureComputeCloudNativeApplicationModelWithDefaults instantiates a new AzureComputeCloudNativeApplicationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *AzureComputeCloudNativeApplicationModel) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *AzureComputeCloudNativeApplicationModel) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *AzureComputeCloudNativeApplicationModel) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetSecret

`func (o *AzureComputeCloudNativeApplicationModel) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *AzureComputeCloudNativeApplicationModel) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *AzureComputeCloudNativeApplicationModel) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetTenantId

`func (o *AzureComputeCloudNativeApplicationModel) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureComputeCloudNativeApplicationModel) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureComputeCloudNativeApplicationModel) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


