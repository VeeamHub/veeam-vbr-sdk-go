# AmazonEC2BrowserModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | **string** | ID of a server used to connect to the object storage. | 
**RegionType** | [**EAmazonRegionType**](EAmazonRegionType.md) |  | 
**Regions** | [**[]AmazonEC2RegionBrowserModel**](AmazonEC2RegionBrowserModel.md) | Array of regions. | 

## Methods

### NewAmazonEC2BrowserModelAllOf

`func NewAmazonEC2BrowserModelAllOf(hostId string, regionType EAmazonRegionType, regions []AmazonEC2RegionBrowserModel, ) *AmazonEC2BrowserModelAllOf`

NewAmazonEC2BrowserModelAllOf instantiates a new AmazonEC2BrowserModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonEC2BrowserModelAllOfWithDefaults

`func NewAmazonEC2BrowserModelAllOfWithDefaults() *AmazonEC2BrowserModelAllOf`

NewAmazonEC2BrowserModelAllOfWithDefaults instantiates a new AmazonEC2BrowserModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *AmazonEC2BrowserModelAllOf) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *AmazonEC2BrowserModelAllOf) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *AmazonEC2BrowserModelAllOf) SetHostId(v string)`

SetHostId sets HostId field to given value.


### GetRegionType

`func (o *AmazonEC2BrowserModelAllOf) GetRegionType() EAmazonRegionType`

GetRegionType returns the RegionType field if non-nil, zero value otherwise.

### GetRegionTypeOk

`func (o *AmazonEC2BrowserModelAllOf) GetRegionTypeOk() (*EAmazonRegionType, bool)`

GetRegionTypeOk returns a tuple with the RegionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionType

`func (o *AmazonEC2BrowserModelAllOf) SetRegionType(v EAmazonRegionType)`

SetRegionType sets RegionType field to given value.


### GetRegions

`func (o *AmazonEC2BrowserModelAllOf) GetRegions() []AmazonEC2RegionBrowserModel`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *AmazonEC2BrowserModelAllOf) GetRegionsOk() (*[]AmazonEC2RegionBrowserModel, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *AmazonEC2BrowserModelAllOf) SetRegions(v []AmazonEC2RegionBrowserModel)`

SetRegions sets Regions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


