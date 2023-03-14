# AzureComputeCloudCredentialsSubscriptionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID that Veeam Backup &amp; Replication assigned to the Azure subscription. | 
**AzureSubscriptionId** | **string** | Original Azure subscription ID. | 
**AzureSubscriptionName** | Pointer to **string** | Azure subscription name. | [optional] 

## Methods

### NewAzureComputeCloudCredentialsSubscriptionInfo

`func NewAzureComputeCloudCredentialsSubscriptionInfo(id string, azureSubscriptionId string, ) *AzureComputeCloudCredentialsSubscriptionInfo`

NewAzureComputeCloudCredentialsSubscriptionInfo instantiates a new AzureComputeCloudCredentialsSubscriptionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureComputeCloudCredentialsSubscriptionInfoWithDefaults

`func NewAzureComputeCloudCredentialsSubscriptionInfoWithDefaults() *AzureComputeCloudCredentialsSubscriptionInfo`

NewAzureComputeCloudCredentialsSubscriptionInfoWithDefaults instantiates a new AzureComputeCloudCredentialsSubscriptionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AzureComputeCloudCredentialsSubscriptionInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureComputeCloudCredentialsSubscriptionInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureComputeCloudCredentialsSubscriptionInfo) SetId(v string)`

SetId sets Id field to given value.


### GetAzureSubscriptionId

`func (o *AzureComputeCloudCredentialsSubscriptionInfo) GetAzureSubscriptionId() string`

GetAzureSubscriptionId returns the AzureSubscriptionId field if non-nil, zero value otherwise.

### GetAzureSubscriptionIdOk

`func (o *AzureComputeCloudCredentialsSubscriptionInfo) GetAzureSubscriptionIdOk() (*string, bool)`

GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionId

`func (o *AzureComputeCloudCredentialsSubscriptionInfo) SetAzureSubscriptionId(v string)`

SetAzureSubscriptionId sets AzureSubscriptionId field to given value.


### GetAzureSubscriptionName

`func (o *AzureComputeCloudCredentialsSubscriptionInfo) GetAzureSubscriptionName() string`

GetAzureSubscriptionName returns the AzureSubscriptionName field if non-nil, zero value otherwise.

### GetAzureSubscriptionNameOk

`func (o *AzureComputeCloudCredentialsSubscriptionInfo) GetAzureSubscriptionNameOk() (*string, bool)`

GetAzureSubscriptionNameOk returns a tuple with the AzureSubscriptionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionName

`func (o *AzureComputeCloudCredentialsSubscriptionInfo) SetAzureSubscriptionName(v string)`

SetAzureSubscriptionName sets AzureSubscriptionName field to given value.

### HasAzureSubscriptionName

`func (o *AzureComputeCloudCredentialsSubscriptionInfo) HasAzureSubscriptionName() bool`

HasAzureSubscriptionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


