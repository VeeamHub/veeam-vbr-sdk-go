# AmazonSnowballEdgeBrowserDestinationSpecAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostId** | Pointer to **string** | ID of a server you want to use to connect to the object storage. You can use the backup server or any Microsoft Windows or Linux server added to your backup infrastructure. By default, the backup server ID is used. | [optional] 
**ConnectionPoint** | **string** | Service point address and port number of the AWS Snowball Edge device. | 
**RegionId** | **string** | For AWS Snowball Edge, the region is *snow*. | 
**BucketName** | **string** | Name of the bucket where you want to store your backup data. | 

## Methods

### NewAmazonSnowballEdgeBrowserDestinationSpecAllOf

`func NewAmazonSnowballEdgeBrowserDestinationSpecAllOf(connectionPoint string, regionId string, bucketName string, ) *AmazonSnowballEdgeBrowserDestinationSpecAllOf`

NewAmazonSnowballEdgeBrowserDestinationSpecAllOf instantiates a new AmazonSnowballEdgeBrowserDestinationSpecAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonSnowballEdgeBrowserDestinationSpecAllOfWithDefaults

`func NewAmazonSnowballEdgeBrowserDestinationSpecAllOfWithDefaults() *AmazonSnowballEdgeBrowserDestinationSpecAllOf`

NewAmazonSnowballEdgeBrowserDestinationSpecAllOfWithDefaults instantiates a new AmazonSnowballEdgeBrowserDestinationSpecAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostId

`func (o *AmazonSnowballEdgeBrowserDestinationSpecAllOf) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *AmazonSnowballEdgeBrowserDestinationSpecAllOf) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *AmazonSnowballEdgeBrowserDestinationSpecAllOf) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *AmazonSnowballEdgeBrowserDestinationSpecAllOf) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetConnectionPoint

`func (o *AmazonSnowballEdgeBrowserDestinationSpecAllOf) GetConnectionPoint() string`

GetConnectionPoint returns the ConnectionPoint field if non-nil, zero value otherwise.

### GetConnectionPointOk

`func (o *AmazonSnowballEdgeBrowserDestinationSpecAllOf) GetConnectionPointOk() (*string, bool)`

GetConnectionPointOk returns a tuple with the ConnectionPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPoint

`func (o *AmazonSnowballEdgeBrowserDestinationSpecAllOf) SetConnectionPoint(v string)`

SetConnectionPoint sets ConnectionPoint field to given value.


### GetRegionId

`func (o *AmazonSnowballEdgeBrowserDestinationSpecAllOf) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *AmazonSnowballEdgeBrowserDestinationSpecAllOf) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *AmazonSnowballEdgeBrowserDestinationSpecAllOf) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetBucketName

`func (o *AmazonSnowballEdgeBrowserDestinationSpecAllOf) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *AmazonSnowballEdgeBrowserDestinationSpecAllOf) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *AmazonSnowballEdgeBrowserDestinationSpecAllOf) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


