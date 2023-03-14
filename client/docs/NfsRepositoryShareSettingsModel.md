# NfsRepositoryShareSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharePath** | **string** | Path to the shared folder that is used as a backup repository. | 
**GatewayServer** | Pointer to [**RepositoryShareGatewayModel**](RepositoryShareGatewayModel.md) |  | [optional] 

## Methods

### NewNfsRepositoryShareSettingsModel

`func NewNfsRepositoryShareSettingsModel(sharePath string, ) *NfsRepositoryShareSettingsModel`

NewNfsRepositoryShareSettingsModel instantiates a new NfsRepositoryShareSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfsRepositoryShareSettingsModelWithDefaults

`func NewNfsRepositoryShareSettingsModelWithDefaults() *NfsRepositoryShareSettingsModel`

NewNfsRepositoryShareSettingsModelWithDefaults instantiates a new NfsRepositoryShareSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharePath

`func (o *NfsRepositoryShareSettingsModel) GetSharePath() string`

GetSharePath returns the SharePath field if non-nil, zero value otherwise.

### GetSharePathOk

`func (o *NfsRepositoryShareSettingsModel) GetSharePathOk() (*string, bool)`

GetSharePathOk returns a tuple with the SharePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePath

`func (o *NfsRepositoryShareSettingsModel) SetSharePath(v string)`

SetSharePath sets SharePath field to given value.


### GetGatewayServer

`func (o *NfsRepositoryShareSettingsModel) GetGatewayServer() RepositoryShareGatewayModel`

GetGatewayServer returns the GatewayServer field if non-nil, zero value otherwise.

### GetGatewayServerOk

`func (o *NfsRepositoryShareSettingsModel) GetGatewayServerOk() (*RepositoryShareGatewayModel, bool)`

GetGatewayServerOk returns a tuple with the GatewayServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayServer

`func (o *NfsRepositoryShareSettingsModel) SetGatewayServer(v RepositoryShareGatewayModel)`

SetGatewayServer sets GatewayServer field to given value.

### HasGatewayServer

`func (o *NfsRepositoryShareSettingsModel) HasGatewayServer() bool`

HasGatewayServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


