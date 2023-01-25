# S3CompatibleBrowserDestinationSpecAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | Pointer to **string** | ID of a server you want to use to connect to the object storage. You can use the backup server or any Microsoft Windows or Linux server added to your backup infrastructure. By default, the backup server ID is used. | [optional] 
**ConnectionPoint** | **string** | Endpoint address and port number of the S3 compatible storage. | 
**RegionId** | **string** | Region where the bucket is located. | 
**BucketName** | **string** | Name of the bucket where you want to store your backup data. | 

## Methods

### NewS3CompatibleBrowserDestinationSpecAllOf

`func NewS3CompatibleBrowserDestinationSpecAllOf(connectionPoint string, regionId string, bucketName string, ) *S3CompatibleBrowserDestinationSpecAllOf`

NewS3CompatibleBrowserDestinationSpecAllOf instantiates a new S3CompatibleBrowserDestinationSpecAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3CompatibleBrowserDestinationSpecAllOfWithDefaults

`func NewS3CompatibleBrowserDestinationSpecAllOfWithDefaults() *S3CompatibleBrowserDestinationSpecAllOf`

NewS3CompatibleBrowserDestinationSpecAllOfWithDefaults instantiates a new S3CompatibleBrowserDestinationSpecAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *S3CompatibleBrowserDestinationSpecAllOf) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *S3CompatibleBrowserDestinationSpecAllOf) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *S3CompatibleBrowserDestinationSpecAllOf) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *S3CompatibleBrowserDestinationSpecAllOf) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetConnectionPoint

`func (o *S3CompatibleBrowserDestinationSpecAllOf) GetConnectionPoint() string`

GetConnectionPoint returns the ConnectionPoint field if non-nil, zero value otherwise.

### GetConnectionPointOk

`func (o *S3CompatibleBrowserDestinationSpecAllOf) GetConnectionPointOk() (*string, bool)`

GetConnectionPointOk returns a tuple with the ConnectionPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPoint

`func (o *S3CompatibleBrowserDestinationSpecAllOf) SetConnectionPoint(v string)`

SetConnectionPoint sets ConnectionPoint field to given value.


### GetRegionId

`func (o *S3CompatibleBrowserDestinationSpecAllOf) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *S3CompatibleBrowserDestinationSpecAllOf) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *S3CompatibleBrowserDestinationSpecAllOf) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetBucketName

`func (o *S3CompatibleBrowserDestinationSpecAllOf) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *S3CompatibleBrowserDestinationSpecAllOf) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *S3CompatibleBrowserDestinationSpecAllOf) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


