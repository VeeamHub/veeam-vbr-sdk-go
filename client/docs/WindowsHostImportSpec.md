# WindowsHostImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Full DNS name or IP address of the server. | 
**Description** | **string** | Description of the server. | 
**Type** | [**EManagedServerType**](EManagedServerType.md) |  | 
**Credentials** | Pointer to [**CredentialsImportModel**](CredentialsImportModel.md) |  | [optional] 
**NetworkSettings** | Pointer to [**WindowsHostPortsModel**](WindowsHostPortsModel.md) |  | [optional] 

## Methods

### NewWindowsHostImportSpec

`func NewWindowsHostImportSpec(name string, description string, type_ EManagedServerType, ) *WindowsHostImportSpec`

NewWindowsHostImportSpec instantiates a new WindowsHostImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWindowsHostImportSpecWithDefaults

`func NewWindowsHostImportSpecWithDefaults() *WindowsHostImportSpec`

NewWindowsHostImportSpecWithDefaults instantiates a new WindowsHostImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WindowsHostImportSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WindowsHostImportSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WindowsHostImportSpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WindowsHostImportSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WindowsHostImportSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WindowsHostImportSpec) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *WindowsHostImportSpec) GetType() EManagedServerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WindowsHostImportSpec) GetTypeOk() (*EManagedServerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WindowsHostImportSpec) SetType(v EManagedServerType)`

SetType sets Type field to given value.


### GetCredentials

`func (o *WindowsHostImportSpec) GetCredentials() CredentialsImportModel`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *WindowsHostImportSpec) GetCredentialsOk() (*CredentialsImportModel, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *WindowsHostImportSpec) SetCredentials(v CredentialsImportModel)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *WindowsHostImportSpec) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetNetworkSettings

`func (o *WindowsHostImportSpec) GetNetworkSettings() WindowsHostPortsModel`

GetNetworkSettings returns the NetworkSettings field if non-nil, zero value otherwise.

### GetNetworkSettingsOk

`func (o *WindowsHostImportSpec) GetNetworkSettingsOk() (*WindowsHostPortsModel, bool)`

GetNetworkSettingsOk returns a tuple with the NetworkSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSettings

`func (o *WindowsHostImportSpec) SetNetworkSettings(v WindowsHostPortsModel)`

SetNetworkSettings sets NetworkSettings field to given value.

### HasNetworkSettings

`func (o *WindowsHostImportSpec) HasNetworkSettings() bool`

HasNetworkSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


