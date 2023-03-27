# AmazonSnowballEdgeRegionBrowserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Region ID. | [optional] 
**Name** | Pointer to **string** | Region name. | [optional] 
**Buckets** | Pointer to [**[]AmazonSnowballEdgeBucketBrowserModel**](AmazonSnowballEdgeBucketBrowserModel.md) | Array of buckets located in the region. | [optional] 

## Methods

### NewAmazonSnowballEdgeRegionBrowserModel

`func NewAmazonSnowballEdgeRegionBrowserModel() *AmazonSnowballEdgeRegionBrowserModel`

NewAmazonSnowballEdgeRegionBrowserModel instantiates a new AmazonSnowballEdgeRegionBrowserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonSnowballEdgeRegionBrowserModelWithDefaults

`func NewAmazonSnowballEdgeRegionBrowserModelWithDefaults() *AmazonSnowballEdgeRegionBrowserModel`

NewAmazonSnowballEdgeRegionBrowserModelWithDefaults instantiates a new AmazonSnowballEdgeRegionBrowserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AmazonSnowballEdgeRegionBrowserModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AmazonSnowballEdgeRegionBrowserModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AmazonSnowballEdgeRegionBrowserModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AmazonSnowballEdgeRegionBrowserModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AmazonSnowballEdgeRegionBrowserModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AmazonSnowballEdgeRegionBrowserModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AmazonSnowballEdgeRegionBrowserModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AmazonSnowballEdgeRegionBrowserModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBuckets

`func (o *AmazonSnowballEdgeRegionBrowserModel) GetBuckets() []AmazonSnowballEdgeBucketBrowserModel`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *AmazonSnowballEdgeRegionBrowserModel) GetBucketsOk() (*[]AmazonSnowballEdgeBucketBrowserModel, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *AmazonSnowballEdgeRegionBrowserModel) SetBuckets(v []AmazonSnowballEdgeBucketBrowserModel)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *AmazonSnowballEdgeRegionBrowserModel) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


