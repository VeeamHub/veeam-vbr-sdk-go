# AzureComputeCloudCredentialsModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionName** | **string** | Name under which the cloud credentials record is shown in Veeam Backup &amp; Replication. | 
**Deployment** | [**AzureComputeCloudCredentialsDeploymentModel**](AzureComputeCloudCredentialsDeploymentModel.md) |  | 
**Subscription** | [**AzureComputeCloudCredentialsSubscriptionModel**](AzureComputeCloudCredentialsSubscriptionModel.md) |  | 
**Tag** | Pointer to **string** | Tag used to identify the cloud credentials record. | [optional] 

## Methods

### NewAzureComputeCloudCredentialsModelAllOf

`func NewAzureComputeCloudCredentialsModelAllOf(connectionName string, deployment AzureComputeCloudCredentialsDeploymentModel, subscription AzureComputeCloudCredentialsSubscriptionModel, ) *AzureComputeCloudCredentialsModelAllOf`

NewAzureComputeCloudCredentialsModelAllOf instantiates a new AzureComputeCloudCredentialsModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureComputeCloudCredentialsModelAllOfWithDefaults

`func NewAzureComputeCloudCredentialsModelAllOfWithDefaults() *AzureComputeCloudCredentialsModelAllOf`

NewAzureComputeCloudCredentialsModelAllOfWithDefaults instantiates a new AzureComputeCloudCredentialsModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionName

`func (o *AzureComputeCloudCredentialsModelAllOf) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *AzureComputeCloudCredentialsModelAllOf) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *AzureComputeCloudCredentialsModelAllOf) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.


### GetDeployment

`func (o *AzureComputeCloudCredentialsModelAllOf) GetDeployment() AzureComputeCloudCredentialsDeploymentModel`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *AzureComputeCloudCredentialsModelAllOf) GetDeploymentOk() (*AzureComputeCloudCredentialsDeploymentModel, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *AzureComputeCloudCredentialsModelAllOf) SetDeployment(v AzureComputeCloudCredentialsDeploymentModel)`

SetDeployment sets Deployment field to given value.


### GetSubscription

`func (o *AzureComputeCloudCredentialsModelAllOf) GetSubscription() AzureComputeCloudCredentialsSubscriptionModel`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *AzureComputeCloudCredentialsModelAllOf) GetSubscriptionOk() (*AzureComputeCloudCredentialsSubscriptionModel, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *AzureComputeCloudCredentialsModelAllOf) SetSubscription(v AzureComputeCloudCredentialsSubscriptionModel)`

SetSubscription sets Subscription field to given value.


### GetTag

`func (o *AzureComputeCloudCredentialsModelAllOf) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AzureComputeCloudCredentialsModelAllOf) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AzureComputeCloudCredentialsModelAllOf) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *AzureComputeCloudCredentialsModelAllOf) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


