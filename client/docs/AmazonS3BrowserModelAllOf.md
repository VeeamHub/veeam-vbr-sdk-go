# AmazonS3BrowserModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | **string** | ID of a server used to connect to the object storage. | 
**RegionType** | [**EAmazonRegionType**](EAmazonRegionType.md) |  | 
**Regions** | Pointer to [**[]AmazonS3RegionBrowserModel**](AmazonS3RegionBrowserModel.md) | Array of AWS regions belonged to the region type. | [optional] 

## Methods

### NewAmazonS3BrowserModelAllOf

`func NewAmazonS3BrowserModelAllOf(hostId string, regionType EAmazonRegionType, ) *AmazonS3BrowserModelAllOf`

NewAmazonS3BrowserModelAllOf instantiates a new AmazonS3BrowserModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonS3BrowserModelAllOfWithDefaults

`func NewAmazonS3BrowserModelAllOfWithDefaults() *AmazonS3BrowserModelAllOf`

NewAmazonS3BrowserModelAllOfWithDefaults instantiates a new AmazonS3BrowserModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *AmazonS3BrowserModelAllOf) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *AmazonS3BrowserModelAllOf) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *AmazonS3BrowserModelAllOf) SetHostId(v string)`

SetHostId sets HostId field to given value.


### GetRegionType

`func (o *AmazonS3BrowserModelAllOf) GetRegionType() EAmazonRegionType`

GetRegionType returns the RegionType field if non-nil, zero value otherwise.

### GetRegionTypeOk

`func (o *AmazonS3BrowserModelAllOf) GetRegionTypeOk() (*EAmazonRegionType, bool)`

GetRegionTypeOk returns a tuple with the RegionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionType

`func (o *AmazonS3BrowserModelAllOf) SetRegionType(v EAmazonRegionType)`

SetRegionType sets RegionType field to given value.


### GetRegions

`func (o *AmazonS3BrowserModelAllOf) GetRegions() []AmazonS3RegionBrowserModel`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *AmazonS3BrowserModelAllOf) GetRegionsOk() (*[]AmazonS3RegionBrowserModel, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *AmazonS3BrowserModelAllOf) SetRegions(v []AmazonS3RegionBrowserModel)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *AmazonS3BrowserModelAllOf) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


