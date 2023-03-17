# SmbRepositoryShareSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharePath** | **string** | Path to the shared folder that is used as a backup repository. | 
**CredentialsId** | **string** | ID of the credentials record used to access the shared folder. | 
**GatewayServer** | Pointer to [**RepositoryShareGatewayModel**](RepositoryShareGatewayModel.md) |  | [optional] 

## Methods

### NewSmbRepositoryShareSettingsModel

`func NewSmbRepositoryShareSettingsModel(sharePath string, credentialsId string, ) *SmbRepositoryShareSettingsModel`

NewSmbRepositoryShareSettingsModel instantiates a new SmbRepositoryShareSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmbRepositoryShareSettingsModelWithDefaults

`func NewSmbRepositoryShareSettingsModelWithDefaults() *SmbRepositoryShareSettingsModel`

NewSmbRepositoryShareSettingsModelWithDefaults instantiates a new SmbRepositoryShareSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharePath

`func (o *SmbRepositoryShareSettingsModel) GetSharePath() string`

GetSharePath returns the SharePath field if non-nil, zero value otherwise.

### GetSharePathOk

`func (o *SmbRepositoryShareSettingsModel) GetSharePathOk() (*string, bool)`

GetSharePathOk returns a tuple with the SharePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePath

`func (o *SmbRepositoryShareSettingsModel) SetSharePath(v string)`

SetSharePath sets SharePath field to given value.


### GetCredentialsId

`func (o *SmbRepositoryShareSettingsModel) GetCredentialsId() string`

GetCredentialsId returns the CredentialsId field if non-nil, zero value otherwise.

### GetCredentialsIdOk

`func (o *SmbRepositoryShareSettingsModel) GetCredentialsIdOk() (*string, bool)`

GetCredentialsIdOk returns a tuple with the CredentialsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsId

`func (o *SmbRepositoryShareSettingsModel) SetCredentialsId(v string)`

SetCredentialsId sets CredentialsId field to given value.


### GetGatewayServer

`func (o *SmbRepositoryShareSettingsModel) GetGatewayServer() RepositoryShareGatewayModel`

GetGatewayServer returns the GatewayServer field if non-nil, zero value otherwise.

### GetGatewayServerOk

`func (o *SmbRepositoryShareSettingsModel) GetGatewayServerOk() (*RepositoryShareGatewayModel, bool)`

GetGatewayServerOk returns a tuple with the GatewayServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayServer

`func (o *SmbRepositoryShareSettingsModel) SetGatewayServer(v RepositoryShareGatewayModel)`

SetGatewayServer sets GatewayServer field to given value.

### HasGatewayServer

`func (o *SmbRepositoryShareSettingsModel) HasGatewayServer() bool`

HasGatewayServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


