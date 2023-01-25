# IBMCloudStorageRegionBrowserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Region ID. | [optional] 
**Name** | Pointer to **string** | Region name. | [optional] 
**Buckets** | Pointer to [**[]IBMCloudStorageBucketBrowserModel**](IBMCloudStorageBucketBrowserModel.md) | Array of buckets located in the region. | [optional] 

## Methods

### NewIBMCloudStorageRegionBrowserModel

`func NewIBMCloudStorageRegionBrowserModel() *IBMCloudStorageRegionBrowserModel`

NewIBMCloudStorageRegionBrowserModel instantiates a new IBMCloudStorageRegionBrowserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIBMCloudStorageRegionBrowserModelWithDefaults

`func NewIBMCloudStorageRegionBrowserModelWithDefaults() *IBMCloudStorageRegionBrowserModel`

NewIBMCloudStorageRegionBrowserModelWithDefaults instantiates a new IBMCloudStorageRegionBrowserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IBMCloudStorageRegionBrowserModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IBMCloudStorageRegionBrowserModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IBMCloudStorageRegionBrowserModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IBMCloudStorageRegionBrowserModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IBMCloudStorageRegionBrowserModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IBMCloudStorageRegionBrowserModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IBMCloudStorageRegionBrowserModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IBMCloudStorageRegionBrowserModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBuckets

`func (o *IBMCloudStorageRegionBrowserModel) GetBuckets() []IBMCloudStorageBucketBrowserModel`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *IBMCloudStorageRegionBrowserModel) GetBucketsOk() (*[]IBMCloudStorageBucketBrowserModel, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *IBMCloudStorageRegionBrowserModel) SetBuckets(v []IBMCloudStorageBucketBrowserModel)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *IBMCloudStorageRegionBrowserModel) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


