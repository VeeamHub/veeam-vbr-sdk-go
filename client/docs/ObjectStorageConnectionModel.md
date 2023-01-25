# ObjectStorageConnectionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | [**ERepositoryConnectionType**](ERepositoryConnectionType.md) |  | 
**GatewayServerIds** | Pointer to **[]string** | Array of gateway server IDs. The value is *null* if the connection type is *Direct*. | [optional] 

## Methods

### NewObjectStorageConnectionModel

`func NewObjectStorageConnectionModel(connectionType ERepositoryConnectionType, ) *ObjectStorageConnectionModel`

NewObjectStorageConnectionModel instantiates a new ObjectStorageConnectionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageConnectionModelWithDefaults

`func NewObjectStorageConnectionModelWithDefaults() *ObjectStorageConnectionModel`

NewObjectStorageConnectionModelWithDefaults instantiates a new ObjectStorageConnectionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *ObjectStorageConnectionModel) GetConnectionType() ERepositoryConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *ObjectStorageConnectionModel) GetConnectionTypeOk() (*ERepositoryConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *ObjectStorageConnectionModel) SetConnectionType(v ERepositoryConnectionType)`

SetConnectionType sets ConnectionType field to given value.


### GetGatewayServerIds

`func (o *ObjectStorageConnectionModel) GetGatewayServerIds() []string`

GetGatewayServerIds returns the GatewayServerIds field if non-nil, zero value otherwise.

### GetGatewayServerIdsOk

`func (o *ObjectStorageConnectionModel) GetGatewayServerIdsOk() (*[]string, bool)`

GetGatewayServerIdsOk returns a tuple with the GatewayServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayServerIds

`func (o *ObjectStorageConnectionModel) SetGatewayServerIds(v []string)`

SetGatewayServerIds sets GatewayServerIds field to given value.

### HasGatewayServerIds

`func (o *ObjectStorageConnectionModel) HasGatewayServerIds() bool`

HasGatewayServerIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


