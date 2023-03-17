# LinuxHostModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshSettings** | Pointer to [**LinuxHostSSHSettingsModel**](LinuxHostSSHSettingsModel.md) |  | [optional] 
**CredentialsStorageType** | Pointer to [**ECredentialsStorageType**](ECredentialsStorageType.md) |  | [optional] 

## Methods

### NewLinuxHostModel

`func NewLinuxHostModel() *LinuxHostModel`

NewLinuxHostModel instantiates a new LinuxHostModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinuxHostModelWithDefaults

`func NewLinuxHostModelWithDefaults() *LinuxHostModel`

NewLinuxHostModelWithDefaults instantiates a new LinuxHostModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSshSettings

`func (o *LinuxHostModel) GetSshSettings() LinuxHostSSHSettingsModel`

GetSshSettings returns the SshSettings field if non-nil, zero value otherwise.

### GetSshSettingsOk

`func (o *LinuxHostModel) GetSshSettingsOk() (*LinuxHostSSHSettingsModel, bool)`

GetSshSettingsOk returns a tuple with the SshSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshSettings

`func (o *LinuxHostModel) SetSshSettings(v LinuxHostSSHSettingsModel)`

SetSshSettings sets SshSettings field to given value.

### HasSshSettings

`func (o *LinuxHostModel) HasSshSettings() bool`

HasSshSettings returns a boolean if a field has been set.

### GetCredentialsStorageType

`func (o *LinuxHostModel) GetCredentialsStorageType() ECredentialsStorageType`

GetCredentialsStorageType returns the CredentialsStorageType field if non-nil, zero value otherwise.

### GetCredentialsStorageTypeOk

`func (o *LinuxHostModel) GetCredentialsStorageTypeOk() (*ECredentialsStorageType, bool)`

GetCredentialsStorageTypeOk returns a tuple with the CredentialsStorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsStorageType

`func (o *LinuxHostModel) SetCredentialsStorageType(v ECredentialsStorageType)`

SetCredentialsStorageType sets CredentialsStorageType field to given value.

### HasCredentialsStorageType

`func (o *LinuxHostModel) HasCredentialsStorageType() bool`

HasCredentialsStorageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


