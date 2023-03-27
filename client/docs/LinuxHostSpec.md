# LinuxHostSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshSettings** | Pointer to [**LinuxHostSSHSettingsModel**](LinuxHostSSHSettingsModel.md) |  | [optional] 
**SshFingerprint** | **string** | SSH key fingerprint used to verify the server identity. For details on how to get the fingerprint, see [Request TLS Certificate or SSH Fingerprint](#tag/Connection/operation/GetConnectionCertificate). | 

## Methods

### NewLinuxHostSpec

`func NewLinuxHostSpec(sshFingerprint string, ) *LinuxHostSpec`

NewLinuxHostSpec instantiates a new LinuxHostSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinuxHostSpecWithDefaults

`func NewLinuxHostSpecWithDefaults() *LinuxHostSpec`

NewLinuxHostSpecWithDefaults instantiates a new LinuxHostSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSshSettings

`func (o *LinuxHostSpec) GetSshSettings() LinuxHostSSHSettingsModel`

GetSshSettings returns the SshSettings field if non-nil, zero value otherwise.

### GetSshSettingsOk

`func (o *LinuxHostSpec) GetSshSettingsOk() (*LinuxHostSSHSettingsModel, bool)`

GetSshSettingsOk returns a tuple with the SshSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshSettings

`func (o *LinuxHostSpec) SetSshSettings(v LinuxHostSSHSettingsModel)`

SetSshSettings sets SshSettings field to given value.

### HasSshSettings

`func (o *LinuxHostSpec) HasSshSettings() bool`

HasSshSettings returns a boolean if a field has been set.

### GetSshFingerprint

`func (o *LinuxHostSpec) GetSshFingerprint() string`

GetSshFingerprint returns the SshFingerprint field if non-nil, zero value otherwise.

### GetSshFingerprintOk

`func (o *LinuxHostSpec) GetSshFingerprintOk() (*string, bool)`

GetSshFingerprintOk returns a tuple with the SshFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshFingerprint

`func (o *LinuxHostSpec) SetSshFingerprint(v string)`

SetSshFingerprint sets SshFingerprint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


