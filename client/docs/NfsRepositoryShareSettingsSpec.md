# NfsRepositoryShareSettingsSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharePath** | **string** | Path to the shared folder that that is used as a backup repository. | 
**GatewayServer** | Pointer to [**RepositoryShareGatewayImportSpec**](RepositoryShareGatewayImportSpec.md) |  | [optional] 

## Methods

### NewNfsRepositoryShareSettingsSpec

`func NewNfsRepositoryShareSettingsSpec(sharePath string, ) *NfsRepositoryShareSettingsSpec`

NewNfsRepositoryShareSettingsSpec instantiates a new NfsRepositoryShareSettingsSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfsRepositoryShareSettingsSpecWithDefaults

`func NewNfsRepositoryShareSettingsSpecWithDefaults() *NfsRepositoryShareSettingsSpec`

NewNfsRepositoryShareSettingsSpecWithDefaults instantiates a new NfsRepositoryShareSettingsSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharePath

`func (o *NfsRepositoryShareSettingsSpec) GetSharePath() string`

GetSharePath returns the SharePath field if non-nil, zero value otherwise.

### GetSharePathOk

`func (o *NfsRepositoryShareSettingsSpec) GetSharePathOk() (*string, bool)`

GetSharePathOk returns a tuple with the SharePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePath

`func (o *NfsRepositoryShareSettingsSpec) SetSharePath(v string)`

SetSharePath sets SharePath field to given value.


### GetGatewayServer

`func (o *NfsRepositoryShareSettingsSpec) GetGatewayServer() RepositoryShareGatewayImportSpec`

GetGatewayServer returns the GatewayServer field if non-nil, zero value otherwise.

### GetGatewayServerOk

`func (o *NfsRepositoryShareSettingsSpec) GetGatewayServerOk() (*RepositoryShareGatewayImportSpec, bool)`

GetGatewayServerOk returns a tuple with the GatewayServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayServer

`func (o *NfsRepositoryShareSettingsSpec) SetGatewayServer(v RepositoryShareGatewayImportSpec)`

SetGatewayServer sets GatewayServer field to given value.

### HasGatewayServer

`func (o *NfsRepositoryShareSettingsSpec) HasGatewayServer() bool`

HasGatewayServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


