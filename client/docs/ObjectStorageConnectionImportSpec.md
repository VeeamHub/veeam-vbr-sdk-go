# ObjectStorageConnectionImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | [**ERepositoryConnectionType**](ERepositoryConnectionType.md) |  | 
**GatewayServers** | Pointer to **[]string** | Array of gateway server IDs. The value is *null* if the connection type is *Direct*. | [optional] 

## Methods

### NewObjectStorageConnectionImportSpec

`func NewObjectStorageConnectionImportSpec(connectionType ERepositoryConnectionType, ) *ObjectStorageConnectionImportSpec`

NewObjectStorageConnectionImportSpec instantiates a new ObjectStorageConnectionImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStorageConnectionImportSpecWithDefaults

`func NewObjectStorageConnectionImportSpecWithDefaults() *ObjectStorageConnectionImportSpec`

NewObjectStorageConnectionImportSpecWithDefaults instantiates a new ObjectStorageConnectionImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *ObjectStorageConnectionImportSpec) GetConnectionType() ERepositoryConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *ObjectStorageConnectionImportSpec) GetConnectionTypeOk() (*ERepositoryConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *ObjectStorageConnectionImportSpec) SetConnectionType(v ERepositoryConnectionType)`

SetConnectionType sets ConnectionType field to given value.


### GetGatewayServers

`func (o *ObjectStorageConnectionImportSpec) GetGatewayServers() []string`

GetGatewayServers returns the GatewayServers field if non-nil, zero value otherwise.

### GetGatewayServersOk

`func (o *ObjectStorageConnectionImportSpec) GetGatewayServersOk() (*[]string, bool)`

GetGatewayServersOk returns a tuple with the GatewayServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayServers

`func (o *ObjectStorageConnectionImportSpec) SetGatewayServers(v []string)`

SetGatewayServers sets GatewayServers field to given value.

### HasGatewayServers

`func (o *ObjectStorageConnectionImportSpec) HasGatewayServers() bool`

HasGatewayServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


