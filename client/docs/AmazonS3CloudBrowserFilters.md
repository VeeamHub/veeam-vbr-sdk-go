# AmazonS3CloudBrowserFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionId** | **string** | Filters buckets by AWS region where an Amazon S3 data center is located. | 
**BucketName** | Pointer to **string** | Filters buckets by bucket name. | [optional] 

## Methods

### NewAmazonS3CloudBrowserFilters

`func NewAmazonS3CloudBrowserFilters(regionId string, ) *AmazonS3CloudBrowserFilters`

NewAmazonS3CloudBrowserFilters instantiates a new AmazonS3CloudBrowserFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonS3CloudBrowserFiltersWithDefaults

`func NewAmazonS3CloudBrowserFiltersWithDefaults() *AmazonS3CloudBrowserFilters`

NewAmazonS3CloudBrowserFiltersWithDefaults instantiates a new AmazonS3CloudBrowserFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionId

`func (o *AmazonS3CloudBrowserFilters) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *AmazonS3CloudBrowserFilters) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *AmazonS3CloudBrowserFilters) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetBucketName

`func (o *AmazonS3CloudBrowserFilters) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *AmazonS3CloudBrowserFilters) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *AmazonS3CloudBrowserFilters) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *AmazonS3CloudBrowserFilters) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


