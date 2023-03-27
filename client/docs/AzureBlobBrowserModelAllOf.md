# AzureBlobBrowserModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | Pointer to **string** | ID of a server used to connect to the object storage. | [optional] 
**RegionType** | Pointer to [**EAzureRegionType**](EAzureRegionType.md) |  | [optional] 
**Containers** | Pointer to [**[]AzureBlobContainerBrowserModel**](AzureBlobContainerBrowserModel.md) | Array of containers that reside in the Azure storage account. | [optional] 

## Methods

### NewAzureBlobBrowserModelAllOf

`func NewAzureBlobBrowserModelAllOf() *AzureBlobBrowserModelAllOf`

NewAzureBlobBrowserModelAllOf instantiates a new AzureBlobBrowserModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureBlobBrowserModelAllOfWithDefaults

`func NewAzureBlobBrowserModelAllOfWithDefaults() *AzureBlobBrowserModelAllOf`

NewAzureBlobBrowserModelAllOfWithDefaults instantiates a new AzureBlobBrowserModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *AzureBlobBrowserModelAllOf) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *AzureBlobBrowserModelAllOf) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *AzureBlobBrowserModelAllOf) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *AzureBlobBrowserModelAllOf) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetRegionType

`func (o *AzureBlobBrowserModelAllOf) GetRegionType() EAzureRegionType`

GetRegionType returns the RegionType field if non-nil, zero value otherwise.

### GetRegionTypeOk

`func (o *AzureBlobBrowserModelAllOf) GetRegionTypeOk() (*EAzureRegionType, bool)`

GetRegionTypeOk returns a tuple with the RegionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionType

`func (o *AzureBlobBrowserModelAllOf) SetRegionType(v EAzureRegionType)`

SetRegionType sets RegionType field to given value.

### HasRegionType

`func (o *AzureBlobBrowserModelAllOf) HasRegionType() bool`

HasRegionType returns a boolean if a field has been set.

### GetContainers

`func (o *AzureBlobBrowserModelAllOf) GetContainers() []AzureBlobContainerBrowserModel`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *AzureBlobBrowserModelAllOf) GetContainersOk() (*[]AzureBlobContainerBrowserModel, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *AzureBlobBrowserModelAllOf) SetContainers(v []AzureBlobContainerBrowserModel)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *AzureBlobBrowserModelAllOf) HasContainers() bool`

HasContainers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


