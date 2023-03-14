# S3CompatibleRegionBrowserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Region ID. | [optional] 
**Name** | Pointer to **string** | Region name. | [optional] 
**Buckets** | Pointer to [**[]S3CompatibleBucketBrowserModel**](S3CompatibleBucketBrowserModel.md) | Array of buckets located in the region. | [optional] 

## Methods

### NewS3CompatibleRegionBrowserModel

`func NewS3CompatibleRegionBrowserModel() *S3CompatibleRegionBrowserModel`

NewS3CompatibleRegionBrowserModel instantiates a new S3CompatibleRegionBrowserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3CompatibleRegionBrowserModelWithDefaults

`func NewS3CompatibleRegionBrowserModelWithDefaults() *S3CompatibleRegionBrowserModel`

NewS3CompatibleRegionBrowserModelWithDefaults instantiates a new S3CompatibleRegionBrowserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *S3CompatibleRegionBrowserModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *S3CompatibleRegionBrowserModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *S3CompatibleRegionBrowserModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *S3CompatibleRegionBrowserModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *S3CompatibleRegionBrowserModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *S3CompatibleRegionBrowserModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *S3CompatibleRegionBrowserModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *S3CompatibleRegionBrowserModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBuckets

`func (o *S3CompatibleRegionBrowserModel) GetBuckets() []S3CompatibleBucketBrowserModel`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *S3CompatibleRegionBrowserModel) GetBucketsOk() (*[]S3CompatibleBucketBrowserModel, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *S3CompatibleRegionBrowserModel) SetBuckets(v []S3CompatibleBucketBrowserModel)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *S3CompatibleRegionBrowserModel) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


