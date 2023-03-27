# WasabiCloudStorageBrowserSpecAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionId** | **string** | Region ID. | 
**GatewayServerId** | Pointer to **string** | ID of a gateway server you want to use to connect to the object storage. Specify this parameter to check internet connection of the server. As a gateway server you can use the backup server or any Microsoft Windows or Linux server added to your backup infrastructure. By default, the backup server ID is used. | [optional] 

## Methods

### NewWasabiCloudStorageBrowserSpecAllOf

`func NewWasabiCloudStorageBrowserSpecAllOf(regionId string, ) *WasabiCloudStorageBrowserSpecAllOf`

NewWasabiCloudStorageBrowserSpecAllOf instantiates a new WasabiCloudStorageBrowserSpecAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWasabiCloudStorageBrowserSpecAllOfWithDefaults

`func NewWasabiCloudStorageBrowserSpecAllOfWithDefaults() *WasabiCloudStorageBrowserSpecAllOf`

NewWasabiCloudStorageBrowserSpecAllOfWithDefaults instantiates a new WasabiCloudStorageBrowserSpecAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionId

`func (o *WasabiCloudStorageBrowserSpecAllOf) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *WasabiCloudStorageBrowserSpecAllOf) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *WasabiCloudStorageBrowserSpecAllOf) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetGatewayServerId

`func (o *WasabiCloudStorageBrowserSpecAllOf) GetGatewayServerId() string`

GetGatewayServerId returns the GatewayServerId field if non-nil, zero value otherwise.

### GetGatewayServerIdOk

`func (o *WasabiCloudStorageBrowserSpecAllOf) GetGatewayServerIdOk() (*string, bool)`

GetGatewayServerIdOk returns a tuple with the GatewayServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayServerId

`func (o *WasabiCloudStorageBrowserSpecAllOf) SetGatewayServerId(v string)`

SetGatewayServerId sets GatewayServerId field to given value.

### HasGatewayServerId

`func (o *WasabiCloudStorageBrowserSpecAllOf) HasGatewayServerId() bool`

HasGatewayServerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


