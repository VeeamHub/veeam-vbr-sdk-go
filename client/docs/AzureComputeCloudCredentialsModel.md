# AzureComputeCloudCredentialsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionName** | **string** | Name under which the cloud credentials record is shown in Veeam Backup &amp; Replication. | 
**Deployment** | [**AzureComputeCloudCredentialsDeploymentModel**](AzureComputeCloudCredentialsDeploymentModel.md) |  | 
**Subscription** | [**AzureComputeCloudCredentialsSubscriptionModel**](AzureComputeCloudCredentialsSubscriptionModel.md) |  | 
**Tag** | Pointer to **string** | Tag used to identify the cloud credentials record. | [optional] 

## Methods

### NewAzureComputeCloudCredentialsModel

`func NewAzureComputeCloudCredentialsModel(connectionName string, deployment AzureComputeCloudCredentialsDeploymentModel, subscription AzureComputeCloudCredentialsSubscriptionModel, ) *AzureComputeCloudCredentialsModel`

NewAzureComputeCloudCredentialsModel instantiates a new AzureComputeCloudCredentialsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureComputeCloudCredentialsModelWithDefaults

`func NewAzureComputeCloudCredentialsModelWithDefaults() *AzureComputeCloudCredentialsModel`

NewAzureComputeCloudCredentialsModelWithDefaults instantiates a new AzureComputeCloudCredentialsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionName

`func (o *AzureComputeCloudCredentialsModel) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *AzureComputeCloudCredentialsModel) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *AzureComputeCloudCredentialsModel) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.


### GetDeployment

`func (o *AzureComputeCloudCredentialsModel) GetDeployment() AzureComputeCloudCredentialsDeploymentModel`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *AzureComputeCloudCredentialsModel) GetDeploymentOk() (*AzureComputeCloudCredentialsDeploymentModel, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *AzureComputeCloudCredentialsModel) SetDeployment(v AzureComputeCloudCredentialsDeploymentModel)`

SetDeployment sets Deployment field to given value.


### GetSubscription

`func (o *AzureComputeCloudCredentialsModel) GetSubscription() AzureComputeCloudCredentialsSubscriptionModel`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *AzureComputeCloudCredentialsModel) GetSubscriptionOk() (*AzureComputeCloudCredentialsSubscriptionModel, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *AzureComputeCloudCredentialsModel) SetSubscription(v AzureComputeCloudCredentialsSubscriptionModel)`

SetSubscription sets Subscription field to given value.


### GetTag

`func (o *AzureComputeCloudCredentialsModel) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AzureComputeCloudCredentialsModel) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AzureComputeCloudCredentialsModel) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *AzureComputeCloudCredentialsModel) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


