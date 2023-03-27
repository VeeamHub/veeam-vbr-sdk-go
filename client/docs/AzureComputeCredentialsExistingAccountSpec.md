# AzureComputeCredentialsExistingAccountSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deployment** | [**AzureComputeCloudCredentialsDeploymentModel**](AzureComputeCloudCredentialsDeploymentModel.md) |  | 
**Subscription** | [**AzureComputeCloudCredentialsSubscriptionSpec**](AzureComputeCloudCredentialsSubscriptionSpec.md) |  | 

## Methods

### NewAzureComputeCredentialsExistingAccountSpec

`func NewAzureComputeCredentialsExistingAccountSpec(deployment AzureComputeCloudCredentialsDeploymentModel, subscription AzureComputeCloudCredentialsSubscriptionSpec, ) *AzureComputeCredentialsExistingAccountSpec`

NewAzureComputeCredentialsExistingAccountSpec instantiates a new AzureComputeCredentialsExistingAccountSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureComputeCredentialsExistingAccountSpecWithDefaults

`func NewAzureComputeCredentialsExistingAccountSpecWithDefaults() *AzureComputeCredentialsExistingAccountSpec`

NewAzureComputeCredentialsExistingAccountSpecWithDefaults instantiates a new AzureComputeCredentialsExistingAccountSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployment

`func (o *AzureComputeCredentialsExistingAccountSpec) GetDeployment() AzureComputeCloudCredentialsDeploymentModel`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *AzureComputeCredentialsExistingAccountSpec) GetDeploymentOk() (*AzureComputeCloudCredentialsDeploymentModel, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *AzureComputeCredentialsExistingAccountSpec) SetDeployment(v AzureComputeCloudCredentialsDeploymentModel)`

SetDeployment sets Deployment field to given value.


### GetSubscription

`func (o *AzureComputeCredentialsExistingAccountSpec) GetSubscription() AzureComputeCloudCredentialsSubscriptionSpec`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *AzureComputeCredentialsExistingAccountSpec) GetSubscriptionOk() (*AzureComputeCloudCredentialsSubscriptionSpec, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *AzureComputeCredentialsExistingAccountSpec) SetSubscription(v AzureComputeCloudCredentialsSubscriptionSpec)`

SetSubscription sets Subscription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


