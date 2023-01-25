# ManagedServerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the server. | 
**Name** | **string** | Full DNS name or IP address of the server. | 
**Description** | **string** | Description of the server. | 
**Type** | [**EManagedServerType**](EManagedServerType.md) |  | 
**CredentialsId** | **string** | ID of a credentials record used to connect to the server. | 
**NetworkSettings** | Pointer to [**WindowsHostPortsModel**](WindowsHostPortsModel.md) |  | [optional] 
**ViHostType** | Pointer to [**EViHostType**](EViHostType.md) |  | [optional] 
**Port** | **int32** | Port used to communicate with the server. | 
**SshSettings** | Pointer to [**LinuxHostSSHSettingsModel**](LinuxHostSSHSettingsModel.md) |  | [optional] 
**CredentialsStorageType** | Pointer to [**ECredentialsStorageType**](ECredentialsStorageType.md) |  | [optional] 

## Methods

### NewManagedServerModel

`func NewManagedServerModel(id string, name string, description string, type_ EManagedServerType, credentialsId string, port int32, ) *ManagedServerModel`

NewManagedServerModel instantiates a new ManagedServerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedServerModelWithDefaults

`func NewManagedServerModelWithDefaults() *ManagedServerModel`

NewManagedServerModelWithDefaults instantiates a new ManagedServerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManagedServerModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagedServerModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagedServerModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ManagedServerModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagedServerModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagedServerModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ManagedServerModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManagedServerModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManagedServerModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *ManagedServerModel) GetType() EManagedServerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManagedServerModel) GetTypeOk() (*EManagedServerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManagedServerModel) SetType(v EManagedServerType)`

SetType sets Type field to given value.


### GetCredentialsId

`func (o *ManagedServerModel) GetCredentialsId() string`

GetCredentialsId returns the CredentialsId field if non-nil, zero value otherwise.

### GetCredentialsIdOk

`func (o *ManagedServerModel) GetCredentialsIdOk() (*string, bool)`

GetCredentialsIdOk returns a tuple with the CredentialsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsId

`func (o *ManagedServerModel) SetCredentialsId(v string)`

SetCredentialsId sets CredentialsId field to given value.


### GetNetworkSettings

`func (o *ManagedServerModel) GetNetworkSettings() WindowsHostPortsModel`

GetNetworkSettings returns the NetworkSettings field if non-nil, zero value otherwise.

### GetNetworkSettingsOk

`func (o *ManagedServerModel) GetNetworkSettingsOk() (*WindowsHostPortsModel, bool)`

GetNetworkSettingsOk returns a tuple with the NetworkSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSettings

`func (o *ManagedServerModel) SetNetworkSettings(v WindowsHostPortsModel)`

SetNetworkSettings sets NetworkSettings field to given value.

### HasNetworkSettings

`func (o *ManagedServerModel) HasNetworkSettings() bool`

HasNetworkSettings returns a boolean if a field has been set.

### GetViHostType

`func (o *ManagedServerModel) GetViHostType() EViHostType`

GetViHostType returns the ViHostType field if non-nil, zero value otherwise.

### GetViHostTypeOk

`func (o *ManagedServerModel) GetViHostTypeOk() (*EViHostType, bool)`

GetViHostTypeOk returns a tuple with the ViHostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViHostType

`func (o *ManagedServerModel) SetViHostType(v EViHostType)`

SetViHostType sets ViHostType field to given value.

### HasViHostType

`func (o *ManagedServerModel) HasViHostType() bool`

HasViHostType returns a boolean if a field has been set.

### GetPort

`func (o *ManagedServerModel) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ManagedServerModel) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ManagedServerModel) SetPort(v int32)`

SetPort sets Port field to given value.


### GetSshSettings

`func (o *ManagedServerModel) GetSshSettings() LinuxHostSSHSettingsModel`

GetSshSettings returns the SshSettings field if non-nil, zero value otherwise.

### GetSshSettingsOk

`func (o *ManagedServerModel) GetSshSettingsOk() (*LinuxHostSSHSettingsModel, bool)`

GetSshSettingsOk returns a tuple with the SshSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshSettings

`func (o *ManagedServerModel) SetSshSettings(v LinuxHostSSHSettingsModel)`

SetSshSettings sets SshSettings field to given value.

### HasSshSettings

`func (o *ManagedServerModel) HasSshSettings() bool`

HasSshSettings returns a boolean if a field has been set.

### GetCredentialsStorageType

`func (o *ManagedServerModel) GetCredentialsStorageType() ECredentialsStorageType`

GetCredentialsStorageType returns the CredentialsStorageType field if non-nil, zero value otherwise.

### GetCredentialsStorageTypeOk

`func (o *ManagedServerModel) GetCredentialsStorageTypeOk() (*ECredentialsStorageType, bool)`

GetCredentialsStorageTypeOk returns a tuple with the CredentialsStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsStorageType

`func (o *ManagedServerModel) SetCredentialsStorageType(v ECredentialsStorageType)`

SetCredentialsStorageType sets CredentialsStorageType field to given value.

### HasCredentialsStorageType

`func (o *ManagedServerModel) HasCredentialsStorageType() bool`

HasCredentialsStorageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


