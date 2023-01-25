# CloudBrowserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialsId** | **string** | ID of the cloud credentials record. | 
**ServiceType** | [**ECloudServiceType**](ECloudServiceType.md) |  | 
**HostId** | **string** | ID of a server used to connect to the object storage. | 
**RegionType** | [**EAmazonRegionType**](EAmazonRegionType.md) |  | 
**Containers** | Pointer to [**[]AzureDataBoxContainerBrowserModel**](AzureDataBoxContainerBrowserModel.md) | Array of containers that reside in the Azure storage account. | [optional] 
**Regions** | [**[]WasabiCloudStorageRegionBrowserModel**](WasabiCloudStorageRegionBrowserModel.md) | Array of regions. | 
**ConnectionPoint** | Pointer to **string** | Service point address and port number of the S3 compatible storage. | [optional] 
**Subscriptions** | Pointer to [**[]AzureSubscriptionBrowserModel**](AzureSubscriptionBrowserModel.md) | Array of Azure subscriptions associated with the account. | [optional] 

## Methods

### NewCloudBrowserModel

`func NewCloudBrowserModel(credentialsId string, serviceType ECloudServiceType, hostId string, regionType EAmazonRegionType, regions []WasabiCloudStorageRegionBrowserModel, ) *CloudBrowserModel`

NewCloudBrowserModel instantiates a new CloudBrowserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBrowserModelWithDefaults

`func NewCloudBrowserModelWithDefaults() *CloudBrowserModel`

NewCloudBrowserModelWithDefaults instantiates a new CloudBrowserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialsId

`func (o *CloudBrowserModel) GetCredentialsId() string`

GetCredentialsId returns the CredentialsId field if non-nil, zero value otherwise.

### GetCredentialsIdOk

`func (o *CloudBrowserModel) GetCredentialsIdOk() (*string, bool)`

GetCredentialsIdOk returns a tuple with the CredentialsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsId

`func (o *CloudBrowserModel) SetCredentialsId(v string)`

SetCredentialsId sets CredentialsId field to given value.


### GetServiceType

`func (o *CloudBrowserModel) GetServiceType() ECloudServiceType`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *CloudBrowserModel) GetServiceTypeOk() (*ECloudServiceType, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *CloudBrowserModel) SetServiceType(v ECloudServiceType)`

SetServiceType sets ServiceType field to given value.


### GetHostId

`func (o *CloudBrowserModel) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *CloudBrowserModel) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *CloudBrowserModel) SetHostId(v string)`

SetHostId sets HostId field to given value.


### GetRegionType

`func (o *CloudBrowserModel) GetRegionType() EAmazonRegionType`

GetRegionType returns the RegionType field if non-nil, zero value otherwise.

### GetRegionTypeOk

`func (o *CloudBrowserModel) GetRegionTypeOk() (*EAmazonRegionType, bool)`

GetRegionTypeOk returns a tuple with the RegionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionType

`func (o *CloudBrowserModel) SetRegionType(v EAmazonRegionType)`

SetRegionType sets RegionType field to given value.


### GetContainers

`func (o *CloudBrowserModel) GetContainers() []AzureDataBoxContainerBrowserModel`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *CloudBrowserModel) GetContainersOk() (*[]AzureDataBoxContainerBrowserModel, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *CloudBrowserModel) SetContainers(v []AzureDataBoxContainerBrowserModel)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *CloudBrowserModel) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetRegions

`func (o *CloudBrowserModel) GetRegions() []WasabiCloudStorageRegionBrowserModel`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *CloudBrowserModel) GetRegionsOk() (*[]WasabiCloudStorageRegionBrowserModel, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *CloudBrowserModel) SetRegions(v []WasabiCloudStorageRegionBrowserModel)`

SetRegions sets Regions field to given value.


### GetConnectionPoint

`func (o *CloudBrowserModel) GetConnectionPoint() string`

GetConnectionPoint returns the ConnectionPoint field if non-nil, zero value otherwise.

### GetConnectionPointOk

`func (o *CloudBrowserModel) GetConnectionPointOk() (*string, bool)`

GetConnectionPointOk returns a tuple with the ConnectionPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPoint

`func (o *CloudBrowserModel) SetConnectionPoint(v string)`

SetConnectionPoint sets ConnectionPoint field to given value.

### HasConnectionPoint

`func (o *CloudBrowserModel) HasConnectionPoint() bool`

HasConnectionPoint returns a boolean if a field has been set.

### GetSubscriptions

`func (o *CloudBrowserModel) GetSubscriptions() []AzureSubscriptionBrowserModel`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *CloudBrowserModel) GetSubscriptionsOk() (*[]AzureSubscriptionBrowserModel, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *CloudBrowserModel) SetSubscriptions(v []AzureSubscriptionBrowserModel)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *CloudBrowserModel) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


