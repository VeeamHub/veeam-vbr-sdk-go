# SmbStorageModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Share** | [**SmbRepositoryShareSettingsModel**](SmbRepositoryShareSettingsModel.md) |  | 
**Repository** | [**NetworkRepositorySettingsModel**](NetworkRepositorySettingsModel.md) |  | 
**MountServer** | [**MountServerSettingsModel**](MountServerSettingsModel.md) |  | 

## Methods

### NewSmbStorageModel

`func NewSmbStorageModel(share SmbRepositoryShareSettingsModel, repository NetworkRepositorySettingsModel, mountServer MountServerSettingsModel, ) *SmbStorageModel`

NewSmbStorageModel instantiates a new SmbStorageModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmbStorageModelWithDefaults

`func NewSmbStorageModelWithDefaults() *SmbStorageModel`

NewSmbStorageModelWithDefaults instantiates a new SmbStorageModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShare

`func (o *SmbStorageModel) GetShare() SmbRepositoryShareSettingsModel`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *SmbStorageModel) GetShareOk() (*SmbRepositoryShareSettingsModel, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *SmbStorageModel) SetShare(v SmbRepositoryShareSettingsModel)`

SetShare sets Share field to given value.


### GetRepository

`func (o *SmbStorageModel) GetRepository() NetworkRepositorySettingsModel`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *SmbStorageModel) GetRepositoryOk() (*NetworkRepositorySettingsModel, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *SmbStorageModel) SetRepository(v NetworkRepositorySettingsModel)`

SetRepository sets Repository field to given value.


### GetMountServer

`func (o *SmbStorageModel) GetMountServer() MountServerSettingsModel`

GetMountServer returns the MountServer field if non-nil, zero value otherwise.

### GetMountServerOk

`func (o *SmbStorageModel) GetMountServerOk() (*MountServerSettingsModel, bool)`

GetMountServerOk returns a tuple with the MountServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountServer

`func (o *SmbStorageModel) SetMountServer(v MountServerSettingsModel)`

SetMountServer sets MountServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


