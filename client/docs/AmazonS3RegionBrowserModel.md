# AmazonS3RegionBrowserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Region ID. | [optional] 
**Name** | Pointer to **string** | Region name. | [optional] 
**Buckets** | Pointer to [**[]AmazonS3BucketBrowserModel**](AmazonS3BucketBrowserModel.md) | Array of buckets located in the region. | [optional] 

## Methods

### NewAmazonS3RegionBrowserModel

`func NewAmazonS3RegionBrowserModel() *AmazonS3RegionBrowserModel`

NewAmazonS3RegionBrowserModel instantiates a new AmazonS3RegionBrowserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonS3RegionBrowserModelWithDefaults

`func NewAmazonS3RegionBrowserModelWithDefaults() *AmazonS3RegionBrowserModel`

NewAmazonS3RegionBrowserModelWithDefaults instantiates a new AmazonS3RegionBrowserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AmazonS3RegionBrowserModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AmazonS3RegionBrowserModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AmazonS3RegionBrowserModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AmazonS3RegionBrowserModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AmazonS3RegionBrowserModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AmazonS3RegionBrowserModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AmazonS3RegionBrowserModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AmazonS3RegionBrowserModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBuckets

`func (o *AmazonS3RegionBrowserModel) GetBuckets() []AmazonS3BucketBrowserModel`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *AmazonS3RegionBrowserModel) GetBucketsOk() (*[]AmazonS3BucketBrowserModel, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *AmazonS3RegionBrowserModel) SetBuckets(v []AmazonS3BucketBrowserModel)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *AmazonS3RegionBrowserModel) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


