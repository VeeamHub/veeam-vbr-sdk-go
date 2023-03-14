# AzureBlobBrowserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | Pointer to **string** | ID of a server used to connect to the object storage. | [optional] 
**RegionType** | Pointer to [**EAzureRegionType**](EAzureRegionType.md) |  | [optional] 
**Containers** | Pointer to [**[]AzureBlobContainerBrowserModel**](AzureBlobContainerBrowserModel.md) | Array of containers that reside in the Azure storage account. | [optional] 

## Methods

### NewAzureBlobBrowserModel

`func NewAzureBlobBrowserModel() *AzureBlobBrowserModel`

NewAzureBlobBrowserModel instantiates a new AzureBlobBrowserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureBlobBrowserModelWithDefaults

`func NewAzureBlobBrowserModelWithDefaults() *AzureBlobBrowserModel`

NewAzureBlobBrowserModelWithDefaults instantiates a new AzureBlobBrowserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *AzureBlobBrowserModel) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *AzureBlobBrowserModel) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *AzureBlobBrowserModel) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *AzureBlobBrowserModel) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetRegionType

`func (o *AzureBlobBrowserModel) GetRegionType() EAzureRegionType`

GetRegionType returns the RegionType field if non-nil, zero value otherwise.

### GetRegionTypeOk

`func (o *AzureBlobBrowserModel) GetRegionTypeOk() (*EAzureRegionType, bool)`

GetRegionTypeOk returns a tuple with the RegionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionType

`func (o *AzureBlobBrowserModel) SetRegionType(v EAzureRegionType)`

SetRegionType sets RegionType field to given value.

### HasRegionType

`func (o *AzureBlobBrowserModel) HasRegionType() bool`

HasRegionType returns a boolean if a field has been set.

### GetContainers

`func (o *AzureBlobBrowserModel) GetContainers() []AzureBlobContainerBrowserModel`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *AzureBlobBrowserModel) GetContainersOk() (*[]AzureBlobContainerBrowserModel, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *AzureBlobBrowserModel) SetContainers(v []AzureBlobContainerBrowserModel)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *AzureBlobBrowserModel) HasContainers() bool`

HasContainers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


