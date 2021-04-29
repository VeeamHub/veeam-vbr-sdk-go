# LinuxHostImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Full DNS name or IP address of the server. | 
**Description** | **string** | Description of the server. | 
**Type** | [**EManagedServerType**](EManagedServerType.md) |  | 
**Credentials** | [**CredentialsImportModel**](CredentialsImportModel.md) |  | 
**SshSettings** | Pointer to [**LinuxHostSSHSettingsModel**](LinuxHostSSHSettingsModel.md) |  | [optional] 
**SshFingerprint** | **string** | SSH key fingerprint used to verify the server identity. | 

## Methods

### NewLinuxHostImportSpec

`func NewLinuxHostImportSpec(name string, description string, type_ EManagedServerType, credentials CredentialsImportModel, sshFingerprint string, ) *LinuxHostImportSpec`

NewLinuxHostImportSpec instantiates a new LinuxHostImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinuxHostImportSpecWithDefaults

`func NewLinuxHostImportSpecWithDefaults() *LinuxHostImportSpec`

NewLinuxHostImportSpecWithDefaults instantiates a new LinuxHostImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LinuxHostImportSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LinuxHostImportSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LinuxHostImportSpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *LinuxHostImportSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LinuxHostImportSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LinuxHostImportSpec) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetType

`func (o *LinuxHostImportSpec) GetType() EManagedServerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinuxHostImportSpec) GetTypeOk() (*EManagedServerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinuxHostImportSpec) SetType(v EManagedServerType)`

SetType sets Type field to given value.


### GetCredentials

`func (o *LinuxHostImportSpec) GetCredentials() CredentialsImportModel`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *LinuxHostImportSpec) GetCredentialsOk() (*CredentialsImportModel, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *LinuxHostImportSpec) SetCredentials(v CredentialsImportModel)`

SetCredentials sets Credentials field to given value.


### GetSshSettings

`func (o *LinuxHostImportSpec) GetSshSettings() LinuxHostSSHSettingsModel`

GetSshSettings returns the SshSettings field if non-nil, zero value otherwise.

### GetSshSettingsOk

`func (o *LinuxHostImportSpec) GetSshSettingsOk() (*LinuxHostSSHSettingsModel, bool)`

GetSshSettingsOk returns a tuple with the SshSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshSettings

`func (o *LinuxHostImportSpec) SetSshSettings(v LinuxHostSSHSettingsModel)`

SetSshSettings sets SshSettings field to given value.

### HasSshSettings

`func (o *LinuxHostImportSpec) HasSshSettings() bool`

HasSshSettings returns a boolean if a field has been set.

### GetSshFingerprint

`func (o *LinuxHostImportSpec) GetSshFingerprint() string`

GetSshFingerprint returns the SshFingerprint field if non-nil, zero value otherwise.

### GetSshFingerprintOk

`func (o *LinuxHostImportSpec) GetSshFingerprintOk() (*string, bool)`

GetSshFingerprintOk returns a tuple with the SshFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshFingerprint

`func (o *LinuxHostImportSpec) SetSshFingerprint(v string)`

SetSshFingerprint sets SshFingerprint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


