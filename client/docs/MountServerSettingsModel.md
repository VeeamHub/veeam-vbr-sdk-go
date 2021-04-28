# MountServerSettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MountServerId** | **string** | ID of the mount server. | 
**WriteCacheFolder** | **string** | Path to the folder used for writing cache during mount operations. | 
**VPowerNFSEnabled** | **bool** | If *true*, the vPower NFS Service is enabled on the mount server. | 
**VPowerNFSPortSettings** | Pointer to [**VPowerNFSPortSettingsModel**](VPowerNFSPortSettingsModel.md) |  | [optional] 

## Methods

### NewMountServerSettingsModel

`func NewMountServerSettingsModel(mountServerId string, writeCacheFolder string, vPowerNFSEnabled bool, ) *MountServerSettingsModel`

NewMountServerSettingsModel instantiates a new MountServerSettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMountServerSettingsModelWithDefaults

`func NewMountServerSettingsModelWithDefaults() *MountServerSettingsModel`

NewMountServerSettingsModelWithDefaults instantiates a new MountServerSettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMountServerId

`func (o *MountServerSettingsModel) GetMountServerId() string`

GetMountServerId returns the MountServerId field if non-nil, zero value otherwise.

### GetMountServerIdOk

`func (o *MountServerSettingsModel) GetMountServerIdOk() (*string, bool)`

GetMountServerIdOk returns a tuple with the MountServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServerId

`func (o *MountServerSettingsModel) SetMountServerId(v string)`

SetMountServerId sets MountServerId field to given value.


### GetWriteCacheFolder

`func (o *MountServerSettingsModel) GetWriteCacheFolder() string`

GetWriteCacheFolder returns the WriteCacheFolder field if non-nil, zero value otherwise.

### GetWriteCacheFolderOk

`func (o *MountServerSettingsModel) GetWriteCacheFolderOk() (*string, bool)`

GetWriteCacheFolderOk returns a tuple with the WriteCacheFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteCacheFolder

`func (o *MountServerSettingsModel) SetWriteCacheFolder(v string)`

SetWriteCacheFolder sets WriteCacheFolder field to given value.


### GetVPowerNFSEnabled

`func (o *MountServerSettingsModel) GetVPowerNFSEnabled() bool`

GetVPowerNFSEnabled returns the VPowerNFSEnabled field if non-nil, zero value otherwise.

### GetVPowerNFSEnabledOk

`func (o *MountServerSettingsModel) GetVPowerNFSEnabledOk() (*bool, bool)`

GetVPowerNFSEnabledOk returns a tuple with the VPowerNFSEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVPowerNFSEnabled

`func (o *MountServerSettingsModel) SetVPowerNFSEnabled(v bool)`

SetVPowerNFSEnabled sets VPowerNFSEnabled field to given value.


### GetVPowerNFSPortSettings

`func (o *MountServerSettingsModel) GetVPowerNFSPortSettings() VPowerNFSPortSettingsModel`

GetVPowerNFSPortSettings returns the VPowerNFSPortSettings field if non-nil, zero value otherwise.

### GetVPowerNFSPortSettingsOk

`func (o *MountServerSettingsModel) GetVPowerNFSPortSettingsOk() (*VPowerNFSPortSettingsModel, bool)`

GetVPowerNFSPortSettingsOk returns a tuple with the VPowerNFSPortSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVPowerNFSPortSettings

`func (o *MountServerSettingsModel) SetVPowerNFSPortSettings(v VPowerNFSPortSettingsModel)`

SetVPowerNFSPortSettings sets VPowerNFSPortSettings field to given value.

### HasVPowerNFSPortSettings

`func (o *MountServerSettingsModel) HasVPowerNFSPortSettings() bool`

HasVPowerNFSPortSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


