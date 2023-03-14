# WasabiCloudStorageRegionBrowserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Region ID. | [optional] 
**Name** | Pointer to **string** | Region name. | [optional] 
**Buckets** | Pointer to [**[]WasabiCloudStorageBucketBrowserModel**](WasabiCloudStorageBucketBrowserModel.md) | Array of buckets located in the region. | [optional] 

## Methods

### NewWasabiCloudStorageRegionBrowserModel

`func NewWasabiCloudStorageRegionBrowserModel() *WasabiCloudStorageRegionBrowserModel`

NewWasabiCloudStorageRegionBrowserModel instantiates a new WasabiCloudStorageRegionBrowserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWasabiCloudStorageRegionBrowserModelWithDefaults

`func NewWasabiCloudStorageRegionBrowserModelWithDefaults() *WasabiCloudStorageRegionBrowserModel`

NewWasabiCloudStorageRegionBrowserModelWithDefaults instantiates a new WasabiCloudStorageRegionBrowserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WasabiCloudStorageRegionBrowserModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WasabiCloudStorageRegionBrowserModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WasabiCloudStorageRegionBrowserModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WasabiCloudStorageRegionBrowserModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WasabiCloudStorageRegionBrowserModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WasabiCloudStorageRegionBrowserModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WasabiCloudStorageRegionBrowserModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WasabiCloudStorageRegionBrowserModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBuckets

`func (o *WasabiCloudStorageRegionBrowserModel) GetBuckets() []WasabiCloudStorageBucketBrowserModel`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *WasabiCloudStorageRegionBrowserModel) GetBucketsOk() (*[]WasabiCloudStorageBucketBrowserModel, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *WasabiCloudStorageRegionBrowserModel) SetBuckets(v []WasabiCloudStorageBucketBrowserModel)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *WasabiCloudStorageRegionBrowserModel) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


