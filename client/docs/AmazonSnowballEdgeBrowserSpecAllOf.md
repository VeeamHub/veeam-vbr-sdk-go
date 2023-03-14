# AmazonSnowballEdgeBrowserSpecAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionId** | Pointer to **string** | To connect to the AWS Snowball Edge device, specify the &#x60;snow&#x60; value. | [optional] 
**ServicePoint** | Pointer to **string** | Service point address and port number of the AWS Snowball Edge device. | [optional] 
**GatewayServerId** | Pointer to **string** | ID of a gateway server you want to use to connect to the object storage. Specify this parameter to check internet connection of the server. As a gateway server you can use the backup server or any Microsoft Windows or Linux server added to your backup infrastructure. By default, the backup server ID is used. | [optional] 

## Methods

### NewAmazonSnowballEdgeBrowserSpecAllOf

`func NewAmazonSnowballEdgeBrowserSpecAllOf() *AmazonSnowballEdgeBrowserSpecAllOf`

NewAmazonSnowballEdgeBrowserSpecAllOf instantiates a new AmazonSnowballEdgeBrowserSpecAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonSnowballEdgeBrowserSpecAllOfWithDefaults

`func NewAmazonSnowballEdgeBrowserSpecAllOfWithDefaults() *AmazonSnowballEdgeBrowserSpecAllOf`

NewAmazonSnowballEdgeBrowserSpecAllOfWithDefaults instantiates a new AmazonSnowballEdgeBrowserSpecAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionId

`func (o *AmazonSnowballEdgeBrowserSpecAllOf) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *AmazonSnowballEdgeBrowserSpecAllOf) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *AmazonSnowballEdgeBrowserSpecAllOf) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *AmazonSnowballEdgeBrowserSpecAllOf) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetServicePoint

`func (o *AmazonSnowballEdgeBrowserSpecAllOf) GetServicePoint() string`

GetServicePoint returns the ServicePoint field if non-nil, zero value otherwise.

### GetServicePointOk

`func (o *AmazonSnowballEdgeBrowserSpecAllOf) GetServicePointOk() (*string, bool)`

GetServicePointOk returns a tuple with the ServicePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePoint

`func (o *AmazonSnowballEdgeBrowserSpecAllOf) SetServicePoint(v string)`

SetServicePoint sets ServicePoint field to given value.

### HasServicePoint

`func (o *AmazonSnowballEdgeBrowserSpecAllOf) HasServicePoint() bool`

HasServicePoint returns a boolean if a field has been set.

### GetGatewayServerId

`func (o *AmazonSnowballEdgeBrowserSpecAllOf) GetGatewayServerId() string`

GetGatewayServerId returns the GatewayServerId field if non-nil, zero value otherwise.

### GetGatewayServerIdOk

`func (o *AmazonSnowballEdgeBrowserSpecAllOf) GetGatewayServerIdOk() (*string, bool)`

GetGatewayServerIdOk returns a tuple with the GatewayServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayServerId

`func (o *AmazonSnowballEdgeBrowserSpecAllOf) SetGatewayServerId(v string)`

SetGatewayServerId sets GatewayServerId field to given value.

### HasGatewayServerId

`func (o *AmazonSnowballEdgeBrowserSpecAllOf) HasGatewayServerId() bool`

HasGatewayServerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


