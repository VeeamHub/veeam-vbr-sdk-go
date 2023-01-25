# CloudCredentialsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the cloud credentials record. | 
**Description** | Pointer to **string** | Description of the cloud credentials record. | [optional] 
**Type** | [**ECloudCredentialsType**](ECloudCredentialsType.md) |  | 
**AccessKey** | **string** | Access ID of a Google HMAC key. | 
**Tag** | Pointer to **string** | Tag used to identify the cloud credentials record. | [optional] 
**ConnectionName** | **string** | Name under which the cloud credentials record is shown in Veeam Backup &amp; Replication. | 
**Deployment** | [**AzureComputeCloudCredentialsDeploymentModel**](AzureComputeCloudCredentialsDeploymentModel.md) |  | 
**Subscription** | [**AzureComputeCloudCredentialsSubscriptionModel**](AzureComputeCloudCredentialsSubscriptionModel.md) |  | 
**Account** | **string** | Name of the Azure storage account. | 

## Methods

### NewCloudCredentialsModel

`func NewCloudCredentialsModel(id string, type_ ECloudCredentialsType, accessKey string, connectionName string, deployment AzureComputeCloudCredentialsDeploymentModel, subscription AzureComputeCloudCredentialsSubscriptionModel, account string, ) *CloudCredentialsModel`

NewCloudCredentialsModel instantiates a new CloudCredentialsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCredentialsModelWithDefaults

`func NewCloudCredentialsModelWithDefaults() *CloudCredentialsModel`

NewCloudCredentialsModelWithDefaults instantiates a new CloudCredentialsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudCredentialsModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudCredentialsModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudCredentialsModel) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *CloudCredentialsModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudCredentialsModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudCredentialsModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudCredentialsModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *CloudCredentialsModel) GetType() ECloudCredentialsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudCredentialsModel) GetTypeOk() (*ECloudCredentialsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudCredentialsModel) SetType(v ECloudCredentialsType)`

SetType sets Type field to given value.


### GetAccessKey

`func (o *CloudCredentialsModel) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *CloudCredentialsModel) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *CloudCredentialsModel) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetTag

`func (o *CloudCredentialsModel) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CloudCredentialsModel) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CloudCredentialsModel) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CloudCredentialsModel) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetConnectionName

`func (o *CloudCredentialsModel) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *CloudCredentialsModel) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *CloudCredentialsModel) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.


### GetDeployment

`func (o *CloudCredentialsModel) GetDeployment() AzureComputeCloudCredentialsDeploymentModel`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *CloudCredentialsModel) GetDeploymentOk() (*AzureComputeCloudCredentialsDeploymentModel, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *CloudCredentialsModel) SetDeployment(v AzureComputeCloudCredentialsDeploymentModel)`

SetDeployment sets Deployment field to given value.


### GetSubscription

`func (o *CloudCredentialsModel) GetSubscription() AzureComputeCloudCredentialsSubscriptionModel`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *CloudCredentialsModel) GetSubscriptionOk() (*AzureComputeCloudCredentialsSubscriptionModel, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *CloudCredentialsModel) SetSubscription(v AzureComputeCloudCredentialsSubscriptionModel)`

SetSubscription sets Subscription field to given value.


### GetAccount

`func (o *CloudCredentialsModel) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CloudCredentialsModel) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CloudCredentialsModel) SetAccount(v string)`

SetAccount sets Account field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


