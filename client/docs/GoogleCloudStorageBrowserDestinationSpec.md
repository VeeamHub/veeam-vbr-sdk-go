# GoogleCloudStorageBrowserDestinationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | Pointer to **string** | ID of a server you want to use to connect to the object storage. You can use the backup server or any Microsoft Windows or Linux server added to your backup infrastructure. By default, the backup server ID is used. | [optional] 
**RegionId** | **string** | Data center region where the bucket is located. | 
**BucketName** | **string** | Name of the bucket where you want to store your backup data. | 

## Methods

### NewGoogleCloudStorageBrowserDestinationSpec

`func NewGoogleCloudStorageBrowserDestinationSpec(regionId string, bucketName string, ) *GoogleCloudStorageBrowserDestinationSpec`

NewGoogleCloudStorageBrowserDestinationSpec instantiates a new GoogleCloudStorageBrowserDestinationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudStorageBrowserDestinationSpecWithDefaults

`func NewGoogleCloudStorageBrowserDestinationSpecWithDefaults() *GoogleCloudStorageBrowserDestinationSpec`

NewGoogleCloudStorageBrowserDestinationSpecWithDefaults instantiates a new GoogleCloudStorageBrowserDestinationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *GoogleCloudStorageBrowserDestinationSpec) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *GoogleCloudStorageBrowserDestinationSpec) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *GoogleCloudStorageBrowserDestinationSpec) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *GoogleCloudStorageBrowserDestinationSpec) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetRegionId

`func (o *GoogleCloudStorageBrowserDestinationSpec) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *GoogleCloudStorageBrowserDestinationSpec) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *GoogleCloudStorageBrowserDestinationSpec) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetBucketName

`func (o *GoogleCloudStorageBrowserDestinationSpec) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *GoogleCloudStorageBrowserDestinationSpec) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *GoogleCloudStorageBrowserDestinationSpec) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


