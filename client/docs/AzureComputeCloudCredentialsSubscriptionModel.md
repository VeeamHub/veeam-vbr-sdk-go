# AzureComputeCloudCredentialsSubscriptionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | ID of a tenant where the Azure Active Directory application is registered in. | 
**ApplicationId** | **string** | Client ID assigned to the Azure Active Directory application. | 
**Secret** | Pointer to **string** | (For password-based authentication) Client secret assigned to the Azure Active Directory application. | [optional] 
**Subscriptions** | Pointer to [**[]AzureComputeCloudCredentialsSubscriptionInfo**](AzureComputeCloudCredentialsSubscriptionInfo.md) | Array of Azure subscriptions associated with the account. | [optional] 

## Methods

### NewAzureComputeCloudCredentialsSubscriptionModel

`func NewAzureComputeCloudCredentialsSubscriptionModel(tenantId string, applicationId string, ) *AzureComputeCloudCredentialsSubscriptionModel`

NewAzureComputeCloudCredentialsSubscriptionModel instantiates a new AzureComputeCloudCredentialsSubscriptionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureComputeCloudCredentialsSubscriptionModelWithDefaults

`func NewAzureComputeCloudCredentialsSubscriptionModelWithDefaults() *AzureComputeCloudCredentialsSubscriptionModel`

NewAzureComputeCloudCredentialsSubscriptionModelWithDefaults instantiates a new AzureComputeCloudCredentialsSubscriptionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *AzureComputeCloudCredentialsSubscriptionModel) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureComputeCloudCredentialsSubscriptionModel) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureComputeCloudCredentialsSubscriptionModel) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetApplicationId

`func (o *AzureComputeCloudCredentialsSubscriptionModel) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *AzureComputeCloudCredentialsSubscriptionModel) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *AzureComputeCloudCredentialsSubscriptionModel) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetSecret

`func (o *AzureComputeCloudCredentialsSubscriptionModel) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *AzureComputeCloudCredentialsSubscriptionModel) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *AzureComputeCloudCredentialsSubscriptionModel) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *AzureComputeCloudCredentialsSubscriptionModel) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetSubscriptions

`func (o *AzureComputeCloudCredentialsSubscriptionModel) GetSubscriptions() []AzureComputeCloudCredentialsSubscriptionInfo`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *AzureComputeCloudCredentialsSubscriptionModel) GetSubscriptionsOk() (*[]AzureComputeCloudCredentialsSubscriptionInfo, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *AzureComputeCloudCredentialsSubscriptionModel) SetSubscriptions(v []AzureComputeCloudCredentialsSubscriptionInfo)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *AzureComputeCloudCredentialsSubscriptionModel) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


