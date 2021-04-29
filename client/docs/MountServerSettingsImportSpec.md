# MountServerSettingsImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MountServerName** | **string** | Name of the mount server. | 
**WriteCacheFolder** | **string** | Path to the folder used for writing cache during mount operations. | 
**VPowerNFSEnabled** | **bool** | If *true*, the vPower NFS Service is enabled on the mount server. | 
**VPowerNFSPortSettings** | Pointer to [**VPowerNFSPortSettingsModel**](VPowerNFSPortSettingsModel.md) |  | [optional] 

## Methods

### NewMountServerSettingsImportSpec

`func NewMountServerSettingsImportSpec(mountServerName string, writeCacheFolder string, vPowerNFSEnabled bool, ) *MountServerSettingsImportSpec`

NewMountServerSettingsImportSpec instantiates a new MountServerSettingsImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMountServerSettingsImportSpecWithDefaults

`func NewMountServerSettingsImportSpecWithDefaults() *MountServerSettingsImportSpec`

NewMountServerSettingsImportSpecWithDefaults instantiates a new MountServerSettingsImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMountServerName

`func (o *MountServerSettingsImportSpec) GetMountServerName() string`

GetMountServerName returns the MountServerName field if non-nil, zero value otherwise.

### GetMountServerNameOk

`func (o *MountServerSettingsImportSpec) GetMountServerNameOk() (*string, bool)`

GetMountServerNameOk returns a tuple with the MountServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServerName

`func (o *MountServerSettingsImportSpec) SetMountServerName(v string)`

SetMountServerName sets MountServerName field to given value.


### GetWriteCacheFolder

`func (o *MountServerSettingsImportSpec) GetWriteCacheFolder() string`

GetWriteCacheFolder returns the WriteCacheFolder field if non-nil, zero value otherwise.

### GetWriteCacheFolderOk

`func (o *MountServerSettingsImportSpec) GetWriteCacheFolderOk() (*string, bool)`

GetWriteCacheFolderOk returns a tuple with the WriteCacheFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteCacheFolder

`func (o *MountServerSettingsImportSpec) SetWriteCacheFolder(v string)`

SetWriteCacheFolder sets WriteCacheFolder field to given value.


### GetVPowerNFSEnabled

`func (o *MountServerSettingsImportSpec) GetVPowerNFSEnabled() bool`

GetVPowerNFSEnabled returns the VPowerNFSEnabled field if non-nil, zero value otherwise.

### GetVPowerNFSEnabledOk

`func (o *MountServerSettingsImportSpec) GetVPowerNFSEnabledOk() (*bool, bool)`

GetVPowerNFSEnabledOk returns a tuple with the VPowerNFSEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVPowerNFSEnabled

`func (o *MountServerSettingsImportSpec) SetVPowerNFSEnabled(v bool)`

SetVPowerNFSEnabled sets VPowerNFSEnabled field to given value.


### GetVPowerNFSPortSettings

`func (o *MountServerSettingsImportSpec) GetVPowerNFSPortSettings() VPowerNFSPortSettingsModel`

GetVPowerNFSPortSettings returns the VPowerNFSPortSettings field if non-nil, zero value otherwise.

### GetVPowerNFSPortSettingsOk

`func (o *MountServerSettingsImportSpec) GetVPowerNFSPortSettingsOk() (*VPowerNFSPortSettingsModel, bool)`

GetVPowerNFSPortSettingsOk returns a tuple with the VPowerNFSPortSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVPowerNFSPortSettings

`func (o *MountServerSettingsImportSpec) SetVPowerNFSPortSettings(v VPowerNFSPortSettingsModel)`

SetVPowerNFSPortSettings sets VPowerNFSPortSettings field to given value.

### HasVPowerNFSPortSettings

`func (o *MountServerSettingsImportSpec) HasVPowerNFSPortSettings() bool`

HasVPowerNFSPortSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


