# AmazonSnowballEdgeBrowserDestinationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | Pointer to **string** | ID of a server you want to use to connect to the object storage. You can use the backup server or any Microsoft Windows or Linux server added to your backup infrastructure. By default, the backup server ID is used. | [optional] 
**ConnectionPoint** | **string** | Service point address and port number of the AWS Snowball Edge device. | 
**RegionId** | **string** | For AWS Snowball Edge, the region is *snow*. | 
**BucketName** | **string** | Name of the bucket where you want to store your backup data. | 

## Methods

### NewAmazonSnowballEdgeBrowserDestinationSpec

`func NewAmazonSnowballEdgeBrowserDestinationSpec(connectionPoint string, regionId string, bucketName string, ) *AmazonSnowballEdgeBrowserDestinationSpec`

NewAmazonSnowballEdgeBrowserDestinationSpec instantiates a new AmazonSnowballEdgeBrowserDestinationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonSnowballEdgeBrowserDestinationSpecWithDefaults

`func NewAmazonSnowballEdgeBrowserDestinationSpecWithDefaults() *AmazonSnowballEdgeBrowserDestinationSpec`

NewAmazonSnowballEdgeBrowserDestinationSpecWithDefaults instantiates a new AmazonSnowballEdgeBrowserDestinationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *AmazonSnowballEdgeBrowserDestinationSpec) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *AmazonSnowballEdgeBrowserDestinationSpec) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *AmazonSnowballEdgeBrowserDestinationSpec) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *AmazonSnowballEdgeBrowserDestinationSpec) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetConnectionPoint

`func (o *AmazonSnowballEdgeBrowserDestinationSpec) GetConnectionPoint() string`

GetConnectionPoint returns the ConnectionPoint field if non-nil, zero value otherwise.

### GetConnectionPointOk

`func (o *AmazonSnowballEdgeBrowserDestinationSpec) GetConnectionPointOk() (*string, bool)`

GetConnectionPointOk returns a tuple with the ConnectionPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPoint

`func (o *AmazonSnowballEdgeBrowserDestinationSpec) SetConnectionPoint(v string)`

SetConnectionPoint sets ConnectionPoint field to given value.


### GetRegionId

`func (o *AmazonSnowballEdgeBrowserDestinationSpec) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *AmazonSnowballEdgeBrowserDestinationSpec) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *AmazonSnowballEdgeBrowserDestinationSpec) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetBucketName

`func (o *AmazonSnowballEdgeBrowserDestinationSpec) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *AmazonSnowballEdgeBrowserDestinationSpec) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *AmazonSnowballEdgeBrowserDestinationSpec) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


