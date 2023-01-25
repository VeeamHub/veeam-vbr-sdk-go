# S3CompatibleBucketBrowserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Bucket name. | [optional] 
**Folders** | Pointer to **[]string** | Array of folders located in the bucket. | [optional] 

## Methods

### NewS3CompatibleBucketBrowserModel

`func NewS3CompatibleBucketBrowserModel() *S3CompatibleBucketBrowserModel`

NewS3CompatibleBucketBrowserModel instantiates a new S3CompatibleBucketBrowserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3CompatibleBucketBrowserModelWithDefaults

`func NewS3CompatibleBucketBrowserModelWithDefaults() *S3CompatibleBucketBrowserModel`

NewS3CompatibleBucketBrowserModelWithDefaults instantiates a new S3CompatibleBucketBrowserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *S3CompatibleBucketBrowserModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *S3CompatibleBucketBrowserModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *S3CompatibleBucketBrowserModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *S3CompatibleBucketBrowserModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFolders

`func (o *S3CompatibleBucketBrowserModel) GetFolders() []string`

GetFolders returns the Folders field if non-nil, zero value otherwise.

### GetFoldersOk

`func (o *S3CompatibleBucketBrowserModel) GetFoldersOk() (*[]string, bool)`

GetFoldersOk returns a tuple with the Folders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolders

`func (o *S3CompatibleBucketBrowserModel) SetFolders(v []string)`

SetFolders sets Folders field to given value.

### HasFolders

`func (o *S3CompatibleBucketBrowserModel) HasFolders() bool`

HasFolders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


