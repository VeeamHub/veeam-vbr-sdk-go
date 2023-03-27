# SmbRepositoryShareSettingsSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharePath** | **string** | Path to the shared folder that is used as a backup repository. | 
**Credentials** | [**CredentialsImportModel**](CredentialsImportModel.md) |  | 
**GatewayServer** | Pointer to [**RepositoryShareGatewayImportSpec**](RepositoryShareGatewayImportSpec.md) |  | [optional] 

## Methods

### NewSmbRepositoryShareSettingsSpec

`func NewSmbRepositoryShareSettingsSpec(sharePath string, credentials CredentialsImportModel, ) *SmbRepositoryShareSettingsSpec`

NewSmbRepositoryShareSettingsSpec instantiates a new SmbRepositoryShareSettingsSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmbRepositoryShareSettingsSpecWithDefaults

`func NewSmbRepositoryShareSettingsSpecWithDefaults() *SmbRepositoryShareSettingsSpec`

NewSmbRepositoryShareSettingsSpecWithDefaults instantiates a new SmbRepositoryShareSettingsSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharePath

`func (o *SmbRepositoryShareSettingsSpec) GetSharePath() string`

GetSharePath returns the SharePath field if non-nil, zero value otherwise.

### GetSharePathOk

`func (o *SmbRepositoryShareSettingsSpec) GetSharePathOk() (*string, bool)`

GetSharePathOk returns a tuple with the SharePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePath

`func (o *SmbRepositoryShareSettingsSpec) SetSharePath(v string)`

SetSharePath sets SharePath field to given value.


### GetCredentials

`func (o *SmbRepositoryShareSettingsSpec) GetCredentials() CredentialsImportModel`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *SmbRepositoryShareSettingsSpec) GetCredentialsOk() (*CredentialsImportModel, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *SmbRepositoryShareSettingsSpec) SetCredentials(v CredentialsImportModel)`

SetCredentials sets Credentials field to given value.


### GetGatewayServer

`func (o *SmbRepositoryShareSettingsSpec) GetGatewayServer() RepositoryShareGatewayImportSpec`

GetGatewayServer returns the GatewayServer field if non-nil, zero value otherwise.

### GetGatewayServerOk

`func (o *SmbRepositoryShareSettingsSpec) GetGatewayServerOk() (*RepositoryShareGatewayImportSpec, bool)`

GetGatewayServerOk returns a tuple with the GatewayServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayServer

`func (o *SmbRepositoryShareSettingsSpec) SetGatewayServer(v RepositoryShareGatewayImportSpec)`

SetGatewayServer sets GatewayServer field to given value.

### HasGatewayServer

`func (o *SmbRepositoryShareSettingsSpec) HasGatewayServer() bool`

HasGatewayServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


