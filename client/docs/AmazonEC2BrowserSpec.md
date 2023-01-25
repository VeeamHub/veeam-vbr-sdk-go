# AmazonEC2BrowserSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionType** | [**EAmazonRegionType**](EAmazonRegionType.md) |  | 
**Filters** | Pointer to [**AmazonEC2BrowserFilters**](AmazonEC2BrowserFilters.md) |  | [optional] 

## Methods

### NewAmazonEC2BrowserSpec

`func NewAmazonEC2BrowserSpec(regionType EAmazonRegionType, ) *AmazonEC2BrowserSpec`

NewAmazonEC2BrowserSpec instantiates a new AmazonEC2BrowserSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonEC2BrowserSpecWithDefaults

`func NewAmazonEC2BrowserSpecWithDefaults() *AmazonEC2BrowserSpec`

NewAmazonEC2BrowserSpecWithDefaults instantiates a new AmazonEC2BrowserSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionType

`func (o *AmazonEC2BrowserSpec) GetRegionType() EAmazonRegionType`

GetRegionType returns the RegionType field if non-nil, zero value otherwise.

### GetRegionTypeOk

`func (o *AmazonEC2BrowserSpec) GetRegionTypeOk() (*EAmazonRegionType, bool)`

GetRegionTypeOk returns a tuple with the RegionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionType

`func (o *AmazonEC2BrowserSpec) SetRegionType(v EAmazonRegionType)`

SetRegionType sets RegionType field to given value.


### GetFilters

`func (o *AmazonEC2BrowserSpec) GetFilters() AmazonEC2BrowserFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *AmazonEC2BrowserSpec) GetFiltersOk() (*AmazonEC2BrowserFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *AmazonEC2BrowserSpec) SetFilters(v AmazonEC2BrowserFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *AmazonEC2BrowserSpec) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


